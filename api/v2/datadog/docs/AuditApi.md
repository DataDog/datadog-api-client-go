# AuditApi

All URIs are relative to *https://api.datadoghq.com*

| Method                                             | HTTP request                         | Description                     |
| -------------------------------------------------- | ------------------------------------ | ------------------------------- |
| [**ListAuditLogs**](AuditApi.md#ListAuditLogs)     | **Get** /api/v2/audit/events         | Get a list of Audit Logs events |
| [**SearchAuditLogs**](AuditApi.md#SearchAuditLogs) | **Post** /api/v2/audit/events/search | Search Audit Logs events        |

## ListAuditLogs

> AuditLogsEventsResponse ListAuditLogs(ctx, datadog.ListAuditLogsOptionalParameters{})

List endpoint returns events that match a Audit Logs search query.
[Results are paginated][1].

Use this endpoint to see your latest Audit Logs events.

[1]: https://docs.datadoghq.com/logs/guide/collect-multiple-logs-with-pagination

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

    filterQuery := "@type:session @application_id:xxxx" // string | Search query following Audit Logs syntax. (optional)
    filterFrom := time.Now() // time.Time | Minimum timestamp for requested events. (optional)
    filterTo := time.Now() // time.Time | Maximum timestamp for requested events. (optional)
    sort := datadog.AuditLogsSort("timestamp") // AuditLogsSort | Order of events in results. (optional)
    pageCursor := "eyJzdGFydEF0IjoiQVFBQUFYS2tMS3pPbm40NGV3QUFBQUJCV0V0clRFdDZVbG8zY3pCRmNsbHJiVmxDWlEifQ==" // string | List following results with a cursor provided in the previous query. (optional)
    pageLimit := int32(25) // int32 | Maximum number of events in the response. (optional) (default to 10)
    optionalParams := datadog.ListAuditLogsOptionalParameters{
        FilterQuery: &filterQuery,
        FilterFrom: &filterFrom,
        FilterTo: &filterTo,
        Sort: &sort,
        PageCursor: &pageCursor,
        PageLimit: &pageLimit,
    }

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.AuditApi.ListAuditLogs(ctx, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuditApi.ListAuditLogs`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAuditLogs`: AuditLogsEventsResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from AuditApi.ListAuditLogs:\n%s\n", responseContent)
}
```

### Required Parameters

### Optional Parameters

Other parameters are passed through a pointer to a ListAuditLogsOptionalParameters struct.

| Name            | Type                                  | Description                                                          | Notes           |
| --------------- | ------------------------------------- | -------------------------------------------------------------------- | --------------- |
| **filterQuery** | **string**                            | Search query following Audit Logs syntax.                            |
| **filterFrom**  | **time.Time**                         | Minimum timestamp for requested events.                              |
| **filterTo**    | **time.Time**                         | Maximum timestamp for requested events.                              |
| **sort**        | [**AuditLogsSort**](AuditLogsSort.md) | Order of events in results.                                          |
| **pageCursor**  | **string**                            | List following results with a cursor provided in the previous query. |
| **pageLimit**   | **int32**                             | Maximum number of events in the response.                            | [default to 10] |

### Return type

[**AuditLogsEventsResponse**](AuditLogsEventsResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SearchAuditLogs

> AuditLogsEventsResponse SearchAuditLogs(ctx, datadog.SearchAuditLogsOptionalParameters{})

List endpoint returns Audit Logs events that match an Audit search query.
[Results are paginated][1].

Use this endpoint to build complex Audit Logs events filtering and search.

[1]: https://docs.datadoghq.com/logs/guide/collect-multiple-logs-with-pagination

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

    body := *datadog.NewAuditLogsSearchEventsRequest() // AuditLogsSearchEventsRequest |  (optional)
    optionalParams := datadog.SearchAuditLogsOptionalParameters{
        Body: &body,
    }

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.AuditApi.SearchAuditLogs(ctx, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuditApi.SearchAuditLogs`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchAuditLogs`: AuditLogsEventsResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from AuditApi.SearchAuditLogs:\n%s\n", responseContent)
}
```

### Required Parameters

### Optional Parameters

Other parameters are passed through a pointer to a SearchAuditLogsOptionalParameters struct.

| Name     | Type                                                                | Description | Notes |
| -------- | ------------------------------------------------------------------- | ----------- | ----- |
| **body** | [**AuditLogsSearchEventsRequest**](AuditLogsSearchEventsRequest.md) |             |

### Return type

[**AuditLogsEventsResponse**](AuditLogsEventsResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
