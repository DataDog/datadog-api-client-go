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

> LogsMetricResponse CreateLogsMetric(ctx).Body(body).Execute()

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
    ctx := context.WithValue(
        context.Background(),
        datadog.ContextAPIKeys,
        map[string]datadog.APIKey{
            "apiKeyAuth": {
                Key: os.Getenv("DD_CLIENT_API_KEY"),
            },
            "appKeyAuth": {
                Key: os.Getenv("DD_CLIENT_APP_KEY"),
            },
        },
    )

    body := *datadog.NewLogsMetricCreateRequest(*datadog.NewLogsMetricCreateData(*datadog.NewLogsMetricCreateAttributes(*datadog.NewLogsMetricCompute(datadog.LogsMetricComputeAggregationType("count"))), "logs.page.load.count", datadog.LogsMetricType("logs_metrics"))) // LogsMetricCreateRequest | The definition of the new log-based metric.

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.LogsMetricsApi.CreateLogsMetric(ctx).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsMetricsApi.CreateLogsMetric``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLogsMetric`: LogsMetricResponse
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from LogsMetricsApi.CreateLogsMetric:\n%s\n", response_content)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLogsMetricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**LogsMetricCreateRequest**](LogsMetricCreateRequest.md) | The definition of the new log-based metric. | 

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

> DeleteLogsMetric(ctx, metricId).Execute()

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
    ctx := context.WithValue(
        context.Background(),
        datadog.ContextAPIKeys,
        map[string]datadog.APIKey{
            "apiKeyAuth": {
                Key: os.Getenv("DD_CLIENT_API_KEY"),
            },
            "appKeyAuth": {
                Key: os.Getenv("DD_CLIENT_APP_KEY"),
            },
        },
    )

    metricId := "metricId_example" // string | The name of the log-based metric.

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    r, err := api_client.LogsMetricsApi.DeleteLogsMetric(ctx, metricId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsMetricsApi.DeleteLogsMetric``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricId** | **string** | The name of the log-based metric. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLogsMetricRequest struct via the builder pattern


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


## GetLogsMetric

> LogsMetricResponse GetLogsMetric(ctx, metricId).Execute()

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
    ctx := context.WithValue(
        context.Background(),
        datadog.ContextAPIKeys,
        map[string]datadog.APIKey{
            "apiKeyAuth": {
                Key: os.Getenv("DD_CLIENT_API_KEY"),
            },
            "appKeyAuth": {
                Key: os.Getenv("DD_CLIENT_APP_KEY"),
            },
        },
    )

    metricId := "metricId_example" // string | The name of the log-based metric.

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.LogsMetricsApi.GetLogsMetric(ctx, metricId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsMetricsApi.GetLogsMetric``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLogsMetric`: LogsMetricResponse
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from LogsMetricsApi.GetLogsMetric:\n%s\n", response_content)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricId** | **string** | The name of the log-based metric. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogsMetricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> LogsMetricsResponse ListLogsMetrics(ctx).Execute()

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
    ctx := context.WithValue(
        context.Background(),
        datadog.ContextAPIKeys,
        map[string]datadog.APIKey{
            "apiKeyAuth": {
                Key: os.Getenv("DD_CLIENT_API_KEY"),
            },
            "appKeyAuth": {
                Key: os.Getenv("DD_CLIENT_APP_KEY"),
            },
        },
    )


    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.LogsMetricsApi.ListLogsMetrics(ctx).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsMetricsApi.ListLogsMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLogsMetrics`: LogsMetricsResponse
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from LogsMetricsApi.ListLogsMetrics:\n%s\n", response_content)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListLogsMetricsRequest struct via the builder pattern


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

> LogsMetricResponse UpdateLogsMetric(ctx, metricId).Body(body).Execute()

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
    ctx := context.WithValue(
        context.Background(),
        datadog.ContextAPIKeys,
        map[string]datadog.APIKey{
            "apiKeyAuth": {
                Key: os.Getenv("DD_CLIENT_API_KEY"),
            },
            "appKeyAuth": {
                Key: os.Getenv("DD_CLIENT_APP_KEY"),
            },
        },
    )

    metricId := "metricId_example" // string | The name of the log-based metric.
    body := *datadog.NewLogsMetricUpdateRequest(*datadog.NewLogsMetricUpdateData(*datadog.NewLogsMetricUpdateAttributes(), datadog.LogsMetricType("logs_metrics"))) // LogsMetricUpdateRequest | New definition of the log-based metric.

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.LogsMetricsApi.UpdateLogsMetric(ctx, metricId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsMetricsApi.UpdateLogsMetric``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLogsMetric`: LogsMetricResponse
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from LogsMetricsApi.UpdateLogsMetric:\n%s\n", response_content)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricId** | **string** | The name of the log-based metric. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLogsMetricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**LogsMetricUpdateRequest**](LogsMetricUpdateRequest.md) | New definition of the log-based metric. | 

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

