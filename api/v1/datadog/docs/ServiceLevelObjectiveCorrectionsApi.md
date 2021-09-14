# ServiceLevelObjectiveCorrectionsApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------ | ------------ | ------------
[**CreateSLOCorrection**](ServiceLevelObjectiveCorrectionsApi.md#CreateSLOCorrection) | **Post** /api/v1/slo/correction | Create an SLO correction
[**DeleteSLOCorrection**](ServiceLevelObjectiveCorrectionsApi.md#DeleteSLOCorrection) | **Delete** /api/v1/slo/correction/{slo_correction_id} | Delete an SLO correction
[**GetSLOCorrection**](ServiceLevelObjectiveCorrectionsApi.md#GetSLOCorrection) | **Get** /api/v1/slo/correction/{slo_correction_id} | Get an SLO correction for an SLO
[**ListSLOCorrection**](ServiceLevelObjectiveCorrectionsApi.md#ListSLOCorrection) | **Get** /api/v1/slo/correction | Get all SLO corrections
[**UpdateSLOCorrection**](ServiceLevelObjectiveCorrectionsApi.md#UpdateSLOCorrection) | **Patch** /api/v1/slo/correction/{slo_correction_id} | Update an SLO correction



## CreateSLOCorrection

> SLOCorrectionResponse CreateSLOCorrection(ctx, body)

Create an SLO Correction.

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

    body := *datadog.NewSLOCorrectionCreateRequest() // SLOCorrectionCreateRequest | Create an SLO Correction

    configuration := datadog.NewConfiguration()
    configuration.SetUnstableOperationEnabled("CreateSLOCorrection", true)

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceLevelObjectiveCorrectionsApi.CreateSLOCorrection(ctx, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceLevelObjectiveCorrectionsApi.CreateSLOCorrection`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSLOCorrection`: SLOCorrectionResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from ServiceLevelObjectiveCorrectionsApi.CreateSLOCorrection:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**body** | [**SLOCorrectionCreateRequest**](SLOCorrectionCreateRequest.md) | Create an SLO Correction | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**SLOCorrectionResponse**](SLOCorrectionResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSLOCorrection

> DeleteSLOCorrection(ctx, sloCorrectionId)

Permanently delete the specified SLO correction object.

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

    sloCorrectionId := "sloCorrectionId_example" // string | The ID of the SLO correction object.

    configuration := datadog.NewConfiguration()
    configuration.SetUnstableOperationEnabled("DeleteSLOCorrection", true)

    apiClient := datadog.NewAPIClient(configuration)
    r, err := apiClient.ServiceLevelObjectiveCorrectionsApi.DeleteSLOCorrection(ctx, sloCorrectionId)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceLevelObjectiveCorrectionsApi.DeleteSLOCorrection`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**sloCorrectionId** | **string** | The ID of the SLO correction object. | 


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


## GetSLOCorrection

> SLOCorrectionResponse GetSLOCorrection(ctx, sloCorrectionId)

Get an SLO correction.

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

    sloCorrectionId := "sloCorrectionId_example" // string | The ID of the SLO correction object.

    configuration := datadog.NewConfiguration()
    configuration.SetUnstableOperationEnabled("GetSLOCorrection", true)

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceLevelObjectiveCorrectionsApi.GetSLOCorrection(ctx, sloCorrectionId)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceLevelObjectiveCorrectionsApi.GetSLOCorrection`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSLOCorrection`: SLOCorrectionResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from ServiceLevelObjectiveCorrectionsApi.GetSLOCorrection:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**sloCorrectionId** | **string** | The ID of the SLO correction object. | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**SLOCorrectionResponse**](SLOCorrectionResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSLOCorrection

> SLOCorrectionListResponse ListSLOCorrection(ctx)

Get all Service Level Objective corrections.

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
    configuration.SetUnstableOperationEnabled("ListSLOCorrection", true)

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceLevelObjectiveCorrectionsApi.ListSLOCorrection(ctx)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceLevelObjectiveCorrectionsApi.ListSLOCorrection`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSLOCorrection`: SLOCorrectionListResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from ServiceLevelObjectiveCorrectionsApi.ListSLOCorrection:\n%s\n", responseContent)
}
```

### Required Parameters

This endpoint does not need any parameter.


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**SLOCorrectionListResponse**](SLOCorrectionListResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSLOCorrection

> SLOCorrectionResponse UpdateSLOCorrection(ctx, sloCorrectionId, body)

Update the specified SLO correction object object.

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

    sloCorrectionId := "sloCorrectionId_example" // string | The ID of the SLO correction object.
    body := *datadog.NewSLOCorrectionUpdateRequest() // SLOCorrectionUpdateRequest | The edited SLO correction object.

    configuration := datadog.NewConfiguration()
    configuration.SetUnstableOperationEnabled("UpdateSLOCorrection", true)

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceLevelObjectiveCorrectionsApi.UpdateSLOCorrection(ctx, sloCorrectionId, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceLevelObjectiveCorrectionsApi.UpdateSLOCorrection`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSLOCorrection`: SLOCorrectionResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from ServiceLevelObjectiveCorrectionsApi.UpdateSLOCorrection:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**sloCorrectionId** | **string** | The ID of the SLO correction object. |  |
**body** | [**SLOCorrectionUpdateRequest**](SLOCorrectionUpdateRequest.md) | The edited SLO correction object. | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**SLOCorrectionResponse**](SLOCorrectionResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

