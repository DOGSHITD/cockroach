// Copyright 2016 The Cockroach Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.

package distsqlrun

import (
	"context"
	"math"
	"sort"
	"strings"
	"testing"

	"github.com/cockroachdb/cockroach/pkg/settings/cluster"
	"github.com/cockroachdb/cockroach/pkg/sql/sem/tree"
	"github.com/cockroachdb/cockroach/pkg/sql/sqlbase"
	"github.com/cockroachdb/cockroach/pkg/util"
	"github.com/cockroachdb/cockroach/pkg/util/leaktest"
)

// TODO(irfansharif): Add tests to verify the following aggregation functions:
//      AVG
//      BOOL_AND
//      BOOL_OR
//      CONCAT_AGG
//      JSON_AGG
//      JSONB_AGG
//      STDDEV
//      VARIANCE
func TestAggregator(t *testing.T) {
	defer leaktest.AfterTest(t)()

	v := [15]sqlbase.EncDatum{}
	null := sqlbase.EncDatum{Datum: tree.DNull}
	for i := range v {
		v[i] = sqlbase.DatumToEncDatum(intType, tree.NewDInt(tree.DInt(i)))
	}

	boolTrue := sqlbase.DatumToEncDatum(boolType, tree.DBoolTrue)
	boolFalse := sqlbase.DatumToEncDatum(boolType, tree.DBoolFalse)
	boolNULL := sqlbase.DatumToEncDatum(boolType, tree.DNull)

	colPtr := func(idx uint32) *uint32 { return &idx }

	testCases := []struct {
		spec        AggregatorSpec
		inputTypes  []sqlbase.ColumnType
		input       sqlbase.EncDatumRows
		outputTypes []sqlbase.ColumnType
		expected    sqlbase.EncDatumRows
	}{
		{
			// SELECT min(@0), max(@0), count(@0), avg(@0), sum(@0), stddev(@0),
			// variance(@0) GROUP BY [] (no rows).
			spec: AggregatorSpec{
				Aggregations: []AggregatorSpec_Aggregation{
					{
						Func:   AggregatorSpec_MIN,
						ColIdx: []uint32{0},
					},
					{
						Func:   AggregatorSpec_MAX,
						ColIdx: []uint32{0},
					},
					{
						Func:   AggregatorSpec_COUNT,
						ColIdx: []uint32{0},
					},
					{
						Func:   AggregatorSpec_AVG,
						ColIdx: []uint32{0},
					},
					{
						Func:   AggregatorSpec_SUM,
						ColIdx: []uint32{0},
					},
					{
						Func:   AggregatorSpec_STDDEV,
						ColIdx: []uint32{0},
					},
					{
						Func:   AggregatorSpec_VARIANCE,
						ColIdx: []uint32{0},
					},
				},
			},
			inputTypes: oneIntCol,
			input:      sqlbase.EncDatumRows{},
			outputTypes: []sqlbase.ColumnType{
				intType, // MIN
				intType, // MAX
				intType, // COUNT
				decType, // AVG
				decType, // SUM
				decType, // STDDEV
				decType, // VARIANCE
			},
			expected: sqlbase.EncDatumRows{
				{null, null, v[0], null, null, null, null},
			},
		},
		{
			// SELECT @2, count(@1), GROUP BY @2.
			spec: AggregatorSpec{
				GroupCols: []uint32{1},
				Aggregations: []AggregatorSpec_Aggregation{
					{
						Func:   AggregatorSpec_ANY_NOT_NULL,
						ColIdx: []uint32{1},
					},
					{
						Func:   AggregatorSpec_COUNT,
						ColIdx: []uint32{0},
					},
				},
			},
			inputTypes: twoIntCols,
			input: sqlbase.EncDatumRows{
				{v[1], v[2]},
				{v[3], null},
				{v[6], v[2]},
				{v[7], v[2]},
				{v[8], v[4]},
			},
			outputTypes: twoIntCols,
			expected: sqlbase.EncDatumRows{
				{null, v[1]},
				{v[4], v[1]},
				{v[2], v[3]},
			},
		},
		{
			// SELECT @2, count(@1), GROUP BY @2.
			spec: AggregatorSpec{
				GroupCols: []uint32{1},
				Aggregations: []AggregatorSpec_Aggregation{
					{
						Func:   AggregatorSpec_ANY_NOT_NULL,
						ColIdx: []uint32{1},
					},
					{
						Func:   AggregatorSpec_COUNT,
						ColIdx: []uint32{0},
					},
				},
			},
			inputTypes: twoIntCols,
			input: sqlbase.EncDatumRows{
				{v[1], v[2]},
				{v[3], v[4]},
				{v[6], v[2]},
				{v[7], v[2]},
				{v[8], v[4]},
			},
			outputTypes: twoIntCols,
			expected: sqlbase.EncDatumRows{
				{v[4], v[2]},
				{v[2], v[3]},
			},
		},
		{
			// SELECT @2, count(@1), GROUP BY @2 (ordering: @2+).
			spec: AggregatorSpec{
				GroupCols:        []uint32{1},
				OrderedGroupCols: []uint32{1},
				Aggregations: []AggregatorSpec_Aggregation{
					{
						Func:   AggregatorSpec_ANY_NOT_NULL,
						ColIdx: []uint32{1},
					},
					{
						Func:   AggregatorSpec_COUNT,
						ColIdx: []uint32{0},
					},
				},
			},
			inputTypes: twoIntCols,
			input: sqlbase.EncDatumRows{
				{v[1], v[2]},
				{v[6], v[2]},
				{v[7], v[2]},
				{v[3], v[4]},
				{v[8], v[4]},
			},
			outputTypes: twoIntCols,
			expected: sqlbase.EncDatumRows{
				{v[2], v[3]},
				{v[4], v[2]},
			},
		},
		{
			// SELECT @2, sum(@1), GROUP BY @2.
			spec: AggregatorSpec{
				GroupCols: []uint32{1},
				Aggregations: []AggregatorSpec_Aggregation{
					{
						Func:   AggregatorSpec_ANY_NOT_NULL,
						ColIdx: []uint32{1},
					},
					{
						Func:   AggregatorSpec_SUM,
						ColIdx: []uint32{0},
					},
				},
			},
			inputTypes: twoIntCols,
			input: sqlbase.EncDatumRows{
				{v[1], v[2]},
				{v[3], v[4]},
				{v[6], v[2]},
				{v[7], v[2]},
				{v[8], v[4]},
			},
			outputTypes: []sqlbase.ColumnType{
				intType, // ANY_NOT_NULL
				decType, // SUM
			},
			expected: sqlbase.EncDatumRows{
				{v[2], v[14]},
				{v[4], v[11]},
			},
		},
		{
			// SELECT @2, sum(@1), GROUP BY @2 (ordering: @2+).
			spec: AggregatorSpec{
				GroupCols:        []uint32{1},
				OrderedGroupCols: []uint32{1},
				Aggregations: []AggregatorSpec_Aggregation{
					{
						Func:   AggregatorSpec_ANY_NOT_NULL,
						ColIdx: []uint32{1},
					},
					{
						Func:   AggregatorSpec_SUM,
						ColIdx: []uint32{0},
					},
				},
			},
			inputTypes: twoIntCols,
			input: sqlbase.EncDatumRows{
				{v[1], v[2]},
				{v[6], v[2]},
				{v[7], v[2]},
				{v[8], v[4]},
				{v[3], v[4]},
			},
			outputTypes: []sqlbase.ColumnType{
				intType, // ANY_NOT_NULL
				decType, // SUM
			},
			expected: sqlbase.EncDatumRows{
				{v[2], v[14]},
				{v[4], v[11]},
			},
		},
		{
			// SELECT count(@1), sum(@1), GROUP BY [] (empty group key).
			spec: AggregatorSpec{
				Aggregations: []AggregatorSpec_Aggregation{
					{
						Func:   AggregatorSpec_COUNT,
						ColIdx: []uint32{0},
					},
					{
						Func:   AggregatorSpec_SUM,
						ColIdx: []uint32{0},
					},
				},
			},
			inputTypes: twoIntCols,
			input: sqlbase.EncDatumRows{
				{v[1], v[2]},
				{v[1], v[4]},
				{v[3], v[2]},
				{v[4], v[2]},
				{v[5], v[4]},
			},
			outputTypes: []sqlbase.ColumnType{
				intType, // COUNT
				decType, // SUM
			},
			expected: sqlbase.EncDatumRows{
				{v[5], v[14]},
			},
		},
		{
			// SELECT SUM DISTINCT (@1), GROUP BY [] (empty group key).
			spec: AggregatorSpec{
				Aggregations: []AggregatorSpec_Aggregation{
					{
						Func:     AggregatorSpec_SUM,
						Distinct: true,
						ColIdx:   []uint32{0},
					},
				},
			},
			inputTypes: oneIntCol,
			input: sqlbase.EncDatumRows{
				{v[2]},
				{v[4]},
				{v[2]},
				{v[2]},
				{v[4]},
			},
			outputTypes: []sqlbase.ColumnType{decType},
			expected: sqlbase.EncDatumRows{
				{v[6]},
			},
		},
		{
			// SELECT @1, GROUP BY [] (empty group key).
			spec: AggregatorSpec{
				Aggregations: []AggregatorSpec_Aggregation{
					{
						Func:   AggregatorSpec_ANY_NOT_NULL,
						ColIdx: []uint32{0},
					},
				},
			},
			inputTypes: oneIntCol,
			input: sqlbase.EncDatumRows{
				{v[1]},
				{v[1]},
				{v[1]},
			},
			outputTypes: oneIntCol,
			expected: sqlbase.EncDatumRows{
				{v[1]},
			},
		},
		{
			// SELECT max(@1), min(@2), count(@2), COUNT DISTINCT (@2), GROUP BY [] (empty group key).
			spec: AggregatorSpec{
				Aggregations: []AggregatorSpec_Aggregation{
					{
						Func:   AggregatorSpec_MAX,
						ColIdx: []uint32{0},
					},
					{
						Func:   AggregatorSpec_MIN,
						ColIdx: []uint32{1},
					},
					{
						Func:   AggregatorSpec_COUNT,
						ColIdx: []uint32{1},
					},
					{
						Func:     AggregatorSpec_COUNT,
						Distinct: true,
						ColIdx:   []uint32{1},
					},
				},
			},
			inputTypes: twoIntCols,
			input: sqlbase.EncDatumRows{
				{v[2], v[2]},
				{v[1], v[4]},
				{v[3], v[2]},
				{v[4], v[2]},
				{v[5], v[4]},
			},
			outputTypes: []sqlbase.ColumnType{intType, intType, intType, intType},
			expected: sqlbase.EncDatumRows{
				{v[5], v[2], v[5], v[2]},
			},
		},
		{
			// SELECT max(@1) FILTER @2, count(@3) FILTER @4, COUNT_ROWS FILTER @4
			spec: AggregatorSpec{
				Aggregations: []AggregatorSpec_Aggregation{
					{
						Func:         AggregatorSpec_MAX,
						ColIdx:       []uint32{0},
						FilterColIdx: colPtr(1),
					},
					{
						Func:         AggregatorSpec_COUNT,
						ColIdx:       []uint32{2},
						FilterColIdx: colPtr(3),
					},
					{
						Func:         AggregatorSpec_COUNT_ROWS,
						FilterColIdx: colPtr(3),
					},
				},
			},
			inputTypes: []sqlbase.ColumnType{intType, boolType, intType, boolType},
			input: sqlbase.EncDatumRows{
				{v[1], boolTrue, v[1], boolTrue},
				{v[5], boolFalse, v[1], boolFalse},
				{v[2], boolTrue, v[1], boolNULL},
				{v[3], boolNULL, v[1], boolTrue},
				{v[2], boolTrue, v[1], boolTrue},
			},
			outputTypes: threeIntCols,
			expected: sqlbase.EncDatumRows{
				{v[2], v[3], v[3]},
			},
		},
	}

	for _, c := range testCases {
		t.Run("", func(t *testing.T) {
			ags := c.spec

			in := NewRowBuffer(c.inputTypes, c.input, RowBufferArgs{})
			out := NewRowBuffer(c.outputTypes, nil /* rows */, RowBufferArgs{})
			st := cluster.MakeTestingClusterSettings()
			evalCtx := tree.MakeTestingEvalContext(st)
			defer evalCtx.Stop(context.Background())
			flowCtx := FlowCtx{
				Settings: st,
				EvalCtx:  &evalCtx,
			}

			ag, err := newAggregator(&flowCtx, 0 /* processorID */, &ags, in, &PostProcessSpec{}, out)
			if err != nil {
				t.Fatal(err)
			}

			ag.Run(context.Background(), nil /* wg */)

			var expected []string
			for _, row := range c.expected {
				expected = append(expected, row.String(c.outputTypes))
			}
			sort.Strings(expected)
			expStr := strings.Join(expected, "")

			var rets []string
			for {
				row := out.NextNoMeta(t)
				if row == nil {
					break
				}
				rets = append(rets, row.String(c.outputTypes))
			}
			sort.Strings(rets)
			retStr := strings.Join(rets, "")

			if expStr != retStr {
				t.Errorf("invalid results; expected:\n   %s\ngot:\n   %s",
					expStr, retStr)
			}
		})
	}
}

