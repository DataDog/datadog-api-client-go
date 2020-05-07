# \ServiceLevelObjectivesApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckCanDeleteSLO**](ServiceLevelObjectivesApi.md#CheckCanDeleteSLO) | **Get** /api/v1/slo/can_delete | Check if SLOs can be safely deleted
[**CreateSLO**](ServiceLevelObjectivesApi.md#CreateSLO) | **Post** /api/v1/slo | Create a SLO object
[**DeleteSLO**](ServiceLevelObjectivesApi.md#DeleteSLO) | **Delete** /api/v1/slo/{slo_id} | Delete a SLO
[**DeleteSLOTimeframeInBulk**](ServiceLevelObjectivesApi.md#DeleteSLOTimeframeInBulk) | **Post** /api/v1/slo/bulk_delete | Bulk Delete SLO Timeframes
[**GetSLO**](ServiceLevelObjectivesApi.md#GetSLO) | **Get** /api/v1/slo/{slo_id} | Get a SLO&#39;s details
[**GetSLOHistory**](ServiceLevelObjectivesApi.md#GetSLOHistory) | **Get** /api/v1/slo/{slo_id}/history | Get an SLO&#39;s history
[**ListSLOs**](ServiceLevelObjectivesApi.md#ListSLOs) | **Get** /api/v1/slo | Search SLOs
[**UpdateSLO**](ServiceLevelObjectivesApi.md#UpdateSLO) | **Put** /api/v1/slo/{slo_id} | Update a SLO



## CheckCanDeleteSLO

> CheckCanDeleteSLOResponse CheckCanDeleteSLO(ctx).Ids(ids).Execute()

Check if SLOs can be safely deleted



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
    ids := "ids_example" // string | A comma separated list of the IDs of the service level objectives objects.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceLevelObjectivesApi.CheckCanDeleteSLO(context.Background(), ids).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceLevelObjectivesApi.CheckCanDeleteSLO``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckCanDeleteSLO`: CheckCanDeleteSLOResponse
    fmt.Fprintf(os.Stdout, "Response from `ServiceLevelObjectivesApi.CheckCanDeleteSLO`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckCanDeleteSLORequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **string** | A comma separated list of the IDs of the service level objectives objects. | 

### Return type

[**CheckCanDeleteSLOResponse**](CheckCanDeleteSLOResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSLO

> SLOListResponse CreateSLO(ctx).Body(body).Execute()

Create a SLO object



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
    body := openapiclient.ServiceLevelObjectiveRequest{CreatedAt: int64(123), Creator: openapiclient.Creator{Email: "Email_example", Handle: "Handle_example", Name: "Name_example"}, Description: "Description_example", Groups: []string{"Groups_example"), Id: "Id_example", ModifiedAt: int64(123), MonitorIds: []int64{int64(123)), Name: "Name_example", Query: openapiclient.ServiceLevelObjective_query{Denominator: "Denominator_example", Numerator: "Numerator_example"}, Tags: []string{"Tags_example"), Thresholds: []SLOThreshold{openapiclient.SLOThreshold{Target: 123, TargetDisplay: "TargetDisplay_example", Timeframe: openapiclient.SLOTimeframe{}, Warning: 123, WarningDisplay: "WarningDisplay_example"}), Type: openapiclient.SLOType{}} // ServiceLevelObjectiveRequest | Service level objective request object.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceLevelObjectivesApi.CreateSLO(context.Background(), body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceLevelObjectivesApi.CreateSLO``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSLO`: SLOListResponse
    fmt.Fprintf(os.Stdout, "Response from `ServiceLevelObjectivesApi.CreateSLO`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSLORequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ServiceLevelObjectiveRequest**](ServiceLevelObjectiveRequest.md) | Service level objective request object. | 

### Return type

[**SLOListResponse**](SLOListResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSLO

> SLODeleteResponse DeleteSLO(ctx, sloId).Force(force).Execute()

Delete a SLO



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
    sloId := "sloId_example" // string | The ID of the service level objective.
    force := "force_example" // string | Delete the monitor even if it's referenced by other resources (e.g. SLO, composite monitor). (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceLevelObjectivesApi.DeleteSLO(context.Background(), sloId).Force(force).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceLevelObjectivesApi.DeleteSLO``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSLO`: SLODeleteResponse
    fmt.Fprintf(os.Stdout, "Response from `ServiceLevelObjectivesApi.DeleteSLO`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sloId** | **string** | The ID of the service level objective. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSLORequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **string** | Delete the monitor even if it&#39;s referenced by other resources (e.g. SLO, composite monitor). | 

### Return type

[**SLODeleteResponse**](SLODeleteResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSLOTimeframeInBulk

> SLOBulkDeleteResponse DeleteSLOTimeframeInBulk(ctx).Body(body).Execute()

Bulk Delete SLO Timeframes



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
    body := map[string]string{ "Key" = "Value" } // map[string][]SLOTimeframe | Delete multiple service level objective objects request body.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceLevelObjectivesApi.DeleteSLOTimeframeInBulk(context.Background(), body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceLevelObjectivesApi.DeleteSLOTimeframeInBulk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSLOTimeframeInBulk`: SLOBulkDeleteResponse
    fmt.Fprintf(os.Stdout, "Response from `ServiceLevelObjectivesApi.DeleteSLOTimeframeInBulk`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSLOTimeframeInBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**map[string][]SLOTimeframe**](array.md) | Delete multiple service level objective objects request body. | 

### Return type

[**SLOBulkDeleteResponse**](SLOBulkDeleteResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSLO

> SLOResponse GetSLO(ctx, sloId).Execute()

Get a SLO's details



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
    sloId := "sloId_example" // string | The ID of the service level objective object.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceLevelObjectivesApi.GetSLO(context.Background(), sloId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceLevelObjectivesApi.GetSLO``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSLO`: SLOResponse
    fmt.Fprintf(os.Stdout, "Response from `ServiceLevelObjectivesApi.GetSLO`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sloId** | **string** | The ID of the service level objective object. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSLORequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SLOResponse**](SLOResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSLOHistory

> SLOHistoryResponse GetSLOHistory(ctx, sloId).FromTs(fromTs).ToTs(toTs).Execute()

Get an SLO's history



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
    sloId := "sloId_example" // string | The ID of the service level objective object.
    fromTs := 987 // int64 | The `from` timestamp for the query window in epoch seconds.
    toTs := 987 // int64 | The `to` timestamp for the query window in epoch seconds.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceLevelObjectivesApi.GetSLOHistory(context.Background(), sloId, fromTs, toTs).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceLevelObjectivesApi.GetSLOHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSLOHistory`: SLOHistoryResponse
    fmt.Fprintf(os.Stdout, "Response from `ServiceLevelObjectivesApi.GetSLOHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sloId** | **string** | The ID of the service level objective object. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSLOHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fromTs** | **int64** | The &#x60;from&#x60; timestamp for the query window in epoch seconds. | 
 **toTs** | **int64** | The &#x60;to&#x60; timestamp for the query window in epoch seconds. | 

### Return type

[**SLOHistoryResponse**](SLOHistoryResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSLOs

> SLOListResponse ListSLOs(ctx).Ids(ids).Execute()

Search SLOs



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
    ids := "ids_example" // string | A comma separated list of the IDs of the service level objectives objects.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceLevelObjectivesApi.ListSLOs(context.Background(), ids).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceLevelObjectivesApi.ListSLOs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSLOs`: SLOListResponse
    fmt.Fprintf(os.Stdout, "Response from `ServiceLevelObjectivesApi.ListSLOs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSLOsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **string** | A comma separated list of the IDs of the service level objectives objects. | 

### Return type

[**SLOListResponse**](SLOListResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSLO

> SLOListResponse UpdateSLO(ctx, sloId).Body(body).Execute()

Update a SLO



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
    sloId := "sloId_example" // string | The ID of the service level objective object.
    body := openapiclient.ServiceLevelObjective{CreatedAt: int64(123), Creator: openapiclient.Creator{Email: "Email_example", Handle: "Handle_example", Name: "Name_example"}, Description: "Description_example", Groups: []string{"Groups_example"), Id: "Id_example", ModifiedAt: int64(123), MonitorIds: []int64{int64(123)), MonitorTags: []string{"MonitorTags_example"), Name: "Name_example", Query: openapiclient.ServiceLevelObjective_query{Denominator: "Denominator_example", Numerator: "Numerator_example"}, Tags: []string{"Tags_example"), Thresholds: []SLOThreshold{openapiclient.SLOThreshold{Target: 123, TargetDisplay: "TargetDisplay_example", Timeframe: openapiclient.SLOTimeframe{}, Warning: 123, WarningDisplay: "WarningDisplay_example"}), Type: openapiclient.SLOType{}} // ServiceLevelObjective | The edited service level objective request object.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceLevelObjectivesApi.UpdateSLO(context.Background(), sloId, body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceLevelObjectivesApi.UpdateSLO``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSLO`: SLOListResponse
    fmt.Fprintf(os.Stdout, "Response from `ServiceLevelObjectivesApi.UpdateSLO`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sloId** | **string** | The ID of the service level objective object. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSLORequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ServiceLevelObjective**](ServiceLevelObjective.md) | The edited service level objective request object. | 

### Return type

[**SLOListResponse**](SLOListResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

