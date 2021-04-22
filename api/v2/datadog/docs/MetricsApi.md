# \MetricsApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTagConfiguration**](MetricsApi.md#CreateTagConfiguration) | **Post** /api/v2/metrics/{metric_name}/tags | Create a tag configuration
[**DeleteTagConfiguration**](MetricsApi.md#DeleteTagConfiguration) | **Delete** /api/v2/metrics/{metric_name}/tags | Delete a tag configuration
[**ListTagConfigurationByName**](MetricsApi.md#ListTagConfigurationByName) | **Get** /api/v2/metrics/{metric_name}/tags | List tag configuration by name
[**ListTagConfigurations**](MetricsApi.md#ListTagConfigurations) | **Get** /api/v2/metrics | List tag configurations
[**ListTagsByMetricName**](MetricsApi.md#ListTagsByMetricName) | **Get** /api/v2/metrics/{metric_name}/all-tags | List tags by metric name
[**ListVolumesByMetricName**](MetricsApi.md#ListVolumesByMetricName) | **Get** /api/v2/metrics/{metric_name}/volumes | List distinct metric volumes by metric name
[**UpdateTagConfiguration**](MetricsApi.md#UpdateTagConfiguration) | **Patch** /api/v2/metrics/{metric_name}/tags | Update a tag configuration



## CreateTagConfiguration

> MetricTagConfigurationResponse CreateTagConfiguration(ctx, metricName).Body(body).Execute()

Create a tag configuration



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
    resp, r, err := apiClient.MetricsApi.CreateTagConfiguration(ctx, metricName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.CreateTagConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTagConfiguration`: MetricTagConfigurationResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from MetricsApi.CreateTagConfiguration:\n%s\n", responseContent)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricName** | **string** | The name of the metric. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTagConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**MetricTagConfigurationCreateRequest**](MetricTagConfigurationCreateRequest.md) |  | 

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

> DeleteTagConfiguration(ctx, metricName).Execute()

Delete a tag configuration



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
    r, err := apiClient.MetricsApi.DeleteTagConfiguration(ctx, metricName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.DeleteTagConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricName** | **string** | The name of the metric. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTagConfigurationRequest struct via the builder pattern


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


## ListTagConfigurationByName

> MetricTagConfigurationResponse ListTagConfigurationByName(ctx, metricName).Execute()

List tag configuration by name



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
    resp, r, err := apiClient.MetricsApi.ListTagConfigurationByName(ctx, metricName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.ListTagConfigurationByName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTagConfigurationByName`: MetricTagConfigurationResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from MetricsApi.ListTagConfigurationByName:\n%s\n", responseContent)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricName** | **string** | The name of the metric. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTagConfigurationByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> MetricsAndMetricTagConfigurationsResponse ListTagConfigurations(ctx).FilterConfigured(filterConfigured).FilterTagsConfigured(filterTagsConfigured).FilterMetricType(filterMetricType).FilterIncludePercentiles(filterIncludePercentiles).FilterTags(filterTags).WindowSeconds(windowSeconds).Execute()

List tag configurations



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

    configuration := datadog.NewConfiguration()
    configuration.SetUnstableOperationEnabled("ListTagConfigurations", true)

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.MetricsApi.ListTagConfigurations(ctx).FilterConfigured(filterConfigured).FilterTagsConfigured(filterTagsConfigured).FilterMetricType(filterMetricType).FilterIncludePercentiles(filterIncludePercentiles).FilterTags(filterTags).WindowSeconds(windowSeconds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.ListTagConfigurations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTagConfigurations`: MetricsAndMetricTagConfigurationsResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from MetricsApi.ListTagConfigurations:\n%s\n", responseContent)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTagConfigurationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterConfigured** | **bool** | Filter metrics that have configured tags. | 
 **filterTagsConfigured** | **string** | Filter tag configurations by configured tags. | 
 **filterMetricType** | [**MetricTagConfigurationMetricTypes**](MetricTagConfigurationMetricTypes.md) | Filter tag configurations by metric type. | [default to &quot;gauge&quot;]
 **filterIncludePercentiles** | **bool** | Filter distributions with additional percentile aggregations enabled or disabled. | 
 **filterTags** | **string** | Filter metrics that have been submitted with the given tags. Supports boolean and wildcard expressions. Cannot be combined with other filters. | 
 **windowSeconds** | **int64** | The number of seconds of look back (from now) to apply to a filter[tag] query. Defaults value is 3600 (1 hour), maximum value is 172,800 (2 days). | 

### Return type

[**MetricsAndMetricTagConfigurationsResponse**](MetricsAndMetricTagConfigurationsResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTagsByMetricName

> MetricAllTagsResponse ListTagsByMetricName(ctx, metricName).Execute()

List tags by metric name



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
    resp, r, err := apiClient.MetricsApi.ListTagsByMetricName(ctx, metricName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.ListTagsByMetricName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTagsByMetricName`: MetricAllTagsResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from MetricsApi.ListTagsByMetricName:\n%s\n", responseContent)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricName** | **string** | The name of the metric. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTagsByMetricNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MetricAllTagsResponse**](MetricAllTagsResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVolumesByMetricName

> MetricVolumesResponse ListVolumesByMetricName(ctx, metricName).Execute()

List distinct metric volumes by metric name



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
    resp, r, err := apiClient.MetricsApi.ListVolumesByMetricName(ctx, metricName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.ListVolumesByMetricName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVolumesByMetricName`: MetricVolumesResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from MetricsApi.ListVolumesByMetricName:\n%s\n", responseContent)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricName** | **string** | The name of the metric. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListVolumesByMetricNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MetricVolumesResponse**](MetricVolumesResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTagConfiguration

> MetricTagConfigurationResponse UpdateTagConfiguration(ctx, metricName).Body(body).Execute()

Update a tag configuration



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
    resp, r, err := apiClient.MetricsApi.UpdateTagConfiguration(ctx, metricName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.UpdateTagConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTagConfiguration`: MetricTagConfigurationResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from MetricsApi.UpdateTagConfiguration:\n%s\n", responseContent)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricName** | **string** | The name of the metric. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTagConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**MetricTagConfigurationUpdateRequest**](MetricTagConfigurationUpdateRequest.md) |  | 

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

