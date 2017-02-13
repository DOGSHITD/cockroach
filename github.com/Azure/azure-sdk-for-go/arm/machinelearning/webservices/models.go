package webservices

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator 1.0.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"net/http"
)

// AssetType enumerates the values for asset type.
type AssetType string

const (
	// AssetTypeModule specifies the asset type module state for asset type.
	AssetTypeModule AssetType = "Module"
	// AssetTypeResource specifies the asset type resource state for asset
	// type.
	AssetTypeResource AssetType = "Resource"
)

// ColumnFormat enumerates the values for column format.
type ColumnFormat string

const (
	// Byte specifies the byte state for column format.
	Byte ColumnFormat = "Byte"
	// Char specifies the char state for column format.
	Char ColumnFormat = "Char"
	// Complex128 specifies the complex 128 state for column format.
	Complex128 ColumnFormat = "Complex128"
	// Complex64 specifies the complex 64 state for column format.
	Complex64 ColumnFormat = "Complex64"
	// DateTime specifies the date time state for column format.
	DateTime ColumnFormat = "Date-time"
	// DateTimeOffset specifies the date time offset state for column format.
	DateTimeOffset ColumnFormat = "Date-timeOffset"
	// Double specifies the double state for column format.
	Double ColumnFormat = "Double"
	// Duration specifies the duration state for column format.
	Duration ColumnFormat = "Duration"
	// Float specifies the float state for column format.
	Float ColumnFormat = "Float"
	// Int16 specifies the int 16 state for column format.
	Int16 ColumnFormat = "Int16"
	// Int32 specifies the int 32 state for column format.
	Int32 ColumnFormat = "Int32"
	// Int64 specifies the int 64 state for column format.
	Int64 ColumnFormat = "Int64"
	// Int8 specifies the int 8 state for column format.
	Int8 ColumnFormat = "Int8"
	// Uint16 specifies the uint 16 state for column format.
	Uint16 ColumnFormat = "Uint16"
	// Uint32 specifies the uint 32 state for column format.
	Uint32 ColumnFormat = "Uint32"
	// Uint64 specifies the uint 64 state for column format.
	Uint64 ColumnFormat = "Uint64"
	// Uint8 specifies the uint 8 state for column format.
	Uint8 ColumnFormat = "Uint8"
)

// ColumnType enumerates the values for column type.
type ColumnType string

const (
	// Boolean specifies the boolean state for column type.
	Boolean ColumnType = "Boolean"
	// Integer specifies the integer state for column type.
	Integer ColumnType = "Integer"
	// Number specifies the number state for column type.
	Number ColumnType = "Number"
	// String specifies the string state for column type.
	String ColumnType = "String"
)

// DiagnosticsLevel enumerates the values for diagnostics level.
type DiagnosticsLevel string

const (
	// All specifies the all state for diagnostics level.
	All DiagnosticsLevel = "All"
	// Error specifies the error state for diagnostics level.
	Error DiagnosticsLevel = "Error"
	// None specifies the none state for diagnostics level.
	None DiagnosticsLevel = "None"
)

// InputPortType enumerates the values for input port type.
type InputPortType string

const (
	// Dataset specifies the dataset state for input port type.
	Dataset InputPortType = "Dataset"
)

// OutputPortType enumerates the values for output port type.
type OutputPortType string

const (
	// OutputPortTypeDataset specifies the output port type dataset state for
	// output port type.
	OutputPortTypeDataset OutputPortType = "Dataset"
)

// ParameterType enumerates the values for parameter type.
type ParameterType string

