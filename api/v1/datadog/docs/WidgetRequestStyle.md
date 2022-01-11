# WidgetRequestStyle

## Properties

| Name          | Type                                                 | Description                           | Notes      |
| ------------- | ---------------------------------------------------- | ------------------------------------- | ---------- |
| **LineType**  | Pointer to [**WidgetLineType**](WidgetLineType.md)   |                                       | [optional] |
| **LineWidth** | Pointer to [**WidgetLineWidth**](WidgetLineWidth.md) |                                       | [optional] |
| **Palette**   | Pointer to **string**                                | Color palette to apply to the widget. | [optional] |

## Methods

### NewWidgetRequestStyle

`func NewWidgetRequestStyle() *WidgetRequestStyle`

NewWidgetRequestStyle instantiates a new WidgetRequestStyle object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewWidgetRequestStyleWithDefaults

`func NewWidgetRequestStyleWithDefaults() *WidgetRequestStyle`

NewWidgetRequestStyleWithDefaults instantiates a new WidgetRequestStyle object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetLineType

`func (o *WidgetRequestStyle) GetLineType() WidgetLineType`

GetLineType returns the LineType field if non-nil, zero value otherwise.

### GetLineTypeOk

`func (o *WidgetRequestStyle) GetLineTypeOk() (*WidgetLineType, bool)`

GetLineTypeOk returns a tuple with the LineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineType

`func (o *WidgetRequestStyle) SetLineType(v WidgetLineType)`

SetLineType sets LineType field to given value.

### HasLineType

`func (o *WidgetRequestStyle) HasLineType() bool`

HasLineType returns a boolean if a field has been set.

### GetLineWidth

`func (o *WidgetRequestStyle) GetLineWidth() WidgetLineWidth`

GetLineWidth returns the LineWidth field if non-nil, zero value otherwise.

### GetLineWidthOk

`func (o *WidgetRequestStyle) GetLineWidthOk() (*WidgetLineWidth, bool)`

GetLineWidthOk returns a tuple with the LineWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineWidth

`func (o *WidgetRequestStyle) SetLineWidth(v WidgetLineWidth)`

SetLineWidth sets LineWidth field to given value.

### HasLineWidth

`func (o *WidgetRequestStyle) HasLineWidth() bool`

HasLineWidth returns a boolean if a field has been set.

### GetPalette

`func (o *WidgetRequestStyle) GetPalette() string`

GetPalette returns the Palette field if non-nil, zero value otherwise.

### GetPaletteOk

`func (o *WidgetRequestStyle) GetPaletteOk() (*string, bool)`

GetPaletteOk returns a tuple with the Palette field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPalette

`func (o *WidgetRequestStyle) SetPalette(v string)`

SetPalette sets Palette field to given value.

### HasPalette

`func (o *WidgetRequestStyle) HasPalette() bool`

HasPalette returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
