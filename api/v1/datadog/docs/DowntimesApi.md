# \DowntimesApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelDowntime**](DowntimesApi.md#CancelDowntime) | **Delete** /api/v1/downtime/{downtime_id} | Cancel a downtime
[**CancelDowntimesByScope**](DowntimesApi.md#CancelDowntimesByScope) | **Post** /api/v1/downtime/cancel/by_scope | Cancel downtimes by scope
[**CreateDowntime**](DowntimesApi.md#CreateDowntime) | **Post** /api/v1/downtime | Schedule a downtime
[**GetDowntime**](DowntimesApi.md#GetDowntime) | **Get** /api/v1/downtime/{downtime_id} | Get a downtime
[**ListDowntimes**](DowntimesApi.md#ListDowntimes) | **Get** /api/v1/downtime | Get all downtimes
[**UpdateDowntime**](DowntimesApi.md#UpdateDowntime) | **Put** /api/v1/downtime/{downtime_id} | Update a downtime



## CancelDowntime

> CancelDowntime(ctx, downtimeId)

Cancel a downtime

Cancel a downtime.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**downtimeId** | **int64**| ID of the downtime to cancel. | 

### Return type

 (empty response body)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelDowntimesByScope

> CanceledDowntimesIds CancelDowntimesByScope(ctx, body)

Cancel downtimes by scope

Delete all downtimes that match the scope of `X`.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**CancelDowntimesByScopeRequest**](CancelDowntimesByScopeRequest.md)| Scope to cancel downtimes for. | 

### Return type

[**CanceledDowntimesIds**](CanceledDowntimesIds.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDowntime

> Downtime CreateDowntime(ctx, body)

Schedule a downtime

Schedule a downtime.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**Downtime**](Downtime.md)| Schedule a downtime request body. | 

### Return type

[**Downtime**](Downtime.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDowntime

> Downtime GetDowntime(ctx, downtimeId)

Get a downtime

Get downtime detail by `downtime_id`.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**downtimeId** | **int64**| ID of the downtime to fetch. | 

### Return type

[**Downtime**](Downtime.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDowntimes

> []Downtime ListDowntimes(ctx, optional)

Get all downtimes

Get all scheduled downtimes.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListDowntimesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListDowntimesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **currentOnly** | **optional.Bool**| Only return downtimes that are active when the request is made. | 

### Return type

[**[]Downtime**](Downtime.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDowntime

> Downtime UpdateDowntime(ctx, downtimeId, body)

Update a downtime

Update a single downtime by `downtime_id`.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**downtimeId** | **int64**| ID of the downtime to update. | 
**body** | [**Downtime**](Downtime.md)| Update a downtime request body. | 

### Return type

[**Downtime**](Downtime.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

