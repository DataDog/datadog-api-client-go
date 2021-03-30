# \DashboardsApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDashboard**](DashboardsApi.md#CreateDashboard) | **Post** /api/v1/dashboard | Create a new dashboard
[**DeleteDashboard**](DashboardsApi.md#DeleteDashboard) | **Delete** /api/v1/dashboard/{dashboard_id} | Delete a dashboard
[**GetDashboard**](DashboardsApi.md#GetDashboard) | **Get** /api/v1/dashboard/{dashboard_id} | Get a dashboard
[**ListDashboards**](DashboardsApi.md#ListDashboards) | **Get** /api/v1/dashboard | Get all dashboards
[**UpdateDashboard**](DashboardsApi.md#UpdateDashboard) | **Put** /api/v1/dashboard/{dashboard_id} | Update a dashboard



## CreateDashboard

> Dashboard CreateDashboard(ctx).Body(body).Execute()

Create a new dashboard



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

    body := *datadog.NewDashboard(datadog.DashboardLayoutType("ordered"), "Title_example", []datadog.Widget{*datadog.NewWidget(datadog.WidgetDefinition{AlertGraphWidgetDefinition: datadog.NewAlertGraphWidgetDefinition("AlertId_example", datadog.AlertGraphWidgetDefinitionType("alert_graph"), datadog.WidgetVizType("timeseries"))})}) // Dashboard | Create a dashboard request body.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.DashboardsApi.CreateDashboard(ctx).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardsApi.CreateDashboard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDashboard`: Dashboard
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from DashboardsApi.CreateDashboard:\n%s\n", responseContent)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Dashboard**](Dashboard.md) | Create a dashboard request body. | 

### Return type

[**Dashboard**](Dashboard.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDashboard

> DashboardDeleteResponse DeleteDashboard(ctx, dashboardId).Execute()

Delete a dashboard



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

    dashboardId := "dashboardId_example" // string | The ID of the dashboard.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.DashboardsApi.DeleteDashboard(ctx, dashboardId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardsApi.DeleteDashboard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDashboard`: DashboardDeleteResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from DashboardsApi.DeleteDashboard:\n%s\n", responseContent)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardId** | **string** | The ID of the dashboard. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DashboardDeleteResponse**](DashboardDeleteResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDashboard

> Dashboard GetDashboard(ctx, dashboardId).Execute()

Get a dashboard



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

    dashboardId := "dashboardId_example" // string | The ID of the dashboard.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.DashboardsApi.GetDashboard(ctx, dashboardId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardsApi.GetDashboard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDashboard`: Dashboard
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from DashboardsApi.GetDashboard:\n%s\n", responseContent)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardId** | **string** | The ID of the dashboard. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Dashboard**](Dashboard.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDashboards

> DashboardSummary ListDashboards(ctx).Execute()

Get all dashboards



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
    resp, r, err := apiClient.DashboardsApi.ListDashboards(ctx).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardsApi.ListDashboards``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDashboards`: DashboardSummary
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from DashboardsApi.ListDashboards:\n%s\n", responseContent)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListDashboardsRequest struct via the builder pattern


### Return type

[**DashboardSummary**](DashboardSummary.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDashboard

> Dashboard UpdateDashboard(ctx, dashboardId).Body(body).Execute()

Update a dashboard



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

    dashboardId := "dashboardId_example" // string | The ID of the dashboard.
    body := *datadog.NewDashboard(datadog.DashboardLayoutType("ordered"), "Title_example", []datadog.Widget{*datadog.NewWidget(datadog.WidgetDefinition{AlertGraphWidgetDefinition: datadog.NewAlertGraphWidgetDefinition("AlertId_example", datadog.AlertGraphWidgetDefinitionType("alert_graph"), datadog.WidgetVizType("timeseries"))})}) // Dashboard | Update Dashboard request body.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.DashboardsApi.UpdateDashboard(ctx, dashboardId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardsApi.UpdateDashboard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDashboard`: Dashboard
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from DashboardsApi.UpdateDashboard:\n%s\n", responseContent)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardId** | **string** | The ID of the dashboard. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**Dashboard**](Dashboard.md) | Update Dashboard request body. | 

### Return type

[**Dashboard**](Dashboard.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

