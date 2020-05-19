# NoteWidgetDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackgroundColor** | Pointer to **string** | Background color of the note. | [optional] 
**Content** | Pointer to **string** | Content of the note. | 
**FontSize** | Pointer to **string** | Size of the text. | [optional] 
**ShowTick** | Pointer to **bool** | Whether to show a tick or not. | [optional] 
**TextAlign** | Pointer to [**WidgetTextAlign**](WidgetTextAlign.md) |  | [optional] 
**TickEdge** | Pointer to [**WidgetTickEdge**](WidgetTickEdge.md) |  | [optional] 
**TickPos** | Pointer to **string** | Where to position the tick on an edge. | [optional] 
**Type** | Pointer to [**NoteWidgetDefinitionType**](NoteWidgetDefinitionType.md) |  | [default to "note"]

## Methods

### NewNoteWidgetDefinition

`func NewNoteWidgetDefinition(content string, type_ NoteWidgetDefinitionType, ) *NoteWidgetDefinition`

NewNoteWidgetDefinition instantiates a new NoteWidgetDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNoteWidgetDefinitionWithDefaults

`func NewNoteWidgetDefinitionWithDefaults() *NoteWidgetDefinition`

NewNoteWidgetDefinitionWithDefaults instantiates a new NoteWidgetDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackgroundColor

`func (o *NoteWidgetDefinition) GetBackgroundColor() string`

GetBackgroundColor returns the BackgroundColor field if non-nil, zero value otherwise.

### GetBackgroundColorOk

`func (o *NoteWidgetDefinition) GetBackgroundColorOk() (*string, bool)`

GetBackgroundColorOk returns a tuple with the BackgroundColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundColor

`func (o *NoteWidgetDefinition) SetBackgroundColor(v string)`

SetBackgroundColor sets BackgroundColor field to given value.

### HasBackgroundColor

`func (o *NoteWidgetDefinition) HasBackgroundColor() bool`

HasBackgroundColor returns a boolean if a field has been set.

### GetContent

`func (o *NoteWidgetDefinition) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *NoteWidgetDefinition) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *NoteWidgetDefinition) SetContent(v string)`

SetContent sets Content field to given value.


### GetFontSize

`func (o *NoteWidgetDefinition) GetFontSize() string`

GetFontSize returns the FontSize field if non-nil, zero value otherwise.

### GetFontSizeOk

`func (o *NoteWidgetDefinition) GetFontSizeOk() (*string, bool)`

GetFontSizeOk returns a tuple with the FontSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFontSize

`func (o *NoteWidgetDefinition) SetFontSize(v string)`

SetFontSize sets FontSize field to given value.

### HasFontSize

`func (o *NoteWidgetDefinition) HasFontSize() bool`

HasFontSize returns a boolean if a field has been set.

### GetShowTick

`func (o *NoteWidgetDefinition) GetShowTick() bool`

GetShowTick returns the ShowTick field if non-nil, zero value otherwise.

### GetShowTickOk

`func (o *NoteWidgetDefinition) GetShowTickOk() (*bool, bool)`

GetShowTickOk returns a tuple with the ShowTick field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowTick

`func (o *NoteWidgetDefinition) SetShowTick(v bool)`

SetShowTick sets ShowTick field to given value.

### HasShowTick

`func (o *NoteWidgetDefinition) HasShowTick() bool`

HasShowTick returns a boolean if a field has been set.

### GetTextAlign

`func (o *NoteWidgetDefinition) GetTextAlign() WidgetTextAlign`

GetTextAlign returns the TextAlign field if non-nil, zero value otherwise.

### GetTextAlignOk

`func (o *NoteWidgetDefinition) GetTextAlignOk() (*WidgetTextAlign, bool)`

GetTextAlignOk returns a tuple with the TextAlign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextAlign

`func (o *NoteWidgetDefinition) SetTextAlign(v WidgetTextAlign)`

SetTextAlign sets TextAlign field to given value.

### HasTextAlign

`func (o *NoteWidgetDefinition) HasTextAlign() bool`

HasTextAlign returns a boolean if a field has been set.

### GetTickEdge

`func (o *NoteWidgetDefinition) GetTickEdge() WidgetTickEdge`

GetTickEdge returns the TickEdge field if non-nil, zero value otherwise.

### GetTickEdgeOk

`func (o *NoteWidgetDefinition) GetTickEdgeOk() (*WidgetTickEdge, bool)`

GetTickEdgeOk returns a tuple with the TickEdge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickEdge

`func (o *NoteWidgetDefinition) SetTickEdge(v WidgetTickEdge)`

SetTickEdge sets TickEdge field to given value.

### HasTickEdge

`func (o *NoteWidgetDefinition) HasTickEdge() bool`

HasTickEdge returns a boolean if a field has been set.

### GetTickPos

`func (o *NoteWidgetDefinition) GetTickPos() string`

GetTickPos returns the TickPos field if non-nil, zero value otherwise.

### GetTickPosOk

`func (o *NoteWidgetDefinition) GetTickPosOk() (*string, bool)`

GetTickPosOk returns a tuple with the TickPos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickPos

`func (o *NoteWidgetDefinition) SetTickPos(v string)`

SetTickPos sets TickPos field to given value.

### HasTickPos

`func (o *NoteWidgetDefinition) HasTickPos() bool`

HasTickPos returns a boolean if a field has been set.

### GetType

`func (o *NoteWidgetDefinition) GetType() NoteWidgetDefinitionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NoteWidgetDefinition) GetTypeOk() (*NoteWidgetDefinitionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NoteWidgetDefinition) SetType(v NoteWidgetDefinitionType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


