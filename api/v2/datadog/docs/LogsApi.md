# LogsApi

All URIs are relative to *https://api.datadoghq.com*

| Method                                        | HTTP request                              | Description        |
| --------------------------------------------- | ----------------------------------------- | ------------------ |
| [**AggregateLogs**](LogsApi.md#AggregateLogs) | **Post** /api/v2/logs/analytics/aggregate | Aggregate events   |
| [**ListLogs**](LogsApi.md#ListLogs)           | **Post** /api/v2/logs/events/search       | Search logs        |
| [**ListLogsGet**](LogsApi.md#ListLogsGet)     | **Get** /api/v2/logs/events               | Get a list of logs |
| [**SubmitLog**](LogsApi.md#SubmitLog)         | **Post** /api/v2/logs                     | Send logs          |

## AggregateLogs

> LogsAggregateResponse AggregateLogs(ctx, body)

The API endpoint to aggregate events into buckets and compute metrics and timeseries.

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
        fmt.Fprintf(os.Stderr, "Error when calling `LogsApi.AggregateLogs`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AggregateLogs`: LogsAggregateResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from LogsApi.AggregateLogs:\n%s\n", responseContent)
}
```

### Required Parameters

| Name     | Type                                                | Description                                                                 | Notes |
| -------- | --------------------------------------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**  | **context.Context**                                 | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **body** | [**LogsAggregateRequest**](LogsAggregateRequest.md) |                                                                             |

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

List endpoint returns logs that match a log search query.
[Results are paginated][1].

Use this endpoint to build complex logs filtering and search.

**If you are considering archiving logs for your organization,
consider use of the Datadog archive capabilities instead of the log list API.
See [Datadog Logs Archive documentation][2].**

[1]: /logs/guide/collect-multiple-logs-with-pagination
[2]: https://docs.datadoghq.com/logs/archives

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
        fmt.Fprintf(os.Stderr, "Error when calling `LogsApi.ListLogs`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLogs`: LogsListResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from LogsApi.ListLogs:\n%s\n", responseContent)
}
```

### Required Parameters

### Optional Parameters

Other parameters are passed through a pointer to a ListLogsOptionalParameters struct.

| Name     | Type                                      | Description | Notes |
| -------- | ----------------------------------------- | ----------- | ----- |
| **body** | [**LogsListRequest**](LogsListRequest.md) |             |

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

List endpoint returns logs that match a log search query.
[Results are paginated][1].

Use this endpoint to see your latest logs.

**If you are considering archiving logs for your organization,
consider use of the Datadog archive capabilities instead of the log list API.
See [Datadog Logs Archive documentation][2].**

[1]: /logs/guide/collect-multiple-logs-with-pagination
[2]: https://docs.datadoghq.com/logs/archives

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
        fmt.Fprintf(os.Stderr, "Error when calling `LogsApi.ListLogsGet`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLogsGet`: LogsListResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from LogsApi.ListLogsGet:\n%s\n", responseContent)
}
```

### Required Parameters

### Optional Parameters

Other parameters are passed through a pointer to a ListLogsGetOptionalParameters struct.

| Name            | Type                        | Description                                                                                                 | Notes           |
| --------------- | --------------------------- | ----------------------------------------------------------------------------------------------------------- | --------------- |
| **filterQuery** | **string**                  | Search query following logs syntax.                                                                         |
| **filterIndex** | **string**                  | For customers with multiple indexes, the indexes to search Defaults to &#39;\*&#39; which means all indexes |
| **filterFrom**  | **time.Time**               | Minimum timestamp for requested logs.                                                                       |
| **filterTo**    | **time.Time**               | Maximum timestamp for requested logs.                                                                       |
| **sort**        | [**LogsSort**](LogsSort.md) | Order of logs in results.                                                                                   |
| **pageCursor**  | **string**                  | List following results with a cursor provided in the previous query.                                        |
| **pageLimit**   | **int32**                   | Maximum number of logs in the response.                                                                     | [default to 10] |

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

## SubmitLog

> interface{} SubmitLog(ctx, body, datadog.SubmitLogOptionalParameters{})

Send your logs to your Datadog platform over HTTP. Limits per HTTP request are:

- Maximum content size per payload (uncompressed): 5MB
- Maximum size for a single log: 1MB
- Maximum array size if sending multiple logs in an array: 1000 entries

Any log exceeding 1MB is accepted and truncated by Datadog:

- For a single log request, the API truncates the log at 1MB and returns a 2xx.
- For a multi-logs request, the API processes all logs, truncates only logs larger than 1MB, and returns a 2xx.

Datadog recommends sending your logs compressed.
Add the `Content-Encoding: gzip` header to the request when sending compressed logs.

The status codes answered by the HTTP API are:

- 202: Accepted: the request has been accepted for processing
- 400: Bad request (likely an issue in the payload formatting)
- 401: Unauthorized (likely a missing API Key)
- 403: Permission issue (likely using an invalid API Key)
- 408: Request Timeout, request should be retried after some time
- 413: Payload too large (batch is above 5MB uncompressed)
- 429: Too Many Requests, request should be retried after some time
- 500: Internal Server Error, the server encountered an unexpected condition that prevented it from fulfilling the request, request should be retried after some time
- 503: Service Unavailable, the server is not ready to handle the request probably because it is overloaded, request should be retried after some time

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

    body := []datadog.HTTPLogItem{*datadog.NewHTTPLogItem()} // []HTTPLogItem | Log to send (JSON format).
    contentEncoding := datadog.ContentEncoding("gzip") // ContentEncoding | HTTP header used to compress the media-type. (optional)
    ddtags := "env:prod,user:my-user" // string | Log tags can be passed as query parameters with `text/plain` content type. (optional)
    optionalParams := datadog.SubmitLogOptionalParameters{
        ContentEncoding: &contentEncoding,
        Ddtags: &ddtags,
    }

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.LogsApi.SubmitLog(ctx, body, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsApi.SubmitLog`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitLog`: interface{}
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from LogsApi.SubmitLog:\n%s\n", responseContent)
}
```

### Required Parameters

| Name     | Type                                | Description                                                                 | Notes |
| -------- | ----------------------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**  | **context.Context**                 | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **body** | [**[]HTTPLogItem**](HTTPLogItem.md) | Log to send (JSON format).                                                  |

### Optional Parameters

Other parameters are passed through a pointer to a SubmitLogOptionalParameters struct.

| Name                | Type                                      | Description                                                                          | Notes |
| ------------------- | ----------------------------------------- | ------------------------------------------------------------------------------------ | ----- |
| **contentEncoding** | [**ContentEncoding**](ContentEncoding.md) | HTTP header used to compress the media-type.                                         |
| **ddtags**          | **string**                                | Log tags can be passed as query parameters with &#x60;text/plain&#x60; content type. |

### Return type

**interface{}**

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/logplex-1, text/plain
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
