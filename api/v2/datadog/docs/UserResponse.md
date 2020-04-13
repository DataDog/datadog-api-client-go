# UserResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**UserResponseAttributes**](UserResponseAttributes.md) |  | [optional] 
**Id** | Pointer to **string** | ID of the user. | [optional] 
**Relationships** | Pointer to [**UserResponseRelationships**](UserResponseRelationships.md) |  | [optional] 
**Type** | Pointer to **string** | Users type. | [optional] [readonly] [default to "users"]

## Methods

### NewUserResponse

`func NewUserResponse() *UserResponse`

NewUserResponse instantiates a new UserResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserResponseWithDefaults

`func NewUserResponseWithDefaults() *UserResponse`

NewUserResponseWithDefaults instantiates a new UserResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *UserResponse) GetAttributes() UserResponseAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *UserResponse) GetAttributesOk() (*UserResponseAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *UserResponse) SetAttributes(v UserResponseAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *UserResponse) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *UserResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRelationships

`func (o *UserResponse) GetRelationships() UserResponseRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *UserResponse) GetRelationshipsOk() (*UserResponseRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *UserResponse) SetRelationships(v UserResponseRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *UserResponse) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetType

`func (o *UserResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UserResponse) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


