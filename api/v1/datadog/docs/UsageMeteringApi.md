# \UsageMeteringApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDailyCustomReports**](UsageMeteringApi.md#GetDailyCustomReports) | **Get** /api/v1/daily_custom_reports | Get the list of available daily custom reports
[**GetIncidentManagement**](UsageMeteringApi.md#GetIncidentManagement) | **Get** /api/v1/usage/incident-management | Get hourly usage for incident management
[**GetIngestedSpans**](UsageMeteringApi.md#GetIngestedSpans) | **Get** /api/v1/usage/ingested-spans | Get hourly usage for ingested spans
[**GetMonthlyCustomReports**](UsageMeteringApi.md#GetMonthlyCustomReports) | **Get** /api/v1/monthly_custom_reports | Get the list of available monthly custom reports
[**GetSpecifiedDailyCustomReports**](UsageMeteringApi.md#GetSpecifiedDailyCustomReports) | **Get** /api/v1/daily_custom_reports/{report_id} | Get specified daily custom reports
[**GetSpecifiedMonthlyCustomReports**](UsageMeteringApi.md#GetSpecifiedMonthlyCustomReports) | **Get** /api/v1/monthly_custom_reports/{report_id} | Get specified monthly custom reports
[**GetTracingWithoutLimits**](UsageMeteringApi.md#GetTracingWithoutLimits) | **Get** /api/v1/usage/tracing-without-limits | Get hourly usage for tracing without limits
[**GetUsageAnalyzedLogs**](UsageMeteringApi.md#GetUsageAnalyzedLogs) | **Get** /api/v1/usage/analyzed_logs | Get hourly usage for analyzed logs
[**GetUsageAttribution**](UsageMeteringApi.md#GetUsageAttribution) | **Get** /api/v1/usage/attribution | Get Usage Attribution
[**GetUsageBillableSummary**](UsageMeteringApi.md#GetUsageBillableSummary) | **Get** /api/v1/usage/billable-summary | Get billable usage across your multi-org account
[**GetUsageFargate**](UsageMeteringApi.md#GetUsageFargate) | **Get** /api/v1/usage/fargate | Get hourly usage for Fargate
[**GetUsageHosts**](UsageMeteringApi.md#GetUsageHosts) | **Get** /api/v1/usage/hosts | Get hourly usage for hosts and containers
[**GetUsageIndexedSpans**](UsageMeteringApi.md#GetUsageIndexedSpans) | **Get** /api/v1/usage/indexed-spans | Get hourly usage for indexed spans
[**GetUsageLambda**](UsageMeteringApi.md#GetUsageLambda) | **Get** /api/v1/usage/aws_lambda | Get hourly usage for Lambda
[**GetUsageLogs**](UsageMeteringApi.md#GetUsageLogs) | **Get** /api/v1/usage/logs | Get hourly usage for Logs
[**GetUsageLogsByIndex**](UsageMeteringApi.md#GetUsageLogsByIndex) | **Get** /api/v1/usage/logs_by_index | Get hourly usage for Logs by Index
[**GetUsageNetworkFlows**](UsageMeteringApi.md#GetUsageNetworkFlows) | **Get** /api/v1/usage/network_flows | Get hourly usage for Network Flows
[**GetUsageNetworkHosts**](UsageMeteringApi.md#GetUsageNetworkHosts) | **Get** /api/v1/usage/network_hosts | Get hourly usage for Network Hosts
[**GetUsageProfiling**](UsageMeteringApi.md#GetUsageProfiling) | **Get** /api/v1/usage/profiling | Get hourly usage for profiled hosts
[**GetUsageRumSessions**](UsageMeteringApi.md#GetUsageRumSessions) | **Get** /api/v1/usage/rum_sessions | Get hourly usage for RUM Sessions
[**GetUsageSNMP**](UsageMeteringApi.md#GetUsageSNMP) | **Get** /api/v1/usage/snmp | Get hourly usage for SNMP devices
[**GetUsageSummary**](UsageMeteringApi.md#GetUsageSummary) | **Get** /api/v1/usage/summary | Get usage across your multi-org account
[**GetUsageSynthetics**](UsageMeteringApi.md#GetUsageSynthetics) | **Get** /api/v1/usage/synthetics | Get hourly usage for Synthetics Checks
[**GetUsageSyntheticsAPI**](UsageMeteringApi.md#GetUsageSyntheticsAPI) | **Get** /api/v1/usage/synthetics_api | Get hourly usage for Synthetics API Checks
[**GetUsageSyntheticsBrowser**](UsageMeteringApi.md#GetUsageSyntheticsBrowser) | **Get** /api/v1/usage/synthetics_browser | Get hourly usage for Synthetics Browser Checks
[**GetUsageTimeseries**](UsageMeteringApi.md#GetUsageTimeseries) | **Get** /api/v1/usage/timeseries | Get hourly usage for custom metrics
[**GetUsageTopAvgMetrics**](UsageMeteringApi.md#GetUsageTopAvgMetrics) | **Get** /api/v1/usage/top_avg_metrics | Get top custom metrics by hourly average
[**GetUsageTrace**](UsageMeteringApi.md#GetUsageTrace) | **Get** /api/v1/usage/traces | Get hourly usage for Trace Search



## GetDailyCustomReports

> UsageCustomReportsResponse GetDailyCustomReports(ctx).PageSize(pageSize).PageNumber(pageNumber).SortDir(sortDir).Sort(sort).Execute()

Get the list of available daily custom reports



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

    pageSize := int64(789) // int64 | The number of files to return in the response. `[default=60]`. (optional)
    pageNumber := int64(789) // int64 | The identifier of the first page to return. This parameter is used for the pagination feature `[default=0]`. (optional)
    sortDir := datadog.UsageSortDirection("desc") // UsageSortDirection | The direction to sort by: `[desc, asc]`. (optional) (default to "desc")
    sort := datadog.UsageSort("computed_on") // UsageSort | The field to sort by: `[computed_on, size, start_date, end_date]`. (optional) (default to "start_date")

    configuration := datadog.NewConfiguration()
    configuration.SetUnstableOperationEnabled("GetDailyCustomReports", true)

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.UsageMeteringApi.GetDailyCustomReports(ctx).PageSize(pageSize).PageNumber(pageNumber).SortDir(sortDir).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetDailyCustomReports``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDailyCustomReports`: UsageCustomReportsResponse
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from UsageMeteringApi.GetDailyCustomReports:\n%v\n", response_content)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDailyCustomReportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int64** | The number of files to return in the response. &#x60;[default&#x3D;60]&#x60;. | 
 **pageNumber** | **int64** | The identifier of the first page to return. This parameter is used for the pagination feature &#x60;[default&#x3D;0]&#x60;. | 
 **sortDir** | [**UsageSortDirection**](UsageSortDirection.md) | The direction to sort by: &#x60;[desc, asc]&#x60;. | [default to &quot;desc&quot;]
 **sort** | [**UsageSort**](UsageSort.md) | The field to sort by: &#x60;[computed_on, size, start_date, end_date]&#x60;. | [default to &quot;start_date&quot;]

### Return type

[**UsageCustomReportsResponse**](UsageCustomReportsResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;datetime-format=rfc3339

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIncidentManagement

> UsageIncidentManagementResponse GetIncidentManagement(ctx).StartHr(startHr).EndHr(endHr).Execute()

Get hourly usage for incident management



### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    "time"
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

    startHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: `[YYYY-MM-DDThh]` for usage beginning at this hour.
    endHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: `[YYYY-MM-DDThh]` for usage ending **before** this hour. (optional)

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.UsageMeteringApi.GetIncidentManagement(ctx).StartHr(startHr).EndHr(endHr).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetIncidentManagement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIncidentManagement`: UsageIncidentManagementResponse
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from UsageMeteringApi.GetIncidentManagement:\n%v\n", response_content)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIncidentManagementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: &#x60;[YYYY-MM-DDThh]&#x60; for usage beginning at this hour. | 
 **endHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: &#x60;[YYYY-MM-DDThh]&#x60; for usage ending **before** this hour. | 

### Return type

[**UsageIncidentManagementResponse**](UsageIncidentManagementResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;datetime-format=rfc3339

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIngestedSpans

> UsageIngestedSpansResponse GetIngestedSpans(ctx).StartHr(startHr).EndHr(endHr).Execute()

Get hourly usage for ingested spans



### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    "time"
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

    startHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: `[YYYY-MM-DDThh]` for usage beginning at this hour.
    endHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: `[YYYY-MM-DDThh]` for usage ending **before** this hour. (optional)

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.UsageMeteringApi.GetIngestedSpans(ctx).StartHr(startHr).EndHr(endHr).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetIngestedSpans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIngestedSpans`: UsageIngestedSpansResponse
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from UsageMeteringApi.GetIngestedSpans:\n%v\n", response_content)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIngestedSpansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: &#x60;[YYYY-MM-DDThh]&#x60; for usage beginning at this hour. | 
 **endHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: &#x60;[YYYY-MM-DDThh]&#x60; for usage ending **before** this hour. | 

### Return type

[**UsageIngestedSpansResponse**](UsageIngestedSpansResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;datetime-format=rfc3339

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMonthlyCustomReports

> UsageCustomReportsResponse GetMonthlyCustomReports(ctx).PageSize(pageSize).PageNumber(pageNumber).SortDir(sortDir).Sort(sort).Execute()

Get the list of available monthly custom reports



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

    pageSize := int64(789) // int64 | The number of files to return in the response `[default=60].` (optional)
    pageNumber := int64(789) // int64 | The identifier of the first page to return. This parameter is used for the pagination feature `[default=0]`. (optional)
    sortDir := datadog.UsageSortDirection("desc") // UsageSortDirection | The direction to sort by: `[desc, asc]`. (optional) (default to "desc")
    sort := datadog.UsageSort("computed_on") // UsageSort | The field to sort by: `[computed_on, size, start_date, end_date]`. (optional) (default to "start_date")

    configuration := datadog.NewConfiguration()
    configuration.SetUnstableOperationEnabled("GetMonthlyCustomReports", true)

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.UsageMeteringApi.GetMonthlyCustomReports(ctx).PageSize(pageSize).PageNumber(pageNumber).SortDir(sortDir).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetMonthlyCustomReports``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMonthlyCustomReports`: UsageCustomReportsResponse
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from UsageMeteringApi.GetMonthlyCustomReports:\n%v\n", response_content)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMonthlyCustomReportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int64** | The number of files to return in the response &#x60;[default&#x3D;60].&#x60; | 
 **pageNumber** | **int64** | The identifier of the first page to return. This parameter is used for the pagination feature &#x60;[default&#x3D;0]&#x60;. | 
 **sortDir** | [**UsageSortDirection**](UsageSortDirection.md) | The direction to sort by: &#x60;[desc, asc]&#x60;. | [default to &quot;desc&quot;]
 **sort** | [**UsageSort**](UsageSort.md) | The field to sort by: &#x60;[computed_on, size, start_date, end_date]&#x60;. | [default to &quot;start_date&quot;]

### Return type

[**UsageCustomReportsResponse**](UsageCustomReportsResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;datetime-format=rfc3339

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpecifiedDailyCustomReports

> UsageSpecifiedCustomReportsResponse GetSpecifiedDailyCustomReports(ctx, reportId).Execute()

Get specified daily custom reports



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

    reportId := "reportId_example" // string | The specified ID to search results for.

    configuration := datadog.NewConfiguration()
    configuration.SetUnstableOperationEnabled("GetSpecifiedDailyCustomReports", true)

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.UsageMeteringApi.GetSpecifiedDailyCustomReports(ctx, reportId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetSpecifiedDailyCustomReports``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSpecifiedDailyCustomReports`: UsageSpecifiedCustomReportsResponse
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from UsageMeteringApi.GetSpecifiedDailyCustomReports:\n%v\n", response_content)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **string** | The specified ID to search results for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpecifiedDailyCustomReportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UsageSpecifiedCustomReportsResponse**](UsageSpecifiedCustomReportsResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;datetime-format=rfc3339

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpecifiedMonthlyCustomReports

> UsageSpecifiedCustomReportsResponse GetSpecifiedMonthlyCustomReports(ctx, reportId).Execute()

Get specified monthly custom reports



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

    reportId := "reportId_example" // string | The specified ID to search results for.

    configuration := datadog.NewConfiguration()
    configuration.SetUnstableOperationEnabled("GetSpecifiedMonthlyCustomReports", true)

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.UsageMeteringApi.GetSpecifiedMonthlyCustomReports(ctx, reportId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetSpecifiedMonthlyCustomReports``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSpecifiedMonthlyCustomReports`: UsageSpecifiedCustomReportsResponse
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from UsageMeteringApi.GetSpecifiedMonthlyCustomReports:\n%v\n", response_content)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **string** | The specified ID to search results for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpecifiedMonthlyCustomReportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UsageSpecifiedCustomReportsResponse**](UsageSpecifiedCustomReportsResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;datetime-format=rfc3339

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTracingWithoutLimits

> UsageTracingWithoutLimitsResponse GetTracingWithoutLimits(ctx).StartHr(startHr).EndHr(endHr).Execute()

Get hourly usage for tracing without limits



### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    "time"
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

    startHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: `[YYYY-MM-DDThh]` for usage beginning at this hour.
    endHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: `[YYYY-MM-DDThh]` for usage ending **before** this hour. (optional)

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.UsageMeteringApi.GetTracingWithoutLimits(ctx).StartHr(startHr).EndHr(endHr).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetTracingWithoutLimits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTracingWithoutLimits`: UsageTracingWithoutLimitsResponse
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from UsageMeteringApi.GetTracingWithoutLimits:\n%v\n", response_content)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTracingWithoutLimitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: &#x60;[YYYY-MM-DDThh]&#x60; for usage beginning at this hour. | 
 **endHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: &#x60;[YYYY-MM-DDThh]&#x60; for usage ending **before** this hour. | 

### Return type

[**UsageTracingWithoutLimitsResponse**](UsageTracingWithoutLimitsResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;datetime-format=rfc3339

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsageAnalyzedLogs

> UsageAnalyzedLogsResponse GetUsageAnalyzedLogs(ctx).StartHr(startHr).EndHr(endHr).Execute()

Get hourly usage for analyzed logs



### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    "time"
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

    startHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: `[YYYY-MM-DDThh]` for usage beginning at this hour.
    endHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: `[YYYY-MM-DDThh]` for usage ending **before** this hour. (optional)

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.UsageMeteringApi.GetUsageAnalyzedLogs(ctx).StartHr(startHr).EndHr(endHr).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetUsageAnalyzedLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsageAnalyzedLogs`: UsageAnalyzedLogsResponse
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from UsageMeteringApi.GetUsageAnalyzedLogs:\n%v\n", response_content)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsageAnalyzedLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: &#x60;[YYYY-MM-DDThh]&#x60; for usage beginning at this hour. | 
 **endHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: &#x60;[YYYY-MM-DDThh]&#x60; for usage ending **before** this hour. | 

### Return type

[**UsageAnalyzedLogsResponse**](UsageAnalyzedLogsResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;datetime-format=rfc3339

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsageAttribution

> UsageAttributionResponse GetUsageAttribution(ctx).StartMonth(startMonth).Fields(fields).EndMonth(endMonth).SortDirection(sortDirection).SortName(sortName).Execute()

Get Usage Attribution



### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    "time"
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

    startMonth := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to month: `[YYYY-MM]` for usage beginning in this month. Maximum of 15 months ago.
    fields := "fields_example" // string | The specified field to search results for.
    endMonth := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to month: `[YYYY-MM]` for usage ending this month. (optional)
    sortDirection := datadog.UsageSortDirection("desc") // UsageSortDirection | The direction to sort by: `[desc, asc]`. (optional) (default to "desc")
    sortName := datadog.UsageAttributionSort("api_percentage") // UsageAttributionSort | The field to sort by. (optional) (default to "custom_timeseries_usage")

    configuration := datadog.NewConfiguration()
    configuration.SetUnstableOperationEnabled("GetUsageAttribution", true)

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.UsageMeteringApi.GetUsageAttribution(ctx).StartMonth(startMonth).Fields(fields).EndMonth(endMonth).SortDirection(sortDirection).SortName(sortName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetUsageAttribution``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsageAttribution`: UsageAttributionResponse
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from UsageMeteringApi.GetUsageAttribution:\n%v\n", response_content)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsageAttributionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startMonth** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to month: &#x60;[YYYY-MM]&#x60; for usage beginning in this month. Maximum of 15 months ago. | 
 **fields** | **string** | The specified field to search results for. | 
 **endMonth** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to month: &#x60;[YYYY-MM]&#x60; for usage ending this month. | 
 **sortDirection** | [**UsageSortDirection**](UsageSortDirection.md) | The direction to sort by: &#x60;[desc, asc]&#x60;. | [default to &quot;desc&quot;]
 **sortName** | [**UsageAttributionSort**](UsageAttributionSort.md) | The field to sort by. | [default to &quot;custom_timeseries_usage&quot;]

### Return type

[**UsageAttributionResponse**](UsageAttributionResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;datetime-format=rfc3339

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsageBillableSummary

> UsageBillableSummaryResponse GetUsageBillableSummary(ctx).Month(month).Execute()

Get billable usage across your multi-org account



### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    "time"
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

    month := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to month: `[YYYY-MM]` for usage starting this month. (optional)

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.UsageMeteringApi.GetUsageBillableSummary(ctx).Month(month).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetUsageBillableSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsageBillableSummary`: UsageBillableSummaryResponse
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from UsageMeteringApi.GetUsageBillableSummary:\n%v\n", response_content)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsageBillableSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **month** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to month: &#x60;[YYYY-MM]&#x60; for usage starting this month. | 

### Return type

[**UsageBillableSummaryResponse**](UsageBillableSummaryResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;datetime-format=rfc3339

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsageFargate

> UsageFargateResponse GetUsageFargate(ctx).StartHr(startHr).EndHr(endHr).Execute()

Get hourly usage for Fargate



### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    "time"
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

    startHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage beginning at this hour.
    endHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage ending **before** this hour. (optional)

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.UsageMeteringApi.GetUsageFargate(ctx).StartHr(startHr).EndHr(endHr).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetUsageFargate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsageFargate`: UsageFargateResponse
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from UsageMeteringApi.GetUsageFargate:\n%v\n", response_content)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsageFargateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage beginning at this hour. | 
 **endHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage ending **before** this hour. | 

### Return type

[**UsageFargateResponse**](UsageFargateResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;datetime-format=rfc3339

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsageHosts

> UsageHostsResponse GetUsageHosts(ctx).StartHr(startHr).EndHr(endHr).Execute()

Get hourly usage for hosts and containers



### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    "time"
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

    startHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage beginning at this hour.
    endHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage ending **before** this hour. (optional)

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.UsageMeteringApi.GetUsageHosts(ctx).StartHr(startHr).EndHr(endHr).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetUsageHosts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsageHosts`: UsageHostsResponse
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from UsageMeteringApi.GetUsageHosts:\n%v\n", response_content)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsageHostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage beginning at this hour. | 
 **endHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage ending **before** this hour. | 

### Return type

[**UsageHostsResponse**](UsageHostsResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;datetime-format=rfc3339

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsageIndexedSpans

> UsageIndexedSpansResponse GetUsageIndexedSpans(ctx).StartHr(startHr).EndHr(endHr).Execute()

Get hourly usage for indexed spans



### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    "time"
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

    startHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: `[YYYY-MM-DDThh]` for usage beginning at this hour.
    endHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: `[YYYY-MM-DDThh]` for usage ending **before** this hour. (optional)

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.UsageMeteringApi.GetUsageIndexedSpans(ctx).StartHr(startHr).EndHr(endHr).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetUsageIndexedSpans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsageIndexedSpans`: UsageIndexedSpansResponse
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from UsageMeteringApi.GetUsageIndexedSpans:\n%v\n", response_content)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsageIndexedSpansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: &#x60;[YYYY-MM-DDThh]&#x60; for usage beginning at this hour. | 
 **endHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: &#x60;[YYYY-MM-DDThh]&#x60; for usage ending **before** this hour. | 

### Return type

[**UsageIndexedSpansResponse**](UsageIndexedSpansResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;datetime-format=rfc3339

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsageLambda

> UsageLambdaResponse GetUsageLambda(ctx).StartHr(startHr).EndHr(endHr).Execute()

Get hourly usage for Lambda



### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    "time"
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

    startHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage beginning at this hour.
    endHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage ending **before** this hour. (optional)

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.UsageMeteringApi.GetUsageLambda(ctx).StartHr(startHr).EndHr(endHr).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetUsageLambda``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsageLambda`: UsageLambdaResponse
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from UsageMeteringApi.GetUsageLambda:\n%v\n", response_content)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsageLambdaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage beginning at this hour. | 
 **endHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage ending **before** this hour. | 

### Return type

[**UsageLambdaResponse**](UsageLambdaResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;datetime-format=rfc3339

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsageLogs

> UsageLogsResponse GetUsageLogs(ctx).StartHr(startHr).EndHr(endHr).Execute()

Get hourly usage for Logs



### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    "time"
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

    startHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage beginning at this hour.
    endHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage ending **before** this hour. (optional)

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.UsageMeteringApi.GetUsageLogs(ctx).StartHr(startHr).EndHr(endHr).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetUsageLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsageLogs`: UsageLogsResponse
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from UsageMeteringApi.GetUsageLogs:\n%v\n", response_content)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsageLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage beginning at this hour. | 
 **endHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage ending **before** this hour. | 

### Return type

[**UsageLogsResponse**](UsageLogsResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;datetime-format=rfc3339

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsageLogsByIndex

> UsageLogsByIndexResponse GetUsageLogsByIndex(ctx).StartHr(startHr).EndHr(endHr).IndexName(indexName).Execute()

Get hourly usage for Logs by Index



### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    "time"
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

    startHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage beginning at this hour.
    endHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage ending **before** this hour. (optional)
    indexName := []string{"Inner_example"} // []string | Comma-separated list of log index names. (optional)

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.UsageMeteringApi.GetUsageLogsByIndex(ctx).StartHr(startHr).EndHr(endHr).IndexName(indexName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetUsageLogsByIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsageLogsByIndex`: UsageLogsByIndexResponse
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from UsageMeteringApi.GetUsageLogsByIndex:\n%v\n", response_content)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsageLogsByIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage beginning at this hour. | 
 **endHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage ending **before** this hour. | 
 **indexName** | **[]string** | Comma-separated list of log index names. | 

### Return type

[**UsageLogsByIndexResponse**](UsageLogsByIndexResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;datetime-format=rfc3339

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsageNetworkFlows

> UsageNetworkFlowsResponse GetUsageNetworkFlows(ctx).StartHr(startHr).EndHr(endHr).Execute()

Get hourly usage for Network Flows



### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    "time"
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

    startHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: `[YYYY-MM-DDThh]` for usage beginning at this hour.
    endHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: `[YYYY-MM-DDThh]` for usage ending **before** this hour. (optional)

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.UsageMeteringApi.GetUsageNetworkFlows(ctx).StartHr(startHr).EndHr(endHr).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetUsageNetworkFlows``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsageNetworkFlows`: UsageNetworkFlowsResponse
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from UsageMeteringApi.GetUsageNetworkFlows:\n%v\n", response_content)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsageNetworkFlowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: &#x60;[YYYY-MM-DDThh]&#x60; for usage beginning at this hour. | 
 **endHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: &#x60;[YYYY-MM-DDThh]&#x60; for usage ending **before** this hour. | 

### Return type

[**UsageNetworkFlowsResponse**](UsageNetworkFlowsResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;datetime-format=rfc3339

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsageNetworkHosts

> UsageNetworkHostsResponse GetUsageNetworkHosts(ctx).StartHr(startHr).EndHr(endHr).Execute()

Get hourly usage for Network Hosts



### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    "time"
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

    startHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage beginning at this hour.
    endHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage ending **before** this hour. (optional)

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.UsageMeteringApi.GetUsageNetworkHosts(ctx).StartHr(startHr).EndHr(endHr).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetUsageNetworkHosts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsageNetworkHosts`: UsageNetworkHostsResponse
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from UsageMeteringApi.GetUsageNetworkHosts:\n%v\n", response_content)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsageNetworkHostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage beginning at this hour. | 
 **endHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage ending **before** this hour. | 

### Return type

[**UsageNetworkHostsResponse**](UsageNetworkHostsResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;datetime-format=rfc3339

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsageProfiling

> UsageProfilingResponse GetUsageProfiling(ctx).StartHr(startHr).EndHr(endHr).Execute()

Get hourly usage for profiled hosts



### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    "time"
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

    startHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: `[YYYY-MM-DDThh]` for usage beginning at this hour.
    endHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: `[YYYY-MM-DDThh]` for usage ending **before** this hour. (optional)

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.UsageMeteringApi.GetUsageProfiling(ctx).StartHr(startHr).EndHr(endHr).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetUsageProfiling``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsageProfiling`: UsageProfilingResponse
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from UsageMeteringApi.GetUsageProfiling:\n%v\n", response_content)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsageProfilingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: &#x60;[YYYY-MM-DDThh]&#x60; for usage beginning at this hour. | 
 **endHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: &#x60;[YYYY-MM-DDThh]&#x60; for usage ending **before** this hour. | 

### Return type

[**UsageProfilingResponse**](UsageProfilingResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;datetime-format=rfc3339

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsageRumSessions

> UsageRumSessionsResponse GetUsageRumSessions(ctx).StartHr(startHr).EndHr(endHr).Type_(type_).Execute()

Get hourly usage for RUM Sessions



### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    "time"
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

    startHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage beginning at this hour.
    endHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage ending **before** this hour. (optional)
    type_ := "type__example" // string | RUM type: `[browser, mobile]`. Defaults to `browser`. (optional)

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.UsageMeteringApi.GetUsageRumSessions(ctx).StartHr(startHr).EndHr(endHr).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetUsageRumSessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsageRumSessions`: UsageRumSessionsResponse
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from UsageMeteringApi.GetUsageRumSessions:\n%v\n", response_content)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsageRumSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage beginning at this hour. | 
 **endHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage ending **before** this hour. | 
 **type_** | **string** | RUM type: &#x60;[browser, mobile]&#x60;. Defaults to &#x60;browser&#x60;. | 

### Return type

[**UsageRumSessionsResponse**](UsageRumSessionsResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;datetime-format=rfc3339

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsageSNMP

> UsageSNMPResponse GetUsageSNMP(ctx).StartHr(startHr).EndHr(endHr).Execute()

Get hourly usage for SNMP devices



### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    "time"
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

    startHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: `[YYYY-MM-DDThh]` for usage beginning at this hour.
    endHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: `[YYYY-MM-DDThh]` for usage ending **before** this hour. (optional)

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.UsageMeteringApi.GetUsageSNMP(ctx).StartHr(startHr).EndHr(endHr).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetUsageSNMP``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsageSNMP`: UsageSNMPResponse
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from UsageMeteringApi.GetUsageSNMP:\n%v\n", response_content)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsageSNMPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: &#x60;[YYYY-MM-DDThh]&#x60; for usage beginning at this hour. | 
 **endHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: &#x60;[YYYY-MM-DDThh]&#x60; for usage ending **before** this hour. | 

### Return type

[**UsageSNMPResponse**](UsageSNMPResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;datetime-format=rfc3339

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsageSummary

> UsageSummaryResponse GetUsageSummary(ctx).StartMonth(startMonth).EndMonth(endMonth).IncludeOrgDetails(includeOrgDetails).Execute()

Get usage across your multi-org account



### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    "time"
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

    startMonth := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to month: `[YYYY-MM]` for usage beginning in this month. Maximum of 15 months ago.
    endMonth := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to month: `[YYYY-MM]` for usage ending this month. (optional)
    includeOrgDetails := true // bool | Include usage summaries for each sub-org. (optional)

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.UsageMeteringApi.GetUsageSummary(ctx).StartMonth(startMonth).EndMonth(endMonth).IncludeOrgDetails(includeOrgDetails).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetUsageSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsageSummary`: UsageSummaryResponse
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from UsageMeteringApi.GetUsageSummary:\n%v\n", response_content)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsageSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startMonth** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to month: &#x60;[YYYY-MM]&#x60; for usage beginning in this month. Maximum of 15 months ago. | 
 **endMonth** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to month: &#x60;[YYYY-MM]&#x60; for usage ending this month. | 
 **includeOrgDetails** | **bool** | Include usage summaries for each sub-org. | 

### Return type

[**UsageSummaryResponse**](UsageSummaryResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;datetime-format=rfc3339

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsageSynthetics

> UsageSyntheticsResponse GetUsageSynthetics(ctx).StartHr(startHr).EndHr(endHr).Execute()

Get hourly usage for Synthetics Checks



### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    "time"
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

    startHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage beginning at this hour.
    endHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage ending **before** this hour. (optional)

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.UsageMeteringApi.GetUsageSynthetics(ctx).StartHr(startHr).EndHr(endHr).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetUsageSynthetics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsageSynthetics`: UsageSyntheticsResponse
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from UsageMeteringApi.GetUsageSynthetics:\n%v\n", response_content)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsageSyntheticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage beginning at this hour. | 
 **endHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage ending **before** this hour. | 

### Return type

[**UsageSyntheticsResponse**](UsageSyntheticsResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;datetime-format=rfc3339

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsageSyntheticsAPI

> UsageSyntheticsAPIResponse GetUsageSyntheticsAPI(ctx).StartHr(startHr).EndHr(endHr).Execute()

Get hourly usage for Synthetics API Checks



### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    "time"
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

    startHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage beginning at this hour.
    endHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage ending **before** this hour. (optional)

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.UsageMeteringApi.GetUsageSyntheticsAPI(ctx).StartHr(startHr).EndHr(endHr).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetUsageSyntheticsAPI``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsageSyntheticsAPI`: UsageSyntheticsAPIResponse
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from UsageMeteringApi.GetUsageSyntheticsAPI:\n%v\n", response_content)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsageSyntheticsAPIRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage beginning at this hour. | 
 **endHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage ending **before** this hour. | 

### Return type

[**UsageSyntheticsAPIResponse**](UsageSyntheticsAPIResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;datetime-format=rfc3339

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsageSyntheticsBrowser

> UsageSyntheticsBrowserResponse GetUsageSyntheticsBrowser(ctx).StartHr(startHr).EndHr(endHr).Execute()

Get hourly usage for Synthetics Browser Checks



### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    "time"
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

    startHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage beginning at this hour.
    endHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage ending **before** this hour. (optional)

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.UsageMeteringApi.GetUsageSyntheticsBrowser(ctx).StartHr(startHr).EndHr(endHr).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetUsageSyntheticsBrowser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsageSyntheticsBrowser`: UsageSyntheticsBrowserResponse
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from UsageMeteringApi.GetUsageSyntheticsBrowser:\n%v\n", response_content)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsageSyntheticsBrowserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage beginning at this hour. | 
 **endHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage ending **before** this hour. | 

### Return type

[**UsageSyntheticsBrowserResponse**](UsageSyntheticsBrowserResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;datetime-format=rfc3339

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsageTimeseries

> UsageTimeseriesResponse GetUsageTimeseries(ctx).StartHr(startHr).EndHr(endHr).Execute()

Get hourly usage for custom metrics



### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    "time"
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

    startHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage beginning at this hour.
    endHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage ending **before** this hour. (optional)

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.UsageMeteringApi.GetUsageTimeseries(ctx).StartHr(startHr).EndHr(endHr).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetUsageTimeseries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsageTimeseries`: UsageTimeseriesResponse
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from UsageMeteringApi.GetUsageTimeseries:\n%v\n", response_content)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsageTimeseriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage beginning at this hour. | 
 **endHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage ending **before** this hour. | 

### Return type

[**UsageTimeseriesResponse**](UsageTimeseriesResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;datetime-format=rfc3339

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsageTopAvgMetrics

> UsageTopAvgMetricsResponse GetUsageTopAvgMetrics(ctx).Month(month).Names(names).Limit(limit).Execute()

Get top custom metrics by hourly average



### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    "time"
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

    month := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to month: [YYYY-MM] for usage beginning at this hour.
    names := []string{"Inner_example"} // []string | Comma-separated list of metric names. (optional)
    limit := int32(56) // int32 | Maximum number of results to return (between 1 and 5000) - defaults to 500 results if limit not specified. (optional) (default to 500)

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.UsageMeteringApi.GetUsageTopAvgMetrics(ctx).Month(month).Names(names).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetUsageTopAvgMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsageTopAvgMetrics`: UsageTopAvgMetricsResponse
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from UsageMeteringApi.GetUsageTopAvgMetrics:\n%v\n", response_content)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsageTopAvgMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **month** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to month: [YYYY-MM] for usage beginning at this hour. | 
 **names** | **[]string** | Comma-separated list of metric names. | 
 **limit** | **int32** | Maximum number of results to return (between 1 and 5000) - defaults to 500 results if limit not specified. | [default to 500]

### Return type

[**UsageTopAvgMetricsResponse**](UsageTopAvgMetricsResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;datetime-format=rfc3339

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsageTrace

> UsageTraceResponse GetUsageTrace(ctx).StartHr(startHr).EndHr(endHr).Execute()

Get hourly usage for Trace Search



### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    "time"
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

    startHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage beginning at this hour.
    endHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage ending **before** this hour. (optional)

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.UsageMeteringApi.GetUsageTrace(ctx).StartHr(startHr).EndHr(endHr).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetUsageTrace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsageTrace`: UsageTraceResponse
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from UsageMeteringApi.GetUsageTrace:\n%v\n", response_content)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsageTraceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage beginning at this hour. | 
 **endHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage ending **before** this hour. | 

### Return type

[**UsageTraceResponse**](UsageTraceResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;datetime-format=rfc3339

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

