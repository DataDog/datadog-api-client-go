# DashboardListsApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------ | ------------ | ------------
[**CreateDashboardListItems**](DashboardListsApi.md#CreateDashboardListItems) | **Post** /api/v2/dashboard/lists/manual/{dashboard_list_id}/dashboards | Add Items to a Dashboard List
[**DeleteDashboardListItems**](DashboardListsApi.md#DeleteDashboardListItems) | **Delete** /api/v2/dashboard/lists/manual/{dashboard_list_id}/dashboards | Delete items from a dashboard list
[**GetDashboardListItems**](DashboardListsApi.md#GetDashboardListItems) | **Get** /api/v2/dashboard/lists/manual/{dashboard_list_id}/dashboards | Get items of a Dashboard List
[**UpdateDashboardListItems**](DashboardListsApi.md#UpdateDashboardListItems) | **Put** /api/v2/dashboard/lists/manual/{dashboard_list_id}/dashboards | Update items of a dashboard list



## CreateDashboardListItems

> DashboardListAddItemsResponse CreateDashboardListItems(ctx, dashboardListId, body)

Add dashboards to an existing dashboard list.

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

    dashboardListId := int64(789) // int64 | ID of the dashboard list to add items to.
    body := *datadog.NewDashboardListAddItemsRequest() // DashboardListAddItemsRequest | Dashboards to add to the dashboard list.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.DashboardListsApi.CreateDashboardListItems(ctx, dashboardListId, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardListsApi.CreateDashboardListItems`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDashboardListItems`: DashboardListAddItemsResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from DashboardListsApi.CreateDashboardListItems:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**dashboardListId** | **int64** | ID of the dashboard list to add items to. |  |
**body** | [**DashboardListAddItemsRequest**](DashboardListAddItemsRequest.md) | Dashboards to add to the dashboard list. | 


### Optional Parameters

This endpoint does not have optional parameters.


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

> DashboardListDeleteItemsResponse DeleteDashboardListItems(ctx, dashboardListId, body)

Delete dashboards from an existing dashboard list.

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

    dashboardListId := int64(789) // int64 | ID of the dashboard list to delete items from.
    body := *datadog.NewDashboardListDeleteItemsRequest() // DashboardListDeleteItemsRequest | Dashboards to delete from the dashboard list.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.DashboardListsApi.DeleteDashboardListItems(ctx, dashboardListId, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardListsApi.DeleteDashboardListItems`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDashboardListItems`: DashboardListDeleteItemsResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from DashboardListsApi.DeleteDashboardListItems:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**dashboardListId** | **int64** | ID of the dashboard list to delete items from. |  |
**body** | [**DashboardListDeleteItemsRequest**](DashboardListDeleteItemsRequest.md) | Dashboards to delete from the dashboard list. | 


### Optional Parameters

This endpoint does not have optional parameters.


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

> DashboardListItems GetDashboardListItems(ctx, dashboardListId)

Fetch the dashboard list’s dashboard definitions.

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

    dashboardListId := int64(789) // int64 | ID of the dashboard list to get items from.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.DashboardListsApi.GetDashboardListItems(ctx, dashboardListId)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardListsApi.GetDashboardListItems`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDashboardListItems`: DashboardListItems
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from DashboardListsApi.GetDashboardListItems:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**dashboardListId** | **int64** | ID of the dashboard list to get items from. | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**DashboardListItems**](DashboardListItems.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDashboardListItems

> DashboardListUpdateItemsResponse UpdateDashboardListItems(ctx, dashboardListId, body)

Update dashboards of an existing dashboard list.

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

    dashboardListId := int64(789) // int64 | ID of the dashboard list to update items from.
    body := *datadog.NewDashboardListUpdateItemsRequest() // DashboardListUpdateItemsRequest | New dashboards of the dashboard list.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.DashboardListsApi.UpdateDashboardListItems(ctx, dashboardListId, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardListsApi.UpdateDashboardListItems`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDashboardListItems`: DashboardListUpdateItemsResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from DashboardListsApi.UpdateDashboardListItems:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**dashboardListId** | **int64** | ID of the dashboard list to update items from. |  |
**body** | [**DashboardListUpdateItemsRequest**](DashboardListUpdateItemsRequest.md) | New dashboards of the dashboard list. | 


### Optional Parameters

This endpoint does not have optional parameters.


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

