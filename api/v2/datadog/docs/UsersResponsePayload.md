# UsersResponsePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]UserResponse**](UserResponse.md) | Array of returned users. | [optional] 
**Included** | Pointer to [**[]UserResponseIncludedItem**](UserResponseIncludedItem.md) | Array of objects related to the users. | [optional] 
**Meta** | Pointer to [**ResponseMetaAttributes**](ResponseMetaAttributes.md) |  | [optional] 

## Methods

### NewUsersResponsePayload

`func NewUsersResponsePayload() *UsersResponsePayload`

NewUsersResponsePayload instantiates a new UsersResponsePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersResponsePayloadWithDefaults

`func NewUsersResponsePayloadWithDefaults() *UsersResponsePayload`

NewUsersResponsePayloadWithDefaults instantiates a new UsersResponsePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *UsersResponsePayload) GetData() []UserResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UsersResponsePayload) GetDataOk() (*[]UserResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UsersResponsePayload) SetData(v []UserResponse)`

SetData sets Data field to given value.

### HasData

`func (o *UsersResponsePayload) HasData() bool`

HasData returns a boolean if a field has been set.

### GetIncluded

`func (o *UsersResponsePayload) GetIncluded() []UserResponseIncludedItem`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *UsersResponsePayload) GetIncludedOk() (*[]UserResponseIncludedItem, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *UsersResponsePayload) SetIncluded(v []UserResponseIncludedItem)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *UsersResponsePayload) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetMeta

`func (o *UsersResponsePayload) GetMeta() ResponseMetaAttributes`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *UsersResponsePayload) GetMetaOk() (*ResponseMetaAttributes, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *UsersResponsePayload) SetMeta(v ResponseMetaAttributes)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *UsersResponsePayload) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


