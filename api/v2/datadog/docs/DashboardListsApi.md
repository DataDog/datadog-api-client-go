# \DashboardListsApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDashboardListItems**](DashboardListsApi.md#CreateDashboardListItems) | **Post** /api/v2/dashboard/lists/manual/{dashboard_list_id}/dashboards | Add Items to a Dashboard List
[**DeleteDashboardListItems**](DashboardListsApi.md#DeleteDashboardListItems) | **Delete** /api/v2/dashboard/lists/manual/{dashboard_list_id}/dashboards | Delete items from a dashboard list
[**GetDashboardListItems**](DashboardListsApi.md#GetDashboardListItems) | **Get** /api/v2/dashboard/lists/manual/{dashboard_list_id}/dashboards | Get a Dashboard List
[**UpdateDashboardListItems**](DashboardListsApi.md#UpdateDashboardListItems) | **Put** /api/v2/dashboard/lists/manual/{dashboard_list_id}/dashboards | Update items of a dashboard list



## CreateDashboardListItems

> DashboardListAddItemsResponse CreateDashboardListItems(ctx, dashboardListId).Body(body).Execute()

Add Items to a Dashboard List



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

    dashboardListId := 987 // int64 | ID of the dashboard list to add items to.
    body := *datadog.NewDashboardListAddItemsRequest() // DashboardListAddItemsRequest | Dashboards to add to the dashboard list.

    configuration := datadog.NewConfiguration()
    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.DashboardListsApi.CreateDashboardListItems(context.Background(), dashboardListId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardListsApi.CreateDashboardListItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDashboardListItems`: DashboardListAddItemsResponse
    fmt.Fprintf(os.Stdout, "Response from `DashboardListsApi.CreateDashboardListItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardListId** | **int64** | ID of the dashboard list to add items to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDashboardListItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DashboardListAddItemsRequest**](DashboardListAddItemsRequest.md) | Dashboards to add to the dashboard list. | 

### Return type

[**DashboardListAddItemsResponse**](DashboardListAddItemsResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDashboardListItems

> DashboardListDeleteItemsResponse DeleteDashboardListItems(ctx, dashboardListId).Body(body).Execute()

Delete items from a dashboard list



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

    dashboardListId := 987 // int64 | ID of the dashboard list to delete items from.
    body := *datadog.NewDashboardListDeleteItemsRequest() // DashboardListDeleteItemsRequest | Dashboards to delete from the dashboard list.

    configuration := datadog.NewConfiguration()
    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.DashboardListsApi.DeleteDashboardListItems(context.Background(), dashboardListId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardListsApi.DeleteDashboardListItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDashboardListItems`: DashboardListDeleteItemsResponse
    fmt.Fprintf(os.Stdout, "Response from `DashboardListsApi.DeleteDashboardListItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardListId** | **int64** | ID of the dashboard list to delete items from. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDashboardListItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DashboardListDeleteItemsRequest**](DashboardListDeleteItemsRequest.md) | Dashboards to delete from the dashboard list. | 

### Return type

[**DashboardListDeleteItemsResponse**](DashboardListDeleteItemsResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDashboardListItems

> DashboardListItems GetDashboardListItems(ctx, dashboardListId).Execute()

Get a Dashboard List



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

    dashboardListId := 987 // int64 | ID of the dashboard list to get items from.

    configuration := datadog.NewConfiguration()
    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.DashboardListsApi.GetDashboardListItems(context.Background(), dashboardListId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardListsApi.GetDashboardListItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDashboardListItems`: DashboardListItems
    fmt.Fprintf(os.Stdout, "Response from `DashboardListsApi.GetDashboardListItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardListId** | **int64** | ID of the dashboard list to get items from. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDashboardListItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DashboardListItems**](DashboardListItems.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDashboardListItems

> DashboardListUpdateItemsResponse UpdateDashboardListItems(ctx, dashboardListId).Body(body).Execute()

Update items of a dashboard list



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

    dashboardListId := 987 // int64 | ID of the dashboard list to update items from.
    body := *datadog.NewDashboardListUpdateItemsRequest() // DashboardListUpdateItemsRequest | New dashboards of the dashboard list.

    configuration := datadog.NewConfiguration()
    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.DashboardListsApi.UpdateDashboardListItems(context.Background(), dashboardListId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardListsApi.UpdateDashboardListItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDashboardListItems`: DashboardListUpdateItemsResponse
    fmt.Fprintf(os.Stdout, "Response from `DashboardListsApi.UpdateDashboardListItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardListId** | **int64** | ID of the dashboard list to update items from. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDashboardListItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DashboardListUpdateItemsRequest**](DashboardListUpdateItemsRequest.md) | New dashboards of the dashboard list. | 

### Return type

[**DashboardListUpdateItemsResponse**](DashboardListUpdateItemsResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

