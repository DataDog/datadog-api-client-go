# DowntimesApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------ | ------------ | ------------
[**CancelDowntime**](DowntimesApi.md#CancelDowntime) | **Delete** /api/v1/downtime/{downtime_id} | Cancel a downtime
[**CancelDowntimesByScope**](DowntimesApi.md#CancelDowntimesByScope) | **Post** /api/v1/downtime/cancel/by_scope | Cancel downtimes by scope
[**CreateDowntime**](DowntimesApi.md#CreateDowntime) | **Post** /api/v1/downtime | Schedule a downtime
[**GetDowntime**](DowntimesApi.md#GetDowntime) | **Get** /api/v1/downtime/{downtime_id} | Get a downtime
[**ListDowntimes**](DowntimesApi.md#ListDowntimes) | **Get** /api/v1/downtime | Get all downtimes
[**ListMonitorDowntimes**](DowntimesApi.md#ListMonitorDowntimes) | **Get** /api/v1/monitor/{monitor_id}/downtimes | Get all downtimes for a monitor
[**UpdateDowntime**](DowntimesApi.md#UpdateDowntime) | **Put** /api/v1/downtime/{downtime_id} | Update a downtime



## CancelDowntime

> CancelDowntime(ctx, downtimeId)

Cancel a downtime.

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

    downtimeId := int64(123456) // int64 | ID of the downtime to cancel.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    r, err := apiClient.DowntimesApi.CancelDowntime(ctx, downtimeId)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DowntimesApi.CancelDowntime`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**downtimeId** | **int64** | ID of the downtime to cancel. | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

 (empty response body)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelDowntimesByScope

> CanceledDowntimesIds CancelDowntimesByScope(ctx, body)

Delete all downtimes that match the scope of `X`.

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

    body := *datadog.NewCancelDowntimesByScopeRequest("host:myserver") // CancelDowntimesByScopeRequest | Scope to cancel downtimes for.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.DowntimesApi.CancelDowntimesByScope(ctx, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DowntimesApi.CancelDowntimesByScope`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CancelDowntimesByScope`: CanceledDowntimesIds
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from DowntimesApi.CancelDowntimesByScope:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**body** | [**CancelDowntimesByScopeRequest**](CancelDowntimesByScopeRequest.md) | Scope to cancel downtimes for. | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**CanceledDowntimesIds**](CanceledDowntimesIds.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDowntime

> Downtime CreateDowntime(ctx, body)

Schedule a downtime.

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

    body := *datadog.NewDowntime() // Downtime | Schedule a downtime request body.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.DowntimesApi.CreateDowntime(ctx, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DowntimesApi.CreateDowntime`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDowntime`: Downtime
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from DowntimesApi.CreateDowntime:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**body** | [**Downtime**](Downtime.md) | Schedule a downtime request body. | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**Downtime**](Downtime.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDowntime

> Downtime GetDowntime(ctx, downtimeId)

Get downtime detail by `downtime_id`.

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

    downtimeId := int64(123456) // int64 | ID of the downtime to fetch.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.DowntimesApi.GetDowntime(ctx, downtimeId)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DowntimesApi.GetDowntime`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDowntime`: Downtime
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from DowntimesApi.GetDowntime:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**downtimeId** | **int64** | ID of the downtime to fetch. | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**Downtime**](Downtime.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDowntimes

> []Downtime ListDowntimes(ctx, datadog.ListDowntimesOptionalParameters{})

Get all scheduled downtimes.

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

    currentOnly := true // bool | Only return downtimes that are active when the request is made. (optional)
    optionalParams := datadog.ListDowntimesOptionalParameters{
        CurrentOnly: &currentOnly,
    }

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.DowntimesApi.ListDowntimes(ctx, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DowntimesApi.ListDowntimes`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDowntimes`: []Downtime
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from DowntimesApi.ListDowntimes:\n%s\n", responseContent)
}
```

### Required Parameters




### Optional Parameters


Other parameters are passed through a pointer to a ListDowntimesOptionalParameters struct.


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**currentOnly** | **bool** | Only return downtimes that are active when the request is made. | 

### Return type

[**[]Downtime**](Downtime.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMonitorDowntimes

> []Downtime ListMonitorDowntimes(ctx, monitorId)

Get all downtimes for the specified monitor

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

    monitorId := int64(789) // int64 | The id of the monitor

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.DowntimesApi.ListMonitorDowntimes(ctx, monitorId)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DowntimesApi.ListMonitorDowntimes`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMonitorDowntimes`: []Downtime
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from DowntimesApi.ListMonitorDowntimes:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**monitorId** | **int64** | The id of the monitor | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**[]Downtime**](Downtime.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDowntime

> Downtime UpdateDowntime(ctx, downtimeId, body)

Update a single downtime by `downtime_id`.

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

    downtimeId := int64(123456) // int64 | ID of the downtime to update.
    body := *datadog.NewDowntime() // Downtime | Update a downtime request body.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.DowntimesApi.UpdateDowntime(ctx, downtimeId, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DowntimesApi.UpdateDowntime`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDowntime`: Downtime
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from DowntimesApi.UpdateDowntime:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**downtimeId** | **int64** | ID of the downtime to update. |  |
**body** | [**Downtime**](Downtime.md) | Update a downtime request body. | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**Downtime**](Downtime.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

