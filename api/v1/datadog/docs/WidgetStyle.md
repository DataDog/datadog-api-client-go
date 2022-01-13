# WidgetStyle

## Properties

| Name        | Type                  | Description                           | Notes      |
| ----------- | --------------------- | ------------------------------------- | ---------- |
| **Palette** | Pointer to **string** | Color palette to apply to the widget. | [optional] |

## Methods

### NewWidgetStyle

`func NewWidgetStyle() *WidgetStyle`

NewWidgetStyle instantiates a new WidgetStyle object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewWidgetStyleWithDefaults

`func NewWidgetStyleWithDefaults() *WidgetStyle`

NewWidgetStyleWithDefaults instantiates a new WidgetStyle object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetPalette

`func (o *WidgetStyle) GetPalette() string`

GetPalette returns the Palette field if non-nil, zero value otherwise.

### GetPaletteOk

`func (o *WidgetStyle) GetPaletteOk() (*string, bool)`

GetPaletteOk returns a tuple with the Palette field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPalette

`func (o *WidgetStyle) SetPalette(v string)`

SetPalette sets Palette field to given value.

### HasPalette

`func (o *WidgetStyle) HasPalette() bool`

HasPalette returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
