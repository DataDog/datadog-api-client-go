# NotebooksApi

All URIs are relative to *https://api.datadoghq.com*

| Method                                               | HTTP request                               | Description       |
| ---------------------------------------------------- | ------------------------------------------ | ----------------- |
| [**CreateNotebook**](NotebooksApi.md#CreateNotebook) | **Post** /api/v1/notebooks                 | Create a notebook |
| [**DeleteNotebook**](NotebooksApi.md#DeleteNotebook) | **Delete** /api/v1/notebooks/{notebook_id} | Delete a notebook |
| [**GetNotebook**](NotebooksApi.md#GetNotebook)       | **Get** /api/v1/notebooks/{notebook_id}    | Get a notebook    |
| [**ListNotebooks**](NotebooksApi.md#ListNotebooks)   | **Get** /api/v1/notebooks                  | Get all notebooks |
| [**UpdateNotebook**](NotebooksApi.md#UpdateNotebook) | **Put** /api/v1/notebooks/{notebook_id}    | Update a notebook |

## CreateNotebook

> NotebookResponse CreateNotebook(ctx, body)

Create a notebook using the specified options.

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

    body := *datadog.NewNotebookCreateRequest(*datadog.NewNotebookCreateData(*datadog.NewNotebookCreateDataAttributes([]datadog.NotebookCellCreateRequest{*datadog.NewNotebookCellCreateRequest(datadog.NotebookCellCreateRequestAttributes{NotebookDistributionCellAttributes: datadog.NewNotebookDistributionCellAttributes(*datadog.NewDistributionWidgetDefinition([]datadog.DistributionWidgetRequest{*datadog.NewDistributionWidgetRequest()}, datadog.DistributionWidgetDefinitionType("distribution")))}, datadog.NotebookCellResourceType("notebook_cells"))}, "Example Notebook", datadog.NotebookGlobalTime{NotebookAbsoluteTime: datadog.NewNotebookAbsoluteTime(time.Now(), time.Now())}), datadog.NotebookResourceType("notebooks"))) // NotebookCreateRequest | The JSON description of the notebook you want to create.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.NotebooksApi.CreateNotebook(ctx, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotebooksApi.CreateNotebook`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNotebook`: NotebookResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from NotebooksApi.CreateNotebook:\n%s\n", responseContent)
}
```

### Required Parameters

| Name     | Type                                                  | Description                                                                 | Notes |
| -------- | ----------------------------------------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**  | **context.Context**                                   | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **body** | [**NotebookCreateRequest**](NotebookCreateRequest.md) | The JSON description of the notebook you want to create.                    |

### Optional Parameters

This endpoint does not have optional parameters.

### Return type

[**NotebookResponse**](NotebookResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## DeleteNotebook

> DeleteNotebook(ctx, notebookId)

Delete a notebook using the specified ID.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
    ctx := datadog.NewDefaultContext(context.Background())

    notebookId := int64(789) // int64 | Unique ID, assigned when you create the notebook.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    r, err := apiClient.NotebooksApi.DeleteNotebook(ctx, notebookId)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotebooksApi.DeleteNotebook`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Required Parameters

| Name           | Type                | Description                                                                 | Notes |
| -------------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**        | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **notebookId** | **int64**           | Unique ID, assigned when you create the notebook.                           |

### Optional Parameters

This endpoint does not have optional parameters.

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

## GetNotebook

> NotebookResponse GetNotebook(ctx, notebookId)

Get a notebook using the specified notebook ID.

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

    notebookId := int64(789) // int64 | Unique ID, assigned when you create the notebook.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.NotebooksApi.GetNotebook(ctx, notebookId)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotebooksApi.GetNotebook`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNotebook`: NotebookResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from NotebooksApi.GetNotebook:\n%s\n", responseContent)
}
```

### Required Parameters

| Name           | Type                | Description                                                                 | Notes |
| -------------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**        | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **notebookId** | **int64**           | Unique ID, assigned when you create the notebook.                           |

### Optional Parameters

This endpoint does not have optional parameters.

### Return type

[**NotebookResponse**](NotebookResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListNotebooks

> NotebooksResponse ListNotebooks(ctx, datadog.ListNotebooksOptionalParameters{})

Get all notebooks. This can also be used to search for notebooks with a particular `query` in the notebook
`name` or author `handle`.

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

    authorHandle := "test@datadoghq.com" // string | Return notebooks created by the given `author_handle`. (optional)
    excludeAuthorHandle := "test@datadoghq.com" // string | Return notebooks not created by the given `author_handle`. (optional)
    start := int64(0) // int64 | The index of the first notebook you want returned. (optional)
    count := int64(5) // int64 | The number of notebooks to be returned. (optional)
    sortField := "modified" // string | Sort by field `modified`, `name`, or `created`. (optional) (default to "modified")
    sortDir := "desc" // string | Sort by direction `asc` or `desc`. (optional) (default to "desc")
    query := "postmortem" // string | Return only notebooks with `query` string in notebook name or author handle. (optional)
    includeCells := false // bool | Value of `false` excludes the `cells` and global `time` for each notebook. (optional) (default to true)
    isTemplate := false // bool | True value returns only template notebooks. Default is false (returns only non-template notebooks). (optional) (default to false)
    type_ := "investigation" // string | If type is provided, returns only notebooks with that metadata type. Default does not have type filtering. (optional)
    optionalParams := datadog.ListNotebooksOptionalParameters{
        AuthorHandle: &authorHandle,
        ExcludeAuthorHandle: &excludeAuthorHandle,
        Start: &start,
        Count: &count,
        SortField: &sortField,
        SortDir: &sortDir,
        Query: &query,
        IncludeCells: &includeCells,
        IsTemplate: &isTemplate,
        Type_: &type_,
    }

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.NotebooksApi.ListNotebooks(ctx, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotebooksApi.ListNotebooks`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNotebooks`: NotebooksResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from NotebooksApi.ListNotebooks:\n%s\n", responseContent)
}
```

### Required Parameters

### Optional Parameters

Other parameters are passed through a pointer to a ListNotebooksOptionalParameters struct.

| Name                    | Type       | Description                                                                                                | Notes                             |
| ----------------------- | ---------- | ---------------------------------------------------------------------------------------------------------- | --------------------------------- |
| **authorHandle**        | **string** | Return notebooks created by the given &#x60;author_handle&#x60;.                                           |
| **excludeAuthorHandle** | **string** | Return notebooks not created by the given &#x60;author_handle&#x60;.                                       |
| **start**               | **int64**  | The index of the first notebook you want returned.                                                         |
| **count**               | **int64**  | The number of notebooks to be returned.                                                                    |
| **sortField**           | **string** | Sort by field &#x60;modified&#x60;, &#x60;name&#x60;, or &#x60;created&#x60;.                              | [default to &quot;modified&quot;] |
| **sortDir**             | **string** | Sort by direction &#x60;asc&#x60; or &#x60;desc&#x60;.                                                     | [default to &quot;desc&quot;]     |
| **query**               | **string** | Return only notebooks with &#x60;query&#x60; string in notebook name or author handle.                     |
| **includeCells**        | **bool**   | Value of &#x60;false&#x60; excludes the &#x60;cells&#x60; and global &#x60;time&#x60; for each notebook.   | [default to true]                 |
| **isTemplate**          | **bool**   | True value returns only template notebooks. Default is false (returns only non-template notebooks).        | [default to false]                |
| **type\_**              | **string** | If type is provided, returns only notebooks with that metadata type. Default does not have type filtering. |

### Return type

[**NotebooksResponse**](NotebooksResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateNotebook

> NotebookResponse UpdateNotebook(ctx, notebookId, body)

Update a notebook using the specified ID.

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

    notebookId := int64(789) // int64 | Unique ID, assigned when you create the notebook.
    body := *datadog.NewNotebookUpdateRequest(*datadog.NewNotebookUpdateData(*datadog.NewNotebookUpdateDataAttributes([]datadog.NotebookUpdateCell{datadog.NotebookUpdateCell{NotebookCellCreateRequest: datadog.NewNotebookCellCreateRequest(datadog.NotebookCellCreateRequestAttributes{NotebookDistributionCellAttributes: datadog.NewNotebookDistributionCellAttributes(*datadog.NewDistributionWidgetDefinition([]datadog.DistributionWidgetRequest{*datadog.NewDistributionWidgetRequest()}, datadog.DistributionWidgetDefinitionType("distribution")))}, datadog.NotebookCellResourceType("notebook_cells"))}}, "Example Notebook", datadog.NotebookGlobalTime{NotebookAbsoluteTime: datadog.NewNotebookAbsoluteTime(time.Now(), time.Now())}), datadog.NotebookResourceType("notebooks"))) // NotebookUpdateRequest | Update notebook request body.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.NotebooksApi.UpdateNotebook(ctx, notebookId, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotebooksApi.UpdateNotebook`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNotebook`: NotebookResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from NotebooksApi.UpdateNotebook:\n%s\n", responseContent)
}
```

### Required Parameters

| Name           | Type                                                  | Description                                                                 | Notes |
| -------------- | ----------------------------------------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**        | **context.Context**                                   | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **notebookId** | **int64**                                             | Unique ID, assigned when you create the notebook.                           |       |
| **body**       | [**NotebookUpdateRequest**](NotebookUpdateRequest.md) | Update notebook request body.                                               |

### Optional Parameters

This endpoint does not have optional parameters.

### Return type

[**NotebookResponse**](NotebookResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
