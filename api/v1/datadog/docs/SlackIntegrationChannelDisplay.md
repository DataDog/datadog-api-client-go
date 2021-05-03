# SlackIntegrationChannelDisplay

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Message** | Pointer to **bool** | Show the main body of the alert event. | [optional] [default to true]
**Notified** | Pointer to **bool** | Show the list of @-handles in the alert event. | [optional] [default to true]
**Snapshot** | Pointer to **bool** | Show the alert event&#39;s snapshot image. | [optional] [default to true]
**Tags** | Pointer to **bool** | Show the scopes on which the monitor alerted. | [optional] [default to true]

## Methods

### NewSlackIntegrationChannelDisplay

`func NewSlackIntegrationChannelDisplay() *SlackIntegrationChannelDisplay`

NewSlackIntegrationChannelDisplay instantiates a new SlackIntegrationChannelDisplay object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlackIntegrationChannelDisplayWithDefaults

`func NewSlackIntegrationChannelDisplayWithDefaults() *SlackIntegrationChannelDisplay`

NewSlackIntegrationChannelDisplayWithDefaults instantiates a new SlackIntegrationChannelDisplay object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *SlackIntegrationChannelDisplay) GetMessage() bool`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SlackIntegrationChannelDisplay) GetMessageOk() (*bool, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SlackIntegrationChannelDisplay) SetMessage(v bool)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SlackIntegrationChannelDisplay) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetNotified

`func (o *SlackIntegrationChannelDisplay) GetNotified() bool`

GetNotified returns the Notified field if non-nil, zero value otherwise.

### GetNotifiedOk

`func (o *SlackIntegrationChannelDisplay) GetNotifiedOk() (*bool, bool)`

GetNotifiedOk returns a tuple with the Notified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotified

`func (o *SlackIntegrationChannelDisplay) SetNotified(v bool)`

SetNotified sets Notified field to given value.

### HasNotified

`func (o *SlackIntegrationChannelDisplay) HasNotified() bool`

HasNotified returns a boolean if a field has been set.

### GetSnapshot

`func (o *SlackIntegrationChannelDisplay) GetSnapshot() bool`

GetSnapshot returns the Snapshot field if non-nil, zero value otherwise.

### GetSnapshotOk

`func (o *SlackIntegrationChannelDisplay) GetSnapshotOk() (*bool, bool)`

GetSnapshotOk returns a tuple with the Snapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshot

`func (o *SlackIntegrationChannelDisplay) SetSnapshot(v bool)`

SetSnapshot sets Snapshot field to given value.

### HasSnapshot

`func (o *SlackIntegrationChannelDisplay) HasSnapshot() bool`

HasSnapshot returns a boolean if a field has been set.

### GetTags

`func (o *SlackIntegrationChannelDisplay) GetTags() bool`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SlackIntegrationChannelDisplay) GetTagsOk() (*bool, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SlackIntegrationChannelDisplay) SetTags(v bool)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SlackIntegrationChannelDisplay) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


