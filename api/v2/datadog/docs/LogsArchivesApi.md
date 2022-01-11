# LogsArchivesApi

All URIs are relative to *https://api.datadoghq.com*

| Method                                                                  | HTTP request                                                 | Description                    |
| ----------------------------------------------------------------------- | ------------------------------------------------------------ | ------------------------------ |
| [**AddReadRoleToArchive**](LogsArchivesApi.md#AddReadRoleToArchive)     | **Post** /api/v2/logs/config/archives/{archive_id}/readers   | Grant role to an archive       |
| [**CreateLogsArchive**](LogsArchivesApi.md#CreateLogsArchive)           | **Post** /api/v2/logs/config/archives                        | Create an archive              |
| [**DeleteLogsArchive**](LogsArchivesApi.md#DeleteLogsArchive)           | **Delete** /api/v2/logs/config/archives/{archive_id}         | Delete an archive              |
| [**GetLogsArchive**](LogsArchivesApi.md#GetLogsArchive)                 | **Get** /api/v2/logs/config/archives/{archive_id}            | Get an archive                 |
| [**GetLogsArchiveOrder**](LogsArchivesApi.md#GetLogsArchiveOrder)       | **Get** /api/v2/logs/config/archive-order                    | Get archive order              |
| [**ListArchiveReadRoles**](LogsArchivesApi.md#ListArchiveReadRoles)     | **Get** /api/v2/logs/config/archives/{archive_id}/readers    | List read roles for an archive |
| [**ListLogsArchives**](LogsArchivesApi.md#ListLogsArchives)             | **Get** /api/v2/logs/config/archives                         | Get all archives               |
| [**RemoveRoleFromArchive**](LogsArchivesApi.md#RemoveRoleFromArchive)   | **Delete** /api/v2/logs/config/archives/{archive_id}/readers | Revoke role from an archive    |
| [**UpdateLogsArchive**](LogsArchivesApi.md#UpdateLogsArchive)           | **Put** /api/v2/logs/config/archives/{archive_id}            | Update an archive              |
| [**UpdateLogsArchiveOrder**](LogsArchivesApi.md#UpdateLogsArchiveOrder) | **Put** /api/v2/logs/config/archive-order                    | Update archive order           |

## AddReadRoleToArchive

> AddReadRoleToArchive(ctx, archiveId, body)

Adds a read role to an archive. ([Roles API](https://docs.datadoghq.com/api/v2/roles/))

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
    ctx := datadog.NewDefaultContext(context.Background())

    archiveId := "archiveId_example" // string | The ID of the archive.
    body := *datadog.NewRelationshipToRole() // RelationshipToRole |

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    r, err := apiClient.LogsArchivesApi.AddReadRoleToArchive(ctx, archiveId, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsArchivesApi.AddReadRoleToArchive`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Required Parameters

| Name          | Type                                            | Description                                                                 | Notes |
| ------------- | ----------------------------------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**       | **context.Context**                             | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **archiveId** | **string**                                      | The ID of the archive.                                                      |       |
| **body**      | [**RelationshipToRole**](RelationshipToRole.md) |                                                                             |

### Optional Parameters

This endpoint does not have optional parameters.

### Return type

(empty response body)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CreateLogsArchive

> LogsArchive CreateLogsArchive(ctx, body)

Create an archive in your organization.

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

    body := *datadog.NewLogsArchiveCreateRequest() // LogsArchiveCreateRequest | The definition of the new archive.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.LogsArchivesApi.CreateLogsArchive(ctx, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsArchivesApi.CreateLogsArchive`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLogsArchive`: LogsArchive
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from LogsArchivesApi.CreateLogsArchive:\n%s\n", responseContent)
}
```

### Required Parameters

| Name     | Type                                                        | Description                                                                 | Notes |
| -------- | ----------------------------------------------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**  | **context.Context**                                         | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **body** | [**LogsArchiveCreateRequest**](LogsArchiveCreateRequest.md) | The definition of the new archive.                                          |

### Optional Parameters

This endpoint does not have optional parameters.

### Return type

[**LogsArchive**](LogsArchive.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## DeleteLogsArchive

> DeleteLogsArchive(ctx, archiveId)

Delete a given archive from your organization.

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
    ctx := datadog.NewDefaultContext(context.Background())

    archiveId := "archiveId_example" // string | The ID of the archive.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    r, err := apiClient.LogsArchivesApi.DeleteLogsArchive(ctx, archiveId)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsArchivesApi.DeleteLogsArchive`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Required Parameters

| Name          | Type                | Description                                                                 | Notes |
| ------------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**       | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **archiveId** | **string**          | The ID of the archive.                                                      |

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

## GetLogsArchive

> LogsArchive GetLogsArchive(ctx, archiveId)

Get a specific archive from your organization.

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

    archiveId := "archiveId_example" // string | The ID of the archive.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.LogsArchivesApi.GetLogsArchive(ctx, archiveId)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsArchivesApi.GetLogsArchive`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLogsArchive`: LogsArchive
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from LogsArchivesApi.GetLogsArchive:\n%s\n", responseContent)
}
```

### Required Parameters

| Name          | Type                | Description                                                                 | Notes |
| ------------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**       | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **archiveId** | **string**          | The ID of the archive.                                                      |

### Optional Parameters

This endpoint does not have optional parameters.

### Return type

[**LogsArchive**](LogsArchive.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetLogsArchiveOrder

> LogsArchiveOrder GetLogsArchiveOrder(ctx)

Get the current order of your archives.
This endpoint takes no JSON arguments.

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


    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.LogsArchivesApi.GetLogsArchiveOrder(ctx)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsArchivesApi.GetLogsArchiveOrder`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLogsArchiveOrder`: LogsArchiveOrder
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from LogsArchivesApi.GetLogsArchiveOrder:\n%s\n", responseContent)
}
```

### Required Parameters

This endpoint does not need any parameter.

### Optional Parameters

This endpoint does not have optional parameters.

### Return type

[**LogsArchiveOrder**](LogsArchiveOrder.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListArchiveReadRoles

> RolesResponse ListArchiveReadRoles(ctx, archiveId)

Returns all read roles a given archive is restricted to.

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

    archiveId := "archiveId_example" // string | The ID of the archive.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.LogsArchivesApi.ListArchiveReadRoles(ctx, archiveId)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsArchivesApi.ListArchiveReadRoles`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListArchiveReadRoles`: RolesResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from LogsArchivesApi.ListArchiveReadRoles:\n%s\n", responseContent)
}
```

### Required Parameters

| Name          | Type                | Description                                                                 | Notes |
| ------------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**       | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **archiveId** | **string**          | The ID of the archive.                                                      |

### Optional Parameters

This endpoint does not have optional parameters.

### Return type

[**RolesResponse**](RolesResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListLogsArchives

> LogsArchives ListLogsArchives(ctx)

Get the list of configured logs archives with their definitions.

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


    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.LogsArchivesApi.ListLogsArchives(ctx)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsArchivesApi.ListLogsArchives`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLogsArchives`: LogsArchives
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from LogsArchivesApi.ListLogsArchives:\n%s\n", responseContent)
}
```

### Required Parameters

This endpoint does not need any parameter.

### Optional Parameters

This endpoint does not have optional parameters.

### Return type

[**LogsArchives**](LogsArchives.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## RemoveRoleFromArchive

> RemoveRoleFromArchive(ctx, archiveId, body)

Removes a role from an archive. ([Roles API](https://docs.datadoghq.com/api/v2/roles/))

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
    ctx := datadog.NewDefaultContext(context.Background())

    archiveId := "archiveId_example" // string | The ID of the archive.
    body := *datadog.NewRelationshipToRole() // RelationshipToRole |

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    r, err := apiClient.LogsArchivesApi.RemoveRoleFromArchive(ctx, archiveId, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsArchivesApi.RemoveRoleFromArchive`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Required Parameters

| Name          | Type                                            | Description                                                                 | Notes |
| ------------- | ----------------------------------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**       | **context.Context**                             | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **archiveId** | **string**                                      | The ID of the archive.                                                      |       |
| **body**      | [**RelationshipToRole**](RelationshipToRole.md) |                                                                             |

### Optional Parameters

This endpoint does not have optional parameters.

### Return type

(empty response body)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateLogsArchive

> LogsArchive UpdateLogsArchive(ctx, archiveId, body)

Update a given archive configuration.

**Note**: Using this method updates your archive configuration by **replacing**
your current configuration with the new one sent to your Datadog organization.

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

    archiveId := "archiveId_example" // string | The ID of the archive.
    body := *datadog.NewLogsArchiveCreateRequest() // LogsArchiveCreateRequest | New definition of the archive.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.LogsArchivesApi.UpdateLogsArchive(ctx, archiveId, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsArchivesApi.UpdateLogsArchive`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLogsArchive`: LogsArchive
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from LogsArchivesApi.UpdateLogsArchive:\n%s\n", responseContent)
}
```

### Required Parameters

| Name          | Type                                                        | Description                                                                 | Notes |
| ------------- | ----------------------------------------------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**       | **context.Context**                                         | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **archiveId** | **string**                                                  | The ID of the archive.                                                      |       |
| **body**      | [**LogsArchiveCreateRequest**](LogsArchiveCreateRequest.md) | New definition of the archive.                                              |

### Optional Parameters

This endpoint does not have optional parameters.

### Return type

[**LogsArchive**](LogsArchive.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateLogsArchiveOrder

> LogsArchiveOrder UpdateLogsArchiveOrder(ctx, body)

Update the order of your archives. Since logs are processed sequentially, reordering an archive may change
the structure and content of the data processed by other archives.

**Note**: Using the `PUT` method updates your archive's order by replacing the current order
with the new one.

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

    body := *datadog.NewLogsArchiveOrder() // LogsArchiveOrder | An object containing the new ordered list of archive IDs.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.LogsArchivesApi.UpdateLogsArchiveOrder(ctx, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsArchivesApi.UpdateLogsArchiveOrder`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLogsArchiveOrder`: LogsArchiveOrder
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from LogsArchivesApi.UpdateLogsArchiveOrder:\n%s\n", responseContent)
}
```

### Required Parameters

| Name     | Type                                        | Description                                                                 | Notes |
| -------- | ------------------------------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**  | **context.Context**                         | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **body** | [**LogsArchiveOrder**](LogsArchiveOrder.md) | An object containing the new ordered list of archive IDs.                   |

### Optional Parameters

This endpoint does not have optional parameters.

### Return type

[**LogsArchiveOrder**](LogsArchiveOrder.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
