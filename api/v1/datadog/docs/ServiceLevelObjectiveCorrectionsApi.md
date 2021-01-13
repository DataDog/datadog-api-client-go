# \ServiceLevelObjectiveCorrectionsApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSLOCorrection**](ServiceLevelObjectiveCorrectionsApi.md#CreateSLOCorrection) | **Post** /api/v1/slo/correction | Create an SLO correction
[**DeleteSLOCorrection**](ServiceLevelObjectiveCorrectionsApi.md#DeleteSLOCorrection) | **Delete** /api/v1/slo/correction/{slo_correction_id} | Delete an SLO Correction
[**GetSLOCorrection**](ServiceLevelObjectiveCorrectionsApi.md#GetSLOCorrection) | **Get** /api/v1/slo/correction/{slo_correction_id} | Get an SLO correction for an SLO
[**ListSLOCorrection**](ServiceLevelObjectiveCorrectionsApi.md#ListSLOCorrection) | **Get** /api/v1/slo/correction | Get all SLO corrections
[**UpdateSLOCorrection**](ServiceLevelObjectiveCorrectionsApi.md#UpdateSLOCorrection) | **Patch** /api/v1/slo/correction/{slo_correction_id} | Update an SLO Correction



## CreateSLOCorrection

> SLOCorrectionResponse CreateSLOCorrection(ctx).Body(body).Execute()

Create an SLO correction



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

    body := *datadog.NewSLOCorrectionCreateRequest() // SLOCorrectionCreateRequest | Create an SLO Correction

    configuration := datadog.NewConfiguration()
    configuration.SetUnstableOperationEnabled("CreateSLOCorrection", true)

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceLevelObjectiveCorrectionsApi.CreateSLOCorrection(ctx).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceLevelObjectiveCorrectionsApi.CreateSLOCorrection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSLOCorrection`: SLOCorrectionResponse
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from ServiceLevelObjectiveCorrectionsApi.CreateSLOCorrection:\n%s\n", response_content)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSLOCorrectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SLOCorrectionCreateRequest**](SLOCorrectionCreateRequest.md) | Create an SLO Correction | 

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

> DeleteSLOCorrection(ctx, sloCorrectionId).Execute()

Delete an SLO Correction



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

    sloCorrectionId := "sloCorrectionId_example" // string | The ID of the SLO correction object

    configuration := datadog.NewConfiguration()
    configuration.SetUnstableOperationEnabled("DeleteSLOCorrection", true)

    api_client := datadog.NewAPIClient(configuration)
    r, err := api_client.ServiceLevelObjectiveCorrectionsApi.DeleteSLOCorrection(ctx, sloCorrectionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceLevelObjectiveCorrectionsApi.DeleteSLOCorrection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sloCorrectionId** | **string** | The ID of the SLO correction object | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSLOCorrectionRequest struct via the builder pattern


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


## GetSLOCorrection

> SLOCorrectionResponse GetSLOCorrection(ctx, sloCorrectionId).Execute()

Get an SLO correction for an SLO



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

    sloCorrectionId := "sloCorrectionId_example" // string | The ID of the SLO correction object

    configuration := datadog.NewConfiguration()
    configuration.SetUnstableOperationEnabled("GetSLOCorrection", true)

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceLevelObjectiveCorrectionsApi.GetSLOCorrection(ctx, sloCorrectionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceLevelObjectiveCorrectionsApi.GetSLOCorrection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSLOCorrection`: SLOCorrectionResponse
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from ServiceLevelObjectiveCorrectionsApi.GetSLOCorrection:\n%s\n", response_content)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sloCorrectionId** | **string** | The ID of the SLO correction object | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSLOCorrectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> SLOCorrectionListResponse ListSLOCorrection(ctx).Execute()

Get all SLO corrections



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
    configuration.SetUnstableOperationEnabled("ListSLOCorrection", true)

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceLevelObjectiveCorrectionsApi.ListSLOCorrection(ctx).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceLevelObjectiveCorrectionsApi.ListSLOCorrection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSLOCorrection`: SLOCorrectionListResponse
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from ServiceLevelObjectiveCorrectionsApi.ListSLOCorrection:\n%s\n", response_content)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListSLOCorrectionRequest struct via the builder pattern


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

> SLOCorrectionResponse UpdateSLOCorrection(ctx, sloCorrectionId).Body(body).Execute()

Update an SLO Correction



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

    sloCorrectionId := "sloCorrectionId_example" // string | The ID of the SLO correction object
    body := *datadog.NewSLOCorrectionUpdateRequest() // SLOCorrectionUpdateRequest | The edited SLO correction object.

    configuration := datadog.NewConfiguration()
    configuration.SetUnstableOperationEnabled("UpdateSLOCorrection", true)

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceLevelObjectiveCorrectionsApi.UpdateSLOCorrection(ctx, sloCorrectionId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceLevelObjectiveCorrectionsApi.UpdateSLOCorrection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSLOCorrection`: SLOCorrectionResponse
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from ServiceLevelObjectiveCorrectionsApi.UpdateSLOCorrection:\n%s\n", response_content)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sloCorrectionId** | **string** | The ID of the SLO correction object | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSLOCorrectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**SLOCorrectionUpdateRequest**](SLOCorrectionUpdateRequest.md) | The edited SLO correction object. | 

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

