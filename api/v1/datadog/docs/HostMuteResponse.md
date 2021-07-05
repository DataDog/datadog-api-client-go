# HostMuteResponse

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Action** | Pointer to **string** | Action applied to the hosts. | [optional] 
**End** | Pointer to **int64** | POSIX timestamp in seconds when the host is unmuted. | [optional] 
**Hostname** | Pointer to **string** | The host name. | [optional] 
**Message** | Pointer to **string** | Message associated with the mute. | [optional] 

## Methods

### NewHostMuteResponse

`func NewHostMuteResponse() *HostMuteResponse`

NewHostMuteResponse instantiates a new HostMuteResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewHostMuteResponseWithDefaults

`func NewHostMuteResponseWithDefaults() *HostMuteResponse`

NewHostMuteResponseWithDefaults instantiates a new HostMuteResponse object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAction

`func (o *HostMuteResponse) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *HostMuteResponse) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *HostMuteResponse) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *HostMuteResponse) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetEnd

`func (o *HostMuteResponse) GetEnd() int64`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *HostMuteResponse) GetEndOk() (*int64, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *HostMuteResponse) SetEnd(v int64)`

SetEnd sets End field to given value.

### HasEnd

`func (o *HostMuteResponse) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetHostname

`func (o *HostMuteResponse) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *HostMuteResponse) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *HostMuteResponse) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *HostMuteResponse) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetMessage

`func (o *HostMuteResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *HostMuteResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *HostMuteResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *HostMuteResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


