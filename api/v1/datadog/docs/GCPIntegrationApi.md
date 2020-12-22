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

    body := *datadog.NewGCPAccount() // GCPAccount | Create a Datadog-GCP integration.

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.GCPIntegrationApi.CreateGCPIntegration(ctx).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GCPIntegrationApi.CreateGCPIntegration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateGCPIntegration`: interface{}
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from GCPIntegrationApi.CreateGCPIntegration:\n%v\n", response_content)
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

    body := *datadog.NewGCPAccount() // GCPAccount | Delete a given Datadog-GCP integration.

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.GCPIntegrationApi.DeleteGCPIntegration(ctx).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GCPIntegrationApi.DeleteGCPIntegration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteGCPIntegration`: interface{}
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from GCPIntegrationApi.DeleteGCPIntegration:\n%v\n", response_content)
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


    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.GCPIntegrationApi.ListGCPIntegration(ctx).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GCPIntegrationApi.ListGCPIntegration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGCPIntegration`: []GCPAccount
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from GCPIntegrationApi.ListGCPIntegration:\n%v\n", response_content)
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

    body := *datadog.NewGCPAccount() // GCPAccount | Update a Datadog-GCP integration.

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.GCPIntegrationApi.UpdateGCPIntegration(ctx).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GCPIntegrationApi.UpdateGCPIntegration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateGCPIntegration`: interface{}
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from GCPIntegrationApi.UpdateGCPIntegration:\n%v\n", response_content)
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