func BenchmarkAggregation(b *testing.B) {
	const numCols = 1
	const numRows = 1000

	aggFuncs := []AggregatorSpec_Func{
		AggregatorSpec_ANY_NOT_NULL,
		AggregatorSpec_AVG,
		AggregatorSpec_COUNT,
		AggregatorSpec_MAX,
		AggregatorSpec_MIN,
		AggregatorSpec_STDDEV,
		AggregatorSpec_SUM,
		AggregatorSpec_SUM_INT,
		AggregatorSpec_VARIANCE,
		AggregatorSpec_XOR_AGG,
	}

	ctx := context.Background()
	st := cluster.MakeTestingClusterSettings()
	evalCtx := tree.MakeTestingEvalContext(st)
	defer evalCtx.Stop(ctx)

	flowCtx := &FlowCtx{
		Settings: st,
		EvalCtx:  &evalCtx,
	}

	for _, aggFunc := range aggFuncs {
		b.Run(aggFunc.String(), func(b *testing.B) {
			spec := &AggregatorSpec{
				Aggregations: []AggregatorSpec_Aggregation{
					{
						Func:   aggFunc,
						ColIdx: []uint32{0},
					},
				},
			}
			post := &PostProcessSpec{}
			disposer := &RowDisposer{}
			input := NewRepeatableRowSource(oneIntCol, makeIntRows(numRows, numCols))

			b.SetBytes(int64(8 * numRows * numCols))
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				d, err := newAggregator(flowCtx, 0 /* processorID */, spec, input, post, disposer)
				if err != nil {
					b.Fatal(err)
				}
				d.Run(context.TODO(), nil)
				input.Reset()
			}
			b.StopTimer()
		})
	}
}

