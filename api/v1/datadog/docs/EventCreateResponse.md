# EventCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertType** | Pointer to [**EventAlertType**](EventAlertType.md) |  | [optional] 
**DateHappened** | Pointer to **int64** | POSIX timestamp of the event. Must be sent as an integer (i.e. no quotes). Limited to events no older than 7 days. | [optional] 
**DeviceName** | Pointer to **string** | A device name. | [optional] 
**Host** | Pointer to **string** | Host name to associate with the event. Any tags associated with the host are also applied to this event. | [optional] 
**Id** | Pointer to **int64** | Integer ID of the event. | [optional] [readonly] 
**Payload** | Pointer to **string** | Payload of the event. | [optional] [readonly] 
**Priority** | Pointer to [**EventPriority**](EventPriority.md) |  | [optional] 
**RelatedEventId** | Pointer to **int64** | ID of the parent event. Must be sent as an integer (i.e. no quotes). | [optional] 
**SourceTypeName** | Pointer to **string** | The type of event being posted. Option examples include nagios, hudson, jenkins, my_apps, chef, puppet, git, bitbucket, etc. A complete list of source attribute values [available here](https://docs.datadoghq.com/integrations/faq/list-of-api-source-attribute-value). | [optional] 
**Status** | Pointer to **string** | A status. | [optional] 
**Tags** | Pointer to **[]string** | A list of tags to apply to the event. | [optional] 
**Text** | Pointer to **string** | The body of the event. Limited to 4000 characters. The text supports markdown. Use &#x60;msg_text&#x60; with the Datadog Ruby library. | [optional] 
**Title** | Pointer to **string** | The event title. Limited to 100 characters. Use &#x60;msg_title&#x60; with the Datadog Ruby library. | [optional] 
**Url** | Pointer to **string** | URL of the event. | [optional] [readonly] 

## Methods

### NewEventCreateResponse

`func NewEventCreateResponse() *EventCreateResponse`

NewEventCreateResponse instantiates a new EventCreateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventCreateResponseWithDefaults

`func NewEventCreateResponseWithDefaults() *EventCreateResponse`

NewEventCreateResponseWithDefaults instantiates a new EventCreateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertType

`func (o *EventCreateResponse) GetAlertType() EventAlertType`

GetAlertType returns the AlertType field if non-nil, zero value otherwise.

### GetAlertTypeOk

`func (o *EventCreateResponse) GetAlertTypeOk() (*EventAlertType, bool)`

GetAlertTypeOk returns a tuple with the AlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertType

`func (o *EventCreateResponse) SetAlertType(v EventAlertType)`

SetAlertType sets AlertType field to given value.

### HasAlertType

`func (o *EventCreateResponse) HasAlertType() bool`

HasAlertType returns a boolean if a field has been set.

### GetDateHappened

`func (o *EventCreateResponse) GetDateHappened() int64`

GetDateHappened returns the DateHappened field if non-nil, zero value otherwise.

### GetDateHappenedOk

`func (o *EventCreateResponse) GetDateHappenedOk() (*int64, bool)`

GetDateHappenedOk returns a tuple with the DateHappened field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateHappened

`func (o *EventCreateResponse) SetDateHappened(v int64)`

SetDateHappened sets DateHappened field to given value.

### HasDateHappened

`func (o *EventCreateResponse) HasDateHappened() bool`

HasDateHappened returns a boolean if a field has been set.

### GetDeviceName

`func (o *EventCreateResponse) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *EventCreateResponse) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *EventCreateResponse) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *EventCreateResponse) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetHost

`func (o *EventCreateResponse) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *EventCreateResponse) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *EventCreateResponse) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *EventCreateResponse) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *EventCreateResponse) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventCreateResponse) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EventCreateResponse) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *EventCreateResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPayload

`func (o *EventCreateResponse) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *EventCreateResponse) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *EventCreateResponse) SetPayload(v string)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *EventCreateResponse) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetPriority

`func (o *EventCreateResponse) GetPriority() EventPriority`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *EventCreateResponse) GetPriorityOk() (*EventPriority, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *EventCreateResponse) SetPriority(v EventPriority)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *EventCreateResponse) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetRelatedEventId

`func (o *EventCreateResponse) GetRelatedEventId() int64`

GetRelatedEventId returns the RelatedEventId field if non-nil, zero value otherwise.

### GetRelatedEventIdOk

`func (o *EventCreateResponse) GetRelatedEventIdOk() (*int64, bool)`

GetRelatedEventIdOk returns a tuple with the RelatedEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedEventId

`func (o *EventCreateResponse) SetRelatedEventId(v int64)`

SetRelatedEventId sets RelatedEventId field to given value.

### HasRelatedEventId

`func (o *EventCreateResponse) HasRelatedEventId() bool`

HasRelatedEventId returns a boolean if a field has been set.

### GetSourceTypeName

`func (o *EventCreateResponse) GetSourceTypeName() string`

GetSourceTypeName returns the SourceTypeName field if non-nil, zero value otherwise.

### GetSourceTypeNameOk

`func (o *EventCreateResponse) GetSourceTypeNameOk() (*string, bool)`

GetSourceTypeNameOk returns a tuple with the SourceTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTypeName

`func (o *EventCreateResponse) SetSourceTypeName(v string)`

SetSourceTypeName sets SourceTypeName field to given value.

### HasSourceTypeName

`func (o *EventCreateResponse) HasSourceTypeName() bool`

HasSourceTypeName returns a boolean if a field has been set.

### GetStatus

`func (o *EventCreateResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EventCreateResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EventCreateResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EventCreateResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *EventCreateResponse) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *EventCreateResponse) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *EventCreateResponse) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *EventCreateResponse) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetText

`func (o *EventCreateResponse) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *EventCreateResponse) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *EventCreateResponse) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *EventCreateResponse) HasText() bool`

HasText returns a boolean if a field has been set.

### GetTitle

`func (o *EventCreateResponse) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *EventCreateResponse) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *EventCreateResponse) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *EventCreateResponse) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUrl

`func (o *EventCreateResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *EventCreateResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *EventCreateResponse) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *EventCreateResponse) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


