# \DashboardListsApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDashboardList**](DashboardListsApi.md#CreateDashboardList) | **Post** /api/v1/dashboard/lists/manual | Create a dashboard list
[**DeleteDashboardList**](DashboardListsApi.md#DeleteDashboardList) | **Delete** /api/v1/dashboard/lists/manual/{list_id} | Delete a dashboard list
[**GetDashboardList**](DashboardListsApi.md#GetDashboardList) | **Get** /api/v1/dashboard/lists/manual/{list_id} | Get a dashboard list
[**ListDashboardLists**](DashboardListsApi.md#ListDashboardLists) | **Get** /api/v1/dashboard/lists/manual | Get all dashboard lists
[**UpdateDashboardList**](DashboardListsApi.md#UpdateDashboardList) | **Put** /api/v1/dashboard/lists/manual/{list_id} | Update a dashboard list



## CreateDashboardList

> DashboardList CreateDashboardList(ctx).Body(body).Execute()

Create a dashboard list



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

    body := *datadog.NewDashboardList("My Dashboard") // DashboardList | Create a dashboard list request body.

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.DashboardListsApi.CreateDashboardList(ctx).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardListsApi.CreateDashboardList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDashboardList`: DashboardList
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from DashboardListsApi.CreateDashboardList:\n%v\n", response_content)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDashboardListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DashboardList**](DashboardList.md) | Create a dashboard list request body. | 

### Return type

[**DashboardList**](DashboardList.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDashboardList

> DashboardListDeleteResponse DeleteDashboardList(ctx, listId).Execute()

Delete a dashboard list



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

    listId := int64(789) // int64 | ID of the dashboard list to delete.

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.DashboardListsApi.DeleteDashboardList(ctx, listId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardListsApi.DeleteDashboardList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDashboardList`: DashboardListDeleteResponse
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from DashboardListsApi.DeleteDashboardList:\n%v\n", response_content)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **int64** | ID of the dashboard list to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDashboardListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DashboardListDeleteResponse**](DashboardListDeleteResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDashboardList

> DashboardList GetDashboardList(ctx, listId).Execute()

Get a dashboard list



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

    listId := int64(789) // int64 | ID of the dashboard list to fetch.

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.DashboardListsApi.GetDashboardList(ctx, listId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardListsApi.GetDashboardList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDashboardList`: DashboardList
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from DashboardListsApi.GetDashboardList:\n%v\n", response_content)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **int64** | ID of the dashboard list to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDashboardListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DashboardList**](DashboardList.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDashboardLists

> DashboardListListResponse ListDashboardLists(ctx).Execute()

Get all dashboard lists



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
    resp, r, err := api_client.DashboardListsApi.ListDashboardLists(ctx).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardListsApi.ListDashboardLists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDashboardLists`: DashboardListListResponse
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from DashboardListsApi.ListDashboardLists:\n%v\n", response_content)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListDashboardListsRequest struct via the builder pattern


### Return type

[**DashboardListListResponse**](DashboardListListResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDashboardList

> DashboardList UpdateDashboardList(ctx, listId).Body(body).Execute()

Update a dashboard list



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

    listId := int64(789) // int64 | ID of the dashboard list to update.
    body := *datadog.NewDashboardList("My Dashboard") // DashboardList | Update a dashboard list request body.

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.DashboardListsApi.UpdateDashboardList(ctx, listId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardListsApi.UpdateDashboardList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDashboardList`: DashboardList
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from DashboardListsApi.UpdateDashboardList:\n%v\n", response_content)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **int64** | ID of the dashboard list to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDashboardListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DashboardList**](DashboardList.md) | Update a dashboard list request body. | 

### Return type

[**DashboardList**](DashboardList.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

