# \IncidentServicesApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIncidentService**](IncidentServicesApi.md#CreateIncidentService) | **Post** /api/v2/services | Create a new incident service
[**DeleteIncidentService**](IncidentServicesApi.md#DeleteIncidentService) | **Delete** /api/v2/services/{service_id} | Delete an existing incident service
[**GetIncidentService**](IncidentServicesApi.md#GetIncidentService) | **Get** /api/v2/services/{service_id} | Get details of an incident service
[**ListIncidentServices**](IncidentServicesApi.md#ListIncidentServices) | **Get** /api/v2/services | Get a list of all incident services
[**UpdateIncidentService**](IncidentServicesApi.md#UpdateIncidentService) | **Patch** /api/v2/services/{service_id} | Update an existing incident service



## CreateIncidentService

> IncidentServiceResponse CreateIncidentService(ctx).Body(body).Execute()

Create a new incident service



### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
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

    body := *datadog.NewIncidentServiceCreateRequest(*datadog.NewIncidentServiceCreateData(datadog.IncidentServiceType("services"))) // IncidentServiceCreateRequest | Incident Service Payload.

    configuration := datadog.NewConfiguration()
    configuration.SetUnstableOperationEnabled("CreateIncidentService", true)

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.IncidentServicesApi.CreateIncidentService(ctx).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IncidentServicesApi.CreateIncidentService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIncidentService`: IncidentServiceResponse
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from IncidentServicesApi.CreateIncidentService:\n%v\n", response_content)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIncidentServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**IncidentServiceCreateRequest**](IncidentServiceCreateRequest.md) | Incident Service Payload. | 

### Return type

[**IncidentServiceResponse**](IncidentServiceResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIncidentService

> DeleteIncidentService(ctx, serviceId).Execute()

Delete an existing incident service



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
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

    serviceId := "serviceId_example" // string | The ID of the incident service.

    configuration := datadog.NewConfiguration()
    configuration.SetUnstableOperationEnabled("DeleteIncidentService", true)

    api_client := datadog.NewAPIClient(configuration)
    r, err := api_client.IncidentServicesApi.DeleteIncidentService(ctx, serviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IncidentServicesApi.DeleteIncidentService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The ID of the incident service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIncidentServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetIncidentService

> IncidentServiceResponse GetIncidentService(ctx, serviceId).Include(include).Execute()

Get details of an incident service



### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
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

    serviceId := "serviceId_example" // string | The ID of the incident service.
    include := "include_example" // string | Specifies which types of related objects should be included in the response. (optional)

    configuration := datadog.NewConfiguration()
    configuration.SetUnstableOperationEnabled("GetIncidentService", true)

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.IncidentServicesApi.GetIncidentService(ctx, serviceId).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IncidentServicesApi.GetIncidentService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIncidentService`: IncidentServiceResponse
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from IncidentServicesApi.GetIncidentService:\n%v\n", response_content)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The ID of the incident service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIncidentServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **string** | Specifies which types of related objects should be included in the response. | 

### Return type

[**IncidentServiceResponse**](IncidentServiceResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIncidentServices

> IncidentServicesResponse ListIncidentServices(ctx).Include(include).PageSize(pageSize).PageOffset(pageOffset).Filter(filter).Execute()

Get a list of all incident services



### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
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

    include := "include_example" // string | Specifies which types of related objects should be included in the response. (optional)
    pageSize := int64(789) // int64 | Size for a given page. (optional) (default to 10)
    pageOffset := int64(789) // int64 | Specific offset to use as the beginning of the returned page. (optional) (default to 0)
    filter := "ExampleServiceName" // string | A search query that filters services by name. (optional)

    configuration := datadog.NewConfiguration()
    configuration.SetUnstableOperationEnabled("ListIncidentServices", true)

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.IncidentServicesApi.ListIncidentServices(ctx).Include(include).PageSize(pageSize).PageOffset(pageOffset).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IncidentServicesApi.ListIncidentServices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIncidentServices`: IncidentServicesResponse
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from IncidentServicesApi.ListIncidentServices:\n%v\n", response_content)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListIncidentServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **string** | Specifies which types of related objects should be included in the response. | 
 **pageSize** | **int64** | Size for a given page. | [default to 10]
 **pageOffset** | **int64** | Specific offset to use as the beginning of the returned page. | [default to 0]
 **filter** | **string** | A search query that filters services by name. | 

### Return type

[**IncidentServicesResponse**](IncidentServicesResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIncidentService

> IncidentServiceResponse UpdateIncidentService(ctx, serviceId).Body(body).Execute()

Update an existing incident service



### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
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

    serviceId := "serviceId_example" // string | The ID of the incident service.
    body := *datadog.NewIncidentServiceUpdateRequest(*datadog.NewIncidentServiceUpdateData("00000000-0000-0000-0000-000000000000", datadog.IncidentServiceType("services"))) // IncidentServiceUpdateRequest | Incident Service Payload.

    configuration := datadog.NewConfiguration()
    configuration.SetUnstableOperationEnabled("UpdateIncidentService", true)

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.IncidentServicesApi.UpdateIncidentService(ctx, serviceId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IncidentServicesApi.UpdateIncidentService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIncidentService`: IncidentServiceResponse
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from IncidentServicesApi.UpdateIncidentService:\n%v\n", response_content)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | The ID of the incident service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIncidentServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**IncidentServiceUpdateRequest**](IncidentServiceUpdateRequest.md) | Incident Service Payload. | 

### Return type

[**IncidentServiceResponse**](IncidentServiceResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

