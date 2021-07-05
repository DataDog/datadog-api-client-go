# WidgetLayout

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Height** | **int64** | The height of the widget. Should be a non-negative integer. | 
**IsColumnBreak** | Pointer to **bool** | Whether the widget should be the first one on the second column in high density or not. **Note**: Only for the **new dashboard layout** and only one widget in the dashboard should have this property set to &#x60;true&#x60;. | [optional] 
**Width** | **int64** | The width of the widget. Should be a non-negative integer. | 
**X** | **int64** | The position of the widget on the x (horizontal) axis. Should be a non-negative integer. | 
**Y** | **int64** | The position of the widget on the y (vertical) axis. Should be a non-negative integer. | 

## Methods

### NewWidgetLayout

`func NewWidgetLayout(height int64, width int64, x int64, y int64) *WidgetLayout`

NewWidgetLayout instantiates a new WidgetLayout object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewWidgetLayoutWithDefaults

`func NewWidgetLayoutWithDefaults() *WidgetLayout`

NewWidgetLayoutWithDefaults instantiates a new WidgetLayout object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetHeight

`func (o *WidgetLayout) GetHeight() int64`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *WidgetLayout) GetHeightOk() (*int64, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *WidgetLayout) SetHeight(v int64)`

SetHeight sets Height field to given value.


### GetIsColumnBreak

`func (o *WidgetLayout) GetIsColumnBreak() bool`

GetIsColumnBreak returns the IsColumnBreak field if non-nil, zero value otherwise.

### GetIsColumnBreakOk

`func (o *WidgetLayout) GetIsColumnBreakOk() (*bool, bool)`

GetIsColumnBreakOk returns a tuple with the IsColumnBreak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsColumnBreak

`func (o *WidgetLayout) SetIsColumnBreak(v bool)`

SetIsColumnBreak sets IsColumnBreak field to given value.

### HasIsColumnBreak

`func (o *WidgetLayout) HasIsColumnBreak() bool`

HasIsColumnBreak returns a boolean if a field has been set.

### GetWidth

`func (o *WidgetLayout) GetWidth() int64`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *WidgetLayout) GetWidthOk() (*int64, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *WidgetLayout) SetWidth(v int64)`

SetWidth sets Width field to given value.


### GetX

`func (o *WidgetLayout) GetX() int64`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *WidgetLayout) GetXOk() (*int64, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *WidgetLayout) SetX(v int64)`

SetX sets X field to given value.


### GetY

`func (o *WidgetLayout) GetY() int64`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *WidgetLayout) GetYOk() (*int64, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *WidgetLayout) SetY(v int64)`

SetY sets Y field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


