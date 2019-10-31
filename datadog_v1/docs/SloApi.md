# \SloApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkPartialDeleteSlo**](SloApi.md#BulkPartialDeleteSlo) | **Post** /api/v1/slo/bulk_delete | Delete (or partially delete) multiple service level objective objects.
[**CheckCanDeleteSlo**](SloApi.md#CheckCanDeleteSlo) | **Get** /api/v1/slo/can_delete | Check if SLOs can be safely deleted.
[**CreateSlo**](SloApi.md#CreateSlo) | **Post** /api/v1/slo | Create a service level objective object.
[**DeleteSlo**](SloApi.md#DeleteSlo) | **Delete** /api/v1/slo/{slo_id} | Delete the specified service level objective object.
[**EditSlo**](SloApi.md#EditSlo) | **Put** /api/v1/slo/{slo_id} | Edit the specified service level objective
[**GetSlo**](SloApi.md#GetSlo) | **Get** /api/v1/slo/{slo_id} | Get a service level objective object
[**GetSlos**](SloApi.md#GetSlos) | **Get** /api/v1/slo | Get multiple service level objective objects by their IDs.
[**HistoryForSlo**](SloApi.md#HistoryForSlo) | **Get** /api/v1/slo/{slo_id}/history | Get the history of the service level objective.



## BulkPartialDeleteSlo

> ServiceLevelObjectivesBulkDeleted BulkPartialDeleteSlo(ctx, requestBody)
Delete (or partially delete) multiple service level objective objects.

### Overview
Delete (or partially delete) multiple service level objective objects.
This endpoint facilitates deletion of one or more thresholds for one or more service level objective objects. If all thresholds are deleted, the service level objective object is deleted as well.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestBody** | [**map[string][]SloTimeframe**](array.md)| Thresholds by service level objective object ID | 

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


## CheckCanDeleteSlo

> CheckCanDeleteServiceLevelObjectiveResponse CheckCanDeleteSlo(ctx, ids)
Check if SLOs can be safely deleted.

### Overview
Check if an SLO can be safely deleted without disrupting dashboards for example.
### Arguments
* **`ids`** [*required*]: The ID (csv) of the service level objective objects to check.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ids** | **string**| A comma separated list of the IDs of the service level objectives objects (e.g. \&quot;id1,id2,id3\&quot;). | 

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


## CreateSlo

> ServiceLevelObjectiveResponse CreateSlo(ctx, serviceLevelObjective)
Create a service level objective object.

### Overview
Create a service level objective object.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceLevelObjective** | [**ServiceLevelObjective**](ServiceLevelObjective.md)| Service level objective request object | 

### Return type

[**ServiceLevelObjectiveResponse**](ServiceLevelObjectiveResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSlo

> ServiceLevelObjectiveDeleted DeleteSlo(ctx, sloId)
Delete the specified service level objective object.

### Overview
Delete the specified service level objective object.
### Arguments
* **`slo_id`** [*required*]: The ID of the service level objective object

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sloId** | **int64**| The id of the service level objective | 

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


## EditSlo

> ServiceLevelObjectiveResponse EditSlo(ctx, sloId, serviceLevelObjective)
Edit the specified service level objective

### Overview
Edit the specified service level objective object.
### Arguments
* **`slo_id`** [*required*]: The ID of the service level objective object

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sloId** | **string**| The ID of the service level objective object | 
**serviceLevelObjective** | [**ServiceLevelObjective**](ServiceLevelObjective.md)| The edited service level objective request object. | 

### Return type

[**ServiceLevelObjectiveResponse**](ServiceLevelObjectiveResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSlo

> ServiceLevelObjectiveResponse GetSlo(ctx, sloId)
Get a service level objective object

### Overview
Get a service level objective object.
### Arguments
* **`slo_id`** [*required*]: The ID of the service level objective object

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sloId** | **string**| The ID of the service level objective object | 

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


## GetSlos

> ServiceLevelObjectiveResponse GetSlos(ctx, ids)
Get multiple service level objective objects by their IDs.

### Overview
Get multiple service level objective objects by their IDs.
### Arguments
* **`ids`** [*required*]: A comma separated list of the IDs of the service level
  objectives objects (e.g. "id1,id2,id3").

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ids** | **string**| A comma separated list of the IDs of the service level objectives objects (e.g. \&quot;id1,id2,id3\&quot;). | 

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


## HistoryForSlo

> HistoryServiceLevelObjectiveResponse HistoryForSlo(ctx, sloId, fromTs, toTs)
Get the history of the service level objective.

### Overview
Get the SLO history data
### Arguments
* **`slo_id`** [*required*]: The ID of the service level objective object
* **`from_ts`** [*required*]: The `from` timestamp in epoch seconds for the query timeframe
* **`to_ts`** [*required*]: The `to` timestamp in epoch seconds for the query timeframe

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sloId** | **string**| The ID of the service level objective object | 
**fromTs** | **string**| The &#x60;from&#x60; timestamp for the query window in epoch seconds | 
**toTs** | **string**| The &#x60;to&#x60; timestamp for the query window in epoch seconds | 

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

