# HostMuteSettings

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**End** | Pointer to **int64** | POSIX timestamp in seconds when the host is unmuted. If omitted, the host remains muted until explicitly unmuted. | [optional] 
**Message** | Pointer to **string** | Message to associate with the muting of this host. | [optional] 
**Override** | Pointer to **bool** | If true and the host is already muted, replaces existing host mute settings. | [optional] 

## Methods

### NewHostMuteSettings

`func NewHostMuteSettings() *HostMuteSettings`

NewHostMuteSettings instantiates a new HostMuteSettings object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewHostMuteSettingsWithDefaults

`func NewHostMuteSettingsWithDefaults() *HostMuteSettings`

NewHostMuteSettingsWithDefaults instantiates a new HostMuteSettings object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetEnd

`func (o *HostMuteSettings) GetEnd() int64`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *HostMuteSettings) GetEndOk() (*int64, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *HostMuteSettings) SetEnd(v int64)`

SetEnd sets End field to given value.

### HasEnd

`func (o *HostMuteSettings) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetMessage

`func (o *HostMuteSettings) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *HostMuteSettings) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *HostMuteSettings) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *HostMuteSettings) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetOverride

`func (o *HostMuteSettings) GetOverride() bool`

GetOverride returns the Override field if non-nil, zero value otherwise.

### GetOverrideOk

`func (o *HostMuteSettings) GetOverrideOk() (*bool, bool)`

GetOverrideOk returns a tuple with the Override field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverride

`func (o *HostMuteSettings) SetOverride(v bool)`

SetOverride sets Override field to given value.

### HasOverride

`func (o *HostMuteSettings) HasOverride() bool`

HasOverride returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


