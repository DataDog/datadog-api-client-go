# EventTimelineWidgetDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | Pointer to **string** | Query to filter the event timeline with | 
**TagsExecution** | Pointer to **string** | The execution method for multi-value filters. Can be either and or or | [optional] 
**Time** | Pointer to [**WidgetTime**](WidgetTime.md) |  | [optional] 
**Title** | Pointer to **string** | Title of the widget | [optional] 
**TitleAlign** | Pointer to [**WidgetTextAlign**](WidgetTextAlign.md) |  | [optional] 
**TitleSize** | Pointer to **string** | Size of the title | [optional] 
**Type** | Pointer to **string** | Type of the widget | [readonly] [default to "event_timeline"]

## Methods

### NewEventTimelineWidgetDefinition

`func NewEventTimelineWidgetDefinition(query string, type_ string, ) *EventTimelineWidgetDefinition`

NewEventTimelineWidgetDefinition instantiates a new EventTimelineWidgetDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventTimelineWidgetDefinitionWithDefaults

`func NewEventTimelineWidgetDefinitionWithDefaults() *EventTimelineWidgetDefinition`

NewEventTimelineWidgetDefinitionWithDefaults instantiates a new EventTimelineWidgetDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *EventTimelineWidgetDefinition) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *EventTimelineWidgetDefinition) GetQueryOk() (string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasQuery

`func (o *EventTimelineWidgetDefinition) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### SetQuery

`func (o *EventTimelineWidgetDefinition) SetQuery(v string)`

SetQuery gets a reference to the given string and assigns it to the Query field.

### GetTagsExecution

`func (o *EventTimelineWidgetDefinition) GetTagsExecution() string`

GetTagsExecution returns the TagsExecution field if non-nil, zero value otherwise.

### GetTagsExecutionOk

`func (o *EventTimelineWidgetDefinition) GetTagsExecutionOk() (string, bool)`

GetTagsExecutionOk returns a tuple with the TagsExecution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTagsExecution

`func (o *EventTimelineWidgetDefinition) HasTagsExecution() bool`

HasTagsExecution returns a boolean if a field has been set.

### SetTagsExecution

`func (o *EventTimelineWidgetDefinition) SetTagsExecution(v string)`

SetTagsExecution gets a reference to the given string and assigns it to the TagsExecution field.

### GetTime

`func (o *EventTimelineWidgetDefinition) GetTime() WidgetTime`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *EventTimelineWidgetDefinition) GetTimeOk() (WidgetTime, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTime

`func (o *EventTimelineWidgetDefinition) HasTime() bool`

HasTime returns a boolean if a field has been set.

### SetTime

`func (o *EventTimelineWidgetDefinition) SetTime(v WidgetTime)`

SetTime gets a reference to the given WidgetTime and assigns it to the Time field.

### GetTitle

`func (o *EventTimelineWidgetDefinition) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *EventTimelineWidgetDefinition) GetTitleOk() (string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitle

`func (o *EventTimelineWidgetDefinition) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitle

`func (o *EventTimelineWidgetDefinition) SetTitle(v string)`

SetTitle gets a reference to the given string and assigns it to the Title field.

### GetTitleAlign

`func (o *EventTimelineWidgetDefinition) GetTitleAlign() WidgetTextAlign`

GetTitleAlign returns the TitleAlign field if non-nil, zero value otherwise.

### GetTitleAlignOk

`func (o *EventTimelineWidgetDefinition) GetTitleAlignOk() (WidgetTextAlign, bool)`

GetTitleAlignOk returns a tuple with the TitleAlign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitleAlign

`func (o *EventTimelineWidgetDefinition) HasTitleAlign() bool`

HasTitleAlign returns a boolean if a field has been set.

### SetTitleAlign

`func (o *EventTimelineWidgetDefinition) SetTitleAlign(v WidgetTextAlign)`

SetTitleAlign gets a reference to the given WidgetTextAlign and assigns it to the TitleAlign field.

### GetTitleSize

`func (o *EventTimelineWidgetDefinition) GetTitleSize() string`

GetTitleSize returns the TitleSize field if non-nil, zero value otherwise.

### GetTitleSizeOk

`func (o *EventTimelineWidgetDefinition) GetTitleSizeOk() (string, bool)`

GetTitleSizeOk returns a tuple with the TitleSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitleSize

`func (o *EventTimelineWidgetDefinition) HasTitleSize() bool`

HasTitleSize returns a boolean if a field has been set.

### SetTitleSize

`func (o *EventTimelineWidgetDefinition) SetTitleSize(v string)`

SetTitleSize gets a reference to the given string and assigns it to the TitleSize field.

### GetType

`func (o *EventTimelineWidgetDefinition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventTimelineWidgetDefinition) GetTypeOk() (string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasType

`func (o *EventTimelineWidgetDefinition) HasType() bool`

HasType returns a boolean if a field has been set.

### SetType

`func (o *EventTimelineWidgetDefinition) SetType(v string)`

SetType gets a reference to the given string and assigns it to the Type field.


### AsWidgetDefinition

`func (s *EventTimelineWidgetDefinition) AsWidgetDefinition() WidgetDefinition`

Convenience method to wrap this instance of EventTimelineWidgetDefinition in WidgetDefinition

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


