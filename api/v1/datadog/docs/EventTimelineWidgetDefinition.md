# EventTimelineWidgetDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | Pointer to **string** | Query to filter the event timeline with. | 
**TagsExecution** | Pointer to **string** | The execution method for multi-value filters. Can be either and or or. | [optional] 
**Time** | Pointer to [**WidgetTime**](WidgetTime.md) |  | [optional] 
**Title** | Pointer to **string** | Title of the widget. | [optional] 
**TitleAlign** | Pointer to [**WidgetTextAlign**](WidgetTextAlign.md) |  | [optional] 
**TitleSize** | Pointer to **string** | Size of the title. | [optional] 
**Type** | Pointer to [**EventTimelineWidgetDefinitionType**](EventTimelineWidgetDefinitionType.md) |  | [default to "event_timeline"]

## Methods

### NewEventTimelineWidgetDefinition

`func NewEventTimelineWidgetDefinition(query string, type_ EventTimelineWidgetDefinitionType, ) *EventTimelineWidgetDefinition`

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

`func (o *EventTimelineWidgetDefinition) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *EventTimelineWidgetDefinition) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetTagsExecution

`func (o *EventTimelineWidgetDefinition) GetTagsExecution() string`

GetTagsExecution returns the TagsExecution field if non-nil, zero value otherwise.

### GetTagsExecutionOk

`func (o *EventTimelineWidgetDefinition) GetTagsExecutionOk() (*string, bool)`

GetTagsExecutionOk returns a tuple with the TagsExecution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagsExecution

`func (o *EventTimelineWidgetDefinition) SetTagsExecution(v string)`

SetTagsExecution sets TagsExecution field to given value.

### HasTagsExecution

`func (o *EventTimelineWidgetDefinition) HasTagsExecution() bool`

HasTagsExecution returns a boolean if a field has been set.

### GetTime

`func (o *EventTimelineWidgetDefinition) GetTime() WidgetTime`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *EventTimelineWidgetDefinition) GetTimeOk() (*WidgetTime, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *EventTimelineWidgetDefinition) SetTime(v WidgetTime)`

SetTime sets Time field to given value.

### HasTime

`func (o *EventTimelineWidgetDefinition) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetTitle

`func (o *EventTimelineWidgetDefinition) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *EventTimelineWidgetDefinition) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *EventTimelineWidgetDefinition) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *EventTimelineWidgetDefinition) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTitleAlign

`func (o *EventTimelineWidgetDefinition) GetTitleAlign() WidgetTextAlign`

GetTitleAlign returns the TitleAlign field if non-nil, zero value otherwise.

### GetTitleAlignOk

`func (o *EventTimelineWidgetDefinition) GetTitleAlignOk() (*WidgetTextAlign, bool)`

GetTitleAlignOk returns a tuple with the TitleAlign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleAlign

`func (o *EventTimelineWidgetDefinition) SetTitleAlign(v WidgetTextAlign)`

SetTitleAlign sets TitleAlign field to given value.

### HasTitleAlign

`func (o *EventTimelineWidgetDefinition) HasTitleAlign() bool`

HasTitleAlign returns a boolean if a field has been set.

### GetTitleSize

`func (o *EventTimelineWidgetDefinition) GetTitleSize() string`

GetTitleSize returns the TitleSize field if non-nil, zero value otherwise.

### GetTitleSizeOk

`func (o *EventTimelineWidgetDefinition) GetTitleSizeOk() (*string, bool)`

GetTitleSizeOk returns a tuple with the TitleSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleSize

`func (o *EventTimelineWidgetDefinition) SetTitleSize(v string)`

SetTitleSize sets TitleSize field to given value.

### HasTitleSize

`func (o *EventTimelineWidgetDefinition) HasTitleSize() bool`

HasTitleSize returns a boolean if a field has been set.

### GetType

`func (o *EventTimelineWidgetDefinition) GetType() EventTimelineWidgetDefinitionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventTimelineWidgetDefinition) GetTypeOk() (*EventTimelineWidgetDefinitionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventTimelineWidgetDefinition) SetType(v EventTimelineWidgetDefinitionType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


