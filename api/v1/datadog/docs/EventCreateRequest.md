# EventCreateRequest

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**AggregationKey** | Pointer to **string** | An arbitrary string to use for aggregation. Limited to 100 characters. If you specify a key, all events using that key are grouped together in the Event Stream. | [optional] 
**AlertType** | Pointer to [**EventAlertType**](EventAlertType.md) |  | [optional] 
**DateHappened** | Pointer to **int64** | POSIX timestamp of the event. Must be sent as an integer (that is no quotes). Limited to events no older than 7 days. | [optional] 
**DeviceName** | Pointer to **string** | A device name. | [optional] 
**Host** | Pointer to **string** | Host name to associate with the event. Any tags associated with the host are also applied to this event. | [optional] 
**Id** | Pointer to **int64** | Integer ID of the event. | [optional] [readonly] 
**Payload** | Pointer to **string** | Payload of the event. | [optional] [readonly] 
**Priority** | Pointer to [**EventPriority**](EventPriority.md) |  | [optional] 
**RelatedEventId** | Pointer to **int64** | ID of the parent event. Must be sent as an integer (that is no quotes). | [optional] 
**SourceTypeName** | Pointer to **string** | The type of event being posted. Option examples include nagios, hudson, jenkins, my_apps, chef, puppet, git, bitbucket, etc. A complete list of source attribute values [available here](https://docs.datadoghq.com/integrations/faq/list-of-api-source-attribute-value). | [optional] 
**Tags** | Pointer to **[]string** | A list of tags to apply to the event. | [optional] 
**Text** | **string** | The body of the event. Limited to 4000 characters. The text supports markdown. To use markdown in the event text, start the text block with &#x60;%%% \\n&#x60; and end the text block with &#x60;\\n %%%&#x60;. Use &#x60;msg_text&#x60; with the Datadog Ruby library. | 
**Title** | **string** | The event title. | 
**Url** | Pointer to **string** | URL of the event. | [optional] [readonly] 

## Methods

### NewEventCreateRequest

`func NewEventCreateRequest(text string, title string) *EventCreateRequest`

NewEventCreateRequest instantiates a new EventCreateRequest object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewEventCreateRequestWithDefaults

`func NewEventCreateRequestWithDefaults() *EventCreateRequest`

NewEventCreateRequestWithDefaults instantiates a new EventCreateRequest object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAggregationKey

`func (o *EventCreateRequest) GetAggregationKey() string`

GetAggregationKey returns the AggregationKey field if non-nil, zero value otherwise.

### GetAggregationKeyOk

`func (o *EventCreateRequest) GetAggregationKeyOk() (*string, bool)`

GetAggregationKeyOk returns a tuple with the AggregationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationKey

`func (o *EventCreateRequest) SetAggregationKey(v string)`

SetAggregationKey sets AggregationKey field to given value.

### HasAggregationKey

`func (o *EventCreateRequest) HasAggregationKey() bool`

HasAggregationKey returns a boolean if a field has been set.

### GetAlertType

`func (o *EventCreateRequest) GetAlertType() EventAlertType`

GetAlertType returns the AlertType field if non-nil, zero value otherwise.

### GetAlertTypeOk

`func (o *EventCreateRequest) GetAlertTypeOk() (*EventAlertType, bool)`

GetAlertTypeOk returns a tuple with the AlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertType

`func (o *EventCreateRequest) SetAlertType(v EventAlertType)`

SetAlertType sets AlertType field to given value.

### HasAlertType

`func (o *EventCreateRequest) HasAlertType() bool`

HasAlertType returns a boolean if a field has been set.

### GetDateHappened

`func (o *EventCreateRequest) GetDateHappened() int64`

GetDateHappened returns the DateHappened field if non-nil, zero value otherwise.

### GetDateHappenedOk

`func (o *EventCreateRequest) GetDateHappenedOk() (*int64, bool)`

GetDateHappenedOk returns a tuple with the DateHappened field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateHappened

`func (o *EventCreateRequest) SetDateHappened(v int64)`

SetDateHappened sets DateHappened field to given value.

### HasDateHappened

`func (o *EventCreateRequest) HasDateHappened() bool`

HasDateHappened returns a boolean if a field has been set.

### GetDeviceName

`func (o *EventCreateRequest) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *EventCreateRequest) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *EventCreateRequest) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *EventCreateRequest) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetHost

`func (o *EventCreateRequest) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *EventCreateRequest) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *EventCreateRequest) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *EventCreateRequest) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *EventCreateRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventCreateRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EventCreateRequest) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *EventCreateRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPayload

`func (o *EventCreateRequest) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *EventCreateRequest) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *EventCreateRequest) SetPayload(v string)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *EventCreateRequest) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetPriority

`func (o *EventCreateRequest) GetPriority() EventPriority`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *EventCreateRequest) GetPriorityOk() (*EventPriority, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *EventCreateRequest) SetPriority(v EventPriority)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *EventCreateRequest) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetRelatedEventId

`func (o *EventCreateRequest) GetRelatedEventId() int64`

GetRelatedEventId returns the RelatedEventId field if non-nil, zero value otherwise.

### GetRelatedEventIdOk

`func (o *EventCreateRequest) GetRelatedEventIdOk() (*int64, bool)`

GetRelatedEventIdOk returns a tuple with the RelatedEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedEventId

`func (o *EventCreateRequest) SetRelatedEventId(v int64)`

SetRelatedEventId sets RelatedEventId field to given value.

### HasRelatedEventId

`func (o *EventCreateRequest) HasRelatedEventId() bool`

HasRelatedEventId returns a boolean if a field has been set.

### GetSourceTypeName

`func (o *EventCreateRequest) GetSourceTypeName() string`

GetSourceTypeName returns the SourceTypeName field if non-nil, zero value otherwise.

### GetSourceTypeNameOk

`func (o *EventCreateRequest) GetSourceTypeNameOk() (*string, bool)`

GetSourceTypeNameOk returns a tuple with the SourceTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTypeName

`func (o *EventCreateRequest) SetSourceTypeName(v string)`

SetSourceTypeName sets SourceTypeName field to given value.

### HasSourceTypeName

`func (o *EventCreateRequest) HasSourceTypeName() bool`

HasSourceTypeName returns a boolean if a field has been set.

### GetTags

`func (o *EventCreateRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *EventCreateRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *EventCreateRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *EventCreateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetText

`func (o *EventCreateRequest) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *EventCreateRequest) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *EventCreateRequest) SetText(v string)`

SetText sets Text field to given value.


### GetTitle

`func (o *EventCreateRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *EventCreateRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *EventCreateRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetUrl

`func (o *EventCreateRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *EventCreateRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *EventCreateRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *EventCreateRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


