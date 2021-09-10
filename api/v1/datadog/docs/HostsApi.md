# HostsApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------ | ------------ | ------------
[**GetHostTotals**](HostsApi.md#GetHostTotals) | **Get** /api/v1/hosts/totals | Get the total number of active hosts
[**ListHosts**](HostsApi.md#ListHosts) | **Get** /api/v1/hosts | Get all hosts for your organization
[**MuteHost**](HostsApi.md#MuteHost) | **Post** /api/v1/host/{host_name}/mute | Mute a host
[**UnmuteHost**](HostsApi.md#UnmuteHost) | **Post** /api/v1/host/{host_name}/unmute | Unmute a host



## GetHostTotals

> HostTotals GetHostTotals(ctx, datadog.GetHostTotalsOptionalParameters{})

This endpoint returns the total number of active and up hosts in your Datadog account.
Active means the host has reported in the past hour, and up means it has reported in the past two hours.

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

    from := int64(789) // int64 | Number of seconds from which you want to get total number of active hosts. (optional)
    optionalParams := datadog.GetHostTotalsOptionalParameters{
        From: &from,
    }

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.HostsApi.GetHostTotals(ctx, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HostsApi.GetHostTotals`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHostTotals`: HostTotals
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from HostsApi.GetHostTotals:\n%s\n", responseContent)
}
```

### Required Parameters




### Optional Parameters


Other parameters are passed through a pointer to a GetHostTotalsOptionalParameters struct.


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
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

> HostListResponse ListHosts(ctx, datadog.ListHostsOptionalParameters{})

This endpoint allows searching for hosts by name, alias, or tag.
Hosts live within the past 3 hours are included by default.
Retention is 7 days.
Results are paginated with a max of 1000 results at a time.

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

    filter := "filter_example" // string | String to filter search results. (optional)
    sortField := "sortField_example" // string | Sort hosts by this field. (optional)
    sortDir := "sortDir_example" // string | Direction of sort. Options include `asc` and `desc`. (optional)
    start := int64(789) // int64 | Host result to start search from. (optional)
    count := int64(789) // int64 | Number of hosts to return. Max 1000. (optional)
    from := int64(789) // int64 | Number of seconds since UNIX epoch from which you want to search your hosts. (optional)
    includeMutedHostsData := true // bool | Include information on the muted status of hosts and when the mute expires. (optional)
    includeHostsMetadata := true // bool | Include additional metadata about the hosts (agent_version, machine, platform, processor, etc.). (optional)
    optionalParams := datadog.ListHostsOptionalParameters{
        Filter: &filter,
        SortField: &sortField,
        SortDir: &sortDir,
        Start: &start,
        Count: &count,
        From: &from,
        IncludeMutedHostsData: &includeMutedHostsData,
        IncludeHostsMetadata: &includeHostsMetadata,
    }

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.HostsApi.ListHosts(ctx, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HostsApi.ListHosts`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListHosts`: HostListResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from HostsApi.ListHosts:\n%s\n", responseContent)
}
```

### Required Parameters




### Optional Parameters


Other parameters are passed through a pointer to a ListHostsOptionalParameters struct.


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
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

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MuteHost

> HostMuteResponse MuteHost(ctx, hostName, body)

Mute a host.

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

    hostName := "hostName_example" // string | Name of the host to mute.
    body := *datadog.NewHostMuteSettings() // HostMuteSettings | Mute a host request body.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.HostsApi.MuteHost(ctx, hostName, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HostsApi.MuteHost`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MuteHost`: HostMuteResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from HostsApi.MuteHost:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**hostName** | **string** | Name of the host to mute. |  |
**body** | [**HostMuteSettings**](HostMuteSettings.md) | Mute a host request body. | 


### Optional Parameters

This endpoint does not have optional parameters.


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

> HostMuteResponse UnmuteHost(ctx, hostName)

Unmutes a host. This endpoint takes no JSON arguments.

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

    hostName := "hostName_example" // string | Name of the host to unmute.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.HostsApi.UnmuteHost(ctx, hostName)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HostsApi.UnmuteHost`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnmuteHost`: HostMuteResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from HostsApi.UnmuteHost:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**hostName** | **string** | Name of the host to unmute. | 


### Optional Parameters

This endpoint does not have optional parameters.


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

