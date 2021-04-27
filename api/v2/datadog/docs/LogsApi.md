# \LogsApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AggregateLogs**](LogsApi.md#AggregateLogs) | **Post** /api/v2/logs/analytics/aggregate | Aggregate events
[**ListLogs**](LogsApi.md#ListLogs) | **Post** /api/v2/logs/events/search | Search logs
[**ListLogsGet**](LogsApi.md#ListLogsGet) | **Get** /api/v2/logs/events | Get a list of logs



## AggregateLogs

> LogsAggregateResponse AggregateLogs(ctx, body)

Aggregate events



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

    body := *datadog.NewLogsAggregateRequest() // LogsAggregateRequest | 

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.LogsApi.AggregateLogs(ctx, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsApi.AggregateLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AggregateLogs`: LogsAggregateResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from LogsApi.AggregateLogs:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**LogsAggregateRequest**](LogsAggregateRequest.md) |  | 

### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**LogsAggregateResponse**](LogsAggregateResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLogs

> LogsListResponse ListLogs(ctx, datadog.ListLogsOptionalParameters{})

Search logs



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

    body := *datadog.NewLogsListRequest() // LogsListRequest |  (optional)
    optionalParams := datadog.ListLogsOptionalParameters{
        Body: &body,
    }

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.LogsApi.ListLogs(ctx, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsApi.ListLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLogs`: LogsListResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from LogsApi.ListLogs:\n%s\n", responseContent)
}
```

### Required Parameters



### Optional Parameters


Other parameters are passed through a pointer to a ListLogsOptionalParameters struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**LogsListRequest**](LogsListRequest.md) |  | 

### Return type

[**LogsListResponse**](LogsListResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLogsGet

> LogsListResponse ListLogsGet(ctx, datadog.ListLogsGetOptionalParameters{})

Get a list of logs



### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    "time"
    datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
    ctx := datadog.NewDefaultContext(context.Background())

    filterQuery := "@datacenter:us @role:db" // string | Search query following logs syntax. (optional)
    filterIndex := "main" // string | For customers with multiple indexes, the indexes to search Defaults to '*' which means all indexes (optional)
    filterFrom := time.Now() // time.Time | Minimum timestamp for requested logs. (optional)
    filterTo := time.Now() // time.Time | Maximum timestamp for requested logs. (optional)
    sort := datadog.LogsSort("timestamp") // LogsSort | Order of logs in results. (optional)
    pageCursor := "eyJzdGFydEF0IjoiQVFBQUFYS2tMS3pPbm40NGV3QUFBQUJCV0V0clRFdDZVbG8zY3pCRmNsbHJiVmxDWlEifQ==" // string | List following results with a cursor provided in the previous query. (optional)
    pageLimit := int32(25) // int32 | Maximum number of logs in the response. (optional) (default to 10)
    optionalParams := datadog.ListLogsGetOptionalParameters{
        FilterQuery: &filterQuery,
        FilterIndex: &filterIndex,
        FilterFrom: &filterFrom,
        FilterTo: &filterTo,
        Sort: &sort,
        PageCursor: &pageCursor,
        PageLimit: &pageLimit,
    }

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.LogsApi.ListLogsGet(ctx, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsApi.ListLogsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLogsGet`: LogsListResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from LogsApi.ListLogsGet:\n%s\n", responseContent)
}
```

### Required Parameters



### Optional Parameters


Other parameters are passed through a pointer to a ListLogsGetOptionalParameters struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**filterQuery** | **string** | Search query following logs syntax. | 
**filterIndex** | **string** | For customers with multiple indexes, the indexes to search Defaults to &#39;*&#39; which means all indexes | 
**filterFrom** | **time.Time** | Minimum timestamp for requested logs. | 
**filterTo** | **time.Time** | Maximum timestamp for requested logs. | 
**sort** | [**LogsSort**](LogsSort.md) | Order of logs in results. | 
**pageCursor** | **string** | List following results with a cursor provided in the previous query. | 
**pageLimit** | **int32** | Maximum number of logs in the response. | [default to 10]

### Return type

[**LogsListResponse**](LogsListResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

