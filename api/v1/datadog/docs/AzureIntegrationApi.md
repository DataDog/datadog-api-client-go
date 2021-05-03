# AzureIntegrationApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------ | ------------ | ------------
[**CreateAzureIntegration**](AzureIntegrationApi.md#CreateAzureIntegration) | **Post** /api/v1/integration/azure | Create an Azure integration
[**DeleteAzureIntegration**](AzureIntegrationApi.md#DeleteAzureIntegration) | **Delete** /api/v1/integration/azure | Delete an Azure integration
[**ListAzureIntegration**](AzureIntegrationApi.md#ListAzureIntegration) | **Get** /api/v1/integration/azure | List all Azure integrations
[**UpdateAzureHostFilters**](AzureIntegrationApi.md#UpdateAzureHostFilters) | **Post** /api/v1/integration/azure/host_filters | Update Azure integration host filters
[**UpdateAzureIntegration**](AzureIntegrationApi.md#UpdateAzureIntegration) | **Put** /api/v1/integration/azure | Update an Azure integration



## CreateAzureIntegration

> interface{} CreateAzureIntegration(ctx, body)

Create a Datadog-Azure integration.

Using the `POST` method updates your integration configuration by adding your new
configuration to the existing one in your Datadog organization.

Using the `PUT` method updates your integration configuration by replacing your
current configuration with the new one sent to your Datadog organization.

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

    body := *datadog.NewAzureAccount() // AzureAccount | Create a Datadog-Azure integration for your Datadog account request body.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.AzureIntegrationApi.CreateAzureIntegration(ctx, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AzureIntegrationApi.CreateAzureIntegration`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAzureIntegration`: interface{}
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from AzureIntegrationApi.CreateAzureIntegration:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**body** | [**AzureAccount**](AzureAccount.md) | Create a Datadog-Azure integration for your Datadog account request body. | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

**interface{}**

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAzureIntegration

> interface{} DeleteAzureIntegration(ctx, body)

Delete a given Datadog-Azure integration from your Datadog account.

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

    body := *datadog.NewAzureAccount() // AzureAccount | Delete a given Datadog-Azure integration request body.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.AzureIntegrationApi.DeleteAzureIntegration(ctx, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AzureIntegrationApi.DeleteAzureIntegration`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAzureIntegration`: interface{}
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from AzureIntegrationApi.DeleteAzureIntegration:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**body** | [**AzureAccount**](AzureAccount.md) | Delete a given Datadog-Azure integration request body. | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

**interface{}**

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAzureIntegration

> []AzureAccount ListAzureIntegration(ctx)

List all Datadog-Azure integrations configured in your Datadog account.

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
    resp, r, err := apiClient.AzureIntegrationApi.ListAzureIntegration(ctx)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AzureIntegrationApi.ListAzureIntegration`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAzureIntegration`: []AzureAccount
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from AzureIntegrationApi.ListAzureIntegration:\n%s\n", responseContent)
}
```

### Required Parameters

This endpoint does not need any parameter.


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**[]AzureAccount**](AzureAccount.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAzureHostFilters

> interface{} UpdateAzureHostFilters(ctx, body)

Update the defined list of host filters for a given Datadog-Azure integration.

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

    body := *datadog.NewAzureAccount() // AzureAccount | Update a Datadog-Azure integration's host filters request body.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.AzureIntegrationApi.UpdateAzureHostFilters(ctx, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AzureIntegrationApi.UpdateAzureHostFilters`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAzureHostFilters`: interface{}
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from AzureIntegrationApi.UpdateAzureHostFilters:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**body** | [**AzureAccount**](AzureAccount.md) | Update a Datadog-Azure integration&#39;s host filters request body. | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

**interface{}**

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAzureIntegration

> interface{} UpdateAzureIntegration(ctx, body)

Update a Datadog-Azure integration. Requires an existing `tenant_name` and `client_id`.
Any other fields supplied will overwrite existing values. To overwrite `tenant_name` or `client_id`,
use `new_tenant_name` and `new_client_id`. To leave a field unchanged, do not supply that field in the payload.

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

    body := *datadog.NewAzureAccount() // AzureAccount | Update a Datadog-Azure integration request body.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.AzureIntegrationApi.UpdateAzureIntegration(ctx, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AzureIntegrationApi.UpdateAzureIntegration`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAzureIntegration`: interface{}
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from AzureIntegrationApi.UpdateAzureIntegration:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**body** | [**AzureAccount**](AzureAccount.md) | Update a Datadog-Azure integration request body. | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

**interface{}**

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

