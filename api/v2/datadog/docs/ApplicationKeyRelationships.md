# ApplicationKeyRelationships

## Properties

| Name        | Type                                                       | Description | Notes      |
| ----------- | ---------------------------------------------------------- | ----------- | ---------- |
| **OwnedBy** | Pointer to [**RelationshipToUser**](RelationshipToUser.md) |             | [optional] |

## Methods

### NewApplicationKeyRelationships

`func NewApplicationKeyRelationships() *ApplicationKeyRelationships`

NewApplicationKeyRelationships instantiates a new ApplicationKeyRelationships object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewApplicationKeyRelationshipsWithDefaults

`func NewApplicationKeyRelationshipsWithDefaults() *ApplicationKeyRelationships`

NewApplicationKeyRelationshipsWithDefaults instantiates a new ApplicationKeyRelationships object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetOwnedBy

`func (o *ApplicationKeyRelationships) GetOwnedBy() RelationshipToUser`

GetOwnedBy returns the OwnedBy field if non-nil, zero value otherwise.

### GetOwnedByOk

`func (o *ApplicationKeyRelationships) GetOwnedByOk() (*RelationshipToUser, bool)`

GetOwnedByOk returns a tuple with the OwnedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnedBy

`func (o *ApplicationKeyRelationships) SetOwnedBy(v RelationshipToUser)`

SetOwnedBy sets OwnedBy field to given value.

### HasOwnedBy

`func (o *ApplicationKeyRelationships) HasOwnedBy() bool`

HasOwnedBy returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
