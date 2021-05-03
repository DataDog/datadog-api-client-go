# FormulaAndFunctionEventQueryGroupBy

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Facet** | **string** | Event facet. | 
**Limit** | Pointer to **int64** | Number of groups to return. | [optional] 
**Sort** | Pointer to [**FormulaAndFunctionEventQueryGroupBySort**](FormulaAndFunctionEventQueryGroupBySort.md) |  | [optional] 

## Methods

### NewFormulaAndFunctionEventQueryGroupBy

`func NewFormulaAndFunctionEventQueryGroupBy(facet string, ) *FormulaAndFunctionEventQueryGroupBy`

NewFormulaAndFunctionEventQueryGroupBy instantiates a new FormulaAndFunctionEventQueryGroupBy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormulaAndFunctionEventQueryGroupByWithDefaults

`func NewFormulaAndFunctionEventQueryGroupByWithDefaults() *FormulaAndFunctionEventQueryGroupBy`

NewFormulaAndFunctionEventQueryGroupByWithDefaults instantiates a new FormulaAndFunctionEventQueryGroupBy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFacet

`func (o *FormulaAndFunctionEventQueryGroupBy) GetFacet() string`

GetFacet returns the Facet field if non-nil, zero value otherwise.

### GetFacetOk

`func (o *FormulaAndFunctionEventQueryGroupBy) GetFacetOk() (*string, bool)`

GetFacetOk returns a tuple with the Facet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacet

`func (o *FormulaAndFunctionEventQueryGroupBy) SetFacet(v string)`

SetFacet sets Facet field to given value.


### GetLimit

`func (o *FormulaAndFunctionEventQueryGroupBy) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *FormulaAndFunctionEventQueryGroupBy) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *FormulaAndFunctionEventQueryGroupBy) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *FormulaAndFunctionEventQueryGroupBy) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetSort

`func (o *FormulaAndFunctionEventQueryGroupBy) GetSort() FormulaAndFunctionEventQueryGroupBySort`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *FormulaAndFunctionEventQueryGroupBy) GetSortOk() (*FormulaAndFunctionEventQueryGroupBySort, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *FormulaAndFunctionEventQueryGroupBy) SetSort(v FormulaAndFunctionEventQueryGroupBySort)`

SetSort sets Sort field to given value.

### HasSort

`func (o *FormulaAndFunctionEventQueryGroupBy) HasSort() bool`

HasSort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


