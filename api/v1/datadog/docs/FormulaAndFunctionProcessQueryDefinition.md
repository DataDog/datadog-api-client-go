# FormulaAndFunctionProcessQueryDefinition

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Aggregator** | Pointer to [**FormulaAndFunctionMetricAggregation**](FormulaAndFunctionMetricAggregation.md) |  | [optional] 
**DataSource** | [**FormulaAndFunctionProcessQueryDataSource**](FormulaAndFunctionProcessQueryDataSource.md) |  | 
**IsNormalizedCpu** | Pointer to **bool** | Whether to normalize the CPU percentages. | [optional] 
**Limit** | Pointer to **int64** | Number of hits to return. | [optional] 
**Metric** | **string** | Process metric name. | 
**Name** | **string** | Name of query for use in formulas. | 
**Sort** | Pointer to [**QuerySortOrder**](QuerySortOrder.md) |  | [optional] [default to QUERYSORTORDER_DESC]
**TagFilters** | Pointer to **[]string** | An array of tags to filter by. | [optional] 
**TextFilter** | Pointer to **string** | Text to use as filter. | [optional] 

## Methods

### NewFormulaAndFunctionProcessQueryDefinition

`func NewFormulaAndFunctionProcessQueryDefinition(dataSource FormulaAndFunctionProcessQueryDataSource, metric string, name string) *FormulaAndFunctionProcessQueryDefinition`

NewFormulaAndFunctionProcessQueryDefinition instantiates a new FormulaAndFunctionProcessQueryDefinition object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewFormulaAndFunctionProcessQueryDefinitionWithDefaults

`func NewFormulaAndFunctionProcessQueryDefinitionWithDefaults() *FormulaAndFunctionProcessQueryDefinition`

NewFormulaAndFunctionProcessQueryDefinitionWithDefaults instantiates a new FormulaAndFunctionProcessQueryDefinition object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAggregator

`func (o *FormulaAndFunctionProcessQueryDefinition) GetAggregator() FormulaAndFunctionMetricAggregation`

GetAggregator returns the Aggregator field if non-nil, zero value otherwise.

### GetAggregatorOk

`func (o *FormulaAndFunctionProcessQueryDefinition) GetAggregatorOk() (*FormulaAndFunctionMetricAggregation, bool)`

GetAggregatorOk returns a tuple with the Aggregator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregator

`func (o *FormulaAndFunctionProcessQueryDefinition) SetAggregator(v FormulaAndFunctionMetricAggregation)`

SetAggregator sets Aggregator field to given value.

### HasAggregator

`func (o *FormulaAndFunctionProcessQueryDefinition) HasAggregator() bool`

HasAggregator returns a boolean if a field has been set.

### GetDataSource

`func (o *FormulaAndFunctionProcessQueryDefinition) GetDataSource() FormulaAndFunctionProcessQueryDataSource`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *FormulaAndFunctionProcessQueryDefinition) GetDataSourceOk() (*FormulaAndFunctionProcessQueryDataSource, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *FormulaAndFunctionProcessQueryDefinition) SetDataSource(v FormulaAndFunctionProcessQueryDataSource)`

SetDataSource sets DataSource field to given value.


### GetIsNormalizedCpu

`func (o *FormulaAndFunctionProcessQueryDefinition) GetIsNormalizedCpu() bool`

GetIsNormalizedCpu returns the IsNormalizedCpu field if non-nil, zero value otherwise.

### GetIsNormalizedCpuOk

`func (o *FormulaAndFunctionProcessQueryDefinition) GetIsNormalizedCpuOk() (*bool, bool)`

GetIsNormalizedCpuOk returns a tuple with the IsNormalizedCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNormalizedCpu

`func (o *FormulaAndFunctionProcessQueryDefinition) SetIsNormalizedCpu(v bool)`

SetIsNormalizedCpu sets IsNormalizedCpu field to given value.

### HasIsNormalizedCpu

`func (o *FormulaAndFunctionProcessQueryDefinition) HasIsNormalizedCpu() bool`

HasIsNormalizedCpu returns a boolean if a field has been set.

### GetLimit

`func (o *FormulaAndFunctionProcessQueryDefinition) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *FormulaAndFunctionProcessQueryDefinition) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *FormulaAndFunctionProcessQueryDefinition) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *FormulaAndFunctionProcessQueryDefinition) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetMetric

`func (o *FormulaAndFunctionProcessQueryDefinition) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *FormulaAndFunctionProcessQueryDefinition) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *FormulaAndFunctionProcessQueryDefinition) SetMetric(v string)`

SetMetric sets Metric field to given value.


### GetName

`func (o *FormulaAndFunctionProcessQueryDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FormulaAndFunctionProcessQueryDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FormulaAndFunctionProcessQueryDefinition) SetName(v string)`

SetName sets Name field to given value.


### GetSort

`func (o *FormulaAndFunctionProcessQueryDefinition) GetSort() QuerySortOrder`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *FormulaAndFunctionProcessQueryDefinition) GetSortOk() (*QuerySortOrder, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *FormulaAndFunctionProcessQueryDefinition) SetSort(v QuerySortOrder)`

SetSort sets Sort field to given value.

### HasSort

`func (o *FormulaAndFunctionProcessQueryDefinition) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetTagFilters

`func (o *FormulaAndFunctionProcessQueryDefinition) GetTagFilters() []string`

GetTagFilters returns the TagFilters field if non-nil, zero value otherwise.

### GetTagFiltersOk

`func (o *FormulaAndFunctionProcessQueryDefinition) GetTagFiltersOk() (*[]string, bool)`

GetTagFiltersOk returns a tuple with the TagFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagFilters

`func (o *FormulaAndFunctionProcessQueryDefinition) SetTagFilters(v []string)`

SetTagFilters sets TagFilters field to given value.

### HasTagFilters

`func (o *FormulaAndFunctionProcessQueryDefinition) HasTagFilters() bool`

HasTagFilters returns a boolean if a field has been set.

### GetTextFilter

`func (o *FormulaAndFunctionProcessQueryDefinition) GetTextFilter() string`

GetTextFilter returns the TextFilter field if non-nil, zero value otherwise.

### GetTextFilterOk

`func (o *FormulaAndFunctionProcessQueryDefinition) GetTextFilterOk() (*string, bool)`

GetTextFilterOk returns a tuple with the TextFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextFilter

`func (o *FormulaAndFunctionProcessQueryDefinition) SetTextFilter(v string)`

SetTextFilter sets TextFilter field to given value.

### HasTextFilter

`func (o *FormulaAndFunctionProcessQueryDefinition) HasTextFilter() bool`

HasTextFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


