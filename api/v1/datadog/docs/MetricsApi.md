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

> MetricMetadata GetMetricMetadata(ctx, metricName).Execute()

Get metric metadata



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

    if site, ok := os.LookupEnv("DD_SITE"); ok {
        ctx = context.WithValue(
            ctx,
            datadog.ContextServerVariables,
            map[string]string{"site": site},
        )
    }

    metricName := "metricName_example" // string | Name of the metric for which to get metadata.

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.MetricsApi.GetMetricMetadata(ctx, metricName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.GetMetricMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMetricMetadata`: MetricMetadata
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from MetricsApi.GetMetricMetadata:\n%s\n", response_content)
}
```

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


## ListActiveMetrics

> MetricsListResponse ListActiveMetrics(ctx).From(from).Host(host).Execute()

Get active metrics list



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

    if site, ok := os.LookupEnv("DD_SITE"); ok {
        ctx = context.WithValue(
            ctx,
            datadog.ContextServerVariables,
            map[string]string{"site": site},
        )
    }

    from := int64(789) // int64 | Seconds since the Unix epoch.
    host := "host_example" // string | Hostname for filtering the list of metrics returned. If set, metrics retrieved are those with the corresponding hostname tag. (optional)

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.MetricsApi.ListActiveMetrics(ctx).From(from).Host(host).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.ListActiveMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListActiveMetrics`: MetricsListResponse
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from MetricsApi.ListActiveMetrics:\n%s\n", response_content)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListActiveMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **int64** | Seconds since the Unix epoch. | 
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


## ListMetrics

> MetricSearchResponse ListMetrics(ctx).Q(q).Execute()

Search metrics



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

    if site, ok := os.LookupEnv("DD_SITE"); ok {
        ctx = context.WithValue(
            ctx,
            datadog.ContextServerVariables,
            map[string]string{"site": site},
        )
    }

    q := "q_example" // string | Query string to search metrics upon. Must be prefixed with `metrics:`.

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.MetricsApi.ListMetrics(ctx).Q(q).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.ListMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMetrics`: MetricSearchResponse
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from MetricsApi.ListMetrics:\n%s\n", response_content)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMetricsRequest struct via the builder pattern


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


## QueryMetrics

> MetricsQueryResponse QueryMetrics(ctx).From(from).To(to).Query(query).Execute()

Query timeseries points



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

    if site, ok := os.LookupEnv("DD_SITE"); ok {
        ctx = context.WithValue(
            ctx,
            datadog.ContextServerVariables,
            map[string]string{"site": site},
        )
    }

    from := int64(789) // int64 | Start of the queried time period, seconds since the Unix epoch.
    to := int64(789) // int64 | End of the queried time period, seconds since the Unix epoch.
    query := "query_example" // string | Query string.

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.MetricsApi.QueryMetrics(ctx).From(from).To(to).Query(query).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.QueryMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryMetrics`: MetricsQueryResponse
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from MetricsApi.QueryMetrics:\n%s\n", response_content)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **int64** | Start of the queried time period, seconds since the Unix epoch. | 
 **to** | **int64** | End of the queried time period, seconds since the Unix epoch. | 
 **query** | **string** | Query string. | 

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

> MetricMetadata UpdateMetricMetadata(ctx, metricName).Body(body).Execute()

Edit metric metadata



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

    if site, ok := os.LookupEnv("DD_SITE"); ok {
        ctx = context.WithValue(
            ctx,
            datadog.ContextServerVariables,
            map[string]string{"site": site},
        )
    }

    metricName := "metricName_example" // string | Name of the metric for which to edit metadata.
    body := *datadog.NewMetricMetadata() // MetricMetadata | New metadata.

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.MetricsApi.UpdateMetricMetadata(ctx, metricName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.UpdateMetricMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMetricMetadata`: MetricMetadata
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from MetricsApi.UpdateMetricMetadata:\n%s\n", response_content)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricName** | **string** | Name of the metric for which to edit metadata. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMetricMetadataRequest struct via the builder pattern


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

