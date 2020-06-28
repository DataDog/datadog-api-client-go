# \LogsPipelinesApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLogsPipeline**](LogsPipelinesApi.md#CreateLogsPipeline) | **Post** /api/v1/logs/config/pipelines | Create a pipeline
[**DeleteLogsPipeline**](LogsPipelinesApi.md#DeleteLogsPipeline) | **Delete** /api/v1/logs/config/pipelines/{pipeline_id} | Delete a pipeline
[**GetLogsPipeline**](LogsPipelinesApi.md#GetLogsPipeline) | **Get** /api/v1/logs/config/pipelines/{pipeline_id} | Get a pipeline
[**GetLogsPipelineOrder**](LogsPipelinesApi.md#GetLogsPipelineOrder) | **Get** /api/v1/logs/config/pipeline-order | Get pipeline order
[**ListLogsPipelines**](LogsPipelinesApi.md#ListLogsPipelines) | **Get** /api/v1/logs/config/pipelines | Get all pipelines
[**UpdateLogsPipeline**](LogsPipelinesApi.md#UpdateLogsPipeline) | **Put** /api/v1/logs/config/pipelines/{pipeline_id} | Update a pipeline
[**UpdateLogsPipelineOrder**](LogsPipelinesApi.md#UpdateLogsPipelineOrder) | **Put** /api/v1/logs/config/pipeline-order | Update pipeline order



## CreateLogsPipeline

> LogsPipeline CreateLogsPipeline(ctx, body)

Create a pipeline

Create a pipeline in your organization.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**LogsPipeline**](LogsPipeline.md)| Definition of the new pipeline. | 

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

> DeleteLogsPipeline(ctx, pipelineId)

Delete a pipeline

Delete a given pipeline from your organization. This endpoint takes no JSON arguments.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pipelineId** | **string**| ID of the pipeline to delete. | 

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


## GetLogsPipeline

> LogsPipeline GetLogsPipeline(ctx, pipelineId)

Get a pipeline

Get a specific pipeline from your organization. This endpoint takes no JSON arguments.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pipelineId** | **string**| ID of the pipeline to get. | 

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

> LogsPipelinesOrder GetLogsPipelineOrder(ctx, )

Get pipeline order

Get the current order of your pipelines. This endpoint takes no JSON arguments.

### Required Parameters

This endpoint does not need any parameter.

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


## ListLogsPipelines

> []LogsPipeline ListLogsPipelines(ctx, )

Get all pipelines

Get all pipelines from your organization. This endpoint takes no JSON arguments.

### Required Parameters

This endpoint does not need any parameter.

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


## UpdateLogsPipeline

> LogsPipeline UpdateLogsPipeline(ctx, pipelineId, body)

Update a pipeline

Update a given pipeline configuration to change it’s processors or their order.  **Note**: Using this method updates your pipeline configuration by **replacing** your current configuration with the new one sent to your Datadog organization.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pipelineId** | **string**| ID of the pipeline to delete. | 
**body** | [**LogsPipeline**](LogsPipeline.md)| New definition of the pipeline. | 

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

> LogsPipelinesOrder UpdateLogsPipelineOrder(ctx, body)

Update pipeline order

Update the order of your pipelines. Since logs are processed sequentially, reordering a pipeline may change the structure and content of the data processed by other pipelines and their processors.  **Note**: Using the `PUT` method updates your pipeline order by replacing your current order with the new one sent to your Datadog organization.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**LogsPipelinesOrder**](LogsPipelinesOrder.md)| Object containing the new ordered list of pipeline IDs. | 

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

