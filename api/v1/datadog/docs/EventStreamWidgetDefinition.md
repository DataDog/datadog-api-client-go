# EventStreamWidgetDefinition

## Properties

| Name              | Type                                                                      | Description                                                            | Notes                                                     |
| ----------------- | ------------------------------------------------------------------------- | ---------------------------------------------------------------------- | --------------------------------------------------------- |
| **EventSize**     | Pointer to [**WidgetEventSize**](WidgetEventSize.md)                      |                                                                        | [optional]                                                |
| **Query**         | **string**                                                                | Query to filter the event stream with.                                 |
| **TagsExecution** | Pointer to **string**                                                     | The execution method for multi-value filters. Can be either and or or. | [optional]                                                |
| **Time**          | Pointer to [**WidgetTime**](WidgetTime.md)                                |                                                                        | [optional]                                                |
| **Title**         | Pointer to **string**                                                     | Title of the widget.                                                   | [optional]                                                |
| **TitleAlign**    | Pointer to [**WidgetTextAlign**](WidgetTextAlign.md)                      |                                                                        | [optional]                                                |
| **TitleSize**     | Pointer to **string**                                                     | Size of the title.                                                     | [optional]                                                |
| **Type**          | [**EventStreamWidgetDefinitionType**](EventStreamWidgetDefinitionType.md) |                                                                        | [default to EVENTSTREAMWIDGETDEFINITIONTYPE_EVENT_STREAM] |

## Methods

### NewEventStreamWidgetDefinition

`func NewEventStreamWidgetDefinition(query string, type_ EventStreamWidgetDefinitionType) *EventStreamWidgetDefinition`

NewEventStreamWidgetDefinition instantiates a new EventStreamWidgetDefinition object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewEventStreamWidgetDefinitionWithDefaults

`func NewEventStreamWidgetDefinitionWithDefaults() *EventStreamWidgetDefinition`

NewEventStreamWidgetDefinitionWithDefaults instantiates a new EventStreamWidgetDefinition object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetEventSize

`func (o *EventStreamWidgetDefinition) GetEventSize() WidgetEventSize`

GetEventSize returns the EventSize field if non-nil, zero value otherwise.

### GetEventSizeOk

`func (o *EventStreamWidgetDefinition) GetEventSizeOk() (*WidgetEventSize, bool)`

GetEventSizeOk returns a tuple with the EventSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSize

`func (o *EventStreamWidgetDefinition) SetEventSize(v WidgetEventSize)`

SetEventSize sets EventSize field to given value.

### HasEventSize

`func (o *EventStreamWidgetDefinition) HasEventSize() bool`

HasEventSize returns a boolean if a field has been set.

### GetQuery

`func (o *EventStreamWidgetDefinition) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *EventStreamWidgetDefinition) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *EventStreamWidgetDefinition) SetQuery(v string)`

SetQuery sets Query field to given value.

### GetTagsExecution

`func (o *EventStreamWidgetDefinition) GetTagsExecution() string`

GetTagsExecution returns the TagsExecution field if non-nil, zero value otherwise.

### GetTagsExecutionOk

`func (o *EventStreamWidgetDefinition) GetTagsExecutionOk() (*string, bool)`

GetTagsExecutionOk returns a tuple with the TagsExecution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagsExecution

`func (o *EventStreamWidgetDefinition) SetTagsExecution(v string)`

SetTagsExecution sets TagsExecution field to given value.

### HasTagsExecution

`func (o *EventStreamWidgetDefinition) HasTagsExecution() bool`

HasTagsExecution returns a boolean if a field has been set.

### GetTime

`func (o *EventStreamWidgetDefinition) GetTime() WidgetTime`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *EventStreamWidgetDefinition) GetTimeOk() (*WidgetTime, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *EventStreamWidgetDefinition) SetTime(v WidgetTime)`

SetTime sets Time field to given value.

### HasTime

`func (o *EventStreamWidgetDefinition) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetTitle

`func (o *EventStreamWidgetDefinition) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *EventStreamWidgetDefinition) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *EventStreamWidgetDefinition) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *EventStreamWidgetDefinition) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTitleAlign

`func (o *EventStreamWidgetDefinition) GetTitleAlign() WidgetTextAlign`

GetTitleAlign returns the TitleAlign field if non-nil, zero value otherwise.

### GetTitleAlignOk

`func (o *EventStreamWidgetDefinition) GetTitleAlignOk() (*WidgetTextAlign, bool)`

GetTitleAlignOk returns a tuple with the TitleAlign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleAlign

`func (o *EventStreamWidgetDefinition) SetTitleAlign(v WidgetTextAlign)`

SetTitleAlign sets TitleAlign field to given value.

### HasTitleAlign

`func (o *EventStreamWidgetDefinition) HasTitleAlign() bool`

HasTitleAlign returns a boolean if a field has been set.

### GetTitleSize

`func (o *EventStreamWidgetDefinition) GetTitleSize() string`

GetTitleSize returns the TitleSize field if non-nil, zero value otherwise.

### GetTitleSizeOk

`func (o *EventStreamWidgetDefinition) GetTitleSizeOk() (*string, bool)`

GetTitleSizeOk returns a tuple with the TitleSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleSize

`func (o *EventStreamWidgetDefinition) SetTitleSize(v string)`

SetTitleSize sets TitleSize field to given value.

### HasTitleSize

`func (o *EventStreamWidgetDefinition) HasTitleSize() bool`

HasTitleSize returns a boolean if a field has been set.

### GetType

`func (o *EventStreamWidgetDefinition) GetType() EventStreamWidgetDefinitionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventStreamWidgetDefinition) GetTypeOk() (*EventStreamWidgetDefinitionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventStreamWidgetDefinition) SetType(v EventStreamWidgetDefinitionType)`

SetType sets Type field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
