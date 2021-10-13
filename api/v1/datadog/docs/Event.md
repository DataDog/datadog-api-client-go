# Event

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**AlertType** | Pointer to [**EventAlertType**](EventAlertType.md) |  | [optional] 
**DateHappened** | Pointer to **int64** | POSIX timestamp of the event. Must be sent as an integer (i.e. no quotes). Limited to events no older than 7 days. | [optional] 
**DeviceName** | Pointer to **string** | A device name. | [optional] 
**Host** | Pointer to **string** | Host name to associate with the event. Any tags associated with the host are also applied to this event. | [optional] 
**Id** | Pointer to **int64** | Integer ID of the event. | [optional] [readonly] 
**IdStr** | Pointer to **string** | Handling IDs as large 64-bit numbers can cause loss of accuracy issues with some programming languages. Instead, use the string representation of the Event ID to avoid losing accuracy. | [optional] [readonly] 
**Payload** | Pointer to **string** | Payload of the event. | [optional] [readonly] 
**Priority** | Pointer to [**EventPriority**](EventPriority.md) |  | [optional] 
**SourceTypeName** | Pointer to **string** | The type of event being posted. Option examples include nagios, hudson, jenkins, my_apps, chef, puppet, git, bitbucket, etc. A complete list of source attribute values [available here](https://docs.datadoghq.com/integrations/faq/list-of-api-source-attribute-value). | [optional] 
**Tags** | Pointer to **[]string** | A list of tags to apply to the event. | [optional] 
**Text** | Pointer to **string** | The body of the event. Limited to 4000 characters. The text supports markdown. To use markdown in the event text, start the text block with &#x60;%%% \\n&#x60; and end the text block with &#x60;\\n %%%&#x60;. Use &#x60;msg_text&#x60; with the Datadog Ruby library. | [optional] 
**Title** | Pointer to **string** | The event title. | [optional] 
**Url** | Pointer to **string** | URL of the event. | [optional] [readonly] 

## Methods

### NewEvent

`func NewEvent() *Event`

NewEvent instantiates a new Event object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewEventWithDefaults

`func NewEventWithDefaults() *Event`

NewEventWithDefaults instantiates a new Event object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAlertType

`func (o *Event) GetAlertType() EventAlertType`

GetAlertType returns the AlertType field if non-nil, zero value otherwise.

### GetAlertTypeOk

`func (o *Event) GetAlertTypeOk() (*EventAlertType, bool)`

GetAlertTypeOk returns a tuple with the AlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertType

`func (o *Event) SetAlertType(v EventAlertType)`

SetAlertType sets AlertType field to given value.

### HasAlertType

`func (o *Event) HasAlertType() bool`

HasAlertType returns a boolean if a field has been set.

### GetDateHappened

`func (o *Event) GetDateHappened() int64`

GetDateHappened returns the DateHappened field if non-nil, zero value otherwise.

### GetDateHappenedOk

`func (o *Event) GetDateHappenedOk() (*int64, bool)`

GetDateHappenedOk returns a tuple with the DateHappened field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateHappened

`func (o *Event) SetDateHappened(v int64)`

SetDateHappened sets DateHappened field to given value.

### HasDateHappened

`func (o *Event) HasDateHappened() bool`

HasDateHappened returns a boolean if a field has been set.

### GetDeviceName

`func (o *Event) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *Event) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *Event) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *Event) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetHost

`func (o *Event) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *Event) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *Event) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *Event) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *Event) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Event) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Event) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Event) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdStr

`func (o *Event) GetIdStr() string`

GetIdStr returns the IdStr field if non-nil, zero value otherwise.

### GetIdStrOk

`func (o *Event) GetIdStrOk() (*string, bool)`

GetIdStrOk returns a tuple with the IdStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdStr

`func (o *Event) SetIdStr(v string)`

SetIdStr sets IdStr field to given value.

### HasIdStr

`func (o *Event) HasIdStr() bool`

HasIdStr returns a boolean if a field has been set.

### GetPayload

`func (o *Event) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *Event) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *Event) SetPayload(v string)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *Event) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetPriority

`func (o *Event) GetPriority() EventPriority`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *Event) GetPriorityOk() (*EventPriority, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *Event) SetPriority(v EventPriority)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *Event) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetSourceTypeName

`func (o *Event) GetSourceTypeName() string`

GetSourceTypeName returns the SourceTypeName field if non-nil, zero value otherwise.

### GetSourceTypeNameOk

`func (o *Event) GetSourceTypeNameOk() (*string, bool)`

GetSourceTypeNameOk returns a tuple with the SourceTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTypeName

`func (o *Event) SetSourceTypeName(v string)`

SetSourceTypeName sets SourceTypeName field to given value.

### HasSourceTypeName

`func (o *Event) HasSourceTypeName() bool`

HasSourceTypeName returns a boolean if a field has been set.

### GetTags

`func (o *Event) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Event) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Event) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Event) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetText

`func (o *Event) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *Event) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *Event) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *Event) HasText() bool`

HasText returns a boolean if a field has been set.

### GetTitle

`func (o *Event) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Event) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Event) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Event) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUrl

`func (o *Event) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Event) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Event) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Event) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


