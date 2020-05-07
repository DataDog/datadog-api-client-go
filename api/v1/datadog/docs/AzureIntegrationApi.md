# \AzureIntegrationApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAzureIntegration**](AzureIntegrationApi.md#CreateAzureIntegration) | **Post** /api/v1/integration/azure | Create an Azure integration
[**DeleteAzureIntegration**](AzureIntegrationApi.md#DeleteAzureIntegration) | **Delete** /api/v1/integration/azure | Delete an Azure integration
[**ListAzureIntegration**](AzureIntegrationApi.md#ListAzureIntegration) | **Get** /api/v1/integration/azure | List all Azure integrations
[**UpdateAzureHostFilters**](AzureIntegrationApi.md#UpdateAzureHostFilters) | **Post** /api/v1/integration/azure/host_filters | Update Azure integration host filters
[**UpdateAzureIntegration**](AzureIntegrationApi.md#UpdateAzureIntegration) | **Put** /api/v1/integration/azure | Update an Azure integration



## CreateAzureIntegration

> interface{} CreateAzureIntegration(ctx).Body(body).Execute()

Create an Azure integration



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := openapiclient.AzureAccount{ClientId: "ClientId_example", ClientSecret: "ClientSecret_example", Errors: []string{"Errors_example"), HostFilters: "HostFilters_example", NewClientId: "NewClientId_example", NewTenantName: "NewTenantName_example", TenantName: "TenantName_example"} // AzureAccount | Create a Datadog-Azure integration for your Datadog account request body.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AzureIntegrationApi.CreateAzureIntegration(context.Background(), body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AzureIntegrationApi.CreateAzureIntegration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAzureIntegration`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `AzureIntegrationApi.CreateAzureIntegration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAzureIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AzureAccount**](AzureAccount.md) | Create a Datadog-Azure integration for your Datadog account request body. | 

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

> interface{} DeleteAzureIntegration(ctx).Body(body).Execute()

Delete an Azure integration



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := openapiclient.AzureAccount{ClientId: "ClientId_example", ClientSecret: "ClientSecret_example", Errors: []string{"Errors_example"), HostFilters: "HostFilters_example", NewClientId: "NewClientId_example", NewTenantName: "NewTenantName_example", TenantName: "TenantName_example"} // AzureAccount | Delete a given Datadog-Azure integration request body.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AzureIntegrationApi.DeleteAzureIntegration(context.Background(), body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AzureIntegrationApi.DeleteAzureIntegration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAzureIntegration`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `AzureIntegrationApi.DeleteAzureIntegration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAzureIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AzureAccount**](AzureAccount.md) | Delete a given Datadog-Azure integration request body. | 

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

> []AzureAccount ListAzureIntegration(ctx).Execute()

List all Azure integrations



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AzureIntegrationApi.ListAzureIntegration(context.Background(), ).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AzureIntegrationApi.ListAzureIntegration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAzureIntegration`: []AzureAccount
    fmt.Fprintf(os.Stdout, "Response from `AzureIntegrationApi.ListAzureIntegration`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAzureIntegrationRequest struct via the builder pattern


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

> interface{} UpdateAzureHostFilters(ctx).Body(body).Execute()

Update Azure integration host filters



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body :=  // AzureAccount | Update a Datadog-Azure integration's host filters request body.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AzureIntegrationApi.UpdateAzureHostFilters(context.Background(), body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AzureIntegrationApi.UpdateAzureHostFilters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAzureHostFilters`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `AzureIntegrationApi.UpdateAzureHostFilters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAzureHostFiltersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AzureAccount**](AzureAccount.md) | Update a Datadog-Azure integration&#39;s host filters request body. | 

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

> interface{} UpdateAzureIntegration(ctx).Body(body).Execute()

Update an Azure integration



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body :=  // AzureAccount | Update a Datadog-Azure integration request body.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AzureIntegrationApi.UpdateAzureIntegration(context.Background(), body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AzureIntegrationApi.UpdateAzureIntegration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAzureIntegration`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `AzureIntegrationApi.UpdateAzureIntegration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAzureIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AzureAccount**](AzureAccount.md) | Update a Datadog-Azure integration request body. | 

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

