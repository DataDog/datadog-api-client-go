# TeamRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedBy** | Pointer to [**RelationshipToUser**](RelationshipToUser.md) |  | [optional] 
**LastModifiedBy** | Pointer to [**RelationshipToUser**](RelationshipToUser.md) |  | [optional] 

## Methods

### NewTeamRelationships

`func NewTeamRelationships() *TeamRelationships`

NewTeamRelationships instantiates a new TeamRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamRelationshipsWithDefaults

`func NewTeamRelationshipsWithDefaults() *TeamRelationships`

NewTeamRelationshipsWithDefaults instantiates a new TeamRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedBy

`func (o *TeamRelationships) GetCreatedBy() RelationshipToUser`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *TeamRelationships) GetCreatedByOk() (*RelationshipToUser, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *TeamRelationships) SetCreatedBy(v RelationshipToUser)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *TeamRelationships) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetLastModifiedBy

`func (o *TeamRelationships) GetLastModifiedBy() RelationshipToUser`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *TeamRelationships) GetLastModifiedByOk() (*RelationshipToUser, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *TeamRelationships) SetLastModifiedBy(v RelationshipToUser)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *TeamRelationships) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


