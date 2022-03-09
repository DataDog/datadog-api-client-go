# RUMApi

All URIs are relative to *https://api.datadoghq.com*

| Method                                           | HTTP request                       | Description              |
| ------------------------------------------------ | ---------------------------------- | ------------------------ |
| [**ListRUMEvents**](RUMApi.md#ListRUMEvents)     | **Get** /api/v2/rum/events         | Get a list of RUM events |
| [**SearchRUMEvents**](RUMApi.md#SearchRUMEvents) | **Post** /api/v2/rum/events/search | Search RUM events        |

## ListRUMEvents

> RUMEventsResponse ListRUMEvents(ctx, datadog.ListRUMEventsOptionalParameters{})

List endpoint returns events that match a RUM search query.
[Results are paginated][1].

Use this endpoint to see your latest RUM events.

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

    filterQuery := "@type:session @application_id:xxxx" // string | Search query following RUM syntax. (optional)
    filterFrom := time.Now() // time.Time | Minimum timestamp for requested events. (optional)
    filterTo := time.Now() // time.Time | Maximum timestamp for requested events. (optional)
    sort := datadog.RUMSort("timestamp") // RUMSort | Order of events in results. (optional)
    pageCursor := "eyJzdGFydEF0IjoiQVFBQUFYS2tMS3pPbm40NGV3QUFBQUJCV0V0clRFdDZVbG8zY3pCRmNsbHJiVmxDWlEifQ==" // string | List following results with a cursor provided in the previous query. (optional)
    pageLimit := int32(25) // int32 | Maximum number of events in the response. (optional) (default to 10)
    optionalParams := datadog.ListRUMEventsOptionalParameters{
        FilterQuery: &filterQuery,
        FilterFrom: &filterFrom,
        FilterTo: &filterTo,
        Sort: &sort,
        PageCursor: &pageCursor,
        PageLimit: &pageLimit,
    }

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.RUMApi.ListRUMEvents(ctx, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMApi.ListRUMEvents`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRUMEvents`: RUMEventsResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from RUMApi.ListRUMEvents:\n%s\n", responseContent)
}
```

### Required Parameters

### Optional Parameters

Other parameters are passed through a pointer to a ListRUMEventsOptionalParameters struct.

| Name            | Type                      | Description                                                          | Notes           |
| --------------- | ------------------------- | -------------------------------------------------------------------- | --------------- |
| **filterQuery** | **string**                | Search query following RUM syntax.                                   |
| **filterFrom**  | **time.Time**             | Minimum timestamp for requested events.                              |
| **filterTo**    | **time.Time**             | Maximum timestamp for requested events.                              |
| **sort**        | [**RUMSort**](RUMSort.md) | Order of events in results.                                          |
| **pageCursor**  | **string**                | List following results with a cursor provided in the previous query. |
| **pageLimit**   | **int32**                 | Maximum number of events in the response.                            | [default to 10] |

### Return type

[**RUMEventsResponse**](RUMEventsResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SearchRUMEvents

> RUMEventsResponse SearchRUMEvents(ctx, datadog.SearchRUMEventsOptionalParameters{})

List endpoint returns RUM events that match a RUM search query.
[Results are paginated][1].

Use this endpoint to build complex RUM events filtering and search.

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

    body := *datadog.NewRUMSearchEventsRequest() // RUMSearchEventsRequest |  (optional)
    optionalParams := datadog.SearchRUMEventsOptionalParameters{
        Body: &body,
    }

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.RUMApi.SearchRUMEvents(ctx, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMApi.SearchRUMEvents`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchRUMEvents`: RUMEventsResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from RUMApi.SearchRUMEvents:\n%s\n", responseContent)
}
```

### Required Parameters

### Optional Parameters

Other parameters are passed through a pointer to a SearchRUMEventsOptionalParameters struct.

| Name     | Type                                                    | Description | Notes |
| -------- | ------------------------------------------------------- | ----------- | ----- |
| **body** | [**RUMSearchEventsRequest**](RUMSearchEventsRequest.md) |             |

### Return type

[**RUMEventsResponse**](RUMEventsResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
