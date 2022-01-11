# LogsIndexesApi

All URIs are relative to *https://api.datadoghq.com*

| Method                                                             | HTTP request                               | Description          |
| ------------------------------------------------------------------ | ------------------------------------------ | -------------------- |
| [**CreateLogsIndex**](LogsIndexesApi.md#CreateLogsIndex)           | **Post** /api/v1/logs/config/indexes       | Create an index      |
| [**GetLogsIndex**](LogsIndexesApi.md#GetLogsIndex)                 | **Get** /api/v1/logs/config/indexes/{name} | Get an index         |
| [**GetLogsIndexOrder**](LogsIndexesApi.md#GetLogsIndexOrder)       | **Get** /api/v1/logs/config/index-order    | Get indexes order    |
| [**ListLogIndexes**](LogsIndexesApi.md#ListLogIndexes)             | **Get** /api/v1/logs/config/indexes        | Get all indexes      |
| [**UpdateLogsIndex**](LogsIndexesApi.md#UpdateLogsIndex)           | **Put** /api/v1/logs/config/indexes/{name} | Update an index      |
| [**UpdateLogsIndexOrder**](LogsIndexesApi.md#UpdateLogsIndexOrder) | **Put** /api/v1/logs/config/index-order    | Update indexes order |

## CreateLogsIndex

> LogsIndex CreateLogsIndex(ctx, body)

Creates a new index. Returns the Index object passed in the request body when the request is successful.

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

    body := *datadog.NewLogsIndex(*datadog.NewLogsFilter(), "main") // LogsIndex | Object containing the new index.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.LogsIndexesApi.CreateLogsIndex(ctx, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsIndexesApi.CreateLogsIndex`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLogsIndex`: LogsIndex
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from LogsIndexesApi.CreateLogsIndex:\n%s\n", responseContent)
}
```

### Required Parameters

| Name     | Type                          | Description                                                                 | Notes |
| -------- | ----------------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**  | **context.Context**           | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **body** | [**LogsIndex**](LogsIndex.md) | Object containing the new index.                                            |

### Optional Parameters

This endpoint does not have optional parameters.

### Return type

[**LogsIndex**](LogsIndex.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetLogsIndex

> LogsIndex GetLogsIndex(ctx, name)

Get one log index from your organization. This endpoint takes no JSON arguments.

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

    name := "name_example" // string | Name of the log index.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.LogsIndexesApi.GetLogsIndex(ctx, name)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsIndexesApi.GetLogsIndex`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLogsIndex`: LogsIndex
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from LogsIndexesApi.GetLogsIndex:\n%s\n", responseContent)
}
```

### Required Parameters

| Name     | Type                | Description                                                                 | Notes |
| -------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**  | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **name** | **string**          | Name of the log index.                                                      |

### Optional Parameters

This endpoint does not have optional parameters.

### Return type

[**LogsIndex**](LogsIndex.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetLogsIndexOrder

> LogsIndexesOrder GetLogsIndexOrder(ctx)

Get the current order of your log indexes. This endpoint takes no JSON arguments.

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


    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.LogsIndexesApi.GetLogsIndexOrder(ctx)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsIndexesApi.GetLogsIndexOrder`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLogsIndexOrder`: LogsIndexesOrder
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from LogsIndexesApi.GetLogsIndexOrder:\n%s\n", responseContent)
}
```

### Required Parameters

This endpoint does not need any parameter.

### Optional Parameters

This endpoint does not have optional parameters.

### Return type

[**LogsIndexesOrder**](LogsIndexesOrder.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListLogIndexes

> LogsIndexListResponse ListLogIndexes(ctx)

The Index object describes the configuration of a log index.
This endpoint returns an array of the `LogIndex` objects of your organization.

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


    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.LogsIndexesApi.ListLogIndexes(ctx)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsIndexesApi.ListLogIndexes`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLogIndexes`: LogsIndexListResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from LogsIndexesApi.ListLogIndexes:\n%s\n", responseContent)
}
```

### Required Parameters

This endpoint does not need any parameter.

### Optional Parameters

This endpoint does not have optional parameters.

### Return type

[**LogsIndexListResponse**](LogsIndexListResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateLogsIndex

> LogsIndex UpdateLogsIndex(ctx, name, body)

Update an index as identified by its name.
Returns the Index object passed in the request body when the request is successful.

Using the `PUT` method updates your index’s configuration by **replacing**
your current configuration with the new one sent to your Datadog organization.

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

    name := "name_example" // string | Name of the log index.
    body := *datadog.NewLogsIndexUpdateRequest(*datadog.NewLogsFilter()) // LogsIndexUpdateRequest | Object containing the new `LogsIndexUpdateRequest`.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.LogsIndexesApi.UpdateLogsIndex(ctx, name, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsIndexesApi.UpdateLogsIndex`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLogsIndex`: LogsIndex
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from LogsIndexesApi.UpdateLogsIndex:\n%s\n", responseContent)
}
```

### Required Parameters

| Name     | Type                                                    | Description                                                                 | Notes |
| -------- | ------------------------------------------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**  | **context.Context**                                     | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **name** | **string**                                              | Name of the log index.                                                      |       |
| **body** | [**LogsIndexUpdateRequest**](LogsIndexUpdateRequest.md) | Object containing the new &#x60;LogsIndexUpdateRequest&#x60;.               |

### Optional Parameters

This endpoint does not have optional parameters.

### Return type

[**LogsIndex**](LogsIndex.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateLogsIndexOrder

> LogsIndexesOrder UpdateLogsIndexOrder(ctx, body)

This endpoint updates the index order of your organization.
It returns the index order object passed in the request body when the request is successful.

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

    body := *datadog.NewLogsIndexesOrder([]string{"IndexNames_example"}) // LogsIndexesOrder | Object containing the new ordered list of index names

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.LogsIndexesApi.UpdateLogsIndexOrder(ctx, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsIndexesApi.UpdateLogsIndexOrder`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLogsIndexOrder`: LogsIndexesOrder
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from LogsIndexesApi.UpdateLogsIndexOrder:\n%s\n", responseContent)
}
```

### Required Parameters

| Name     | Type                                        | Description                                                                 | Notes |
| -------- | ------------------------------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**  | **context.Context**                         | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **body** | [**LogsIndexesOrder**](LogsIndexesOrder.md) | Object containing the new ordered list of index names                       |

### Optional Parameters

This endpoint does not have optional parameters.

### Return type

[**LogsIndexesOrder**](LogsIndexesOrder.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
