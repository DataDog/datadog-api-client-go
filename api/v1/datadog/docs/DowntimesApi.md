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

> CancelDowntime(ctx, downtimeId).Execute()

Cancel a downtime



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**downtimeId** | **int64** | ID of the downtime to cancel. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelDowntimeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> CanceledDowntimesIds CancelDowntimesByScope(ctx).Body(body).Execute()

Cancel downtimes by scope



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCancelDowntimesByScopeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CancelDowntimesByScopeRequest**](CancelDowntimesByScopeRequest.md) | Scope to cancel downtimes for. | 

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

> Downtime CreateDowntime(ctx).Body(body).Execute()

Schedule a downtime



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDowntimeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Downtime**](Downtime.md) | Schedule a downtime request body. | 

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

> Downtime GetDowntime(ctx, downtimeId).Execute()

Get a downtime



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**downtimeId** | **int64** | ID of the downtime to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDowntimeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> []Downtime ListDowntimes(ctx).CurrentOnly(currentOnly).Execute()

Get all downtimes



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDowntimesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **currentOnly** | **bool** | Only return downtimes that are active when the request is made. | 

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

> Downtime UpdateDowntime(ctx, downtimeId).Body(body).Execute()

Update a downtime



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**downtimeId** | **int64** | ID of the downtime to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDowntimeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**Downtime**](Downtime.md) | Update a downtime request body. | 

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

