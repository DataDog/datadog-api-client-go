# \SLOApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkPartialDeleteSLO**](SLOApi.md#BulkPartialDeleteSLO) | **Post** /api/v1/slo/bulk_delete | Delete (or partially delete) multiple service level objective objects.
[**CheckCanDeleteSLO**](SLOApi.md#CheckCanDeleteSLO) | **Get** /api/v1/slo/can_delete | Check if SLOs can be safely deleted.
[**CreateSLO**](SLOApi.md#CreateSLO) | **Post** /api/v1/slo | Create a service level objective object.
[**DeleteSLO**](SLOApi.md#DeleteSLO) | **Delete** /api/v1/slo/{slo_id} | Delete the specified service level objective object.
[**EditSLO**](SLOApi.md#EditSLO) | **Put** /api/v1/slo/{slo_id} | Edit the specified service level objective
[**GetSLO**](SLOApi.md#GetSLO) | **Get** /api/v1/slo/{slo_id} | Get a service level objective object
[**GetSLOs**](SLOApi.md#GetSLOs) | **Get** /api/v1/slo | Get multiple service level objective objects by their IDs.
[**HistoryForSLO**](SLOApi.md#HistoryForSLO) | **Get** /api/v1/slo/{slo_id}/history | Get the history of the service level objective.



## BulkPartialDeleteSLO

> ServiceLevelObjectivesBulkDeleted BulkPartialDeleteSLO(ctx).Body(body).Execute()

Delete (or partially delete) multiple service level objective objects.



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkPartialDeleteSLORequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**map[string][]SLOTimeframe**](array.md) | Thresholds by service level objective object ID | 

### Return type

[**ServiceLevelObjectivesBulkDeleted**](ServiceLevelObjectivesBulkDeleted.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckCanDeleteSLO

> CheckCanDeleteServiceLevelObjectiveResponse CheckCanDeleteSLO(ctx).Ids(ids).Execute()

Check if SLOs can be safely deleted.



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckCanDeleteSLORequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **string** | A comma separated list of the IDs of the service level objectives objects (e.g. \&quot;id1,id2,id3\&quot;). | 

### Return type

[**CheckCanDeleteServiceLevelObjectiveResponse**](CheckCanDeleteServiceLevelObjectiveResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSLO

> ServiceLevelObjectiveListResponse CreateSLO(ctx).Body(body).Execute()

Create a service level objective object.



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSLORequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ServiceLevelObjective**](ServiceLevelObjective.md) | Service level objective request object | 

### Return type

[**ServiceLevelObjectiveListResponse**](ServiceLevelObjectiveListResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSLO

> ServiceLevelObjectiveDeleted DeleteSLO(ctx, sloId).Execute()

Delete the specified service level objective object.



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sloId** | **string** | The id of the service level objective | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSLORequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceLevelObjectiveDeleted**](ServiceLevelObjectiveDeleted.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditSLO

> ServiceLevelObjectiveListResponse EditSLO(ctx, sloId).Body(body).Execute()

Edit the specified service level objective



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sloId** | **string** | The ID of the service level objective object | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditSLORequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ServiceLevelObjective**](ServiceLevelObjective.md) | The edited service level objective request object. | 

### Return type

[**ServiceLevelObjectiveListResponse**](ServiceLevelObjectiveListResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSLO

> ServiceLevelObjectiveResponse GetSLO(ctx, sloId).Execute()

Get a service level objective object



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sloId** | **string** | The ID of the service level objective object | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSLORequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceLevelObjectiveResponse**](ServiceLevelObjectiveResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSLOs

> ServiceLevelObjectiveListResponse GetSLOs(ctx).Ids(ids).Execute()

Get multiple service level objective objects by their IDs.



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSLOsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **string** | A comma separated list of the IDs of the service level objectives objects (e.g. \&quot;id1,id2,id3\&quot;). | 

### Return type

[**ServiceLevelObjectiveListResponse**](ServiceLevelObjectiveListResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HistoryForSLO

> HistoryServiceLevelObjectiveResponse HistoryForSLO(ctx, sloId).FromTs(fromTs).ToTs(toTs).Execute()

Get the history of the service level objective.



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sloId** | **string** | The ID of the service level objective object | 

### Other Parameters

Other parameters are passed through a pointer to a apiHistoryForSLORequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fromTs** | **string** | The &#x60;from&#x60; timestamp for the query window in epoch seconds | 
 **toTs** | **string** | The &#x60;to&#x60; timestamp for the query window in epoch seconds | 

### Return type

[**HistoryServiceLevelObjectiveResponse**](HistoryServiceLevelObjectiveResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

