# MonitorState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Groups** | Pointer to [**[]MonitorStateGroup**](MonitorStateGroup.md) |  | [optional] 
**MonitorId** | Pointer to **int32** |  | [optional] 
**OverallState** | Pointer to [**MonitorOverallStates**](MonitorOverallStates.md) |  | [optional] 

## Methods

### GetGroups

`func (o *MonitorState) GetGroups() []MonitorStateGroup`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *MonitorState) GetGroupsOk() ([]MonitorStateGroup, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasGroups

`func (o *MonitorState) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### SetGroups

`func (o *MonitorState) SetGroups(v []MonitorStateGroup)`

SetGroups gets a reference to the given []MonitorStateGroup and assigns it to the Groups field.

### GetMonitorId

`func (o *MonitorState) GetMonitorId() int32`

GetMonitorId returns the MonitorId field if non-nil, zero value otherwise.

### GetMonitorIdOk

`func (o *MonitorState) GetMonitorIdOk() (int32, bool)`

GetMonitorIdOk returns a tuple with the MonitorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMonitorId

`func (o *MonitorState) HasMonitorId() bool`

HasMonitorId returns a boolean if a field has been set.

### SetMonitorId

`func (o *MonitorState) SetMonitorId(v int32)`

SetMonitorId gets a reference to the given int32 and assigns it to the MonitorId field.

### GetOverallState

`func (o *MonitorState) GetOverallState() MonitorOverallStates`

GetOverallState returns the OverallState field if non-nil, zero value otherwise.

### GetOverallStateOk

`func (o *MonitorState) GetOverallStateOk() (MonitorOverallStates, bool)`

GetOverallStateOk returns a tuple with the OverallState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOverallState

`func (o *MonitorState) HasOverallState() bool`

HasOverallState returns a boolean if a field has been set.

### SetOverallState

`func (o *MonitorState) SetOverallState(v MonitorOverallStates)`

SetOverallState gets a reference to the given MonitorOverallStates and assigns it to the OverallState field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


