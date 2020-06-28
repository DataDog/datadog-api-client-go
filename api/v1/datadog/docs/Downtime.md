# Downtime

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | **bool** | If a scheduled downtime currently exists. | [optional] [readonly] 
**Canceled** | Pointer to **int64** | If a scheduled downtime is canceled. | [optional] [readonly] 
**CreatorId** | **int32** | User ID of the downtime creator. | [optional] [readonly] 
**Disabled** | **bool** | If a downtime has been disabled. | [optional] 
**DowntimeType** | **int32** | &#x60;0&#x60; for a downtime applied on &#x60;*&#x60; or all, &#x60;1&#x60; when the downtime is only scoped to hosts, or &#x60;2&#x60; when the downtime is scoped to anything but hosts. | [optional] [readonly] 
**End** | Pointer to **int64** | POSIX timestamp to end the downtime. If not provided, the downtime is in effect indefinitely until you cancel it. | [optional] 
**Id** | **int64** | The downtime ID. | [optional] [readonly] 
**Message** | **string** | A message to include with notifications for this downtime. Email notifications can be sent to specific users by using the same &#x60;@username&#x60; notation as events. | [optional] 
**MonitorId** | Pointer to **int64** | A single monitor to which the downtime applies. If not provided, the downtime applies to all monitors. | [optional] 
**MonitorTags** | **[]string** | A comma-separated list of monitor tags. For example, tags that are applied directly to monitors, not tags that are used in monitor queries (which are filtered by the scope parameter), to which the downtime applies. The resulting downtime applies to monitors that match ALL provided monitor tags. For example, &#x60;service:postgres&#x60; **AND** &#x60;team:frontend&#x60;. | [optional] 
**ParentId** | Pointer to **int64** | ID of the parent Downtime. | [optional] 
**Recurrence** | Pointer to [**DowntimeRecurrence**](DowntimeRecurrence.md) |  | [optional] 
**Scope** | **[]string** | The scope(s) to which the downtime applies. For example, &#x60;host:app2&#x60;. Provide multiple scopes as a comma-separated list like &#x60;env:dev,env:prod&#x60;. The resulting downtime applies to sources that matches ALL provided scopes (&#x60;env:dev&#x60; **AND** &#x60;env:prod&#x60;). | [optional] 
**Start** | **int64** | POSIX timestamp to start the downtime. If not provided, the downtime starts the moment it is created. | [optional] 
**Timezone** | **string** | The timezone for the downtime. | [optional] 
**UpdaterId** | Pointer to **int32** | ID of the last user that updated the downtime. | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


