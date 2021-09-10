# UsageMeteringApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------ | ------------ | ------------
[**GetDailyCustomReports**](UsageMeteringApi.md#GetDailyCustomReports) | **Get** /api/v1/daily_custom_reports | Get the list of available daily custom reports
[**GetIncidentManagement**](UsageMeteringApi.md#GetIncidentManagement) | **Get** /api/v1/usage/incident-management | Get hourly usage for incident management
[**GetIngestedSpans**](UsageMeteringApi.md#GetIngestedSpans) | **Get** /api/v1/usage/ingested-spans | Get hourly usage for ingested spans
[**GetMonthlyCustomReports**](UsageMeteringApi.md#GetMonthlyCustomReports) | **Get** /api/v1/monthly_custom_reports | Get the list of available monthly custom reports
[**GetSpecifiedDailyCustomReports**](UsageMeteringApi.md#GetSpecifiedDailyCustomReports) | **Get** /api/v1/daily_custom_reports/{report_id} | Get specified daily custom reports
[**GetSpecifiedMonthlyCustomReports**](UsageMeteringApi.md#GetSpecifiedMonthlyCustomReports) | **Get** /api/v1/monthly_custom_reports/{report_id} | Get specified monthly custom reports
[**GetUsageAnalyzedLogs**](UsageMeteringApi.md#GetUsageAnalyzedLogs) | **Get** /api/v1/usage/analyzed_logs | Get hourly usage for analyzed logs
[**GetUsageAttribution**](UsageMeteringApi.md#GetUsageAttribution) | **Get** /api/v1/usage/attribution | Get Usage Attribution
[**GetUsageAuditLogs**](UsageMeteringApi.md#GetUsageAuditLogs) | **Get** /api/v1/usage/audit_logs | Get hourly usage for audit logs
[**GetUsageBillableSummary**](UsageMeteringApi.md#GetUsageBillableSummary) | **Get** /api/v1/usage/billable-summary | Get billable usage across your account
[**GetUsageCWS**](UsageMeteringApi.md#GetUsageCWS) | **Get** /api/v1/usage/cws | Get hourly usage for Cloud Workload Security
[**GetUsageCloudSecurityPostureManagement**](UsageMeteringApi.md#GetUsageCloudSecurityPostureManagement) | **Get** /api/v1/usage/cspm | Get hourly usage for CSPM
[**GetUsageDBM**](UsageMeteringApi.md#GetUsageDBM) | **Get** /api/v1/usage/dbm | Get hourly usage for Database Monitoring
[**GetUsageFargate**](UsageMeteringApi.md#GetUsageFargate) | **Get** /api/v1/usage/fargate | Get hourly usage for Fargate
[**GetUsageHosts**](UsageMeteringApi.md#GetUsageHosts) | **Get** /api/v1/usage/hosts | Get hourly usage for hosts and containers
[**GetUsageIndexedSpans**](UsageMeteringApi.md#GetUsageIndexedSpans) | **Get** /api/v1/usage/indexed-spans | Get hourly usage for indexed spans
[**GetUsageInternetOfThings**](UsageMeteringApi.md#GetUsageInternetOfThings) | **Get** /api/v1/usage/iot | Get hourly usage for IoT
[**GetUsageLambda**](UsageMeteringApi.md#GetUsageLambda) | **Get** /api/v1/usage/aws_lambda | Get hourly usage for Lambda
[**GetUsageLogs**](UsageMeteringApi.md#GetUsageLogs) | **Get** /api/v1/usage/logs | Get hourly usage for Logs
[**GetUsageLogsByIndex**](UsageMeteringApi.md#GetUsageLogsByIndex) | **Get** /api/v1/usage/logs_by_index | Get hourly usage for Logs by Index
[**GetUsageLogsByRetention**](UsageMeteringApi.md#GetUsageLogsByRetention) | **Get** /api/v1/usage/logs-by-retention | Get hourly logs usage by retention
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
[**GetUsageTopAvgMetrics**](UsageMeteringApi.md#GetUsageTopAvgMetrics) | **Get** /api/v1/usage/top_avg_metrics | Get all custom metrics by hourly average



## GetDailyCustomReports

> UsageCustomReportsResponse GetDailyCustomReports(ctx, datadog.GetDailyCustomReportsOptionalParameters{})

Get daily custom reports.

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

    pageSize := int64(789) // int64 | The number of files to return in the response. `[default=60]`. (optional)
    pageNumber := int64(789) // int64 | The identifier of the first page to return. This parameter is used for the pagination feature `[default=0]`. (optional)
    sortDir := datadog.UsageSortDirection("desc") // UsageSortDirection | The direction to sort by: `[desc, asc]`. (optional) (default to "desc")
    sort := datadog.UsageSort("computed_on") // UsageSort | The field to sort by: `[computed_on, size, start_date, end_date]`. (optional) (default to "start_date")
    optionalParams := datadog.GetDailyCustomReportsOptionalParameters{
        PageSize: &pageSize,
        PageNumber: &pageNumber,
        SortDir: &sortDir,
        Sort: &sort,
    }

    configuration := datadog.NewConfiguration()
    configuration.SetUnstableOperationEnabled("GetDailyCustomReports", true)

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.UsageMeteringApi.GetDailyCustomReports(ctx, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetDailyCustomReports`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDailyCustomReports`: UsageCustomReportsResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from UsageMeteringApi.GetDailyCustomReports:\n%s\n", responseContent)
}
```

### Required Parameters




### Optional Parameters


Other parameters are passed through a pointer to a GetDailyCustomReportsOptionalParameters struct.


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
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

> UsageIncidentManagementResponse GetIncidentManagement(ctx, startHr, datadog.GetIncidentManagementOptionalParameters{})

Get hourly usage for incident management.

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
    ctx := datadog.NewDefaultContext(context.Background())

    startHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: `[YYYY-MM-DDThh]` for usage beginning at this hour.
    endHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: `[YYYY-MM-DDThh]` for usage ending **before** this hour. (optional)
    optionalParams := datadog.GetIncidentManagementOptionalParameters{
        EndHr: &endHr,
    }

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.UsageMeteringApi.GetIncidentManagement(ctx, startHr, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetIncidentManagement`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIncidentManagement`: UsageIncidentManagementResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from UsageMeteringApi.GetIncidentManagement:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**startHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: &#x60;[YYYY-MM-DDThh]&#x60; for usage beginning at this hour. | 


### Optional Parameters


Other parameters are passed through a pointer to a GetIncidentManagementOptionalParameters struct.


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**endHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: &#x60;[YYYY-MM-DDThh]&#x60; for usage ending **before** this hour. | 

### Return type

[**UsageIncidentManagementResponse**](UsageIncidentManagementResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;datetime-format=rfc3339

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIngestedSpans

> UsageIngestedSpansResponse GetIngestedSpans(ctx, startHr, datadog.GetIngestedSpansOptionalParameters{})

Get hourly usage for ingested spans.

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
    ctx := datadog.NewDefaultContext(context.Background())

    startHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: `[YYYY-MM-DDThh]` for usage beginning at this hour.
    endHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: `[YYYY-MM-DDThh]` for usage ending **before** this hour. (optional)
    optionalParams := datadog.GetIngestedSpansOptionalParameters{
        EndHr: &endHr,
    }

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.UsageMeteringApi.GetIngestedSpans(ctx, startHr, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetIngestedSpans`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIngestedSpans`: UsageIngestedSpansResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from UsageMeteringApi.GetIngestedSpans:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**startHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: &#x60;[YYYY-MM-DDThh]&#x60; for usage beginning at this hour. | 


### Optional Parameters


Other parameters are passed through a pointer to a GetIngestedSpansOptionalParameters struct.


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**endHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: &#x60;[YYYY-MM-DDThh]&#x60; for usage ending **before** this hour. | 

### Return type

[**UsageIngestedSpansResponse**](UsageIngestedSpansResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;datetime-format=rfc3339

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMonthlyCustomReports

> UsageCustomReportsResponse GetMonthlyCustomReports(ctx, datadog.GetMonthlyCustomReportsOptionalParameters{})

Get monthly custom reports.

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

    pageSize := int64(789) // int64 | The number of files to return in the response `[default=60].` (optional)
    pageNumber := int64(789) // int64 | The identifier of the first page to return. This parameter is used for the pagination feature `[default=0]`. (optional)
    sortDir := datadog.UsageSortDirection("desc") // UsageSortDirection | The direction to sort by: `[desc, asc]`. (optional) (default to "desc")
    sort := datadog.UsageSort("computed_on") // UsageSort | The field to sort by: `[computed_on, size, start_date, end_date]`. (optional) (default to "start_date")
    optionalParams := datadog.GetMonthlyCustomReportsOptionalParameters{
        PageSize: &pageSize,
        PageNumber: &pageNumber,
        SortDir: &sortDir,
        Sort: &sort,
    }

    configuration := datadog.NewConfiguration()
    configuration.SetUnstableOperationEnabled("GetMonthlyCustomReports", true)

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.UsageMeteringApi.GetMonthlyCustomReports(ctx, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetMonthlyCustomReports`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMonthlyCustomReports`: UsageCustomReportsResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from UsageMeteringApi.GetMonthlyCustomReports:\n%s\n", responseContent)
}
```

### Required Parameters




### Optional Parameters


Other parameters are passed through a pointer to a GetMonthlyCustomReportsOptionalParameters struct.


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
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

> UsageSpecifiedCustomReportsResponse GetSpecifiedDailyCustomReports(ctx, reportId)

Get specified daily custom reports.

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

    reportId := "reportId_example" // string | Date of the report in the format `YYYY-MM-DD`.

    configuration := datadog.NewConfiguration()
    configuration.SetUnstableOperationEnabled("GetSpecifiedDailyCustomReports", true)

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.UsageMeteringApi.GetSpecifiedDailyCustomReports(ctx, reportId)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetSpecifiedDailyCustomReports`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSpecifiedDailyCustomReports`: UsageSpecifiedCustomReportsResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from UsageMeteringApi.GetSpecifiedDailyCustomReports:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**reportId** | **string** | Date of the report in the format &#x60;YYYY-MM-DD&#x60;. | 


### Optional Parameters

This endpoint does not have optional parameters.


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

> UsageSpecifiedCustomReportsResponse GetSpecifiedMonthlyCustomReports(ctx, reportId)

Get specified monthly custom reports.

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

    reportId := "reportId_example" // string | Date of the report in the format `YYYY-MM-DD`.

    configuration := datadog.NewConfiguration()
    configuration.SetUnstableOperationEnabled("GetSpecifiedMonthlyCustomReports", true)

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.UsageMeteringApi.GetSpecifiedMonthlyCustomReports(ctx, reportId)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetSpecifiedMonthlyCustomReports`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSpecifiedMonthlyCustomReports`: UsageSpecifiedCustomReportsResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from UsageMeteringApi.GetSpecifiedMonthlyCustomReports:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**reportId** | **string** | Date of the report in the format &#x60;YYYY-MM-DD&#x60;. | 


### Optional Parameters

This endpoint does not have optional parameters.


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


## GetUsageAnalyzedLogs

> UsageAnalyzedLogsResponse GetUsageAnalyzedLogs(ctx, startHr, datadog.GetUsageAnalyzedLogsOptionalParameters{})

Get hourly usage for analyzed logs (Security Monitoring).

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
    ctx := datadog.NewDefaultContext(context.Background())

    startHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: `[YYYY-MM-DDThh]` for usage beginning at this hour.
    endHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: `[YYYY-MM-DDThh]` for usage ending **before** this hour. (optional)
    optionalParams := datadog.GetUsageAnalyzedLogsOptionalParameters{
        EndHr: &endHr,
    }

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.UsageMeteringApi.GetUsageAnalyzedLogs(ctx, startHr, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetUsageAnalyzedLogs`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsageAnalyzedLogs`: UsageAnalyzedLogsResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from UsageMeteringApi.GetUsageAnalyzedLogs:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**startHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: &#x60;[YYYY-MM-DDThh]&#x60; for usage beginning at this hour. | 


### Optional Parameters


Other parameters are passed through a pointer to a GetUsageAnalyzedLogsOptionalParameters struct.


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**endHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: &#x60;[YYYY-MM-DDThh]&#x60; for usage ending **before** this hour. | 

### Return type

[**UsageAnalyzedLogsResponse**](UsageAnalyzedLogsResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;datetime-format=rfc3339

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsageAttribution

> UsageAttributionResponse GetUsageAttribution(ctx, startMonth, fields, datadog.GetUsageAttributionOptionalParameters{})

Get Usage Attribution.

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
    ctx := datadog.NewDefaultContext(context.Background())

    startMonth := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to month: `[YYYY-MM]` for usage beginning in this month. Maximum of 15 months ago.
    fields := datadog.UsageAttributionSupportedMetrics("custom_timeseries_usage") // UsageAttributionSupportedMetrics | Comma-separated list of usage types to return, or `*` for all usage types.
    endMonth := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to month: `[YYYY-MM]` for usage ending this month. (optional)
    sortDirection := datadog.UsageSortDirection("desc") // UsageSortDirection | The direction to sort by: `[desc, asc]`. (optional) (default to "desc")
    sortName := datadog.UsageAttributionSort("api_percentage") // UsageAttributionSort | The field to sort by. (optional) (default to "custom_timeseries_usage")
    includeDescendants := true // bool | Include child org usage in the response. Defaults to false. (optional) (default to false)
    optionalParams := datadog.GetUsageAttributionOptionalParameters{
        EndMonth: &endMonth,
        SortDirection: &sortDirection,
        SortName: &sortName,
        IncludeDescendants: &includeDescendants,
    }

    configuration := datadog.NewConfiguration()
    configuration.SetUnstableOperationEnabled("GetUsageAttribution", true)

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.UsageMeteringApi.GetUsageAttribution(ctx, startMonth, fields, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetUsageAttribution`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsageAttribution`: UsageAttributionResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from UsageMeteringApi.GetUsageAttribution:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**startMonth** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to month: &#x60;[YYYY-MM]&#x60; for usage beginning in this month. Maximum of 15 months ago. |  |
**fields** | [**UsageAttributionSupportedMetrics**](UsageAttributionSupportedMetrics.md) | Comma-separated list of usage types to return, or &#x60;*&#x60; for all usage types. | 


### Optional Parameters


Other parameters are passed through a pointer to a GetUsageAttributionOptionalParameters struct.


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**endMonth** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to month: &#x60;[YYYY-MM]&#x60; for usage ending this month. | 
**sortDirection** | [**UsageSortDirection**](UsageSortDirection.md) | The direction to sort by: &#x60;[desc, asc]&#x60;. | [default to &quot;desc&quot;]
**sortName** | [**UsageAttributionSort**](UsageAttributionSort.md) | The field to sort by. | [default to &quot;custom_timeseries_usage&quot;]
**includeDescendants** | **bool** | Include child org usage in the response. Defaults to false. | [default to false]

### Return type

[**UsageAttributionResponse**](UsageAttributionResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;datetime-format=rfc3339

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsageAuditLogs

> UsageAuditLogsResponse GetUsageAuditLogs(ctx, startHr, datadog.GetUsageAuditLogsOptionalParameters{})

Get hourly usage for audit logs.

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
    ctx := datadog.NewDefaultContext(context.Background())

    startHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: `[YYYY-MM-DDThh]` for usage beginning at this hour.
    endHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: `[YYYY-MM-DDThh]` for usage ending **before** this hour. (optional)
    optionalParams := datadog.GetUsageAuditLogsOptionalParameters{
        EndHr: &endHr,
    }

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.UsageMeteringApi.GetUsageAuditLogs(ctx, startHr, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetUsageAuditLogs`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsageAuditLogs`: UsageAuditLogsResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from UsageMeteringApi.GetUsageAuditLogs:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**startHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: &#x60;[YYYY-MM-DDThh]&#x60; for usage beginning at this hour. | 


### Optional Parameters


Other parameters are passed through a pointer to a GetUsageAuditLogsOptionalParameters struct.


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**endHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: &#x60;[YYYY-MM-DDThh]&#x60; for usage ending **before** this hour. | 

### Return type

[**UsageAuditLogsResponse**](UsageAuditLogsResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;datetime-format=rfc3339

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsageBillableSummary

> UsageBillableSummaryResponse GetUsageBillableSummary(ctx, datadog.GetUsageBillableSummaryOptionalParameters{})

Get billable usage across your account.

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
    ctx := datadog.NewDefaultContext(context.Background())

    month := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to month: `[YYYY-MM]` for usage starting this month. (optional)
    optionalParams := datadog.GetUsageBillableSummaryOptionalParameters{
        Month: &month,
    }

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.UsageMeteringApi.GetUsageBillableSummary(ctx, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetUsageBillableSummary`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsageBillableSummary`: UsageBillableSummaryResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from UsageMeteringApi.GetUsageBillableSummary:\n%s\n", responseContent)
}
```

### Required Parameters




### Optional Parameters


Other parameters are passed through a pointer to a GetUsageBillableSummaryOptionalParameters struct.


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**month** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to month: &#x60;[YYYY-MM]&#x60; for usage starting this month. | 

### Return type

[**UsageBillableSummaryResponse**](UsageBillableSummaryResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;datetime-format=rfc3339

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsageCWS

> UsageCWSResponse GetUsageCWS(ctx, startHr, datadog.GetUsageCWSOptionalParameters{})

Get hourly usage for Cloud Workload Security.

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
    ctx := datadog.NewDefaultContext(context.Background())

    startHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: `[YYYY-MM-DDThh]` for usage beginning at this hour.
    endHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: `[YYYY-MM-DDThh]` for usage ending **before** this hour. (optional)
    optionalParams := datadog.GetUsageCWSOptionalParameters{
        EndHr: &endHr,
    }

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.UsageMeteringApi.GetUsageCWS(ctx, startHr, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetUsageCWS`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsageCWS`: UsageCWSResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from UsageMeteringApi.GetUsageCWS:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**startHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: &#x60;[YYYY-MM-DDThh]&#x60; for usage beginning at this hour. | 


### Optional Parameters


Other parameters are passed through a pointer to a GetUsageCWSOptionalParameters struct.


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**endHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: &#x60;[YYYY-MM-DDThh]&#x60; for usage ending **before** this hour. | 

### Return type

[**UsageCWSResponse**](UsageCWSResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;datetime-format=rfc3339

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsageCloudSecurityPostureManagement

> UsageCloudSecurityPostureManagementResponse GetUsageCloudSecurityPostureManagement(ctx, startHr, datadog.GetUsageCloudSecurityPostureManagementOptionalParameters{})

Get hourly usage for Cloud Security Posture Management (CSPM).

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
    ctx := datadog.NewDefaultContext(context.Background())

    startHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: `[YYYY-MM-DDThh]` for usage beginning at this hour.
    endHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: `[YYYY-MM-DDThh]` for usage ending **before** this hour. (optional)
    optionalParams := datadog.GetUsageCloudSecurityPostureManagementOptionalParameters{
        EndHr: &endHr,
    }

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.UsageMeteringApi.GetUsageCloudSecurityPostureManagement(ctx, startHr, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetUsageCloudSecurityPostureManagement`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsageCloudSecurityPostureManagement`: UsageCloudSecurityPostureManagementResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from UsageMeteringApi.GetUsageCloudSecurityPostureManagement:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**startHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: &#x60;[YYYY-MM-DDThh]&#x60; for usage beginning at this hour. | 


### Optional Parameters


Other parameters are passed through a pointer to a GetUsageCloudSecurityPostureManagementOptionalParameters struct.


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**endHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: &#x60;[YYYY-MM-DDThh]&#x60; for usage ending **before** this hour. | 

### Return type

[**UsageCloudSecurityPostureManagementResponse**](UsageCloudSecurityPostureManagementResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;datetime-format=rfc3339

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsageDBM

> UsageDBMResponse GetUsageDBM(ctx, startHr, datadog.GetUsageDBMOptionalParameters{})

Get hourly usage for Database Monitoring

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
    ctx := datadog.NewDefaultContext(context.Background())

    startHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: `[YYYY-MM-DDThh]` for usage beginning at this hour.
    endHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: `[YYYY-MM-DDThh]` for usage ending **before** this hour. (optional)
    optionalParams := datadog.GetUsageDBMOptionalParameters{
        EndHr: &endHr,
    }

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.UsageMeteringApi.GetUsageDBM(ctx, startHr, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetUsageDBM`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsageDBM`: UsageDBMResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from UsageMeteringApi.GetUsageDBM:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**startHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: &#x60;[YYYY-MM-DDThh]&#x60; for usage beginning at this hour. | 


### Optional Parameters


Other parameters are passed through a pointer to a GetUsageDBMOptionalParameters struct.


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**endHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: &#x60;[YYYY-MM-DDThh]&#x60; for usage ending **before** this hour. | 

### Return type

[**UsageDBMResponse**](UsageDBMResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;datetime-format=rfc3339

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsageFargate

> UsageFargateResponse GetUsageFargate(ctx, startHr, datadog.GetUsageFargateOptionalParameters{})

Get hourly usage for [Fargate](https://docs.datadoghq.com/integrations/ecs_fargate/).

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
    ctx := datadog.NewDefaultContext(context.Background())

    startHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage beginning at this hour.
    endHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage ending **before** this hour. (optional)
    optionalParams := datadog.GetUsageFargateOptionalParameters{
        EndHr: &endHr,
    }

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.UsageMeteringApi.GetUsageFargate(ctx, startHr, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetUsageFargate`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsageFargate`: UsageFargateResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from UsageMeteringApi.GetUsageFargate:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**startHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage beginning at this hour. | 


### Optional Parameters


Other parameters are passed through a pointer to a GetUsageFargateOptionalParameters struct.


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**endHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage ending **before** this hour. | 

### Return type

[**UsageFargateResponse**](UsageFargateResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;datetime-format=rfc3339

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsageHosts

> UsageHostsResponse GetUsageHosts(ctx, startHr, datadog.GetUsageHostsOptionalParameters{})

Get hourly usage for hosts and containers.

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
    ctx := datadog.NewDefaultContext(context.Background())

    startHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage beginning at this hour.
    endHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage ending **before** this hour. (optional)
    optionalParams := datadog.GetUsageHostsOptionalParameters{
        EndHr: &endHr,
    }

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.UsageMeteringApi.GetUsageHosts(ctx, startHr, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetUsageHosts`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsageHosts`: UsageHostsResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from UsageMeteringApi.GetUsageHosts:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**startHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage beginning at this hour. | 


### Optional Parameters


Other parameters are passed through a pointer to a GetUsageHostsOptionalParameters struct.


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**endHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage ending **before** this hour. | 

### Return type

[**UsageHostsResponse**](UsageHostsResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;datetime-format=rfc3339

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsageIndexedSpans

> UsageIndexedSpansResponse GetUsageIndexedSpans(ctx, startHr, datadog.GetUsageIndexedSpansOptionalParameters{})

Get hourly usage for indexed spans.

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
    ctx := datadog.NewDefaultContext(context.Background())

    startHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: `[YYYY-MM-DDThh]` for usage beginning at this hour.
    endHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: `[YYYY-MM-DDThh]` for usage ending **before** this hour. (optional)
    optionalParams := datadog.GetUsageIndexedSpansOptionalParameters{
        EndHr: &endHr,
    }

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.UsageMeteringApi.GetUsageIndexedSpans(ctx, startHr, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetUsageIndexedSpans`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsageIndexedSpans`: UsageIndexedSpansResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from UsageMeteringApi.GetUsageIndexedSpans:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**startHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: &#x60;[YYYY-MM-DDThh]&#x60; for usage beginning at this hour. | 


### Optional Parameters


Other parameters are passed through a pointer to a GetUsageIndexedSpansOptionalParameters struct.


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**endHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: &#x60;[YYYY-MM-DDThh]&#x60; for usage ending **before** this hour. | 

### Return type

[**UsageIndexedSpansResponse**](UsageIndexedSpansResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;datetime-format=rfc3339

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsageInternetOfThings

> UsageIoTResponse GetUsageInternetOfThings(ctx, startHr, datadog.GetUsageInternetOfThingsOptionalParameters{})

Get hourly usage for IoT.

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
    ctx := datadog.NewDefaultContext(context.Background())

    startHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: `[YYYY-MM-DDThh]` for usage beginning at this hour.
    endHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: `[YYYY-MM-DDThh]` for usage ending **before** this hour. (optional)
    optionalParams := datadog.GetUsageInternetOfThingsOptionalParameters{
        EndHr: &endHr,
    }

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.UsageMeteringApi.GetUsageInternetOfThings(ctx, startHr, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetUsageInternetOfThings`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsageInternetOfThings`: UsageIoTResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from UsageMeteringApi.GetUsageInternetOfThings:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**startHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: &#x60;[YYYY-MM-DDThh]&#x60; for usage beginning at this hour. | 


### Optional Parameters


Other parameters are passed through a pointer to a GetUsageInternetOfThingsOptionalParameters struct.


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**endHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: &#x60;[YYYY-MM-DDThh]&#x60; for usage ending **before** this hour. | 

### Return type

[**UsageIoTResponse**](UsageIoTResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;datetime-format=rfc3339

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsageLambda

> UsageLambdaResponse GetUsageLambda(ctx, startHr, datadog.GetUsageLambdaOptionalParameters{})

Get hourly usage for lambda.

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
    ctx := datadog.NewDefaultContext(context.Background())

    startHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage beginning at this hour.
    endHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage ending **before** this hour. (optional)
    optionalParams := datadog.GetUsageLambdaOptionalParameters{
        EndHr: &endHr,
    }

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.UsageMeteringApi.GetUsageLambda(ctx, startHr, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetUsageLambda`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsageLambda`: UsageLambdaResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from UsageMeteringApi.GetUsageLambda:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**startHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage beginning at this hour. | 


### Optional Parameters


Other parameters are passed through a pointer to a GetUsageLambdaOptionalParameters struct.


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**endHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage ending **before** this hour. | 

### Return type

[**UsageLambdaResponse**](UsageLambdaResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;datetime-format=rfc3339

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsageLogs

> UsageLogsResponse GetUsageLogs(ctx, startHr, datadog.GetUsageLogsOptionalParameters{})

Get hourly usage for logs.

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
    ctx := datadog.NewDefaultContext(context.Background())

    startHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage beginning at this hour.
    endHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage ending **before** this hour. (optional)
    optionalParams := datadog.GetUsageLogsOptionalParameters{
        EndHr: &endHr,
    }

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.UsageMeteringApi.GetUsageLogs(ctx, startHr, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetUsageLogs`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsageLogs`: UsageLogsResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from UsageMeteringApi.GetUsageLogs:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**startHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage beginning at this hour. | 


### Optional Parameters


Other parameters are passed through a pointer to a GetUsageLogsOptionalParameters struct.


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**endHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage ending **before** this hour. | 

### Return type

[**UsageLogsResponse**](UsageLogsResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;datetime-format=rfc3339

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsageLogsByIndex

> UsageLogsByIndexResponse GetUsageLogsByIndex(ctx, startHr, datadog.GetUsageLogsByIndexOptionalParameters{})

Get hourly usage for logs by index.

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
    ctx := datadog.NewDefaultContext(context.Background())

    startHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage beginning at this hour.
    endHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage ending **before** this hour. (optional)
    indexName := []string{"Inner_example"} // []string | Comma-separated list of log index names. (optional)
    optionalParams := datadog.GetUsageLogsByIndexOptionalParameters{
        EndHr: &endHr,
        IndexName: &indexName,
    }

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.UsageMeteringApi.GetUsageLogsByIndex(ctx, startHr, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetUsageLogsByIndex`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsageLogsByIndex`: UsageLogsByIndexResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from UsageMeteringApi.GetUsageLogsByIndex:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**startHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage beginning at this hour. | 


### Optional Parameters


Other parameters are passed through a pointer to a GetUsageLogsByIndexOptionalParameters struct.


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**endHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage ending **before** this hour. | 
**indexName** | **[]string** | Comma-separated list of log index names. | 

### Return type

[**UsageLogsByIndexResponse**](UsageLogsByIndexResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;datetime-format=rfc3339

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsageLogsByRetention

> UsageLogsByRetentionResponse GetUsageLogsByRetention(ctx, startHr, datadog.GetUsageLogsByRetentionOptionalParameters{})

Get hourly usage for indexed logs by retention period.

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
    ctx := datadog.NewDefaultContext(context.Background())

    startHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: `[YYYY-MM-DDThh]` for usage beginning at this hour.
    endHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: `[YYYY-MM-DDThh]` for usage ending **before** this hour. (optional)
    optionalParams := datadog.GetUsageLogsByRetentionOptionalParameters{
        EndHr: &endHr,
    }

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.UsageMeteringApi.GetUsageLogsByRetention(ctx, startHr, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetUsageLogsByRetention`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsageLogsByRetention`: UsageLogsByRetentionResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from UsageMeteringApi.GetUsageLogsByRetention:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**startHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: &#x60;[YYYY-MM-DDThh]&#x60; for usage beginning at this hour. | 


### Optional Parameters


Other parameters are passed through a pointer to a GetUsageLogsByRetentionOptionalParameters struct.


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**endHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: &#x60;[YYYY-MM-DDThh]&#x60; for usage ending **before** this hour. | 

### Return type

[**UsageLogsByRetentionResponse**](UsageLogsByRetentionResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;datetime-format=rfc3339

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsageNetworkFlows

> UsageNetworkFlowsResponse GetUsageNetworkFlows(ctx, startHr, datadog.GetUsageNetworkFlowsOptionalParameters{})

Get hourly usage for network flows.

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
    ctx := datadog.NewDefaultContext(context.Background())

    startHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: `[YYYY-MM-DDThh]` for usage beginning at this hour.
    endHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: `[YYYY-MM-DDThh]` for usage ending **before** this hour. (optional)
    optionalParams := datadog.GetUsageNetworkFlowsOptionalParameters{
        EndHr: &endHr,
    }

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.UsageMeteringApi.GetUsageNetworkFlows(ctx, startHr, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetUsageNetworkFlows`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsageNetworkFlows`: UsageNetworkFlowsResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from UsageMeteringApi.GetUsageNetworkFlows:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**startHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: &#x60;[YYYY-MM-DDThh]&#x60; for usage beginning at this hour. | 


### Optional Parameters


Other parameters are passed through a pointer to a GetUsageNetworkFlowsOptionalParameters struct.


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**endHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: &#x60;[YYYY-MM-DDThh]&#x60; for usage ending **before** this hour. | 

### Return type

[**UsageNetworkFlowsResponse**](UsageNetworkFlowsResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;datetime-format=rfc3339

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsageNetworkHosts

> UsageNetworkHostsResponse GetUsageNetworkHosts(ctx, startHr, datadog.GetUsageNetworkHostsOptionalParameters{})

Get hourly usage for network hosts.

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
    ctx := datadog.NewDefaultContext(context.Background())

    startHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage beginning at this hour.
    endHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage ending **before** this hour. (optional)
    optionalParams := datadog.GetUsageNetworkHostsOptionalParameters{
        EndHr: &endHr,
    }

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.UsageMeteringApi.GetUsageNetworkHosts(ctx, startHr, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetUsageNetworkHosts`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsageNetworkHosts`: UsageNetworkHostsResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from UsageMeteringApi.GetUsageNetworkHosts:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**startHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage beginning at this hour. | 


### Optional Parameters


Other parameters are passed through a pointer to a GetUsageNetworkHostsOptionalParameters struct.


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**endHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage ending **before** this hour. | 

### Return type

[**UsageNetworkHostsResponse**](UsageNetworkHostsResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;datetime-format=rfc3339

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsageProfiling

> UsageProfilingResponse GetUsageProfiling(ctx, startHr, datadog.GetUsageProfilingOptionalParameters{})

Get hourly usage for profiled hosts.

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
    ctx := datadog.NewDefaultContext(context.Background())

    startHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: `[YYYY-MM-DDThh]` for usage beginning at this hour.
    endHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: `[YYYY-MM-DDThh]` for usage ending **before** this hour. (optional)
    optionalParams := datadog.GetUsageProfilingOptionalParameters{
        EndHr: &endHr,
    }

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.UsageMeteringApi.GetUsageProfiling(ctx, startHr, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetUsageProfiling`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsageProfiling`: UsageProfilingResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from UsageMeteringApi.GetUsageProfiling:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**startHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: &#x60;[YYYY-MM-DDThh]&#x60; for usage beginning at this hour. | 


### Optional Parameters


Other parameters are passed through a pointer to a GetUsageProfilingOptionalParameters struct.


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**endHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: &#x60;[YYYY-MM-DDThh]&#x60; for usage ending **before** this hour. | 

### Return type

[**UsageProfilingResponse**](UsageProfilingResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;datetime-format=rfc3339

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsageRumSessions

> UsageRumSessionsResponse GetUsageRumSessions(ctx, startHr, datadog.GetUsageRumSessionsOptionalParameters{})

Get hourly usage for [RUM](https://docs.datadoghq.com/real_user_monitoring/) Sessions.

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
    ctx := datadog.NewDefaultContext(context.Background())

    startHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage beginning at this hour.
    endHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage ending **before** this hour. (optional)
    type_ := "type__example" // string | RUM type: `[browser, mobile]`. Defaults to `browser`. (optional)
    optionalParams := datadog.GetUsageRumSessionsOptionalParameters{
        EndHr: &endHr,
        Type_: &type_,
    }

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.UsageMeteringApi.GetUsageRumSessions(ctx, startHr, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetUsageRumSessions`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsageRumSessions`: UsageRumSessionsResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from UsageMeteringApi.GetUsageRumSessions:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**startHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage beginning at this hour. | 


### Optional Parameters


Other parameters are passed through a pointer to a GetUsageRumSessionsOptionalParameters struct.


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**endHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage ending **before** this hour. | 
**type_** | **string** | RUM type: &#x60;[browser, mobile]&#x60;. Defaults to &#x60;browser&#x60;. | 

### Return type

[**UsageRumSessionsResponse**](UsageRumSessionsResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;datetime-format=rfc3339

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsageSNMP

> UsageSNMPResponse GetUsageSNMP(ctx, startHr, datadog.GetUsageSNMPOptionalParameters{})

Get hourly usage for SNMP devices.

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
    ctx := datadog.NewDefaultContext(context.Background())

    startHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: `[YYYY-MM-DDThh]` for usage beginning at this hour.
    endHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: `[YYYY-MM-DDThh]` for usage ending **before** this hour. (optional)
    optionalParams := datadog.GetUsageSNMPOptionalParameters{
        EndHr: &endHr,
    }

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.UsageMeteringApi.GetUsageSNMP(ctx, startHr, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetUsageSNMP`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsageSNMP`: UsageSNMPResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from UsageMeteringApi.GetUsageSNMP:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**startHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: &#x60;[YYYY-MM-DDThh]&#x60; for usage beginning at this hour. | 


### Optional Parameters


Other parameters are passed through a pointer to a GetUsageSNMPOptionalParameters struct.


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**endHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: &#x60;[YYYY-MM-DDThh]&#x60; for usage ending **before** this hour. | 

### Return type

[**UsageSNMPResponse**](UsageSNMPResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;datetime-format=rfc3339

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsageSummary

> UsageSummaryResponse GetUsageSummary(ctx, startMonth, datadog.GetUsageSummaryOptionalParameters{})

Get usage across your multi-org account. You must have the multi-org feature enabled.

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
    ctx := datadog.NewDefaultContext(context.Background())

    startMonth := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to month: `[YYYY-MM]` for usage beginning in this month. Maximum of 15 months ago.
    endMonth := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to month: `[YYYY-MM]` for usage ending this month. (optional)
    includeOrgDetails := true // bool | Include usage summaries for each sub-org. (optional)
    optionalParams := datadog.GetUsageSummaryOptionalParameters{
        EndMonth: &endMonth,
        IncludeOrgDetails: &includeOrgDetails,
    }

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.UsageMeteringApi.GetUsageSummary(ctx, startMonth, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetUsageSummary`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsageSummary`: UsageSummaryResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from UsageMeteringApi.GetUsageSummary:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**startMonth** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to month: &#x60;[YYYY-MM]&#x60; for usage beginning in this month. Maximum of 15 months ago. | 


### Optional Parameters


Other parameters are passed through a pointer to a GetUsageSummaryOptionalParameters struct.


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**endMonth** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to month: &#x60;[YYYY-MM]&#x60; for usage ending this month. | 
**includeOrgDetails** | **bool** | Include usage summaries for each sub-org. | 

### Return type

[**UsageSummaryResponse**](UsageSummaryResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;datetime-format=rfc3339

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsageSynthetics

> UsageSyntheticsResponse GetUsageSynthetics(ctx, startHr, datadog.GetUsageSyntheticsOptionalParameters{})

Get hourly usage for [Synthetics checks](https://docs.datadoghq.com/synthetics/).

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
    ctx := datadog.NewDefaultContext(context.Background())

    startHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage beginning at this hour.
    endHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage ending **before** this hour. (optional)
    optionalParams := datadog.GetUsageSyntheticsOptionalParameters{
        EndHr: &endHr,
    }

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.UsageMeteringApi.GetUsageSynthetics(ctx, startHr, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetUsageSynthetics`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsageSynthetics`: UsageSyntheticsResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from UsageMeteringApi.GetUsageSynthetics:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**startHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage beginning at this hour. | 


### Optional Parameters


Other parameters are passed through a pointer to a GetUsageSyntheticsOptionalParameters struct.


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**endHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage ending **before** this hour. | 

### Return type

[**UsageSyntheticsResponse**](UsageSyntheticsResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;datetime-format=rfc3339

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsageSyntheticsAPI

> UsageSyntheticsAPIResponse GetUsageSyntheticsAPI(ctx, startHr, datadog.GetUsageSyntheticsAPIOptionalParameters{})

Get hourly usage for [synthetics API checks](https://docs.datadoghq.com/synthetics/).

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
    ctx := datadog.NewDefaultContext(context.Background())

    startHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage beginning at this hour.
    endHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage ending **before** this hour. (optional)
    optionalParams := datadog.GetUsageSyntheticsAPIOptionalParameters{
        EndHr: &endHr,
    }

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.UsageMeteringApi.GetUsageSyntheticsAPI(ctx, startHr, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetUsageSyntheticsAPI`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsageSyntheticsAPI`: UsageSyntheticsAPIResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from UsageMeteringApi.GetUsageSyntheticsAPI:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**startHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage beginning at this hour. | 


### Optional Parameters


Other parameters are passed through a pointer to a GetUsageSyntheticsAPIOptionalParameters struct.


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**endHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage ending **before** this hour. | 

### Return type

[**UsageSyntheticsAPIResponse**](UsageSyntheticsAPIResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;datetime-format=rfc3339

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsageSyntheticsBrowser

> UsageSyntheticsBrowserResponse GetUsageSyntheticsBrowser(ctx, startHr, datadog.GetUsageSyntheticsBrowserOptionalParameters{})

Get hourly usage for synthetics browser checks.

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
    ctx := datadog.NewDefaultContext(context.Background())

    startHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage beginning at this hour.
    endHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage ending **before** this hour. (optional)
    optionalParams := datadog.GetUsageSyntheticsBrowserOptionalParameters{
        EndHr: &endHr,
    }

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.UsageMeteringApi.GetUsageSyntheticsBrowser(ctx, startHr, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetUsageSyntheticsBrowser`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsageSyntheticsBrowser`: UsageSyntheticsBrowserResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from UsageMeteringApi.GetUsageSyntheticsBrowser:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**startHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage beginning at this hour. | 


### Optional Parameters


Other parameters are passed through a pointer to a GetUsageSyntheticsBrowserOptionalParameters struct.


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**endHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage ending **before** this hour. | 

### Return type

[**UsageSyntheticsBrowserResponse**](UsageSyntheticsBrowserResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;datetime-format=rfc3339

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsageTimeseries

> UsageTimeseriesResponse GetUsageTimeseries(ctx, startHr, datadog.GetUsageTimeseriesOptionalParameters{})

Get hourly usage for [custom metrics](https://docs.datadoghq.com/developers/metrics/custom_metrics/).

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
    ctx := datadog.NewDefaultContext(context.Background())

    startHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage beginning at this hour.
    endHr := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage ending **before** this hour. (optional)
    optionalParams := datadog.GetUsageTimeseriesOptionalParameters{
        EndHr: &endHr,
    }

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.UsageMeteringApi.GetUsageTimeseries(ctx, startHr, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetUsageTimeseries`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsageTimeseries`: UsageTimeseriesResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from UsageMeteringApi.GetUsageTimeseries:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**startHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage beginning at this hour. | 


### Optional Parameters


Other parameters are passed through a pointer to a GetUsageTimeseriesOptionalParameters struct.


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**endHr** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to hour: [YYYY-MM-DDThh] for usage ending **before** this hour. | 

### Return type

[**UsageTimeseriesResponse**](UsageTimeseriesResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;datetime-format=rfc3339

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsageTopAvgMetrics

> UsageTopAvgMetricsResponse GetUsageTopAvgMetrics(ctx, datadog.GetUsageTopAvgMetricsOptionalParameters{})

Get all [custom metrics](https://docs.datadoghq.com/developers/metrics/custom_metrics/) by hourly average. Use the month parameter to get a month-to-date data resolution or use the day parameter to get a daily resolution. One of the two is required, and only one of the two is allowed.

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
    ctx := datadog.NewDefaultContext(context.Background())

    month := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to month: [YYYY-MM] for usage beginning at this hour. (Either month or day should be specified, but not both) (optional)
    day := time.Now() // time.Time | Datetime in ISO-8601 format, UTC, precise to day: [YYYY-MM-DD] for usage beginning at this hour. (Either month or day should be specified, but not both) (optional)
    names := []string{"Inner_example"} // []string | Comma-separated list of metric names. (optional)
    limit := int32(56) // int32 | Maximum number of results to return (between 1 and 5000) - defaults to 500 results if limit not specified. (optional) (default to 500)
    nextRecordId := "nextRecordId_example" // string | List following results with a next_record_id provided in the previous query. (optional)
    optionalParams := datadog.GetUsageTopAvgMetricsOptionalParameters{
        Month: &month,
        Day: &day,
        Names: &names,
        Limit: &limit,
        NextRecordId: &nextRecordId,
    }

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.UsageMeteringApi.GetUsageTopAvgMetrics(ctx, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetUsageTopAvgMetrics`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsageTopAvgMetrics`: UsageTopAvgMetricsResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from UsageMeteringApi.GetUsageTopAvgMetrics:\n%s\n", responseContent)
}
```

### Required Parameters




### Optional Parameters


Other parameters are passed through a pointer to a GetUsageTopAvgMetricsOptionalParameters struct.


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**month** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to month: [YYYY-MM] for usage beginning at this hour. (Either month or day should be specified, but not both) | 
**day** | **time.Time** | Datetime in ISO-8601 format, UTC, precise to day: [YYYY-MM-DD] for usage beginning at this hour. (Either month or day should be specified, but not both) | 
**names** | **[]string** | Comma-separated list of metric names. | 
**limit** | **int32** | Maximum number of results to return (between 1 and 5000) - defaults to 500 results if limit not specified. | [default to 500]
**nextRecordId** | **string** | List following results with a next_record_id provided in the previous query. | 

### Return type

[**UsageTopAvgMetricsResponse**](UsageTopAvgMetricsResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;datetime-format=rfc3339

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

