# \LogsArchivesApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddReadRoleToArchive**](LogsArchivesApi.md#AddReadRoleToArchive) | **Post** /api/v2/logs/config/archives/{archive_id}/readers | Grant role to an archive
[**CreateLogsArchive**](LogsArchivesApi.md#CreateLogsArchive) | **Post** /api/v2/logs/config/archives | Create an archive
[**DeleteLogsArchive**](LogsArchivesApi.md#DeleteLogsArchive) | **Delete** /api/v2/logs/config/archives/{archive_id} | Delete an archive
[**GetLogsArchive**](LogsArchivesApi.md#GetLogsArchive) | **Get** /api/v2/logs/config/archives/{archive_id} | Get an archive
[**GetLogsArchiveOrder**](LogsArchivesApi.md#GetLogsArchiveOrder) | **Get** /api/v2/logs/config/archive-order | Get archive order
[**ListArchiveReadRoles**](LogsArchivesApi.md#ListArchiveReadRoles) | **Get** /api/v2/logs/config/archives/{archive_id}/readers | List read roles for an archive
[**ListLogsArchives**](LogsArchivesApi.md#ListLogsArchives) | **Get** /api/v2/logs/config/archives | Get all archives
[**RemoveRoleFromArchive**](LogsArchivesApi.md#RemoveRoleFromArchive) | **Delete** /api/v2/logs/config/archives/{archive_id}/readers | Revoke role from an archive
[**UpdateLogsArchive**](LogsArchivesApi.md#UpdateLogsArchive) | **Put** /api/v2/logs/config/archives/{archive_id} | Update an archive
[**UpdateLogsArchiveOrder**](LogsArchivesApi.md#UpdateLogsArchiveOrder) | **Put** /api/v2/logs/config/archive-order | Update archive order



## AddReadRoleToArchive

> AddReadRoleToArchive(ctx, archiveId).Body(body).Execute()

Grant role to an archive



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

    archiveId := "archiveId_example" // string | The ID of the archive.
    body := *datadog.NewRelationshipToRole() // RelationshipToRole | 

    configuration := datadog.NewConfiguration()
    configuration.SetUnstableOperationEnabled("AddReadRoleToArchive", true)

    api_client := datadog.NewAPIClient(configuration)
    r, err := api_client.LogsArchivesApi.AddReadRoleToArchive(ctx, archiveId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsArchivesApi.AddReadRoleToArchive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**archiveId** | **string** | The ID of the archive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddReadRoleToArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**RelationshipToRole**](RelationshipToRole.md) |  | 

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

> LogsArchive CreateLogsArchive(ctx).Body(body).Execute()

Create an archive



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

    body := *datadog.NewLogsArchiveCreateRequest() // LogsArchiveCreateRequest | The definition of the new archive.

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.LogsArchivesApi.CreateLogsArchive(ctx).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsArchivesApi.CreateLogsArchive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLogsArchive`: LogsArchive
    fmt.Fprintf(os.Stdout, "Response from `LogsArchivesApi.CreateLogsArchive`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLogsArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**LogsArchiveCreateRequest**](LogsArchiveCreateRequest.md) | The definition of the new archive. | 

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

> DeleteLogsArchive(ctx, archiveId).Execute()

Delete an archive



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

    archiveId := "archiveId_example" // string | The ID of the archive.

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    r, err := api_client.LogsArchivesApi.DeleteLogsArchive(ctx, archiveId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsArchivesApi.DeleteLogsArchive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**archiveId** | **string** | The ID of the archive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLogsArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> LogsArchive GetLogsArchive(ctx, archiveId).Execute()

Get an archive



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

    archiveId := "archiveId_example" // string | The ID of the archive.

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.LogsArchivesApi.GetLogsArchive(ctx, archiveId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsArchivesApi.GetLogsArchive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLogsArchive`: LogsArchive
    fmt.Fprintf(os.Stdout, "Response from `LogsArchivesApi.GetLogsArchive`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**archiveId** | **string** | The ID of the archive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogsArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> LogsArchiveOrder GetLogsArchiveOrder(ctx).Execute()

Get archive order



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


    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.LogsArchivesApi.GetLogsArchiveOrder(ctx).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsArchivesApi.GetLogsArchiveOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLogsArchiveOrder`: LogsArchiveOrder
    fmt.Fprintf(os.Stdout, "Response from `LogsArchivesApi.GetLogsArchiveOrder`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogsArchiveOrderRequest struct via the builder pattern


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

> RolesResponse ListArchiveReadRoles(ctx, archiveId).Execute()

List read roles for an archive



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

    archiveId := "archiveId_example" // string | The ID of the archive.

    configuration := datadog.NewConfiguration()
    configuration.SetUnstableOperationEnabled("ListArchiveReadRoles", true)

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.LogsArchivesApi.ListArchiveReadRoles(ctx, archiveId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsArchivesApi.ListArchiveReadRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListArchiveReadRoles`: RolesResponse
    fmt.Fprintf(os.Stdout, "Response from `LogsArchivesApi.ListArchiveReadRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**archiveId** | **string** | The ID of the archive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListArchiveReadRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> LogsArchives ListLogsArchives(ctx).Execute()

Get all archives



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


    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.LogsArchivesApi.ListLogsArchives(ctx).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsArchivesApi.ListLogsArchives``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLogsArchives`: LogsArchives
    fmt.Fprintf(os.Stdout, "Response from `LogsArchivesApi.ListLogsArchives`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListLogsArchivesRequest struct via the builder pattern


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

> RemoveRoleFromArchive(ctx, archiveId).Body(body).Execute()

Revoke role from an archive



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

    archiveId := "archiveId_example" // string | The ID of the archive.
    body := *datadog.NewRelationshipToRole() // RelationshipToRole | 

    configuration := datadog.NewConfiguration()
    configuration.SetUnstableOperationEnabled("RemoveRoleFromArchive", true)

    api_client := datadog.NewAPIClient(configuration)
    r, err := api_client.LogsArchivesApi.RemoveRoleFromArchive(ctx, archiveId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsArchivesApi.RemoveRoleFromArchive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**archiveId** | **string** | The ID of the archive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveRoleFromArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**RelationshipToRole**](RelationshipToRole.md) |  | 

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

> LogsArchive UpdateLogsArchive(ctx, archiveId).Body(body).Execute()

Update an archive



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

    archiveId := "archiveId_example" // string | The ID of the archive.
    body := *datadog.NewLogsArchiveCreateRequest() // LogsArchiveCreateRequest | New definition of the archive.

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.LogsArchivesApi.UpdateLogsArchive(ctx, archiveId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsArchivesApi.UpdateLogsArchive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLogsArchive`: LogsArchive
    fmt.Fprintf(os.Stdout, "Response from `LogsArchivesApi.UpdateLogsArchive`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**archiveId** | **string** | The ID of the archive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLogsArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**LogsArchiveCreateRequest**](LogsArchiveCreateRequest.md) | New definition of the archive. | 

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

> LogsArchiveOrder UpdateLogsArchiveOrder(ctx).Body(body).Execute()

Update archive order



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

    body := *datadog.NewLogsArchiveOrder() // LogsArchiveOrder | An object containing the new ordered list of archive IDs.

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.LogsArchivesApi.UpdateLogsArchiveOrder(ctx).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsArchivesApi.UpdateLogsArchiveOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLogsArchiveOrder`: LogsArchiveOrder
    fmt.Fprintf(os.Stdout, "Response from `LogsArchivesApi.UpdateLogsArchiveOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLogsArchiveOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**LogsArchiveOrder**](LogsArchiveOrder.md) | An object containing the new ordered list of archive IDs. | 

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

