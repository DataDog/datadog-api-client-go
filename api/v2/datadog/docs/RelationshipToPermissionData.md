# RelationshipToPermissionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the permission. | [optional] 
**Type** | Pointer to [**PermissionsType**](PermissionsType.md) |  | [optional] [default to "permissions"]

## Methods

### NewRelationshipToPermissionData

`func NewRelationshipToPermissionData() *RelationshipToPermissionData`

NewRelationshipToPermissionData instantiates a new RelationshipToPermissionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelationshipToPermissionDataWithDefaults

`func NewRelationshipToPermissionDataWithDefaults() *RelationshipToPermissionData`

NewRelationshipToPermissionDataWithDefaults instantiates a new RelationshipToPermissionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RelationshipToPermissionData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RelationshipToPermissionData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RelationshipToPermissionData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RelationshipToPermissionData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *RelationshipToPermissionData) GetType() PermissionsType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RelationshipToPermissionData) GetTypeOk() (*PermissionsType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RelationshipToPermissionData) SetType(v PermissionsType)`

SetType sets Type field to given value.

### HasType

`func (o *RelationshipToPermissionData) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


