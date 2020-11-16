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

    ids := "id1, id2, id3" // string | A comma separated list of the IDs of the service level objectives objects.

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceLevelObjectivesApi.CheckCanDeleteSLO(ctx).Ids(ids).Execute()
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

    body := *datadog.NewServiceLevelObjectiveRequest("Name_example", []datadog.SLOThreshold{*datadog.NewSLOThreshold(float64(0.0), datadog.SLOTimeframe("7d"))}, datadog.SLOType("metric")) // ServiceLevelObjectiveRequest | Service level objective request object.

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceLevelObjectivesApi.CreateSLO(ctx).Body(body).Execute()
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

    sloId := "sloId_example" // string | The ID of the service level objective.
    force := "force_example" // string | Delete the monitor even if it's referenced by other resources (e.g. SLO, composite monitor). (optional)

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceLevelObjectivesApi.DeleteSLO(ctx, sloId).Force(force).Execute()
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

    body := map[string][]datadog.SLOTimeframe{"key": []datadog.SLOTimeframe{datadog.SLOTimeframe("7d")}} // map[string][]SLOTimeframe | Delete multiple service level objective objects request body.

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceLevelObjectivesApi.DeleteSLOTimeframeInBulk(ctx).Body(body).Execute()
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

    sloId := "sloId_example" // string | The ID of the service level objective object.

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceLevelObjectivesApi.GetSLO(ctx, sloId).Execute()
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

    sloId := "sloId_example" // string | The ID of the service level objective object.
    fromTs := int64(789) // int64 | The `from` timestamp for the query window in epoch seconds.
    toTs := int64(789) // int64 | The `to` timestamp for the query window in epoch seconds.

    configuration := datadog.NewConfiguration()
    configuration.SetUnstableOperationEnabled("GetSLOHistory", true)

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceLevelObjectivesApi.GetSLOHistory(ctx, sloId).FromTs(fromTs).ToTs(toTs).Execute()
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

    ids := "id1, id2, id3" // string | A comma separated list of the IDs of the service level objectives objects.

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceLevelObjectivesApi.ListSLOs(ctx).Ids(ids).Execute()
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

    sloId := "sloId_example" // string | The ID of the service level objective object.
    body := *datadog.NewServiceLevelObjective("Name_example", []datadog.SLOThreshold{*datadog.NewSLOThreshold(float64(0.0), datadog.SLOTimeframe("7d"))}, datadog.SLOType("metric")) // ServiceLevelObjective | The edited service level objective request object.

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceLevelObjectivesApi.UpdateSLO(ctx, sloId).Body(body).Execute()
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

