# Event

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AggregationKey** | Pointer to **string** | An arbitrary string to use for aggregation. Limited to 100 characters. If you specify a key, all events using that key are grouped together in the Event Stream. | [optional] 
**AlertType** | Pointer to **string** | If it is an alert event, set its type between: error, warning, info, and success. | [optional] 
**DateHappened** | Pointer to **int64** | POSIX timestamp of the event. Must be sent as an integer (i.e. no quotes). Limited to events no older than 1 year, 24 days (389 days) | [optional] 
**DeviceName** | Pointer to **[]string** | A list of device names to post the event with. | [optional] 
**Host** | Pointer to **string** | Host name to associate with the event. Any tags associated with the host are also applied to this event. | [optional] 
**Id** | Pointer to **int64** |  | [optional] [readonly] 
**Payload** | Pointer to **string** |  | [optional] [readonly] 
**Priority** | Pointer to **string** | The priority of the event: normal or low. | [optional] 
**RelatedEventId** | Pointer to **int64** | ID of the parent event. Must be sent as an integer (i.e. no quotes). | [optional] 
**SourceTypeName** | Pointer to **string** | The type of event being posted. Options: nagios, hudson, jenkins, my_apps, chef, puppet, git, bitbucket, ... [Complete list of source attribute values](https://docs.datadoghq.com/integrations/faq/list-of-api-source-attribute-value) | [optional] 
**Tags** | Pointer to **[]string** | A list of tags to apply to the event. | [optional] 
**Text** | Pointer to **string** | The body of the event. Limited to 4000 characters. The text supports markdown. Use msg_text with the Datadog Ruby library | 
**Title** | Pointer to **string** | The event title. Limited to 100 characters. Use msg_title with the Datadog Ruby library. | 
**Url** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewEvent

`func NewEvent(text string, title string, ) *Event`

NewEvent instantiates a new Event object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventWithDefaults

`func NewEventWithDefaults() *Event`

NewEventWithDefaults instantiates a new Event object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregationKey

`func (o *Event) GetAggregationKey() string`

GetAggregationKey returns the AggregationKey field if non-nil, zero value otherwise.

### GetAggregationKeyOk

`func (o *Event) GetAggregationKeyOk() (string, bool)`

GetAggregationKeyOk returns a tuple with the AggregationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAggregationKey

`func (o *Event) HasAggregationKey() bool`

HasAggregationKey returns a boolean if a field has been set.

### SetAggregationKey

`func (o *Event) SetAggregationKey(v string)`

SetAggregationKey gets a reference to the given string and assigns it to the AggregationKey field.

### GetAlertType

`func (o *Event) GetAlertType() string`

GetAlertType returns the AlertType field if non-nil, zero value otherwise.

### GetAlertTypeOk

`func (o *Event) GetAlertTypeOk() (string, bool)`

GetAlertTypeOk returns a tuple with the AlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAlertType

`func (o *Event) HasAlertType() bool`

HasAlertType returns a boolean if a field has been set.

### SetAlertType

`func (o *Event) SetAlertType(v string)`

SetAlertType gets a reference to the given string and assigns it to the AlertType field.

### GetDateHappened

`func (o *Event) GetDateHappened() int64`

GetDateHappened returns the DateHappened field if non-nil, zero value otherwise.

### GetDateHappenedOk

`func (o *Event) GetDateHappenedOk() (int64, bool)`

GetDateHappenedOk returns a tuple with the DateHappened field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDateHappened

`func (o *Event) HasDateHappened() bool`

HasDateHappened returns a boolean if a field has been set.

### SetDateHappened

`func (o *Event) SetDateHappened(v int64)`

SetDateHappened gets a reference to the given int64 and assigns it to the DateHappened field.

### GetDeviceName

`func (o *Event) GetDeviceName() []string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *Event) GetDeviceNameOk() ([]string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceName

`func (o *Event) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### SetDeviceName

`func (o *Event) SetDeviceName(v []string)`

SetDeviceName gets a reference to the given []string and assigns it to the DeviceName field.

### GetHost

`func (o *Event) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *Event) GetHostOk() (string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHost

`func (o *Event) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHost

`func (o *Event) SetHost(v string)`

SetHost gets a reference to the given string and assigns it to the Host field.

### GetId

`func (o *Event) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Event) GetIdOk() (int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *Event) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *Event) SetId(v int64)`

SetId gets a reference to the given int64 and assigns it to the Id field.

### GetPayload

`func (o *Event) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *Event) GetPayloadOk() (string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPayload

`func (o *Event) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### SetPayload

`func (o *Event) SetPayload(v string)`

SetPayload gets a reference to the given string and assigns it to the Payload field.

### GetPriority

`func (o *Event) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *Event) GetPriorityOk() (string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPriority

`func (o *Event) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriority

`func (o *Event) SetPriority(v string)`

SetPriority gets a reference to the given string and assigns it to the Priority field.

### GetRelatedEventId

`func (o *Event) GetRelatedEventId() int64`

GetRelatedEventId returns the RelatedEventId field if non-nil, zero value otherwise.

### GetRelatedEventIdOk

`func (o *Event) GetRelatedEventIdOk() (int64, bool)`

GetRelatedEventIdOk returns a tuple with the RelatedEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRelatedEventId

`func (o *Event) HasRelatedEventId() bool`

HasRelatedEventId returns a boolean if a field has been set.

### SetRelatedEventId

`func (o *Event) SetRelatedEventId(v int64)`

SetRelatedEventId gets a reference to the given int64 and assigns it to the RelatedEventId field.

### GetSourceTypeName

`func (o *Event) GetSourceTypeName() string`

GetSourceTypeName returns the SourceTypeName field if non-nil, zero value otherwise.

### GetSourceTypeNameOk

`func (o *Event) GetSourceTypeNameOk() (string, bool)`

GetSourceTypeNameOk returns a tuple with the SourceTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSourceTypeName

`func (o *Event) HasSourceTypeName() bool`

HasSourceTypeName returns a boolean if a field has been set.

### SetSourceTypeName

`func (o *Event) SetSourceTypeName(v string)`

SetSourceTypeName gets a reference to the given string and assigns it to the SourceTypeName field.

### GetTags

`func (o *Event) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Event) GetTagsOk() ([]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTags

`func (o *Event) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTags

`func (o *Event) SetTags(v []string)`

SetTags gets a reference to the given []string and assigns it to the Tags field.

### GetText

`func (o *Event) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *Event) GetTextOk() (string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasText

`func (o *Event) HasText() bool`

HasText returns a boolean if a field has been set.

### SetText

`func (o *Event) SetText(v string)`

SetText gets a reference to the given string and assigns it to the Text field.

### GetTitle

`func (o *Event) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Event) GetTitleOk() (string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitle

`func (o *Event) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitle

`func (o *Event) SetTitle(v string)`

SetTitle gets a reference to the given string and assigns it to the Title field.

### GetUrl

`func (o *Event) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Event) GetUrlOk() (string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUrl

`func (o *Event) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrl

`func (o *Event) SetUrl(v string)`

SetUrl gets a reference to the given string and assigns it to the Url field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


