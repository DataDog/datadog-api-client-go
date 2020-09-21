# UserRelationshipData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A unique identifier that represents the user. | [optional] 
**Type** | Pointer to [**UsersType**](UsersType.md) |  | [optional] [default to "users"]

## Methods

### NewUserRelationshipData

`func NewUserRelationshipData() *UserRelationshipData`

NewUserRelationshipData instantiates a new UserRelationshipData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserRelationshipDataWithDefaults

`func NewUserRelationshipDataWithDefaults() *UserRelationshipData`

NewUserRelationshipDataWithDefaults instantiates a new UserRelationshipData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserRelationshipData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserRelationshipData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserRelationshipData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserRelationshipData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *UserRelationshipData) GetType() UsersType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserRelationshipData) GetTypeOk() (*UsersType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserRelationshipData) SetType(v UsersType)`

SetType sets Type field to given value.

### HasType

`func (o *UserRelationshipData) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


