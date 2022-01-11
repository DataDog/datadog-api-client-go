# LogsApi

All URIs are relative to *https://api.datadoghq.com*

| Method                                | HTTP request                       | Description |
| ------------------------------------- | ---------------------------------- | ----------- |
| [**ListLogs**](LogsApi.md#ListLogs)   | **Post** /api/v1/logs-queries/list | Search logs |
| [**SubmitLog**](LogsApi.md#SubmitLog) | **Post** /v1/input                 | Send logs   |

## ListLogs

> LogsListResponse ListLogs(ctx, body)

List endpoint returns logs that match a log search query.
[Results are paginated][1].

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
    datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
    ctx := datadog.NewDefaultContext(context.Background())

    body := *datadog.NewLogsListRequest(*datadog.NewLogsListRequestTime(time.Now(), time.Now())) // LogsListRequest | Logs filter

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.LogsApi.ListLogs(ctx, body)
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

| Name     | Type                                      | Description                                                                 | Notes |
| -------- | ----------------------------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**  | **context.Context**                       | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **body** | [**LogsListRequest**](LogsListRequest.md) | Logs filter                                                                 |

### Optional Parameters

This endpoint does not have optional parameters.

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

- 200: OK
- 400: Bad request (likely an issue in the payload formatting)
- 403: Permission issue (likely using an invalid API Key)
- 413: Payload too large (batch is above 5MB uncompressed)
- 5xx: Internal error, request should be retried after some time

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
