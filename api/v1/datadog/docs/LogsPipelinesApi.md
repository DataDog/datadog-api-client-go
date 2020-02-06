# \LogsPipelinesApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLogsPipeline**](LogsPipelinesApi.md#CreateLogsPipeline) | **Post** /api/v1/logs/config/pipelines | Create a pipeline in your organization.
[**DeleteLogsPipeline**](LogsPipelinesApi.md#DeleteLogsPipeline) | **Delete** /api/v1/logs/config/pipelines/{pipeline_id} | Delete a given pipeline from your organization.
[**GetAllLogsPipelines**](LogsPipelinesApi.md#GetAllLogsPipelines) | **Get** /api/v1/logs/config/pipelines | Get all pipelines from your organization.
[**GetLogsPipeline**](LogsPipelinesApi.md#GetLogsPipeline) | **Get** /api/v1/logs/config/pipelines/{pipeline_id} | Get a specific pipeline from your organization.
[**GetLogsPipelineOrder**](LogsPipelinesApi.md#GetLogsPipelineOrder) | **Get** /api/v1/logs/config/pipeline-order | Get the current order of your pipelines.
[**UpdateLogsPipeline**](LogsPipelinesApi.md#UpdateLogsPipeline) | **Put** /api/v1/logs/config/pipelines/{pipeline_id} | Update a pipeline in your organization.
[**UpdateLogsPipelineOrder**](LogsPipelinesApi.md#UpdateLogsPipelineOrder) | **Put** /api/v1/logs/config/pipeline-order | Update the order of your pipelines.



## CreateLogsPipeline

> LogsPipeline CreateLogsPipeline(ctx).Body(body).Execute()

Create a pipeline in your organization.



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLogsPipelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**LogsPipeline**](LogsPipeline.md) | Definition of the new pipeline | 

### Return type

[**LogsPipeline**](LogsPipeline.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLogsPipeline

> DeleteLogsPipeline(ctx, pipelineId).Execute()

Delete a given pipeline from your organization.



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pipelineId** | **string** | ID of the pipeline to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLogsPipelineRequest struct via the builder pattern


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


## GetAllLogsPipelines

> []LogsPipeline GetAllLogsPipelines(ctx).Execute()

Get all pipelines from your organization.



### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllLogsPipelinesRequest struct via the builder pattern


### Return type

[**[]LogsPipeline**](LogsPipeline.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogsPipeline

> LogsPipeline GetLogsPipeline(ctx, pipelineId).Execute()

Get a specific pipeline from your organization.



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pipelineId** | **string** | ID of the pipeline to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogsPipelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LogsPipeline**](LogsPipeline.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogsPipelineOrder

> LogsPipelinesOrder GetLogsPipelineOrder(ctx).Execute()

Get the current order of your pipelines.



### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogsPipelineOrderRequest struct via the builder pattern


### Return type

[**LogsPipelinesOrder**](LogsPipelinesOrder.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLogsPipeline

> LogsPipeline UpdateLogsPipeline(ctx, pipelineId).Body(body).Execute()

Update a pipeline in your organization.



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pipelineId** | **string** | ID of the pipeline to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLogsPipelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**LogsPipeline**](LogsPipeline.md) | New definition of the pipeline | 

### Return type

[**LogsPipeline**](LogsPipeline.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLogsPipelineOrder

> LogsPipelinesOrder UpdateLogsPipelineOrder(ctx).Body(body).Execute()

Update the order of your pipelines.



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLogsPipelineOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**LogsPipelinesOrder**](LogsPipelinesOrder.md) | Object containing the new ordered list of pipeline IDs | 

### Return type

[**LogsPipelinesOrder**](LogsPipelinesOrder.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

