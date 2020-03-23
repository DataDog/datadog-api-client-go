# \MetricsApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EditMetricMetadata**](MetricsApi.md#EditMetricMetadata) | **Put** /api/v1/metrics/{metric_name} | Edit metric metadata
[**GetAllActiveMetrics**](MetricsApi.md#GetAllActiveMetrics) | **Get** /api/v1/metrics | Get active metrics list
[**GetMetricMetadata**](MetricsApi.md#GetMetricMetadata) | **Get** /api/v1/metrics/{metric_name} | Get metric metadata
[**QueryMetrics**](MetricsApi.md#QueryMetrics) | **Get** /api/v1/query | Query timeseries points
[**SearchMetrics**](MetricsApi.md#SearchMetrics) | **Get** /api/v1/search | Search metrics
[**SubmitMetrics**](MetricsApi.md#SubmitMetrics) | **Post** /api/v1/series | Submit metrics



## EditMetricMetadata

> MetricMetadata EditMetricMetadata(ctx, metricName).Body(body).Execute()

Edit metric metadata



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricName** | **string** | Name of the metric for which to edit metadata. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditMetricMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**MetricMetadata**](MetricMetadata.md) | New metadata. | 

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


## GetAllActiveMetrics

> MetricsListResponse GetAllActiveMetrics(ctx).From(from).Host(host).Execute()

Get active metrics list



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllActiveMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **int64** | Seconds since the Unix epoch | 
 **host** | **string** | Hostname for filtering the list of metrics returned. If set, metrics retrieved are those with the corresponding hostname tag. | 

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


## GetMetricMetadata

> MetricMetadata GetMetricMetadata(ctx, metricName).Execute()

Get metric metadata



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricName** | **string** | Name of the metric for which to get metadata. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetricMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## QueryMetrics

> MetricsQueryResponse QueryMetrics(ctx).From(from).To(to).Query(query).Execute()

Query timeseries points



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **int64** | Start of the queried time period, seconds since the Unix epoch. | 
 **to** | **int64** | End of the queried time period, seconds since the Unix epoch. | 
 **query** | **string** | Query string | 

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


## SearchMetrics

> MetricSearchResponse SearchMetrics(ctx).Q(q).Execute()

Search metrics



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | Query string to search metrics upon. Must be prefixed with &#x60;metrics:&#x60;. | 

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


## SubmitMetrics

> IntakePayloadAccepted SubmitMetrics(ctx).Body(body).Execute()

Submit metrics



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**MetricsPayload**](MetricsPayload.md) |  | 

### Return type

[**IntakePayloadAccepted**](IntakePayloadAccepted.md)

### Authorization

[apiKeyAuthQuery](../README.md#apiKeyAuthQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