func BenchmarkGrouping(b *testing.B) {
	const numCols = 1
	const numRows = 1000

	ctx := context.Background()
	st := cluster.MakeTestingClusterSettings()
	evalCtx := tree.MakeTestingEvalContext(st)
	defer evalCtx.Stop(ctx)

	flowCtx := &FlowCtx{
		Settings: st,
		EvalCtx:  &evalCtx,
	}
	spec := &AggregatorSpec{
		GroupCols: []uint32{0},
	}
	post := &PostProcessSpec{}
	disposer := &RowDisposer{}
	input := NewRepeatableRowSource(oneIntCol, makeIntRows(numRows, numCols))

	b.SetBytes(int64(8 * numRows * numCols))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d, err := newAggregator(flowCtx, 0 /* processorID */, spec, input, post, disposer)
		if err != nil {
			b.Fatal(err)
		}
		d.Run(context.Background(), nil /* wg */)
		input.Reset()
	}
	b.StopTimer()
}

func benchmarkAggregationWithGrouping(b *testing.B, numOrderedCols int) {
	const numCols = 3
	const groupSize = 10
	var groupedCols = [2]int{0, 1}
	var allOrderedGroupCols = [2]uint32{0, 1}

	aggFuncs := []AggregatorSpec_Func{
		AggregatorSpec_ANY_NOT_NULL,
		AggregatorSpec_AVG,
		AggregatorSpec_COUNT,
		AggregatorSpec_MAX,
		AggregatorSpec_MIN,
		AggregatorSpec_STDDEV,
		AggregatorSpec_SUM,
		AggregatorSpec_SUM_INT,
		AggregatorSpec_VARIANCE,
		AggregatorSpec_XOR_AGG,
	}

	ctx := context.Background()
	st := cluster.MakeTestingClusterSettings()
	evalCtx := tree.MakeTestingEvalContext(st)
	defer evalCtx.Stop(ctx)

	flowCtx := &FlowCtx{
		Settings: st,
		EvalCtx:  &evalCtx,
	}

	for _, aggFunc := range aggFuncs {
		b.Run(aggFunc.String(), func(b *testing.B) {
			spec := &AggregatorSpec{
				GroupCols: []uint32{0, 1},
				Aggregations: []AggregatorSpec_Aggregation{
					{
						Func:   aggFunc,
						ColIdx: []uint32{2},
					},
				},
			}
			spec.OrderedGroupCols = allOrderedGroupCols[:numOrderedCols]
			post := &PostProcessSpec{}
			disposer := &RowDisposer{}
			input := NewRepeatableRowSource(threeIntCols, makeGroupedIntRows(groupSize, numCols, groupedCols[:]))

			b.SetBytes(int64(8 * intPow(groupSize, len(groupedCols)+1) * numCols))
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				d, err := newAggregator(flowCtx, 0 /* processorID */, spec, input, post, disposer)
				if err != nil {
					b.Fatal(err)
				}
				d.Run(context.Background(), nil /* wg */)
				input.Reset()
			}
			b.StopTimer()
		})
	}
}

