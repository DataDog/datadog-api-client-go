# FreeTextWidgetDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Color** | Pointer to **string** | Color of the text | [optional] 
**FontSize** | Pointer to **string** | Size of the text | [optional] 
**Text** | Pointer to **string** | Text to display | 
**TextAlign** | Pointer to [**WidgetTextAlign**](WidgetTextAlign.md) |  | [optional] 
**Type** | Pointer to **string** | Type of the widget | [readonly] [default to "free_text"]

## Methods

### NewFreeTextWidgetDefinition

`func NewFreeTextWidgetDefinition(text string, type_ string, ) *FreeTextWidgetDefinition`

NewFreeTextWidgetDefinition instantiates a new FreeTextWidgetDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFreeTextWidgetDefinitionWithDefaults

`func NewFreeTextWidgetDefinitionWithDefaults() *FreeTextWidgetDefinition`

NewFreeTextWidgetDefinitionWithDefaults instantiates a new FreeTextWidgetDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColor

`func (o *FreeTextWidgetDefinition) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *FreeTextWidgetDefinition) GetColorOk() (string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasColor

`func (o *FreeTextWidgetDefinition) HasColor() bool`

HasColor returns a boolean if a field has been set.

### SetColor

`func (o *FreeTextWidgetDefinition) SetColor(v string)`

SetColor gets a reference to the given string and assigns it to the Color field.

### GetFontSize

`func (o *FreeTextWidgetDefinition) GetFontSize() string`

GetFontSize returns the FontSize field if non-nil, zero value otherwise.

### GetFontSizeOk

`func (o *FreeTextWidgetDefinition) GetFontSizeOk() (string, bool)`

GetFontSizeOk returns a tuple with the FontSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFontSize

`func (o *FreeTextWidgetDefinition) HasFontSize() bool`

HasFontSize returns a boolean if a field has been set.

### SetFontSize

`func (o *FreeTextWidgetDefinition) SetFontSize(v string)`

SetFontSize gets a reference to the given string and assigns it to the FontSize field.

### GetText

`func (o *FreeTextWidgetDefinition) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *FreeTextWidgetDefinition) GetTextOk() (string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasText

`func (o *FreeTextWidgetDefinition) HasText() bool`

HasText returns a boolean if a field has been set.

### SetText

`func (o *FreeTextWidgetDefinition) SetText(v string)`

SetText gets a reference to the given string and assigns it to the Text field.

### GetTextAlign

`func (o *FreeTextWidgetDefinition) GetTextAlign() WidgetTextAlign`

GetTextAlign returns the TextAlign field if non-nil, zero value otherwise.

### GetTextAlignOk

`func (o *FreeTextWidgetDefinition) GetTextAlignOk() (WidgetTextAlign, bool)`

GetTextAlignOk returns a tuple with the TextAlign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTextAlign

`func (o *FreeTextWidgetDefinition) HasTextAlign() bool`

HasTextAlign returns a boolean if a field has been set.

### SetTextAlign

`func (o *FreeTextWidgetDefinition) SetTextAlign(v WidgetTextAlign)`

SetTextAlign gets a reference to the given WidgetTextAlign and assigns it to the TextAlign field.

### GetType

`func (o *FreeTextWidgetDefinition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FreeTextWidgetDefinition) GetTypeOk() (string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasType

`func (o *FreeTextWidgetDefinition) HasType() bool`

HasType returns a boolean if a field has been set.

### SetType

`func (o *FreeTextWidgetDefinition) SetType(v string)`

SetType gets a reference to the given string and assigns it to the Type field.


### AsWidgetDefinition

`func (s *FreeTextWidgetDefinition) AsWidgetDefinition() WidgetDefinition`

Convenience method to wrap this instance of FreeTextWidgetDefinition in WidgetDefinition

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


