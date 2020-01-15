# ServiceLevelObjectiveResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**ServiceLevelObjective**](ServiceLevelObjective.md) |  | 
**Errors** | Pointer to **[]string** | An array of error messages. Each endpoint documents how/whether this field is used. | [optional] 

## Methods

### GetData

`func (o *ServiceLevelObjectiveResponse) GetData() ServiceLevelObjective`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ServiceLevelObjectiveResponse) GetDataOk() (ServiceLevelObjective, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasData

`func (o *ServiceLevelObjectiveResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### SetData

`func (o *ServiceLevelObjectiveResponse) SetData(v ServiceLevelObjective)`

SetData gets a reference to the given ServiceLevelObjective and assigns it to the Data field.

### GetErrors

`func (o *ServiceLevelObjectiveResponse) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ServiceLevelObjectiveResponse) GetErrorsOk() ([]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasErrors

`func (o *ServiceLevelObjectiveResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### SetErrors

`func (o *ServiceLevelObjectiveResponse) SetErrors(v []string)`

SetErrors gets a reference to the given []string and assigns it to the Errors field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


