# ServiceLevelObjectivesBulkDeletedData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deleted** | Pointer to **[]string** | An array of service level objective object IDs that indicates which objects that were completely deleted. | [optional] 
**Updated** | Pointer to **[]string** | An array of service level objective object IDs that indicates which objects that were modified (objects for which at least one threshold was deleted, but that were not completely deleted). | [optional] 

## Methods

### GetDeleted

`func (o *ServiceLevelObjectivesBulkDeletedData) GetDeleted() []string`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *ServiceLevelObjectivesBulkDeletedData) GetDeletedOk() ([]string, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeleted

`func (o *ServiceLevelObjectivesBulkDeletedData) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### SetDeleted

`func (o *ServiceLevelObjectivesBulkDeletedData) SetDeleted(v []string)`

SetDeleted gets a reference to the given []string and assigns it to the Deleted field.

### GetUpdated

`func (o *ServiceLevelObjectivesBulkDeletedData) GetUpdated() []string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *ServiceLevelObjectivesBulkDeletedData) GetUpdatedOk() ([]string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUpdated

`func (o *ServiceLevelObjectivesBulkDeletedData) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### SetUpdated

`func (o *ServiceLevelObjectivesBulkDeletedData) SetUpdated(v []string)`

SetUpdated gets a reference to the given []string and assigns it to the Updated field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


