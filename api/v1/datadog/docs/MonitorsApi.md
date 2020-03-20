# \MonitorsApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMonitor**](MonitorsApi.md#CreateMonitor) | **Post** /api/v1/monitor | Create a monitor
[**DeleteMonitor**](MonitorsApi.md#DeleteMonitor) | **Delete** /api/v1/monitor/{monitor_id} | Delete a monitor
[**EditMonitor**](MonitorsApi.md#EditMonitor) | **Put** /api/v1/monitor/{monitor_id} | Edit a monitor
[**GetAllMonitors**](MonitorsApi.md#GetAllMonitors) | **Get** /api/v1/monitor | Get all monitor details
[**GetMonitor**](MonitorsApi.md#GetMonitor) | **Get** /api/v1/monitor/{monitor_id} | Get a monitor&#39;s details
[**ValidateMonitor**](MonitorsApi.md#ValidateMonitor) | **Post** /api/v1/monitor/validate | Validate a monitor



## CreateMonitor

> Monitor CreateMonitor(ctx).Body(body).Execute()

Create a monitor



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Monitor**](Monitor.md) | Monitor request object | 

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

> DeletedMonitor DeleteMonitor(ctx, monitorId).Execute()

Delete a monitor



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorId** | **int64** | The id of the monitor | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeletedMonitor**](DeletedMonitor.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditMonitor

> Monitor EditMonitor(ctx, monitorId).Body(body).Execute()

Edit a monitor



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorId** | **int64** | The id of the monitor | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**Monitor**](Monitor.md) | Monitor request object | 

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

> []Monitor GetAllMonitors(ctx).GroupStates(groupStates).Name(name).Tags(tags).MonitorTags(monitorTags).WithDowntimes(withDowntimes).Execute()

Get all monitor details



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllMonitorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupStates** | **string** | When specified, shows additional information about the group states. Choose one or more from &#x60;all&#x60;, &#x60;alert&#x60;, &#x60;warn&#x60;, and &#x60;no data&#x60;. | 
 **name** | **string** | A string to filter monitors by name. | 
 **tags** | **string** | A comma separated list indicating what tags, if any, should be used to filter the list of monitorsby scope, e.g. host:host0. | 
 **monitorTags** | **string** | A comma separated list indicating what service and/or custom tags, if any, should be used to filter the list of monitors. Tags created in the Datadog UI automatically have the service key prepended (e.g. service:my-app). | 
 **withDowntimes** | **bool** | If this argument is set to true, then the returned data includes all current downtimes for each monitor. | 

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

> Monitor GetMonitor(ctx, monitorId).GroupStates(groupStates).Execute()

Get a monitor's details



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorId** | **int64** | The id of the monitor | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupStates** | **string** | When specified, shows additional information about the group states. Choose one or more from &#x60;all&#x60;, &#x60;alert&#x60;, &#x60;warn&#x60;, and &#x60;no data&#x60;. | 

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

> Monitor ValidateMonitor(ctx).Body(body).Execute()

Validate a monitor



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Monitor**](Monitor.md) | Monitor request object | 

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

