# MetricsApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------ | ------------ | ------------
[**CreateTagConfiguration**](MetricsApi.md#CreateTagConfiguration) | **Post** /api/v2/metrics/{metric_name}/tags | Create a tag configuration
[**DeleteTagConfiguration**](MetricsApi.md#DeleteTagConfiguration) | **Delete** /api/v2/metrics/{metric_name}/tags | Delete a tag configuration
[**ListTagConfigurationByName**](MetricsApi.md#ListTagConfigurationByName) | **Get** /api/v2/metrics/{metric_name}/tags | List tag configuration by name
[**ListTagConfigurations**](MetricsApi.md#ListTagConfigurations) | **Get** /api/v2/metrics | List tag configurations
[**ListTagsByMetricName**](MetricsApi.md#ListTagsByMetricName) | **Get** /api/v2/metrics/{metric_name}/all-tags | List tags by metric name
[**ListVolumesByMetricName**](MetricsApi.md#ListVolumesByMetricName) | **Get** /api/v2/metrics/{metric_name}/volumes | List distinct metric volumes by metric name
[**UpdateTagConfiguration**](MetricsApi.md#UpdateTagConfiguration) | **Patch** /api/v2/metrics/{metric_name}/tags | Update a tag configuration



## CreateTagConfiguration

> MetricTagConfigurationResponse CreateTagConfiguration(ctx, metricName, body)

Create and define a list of queryable tag keys for an existing count/gauge/rate/distribution metric. Optionally, include percentile aggregations on any distribution metric.
Can only be used with application keys of users with the `Manage Tags for Metrics` permission.

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

    metricName := "dist.http.endpoint.request" // string | The name of the metric.
    body := *datadog.NewMetricTagConfigurationCreateRequest(*datadog.NewMetricTagConfigurationCreateData("test.metric.latency", datadog.MetricTagConfigurationType("manage_tags"))) // MetricTagConfigurationCreateRequest | 

    configuration := datadog.NewConfiguration()
    configuration.SetUnstableOperationEnabled("CreateTagConfiguration", true)

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.MetricsApi.CreateTagConfiguration(ctx, metricName, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.CreateTagConfiguration`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTagConfiguration`: MetricTagConfigurationResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from MetricsApi.CreateTagConfiguration:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**metricName** | **string** | The name of the metric. |  |
**body** | [**MetricTagConfigurationCreateRequest**](MetricTagConfigurationCreateRequest.md) |  | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**MetricTagConfigurationResponse**](MetricTagConfigurationResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTagConfiguration

> DeleteTagConfiguration(ctx, metricName)

Deletes a metric's tag configuration. Can only be used with application
keys from users with the `Manage Tags for Metrics` permission.

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

    metricName := "dist.http.endpoint.request" // string | The name of the metric.

    configuration := datadog.NewConfiguration()
    configuration.SetUnstableOperationEnabled("DeleteTagConfiguration", true)

    apiClient := datadog.NewAPIClient(configuration)
    r, err := apiClient.MetricsApi.DeleteTagConfiguration(ctx, metricName)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.DeleteTagConfiguration`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**metricName** | **string** | The name of the metric. | 


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


## ListTagConfigurationByName

> MetricTagConfigurationResponse ListTagConfigurationByName(ctx, metricName)

Returns the tag configuration for the given metric name.

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

    metricName := "dist.http.endpoint.request" // string | The name of the metric.

    configuration := datadog.NewConfiguration()
    configuration.SetUnstableOperationEnabled("ListTagConfigurationByName", true)

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.MetricsApi.ListTagConfigurationByName(ctx, metricName)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.ListTagConfigurationByName`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTagConfigurationByName`: MetricTagConfigurationResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from MetricsApi.ListTagConfigurationByName:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**metricName** | **string** | The name of the metric. | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**MetricTagConfigurationResponse**](MetricTagConfigurationResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTagConfigurations

> MetricsAndMetricTagConfigurationsResponse ListTagConfigurations(ctx, datadog.ListTagConfigurationsOptionalParameters{})

Returns all configured count/gauge/rate/distribution metric names
(with additional filters if specified).

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

    filterConfigured := true // bool | Filter metrics that have configured tags. (optional)
    filterTagsConfigured := "app" // string | Filter tag configurations by configured tags. (optional)
    filterMetricType := datadog.MetricTagConfigurationMetricTypes("gauge") // MetricTagConfigurationMetricTypes | Filter tag configurations by metric type. (optional) (default to "gauge")
    filterIncludePercentiles := true // bool | Filter distributions with additional percentile aggregations enabled or disabled. (optional)
    filterTags := "env IN (staging,test) AND service:web" // string | Filter metrics that have been submitted with the given tags. Supports boolean and wildcard expressions. Cannot be combined with other filters. (optional)
    windowSeconds := int64(3600) // int64 | The number of seconds of look back (from now) to apply to a filter[tag] query. Defaults value is 3600 (1 hour), maximum value is 172,800 (2 days). (optional)
    optionalParams := datadog.ListTagConfigurationsOptionalParameters{
        FilterConfigured: &filterConfigured,
        FilterTagsConfigured: &filterTagsConfigured,
        FilterMetricType: &filterMetricType,
        FilterIncludePercentiles: &filterIncludePercentiles,
        FilterTags: &filterTags,
        WindowSeconds: &windowSeconds,
    }

    configuration := datadog.NewConfiguration()
    configuration.SetUnstableOperationEnabled("ListTagConfigurations", true)

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.MetricsApi.ListTagConfigurations(ctx, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.ListTagConfigurations`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTagConfigurations`: MetricsAndMetricTagConfigurationsResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from MetricsApi.ListTagConfigurations:\n%s\n", responseContent)
}
```

### Required Parameters




### Optional Parameters


Other parameters are passed through a pointer to a ListTagConfigurationsOptionalParameters struct.


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**filterConfigured** | **bool** | Filter metrics that have configured tags. | 
**filterTagsConfigured** | **string** | Filter tag configurations by configured tags. | 
**filterMetricType** | [**MetricTagConfigurationMetricTypes**](MetricTagConfigurationMetricTypes.md) | Filter tag configurations by metric type. | [default to &quot;gauge&quot;]
**filterIncludePercentiles** | **bool** | Filter distributions with additional percentile aggregations enabled or disabled. | 
**filterTags** | **string** | Filter metrics that have been submitted with the given tags. Supports boolean and wildcard expressions. Cannot be combined with other filters. | 
**windowSeconds** | **int64** | The number of seconds of look back (from now) to apply to a filter[tag] query. Defaults value is 3600 (1 hour), maximum value is 172,800 (2 days). | 

### Return type

[**MetricsAndMetricTagConfigurationsResponse**](MetricsAndMetricTagConfigurationsResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTagsByMetricName

> MetricAllTagsResponse ListTagsByMetricName(ctx, metricName)

View indexed tag key-value pairs for a given metric name.

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

    metricName := "dist.http.endpoint.request" // string | The name of the metric.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.MetricsApi.ListTagsByMetricName(ctx, metricName)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.ListTagsByMetricName`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTagsByMetricName`: MetricAllTagsResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from MetricsApi.ListTagsByMetricName:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**metricName** | **string** | The name of the metric. | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**MetricAllTagsResponse**](MetricAllTagsResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVolumesByMetricName

> MetricVolumesResponse ListVolumesByMetricName(ctx, metricName)

View distinct metrics volumes for the given metric name.

Custom distribution metrics will return both ingested and indexed custom metric volumes.
For Metrics without Limits&trade; beta customers, all metrics will return both ingested/indexed volumes.
Custom metrics generated in-app from other products will return `null` for ingested volumes.

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

    metricName := "dist.http.endpoint.request" // string | The name of the metric.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.MetricsApi.ListVolumesByMetricName(ctx, metricName)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.ListVolumesByMetricName`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVolumesByMetricName`: MetricVolumesResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from MetricsApi.ListVolumesByMetricName:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**metricName** | **string** | The name of the metric. | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**MetricVolumesResponse**](MetricVolumesResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTagConfiguration

> MetricTagConfigurationResponse UpdateTagConfiguration(ctx, metricName, body)

Update the tag configuration of a metric or percentile aggregations of a distribution metric. Can only be used with
application keys from users with the `Manage Tags for Metrics` permission.

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

    metricName := "dist.http.endpoint.request" // string | The name of the metric.
    body := *datadog.NewMetricTagConfigurationUpdateRequest(*datadog.NewMetricTagConfigurationUpdateData("test.metric.latency", datadog.MetricTagConfigurationType("manage_tags"))) // MetricTagConfigurationUpdateRequest | 

    configuration := datadog.NewConfiguration()
    configuration.SetUnstableOperationEnabled("UpdateTagConfiguration", true)

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.MetricsApi.UpdateTagConfiguration(ctx, metricName, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.UpdateTagConfiguration`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTagConfiguration`: MetricTagConfigurationResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from MetricsApi.UpdateTagConfiguration:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**metricName** | **string** | The name of the metric. |  |
**body** | [**MetricTagConfigurationUpdateRequest**](MetricTagConfigurationUpdateRequest.md) |  | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**MetricTagConfigurationResponse**](MetricTagConfigurationResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

