# \LogsIndexesApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllLogIndexes**](LogsIndexesApi.md#GetAllLogIndexes) | **Get** /api/v1/logs/config/indexes | This endpoint returns an array of the LogIndex objects of your organization.
[**GetLogsIndex**](LogsIndexesApi.md#GetLogsIndex) | **Get** /api/v1/logs/config/indexes/{name} | This endpoint returns an Index identified by its name.
[**GetLogsIndexOrder**](LogsIndexesApi.md#GetLogsIndexOrder) | **Get** /api/v1/logs/config/index-order | Get the current order of your log indexes.
[**UpdateLogsIndex**](LogsIndexesApi.md#UpdateLogsIndex) | **Put** /api/v1/logs/config/indexes/{name} | This endpoint updates an Index identified by its name.
[**UpdateLogsIndexOrder**](LogsIndexesApi.md#UpdateLogsIndexOrder) | **Put** /api/v1/logs/config/index-order | Update the order of your log indexes.



## GetAllLogIndexes

> LogsIndexListResponse GetAllLogIndexes(ctx).Execute()

This endpoint returns an array of the LogIndex objects of your organization.



### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllLogIndexesRequest struct via the builder pattern


### Return type

[**LogsIndexListResponse**](LogsIndexListResponse.md)

### Authorization

[apiKeyAuthHeader](../README.md#apiKeyAuthHeader), [appKeyAuthHeader](../README.md#appKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogsIndex

> LogsIndex GetLogsIndex(ctx, name).Execute()

This endpoint returns an Index identified by its name.



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the log index | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogsIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LogsIndex**](LogsIndex.md)

### Authorization

[apiKeyAuthHeader](../README.md#apiKeyAuthHeader), [appKeyAuthHeader](../README.md#appKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogsIndexOrder

> LogsIndexesOrder GetLogsIndexOrder(ctx).Execute()

Get the current order of your log indexes.



### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogsIndexOrderRequest struct via the builder pattern


### Return type

[**LogsIndexesOrder**](LogsIndexesOrder.md)

### Authorization

[apiKeyAuthHeader](../README.md#apiKeyAuthHeader), [appKeyAuthHeader](../README.md#appKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLogsIndex

> LogsIndex UpdateLogsIndex(ctx, name).Body(body).Execute()

This endpoint updates an Index identified by its name.



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the log index | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLogsIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**LogsIndex**](LogsIndex.md) | Object containing the new LogsIndex | 

### Return type

[**LogsIndex**](LogsIndex.md)

### Authorization

[apiKeyAuthHeader](../README.md#apiKeyAuthHeader), [appKeyAuthHeader](../README.md#appKeyAuthHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLogsIndexOrder

> LogsIndexesOrder UpdateLogsIndexOrder(ctx).Body(body).Execute()

Update the order of your log indexes.



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLogsIndexOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**LogsIndexesOrder**](LogsIndexesOrder.md) | Object containing the new ordered list of index names | 

### Return type

[**LogsIndexesOrder**](LogsIndexesOrder.md)

### Authorization

[apiKeyAuthHeader](../README.md#apiKeyAuthHeader), [appKeyAuthHeader](../README.md#appKeyAuthHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

