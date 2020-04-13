# UserResponsePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**UserResponse**](UserResponse.md) |  | [optional] 
**Included** | Pointer to [**[]UserResponseIncludedItem**](UserResponseIncludedItem.md) | Array of objects related to the user. | [optional] 

## Methods

### NewUserResponsePayload

`func NewUserResponsePayload() *UserResponsePayload`

NewUserResponsePayload instantiates a new UserResponsePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserResponsePayloadWithDefaults

`func NewUserResponsePayloadWithDefaults() *UserResponsePayload`

NewUserResponsePayloadWithDefaults instantiates a new UserResponsePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *UserResponsePayload) GetData() UserResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UserResponsePayload) GetDataOk() (*UserResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UserResponsePayload) SetData(v UserResponse)`

SetData sets Data field to given value.

### HasData

`func (o *UserResponsePayload) HasData() bool`

HasData returns a boolean if a field has been set.

### GetIncluded

`func (o *UserResponsePayload) GetIncluded() []UserResponseIncludedItem`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *UserResponsePayload) GetIncludedOk() (*[]UserResponseIncludedItem, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *UserResponsePayload) SetIncluded(v []UserResponseIncludedItem)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *UserResponsePayload) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


