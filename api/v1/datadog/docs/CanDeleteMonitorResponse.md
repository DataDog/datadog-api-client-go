# CanDeleteMonitorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**CanDeleteMonitorResponseData**](CanDeleteMonitorResponse_data.md) |  | 
**Errors** | Pointer to **map[string]string** | A mapping of Monitor ID to strings denoting where it&#39;s used. | [optional] 

## Methods

### GetData

`func (o *CanDeleteMonitorResponse) GetData() CanDeleteMonitorResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CanDeleteMonitorResponse) GetDataOk() (CanDeleteMonitorResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasData

`func (o *CanDeleteMonitorResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### SetData

`func (o *CanDeleteMonitorResponse) SetData(v CanDeleteMonitorResponseData)`

SetData gets a reference to the given CanDeleteMonitorResponseData and assigns it to the Data field.

### GetErrors

`func (o *CanDeleteMonitorResponse) GetErrors() map[string]string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *CanDeleteMonitorResponse) GetErrorsOk() (map[string]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasErrors

`func (o *CanDeleteMonitorResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### SetErrors

`func (o *CanDeleteMonitorResponse) SetErrors(v map[string]string)`

SetErrors gets a reference to the given map[string]string and assigns it to the Errors field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


