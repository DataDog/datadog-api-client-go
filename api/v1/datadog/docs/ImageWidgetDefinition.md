# ImageWidgetDefinition

## Properties

| Name                | Type                                                             | Description                             | Notes                                        |
| ------------------- | ---------------------------------------------------------------- | --------------------------------------- | -------------------------------------------- |
| **HasBackground**   | Pointer to **bool**                                              | Whether to display a background or not. | [optional] [default to true]                 |
| **HasBorder**       | Pointer to **bool**                                              | Whether to display a border or not.     | [optional] [default to true]                 |
| **HorizontalAlign** | Pointer to [**WidgetHorizontalAlign**](WidgetHorizontalAlign.md) |                                         | [optional]                                   |
| **Margin**          | Pointer to [**WidgetMargin**](WidgetMargin.md)                   |                                         | [optional]                                   |
| **Sizing**          | Pointer to [**WidgetImageSizing**](WidgetImageSizing.md)         |                                         | [optional]                                   |
| **Type**            | [**ImageWidgetDefinitionType**](ImageWidgetDefinitionType.md)    |                                         | [default to IMAGEWIDGETDEFINITIONTYPE_IMAGE] |
| **Url**             | **string**                                                       | URL of the image.                       |
| **UrlDarkTheme**    | Pointer to **string**                                            | URL of the image in dark mode.          | [optional]                                   |
| **VerticalAlign**   | Pointer to [**WidgetVerticalAlign**](WidgetVerticalAlign.md)     |                                         | [optional]                                   |

## Methods

### NewImageWidgetDefinition

`func NewImageWidgetDefinition(type_ ImageWidgetDefinitionType, url string) *ImageWidgetDefinition`

NewImageWidgetDefinition instantiates a new ImageWidgetDefinition object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewImageWidgetDefinitionWithDefaults

`func NewImageWidgetDefinitionWithDefaults() *ImageWidgetDefinition`

NewImageWidgetDefinitionWithDefaults instantiates a new ImageWidgetDefinition object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetHasBackground

`func (o *ImageWidgetDefinition) GetHasBackground() bool`

GetHasBackground returns the HasBackground field if non-nil, zero value otherwise.

### GetHasBackgroundOk

`func (o *ImageWidgetDefinition) GetHasBackgroundOk() (*bool, bool)`

GetHasBackgroundOk returns a tuple with the HasBackground field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasBackground

`func (o *ImageWidgetDefinition) SetHasBackground(v bool)`

SetHasBackground sets HasBackground field to given value.

### HasHasBackground

`func (o *ImageWidgetDefinition) HasHasBackground() bool`

HasHasBackground returns a boolean if a field has been set.

### GetHasBorder

`func (o *ImageWidgetDefinition) GetHasBorder() bool`

GetHasBorder returns the HasBorder field if non-nil, zero value otherwise.

### GetHasBorderOk

`func (o *ImageWidgetDefinition) GetHasBorderOk() (*bool, bool)`

GetHasBorderOk returns a tuple with the HasBorder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasBorder

`func (o *ImageWidgetDefinition) SetHasBorder(v bool)`

SetHasBorder sets HasBorder field to given value.

### HasHasBorder

`func (o *ImageWidgetDefinition) HasHasBorder() bool`

HasHasBorder returns a boolean if a field has been set.

### GetHorizontalAlign

`func (o *ImageWidgetDefinition) GetHorizontalAlign() WidgetHorizontalAlign`

GetHorizontalAlign returns the HorizontalAlign field if non-nil, zero value otherwise.

### GetHorizontalAlignOk

`func (o *ImageWidgetDefinition) GetHorizontalAlignOk() (*WidgetHorizontalAlign, bool)`

GetHorizontalAlignOk returns a tuple with the HorizontalAlign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHorizontalAlign

`func (o *ImageWidgetDefinition) SetHorizontalAlign(v WidgetHorizontalAlign)`

SetHorizontalAlign sets HorizontalAlign field to given value.

### HasHorizontalAlign

`func (o *ImageWidgetDefinition) HasHorizontalAlign() bool`

HasHorizontalAlign returns a boolean if a field has been set.

### GetMargin

`func (o *ImageWidgetDefinition) GetMargin() WidgetMargin`

GetMargin returns the Margin field if non-nil, zero value otherwise.

### GetMarginOk

`func (o *ImageWidgetDefinition) GetMarginOk() (*WidgetMargin, bool)`

GetMarginOk returns a tuple with the Margin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMargin

`func (o *ImageWidgetDefinition) SetMargin(v WidgetMargin)`

SetMargin sets Margin field to given value.

### HasMargin

`func (o *ImageWidgetDefinition) HasMargin() bool`

HasMargin returns a boolean if a field has been set.

### GetSizing

`func (o *ImageWidgetDefinition) GetSizing() WidgetImageSizing`

GetSizing returns the Sizing field if non-nil, zero value otherwise.

### GetSizingOk

`func (o *ImageWidgetDefinition) GetSizingOk() (*WidgetImageSizing, bool)`

GetSizingOk returns a tuple with the Sizing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizing

`func (o *ImageWidgetDefinition) SetSizing(v WidgetImageSizing)`

SetSizing sets Sizing field to given value.

### HasSizing

`func (o *ImageWidgetDefinition) HasSizing() bool`

HasSizing returns a boolean if a field has been set.

### GetType

`func (o *ImageWidgetDefinition) GetType() ImageWidgetDefinitionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ImageWidgetDefinition) GetTypeOk() (*ImageWidgetDefinitionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ImageWidgetDefinition) SetType(v ImageWidgetDefinitionType)`

SetType sets Type field to given value.

### GetUrl

`func (o *ImageWidgetDefinition) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ImageWidgetDefinition) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ImageWidgetDefinition) SetUrl(v string)`

SetUrl sets Url field to given value.

### GetUrlDarkTheme

`func (o *ImageWidgetDefinition) GetUrlDarkTheme() string`

GetUrlDarkTheme returns the UrlDarkTheme field if non-nil, zero value otherwise.

### GetUrlDarkThemeOk

`func (o *ImageWidgetDefinition) GetUrlDarkThemeOk() (*string, bool)`

GetUrlDarkThemeOk returns a tuple with the UrlDarkTheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlDarkTheme

`func (o *ImageWidgetDefinition) SetUrlDarkTheme(v string)`

SetUrlDarkTheme sets UrlDarkTheme field to given value.

### HasUrlDarkTheme

`func (o *ImageWidgetDefinition) HasUrlDarkTheme() bool`

HasUrlDarkTheme returns a boolean if a field has been set.

### GetVerticalAlign

`func (o *ImageWidgetDefinition) GetVerticalAlign() WidgetVerticalAlign`

GetVerticalAlign returns the VerticalAlign field if non-nil, zero value otherwise.

### GetVerticalAlignOk

`func (o *ImageWidgetDefinition) GetVerticalAlignOk() (*WidgetVerticalAlign, bool)`

GetVerticalAlignOk returns a tuple with the VerticalAlign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerticalAlign

`func (o *ImageWidgetDefinition) SetVerticalAlign(v WidgetVerticalAlign)`

SetVerticalAlign sets VerticalAlign field to given value.

### HasVerticalAlign

`func (o *ImageWidgetDefinition) HasVerticalAlign() bool`

HasVerticalAlign returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
