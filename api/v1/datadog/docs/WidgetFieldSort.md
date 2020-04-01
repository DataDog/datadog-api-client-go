# WidgetFieldSort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Column** | Pointer to **string** | Facet path for the column | 
**Order** | Pointer to [**WidgetSort**](WidgetSort.md) |  | 

## Methods

### NewWidgetFieldSort

`func NewWidgetFieldSort(column string, order WidgetSort, ) *WidgetFieldSort`

NewWidgetFieldSort instantiates a new WidgetFieldSort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWidgetFieldSortWithDefaults

`func NewWidgetFieldSortWithDefaults() *WidgetFieldSort`

NewWidgetFieldSortWithDefaults instantiates a new WidgetFieldSort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumn

`func (o *WidgetFieldSort) GetColumn() string`

GetColumn returns the Column field if non-nil, zero value otherwise.

### GetColumnOk

`func (o *WidgetFieldSort) GetColumnOk() (*string, bool)`

GetColumnOk returns a tuple with the Column field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumn

`func (o *WidgetFieldSort) SetColumn(v string)`

SetColumn sets Column field to given value.


### GetOrder

`func (o *WidgetFieldSort) GetOrder() WidgetSort`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *WidgetFieldSort) GetOrderOk() (*WidgetSort, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *WidgetFieldSort) SetOrder(v WidgetSort)`

SetOrder sets Order field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


