# \MetricsApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMetricMetadata**](MetricsApi.md#GetMetricMetadata) | **Get** /api/v1/metrics/{metric_name} | Get metric metadata
[**ListActiveMetrics**](MetricsApi.md#ListActiveMetrics) | **Get** /api/v1/metrics | Get active metrics list
[**ListMetrics**](MetricsApi.md#ListMetrics) | **Get** /api/v1/search | Search metrics
[**QueryMetrics**](MetricsApi.md#QueryMetrics) | **Get** /api/v1/query | Query timeseries points
[**UpdateMetricMetadata**](MetricsApi.md#UpdateMetricMetadata) | **Put** /api/v1/metrics/{metric_name} | Edit metric metadata



## GetMetricMetadata

> MetricMetadata GetMetricMetadata(ctx, metricName)

Get metric metadata

Get metadata about a specific metric.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricName** | **string**| Name of the metric for which to get metadata. | 

### Return type

[**MetricMetadata**](MetricMetadata.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListActiveMetrics

> MetricsListResponse ListActiveMetrics(ctx, from, optional)

Get active metrics list

Get the list of actively reporting metrics from a given time until now.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**from** | **int64**| Seconds since the Unix epoch. | 
 **optional** | ***ListActiveMetricsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListActiveMetricsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **host** | **optional.String**| Hostname for filtering the list of metrics returned. If set, metrics retrieved are those with the corresponding hostname tag. | 

### Return type

[**MetricsListResponse**](MetricsListResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMetrics

> MetricSearchResponse ListMetrics(ctx, q)

Search metrics

Search for metrics from the last 24 hours in Datadog.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**q** | **string**| Query string to search metrics upon. Must be prefixed with &#x60;metrics:&#x60;. | 

### Return type

[**MetricSearchResponse**](MetricSearchResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryMetrics

> MetricsQueryResponse QueryMetrics(ctx, from, to, query)

Query timeseries points

Query timeseries points.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**from** | **int64**| Start of the queried time period, seconds since the Unix epoch. | 
**to** | **int64**| End of the queried time period, seconds since the Unix epoch. | 
**query** | **string**| Query string. | 

### Return type

[**MetricsQueryResponse**](MetricsQueryResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMetricMetadata

> MetricMetadata UpdateMetricMetadata(ctx, metricName, body)

Edit metric metadata

Edit metadata of a specific metric. Find out more about [supported types](https://docs.datadoghq.com/developers/metrics).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricName** | **string**| Name of the metric for which to edit metadata. | 
**body** | [**MetricMetadata**](MetricMetadata.md)| New metadata. | 

### Return type

[**MetricMetadata**](MetricMetadata.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

