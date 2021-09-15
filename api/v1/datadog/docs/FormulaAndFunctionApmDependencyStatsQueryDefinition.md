# FormulaAndFunctionApmDependencyStatsQueryDefinition

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**DataSource** | [**FormulaAndFunctionApmDependencyStatsDataSource**](FormulaAndFunctionApmDependencyStatsDataSource.md) |  | 
**Env** | **string** | APM environment. | 
**IsUpstream** | Pointer to **bool** | Determines whether stats for upstream or downstream dependencies should be queried. | [optional] 
**Name** | **string** | Name of query to use in formulas. | 
**OperationName** | **string** | Name of operation on service. | 
**PrimaryTagName** | Pointer to **string** | The name of the second primary tag used within APM; required when &#x60;primary_tag_value&#x60; is specified. See https://docs.datadoghq.com/tracing/guide/setting_primary_tags_to_scope/#add-a-second-primary-tag-in-datadog. | [optional] 
**PrimaryTagValue** | Pointer to **string** | Filter APM data by the second primary tag. &#x60;primary_tag_name&#x60; must also be specified. | [optional] 
**ResourceName** | **string** | APM resource. | 
**Service** | **string** | APM service. | 
**Stat** | [**FormulaAndFunctionApmDependencyStatName**](FormulaAndFunctionApmDependencyStatName.md) |  | 

## Methods

### NewFormulaAndFunctionApmDependencyStatsQueryDefinition

`func NewFormulaAndFunctionApmDependencyStatsQueryDefinition(dataSource FormulaAndFunctionApmDependencyStatsDataSource, env string, name string, operationName string, resourceName string, service string, stat FormulaAndFunctionApmDependencyStatName) *FormulaAndFunctionApmDependencyStatsQueryDefinition`

NewFormulaAndFunctionApmDependencyStatsQueryDefinition instantiates a new FormulaAndFunctionApmDependencyStatsQueryDefinition object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewFormulaAndFunctionApmDependencyStatsQueryDefinitionWithDefaults

`func NewFormulaAndFunctionApmDependencyStatsQueryDefinitionWithDefaults() *FormulaAndFunctionApmDependencyStatsQueryDefinition`

NewFormulaAndFunctionApmDependencyStatsQueryDefinitionWithDefaults instantiates a new FormulaAndFunctionApmDependencyStatsQueryDefinition object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetDataSource

`func (o *FormulaAndFunctionApmDependencyStatsQueryDefinition) GetDataSource() FormulaAndFunctionApmDependencyStatsDataSource`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *FormulaAndFunctionApmDependencyStatsQueryDefinition) GetDataSourceOk() (*FormulaAndFunctionApmDependencyStatsDataSource, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *FormulaAndFunctionApmDependencyStatsQueryDefinition) SetDataSource(v FormulaAndFunctionApmDependencyStatsDataSource)`

SetDataSource sets DataSource field to given value.


### GetEnv

`func (o *FormulaAndFunctionApmDependencyStatsQueryDefinition) GetEnv() string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *FormulaAndFunctionApmDependencyStatsQueryDefinition) GetEnvOk() (*string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *FormulaAndFunctionApmDependencyStatsQueryDefinition) SetEnv(v string)`

SetEnv sets Env field to given value.


### GetIsUpstream

`func (o *FormulaAndFunctionApmDependencyStatsQueryDefinition) GetIsUpstream() bool`

GetIsUpstream returns the IsUpstream field if non-nil, zero value otherwise.

### GetIsUpstreamOk

`func (o *FormulaAndFunctionApmDependencyStatsQueryDefinition) GetIsUpstreamOk() (*bool, bool)`

GetIsUpstreamOk returns a tuple with the IsUpstream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUpstream

`func (o *FormulaAndFunctionApmDependencyStatsQueryDefinition) SetIsUpstream(v bool)`

SetIsUpstream sets IsUpstream field to given value.

### HasIsUpstream

`func (o *FormulaAndFunctionApmDependencyStatsQueryDefinition) HasIsUpstream() bool`

HasIsUpstream returns a boolean if a field has been set.

### GetName

`func (o *FormulaAndFunctionApmDependencyStatsQueryDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FormulaAndFunctionApmDependencyStatsQueryDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FormulaAndFunctionApmDependencyStatsQueryDefinition) SetName(v string)`

SetName sets Name field to given value.


### GetOperationName

`func (o *FormulaAndFunctionApmDependencyStatsQueryDefinition) GetOperationName() string`

GetOperationName returns the OperationName field if non-nil, zero value otherwise.

### GetOperationNameOk

`func (o *FormulaAndFunctionApmDependencyStatsQueryDefinition) GetOperationNameOk() (*string, bool)`

GetOperationNameOk returns a tuple with the OperationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationName

`func (o *FormulaAndFunctionApmDependencyStatsQueryDefinition) SetOperationName(v string)`

SetOperationName sets OperationName field to given value.


### GetPrimaryTagName

`func (o *FormulaAndFunctionApmDependencyStatsQueryDefinition) GetPrimaryTagName() string`

GetPrimaryTagName returns the PrimaryTagName field if non-nil, zero value otherwise.

### GetPrimaryTagNameOk

`func (o *FormulaAndFunctionApmDependencyStatsQueryDefinition) GetPrimaryTagNameOk() (*string, bool)`

GetPrimaryTagNameOk returns a tuple with the PrimaryTagName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryTagName

`func (o *FormulaAndFunctionApmDependencyStatsQueryDefinition) SetPrimaryTagName(v string)`

SetPrimaryTagName sets PrimaryTagName field to given value.

### HasPrimaryTagName

`func (o *FormulaAndFunctionApmDependencyStatsQueryDefinition) HasPrimaryTagName() bool`

HasPrimaryTagName returns a boolean if a field has been set.

### GetPrimaryTagValue

`func (o *FormulaAndFunctionApmDependencyStatsQueryDefinition) GetPrimaryTagValue() string`

GetPrimaryTagValue returns the PrimaryTagValue field if non-nil, zero value otherwise.

### GetPrimaryTagValueOk

`func (o *FormulaAndFunctionApmDependencyStatsQueryDefinition) GetPrimaryTagValueOk() (*string, bool)`

GetPrimaryTagValueOk returns a tuple with the PrimaryTagValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryTagValue

`func (o *FormulaAndFunctionApmDependencyStatsQueryDefinition) SetPrimaryTagValue(v string)`

SetPrimaryTagValue sets PrimaryTagValue field to given value.

### HasPrimaryTagValue

`func (o *FormulaAndFunctionApmDependencyStatsQueryDefinition) HasPrimaryTagValue() bool`

HasPrimaryTagValue returns a boolean if a field has been set.

### GetResourceName

`func (o *FormulaAndFunctionApmDependencyStatsQueryDefinition) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *FormulaAndFunctionApmDependencyStatsQueryDefinition) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *FormulaAndFunctionApmDependencyStatsQueryDefinition) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.


### GetService

`func (o *FormulaAndFunctionApmDependencyStatsQueryDefinition) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *FormulaAndFunctionApmDependencyStatsQueryDefinition) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *FormulaAndFunctionApmDependencyStatsQueryDefinition) SetService(v string)`

SetService sets Service field to given value.


### GetStat

`func (o *FormulaAndFunctionApmDependencyStatsQueryDefinition) GetStat() FormulaAndFunctionApmDependencyStatName`

GetStat returns the Stat field if non-nil, zero value otherwise.

### GetStatOk

`func (o *FormulaAndFunctionApmDependencyStatsQueryDefinition) GetStatOk() (*FormulaAndFunctionApmDependencyStatName, bool)`

GetStatOk returns a tuple with the Stat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStat

`func (o *FormulaAndFunctionApmDependencyStatsQueryDefinition) SetStat(v FormulaAndFunctionApmDependencyStatName)`

SetStat sets Stat field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