const (
	// ParameterTypeBoolean specifies the parameter type boolean state for
	// parameter type.
	ParameterTypeBoolean ParameterType = "Boolean"
	// ParameterTypeColumnPicker specifies the parameter type column picker
	// state for parameter type.
	ParameterTypeColumnPicker ParameterType = "ColumnPicker"
	// ParameterTypeCredential specifies the parameter type credential state
	// for parameter type.
	ParameterTypeCredential ParameterType = "Credential"
	// ParameterTypeDataGatewayName specifies the parameter type data gateway
	// name state for parameter type.
	ParameterTypeDataGatewayName ParameterType = "DataGatewayName"
	// ParameterTypeDouble specifies the parameter type double state for
	// parameter type.
	ParameterTypeDouble ParameterType = "Double"
	// ParameterTypeEnumerated specifies the parameter type enumerated state
	// for parameter type.
	ParameterTypeEnumerated ParameterType = "Enumerated"
	// ParameterTypeFloat specifies the parameter type float state for
	// parameter type.
	ParameterTypeFloat ParameterType = "Float"
	// ParameterTypeInt specifies the parameter type int state for parameter
	// type.
	ParameterTypeInt ParameterType = "Int"
	// ParameterTypeMode specifies the parameter type mode state for parameter
	// type.
	ParameterTypeMode ParameterType = "Mode"
	// ParameterTypeParameterRange specifies the parameter type parameter range
	// state for parameter type.
	ParameterTypeParameterRange ParameterType = "ParameterRange"
	// ParameterTypeScript specifies the parameter type script state for
	// parameter type.
	ParameterTypeScript ParameterType = "Script"
	// ParameterTypeString specifies the parameter type string state for
	// parameter type.
	ParameterTypeString ParameterType = "String"
)

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// Failed specifies the failed state for provisioning state.
	Failed ProvisioningState = "Failed"
	// Provisioning specifies the provisioning state for provisioning state.
	Provisioning ProvisioningState = "Provisioning"
	// Succeeded specifies the succeeded state for provisioning state.
	Succeeded ProvisioningState = "Succeeded"
	// Unknown specifies the unknown state for provisioning state.
	Unknown ProvisioningState = "Unknown"
)

// AssetItem is information about an asset associated with the web service.
type AssetItem struct {
	Name         *string                 `json:"name,omitempty"`
	ID           *string                 `json:"id,omitempty"`
	Type         AssetType               `json:"type,omitempty"`
	LocationInfo *AssetLocation          `json:"locationInfo,omitempty"`
	InputPorts   *map[string]*InputPort  `json:"inputPorts,omitempty"`
	OutputPorts  *map[string]*OutputPort `json:"outputPorts,omitempty"`
	Metadata     *map[string]*string     `json:"metadata,omitempty"`
	Parameters   *[]ModuleAssetParameter `json:"parameters,omitempty"`
}

// AssetLocation is describes the access location for a web service asset.
type AssetLocation struct {
	URI         *string `json:"uri,omitempty"`
	Credentials *string `json:"credentials,omitempty"`
}

// ColumnSpecification is swagger 2.0 schema for a column within the data table
// representing a web service input or output. See Swagger specification:
// http://swagger.io/specification/
type ColumnSpecification struct {
	Type          ColumnType                `json:"type,omitempty"`
	Format        ColumnFormat              `json:"format,omitempty"`
	Enum          *[]map[string]interface{} `json:"enum,omitempty"`
	XMsIsnullable *bool                     `json:"x-ms-isnullable,omitempty"`
	XMsIsordered  *bool                     `json:"x-ms-isordered,omitempty"`
}

// CommitmentPlan is information about the machine learning commitment plan
// associated with the web service.
type CommitmentPlan struct {
	ID *string `json:"id,omitempty"`
}

// DiagnosticsConfiguration is diagnostics settings for an Azure ML web
// service.
type DiagnosticsConfiguration struct {
	Level  DiagnosticsLevel `json:"level,omitempty"`
	Expiry *date.Time       `json:"expiry,omitempty"`
}

// ExampleRequest is sample input data for the service's input(s).
type ExampleRequest struct {
	Inputs           *map[string][][]map[string]interface{} `json:"inputs,omitempty"`
	GlobalParameters *map[string]*map[string]interface{}    `json:"globalParameters,omitempty"`
}

// GraphEdge is defines an edge within the web service's graph.
type GraphEdge struct {
	SourceNodeID *string `json:"sourceNodeId,omitempty"`
	SourcePortID *string `json:"sourcePortId,omitempty"`
	TargetNodeID *string `json:"targetNodeId,omitempty"`
	TargetPortID *string `json:"targetPortId,omitempty"`
}