func BenchmarkOrderedAggregation(b *testing.B) {
	benchmarkAggregationWithGrouping(b, 2 /* numOrderedCols */)
}

func BenchmarkPartiallyOrderedAggregation(b *testing.B) {
	benchmarkAggregationWithGrouping(b, 1 /* numOrderedCols */)
}

func BenchmarkUnorderedAggregation(b *testing.B) {
	benchmarkAggregationWithGrouping(b, 0 /* numOrderedCols */)
}

func intPow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

// makeGroupedIntRows constructs a (groupSize**(len(groupedCols)+1)) x numCols
// table, where columns in groupedCols are sorted in ascending order with column
// priority defined by their position in groupedCols. If used in an aggregation
// where groupedCols are the GROUP BY columns, each group will have a size of
// groupSize. To make the input more interesting for aggregation, group columns
// are repeated.
//
// Examples:
// makeGroupedIntRows(2, 2, []int{1, 0}) ->
// [0 0]
// [0 0]
// [1 0]
// [1 0]
// [0 1]
// [0 1]
// [1 1]
// [1 1]
func makeGroupedIntRows(groupSize, numCols int, groupedCols []int) sqlbase.EncDatumRows {
	numRows := intPow(groupSize, len(groupedCols)+1)
	rows := make(sqlbase.EncDatumRows, numRows)

	groupColSet := util.MakeFastIntSet(groupedCols...)
	getGroupedColVal := func(rowIdx, colIdx int) int {
		rank := -1
		for i, c := range groupedCols {
			if colIdx == c {
				rank = len(groupedCols) - i
				break
			}
		}
		if rank == -1 {
			panic("provided colIdx is not a group column")
		}
		return (rowIdx % intPow(groupSize, rank+1)) / intPow(groupSize, rank)
	}

	for i := range rows {
		rows[i] = make(sqlbase.EncDatumRow, numCols)
		for j := 0; j < numCols; j++ {
			if groupColSet.Contains(j) {
				rows[i][j] = sqlbase.DatumToEncDatum(
					intType, tree.NewDInt(tree.DInt(getGroupedColVal(i, j))))
			} else {
				rows[i][j] = sqlbase.DatumToEncDatum(intType, tree.NewDInt(tree.DInt(i+j)))
			}
		}
	}
	return rows
}