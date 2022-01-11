# RelationshipToPermissions

## Properties

| Name     | Type                                                                             | Description                          | Notes      |
| -------- | -------------------------------------------------------------------------------- | ------------------------------------ | ---------- |
| **Data** | Pointer to [**[]RelationshipToPermissionData**](RelationshipToPermissionData.md) | Relationships to permission objects. | [optional] |

## Methods

### NewRelationshipToPermissions

`func NewRelationshipToPermissions() *RelationshipToPermissions`

NewRelationshipToPermissions instantiates a new RelationshipToPermissions object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewRelationshipToPermissionsWithDefaults

`func NewRelationshipToPermissionsWithDefaults() *RelationshipToPermissions`

NewRelationshipToPermissionsWithDefaults instantiates a new RelationshipToPermissions object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetData

`func (o *RelationshipToPermissions) GetData() []RelationshipToPermissionData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RelationshipToPermissions) GetDataOk() (*[]RelationshipToPermissionData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RelationshipToPermissions) SetData(v []RelationshipToPermissionData)`

SetData sets Data field to given value.

### HasData

`func (o *RelationshipToPermissions) HasData() bool`

HasData returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
