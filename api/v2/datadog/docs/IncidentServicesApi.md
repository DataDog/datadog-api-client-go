# IncidentServicesApi

All URIs are relative to *https://api.datadoghq.com*

| Method                                                                    | HTTP request                             | Description                         |
| ------------------------------------------------------------------------- | ---------------------------------------- | ----------------------------------- |
| [**CreateIncidentService**](IncidentServicesApi.md#CreateIncidentService) | **Post** /api/v2/services                | Create a new incident service       |
| [**DeleteIncidentService**](IncidentServicesApi.md#DeleteIncidentService) | **Delete** /api/v2/services/{service_id} | Delete an existing incident service |
| [**GetIncidentService**](IncidentServicesApi.md#GetIncidentService)       | **Get** /api/v2/services/{service_id}    | Get details of an incident service  |
| [**ListIncidentServices**](IncidentServicesApi.md#ListIncidentServices)   | **Get** /api/v2/services                 | Get a list of all incident services |
| [**UpdateIncidentService**](IncidentServicesApi.md#UpdateIncidentService) | **Patch** /api/v2/services/{service_id}  | Update an existing incident service |

## CreateIncidentService

> IncidentServiceResponse CreateIncidentService(ctx, body)

Creates a new incident service.

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
    ctx := datadog.NewDefaultContext(context.Background())

    body := *datadog.NewIncidentServiceCreateRequest(*datadog.NewIncidentServiceCreateData(datadog.IncidentServiceType("services"))) // IncidentServiceCreateRequest | Incident Service Payload.

    configuration := datadog.NewConfiguration()
    configuration.SetUnstableOperationEnabled("CreateIncidentService", true)

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.IncidentServicesApi.CreateIncidentService(ctx, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IncidentServicesApi.CreateIncidentService`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIncidentService`: IncidentServiceResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from IncidentServicesApi.CreateIncidentService:\n%s\n", responseContent)
}
```

### Required Parameters

| Name     | Type                                                                | Description                                                                 | Notes |
| -------- | ------------------------------------------------------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**  | **context.Context**                                                 | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **body** | [**IncidentServiceCreateRequest**](IncidentServiceCreateRequest.md) | Incident Service Payload.                                                   |

### Optional Parameters

This endpoint does not have optional parameters.

### Return type

[**IncidentServiceResponse**](IncidentServiceResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## DeleteIncidentService

> DeleteIncidentService(ctx, serviceId)

Deletes an existing incident service.

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
    ctx := datadog.NewDefaultContext(context.Background())

    serviceId := "serviceId_example" // string | The ID of the incident service.

    configuration := datadog.NewConfiguration()
    configuration.SetUnstableOperationEnabled("DeleteIncidentService", true)

    apiClient := datadog.NewAPIClient(configuration)
    r, err := apiClient.IncidentServicesApi.DeleteIncidentService(ctx, serviceId)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IncidentServicesApi.DeleteIncidentService`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Required Parameters

| Name          | Type                | Description                                                                 | Notes |
| ------------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**       | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **serviceId** | **string**          | The ID of the incident service.                                             |

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

## GetIncidentService

> IncidentServiceResponse GetIncidentService(ctx, serviceId, datadog.GetIncidentServiceOptionalParameters{})

Get details of an incident service. If the `include[users]` query parameter is provided,
the included attribute will contain the users related to these incident services.

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
    ctx := datadog.NewDefaultContext(context.Background())

    serviceId := "serviceId_example" // string | The ID of the incident service.
    include := datadog.IncidentRelatedObject("users") // IncidentRelatedObject | Specifies which types of related objects should be included in the response. (optional)
    optionalParams := datadog.GetIncidentServiceOptionalParameters{
        Include: &include,
    }

    configuration := datadog.NewConfiguration()
    configuration.SetUnstableOperationEnabled("GetIncidentService", true)

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.IncidentServicesApi.GetIncidentService(ctx, serviceId, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IncidentServicesApi.GetIncidentService`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIncidentService`: IncidentServiceResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from IncidentServicesApi.GetIncidentService:\n%s\n", responseContent)
}
```

### Required Parameters

| Name          | Type                | Description                                                                 | Notes |
| ------------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**       | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **serviceId** | **string**          | The ID of the incident service.                                             |

### Optional Parameters

Other parameters are passed through a pointer to a GetIncidentServiceOptionalParameters struct.

| Name        | Type                                                  | Description                                                                  | Notes |
| ----------- | ----------------------------------------------------- | ---------------------------------------------------------------------------- | ----- |
| **include** | [**IncidentRelatedObject**](IncidentRelatedObject.md) | Specifies which types of related objects should be included in the response. |

### Return type

[**IncidentServiceResponse**](IncidentServiceResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListIncidentServices

> IncidentServicesResponse ListIncidentServices(ctx, datadog.ListIncidentServicesOptionalParameters{})

Get all incident services uploaded for the requesting user's organization. If the `include[users]` query parameter is provided, the included attribute will contain the users related to these incident services.

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
    ctx := datadog.NewDefaultContext(context.Background())

    include := datadog.IncidentRelatedObject("users") // IncidentRelatedObject | Specifies which types of related objects should be included in the response. (optional)
    pageSize := int64(10) // int64 | Size for a given page. (optional) (default to 10)
    pageOffset := int64(0) // int64 | Specific offset to use as the beginning of the returned page. (optional) (default to 0)
    filter := "ExampleServiceName" // string | A search query that filters services by name. (optional)
    optionalParams := datadog.ListIncidentServicesOptionalParameters{
        Include: &include,
        PageSize: &pageSize,
        PageOffset: &pageOffset,
        Filter: &filter,
    }

    configuration := datadog.NewConfiguration()
    configuration.SetUnstableOperationEnabled("ListIncidentServices", true)

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.IncidentServicesApi.ListIncidentServices(ctx, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IncidentServicesApi.ListIncidentServices`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIncidentServices`: IncidentServicesResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from IncidentServicesApi.ListIncidentServices:\n%s\n", responseContent)
}
```

### Required Parameters

### Optional Parameters

Other parameters are passed through a pointer to a ListIncidentServicesOptionalParameters struct.

| Name           | Type                                                  | Description                                                                  | Notes           |
| -------------- | ----------------------------------------------------- | ---------------------------------------------------------------------------- | --------------- |
| **include**    | [**IncidentRelatedObject**](IncidentRelatedObject.md) | Specifies which types of related objects should be included in the response. |
| **pageSize**   | **int64**                                             | Size for a given page.                                                       | [default to 10] |
| **pageOffset** | **int64**                                             | Specific offset to use as the beginning of the returned page.                | [default to 0]  |
| **filter**     | **string**                                            | A search query that filters services by name.                                |

### Return type

[**IncidentServicesResponse**](IncidentServicesResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateIncidentService

> IncidentServiceResponse UpdateIncidentService(ctx, serviceId, body)

Updates an existing incident service. Only provide the attributes which should be updated as this request is a partial update.

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
    ctx := datadog.NewDefaultContext(context.Background())

    serviceId := "serviceId_example" // string | The ID of the incident service.
    body := *datadog.NewIncidentServiceUpdateRequest(*datadog.NewIncidentServiceUpdateData(datadog.IncidentServiceType("services"))) // IncidentServiceUpdateRequest | Incident Service Payload.

    configuration := datadog.NewConfiguration()
    configuration.SetUnstableOperationEnabled("UpdateIncidentService", true)

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.IncidentServicesApi.UpdateIncidentService(ctx, serviceId, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IncidentServicesApi.UpdateIncidentService`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIncidentService`: IncidentServiceResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from IncidentServicesApi.UpdateIncidentService:\n%s\n", responseContent)
}
```

### Required Parameters

| Name          | Type                                                                | Description                                                                 | Notes |
| ------------- | ------------------------------------------------------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**       | **context.Context**                                                 | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **serviceId** | **string**                                                          | The ID of the incident service.                                             |       |
| **body**      | [**IncidentServiceUpdateRequest**](IncidentServiceUpdateRequest.md) | Incident Service Payload.                                                   |

### Optional Parameters

This endpoint does not have optional parameters.

### Return type

[**IncidentServiceResponse**](IncidentServiceResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
