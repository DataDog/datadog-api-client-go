# UsersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]User**](User.md) | Array of returned users. | [optional] 
**Included** | Pointer to [**[]UserResponseIncludedItem**](UserResponseIncludedItem.md) | Array of objects related to the users. | [optional] 
**Meta** | Pointer to [**ResponseMetaAttributes**](ResponseMetaAttributes.md) |  | [optional] 

## Methods

### NewUsersResponse

`func NewUsersResponse() *UsersResponse`

NewUsersResponse instantiates a new UsersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersResponseWithDefaults

`func NewUsersResponseWithDefaults() *UsersResponse`

NewUsersResponseWithDefaults instantiates a new UsersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *UsersResponse) GetData() []User`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UsersResponse) GetDataOk() (*[]User, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UsersResponse) SetData(v []User)`

SetData sets Data field to given value.

### HasData

`func (o *UsersResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetIncluded

`func (o *UsersResponse) GetIncluded() []UserResponseIncludedItem`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *UsersResponse) GetIncludedOk() (*[]UserResponseIncludedItem, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *UsersResponse) SetIncluded(v []UserResponseIncludedItem)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *UsersResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetMeta

`func (o *UsersResponse) GetMeta() ResponseMetaAttributes`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *UsersResponse) GetMetaOk() (*ResponseMetaAttributes, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *UsersResponse) SetMeta(v ResponseMetaAttributes)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *UsersResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


