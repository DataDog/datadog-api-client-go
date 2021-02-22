# TimeSeriesFormulaAndFunctionEventQueryGroupBy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Facet** | **string** | Event facet. | 
**Limit** | Pointer to **int64** | Number of groups to return. | [optional] 
**Sort** | Pointer to [**TimeSeriesFormulaAndFunctionEventQueryGroupBySort**](TimeSeriesFormulaAndFunctionEventQueryGroupBy_sort.md) |  | [optional] 

## Methods

### NewTimeSeriesFormulaAndFunctionEventQueryGroupBy

`func NewTimeSeriesFormulaAndFunctionEventQueryGroupBy(facet string, ) *TimeSeriesFormulaAndFunctionEventQueryGroupBy`

NewTimeSeriesFormulaAndFunctionEventQueryGroupBy instantiates a new TimeSeriesFormulaAndFunctionEventQueryGroupBy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeSeriesFormulaAndFunctionEventQueryGroupByWithDefaults

`func NewTimeSeriesFormulaAndFunctionEventQueryGroupByWithDefaults() *TimeSeriesFormulaAndFunctionEventQueryGroupBy`

NewTimeSeriesFormulaAndFunctionEventQueryGroupByWithDefaults instantiates a new TimeSeriesFormulaAndFunctionEventQueryGroupBy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFacet

`func (o *TimeSeriesFormulaAndFunctionEventQueryGroupBy) GetFacet() string`

GetFacet returns the Facet field if non-nil, zero value otherwise.

### GetFacetOk

`func (o *TimeSeriesFormulaAndFunctionEventQueryGroupBy) GetFacetOk() (*string, bool)`

GetFacetOk returns a tuple with the Facet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacet

`func (o *TimeSeriesFormulaAndFunctionEventQueryGroupBy) SetFacet(v string)`

SetFacet sets Facet field to given value.


### GetLimit

`func (o *TimeSeriesFormulaAndFunctionEventQueryGroupBy) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *TimeSeriesFormulaAndFunctionEventQueryGroupBy) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *TimeSeriesFormulaAndFunctionEventQueryGroupBy) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *TimeSeriesFormulaAndFunctionEventQueryGroupBy) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetSort

`func (o *TimeSeriesFormulaAndFunctionEventQueryGroupBy) GetSort() TimeSeriesFormulaAndFunctionEventQueryGroupBySort`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *TimeSeriesFormulaAndFunctionEventQueryGroupBy) GetSortOk() (*TimeSeriesFormulaAndFunctionEventQueryGroupBySort, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *TimeSeriesFormulaAndFunctionEventQueryGroupBy) SetSort(v TimeSeriesFormulaAndFunctionEventQueryGroupBySort)`

SetSort sets Sort field to given value.

### HasSort

`func (o *TimeSeriesFormulaAndFunctionEventQueryGroupBy) HasSort() bool`

HasSort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


