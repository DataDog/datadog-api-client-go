# \MetricsApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTagConfiguration**](MetricsApi.md#CreateTagConfiguration) | **Post** /api/v2/metrics/{metric_name}/tags | Create a Tag Configuration
[**DeleteTagConfiguration**](MetricsApi.md#DeleteTagConfiguration) | **Delete** /api/v2/metrics/{metric_name}/tags | Delete a Tag Configuration
[**ListTagConfigurationByName**](MetricsApi.md#ListTagConfigurationByName) | **Get** /api/v2/metrics/{metric_name}/tags | List Tag Configuration by Name
[**ListTagConfigurations**](MetricsApi.md#ListTagConfigurations) | **Get** /api/v2/metrics | List Tag Configurations
[**UpdateTagConfiguration**](MetricsApi.md#UpdateTagConfiguration) | **Patch** /api/v2/metrics/{metric_name}/tags | Update a Tag Configuration



## CreateTagConfiguration

> MetricTagConfigurationResponse CreateTagConfiguration(ctx, metricName).Body(body).Execute()

Create a Tag Configuration



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

    if site, ok := os.LookupEnv("DD_SITE"); ok {
        ctx = context.WithValue(
            ctx,
            datadog.ContextServerVariables,
            map[string]string{"site": site},
        )
    }

    metricName := "dist.http.endpoint.request" // string | The name of the metric.
    body := *datadog.NewMetricTagConfigurationCreateRequest(*datadog.NewMetricTagConfigurationCreateData("test.metric.latency", datadog.MetricTagConfigurationType("manage_tags"))) // MetricTagConfigurationCreateRequest | 

    configuration := datadog.NewConfiguration()
    configuration.SetUnstableOperationEnabled("CreateTagConfiguration", true)

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.MetricsApi.CreateTagConfiguration(ctx, metricName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.CreateTagConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTagConfiguration`: MetricTagConfigurationResponse
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from MetricsApi.CreateTagConfiguration:\n%s\n", response_content)
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

Delete a Tag Configuration



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

    if site, ok := os.LookupEnv("DD_SITE"); ok {
        ctx = context.WithValue(
            ctx,
            datadog.ContextServerVariables,
            map[string]string{"site": site},
        )
    }

    metricName := "dist.http.endpoint.request" // string | The name of the metric.

    configuration := datadog.NewConfiguration()
    configuration.SetUnstableOperationEnabled("DeleteTagConfiguration", true)

    api_client := datadog.NewAPIClient(configuration)
    r, err := api_client.MetricsApi.DeleteTagConfiguration(ctx, metricName).Execute()
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

List Tag Configuration by Name



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

    if site, ok := os.LookupEnv("DD_SITE"); ok {
        ctx = context.WithValue(
            ctx,
            datadog.ContextServerVariables,
            map[string]string{"site": site},
        )
    }

    metricName := "dist.http.endpoint.request" // string | The name of the metric.

    configuration := datadog.NewConfiguration()
    configuration.SetUnstableOperationEnabled("ListTagConfigurationByName", true)

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.MetricsApi.ListTagConfigurationByName(ctx, metricName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.ListTagConfigurationByName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTagConfigurationByName`: MetricTagConfigurationResponse
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from MetricsApi.ListTagConfigurationByName:\n%s\n", response_content)
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

> MetricsAndMetricTagConfigurationsResponse ListTagConfigurations(ctx).FilterConfigured(filterConfigured).FilterTagsConfigured(filterTagsConfigured).FilterMetricType(filterMetricType).FilterIncludePercentiles(filterIncludePercentiles).Execute()

List Tag Configurations



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

    if site, ok := os.LookupEnv("DD_SITE"); ok {
        ctx = context.WithValue(
            ctx,
            datadog.ContextServerVariables,
            map[string]string{"site": site},
        )
    }

    filterConfigured := true // bool | Filter metrics that have configured tags. (optional)
    filterTagsConfigured := "app" // string | Filter tag configurations by configured tags. (optional)
    filterMetricType := datadog.MetricTagConfigurationMetricTypes("gauge") // MetricTagConfigurationMetricTypes | Filter tag configurations by metric type. (optional) (default to "gauge")
    filterIncludePercentiles := true // bool | Filter distributions with additional percentile aggregations enabled or disabled. (optional)

    configuration := datadog.NewConfiguration()
    configuration.SetUnstableOperationEnabled("ListTagConfigurations", true)

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.MetricsApi.ListTagConfigurations(ctx).FilterConfigured(filterConfigured).FilterTagsConfigured(filterTagsConfigured).FilterMetricType(filterMetricType).FilterIncludePercentiles(filterIncludePercentiles).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.ListTagConfigurations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTagConfigurations`: MetricsAndMetricTagConfigurationsResponse
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from MetricsApi.ListTagConfigurations:\n%s\n", response_content)
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


## UpdateTagConfiguration

> MetricTagConfigurationResponse UpdateTagConfiguration(ctx, metricName).Body(body).Execute()

Update a Tag Configuration



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

    if site, ok := os.LookupEnv("DD_SITE"); ok {
        ctx = context.WithValue(
            ctx,
            datadog.ContextServerVariables,
            map[string]string{"site": site},
        )
    }

    metricName := "dist.http.endpoint.request" // string | The name of the metric.
    body := *datadog.NewMetricTagConfigurationUpdateRequest(*datadog.NewMetricTagConfigurationUpdateData("test.metric.latency", datadog.MetricTagConfigurationType("manage_tags"))) // MetricTagConfigurationUpdateRequest | 

    configuration := datadog.NewConfiguration()
    configuration.SetUnstableOperationEnabled("UpdateTagConfiguration", true)

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.MetricsApi.UpdateTagConfiguration(ctx, metricName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.UpdateTagConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTagConfiguration`: MetricTagConfigurationResponse
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from MetricsApi.UpdateTagConfiguration:\n%s\n", response_content)
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

