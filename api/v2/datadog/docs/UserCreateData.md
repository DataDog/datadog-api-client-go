# UserCreateData

## Properties

| Name              | Type                                                     | Description | Notes                        |
| ----------------- | -------------------------------------------------------- | ----------- | ---------------------------- |
| **Attributes**    | [**UserCreateAttributes**](UserCreateAttributes.md)      |             |
| **Relationships** | Pointer to [**UserRelationships**](UserRelationships.md) |             | [optional]                   |
| **Type**          | [**UsersType**](UsersType.md)                            |             | [default to USERSTYPE_USERS] |

## Methods

### NewUserCreateData

`func NewUserCreateData(attributes UserCreateAttributes, type_ UsersType) *UserCreateData`

NewUserCreateData instantiates a new UserCreateData object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewUserCreateDataWithDefaults

`func NewUserCreateDataWithDefaults() *UserCreateData`

NewUserCreateDataWithDefaults instantiates a new UserCreateData object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAttributes

`func (o *UserCreateData) GetAttributes() UserCreateAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *UserCreateData) GetAttributesOk() (*UserCreateAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *UserCreateData) SetAttributes(v UserCreateAttributes)`

SetAttributes sets Attributes field to given value.

### GetRelationships

`func (o *UserCreateData) GetRelationships() UserRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *UserCreateData) GetRelationshipsOk() (*UserRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *UserCreateData) SetRelationships(v UserRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *UserCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetType

`func (o *UserCreateData) GetType() UsersType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserCreateData) GetTypeOk() (*UsersType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserCreateData) SetType(v UsersType)`

SetType sets Type field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
