# \HostsApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHostTotals**](HostsApi.md#GetHostTotals) | **Get** /api/v1/hosts/totals | Get the total number of active hosts
[**ListHosts**](HostsApi.md#ListHosts) | **Get** /api/v1/hosts | Get all hosts for your organization
[**MuteHost**](HostsApi.md#MuteHost) | **Post** /api/v1/host/{host_name}/mute | Mute a host
[**UnmuteHost**](HostsApi.md#UnmuteHost) | **Post** /api/v1/host/{host_name}/unmute | Unmute a host



## GetHostTotals

> HostTotals GetHostTotals(ctx).From(from).Execute()

Get the total number of active hosts



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

    from := 987 // int64 | Number of seconds from which you want to get total number of active hosts. (optional)

    configuration := datadog.NewConfiguration()
    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.HostsApi.GetHostTotals(ctx).From(from).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HostsApi.GetHostTotals``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHostTotals`: HostTotals
    fmt.Fprintf(os.Stdout, "Response from `HostsApi.GetHostTotals`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHostTotalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **int64** | Number of seconds from which you want to get total number of active hosts. | 

### Return type

[**HostTotals**](HostTotals.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHosts

> HostListResponse ListHosts(ctx).Filter(filter).SortField(sortField).SortDir(sortDir).Start(start).Count(count).From(from).IncludeMutedHostsData(includeMutedHostsData).IncludeHostsMetadata(includeHostsMetadata).Execute()

Get all hosts for your organization



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

    filter := "filter_example" // string | String to filter search results. (optional)
    sortField := "sortField_example" // string | Sort hosts by this field. (optional)
    sortDir := "sortDir_example" // string | Direction of sort. Options include `asc` and `desc`. (optional)
    start := 987 // int64 | Host result to start search from. (optional)
    count := 987 // int64 | Number of hosts to return. Max 1000. (optional)
    from := 987 // int64 | Number of seconds since UNIX epoch from which you want to search your hosts. (optional)
    includeMutedHostsData := true // bool | Include information on the muted status of hosts and when the mute expires. (optional)
    includeHostsMetadata := true // bool | Include additional metadata about the hosts (agent_version, machine, platform, processor, etc.). (optional)

    configuration := datadog.NewConfiguration()
    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.HostsApi.ListHosts(ctx).Filter(filter).SortField(sortField).SortDir(sortDir).Start(start).Count(count).From(from).IncludeMutedHostsData(includeMutedHostsData).IncludeHostsMetadata(includeHostsMetadata).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HostsApi.ListHosts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListHosts`: HostListResponse
    fmt.Fprintf(os.Stdout, "Response from `HostsApi.ListHosts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListHostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | String to filter search results. | 
 **sortField** | **string** | Sort hosts by this field. | 
 **sortDir** | **string** | Direction of sort. Options include &#x60;asc&#x60; and &#x60;desc&#x60;. | 
 **start** | **int64** | Host result to start search from. | 
 **count** | **int64** | Number of hosts to return. Max 1000. | 
 **from** | **int64** | Number of seconds since UNIX epoch from which you want to search your hosts. | 
 **includeMutedHostsData** | **bool** | Include information on the muted status of hosts and when the mute expires. | 
 **includeHostsMetadata** | **bool** | Include additional metadata about the hosts (agent_version, machine, platform, processor, etc.). | 

### Return type

[**HostListResponse**](HostListResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MuteHost

> HostMuteResponse MuteHost(ctx, hostName).Body(body).Execute()

Mute a host



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

    hostName := "hostName_example" // string | Name of the host to mute.
    body := datadog.HostMuteSettings{End: int64(123), Message: "Message_example", Override: false} // HostMuteSettings | Mute a host request body. (optional)

    configuration := datadog.NewConfiguration()
    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.HostsApi.MuteHost(ctx, hostName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HostsApi.MuteHost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MuteHost`: HostMuteResponse
    fmt.Fprintf(os.Stdout, "Response from `HostsApi.MuteHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostName** | **string** | Name of the host to mute. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMuteHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**HostMuteSettings**](HostMuteSettings.md) | Mute a host request body. | 

### Return type

[**HostMuteResponse**](HostMuteResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnmuteHost

> HostMuteResponse UnmuteHost(ctx, hostName).Execute()

Unmute a host



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

    hostName := "hostName_example" // string | Name of the host to unmute.

    configuration := datadog.NewConfiguration()
    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.HostsApi.UnmuteHost(ctx, hostName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HostsApi.UnmuteHost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnmuteHost`: HostMuteResponse
    fmt.Fprintf(os.Stdout, "Response from `HostsApi.UnmuteHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostName** | **string** | Name of the host to unmute. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnmuteHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HostMuteResponse**](HostMuteResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