// GraphNode is specifies a node in the web service graph. The node can either
// be an input, output or asset node, so only one of the corresponding id
// properties is populated at any given time.
type GraphNode struct {
	AssetID    *string             `json:"assetId,omitempty"`
	InputID    *string             `json:"inputId,omitempty"`
	OutputID   *string             `json:"outputId,omitempty"`
	Parameters *map[string]*string `json:"parameters,omitempty"`
}

// GraphPackage is defines the graph of modules making up the machine learning
// solution.
type GraphPackage struct {
	Nodes           *map[string]*GraphNode      `json:"nodes,omitempty"`
	Edges           *[]GraphEdge                `json:"edges,omitempty"`
	GraphParameters *map[string]*GraphParameter `json:"graphParameters,omitempty"`
}

// GraphParameter is defines a global parameter in the graph.
type GraphParameter struct {
	Description *string               `json:"description,omitempty"`
	Type        ParameterType         `json:"type,omitempty"`
	Links       *[]GraphParameterLink `json:"links,omitempty"`
}

// GraphParameterLink is association link for a graph global parameter to a
// node in the graph.
type GraphParameterLink struct {
	NodeID       *string `json:"nodeId,omitempty"`
	ParameterKey *string `json:"parameterKey,omitempty"`
}

// InputPort is asset input port
type InputPort struct {
	Type InputPortType `json:"type,omitempty"`
}

// Keys is access keys for the web service calls.
type Keys struct {
	autorest.Response `json:"-"`
	Primary           *string `json:"primary,omitempty"`
	Secondary         *string `json:"secondary,omitempty"`
}

// MachineLearningWorkspace is information about the machine learning workspace
// containing the experiment that is source for the web service.
type MachineLearningWorkspace struct {
	ID *string `json:"id,omitempty"`
}

// ModeValueInfo is nested parameter definition.
type ModeValueInfo struct {
	InterfaceString *string                 `json:"interfaceString,omitempty"`
	Parameters      *[]ModuleAssetParameter `json:"parameters,omitempty"`
}

// ModuleAssetParameter is parameter definition for a module asset.
type ModuleAssetParameter struct {
	Name           *string                    `json:"name,omitempty"`
	ParameterType  *string                    `json:"parameterType,omitempty"`
	ModeValuesInfo *map[string]*ModeValueInfo `json:"modeValuesInfo,omitempty"`
}

// OutputPort is asset output port
type OutputPort struct {
	Type OutputPortType `json:"type,omitempty"`
}

// PaginatedWebServicesList is paginated list of web services.
type PaginatedWebServicesList struct {
	autorest.Response `json:"-"`
	Value             *[]WebService `json:"value,omitempty"`
	NextLink          *string       `json:"nextLink,omitempty"`
}

// PaginatedWebServicesListPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client PaginatedWebServicesList) PaginatedWebServicesListPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// Properties is the set of properties specific to the Azure ML web service
// resource.
type Properties struct {
	Title                    *string                          `json:"title,omitempty"`
	Description              *string                          `json:"description,omitempty"`
	CreatedOn                *date.Time                       `json:"createdOn,omitempty"`
	ModifiedOn               *date.Time                       `json:"modifiedOn,omitempty"`
	ProvisioningState        ProvisioningState                `json:"provisioningState,omitempty"`
	Keys                     *Keys                            `json:"keys,omitempty"`
	ReadOnly                 *bool                            `json:"readOnly,omitempty"`
	SwaggerLocation          *string                          `json:"swaggerLocation,omitempty"`
	ExposeSampleData         *bool                            `json:"exposeSampleData,omitempty"`
	RealtimeConfiguration    *RealtimeConfiguration           `json:"realtimeConfiguration,omitempty"`
	Diagnostics              *DiagnosticsConfiguration        `json:"diagnostics,omitempty"`
	StorageAccount           *StorageAccount                  `json:"storageAccount,omitempty"`
	MachineLearningWorkspace *MachineLearningWorkspace        `json:"machineLearningWorkspace,omitempty"`
	CommitmentPlan           *CommitmentPlan                  `json:"commitmentPlan,omitempty"`
	Input                    *ServiceInputOutputSpecification `json:"input,omitempty"`
	Output                   *ServiceInputOutputSpecification `json:"output,omitempty"`
	ExampleRequest           *ExampleRequest                  `json:"exampleRequest,omitempty"`
	Assets                   *map[string]*AssetItem           `json:"assets,omitempty"`
	Parameters               *map[string]*string              `json:"parameters,omitempty"`
}

