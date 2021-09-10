# MetricsApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------ | ------------ | ------------
[**GetMetricMetadata**](MetricsApi.md#GetMetricMetadata) | **Get** /api/v1/metrics/{metric_name} | Get metric metadata
[**ListActiveMetrics**](MetricsApi.md#ListActiveMetrics) | **Get** /api/v1/metrics | Get active metrics list
[**ListMetrics**](MetricsApi.md#ListMetrics) | **Get** /api/v1/search | Search metrics
[**QueryMetrics**](MetricsApi.md#QueryMetrics) | **Get** /api/v1/query | Query timeseries points
[**SubmitMetrics**](MetricsApi.md#SubmitMetrics) | **Post** /api/v1/series | Submit metrics
[**UpdateMetricMetadata**](MetricsApi.md#UpdateMetricMetadata) | **Put** /api/v1/metrics/{metric_name} | Edit metric metadata



## GetMetricMetadata

> MetricMetadata GetMetricMetadata(ctx, metricName)

Get metadata about a specific metric.

### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
    ctx := datadog.NewDefaultContext(context.Background())

    metricName := "metricName_example" // string | Name of the metric for which to get metadata.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.MetricsApi.GetMetricMetadata(ctx, metricName)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.GetMetricMetadata`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMetricMetadata`: MetricMetadata
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from MetricsApi.GetMetricMetadata:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**metricName** | **string** | Name of the metric for which to get metadata. | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**MetricMetadata**](MetricMetadata.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListActiveMetrics

> MetricsListResponse ListActiveMetrics(ctx, from, datadog.ListActiveMetricsOptionalParameters{})

Get the list of actively reporting metrics from a given time until now.

### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
    ctx := datadog.NewDefaultContext(context.Background())

    from := int64(789) // int64 | Seconds since the Unix epoch.
    host := "host_example" // string | Hostname for filtering the list of metrics returned. If set, metrics retrieved are those with the corresponding hostname tag. (optional)
    tagFilter := "env IN (staging,test) AND service:web" // string | Filter metrics that have been submitted with the given tags. Supports boolean and wildcard expressions. Cannot be combined with other filters. (optional)
    optionalParams := datadog.ListActiveMetricsOptionalParameters{
        Host: &host,
        TagFilter: &tagFilter,
    }

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.MetricsApi.ListActiveMetrics(ctx, from, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.ListActiveMetrics`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListActiveMetrics`: MetricsListResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from MetricsApi.ListActiveMetrics:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**from** | **int64** | Seconds since the Unix epoch. | 


### Optional Parameters


Other parameters are passed through a pointer to a ListActiveMetricsOptionalParameters struct.


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**host** | **string** | Hostname for filtering the list of metrics returned. If set, metrics retrieved are those with the corresponding hostname tag. | 
**tagFilter** | **string** | Filter metrics that have been submitted with the given tags. Supports boolean and wildcard expressions. Cannot be combined with other filters. | 

### Return type

[**MetricsListResponse**](MetricsListResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMetrics

> MetricSearchResponse ListMetrics(ctx, q)

Search for metrics from the last 24 hours in Datadog.

### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
    ctx := datadog.NewDefaultContext(context.Background())

    q := "q_example" // string | Query string to search metrics upon. Must be prefixed with `metrics:`.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.MetricsApi.ListMetrics(ctx, q)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.ListMetrics`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMetrics`: MetricSearchResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from MetricsApi.ListMetrics:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**q** | **string** | Query string to search metrics upon. Must be prefixed with &#x60;metrics:&#x60;. | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**MetricSearchResponse**](MetricSearchResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryMetrics

> MetricsQueryResponse QueryMetrics(ctx, from, to, query)

Query timeseries points.

### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
    ctx := datadog.NewDefaultContext(context.Background())

    from := int64(789) // int64 | Start of the queried time period, seconds since the Unix epoch.
    to := int64(789) // int64 | End of the queried time period, seconds since the Unix epoch.
    query := "query_example" // string | Query string.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.MetricsApi.QueryMetrics(ctx, from, to, query)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.QueryMetrics`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryMetrics`: MetricsQueryResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from MetricsApi.QueryMetrics:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**from** | **int64** | Start of the queried time period, seconds since the Unix epoch. |  |
**to** | **int64** | End of the queried time period, seconds since the Unix epoch. |  |
**query** | **string** | Query string. | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**MetricsQueryResponse**](MetricsQueryResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitMetrics

> IntakePayloadAccepted SubmitMetrics(ctx, body)

The metrics end-point allows you to post time-series data that can be graphed on Datadog’s dashboards.
The maximum payload size is 3.2 megabytes (3200000 bytes). Compressed payloads must have a decompressed size of less than 62 megabytes (62914560 bytes).

If you’re submitting metrics directly to the Datadog API without using DogStatsD, expect:

- 64 bits for the timestamp
- 32 bits for the value
- 20 bytes for the metric names
- 50 bytes for the timeseries
- The full payload is approximately 100 bytes. However, with the DogStatsD API,
compression is applied, which reduces the payload size.

### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
    ctx := datadog.NewDefaultContext(context.Background())

    body := *datadog.NewMetricsPayload([]datadog.Series{*datadog.NewSeries("system.load.1", [][]float64{[]float64{float64(123)}})}) // MetricsPayload | 

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.MetricsApi.SubmitMetrics(ctx, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.SubmitMetrics`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitMetrics`: IntakePayloadAccepted
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from MetricsApi.SubmitMetrics:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**body** | [**MetricsPayload**](MetricsPayload.md) |  | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**IntakePayloadAccepted**](IntakePayloadAccepted.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMetricMetadata

> MetricMetadata UpdateMetricMetadata(ctx, metricName, body)

Edit metadata of a specific metric. Find out more about [supported types](https://docs.datadoghq.com/developers/metrics).

### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
    ctx := datadog.NewDefaultContext(context.Background())

    metricName := "metricName_example" // string | Name of the metric for which to edit metadata.
    body := *datadog.NewMetricMetadata() // MetricMetadata | New metadata.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.MetricsApi.UpdateMetricMetadata(ctx, metricName, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.UpdateMetricMetadata`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMetricMetadata`: MetricMetadata
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from MetricsApi.UpdateMetricMetadata:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**metricName** | **string** | Name of the metric for which to edit metadata. |  |
**body** | [**MetricMetadata**](MetricMetadata.md) | New metadata. | 


### Optional Parameters

This endpoint does not have optional parameters.


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

