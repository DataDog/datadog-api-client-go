# \PagerDutyIntegrationApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePagerDutyIntegrationService**](PagerDutyIntegrationApi.md#CreatePagerDutyIntegrationService) | **Post** /api/v1/integration/pagerduty/configuration/services | Create a new service object
[**DeletePagerDutyIntegrationService**](PagerDutyIntegrationApi.md#DeletePagerDutyIntegrationService) | **Delete** /api/v1/integration/pagerduty/configuration/services/{service_name} | Delete a single service object
[**GetPagerDutyIntegrationService**](PagerDutyIntegrationApi.md#GetPagerDutyIntegrationService) | **Get** /api/v1/integration/pagerduty/configuration/services/{service_name} | Get a single service object
[**UpdatePagerDutyIntegrationService**](PagerDutyIntegrationApi.md#UpdatePagerDutyIntegrationService) | **Put** /api/v1/integration/pagerduty/configuration/services/{service_name} | Update a single service object



## CreatePagerDutyIntegrationService

> PagerDutyServiceName CreatePagerDutyIntegrationService(ctx, body)

Create a new service object



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

    body := *datadog.NewPagerDutyService("ServiceKey_example", "ServiceName_example") // PagerDutyService | Create a new service object request body.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.PagerDutyIntegrationApi.CreatePagerDutyIntegrationService(ctx, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PagerDutyIntegrationApi.CreatePagerDutyIntegrationService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePagerDutyIntegrationService`: PagerDutyServiceName
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from PagerDutyIntegrationApi.CreatePagerDutyIntegrationService:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**PagerDutyService**](PagerDutyService.md) | Create a new service object request body. | 

### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**PagerDutyServiceName**](PagerDutyServiceName.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePagerDutyIntegrationService

> DeletePagerDutyIntegrationService(ctx, serviceName)

Delete a single service object



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

    serviceName := "serviceName_example" // string | The service name

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    r, err := apiClient.PagerDutyIntegrationApi.DeletePagerDutyIntegrationService(ctx, serviceName)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PagerDutyIntegrationApi.DeletePagerDutyIntegrationService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceName** | **string** | The service name | 

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


## GetPagerDutyIntegrationService

> PagerDutyServiceName GetPagerDutyIntegrationService(ctx, serviceName)

Get a single service object



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

    serviceName := "serviceName_example" // string | The service name.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.PagerDutyIntegrationApi.GetPagerDutyIntegrationService(ctx, serviceName)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PagerDutyIntegrationApi.GetPagerDutyIntegrationService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPagerDutyIntegrationService`: PagerDutyServiceName
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from PagerDutyIntegrationApi.GetPagerDutyIntegrationService:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceName** | **string** | The service name. | 

### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**PagerDutyServiceName**](PagerDutyServiceName.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePagerDutyIntegrationService

> UpdatePagerDutyIntegrationService(ctx, serviceName, body)

Update a single service object



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

    serviceName := "serviceName_example" // string | The service name
    body := *datadog.NewPagerDutyServiceKey("ServiceKey_example") // PagerDutyServiceKey | Update an existing service object request body.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    r, err := apiClient.PagerDutyIntegrationApi.UpdatePagerDutyIntegrationService(ctx, serviceName, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PagerDutyIntegrationApi.UpdatePagerDutyIntegrationService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceName** | **string** | The service name | 
**body** | [**PagerDutyServiceKey**](PagerDutyServiceKey.md) | Update an existing service object request body. | 

### Optional Parameters

This endpoint does not have optional parameters.


### Return type

 (empty response body)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

