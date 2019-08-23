# \MonitorsApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMonitor**](MonitorsApi.md#CreateMonitor) | **Post** /api/v1/monitor | Create a new Monitor
[**DeleteMonitor**](MonitorsApi.md#DeleteMonitor) | **Delete** /api/v1/monitor/{monitor_id} | Delete the specified monitor.
[**EditMonitor**](MonitorsApi.md#EditMonitor) | **Put** /api/v1/monitor/{monitor_id} | Edit the specified monitor
[**GetAllMonitors**](MonitorsApi.md#GetAllMonitors) | **Get** /api/v1/monitor | Get details about the specified monitor.
[**GetMonitor**](MonitorsApi.md#GetMonitor) | **Get** /api/v1/monitor/{monitor_id} | Get details about the specified monitor.
[**ValidateMonitor**](MonitorsApi.md#ValidateMonitor) | **Post** /api/v1/monitor/validate | 



## CreateMonitor

> Monitor CreateMonitor(ctx, monitor)
Create a new Monitor

### Overview Create a monitor using the specified options ### Arguments * **`Monitor`** [*required*] The Monitor Object to create

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitor** | [**Monitor**](Monitor.md)| Monitor request object | 

### Return type

[**Monitor**](Monitor.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMonitor

> map[string]int64 DeleteMonitor(ctx, monitorId)
Delete the specified monitor.

### Overview Delete the specified monitor ### Arguments * **`monitor_id`** [*required*]: The id of the monitor.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorId** | **int64**| The id of the monitor | 

### Return type

**map[string]int64**

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditMonitor

> Monitor EditMonitor(ctx, monitorId, monitor)
Edit the specified monitor

### Overview Edit the specified monitor. ### Arguments * **`monitor_id`** [*required*]: The id of the monitor.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorId** | **int64**| The id of the monitor | 
**monitor** | [**Monitor**](Monitor.md)| Monitor request object | 

### Return type

[**Monitor**](Monitor.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllMonitors

> []Monitor GetAllMonitors(ctx, optional)
Get details about the specified monitor.

### Overview Get details about the specified monitor from your organization. ### Arguments * **`group_states`** [*optional* *default*=**None**] If this argument is set, the returned data includes additional information (if available) regarding the specified group states, including the last notification timestamp, last resolution timestamp and details about the last time the monitor was triggered. The argument should include a string list indicating what, if any, group states to include. Choose one or more from all, alert, warn, or no data. Example 'alert,warn' * **`name`** [*optional* *default*==**None**] A string to filter monitors by name * **`tags`** [*optional* *default*==**None**] A comma separated list indicating what tags, if any, should be used to filter the list of monitorsby scope, e.g. host:host0 * **`monitor_tags`** [*optional* *default*==**None**] A comma separated list indicating what service and/or custom tags, if any, should be used to filter the list of monitors. Tags created in the Datadog UI automatically have the service key prepended (e.g. service:my-app) * **`with_downtimes`** [*optional* *default*==**true**] If this argument is set to true, then the returned data includes all current downtimes for each monitor.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllMonitorsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAllMonitorsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupStates** | **optional.String**| When specified, shows additional information about the group states. Choose one or more from &#x60;all&#x60;, &#x60;alert&#x60;, &#x60;warn&#x60;, and &#x60;no data&#x60;. | 
 **name** | **optional.String**|  | 
 **tags** | **optional.String**|  | 
 **monitorTags** | **optional.String**|  | 
 **withDowntimes** | **optional.Bool**|  | 

### Return type

[**[]Monitor**](Monitor.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMonitor

> Monitor GetMonitor(ctx, monitorId, optional)
Get details about the specified monitor.

### Overview Get details about the specified monitor from your organization. ### Arguments * **`monitor_id`** [*required*]: The id of the monitor. * **`group_states`** [*optional* *default*=**None**] If this argument is set, the returned data includes additional information (if available) regarding the specified group states, including the last notification timestamp, last resolution timestamp and details about the last time the monitor was triggered. The argument should include a string list indicating what, if any, group states to include. Choose one or more from all, alert, warn, or no data. Example 'alert,warn'

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorId** | **int64**| The id of the monitor | 
 **optional** | ***GetMonitorOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetMonitorOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupStates** | **optional.String**| When specified, shows additional information about the group states. Choose one or more from &#x60;all&#x60;, &#x60;alert&#x60;, &#x60;warn&#x60;, and &#x60;no data&#x60;. | 

### Return type

[**Monitor**](Monitor.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateMonitor

> Monitor ValidateMonitor(ctx, monitor)


### Overview Validate the monitor provided in the request ### Arguments * **`Monitor`** [*required*] The Monitor Object to validate summary: Validate the provided monitor

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitor** | [**Monitor**](Monitor.md)| Monitor request object | 

### Return type

[**Monitor**](Monitor.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

