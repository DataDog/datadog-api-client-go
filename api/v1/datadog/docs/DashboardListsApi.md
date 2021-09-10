# DashboardListsApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------ | ------------ | ------------
[**CreateDashboardList**](DashboardListsApi.md#CreateDashboardList) | **Post** /api/v1/dashboard/lists/manual | Create a dashboard list
[**DeleteDashboardList**](DashboardListsApi.md#DeleteDashboardList) | **Delete** /api/v1/dashboard/lists/manual/{list_id} | Delete a dashboard list
[**GetDashboardList**](DashboardListsApi.md#GetDashboardList) | **Get** /api/v1/dashboard/lists/manual/{list_id} | Get a dashboard list
[**ListDashboardLists**](DashboardListsApi.md#ListDashboardLists) | **Get** /api/v1/dashboard/lists/manual | Get all dashboard lists
[**UpdateDashboardList**](DashboardListsApi.md#UpdateDashboardList) | **Put** /api/v1/dashboard/lists/manual/{list_id} | Update a dashboard list



## CreateDashboardList

> DashboardList CreateDashboardList(ctx, body)

Create an empty dashboard list.

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

    body := *datadog.NewDashboardList("My Dashboard") // DashboardList | Create a dashboard list request body.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.DashboardListsApi.CreateDashboardList(ctx, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardListsApi.CreateDashboardList`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDashboardList`: DashboardList
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from DashboardListsApi.CreateDashboardList:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**body** | [**DashboardList**](DashboardList.md) | Create a dashboard list request body. | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**DashboardList**](DashboardList.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDashboardList

> DashboardListDeleteResponse DeleteDashboardList(ctx, listId)

Delete a dashboard list.

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

    listId := int64(789) // int64 | ID of the dashboard list to delete.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.DashboardListsApi.DeleteDashboardList(ctx, listId)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardListsApi.DeleteDashboardList`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDashboardList`: DashboardListDeleteResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from DashboardListsApi.DeleteDashboardList:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**listId** | **int64** | ID of the dashboard list to delete. | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**DashboardListDeleteResponse**](DashboardListDeleteResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDashboardList

> DashboardList GetDashboardList(ctx, listId)

Fetch an existing dashboard list's definition.

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

    listId := int64(789) // int64 | ID of the dashboard list to fetch.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.DashboardListsApi.GetDashboardList(ctx, listId)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardListsApi.GetDashboardList`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDashboardList`: DashboardList
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from DashboardListsApi.GetDashboardList:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**listId** | **int64** | ID of the dashboard list to fetch. | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**DashboardList**](DashboardList.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDashboardLists

> DashboardListListResponse ListDashboardLists(ctx)

Fetch all of your existing dashboard list definitions.

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
    resp, r, err := apiClient.DashboardListsApi.ListDashboardLists(ctx)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardListsApi.ListDashboardLists`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDashboardLists`: DashboardListListResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from DashboardListsApi.ListDashboardLists:\n%s\n", responseContent)
}
```

### Required Parameters

This endpoint does not need any parameter.


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**DashboardListListResponse**](DashboardListListResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDashboardList

> DashboardList UpdateDashboardList(ctx, listId, body)

Update the name of a dashboard list.

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

    listId := int64(789) // int64 | ID of the dashboard list to update.
    body := *datadog.NewDashboardList("My Dashboard") // DashboardList | Update a dashboard list request body.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.DashboardListsApi.UpdateDashboardList(ctx, listId, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardListsApi.UpdateDashboardList`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDashboardList`: DashboardList
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from DashboardListsApi.UpdateDashboardList:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**listId** | **int64** | ID of the dashboard list to update. |  |
**body** | [**DashboardList**](DashboardList.md) | Update a dashboard list request body. | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**DashboardList**](DashboardList.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

