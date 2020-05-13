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
    dashboardListId := 987 // int64 | ID of the dashboard list to add items to.
    body := datadog.DashboardListItems{Dashboards: []DashboardListItem{datadog.DashboardListItem{Author: datadog.Creator{Email: "Email_example", Handle: "Handle_example", Name: "Name_example"}, Created: "TODO", Icon: "Icon_example", Id: "Id_example", IsFavorite: false, IsReadOnly: false, IsShared: false, Modified: "TODO", Popularity: 123, Title: "Title_example", Type: datadog.DashboardType{}, Url: "Url_example"}), Total: int64(123)} // DashboardListItems | Dashboards to add to the dashboard list.

    configuration := datadog.NewConfiguration()
    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.DashboardListsApi.CreateDashboardListItems(context.Background(), dashboardListId, body).Execute()
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

 **body** | [**DashboardListItems**](DashboardListItems.md) | Dashboards to add to the dashboard list. | 

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
    dashboardListId := 987 // int64 | ID of the dashboard list to delete items from.
    body := datadog.DashboardListItems{Dashboards: []DashboardListItem{datadog.DashboardListItem{Author: datadog.Creator{Email: "Email_example", Handle: "Handle_example", Name: "Name_example"}, Created: "TODO", Icon: "Icon_example", Id: "Id_example", IsFavorite: false, IsReadOnly: false, IsShared: false, Modified: "TODO", Popularity: 123, Title: "Title_example", Type: datadog.DashboardType{}, Url: "Url_example"}), Total: int64(123)} // DashboardListItems | Dashboards to delete from the dashboard list.

    configuration := datadog.NewConfiguration()
    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.DashboardListsApi.DeleteDashboardListItems(context.Background(), dashboardListId, body).Execute()
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

 **body** | [**DashboardListItems**](DashboardListItems.md) | Dashboards to delete from the dashboard list. | 

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

> DashboardListItems UpdateDashboardListItems(ctx, dashboardListId).Body(body).Execute()

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
    dashboardListId := 987 // int64 | ID of the dashboard list to update items from.
    body :=  // DashboardListItems | New dashboards of the dashboard list.

    configuration := datadog.NewConfiguration()
    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.DashboardListsApi.UpdateDashboardListItems(context.Background(), dashboardListId, body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardListsApi.UpdateDashboardListItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDashboardListItems`: DashboardListItems
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

 **body** | [**DashboardListItems**](DashboardListItems.md) | New dashboards of the dashboard list. | 

### Return type

[**DashboardListItems**](DashboardListItems.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

