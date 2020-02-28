# EventStreamWidgetDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventSize** | Pointer to [**WidgetEventSize**](WidgetEventSize.md) |  | [optional] 
**Query** | Pointer to **string** | Query to filter the event stream with | 
**TagsExecution** | Pointer to **string** | The execution method for multi-value filters. Can be either and or or | [optional] 
**Time** | Pointer to [**WidgetTime**](WidgetTime.md) |  | [optional] 
**Title** | Pointer to **string** | Title of the widget | [optional] 
**TitleAlign** | Pointer to [**WidgetTextAlign**](WidgetTextAlign.md) |  | [optional] 
**TitleSize** | Pointer to **string** | Size of the title | [optional] 
**Type** | Pointer to **string** | Type of the widget | [readonly] [default to "event_stream"]

## Methods

### NewEventStreamWidgetDefinition

`func NewEventStreamWidgetDefinition(query string, type_ string, ) *EventStreamWidgetDefinition`

NewEventStreamWidgetDefinition instantiates a new EventStreamWidgetDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventStreamWidgetDefinitionWithDefaults

`func NewEventStreamWidgetDefinitionWithDefaults() *EventStreamWidgetDefinition`

NewEventStreamWidgetDefinitionWithDefaults instantiates a new EventStreamWidgetDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventSize

`func (o *EventStreamWidgetDefinition) GetEventSize() WidgetEventSize`

GetEventSize returns the EventSize field if non-nil, zero value otherwise.

### GetEventSizeOk

`func (o *EventStreamWidgetDefinition) GetEventSizeOk() (WidgetEventSize, bool)`

GetEventSizeOk returns a tuple with the EventSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEventSize

`func (o *EventStreamWidgetDefinition) HasEventSize() bool`

HasEventSize returns a boolean if a field has been set.

### SetEventSize

`func (o *EventStreamWidgetDefinition) SetEventSize(v WidgetEventSize)`

SetEventSize gets a reference to the given WidgetEventSize and assigns it to the EventSize field.

### GetQuery

`func (o *EventStreamWidgetDefinition) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *EventStreamWidgetDefinition) GetQueryOk() (string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasQuery

`func (o *EventStreamWidgetDefinition) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### SetQuery

`func (o *EventStreamWidgetDefinition) SetQuery(v string)`

SetQuery gets a reference to the given string and assigns it to the Query field.

### GetTagsExecution

`func (o *EventStreamWidgetDefinition) GetTagsExecution() string`

GetTagsExecution returns the TagsExecution field if non-nil, zero value otherwise.

### GetTagsExecutionOk

`func (o *EventStreamWidgetDefinition) GetTagsExecutionOk() (string, bool)`

GetTagsExecutionOk returns a tuple with the TagsExecution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTagsExecution

`func (o *EventStreamWidgetDefinition) HasTagsExecution() bool`

HasTagsExecution returns a boolean if a field has been set.

### SetTagsExecution

`func (o *EventStreamWidgetDefinition) SetTagsExecution(v string)`

SetTagsExecution gets a reference to the given string and assigns it to the TagsExecution field.

### GetTime

`func (o *EventStreamWidgetDefinition) GetTime() WidgetTime`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *EventStreamWidgetDefinition) GetTimeOk() (WidgetTime, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTime

`func (o *EventStreamWidgetDefinition) HasTime() bool`

HasTime returns a boolean if a field has been set.

### SetTime

`func (o *EventStreamWidgetDefinition) SetTime(v WidgetTime)`

SetTime gets a reference to the given WidgetTime and assigns it to the Time field.

### GetTitle

`func (o *EventStreamWidgetDefinition) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *EventStreamWidgetDefinition) GetTitleOk() (string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitle

`func (o *EventStreamWidgetDefinition) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitle

`func (o *EventStreamWidgetDefinition) SetTitle(v string)`

SetTitle gets a reference to the given string and assigns it to the Title field.

### GetTitleAlign

`func (o *EventStreamWidgetDefinition) GetTitleAlign() WidgetTextAlign`

GetTitleAlign returns the TitleAlign field if non-nil, zero value otherwise.

### GetTitleAlignOk

`func (o *EventStreamWidgetDefinition) GetTitleAlignOk() (WidgetTextAlign, bool)`

GetTitleAlignOk returns a tuple with the TitleAlign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitleAlign

`func (o *EventStreamWidgetDefinition) HasTitleAlign() bool`

HasTitleAlign returns a boolean if a field has been set.

### SetTitleAlign

`func (o *EventStreamWidgetDefinition) SetTitleAlign(v WidgetTextAlign)`

SetTitleAlign gets a reference to the given WidgetTextAlign and assigns it to the TitleAlign field.

### GetTitleSize

`func (o *EventStreamWidgetDefinition) GetTitleSize() string`

GetTitleSize returns the TitleSize field if non-nil, zero value otherwise.

### GetTitleSizeOk

`func (o *EventStreamWidgetDefinition) GetTitleSizeOk() (string, bool)`

GetTitleSizeOk returns a tuple with the TitleSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitleSize

`func (o *EventStreamWidgetDefinition) HasTitleSize() bool`

HasTitleSize returns a boolean if a field has been set.

### SetTitleSize

`func (o *EventStreamWidgetDefinition) SetTitleSize(v string)`

SetTitleSize gets a reference to the given string and assigns it to the TitleSize field.

### GetType

`func (o *EventStreamWidgetDefinition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventStreamWidgetDefinition) GetTypeOk() (string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasType

`func (o *EventStreamWidgetDefinition) HasType() bool`

HasType returns a boolean if a field has been set.

### SetType

`func (o *EventStreamWidgetDefinition) SetType(v string)`

SetType gets a reference to the given string and assigns it to the Type field.


### AsWidgetDefinition

`func (s *EventStreamWidgetDefinition) AsWidgetDefinition() WidgetDefinition`

Convenience method to wrap this instance of EventStreamWidgetDefinition in WidgetDefinition

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


