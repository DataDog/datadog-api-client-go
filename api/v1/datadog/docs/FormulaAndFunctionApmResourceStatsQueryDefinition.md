# FormulaAndFunctionApmResourceStatsQueryDefinition

## Properties

| Name                | Type                                                                                                | Description                                                                                                                                                                                                               | Notes      |
| ------------------- | --------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------- |
| **DataSource**      | [**FormulaAndFunctionApmResourceStatsDataSource**](FormulaAndFunctionApmResourceStatsDataSource.md) |                                                                                                                                                                                                                           |
| **Env**             | **string**                                                                                          | APM environment.                                                                                                                                                                                                          |
| **GroupBy**         | Pointer to **[]string**                                                                             | Array of fields to group results by.                                                                                                                                                                                      | [optional] |
| **Name**            | **string**                                                                                          | Name of this query to use in formulas.                                                                                                                                                                                    |
| **OperationName**   | Pointer to **string**                                                                               | Name of operation on service.                                                                                                                                                                                             | [optional] |
| **PrimaryTagName**  | Pointer to **string**                                                                               | Name of the second primary tag used within APM. Required when &#x60;primary_tag_value&#x60; is specified. See https://docs.datadoghq.com/tracing/guide/setting_primary_tags_to_scope/#add-a-second-primary-tag-in-datadog | [optional] |
| **PrimaryTagValue** | Pointer to **string**                                                                               | Value of the second primary tag by which to filter APM data. &#x60;primary_tag_name&#x60; must also be specified.                                                                                                         | [optional] |
| **ResourceName**    | Pointer to **string**                                                                               | APM resource name.                                                                                                                                                                                                        | [optional] |
| **Service**         | **string**                                                                                          | APM service name.                                                                                                                                                                                                         |
| **Stat**            | [**FormulaAndFunctionApmResourceStatName**](FormulaAndFunctionApmResourceStatName.md)               |                                                                                                                                                                                                                           |

## Methods

### NewFormulaAndFunctionApmResourceStatsQueryDefinition

`func NewFormulaAndFunctionApmResourceStatsQueryDefinition(dataSource FormulaAndFunctionApmResourceStatsDataSource, env string, name string, service string, stat FormulaAndFunctionApmResourceStatName) *FormulaAndFunctionApmResourceStatsQueryDefinition`

NewFormulaAndFunctionApmResourceStatsQueryDefinition instantiates a new FormulaAndFunctionApmResourceStatsQueryDefinition object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewFormulaAndFunctionApmResourceStatsQueryDefinitionWithDefaults

`func NewFormulaAndFunctionApmResourceStatsQueryDefinitionWithDefaults() *FormulaAndFunctionApmResourceStatsQueryDefinition`

NewFormulaAndFunctionApmResourceStatsQueryDefinitionWithDefaults instantiates a new FormulaAndFunctionApmResourceStatsQueryDefinition object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetDataSource

`func (o *FormulaAndFunctionApmResourceStatsQueryDefinition) GetDataSource() FormulaAndFunctionApmResourceStatsDataSource`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *FormulaAndFunctionApmResourceStatsQueryDefinition) GetDataSourceOk() (*FormulaAndFunctionApmResourceStatsDataSource, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *FormulaAndFunctionApmResourceStatsQueryDefinition) SetDataSource(v FormulaAndFunctionApmResourceStatsDataSource)`

SetDataSource sets DataSource field to given value.

### GetEnv

`func (o *FormulaAndFunctionApmResourceStatsQueryDefinition) GetEnv() string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *FormulaAndFunctionApmResourceStatsQueryDefinition) GetEnvOk() (*string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *FormulaAndFunctionApmResourceStatsQueryDefinition) SetEnv(v string)`

SetEnv sets Env field to given value.

### GetGroupBy

`func (o *FormulaAndFunctionApmResourceStatsQueryDefinition) GetGroupBy() []string`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *FormulaAndFunctionApmResourceStatsQueryDefinition) GetGroupByOk() (*[]string, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *FormulaAndFunctionApmResourceStatsQueryDefinition) SetGroupBy(v []string)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *FormulaAndFunctionApmResourceStatsQueryDefinition) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetName

`func (o *FormulaAndFunctionApmResourceStatsQueryDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FormulaAndFunctionApmResourceStatsQueryDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FormulaAndFunctionApmResourceStatsQueryDefinition) SetName(v string)`

SetName sets Name field to given value.

### GetOperationName

`func (o *FormulaAndFunctionApmResourceStatsQueryDefinition) GetOperationName() string`

GetOperationName returns the OperationName field if non-nil, zero value otherwise.

### GetOperationNameOk

`func (o *FormulaAndFunctionApmResourceStatsQueryDefinition) GetOperationNameOk() (*string, bool)`

GetOperationNameOk returns a tuple with the OperationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationName

`func (o *FormulaAndFunctionApmResourceStatsQueryDefinition) SetOperationName(v string)`

SetOperationName sets OperationName field to given value.

### HasOperationName

`func (o *FormulaAndFunctionApmResourceStatsQueryDefinition) HasOperationName() bool`

HasOperationName returns a boolean if a field has been set.

### GetPrimaryTagName

`func (o *FormulaAndFunctionApmResourceStatsQueryDefinition) GetPrimaryTagName() string`

GetPrimaryTagName returns the PrimaryTagName field if non-nil, zero value otherwise.

### GetPrimaryTagNameOk

`func (o *FormulaAndFunctionApmResourceStatsQueryDefinition) GetPrimaryTagNameOk() (*string, bool)`

GetPrimaryTagNameOk returns a tuple with the PrimaryTagName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryTagName

`func (o *FormulaAndFunctionApmResourceStatsQueryDefinition) SetPrimaryTagName(v string)`

SetPrimaryTagName sets PrimaryTagName field to given value.

### HasPrimaryTagName

`func (o *FormulaAndFunctionApmResourceStatsQueryDefinition) HasPrimaryTagName() bool`

HasPrimaryTagName returns a boolean if a field has been set.

### GetPrimaryTagValue

`func (o *FormulaAndFunctionApmResourceStatsQueryDefinition) GetPrimaryTagValue() string`

GetPrimaryTagValue returns the PrimaryTagValue field if non-nil, zero value otherwise.

### GetPrimaryTagValueOk

`func (o *FormulaAndFunctionApmResourceStatsQueryDefinition) GetPrimaryTagValueOk() (*string, bool)`

GetPrimaryTagValueOk returns a tuple with the PrimaryTagValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryTagValue

`func (o *FormulaAndFunctionApmResourceStatsQueryDefinition) SetPrimaryTagValue(v string)`

SetPrimaryTagValue sets PrimaryTagValue field to given value.

### HasPrimaryTagValue

`func (o *FormulaAndFunctionApmResourceStatsQueryDefinition) HasPrimaryTagValue() bool`

HasPrimaryTagValue returns a boolean if a field has been set.

### GetResourceName

`func (o *FormulaAndFunctionApmResourceStatsQueryDefinition) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *FormulaAndFunctionApmResourceStatsQueryDefinition) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *FormulaAndFunctionApmResourceStatsQueryDefinition) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *FormulaAndFunctionApmResourceStatsQueryDefinition) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### GetService

`func (o *FormulaAndFunctionApmResourceStatsQueryDefinition) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *FormulaAndFunctionApmResourceStatsQueryDefinition) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *FormulaAndFunctionApmResourceStatsQueryDefinition) SetService(v string)`

SetService sets Service field to given value.

### GetStat

`func (o *FormulaAndFunctionApmResourceStatsQueryDefinition) GetStat() FormulaAndFunctionApmResourceStatName`

GetStat returns the Stat field if non-nil, zero value otherwise.

### GetStatOk

`func (o *FormulaAndFunctionApmResourceStatsQueryDefinition) GetStatOk() (*FormulaAndFunctionApmResourceStatName, bool)`

GetStatOk returns a tuple with the Stat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStat

`func (o *FormulaAndFunctionApmResourceStatsQueryDefinition) SetStat(v FormulaAndFunctionApmResourceStatName)`

SetStat sets Stat field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
