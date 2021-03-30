# \PagerDutyIntegrationApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePagerDutyIntegrationService**](PagerDutyIntegrationApi.md#CreatePagerDutyIntegrationService) | **Post** /api/v1/integration/pagerduty/configuration/services | Create a new service object
[**DeletePagerDutyIntegrationService**](PagerDutyIntegrationApi.md#DeletePagerDutyIntegrationService) | **Delete** /api/v1/integration/pagerduty/configuration/services/{service_name} | Delete a single service object
[**GetPagerDutyIntegrationService**](PagerDutyIntegrationApi.md#GetPagerDutyIntegrationService) | **Get** /api/v1/integration/pagerduty/configuration/services/{service_name} | Get a single service object
[**UpdatePagerDutyIntegrationService**](PagerDutyIntegrationApi.md#UpdatePagerDutyIntegrationService) | **Put** /api/v1/integration/pagerduty/configuration/services/{service_name} | Update a single service object



## CreatePagerDutyIntegrationService

> PagerDutyServiceName CreatePagerDutyIntegrationService(ctx).Body(body).Execute()

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
    resp, r, err := apiClient.PagerDutyIntegrationApi.CreatePagerDutyIntegrationService(ctx).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PagerDutyIntegrationApi.CreatePagerDutyIntegrationService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePagerDutyIntegrationService`: PagerDutyServiceName
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from PagerDutyIntegrationApi.CreatePagerDutyIntegrationService:\n%s\n", responseContent)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePagerDutyIntegrationServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PagerDutyService**](PagerDutyService.md) | Create a new service object request body. | 

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

> DeletePagerDutyIntegrationService(ctx, serviceName).Execute()

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
    r, err := apiClient.PagerDutyIntegrationApi.DeletePagerDutyIntegrationService(ctx, serviceName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PagerDutyIntegrationApi.DeletePagerDutyIntegrationService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceName** | **string** | The service name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePagerDutyIntegrationServiceRequest struct via the builder pattern


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


## GetPagerDutyIntegrationService

> PagerDutyServiceName GetPagerDutyIntegrationService(ctx, serviceName).Execute()

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
    resp, r, err := apiClient.PagerDutyIntegrationApi.GetPagerDutyIntegrationService(ctx, serviceName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PagerDutyIntegrationApi.GetPagerDutyIntegrationService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPagerDutyIntegrationService`: PagerDutyServiceName
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from PagerDutyIntegrationApi.GetPagerDutyIntegrationService:\n%s\n", responseContent)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceName** | **string** | The service name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPagerDutyIntegrationServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> UpdatePagerDutyIntegrationService(ctx, serviceName).Body(body).Execute()

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
    r, err := apiClient.PagerDutyIntegrationApi.UpdatePagerDutyIntegrationService(ctx, serviceName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PagerDutyIntegrationApi.UpdatePagerDutyIntegrationService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceName** | **string** | The service name | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePagerDutyIntegrationServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**PagerDutyServiceKey**](PagerDutyServiceKey.md) | Update an existing service object request body. | 

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

