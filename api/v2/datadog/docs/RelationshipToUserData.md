# RelationshipToUserData

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Id** | **string** | A unique identifier that represents the user. | 
**Type** | [**UsersType**](UsersType.md) |  | [default to USERSTYPE_USERS]

## Methods

### NewRelationshipToUserData

`func NewRelationshipToUserData(id string, type_ UsersType) *RelationshipToUserData`

NewRelationshipToUserData instantiates a new RelationshipToUserData object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewRelationshipToUserDataWithDefaults

`func NewRelationshipToUserDataWithDefaults() *RelationshipToUserData`

NewRelationshipToUserDataWithDefaults instantiates a new RelationshipToUserData object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetId

`func (o *RelationshipToUserData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RelationshipToUserData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RelationshipToUserData) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *RelationshipToUserData) GetType() UsersType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RelationshipToUserData) GetTypeOk() (*UsersType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RelationshipToUserData) SetType(v UsersType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


