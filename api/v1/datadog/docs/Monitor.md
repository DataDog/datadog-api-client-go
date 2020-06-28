# Monitor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | [**time.Time**](time.Time.md) | Timestamp of the monitor creation. | [optional] [readonly] 
**Creator** | [**Creator**](Creator.md) |  | [optional] 
**Deleted** | Pointer to [**time.Time**](time.Time.md) | Whether or not the monitor is deleted. (Always &#x60;null&#x60;) | [optional] [readonly] 
**Id** | **int64** | ID of this monitor. | [optional] [readonly] 
**Message** | **string** | A message to include with notifications for this monitor. | [optional] 
**Modified** | [**time.Time**](time.Time.md) | Last timestamp when the monitor was edited. | [optional] [readonly] 
**Multi** | **bool** | Whether or not the monitor is broken down on different groups. | [optional] [readonly] 
**Name** | **string** | The monitor name. | [optional] 
**Options** | [**MonitorOptions**](MonitorOptions.md) |  | [optional] 
**OverallState** | [**MonitorOverallStates**](MonitorOverallStates.md) |  | [optional] 
**Query** | **string** | The monitor query. | [optional] 
**State** | [**MonitorState**](MonitorState.md) |  | [optional] 
**Tags** | **[]string** | Tags associated to your monitor. | [optional] 
**Type** | [**MonitorType**](MonitorType.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


