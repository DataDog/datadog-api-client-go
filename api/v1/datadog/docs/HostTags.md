# HostTags

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Host** | Pointer to **string** | Your host name. | [optional] 
**Tags** | Pointer to **[]string** | A list of tags to apply to the host. | [optional] 

## Methods

### NewHostTags

`func NewHostTags() *HostTags`

NewHostTags instantiates a new HostTags object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewHostTagsWithDefaults

`func NewHostTagsWithDefaults() *HostTags`

NewHostTagsWithDefaults instantiates a new HostTags object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetHost

`func (o *HostTags) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *HostTags) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *HostTags) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *HostTags) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetTags

`func (o *HostTags) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *HostTags) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *HostTags) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *HostTags) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


