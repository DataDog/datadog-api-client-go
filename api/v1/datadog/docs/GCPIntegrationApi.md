# \GCPIntegrationApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGCPIntegration**](GCPIntegrationApi.md#CreateGCPIntegration) | **Post** /api/v1/integration/gcp | Create a GCP integration
[**DeleteGCPIntegration**](GCPIntegrationApi.md#DeleteGCPIntegration) | **Delete** /api/v1/integration/gcp | Delete a GCP integration
[**ListGCPIntegration**](GCPIntegrationApi.md#ListGCPIntegration) | **Get** /api/v1/integration/gcp | List all GCP integrations
[**UpdateGCPIntegration**](GCPIntegrationApi.md#UpdateGCPIntegration) | **Put** /api/v1/integration/gcp | Update a GCP integration



## CreateGCPIntegration

> interface{} CreateGCPIntegration(ctx).Body(body).Execute()

Create a GCP integration



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

    body := *datadog.NewGCPAccount() // GCPAccount | Create a Datadog-GCP integration.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.GCPIntegrationApi.CreateGCPIntegration(ctx).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GCPIntegrationApi.CreateGCPIntegration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateGCPIntegration`: interface{}
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from GCPIntegrationApi.CreateGCPIntegration:\n%s\n", responseContent)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGCPIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GCPAccount**](GCPAccount.md) | Create a Datadog-GCP integration. | 

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


## DeleteGCPIntegration

> interface{} DeleteGCPIntegration(ctx).Body(body).Execute()

Delete a GCP integration



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

    body := *datadog.NewGCPAccount() // GCPAccount | Delete a given Datadog-GCP integration.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.GCPIntegrationApi.DeleteGCPIntegration(ctx).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GCPIntegrationApi.DeleteGCPIntegration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteGCPIntegration`: interface{}
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from GCPIntegrationApi.DeleteGCPIntegration:\n%s\n", responseContent)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGCPIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GCPAccount**](GCPAccount.md) | Delete a given Datadog-GCP integration. | 

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


## ListGCPIntegration

> []GCPAccount ListGCPIntegration(ctx).Execute()

List all GCP integrations



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
    resp, r, err := apiClient.GCPIntegrationApi.ListGCPIntegration(ctx).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GCPIntegrationApi.ListGCPIntegration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGCPIntegration`: []GCPAccount
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from GCPIntegrationApi.ListGCPIntegration:\n%s\n", responseContent)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListGCPIntegrationRequest struct via the builder pattern


### Return type

[**[]GCPAccount**](GCPAccount.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGCPIntegration

> interface{} UpdateGCPIntegration(ctx).Body(body).Execute()

Update a GCP integration



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

    body := *datadog.NewGCPAccount() // GCPAccount | Update a Datadog-GCP integration.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.GCPIntegrationApi.UpdateGCPIntegration(ctx).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GCPIntegrationApi.UpdateGCPIntegration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateGCPIntegration`: interface{}
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from GCPIntegrationApi.UpdateGCPIntegration:\n%s\n", responseContent)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGCPIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GCPAccount**](GCPAccount.md) | Update a Datadog-GCP integration. | 

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

