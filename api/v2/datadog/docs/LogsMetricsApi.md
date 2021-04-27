# \LogsMetricsApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLogsMetric**](LogsMetricsApi.md#CreateLogsMetric) | **Post** /api/v2/logs/config/metrics | Create a log-based metric
[**DeleteLogsMetric**](LogsMetricsApi.md#DeleteLogsMetric) | **Delete** /api/v2/logs/config/metrics/{metric_id} | Delete a log-based metric
[**GetLogsMetric**](LogsMetricsApi.md#GetLogsMetric) | **Get** /api/v2/logs/config/metrics/{metric_id} | Get a log-based metric
[**ListLogsMetrics**](LogsMetricsApi.md#ListLogsMetrics) | **Get** /api/v2/logs/config/metrics | Get all log-based metrics
[**UpdateLogsMetric**](LogsMetricsApi.md#UpdateLogsMetric) | **Patch** /api/v2/logs/config/metrics/{metric_id} | Update a log-based metric



## CreateLogsMetric

> LogsMetricResponse CreateLogsMetric(ctx, body)

Create a log-based metric



### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
    ctx := datadog.NewDefaultContext(context.Background())

    body := *datadog.NewLogsMetricCreateRequest(*datadog.NewLogsMetricCreateData(*datadog.NewLogsMetricCreateAttributes(*datadog.NewLogsMetricCompute(datadog.LogsMetricComputeAggregationType("count"))), "logs.page.load.count", datadog.LogsMetricType("logs_metrics"))) // LogsMetricCreateRequest | The definition of the new log-based metric.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.LogsMetricsApi.CreateLogsMetric(ctx, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsMetricsApi.CreateLogsMetric``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLogsMetric`: LogsMetricResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from LogsMetricsApi.CreateLogsMetric:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**LogsMetricCreateRequest**](LogsMetricCreateRequest.md) | The definition of the new log-based metric. | 

### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**LogsMetricResponse**](LogsMetricResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLogsMetric

> DeleteLogsMetric(ctx, metricId)

Delete a log-based metric



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
    ctx := datadog.NewDefaultContext(context.Background())

    metricId := "metricId_example" // string | The name of the log-based metric.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    r, err := apiClient.LogsMetricsApi.DeleteLogsMetric(ctx, metricId)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsMetricsApi.DeleteLogsMetric``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricId** | **string** | The name of the log-based metric. | 

### Optional Parameters

This endpoint does not have optional parameters.


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


## GetLogsMetric

> LogsMetricResponse GetLogsMetric(ctx, metricId)

Get a log-based metric



### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
    ctx := datadog.NewDefaultContext(context.Background())

    metricId := "metricId_example" // string | The name of the log-based metric.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.LogsMetricsApi.GetLogsMetric(ctx, metricId)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsMetricsApi.GetLogsMetric``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLogsMetric`: LogsMetricResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from LogsMetricsApi.GetLogsMetric:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricId** | **string** | The name of the log-based metric. | 

### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**LogsMetricResponse**](LogsMetricResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLogsMetrics

> LogsMetricsResponse ListLogsMetrics(ctx)

Get all log-based metrics



### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
    ctx := datadog.NewDefaultContext(context.Background())


    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.LogsMetricsApi.ListLogsMetrics(ctx)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsMetricsApi.ListLogsMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLogsMetrics`: LogsMetricsResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from LogsMetricsApi.ListLogsMetrics:\n%s\n", responseContent)
}
```

### Required Parameters

This endpoint does not need any parameter.

### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**LogsMetricsResponse**](LogsMetricsResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLogsMetric

> LogsMetricResponse UpdateLogsMetric(ctx, metricId, body)

Update a log-based metric



### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
    ctx := datadog.NewDefaultContext(context.Background())

    metricId := "metricId_example" // string | The name of the log-based metric.
    body := *datadog.NewLogsMetricUpdateRequest(*datadog.NewLogsMetricUpdateData(*datadog.NewLogsMetricUpdateAttributes(), datadog.LogsMetricType("logs_metrics"))) // LogsMetricUpdateRequest | New definition of the log-based metric.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.LogsMetricsApi.UpdateLogsMetric(ctx, metricId, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsMetricsApi.UpdateLogsMetric``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLogsMetric`: LogsMetricResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from LogsMetricsApi.UpdateLogsMetric:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricId** | **string** | The name of the log-based metric. | 
**body** | [**LogsMetricUpdateRequest**](LogsMetricUpdateRequest.md) | New definition of the log-based metric. | 

### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**LogsMetricResponse**](LogsMetricResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

