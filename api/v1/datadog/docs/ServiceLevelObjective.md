# ServiceLevelObjective

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **int64** | Creation timestamp (UNIX time in seconds)  Always included in service level objective responses. | [optional] [readonly] 
**Creator** | [**Creator**](Creator.md) |  | [optional] 
**Description** | Pointer to **string** | A user-defined description of the service level objective.  Always included in service level objective responses (but may be &#x60;null&#x60;). Optional in create/update requests. | [optional] 
**Groups** | **[]string** | A list of (up to 20) monitor groups that narrow the scope of a monitor service level objective.  Included in service level objective responses if it is not empty. Optional in create/update requests for monitor service level objectives, but may only be used when then length of the &#x60;monitor_ids&#x60; field is one. | [optional] 
**Id** | **string** | A unique identifier for the service level objective object.  Always included in service level objective responses. | [optional] [readonly] 
**ModifiedAt** | **int64** | Modification timestamp (UNIX time in seconds)  Always included in service level objective responses. | [optional] [readonly] 
**MonitorIds** | **[]int64** | A list of monitor ids that defines the scope of a monitor service level objective. **Required if type is &#x60;monitor&#x60;**. | [optional] 
**MonitorTags** | **[]string** | The union of monitor tags for all monitors referenced by the &#x60;monitor_ids&#x60; field. Always included in service level objective responses for monitor service level objectives (but may be empty). Ignored in create/update requests. Does not affect which monitors are included in the service level objective (that is determined entirely by the &#x60;monitor_ids&#x60; field). | [optional] 
**Name** | **string** | The name of the service level objective object. | 
**Query** | [**ServiceLevelObjectiveQuery**](ServiceLevelObjectiveQuery.md) |  | [optional] 
**Tags** | **[]string** | A list of tags associated with this service level objective. Always included in service level objective responses (but may be empty). Optional in create/update requests. | [optional] 
**Thresholds** | [**[]SloThreshold**](SLOThreshold.md) | The thresholds (timeframes and associated targets) for this service level objective object. | 
**Type** | [**SloType**](SLOType.md) |  | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


