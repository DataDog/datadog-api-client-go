# WidgetMarker

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**DisplayType** | Pointer to **string** | Combination of:   - A severity error, warning, ok, or info   - A line type: dashed, solid, or bold In this case of a Distribution widget, this can be set to be &#x60;x_axis_percentile&#x60;.  | [optional] 
**Label** | Pointer to **string** | Label to display over the marker. | [optional] 
**Time** | Pointer to **string** | Timestamp for the widget. | [optional] 
**Value** | **string** | Value to apply. Can be a single value y &#x3D; 15 or a range of values 0 &lt; y &lt; 10. | 

## Methods

### NewWidgetMarker

`func NewWidgetMarker(value string, ) *WidgetMarker`

NewWidgetMarker instantiates a new WidgetMarker object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWidgetMarkerWithDefaults

`func NewWidgetMarkerWithDefaults() *WidgetMarker`

NewWidgetMarkerWithDefaults instantiates a new WidgetMarker object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayType

`func (o *WidgetMarker) GetDisplayType() string`

GetDisplayType returns the DisplayType field if non-nil, zero value otherwise.

### GetDisplayTypeOk

`func (o *WidgetMarker) GetDisplayTypeOk() (*string, bool)`

GetDisplayTypeOk returns a tuple with the DisplayType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayType

`func (o *WidgetMarker) SetDisplayType(v string)`

SetDisplayType sets DisplayType field to given value.

### HasDisplayType

`func (o *WidgetMarker) HasDisplayType() bool`

HasDisplayType returns a boolean if a field has been set.

### GetLabel

`func (o *WidgetMarker) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *WidgetMarker) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *WidgetMarker) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *WidgetMarker) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetTime

`func (o *WidgetMarker) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *WidgetMarker) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *WidgetMarker) SetTime(v string)`

SetTime sets Time field to given value.

### HasTime

`func (o *WidgetMarker) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetValue

`func (o *WidgetMarker) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *WidgetMarker) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *WidgetMarker) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


