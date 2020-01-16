# HostMuteSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | Pointer to **int64** | POSIX timestamp in seconds when the host is unmuted. If omitted, the host remains muted until explicitly unmuted. | [optional] 
**Message** | Pointer to **string** | Message to associate with the muting of this host. | [optional] 
**Override** | Pointer to **bool** | If true and the host is already muted, replaces existing host mute settings. | [optional] 

## Methods

### GetEnd

`func (o *HostMuteSettings) GetEnd() int64`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *HostMuteSettings) GetEndOk() (int64, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnd

`func (o *HostMuteSettings) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### SetEnd

`func (o *HostMuteSettings) SetEnd(v int64)`

SetEnd gets a reference to the given int64 and assigns it to the End field.

### GetMessage

`func (o *HostMuteSettings) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *HostMuteSettings) GetMessageOk() (string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMessage

`func (o *HostMuteSettings) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessage

`func (o *HostMuteSettings) SetMessage(v string)`

SetMessage gets a reference to the given string and assigns it to the Message field.

### GetOverride

`func (o *HostMuteSettings) GetOverride() bool`

GetOverride returns the Override field if non-nil, zero value otherwise.

### GetOverrideOk

`func (o *HostMuteSettings) GetOverrideOk() (bool, bool)`

GetOverrideOk returns a tuple with the Override field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOverride

`func (o *HostMuteSettings) HasOverride() bool`

HasOverride returns a boolean if a field has been set.

### SetOverride

`func (o *HostMuteSettings) SetOverride(v bool)`

SetOverride gets a reference to the given bool and assigns it to the Override field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


