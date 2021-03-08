# WidgetFormulaLimit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int64** | Number of results to return. | [optional] 
**Order** | Pointer to [**QuerySortOrder**](QuerySortOrder.md) |  | [optional] [default to QUERYSORTORDER_DESC]

## Methods

### NewWidgetFormulaLimit

`func NewWidgetFormulaLimit() *WidgetFormulaLimit`

NewWidgetFormulaLimit instantiates a new WidgetFormulaLimit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWidgetFormulaLimitWithDefaults

`func NewWidgetFormulaLimitWithDefaults() *WidgetFormulaLimit`

NewWidgetFormulaLimitWithDefaults instantiates a new WidgetFormulaLimit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *WidgetFormulaLimit) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *WidgetFormulaLimit) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *WidgetFormulaLimit) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *WidgetFormulaLimit) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetOrder

`func (o *WidgetFormulaLimit) GetOrder() QuerySortOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *WidgetFormulaLimit) GetOrderOk() (*QuerySortOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *WidgetFormulaLimit) SetOrder(v QuerySortOrder)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *WidgetFormulaLimit) HasOrder() bool`

HasOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


