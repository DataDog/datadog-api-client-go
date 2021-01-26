# TimeSeriesFormulaAndFunctionEventQueryDefinitionGroupBy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Facet** | **string** | Event facet. | 
**Limit** | Pointer to **int64** | Number of groups to return. | [optional] 
**Sort** | Pointer to [**TimeSeriesFormulaAndFunctionEventQueryDefinitionGroupBySort**](TimeSeriesFormulaAndFunctionEventQueryDefinition_group_by_sort.md) |  | [optional] 

## Methods

### NewTimeSeriesFormulaAndFunctionEventQueryDefinitionGroupBy

`func NewTimeSeriesFormulaAndFunctionEventQueryDefinitionGroupBy(facet string, ) *TimeSeriesFormulaAndFunctionEventQueryDefinitionGroupBy`

NewTimeSeriesFormulaAndFunctionEventQueryDefinitionGroupBy instantiates a new TimeSeriesFormulaAndFunctionEventQueryDefinitionGroupBy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeSeriesFormulaAndFunctionEventQueryDefinitionGroupByWithDefaults

`func NewTimeSeriesFormulaAndFunctionEventQueryDefinitionGroupByWithDefaults() *TimeSeriesFormulaAndFunctionEventQueryDefinitionGroupBy`

NewTimeSeriesFormulaAndFunctionEventQueryDefinitionGroupByWithDefaults instantiates a new TimeSeriesFormulaAndFunctionEventQueryDefinitionGroupBy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFacet

`func (o *TimeSeriesFormulaAndFunctionEventQueryDefinitionGroupBy) GetFacet() string`

GetFacet returns the Facet field if non-nil, zero value otherwise.

### GetFacetOk

`func (o *TimeSeriesFormulaAndFunctionEventQueryDefinitionGroupBy) GetFacetOk() (*string, bool)`

GetFacetOk returns a tuple with the Facet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacet

`func (o *TimeSeriesFormulaAndFunctionEventQueryDefinitionGroupBy) SetFacet(v string)`

SetFacet sets Facet field to given value.


### GetLimit

`func (o *TimeSeriesFormulaAndFunctionEventQueryDefinitionGroupBy) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *TimeSeriesFormulaAndFunctionEventQueryDefinitionGroupBy) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *TimeSeriesFormulaAndFunctionEventQueryDefinitionGroupBy) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *TimeSeriesFormulaAndFunctionEventQueryDefinitionGroupBy) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetSort

`func (o *TimeSeriesFormulaAndFunctionEventQueryDefinitionGroupBy) GetSort() TimeSeriesFormulaAndFunctionEventQueryDefinitionGroupBySort`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *TimeSeriesFormulaAndFunctionEventQueryDefinitionGroupBy) GetSortOk() (*TimeSeriesFormulaAndFunctionEventQueryDefinitionGroupBySort, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *TimeSeriesFormulaAndFunctionEventQueryDefinitionGroupBy) SetSort(v TimeSeriesFormulaAndFunctionEventQueryDefinitionGroupBySort)`

SetSort sets Sort field to given value.

### HasSort

`func (o *TimeSeriesFormulaAndFunctionEventQueryDefinitionGroupBy) HasSort() bool`

HasSort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


