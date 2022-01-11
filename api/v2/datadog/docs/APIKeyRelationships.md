# APIKeyRelationships

## Properties

| Name           | Type                                                       | Description | Notes      |
| -------------- | ---------------------------------------------------------- | ----------- | ---------- |
| **CreatedBy**  | Pointer to [**RelationshipToUser**](RelationshipToUser.md) |             | [optional] |
| **ModifiedBy** | Pointer to [**RelationshipToUser**](RelationshipToUser.md) |             | [optional] |

## Methods

### NewAPIKeyRelationships

`func NewAPIKeyRelationships() *APIKeyRelationships`

NewAPIKeyRelationships instantiates a new APIKeyRelationships object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewAPIKeyRelationshipsWithDefaults

`func NewAPIKeyRelationshipsWithDefaults() *APIKeyRelationships`

NewAPIKeyRelationshipsWithDefaults instantiates a new APIKeyRelationships object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetCreatedBy

`func (o *APIKeyRelationships) GetCreatedBy() RelationshipToUser`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *APIKeyRelationships) GetCreatedByOk() (*RelationshipToUser, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *APIKeyRelationships) SetCreatedBy(v RelationshipToUser)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *APIKeyRelationships) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetModifiedBy

`func (o *APIKeyRelationships) GetModifiedBy() RelationshipToUser`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *APIKeyRelationships) GetModifiedByOk() (*RelationshipToUser, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *APIKeyRelationships) SetModifiedBy(v RelationshipToUser)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *APIKeyRelationships) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
