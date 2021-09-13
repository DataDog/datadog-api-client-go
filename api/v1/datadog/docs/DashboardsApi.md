# DashboardsApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------ | ------------ | ------------
[**CreateDashboard**](DashboardsApi.md#CreateDashboard) | **Post** /api/v1/dashboard | Create a new dashboard
[**DeleteDashboard**](DashboardsApi.md#DeleteDashboard) | **Delete** /api/v1/dashboard/{dashboard_id} | Delete a dashboard
[**DeleteDashboards**](DashboardsApi.md#DeleteDashboards) | **Delete** /api/v1/dashboard | Delete dashboards
[**GetDashboard**](DashboardsApi.md#GetDashboard) | **Get** /api/v1/dashboard/{dashboard_id} | Get a dashboard
[**ListDashboards**](DashboardsApi.md#ListDashboards) | **Get** /api/v1/dashboard | Get all dashboards
[**RestoreDashboards**](DashboardsApi.md#RestoreDashboards) | **Patch** /api/v1/dashboard | Restore deleted dashboards
[**UpdateDashboard**](DashboardsApi.md#UpdateDashboard) | **Put** /api/v1/dashboard/{dashboard_id} | Update a dashboard



## CreateDashboard

> Dashboard CreateDashboard(ctx, body)

Create a dashboard using the specified options. When defining queries in your widgets, take note of which queries should have the `as_count()` or `as_rate()` modifiers appended.
Refer to the following [documentation](https://docs.datadoghq.com/developers/metrics/type_modifiers/?tab=count#in-application-modifiers) for more information on these modifiers.

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
    resp, r, err := apiClient.DashboardsApi.CreateDashboard(ctx, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardsApi.CreateDashboard`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDashboard`: Dashboard
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from DashboardsApi.CreateDashboard:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**body** | [**Dashboard**](Dashboard.md) | Create a dashboard request body. | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**Dashboard**](Dashboard.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDashboard

> DashboardDeleteResponse DeleteDashboard(ctx, dashboardId)

Delete a dashboard using the specified ID.

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
    resp, r, err := apiClient.DashboardsApi.DeleteDashboard(ctx, dashboardId)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardsApi.DeleteDashboard`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDashboard`: DashboardDeleteResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from DashboardsApi.DeleteDashboard:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**dashboardId** | **string** | The ID of the dashboard. | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**DashboardDeleteResponse**](DashboardDeleteResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDashboards

> DeleteDashboards(ctx, body)

Delete dashboards using the specified IDs. If there are any failures, no dashboards will be deleted (partial success is not allowed).

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

    body := *datadog.NewDashboardBulkDeleteRequest([]datadog.DashboardBulkActionData{*datadog.NewDashboardBulkActionData("123-abc-456", datadog.DashboardResourceType("dashboard"))}) // DashboardBulkDeleteRequest | Delete dashboards request body.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    r, err := apiClient.DashboardsApi.DeleteDashboards(ctx, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardsApi.DeleteDashboards`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**body** | [**DashboardBulkDeleteRequest**](DashboardBulkDeleteRequest.md) | Delete dashboards request body. | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

 (empty response body)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDashboard

> Dashboard GetDashboard(ctx, dashboardId)

Get a dashboard using the specified ID.

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
    resp, r, err := apiClient.DashboardsApi.GetDashboard(ctx, dashboardId)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardsApi.GetDashboard`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDashboard`: Dashboard
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from DashboardsApi.GetDashboard:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**dashboardId** | **string** | The ID of the dashboard. | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**Dashboard**](Dashboard.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDashboards

> DashboardSummary ListDashboards(ctx, datadog.ListDashboardsOptionalParameters{})

Get all dashboards.

**Note**: This query will only return custom created or cloned dashboards.
This query will not return preset dashboards.

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

    filterShared := true // bool | When `true`, this query only returns shared custom created or cloned dashboards. (optional)
    optionalParams := datadog.ListDashboardsOptionalParameters{
        FilterShared: &filterShared,
    }

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.DashboardsApi.ListDashboards(ctx, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardsApi.ListDashboards`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDashboards`: DashboardSummary
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from DashboardsApi.ListDashboards:\n%s\n", responseContent)
}
```

### Required Parameters




### Optional Parameters


Other parameters are passed through a pointer to a ListDashboardsOptionalParameters struct.


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**filterShared** | **bool** | When &#x60;true&#x60;, this query only returns shared custom created or cloned dashboards. | 

### Return type

[**DashboardSummary**](DashboardSummary.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestoreDashboards

> RestoreDashboards(ctx, body)

Restore dashboards using the specified IDs. If there are any failures, no dashboards will be restored (partial success is not allowed).

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

    body := *datadog.NewDashboardRestoreRequest([]datadog.DashboardBulkActionData{*datadog.NewDashboardBulkActionData("123-abc-456", datadog.DashboardResourceType("dashboard"))}) // DashboardRestoreRequest | Restore dashboards request body.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    r, err := apiClient.DashboardsApi.RestoreDashboards(ctx, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardsApi.RestoreDashboards`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**body** | [**DashboardRestoreRequest**](DashboardRestoreRequest.md) | Restore dashboards request body. | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

 (empty response body)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDashboard

> Dashboard UpdateDashboard(ctx, dashboardId, body)

Update a dashboard using the specified ID.

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
    resp, r, err := apiClient.DashboardsApi.UpdateDashboard(ctx, dashboardId, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardsApi.UpdateDashboard`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDashboard`: Dashboard
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from DashboardsApi.UpdateDashboard:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**dashboardId** | **string** | The ID of the dashboard. |  |
**body** | [**Dashboard**](Dashboard.md) | Update Dashboard request body. | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**Dashboard**](Dashboard.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

