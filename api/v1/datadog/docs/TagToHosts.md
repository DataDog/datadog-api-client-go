# TagToHosts

## Properties

| Name     | Type                               | Description                          | Notes      |
| -------- | ---------------------------------- | ------------------------------------ | ---------- |
| **Tags** | Pointer to **map[string][]string** | A list of tags to apply to the host. | [optional] |

## Methods

### NewTagToHosts

`func NewTagToHosts() *TagToHosts`

NewTagToHosts instantiates a new TagToHosts object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewTagToHostsWithDefaults

`func NewTagToHostsWithDefaults() *TagToHosts`

NewTagToHostsWithDefaults instantiates a new TagToHosts object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetTags

`func (o *TagToHosts) GetTags() map[string][]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TagToHosts) GetTagsOk() (*map[string][]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TagToHosts) SetTags(v map[string][]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *TagToHosts) HasTags() bool`

HasTags returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
