# \ServiceLevelObjectivesApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckCanDeleteSLO**](ServiceLevelObjectivesApi.md#CheckCanDeleteSLO) | **Get** /api/v1/slo/can_delete | Check if SLOs can be safely deleted
[**CreateSLO**](ServiceLevelObjectivesApi.md#CreateSLO) | **Post** /api/v1/slo | Create a SLO object
[**DeleteSLO**](ServiceLevelObjectivesApi.md#DeleteSLO) | **Delete** /api/v1/slo/{slo_id} | Delete a SLO
[**DeleteSLOTimeframeInBulk**](ServiceLevelObjectivesApi.md#DeleteSLOTimeframeInBulk) | **Post** /api/v1/slo/bulk_delete | Bulk Delete SLO Timeframes
[**GetSLO**](ServiceLevelObjectivesApi.md#GetSLO) | **Get** /api/v1/slo/{slo_id} | Get a SLO&#39;s details
[**GetSLOHistory**](ServiceLevelObjectivesApi.md#GetSLOHistory) | **Get** /api/v1/slo/{slo_id}/history | Get an SLO&#39;s history
[**GetSLOs**](ServiceLevelObjectivesApi.md#GetSLOs) | **Get** /api/v1/slo | Search SLOs
[**UpdateSLO**](ServiceLevelObjectivesApi.md#UpdateSLO) | **Put** /api/v1/slo/{slo_id} | Update a SLO



## CheckCanDeleteSLO

> CheckCanDeleteServiceLevelObjectiveResponse CheckCanDeleteSLO(ctx).Ids(ids).Execute()

Check if SLOs can be safely deleted



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

Create a SLO object



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSLORequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ServiceLevelObjective**](ServiceLevelObjective.md) | Service level objective request object. | 

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

Delete a SLO



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sloId** | **string** | The id of the service level objective. | 

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


## DeleteSLOTimeframeInBulk

> ServiceLevelObjectivesBulkDeleted DeleteSLOTimeframeInBulk(ctx).Body(body).Execute()

Bulk Delete SLO Timeframes



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSLOTimeframeInBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**map[string][]SLOTimeframe**](array.md) | Thresholds by service level objective object ID. | 

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


## GetSLO

> ServiceLevelObjectiveResponse GetSLO(ctx, sloId).Execute()

Get a SLO's details



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sloId** | **string** | The ID of the service level objective object. | 

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


## GetSLOHistory

> HistoryServiceLevelObjectiveResponse GetSLOHistory(ctx, sloId).FromTs(fromTs).ToTs(toTs).Execute()

Get an SLO's history



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sloId** | **string** | The ID of the service level objective object. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSLOHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fromTs** | **string** | The &#x60;from&#x60; timestamp for the query window in epoch seconds. | 
 **toTs** | **string** | The &#x60;to&#x60; timestamp for the query window in epoch seconds. | 

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


## GetSLOs

> ServiceLevelObjectiveListResponse GetSLOs(ctx).Ids(ids).Execute()

Search SLOs



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSLOsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **string** | A comma separated list of the IDs of the service level objectives objects. For example, \&quot;id1,id2,id3\&quot;. | 

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


## UpdateSLO

> ServiceLevelObjectiveListResponse UpdateSLO(ctx, sloId).Body(body).Execute()

Update a SLO



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sloId** | **string** | The ID of the service level objective object. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSLORequest struct via the builder pattern


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

