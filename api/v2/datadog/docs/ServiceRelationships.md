# ServiceRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedBy** | Pointer to [**UserRelationship**](UserRelationship.md) |  | [optional] 
**LastModifiedBy** | Pointer to [**UserRelationship**](UserRelationship.md) |  | [optional] 

## Methods

### NewServiceRelationships

`func NewServiceRelationships() *ServiceRelationships`

NewServiceRelationships instantiates a new ServiceRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceRelationshipsWithDefaults

`func NewServiceRelationshipsWithDefaults() *ServiceRelationships`

NewServiceRelationshipsWithDefaults instantiates a new ServiceRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedBy

`func (o *ServiceRelationships) GetCreatedBy() UserRelationship`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ServiceRelationships) GetCreatedByOk() (*UserRelationship, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ServiceRelationships) SetCreatedBy(v UserRelationship)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ServiceRelationships) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetLastModifiedBy

`func (o *ServiceRelationships) GetLastModifiedBy() UserRelationship`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *ServiceRelationships) GetLastModifiedByOk() (*UserRelationship, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *ServiceRelationships) SetLastModifiedBy(v UserRelationship)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *ServiceRelationships) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


