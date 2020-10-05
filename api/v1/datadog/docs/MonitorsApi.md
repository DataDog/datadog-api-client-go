# \MonitorsApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckCanDeleteMonitor**](MonitorsApi.md#CheckCanDeleteMonitor) | **Get** /api/v1/monitor/can_delete | Check if a monitor can be deleted
[**CreateMonitor**](MonitorsApi.md#CreateMonitor) | **Post** /api/v1/monitor | Create a monitor
[**DeleteMonitor**](MonitorsApi.md#DeleteMonitor) | **Delete** /api/v1/monitor/{monitor_id} | Delete a monitor
[**GetMonitor**](MonitorsApi.md#GetMonitor) | **Get** /api/v1/monitor/{monitor_id} | Get a monitor&#39;s details
[**ListMonitors**](MonitorsApi.md#ListMonitors) | **Get** /api/v1/monitor | Get all monitor details
[**UpdateMonitor**](MonitorsApi.md#UpdateMonitor) | **Put** /api/v1/monitor/{monitor_id} | Edit a monitor
[**ValidateMonitor**](MonitorsApi.md#ValidateMonitor) | **Post** /api/v1/monitor/validate | Validate a monitor



## CheckCanDeleteMonitor

> CheckCanDeleteMonitorResponse CheckCanDeleteMonitor(ctx).MonitorIds(monitorIds).Execute()

Check if a monitor can be deleted



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

    monitorIds := []int64{int64(123)} // []int64 | The IDs of the monitor to check.

    configuration := datadog.NewConfiguration()
    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.MonitorsApi.CheckCanDeleteMonitor(context.Background()).MonitorIds(monitorIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorsApi.CheckCanDeleteMonitor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckCanDeleteMonitor`: CheckCanDeleteMonitorResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitorsApi.CheckCanDeleteMonitor`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckCanDeleteMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **monitorIds** | [**[]int64**](int64.md) | The IDs of the monitor to check. | 

### Return type

[**CheckCanDeleteMonitorResponse**](CheckCanDeleteMonitorResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMonitor

> Monitor CreateMonitor(ctx).Body(body).Execute()

Create a monitor



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

    body := datadog.Monitor{Created: "TODO", Creator: datadog.Creator{Email: "Email_example", Handle: "Handle_example", Name: "Name_example"}, Deleted: "TODO", Id: int64(123), Message: "Message_example", Modified: "TODO", Multi: false, Name: "Name_example", Options: datadog.MonitorOptions{Aggregation: datadog.MonitorOptions_aggregation{GroupBy: "GroupBy_example", Metric: "Metric_example", Type: "Type_example"}, DeviceIds: []MonitorDeviceID{datadog.MonitorDeviceID{}), EnableLogsSample: false, EscalationMessage: "EscalationMessage_example", EvaluationDelay: int64(123), IncludeTags: false, Locked: false, MinFailureDuration: int64(123), MinLocationFailed: int64(123), NewHostDelay: int64(123), NoDataTimeframe: int64(123), NotifyAudit: false, NotifyNoData: false, RenotifyInterval: int64(123), RequireFullWindow: false, Silenced: map[string]string{ "Key" = "Value" }, SyntheticsCheckId: int64(123), ThresholdWindows: datadog.MonitorThresholdWindowOptions{RecoveryWindow: "RecoveryWindow_example", TriggerWindow: "TriggerWindow_example"}, Thresholds: datadog.MonitorThresholds{Critical: 123, CriticalRecovery: 123, Ok: 123, Unknown: 123, Warning: 123, WarningRecovery: 123}, TimeoutH: int64(123)}, OverallState: datadog.MonitorOverallStates{}, Query: "Query_example", State: datadog.MonitorState{Groups: map[string]string{ "Key" = "Value" }}, Tags: []string{"Tags_example"), Type: datadog.MonitorType{}} // Monitor | Create a monitor request body.

    configuration := datadog.NewConfiguration()
    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.MonitorsApi.CreateMonitor(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorsApi.CreateMonitor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMonitor`: Monitor
    fmt.Fprintf(os.Stdout, "Response from `MonitorsApi.CreateMonitor`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Monitor**](Monitor.md) | Create a monitor request body. | 

### Return type

[**Monitor**](Monitor.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMonitor

> DeletedMonitor DeleteMonitor(ctx, monitorId).Force(force).Execute()

Delete a monitor



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

    monitorId := 987 // int64 | The ID of the monitor.
    force := "force_example" // string | Delete the monitor even if it's referenced by other resources (e.g. SLO, composite monitor). (optional)

    configuration := datadog.NewConfiguration()
    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.MonitorsApi.DeleteMonitor(context.Background(), monitorId).Force(force).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorsApi.DeleteMonitor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMonitor`: DeletedMonitor
    fmt.Fprintf(os.Stdout, "Response from `MonitorsApi.DeleteMonitor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorId** | **int64** | The ID of the monitor. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **string** | Delete the monitor even if it&#39;s referenced by other resources (e.g. SLO, composite monitor). | 

### Return type

[**DeletedMonitor**](DeletedMonitor.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMonitor

> Monitor GetMonitor(ctx, monitorId).GroupStates(groupStates).Execute()

Get a monitor's details



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

    monitorId := 987 // int64 | The ID of the monitor
    groupStates := "groupStates_example" // string | When specified, shows additional information about the group states. Choose one or more from `all`, `alert`, `warn`, and `no data`. (optional)

    configuration := datadog.NewConfiguration()
    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.MonitorsApi.GetMonitor(context.Background(), monitorId).GroupStates(groupStates).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorsApi.GetMonitor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMonitor`: Monitor
    fmt.Fprintf(os.Stdout, "Response from `MonitorsApi.GetMonitor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorId** | **int64** | The ID of the monitor | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupStates** | **string** | When specified, shows additional information about the group states. Choose one or more from &#x60;all&#x60;, &#x60;alert&#x60;, &#x60;warn&#x60;, and &#x60;no data&#x60;. | 

### Return type

[**Monitor**](Monitor.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMonitors

> []Monitor ListMonitors(ctx).GroupStates(groupStates).Name(name).Tags(tags).MonitorTags(monitorTags).WithDowntimes(withDowntimes).IdOffset(idOffset).Page(page).PageSize(pageSize).Execute()

Get all monitor details



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

    groupStates := "groupStates_example" // string | When specified, shows additional information about the group states. Choose one or more from `all`, `alert`, `warn`, and `no data`. (optional)
    name := "name_example" // string | A string to filter monitors by name. (optional)
    tags := "tags_example" // string | A comma separated list indicating what tags, if any, should be used to filter the list of monitors by scope. For example, `host:host0`. (optional)
    monitorTags := "monitorTags_example" // string | A comma separated list indicating what service and/or custom tags, if any, should be used to filter the list of monitors. Tags created in the Datadog UI automatically have the service key prepended. For example, `service:my-app`. (optional)
    withDowntimes := true // bool | If this argument is set to true, then the returned data includes all current downtimes for each monitor. (optional)
    idOffset := 987 // int64 | Monitor ID offset. (optional)
    page := 987 // int64 | The page to start paginating from. If this argument is not specified, the request returns all monitors without pagination. (optional)
    pageSize := 987 // int32 | The number of monitors to return per page. If the page argument is not specified, the default behavior returns all monitors without a `page_size` limit. However, if page is specified and `page_size` is not, the argument defaults to 100. (optional)

    configuration := datadog.NewConfiguration()
    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.MonitorsApi.ListMonitors(context.Background()).GroupStates(groupStates).Name(name).Tags(tags).MonitorTags(monitorTags).WithDowntimes(withDowntimes).IdOffset(idOffset).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorsApi.ListMonitors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMonitors`: []Monitor
    fmt.Fprintf(os.Stdout, "Response from `MonitorsApi.ListMonitors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMonitorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupStates** | **string** | When specified, shows additional information about the group states. Choose one or more from &#x60;all&#x60;, &#x60;alert&#x60;, &#x60;warn&#x60;, and &#x60;no data&#x60;. | 
 **name** | **string** | A string to filter monitors by name. | 
 **tags** | **string** | A comma separated list indicating what tags, if any, should be used to filter the list of monitors by scope. For example, &#x60;host:host0&#x60;. | 
 **monitorTags** | **string** | A comma separated list indicating what service and/or custom tags, if any, should be used to filter the list of monitors. Tags created in the Datadog UI automatically have the service key prepended. For example, &#x60;service:my-app&#x60;. | 
 **withDowntimes** | **bool** | If this argument is set to true, then the returned data includes all current downtimes for each monitor. | 
 **idOffset** | **int64** | Monitor ID offset. | 
 **page** | **int64** | The page to start paginating from. If this argument is not specified, the request returns all monitors without pagination. | 
 **pageSize** | **int32** | The number of monitors to return per page. If the page argument is not specified, the default behavior returns all monitors without a &#x60;page_size&#x60; limit. However, if page is specified and &#x60;page_size&#x60; is not, the argument defaults to 100. | 

### Return type

[**[]Monitor**](Monitor.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMonitor

> Monitor UpdateMonitor(ctx, monitorId).Body(body).Execute()

Edit a monitor



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

    monitorId := 987 // int64 | The ID of the monitor.
    body := datadog.MonitorUpdateRequest{Created: "TODO", Creator: datadog.Creator{Email: "Email_example", Handle: "Handle_example", Name: "Name_example"}, Deleted: "TODO", Id: int64(123), Message: "Message_example", Modified: "TODO", Multi: false, Name: "Name_example", Options: datadog.MonitorOptions{Aggregation: datadog.MonitorOptions_aggregation{GroupBy: "GroupBy_example", Metric: "Metric_example", Type: "Type_example"}, DeviceIds: []MonitorDeviceID{datadog.MonitorDeviceID{}), EnableLogsSample: false, EscalationMessage: "EscalationMessage_example", EvaluationDelay: int64(123), IncludeTags: false, Locked: false, MinFailureDuration: int64(123), MinLocationFailed: int64(123), NewHostDelay: int64(123), NoDataTimeframe: int64(123), NotifyAudit: false, NotifyNoData: false, RenotifyInterval: int64(123), RequireFullWindow: false, Silenced: map[string]string{ "Key" = "Value" }, SyntheticsCheckId: int64(123), ThresholdWindows: datadog.MonitorThresholdWindowOptions{RecoveryWindow: "RecoveryWindow_example", TriggerWindow: "TriggerWindow_example"}, Thresholds: datadog.MonitorThresholds{Critical: 123, CriticalRecovery: 123, Ok: 123, Unknown: 123, Warning: 123, WarningRecovery: 123}, TimeoutH: int64(123)}, OverallState: datadog.MonitorOverallStates{}, Query: "Query_example", State: datadog.MonitorState{Groups: map[string]string{ "Key" = "Value" }}, Tags: []string{"Tags_example"), Type: datadog.MonitorType{}} // MonitorUpdateRequest | Edit a monitor request body.

    configuration := datadog.NewConfiguration()
    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.MonitorsApi.UpdateMonitor(context.Background(), monitorId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorsApi.UpdateMonitor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMonitor`: Monitor
    fmt.Fprintf(os.Stdout, "Response from `MonitorsApi.UpdateMonitor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorId** | **int64** | The ID of the monitor. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**MonitorUpdateRequest**](MonitorUpdateRequest.md) | Edit a monitor request body. | 

### Return type

[**Monitor**](Monitor.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateMonitor

> Monitor ValidateMonitor(ctx).Body(body).Execute()

Validate a monitor



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

    body := datadog.Monitor{Created: "TODO", Creator: , Deleted: "TODO", Id: int64(123), Message: "Message_example", Modified: "TODO", Multi: false, Name: "Name_example", Options: , OverallState: , Query: "Query_example", State: , Tags: []string{"Tags_example"), Type: } // Monitor | Monitor request object

    configuration := datadog.NewConfiguration()
    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.MonitorsApi.ValidateMonitor(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorsApi.ValidateMonitor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ValidateMonitor`: Monitor
    fmt.Fprintf(os.Stdout, "Response from `MonitorsApi.ValidateMonitor`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Monitor**](Monitor.md) | Monitor request object | 

### Return type

[**Monitor**](Monitor.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

