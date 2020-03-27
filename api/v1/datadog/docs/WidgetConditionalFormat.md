# WidgetConditionalFormat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comparator** | Pointer to [**WidgetComparator**](WidgetComparator.md) |  | 
**CustomBgColor** | Pointer to **string** | Color palette to apply to the background, same values available as palette. | [optional] 
**CustomFgColor** | Pointer to **string** | Color palette to apply to the foreground, same values available as palette. | [optional] 
**HideValue** | Pointer to **bool** |  | [optional] 
**ImageUrl** | Pointer to **string** | Displays an image as the background. | [optional] 
**Metric** | Pointer to **string** | Metric from the request to correlate this conditional format with | [optional] 
**Palette** | Pointer to [**WidgetPalette**](WidgetPalette.md) |  | 
**Timeframe** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **float64** | Value for the comparator. | 

## Methods

### NewWidgetConditionalFormat

`func NewWidgetConditionalFormat(comparator WidgetComparator, palette WidgetPalette, value float64, ) *WidgetConditionalFormat`

NewWidgetConditionalFormat instantiates a new WidgetConditionalFormat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWidgetConditionalFormatWithDefaults

`func NewWidgetConditionalFormatWithDefaults() *WidgetConditionalFormat`

NewWidgetConditionalFormatWithDefaults instantiates a new WidgetConditionalFormat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComparator

`func (o *WidgetConditionalFormat) GetComparator() WidgetComparator`

GetComparator returns the Comparator field if non-nil, zero value otherwise.

### GetComparatorOk

`func (o *WidgetConditionalFormat) GetComparatorOk() (*WidgetComparator, bool)`

GetComparatorOk returns a tuple with the Comparator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparator

`func (o *WidgetConditionalFormat) SetComparator(v WidgetComparator)`

SetComparator sets Comparator field to given value.


### GetCustomBgColor

`func (o *WidgetConditionalFormat) GetCustomBgColor() string`

GetCustomBgColor returns the CustomBgColor field if non-nil, zero value otherwise.

### GetCustomBgColorOk

`func (o *WidgetConditionalFormat) GetCustomBgColorOk() (*string, bool)`

GetCustomBgColorOk returns a tuple with the CustomBgColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomBgColor

`func (o *WidgetConditionalFormat) SetCustomBgColor(v string)`

SetCustomBgColor sets CustomBgColor field to given value.

### HasCustomBgColor

`func (o *WidgetConditionalFormat) HasCustomBgColor() bool`

HasCustomBgColor returns a boolean if a field has been set.

### GetCustomFgColor

`func (o *WidgetConditionalFormat) GetCustomFgColor() string`

GetCustomFgColor returns the CustomFgColor field if non-nil, zero value otherwise.

### GetCustomFgColorOk

`func (o *WidgetConditionalFormat) GetCustomFgColorOk() (*string, bool)`

GetCustomFgColorOk returns a tuple with the CustomFgColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFgColor

`func (o *WidgetConditionalFormat) SetCustomFgColor(v string)`

SetCustomFgColor sets CustomFgColor field to given value.

### HasCustomFgColor

`func (o *WidgetConditionalFormat) HasCustomFgColor() bool`

HasCustomFgColor returns a boolean if a field has been set.

### GetHideValue

`func (o *WidgetConditionalFormat) GetHideValue() bool`

GetHideValue returns the HideValue field if non-nil, zero value otherwise.

### GetHideValueOk

`func (o *WidgetConditionalFormat) GetHideValueOk() (*bool, bool)`

GetHideValueOk returns a tuple with the HideValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideValue

`func (o *WidgetConditionalFormat) SetHideValue(v bool)`

SetHideValue sets HideValue field to given value.

### HasHideValue

`func (o *WidgetConditionalFormat) HasHideValue() bool`

HasHideValue returns a boolean if a field has been set.

### GetImageUrl

`func (o *WidgetConditionalFormat) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *WidgetConditionalFormat) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *WidgetConditionalFormat) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *WidgetConditionalFormat) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetMetric

`func (o *WidgetConditionalFormat) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *WidgetConditionalFormat) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *WidgetConditionalFormat) SetMetric(v string)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *WidgetConditionalFormat) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetPalette

`func (o *WidgetConditionalFormat) GetPalette() WidgetPalette`

GetPalette returns the Palette field if non-nil, zero value otherwise.

### GetPaletteOk

`func (o *WidgetConditionalFormat) GetPaletteOk() (*WidgetPalette, bool)`

GetPaletteOk returns a tuple with the Palette field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPalette

`func (o *WidgetConditionalFormat) SetPalette(v WidgetPalette)`

SetPalette sets Palette field to given value.


### GetTimeframe

`func (o *WidgetConditionalFormat) GetTimeframe() string`

GetTimeframe returns the Timeframe field if non-nil, zero value otherwise.

### GetTimeframeOk

`func (o *WidgetConditionalFormat) GetTimeframeOk() (*string, bool)`

GetTimeframeOk returns a tuple with the Timeframe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeframe

`func (o *WidgetConditionalFormat) SetTimeframe(v string)`

SetTimeframe sets Timeframe field to given value.

### HasTimeframe

`func (o *WidgetConditionalFormat) HasTimeframe() bool`

HasTimeframe returns a boolean if a field has been set.

### GetValue

`func (o *WidgetConditionalFormat) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *WidgetConditionalFormat) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *WidgetConditionalFormat) SetValue(v float64)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