// PropertiesForGraph is properties specific to a Graph based web service.
type PropertiesForGraph struct {
	Title                    *string                          `json:"title,omitempty"`
	Description              *string                          `json:"description,omitempty"`
	CreatedOn                *date.Time                       `json:"createdOn,omitempty"`
	ModifiedOn               *date.Time                       `json:"modifiedOn,omitempty"`
	ProvisioningState        ProvisioningState                `json:"provisioningState,omitempty"`
	Keys                     *Keys                            `json:"keys,omitempty"`
	ReadOnly                 *bool                            `json:"readOnly,omitempty"`
	SwaggerLocation          *string                          `json:"swaggerLocation,omitempty"`
	ExposeSampleData         *bool                            `json:"exposeSampleData,omitempty"`
	RealtimeConfiguration    *RealtimeConfiguration           `json:"realtimeConfiguration,omitempty"`
	Diagnostics              *DiagnosticsConfiguration        `json:"diagnostics,omitempty"`
	StorageAccount           *StorageAccount                  `json:"storageAccount,omitempty"`
	MachineLearningWorkspace *MachineLearningWorkspace        `json:"machineLearningWorkspace,omitempty"`
	CommitmentPlan           *CommitmentPlan                  `json:"commitmentPlan,omitempty"`
	Input                    *ServiceInputOutputSpecification `json:"input,omitempty"`
	Output                   *ServiceInputOutputSpecification `json:"output,omitempty"`
	ExampleRequest           *ExampleRequest                  `json:"exampleRequest,omitempty"`
	Assets                   *map[string]*AssetItem           `json:"assets,omitempty"`
	Parameters               *map[string]*string              `json:"parameters,omitempty"`
	Package                  *GraphPackage                    `json:"package,omitempty"`
}

// RealtimeConfiguration is holds the available configuration options for an
// Azure ML web service endpoint.
type RealtimeConfiguration struct {
	MaxConcurrentCalls *int32 `json:"maxConcurrentCalls,omitempty"`
}

// Resource is
type Resource struct {
	ID       *string             `json:"id,omitempty"`
	Name     *string             `json:"name,omitempty"`
	Location *string             `json:"location,omitempty"`
	Type     *string             `json:"type,omitempty"`
	Tags     *map[string]*string `json:"tags,omitempty"`
}

// ServiceInputOutputSpecification is the swagger 2.0 schema describing the
// service's inputs or outputs. See Swagger specification:
// http://swagger.io/specification/
type ServiceInputOutputSpecification struct {
	Title       *string                         `json:"title,omitempty"`
	Description *string                         `json:"description,omitempty"`
	Type        *string                         `json:"type,omitempty"`
	Properties  *map[string]*TableSpecification `json:"properties,omitempty"`
}

// StorageAccount is access information for a storage account.
type StorageAccount struct {
	Name *string `json:"name,omitempty"`
	Key  *string `json:"key,omitempty"`
}

// TableSpecification is the swagger 2.0 schema describing a single service
// input or output. See Swagger specification: http://swagger.io/specification/
type TableSpecification struct {
	Title       *string                          `json:"title,omitempty"`
	Description *string                          `json:"description,omitempty"`
	Type        *string                          `json:"type,omitempty"`
	Format      *string                          `json:"format,omitempty"`
	Properties  *map[string]*ColumnSpecification `json:"properties,omitempty"`
}

// WebService is instance of an Azure ML web service resource.
type WebService struct {
	autorest.Response `json:"-"`
	ID                *string             `json:"id,omitempty"`
	Name              *string             `json:"name,omitempty"`
	Location          *string             `json:"location,omitempty"`
	Type              *string             `json:"type,omitempty"`
	Tags              *map[string]*string `json:"tags,omitempty"`
	Properties        *Properties         `json:"properties,omitempty"`
}
