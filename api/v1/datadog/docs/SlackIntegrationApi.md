# \SlackIntegrationApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSlackIntegrationChannel**](SlackIntegrationApi.md#CreateSlackIntegrationChannel) | **Post** /api/v1/integration/slack/configuration/accounts/{account_name}/channels | Create a Slack integration channel
[**GetSlackIntegrationChannel**](SlackIntegrationApi.md#GetSlackIntegrationChannel) | **Get** /api/v1/integration/slack/configuration/accounts/{account_name}/channels/{channel_name} | Get a Slack integration channel
[**GetSlackIntegrationChannels**](SlackIntegrationApi.md#GetSlackIntegrationChannels) | **Get** /api/v1/integration/slack/configuration/accounts/{account_name}/channels | Get all channels in a Slack integration
[**RemoveSlackIntegrationChannel**](SlackIntegrationApi.md#RemoveSlackIntegrationChannel) | **Delete** /api/v1/integration/slack/configuration/accounts/{account_name}/channels/{channel_name} | Remove a Slack integration channel
[**UpdateSlackIntegrationChannel**](SlackIntegrationApi.md#UpdateSlackIntegrationChannel) | **Patch** /api/v1/integration/slack/configuration/accounts/{account_name}/channels/{channel_name} | Update a Slack integration channel



## CreateSlackIntegrationChannel

> SlackIntegrationChannel CreateSlackIntegrationChannel(ctx, accountName, body)

Create a Slack integration channel



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

    accountName := "accountName_example" // string | Your Slack account name.
    body := *datadog.NewSlackIntegrationChannel() // SlackIntegrationChannel | Payload describing Slack channel to be created

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.SlackIntegrationApi.CreateSlackIntegrationChannel(ctx, accountName, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlackIntegrationApi.CreateSlackIntegrationChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSlackIntegrationChannel`: SlackIntegrationChannel
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from SlackIntegrationApi.CreateSlackIntegrationChannel:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** | Your Slack account name. | 
**body** | [**SlackIntegrationChannel**](SlackIntegrationChannel.md) | Payload describing Slack channel to be created | 

### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**SlackIntegrationChannel**](SlackIntegrationChannel.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSlackIntegrationChannel

> SlackIntegrationChannel GetSlackIntegrationChannel(ctx, accountName, channelName)

Get a Slack integration channel



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

    accountName := "accountName_example" // string | Your Slack account name.
    channelName := "channelName_example" // string | The name of the Slack channel being operated on.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.SlackIntegrationApi.GetSlackIntegrationChannel(ctx, accountName, channelName)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlackIntegrationApi.GetSlackIntegrationChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSlackIntegrationChannel`: SlackIntegrationChannel
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from SlackIntegrationApi.GetSlackIntegrationChannel:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** | Your Slack account name. | 
**channelName** | **string** | The name of the Slack channel being operated on. | 

### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**SlackIntegrationChannel**](SlackIntegrationChannel.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSlackIntegrationChannels

> []SlackIntegrationChannel GetSlackIntegrationChannels(ctx, accountName)

Get all channels in a Slack integration



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

    accountName := "accountName_example" // string | Your Slack account name.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.SlackIntegrationApi.GetSlackIntegrationChannels(ctx, accountName)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlackIntegrationApi.GetSlackIntegrationChannels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSlackIntegrationChannels`: []SlackIntegrationChannel
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from SlackIntegrationApi.GetSlackIntegrationChannels:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** | Your Slack account name. | 

### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**[]SlackIntegrationChannel**](SlackIntegrationChannel.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveSlackIntegrationChannel

> RemoveSlackIntegrationChannel(ctx, accountName, channelName)

Remove a Slack integration channel



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

    accountName := "accountName_example" // string | Your Slack account name.
    channelName := "channelName_example" // string | The name of the Slack channel being operated on.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    r, err := apiClient.SlackIntegrationApi.RemoveSlackIntegrationChannel(ctx, accountName, channelName)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlackIntegrationApi.RemoveSlackIntegrationChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** | Your Slack account name. | 
**channelName** | **string** | The name of the Slack channel being operated on. | 

### Optional Parameters

This endpoint does not have optional parameters.


### Return type

 (empty response body)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSlackIntegrationChannel

> SlackIntegrationChannel UpdateSlackIntegrationChannel(ctx, accountName, channelName, body)

Update a Slack integration channel



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

    accountName := "accountName_example" // string | Your Slack account name.
    channelName := "channelName_example" // string | The name of the Slack channel being operated on.
    body := *datadog.NewSlackIntegrationChannel() // SlackIntegrationChannel | Payload describing fields and values to be updated.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.SlackIntegrationApi.UpdateSlackIntegrationChannel(ctx, accountName, channelName, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlackIntegrationApi.UpdateSlackIntegrationChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSlackIntegrationChannel`: SlackIntegrationChannel
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from SlackIntegrationApi.UpdateSlackIntegrationChannel:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** | Your Slack account name. | 
**channelName** | **string** | The name of the Slack channel being operated on. | 
**body** | [**SlackIntegrationChannel**](SlackIntegrationChannel.md) | Payload describing fields and values to be updated. | 

### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**SlackIntegrationChannel**](SlackIntegrationChannel.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

