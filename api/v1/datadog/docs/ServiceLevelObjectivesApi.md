# ServiceLevelObjectivesApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------ | ------------ | ------------
[**CheckCanDeleteSLO**](ServiceLevelObjectivesApi.md#CheckCanDeleteSLO) | **Get** /api/v1/slo/can_delete | Check if SLOs can be safely deleted
[**CreateSLO**](ServiceLevelObjectivesApi.md#CreateSLO) | **Post** /api/v1/slo | Create an SLO object
[**DeleteSLO**](ServiceLevelObjectivesApi.md#DeleteSLO) | **Delete** /api/v1/slo/{slo_id} | Delete an SLO
[**DeleteSLOTimeframeInBulk**](ServiceLevelObjectivesApi.md#DeleteSLOTimeframeInBulk) | **Post** /api/v1/slo/bulk_delete | Bulk Delete SLO Timeframes
[**GetSLO**](ServiceLevelObjectivesApi.md#GetSLO) | **Get** /api/v1/slo/{slo_id} | Get an SLO&#39;s details
[**GetSLOHistory**](ServiceLevelObjectivesApi.md#GetSLOHistory) | **Get** /api/v1/slo/{slo_id}/history | Get an SLO&#39;s history
[**ListSLOs**](ServiceLevelObjectivesApi.md#ListSLOs) | **Get** /api/v1/slo | Get all SLOs
[**UpdateSLO**](ServiceLevelObjectivesApi.md#UpdateSLO) | **Put** /api/v1/slo/{slo_id} | Update an SLO



## CheckCanDeleteSLO

> CheckCanDeleteSLOResponse CheckCanDeleteSLO(ctx, ids)

Check if an SLO can be safely deleted. For example,
assure an SLO can be deleted without disrupting a dashboard.

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

    ids := "id1, id2, id3" // string | A comma separated list of the IDs of the service level objectives objects.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceLevelObjectivesApi.CheckCanDeleteSLO(ctx, ids)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceLevelObjectivesApi.CheckCanDeleteSLO`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckCanDeleteSLO`: CheckCanDeleteSLOResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from ServiceLevelObjectivesApi.CheckCanDeleteSLO:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**ids** | **string** | A comma separated list of the IDs of the service level objectives objects. | 


### Optional Parameters

This endpoint does not have optional parameters.


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

> SLOListResponse CreateSLO(ctx, body)

Create a service level objective object.

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

    body := *datadog.NewServiceLevelObjectiveRequest("Custom Metric SLO", []datadog.SLOThreshold{*datadog.NewSLOThreshold(float64(99.9), datadog.SLOTimeframe("7d"))}, datadog.SLOType("metric")) // ServiceLevelObjectiveRequest | Service level objective request object.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceLevelObjectivesApi.CreateSLO(ctx, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceLevelObjectivesApi.CreateSLO`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSLO`: SLOListResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from ServiceLevelObjectivesApi.CreateSLO:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**body** | [**ServiceLevelObjectiveRequest**](ServiceLevelObjectiveRequest.md) | Service level objective request object. | 


### Optional Parameters

This endpoint does not have optional parameters.


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

> SLODeleteResponse DeleteSLO(ctx, sloId, datadog.DeleteSLOOptionalParameters{})

Permanently delete the specified service level objective object.

If an SLO is used in a dashboard, the `DELETE /v1/slo/` endpoint returns
a 409 conflict error because the SLO is referenced in a dashboard.

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

    sloId := "sloId_example" // string | The ID of the service level objective.
    force := "force_example" // string | Delete the monitor even if it's referenced by other resources (e.g. SLO, composite monitor). (optional)
    optionalParams := datadog.DeleteSLOOptionalParameters{
        Force: &force,
    }

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceLevelObjectivesApi.DeleteSLO(ctx, sloId, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceLevelObjectivesApi.DeleteSLO`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSLO`: SLODeleteResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from ServiceLevelObjectivesApi.DeleteSLO:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**sloId** | **string** | The ID of the service level objective. | 


### Optional Parameters


Other parameters are passed through a pointer to a DeleteSLOOptionalParameters struct.


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
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

> SLOBulkDeleteResponse DeleteSLOTimeframeInBulk(ctx, body)

Delete (or partially delete) multiple service level objective objects.

This endpoint facilitates deletion of one or more thresholds for one or more
service level objective objects. If all thresholds are deleted, the service level
objective object is deleted as well.

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

    body := map[string][]datadog.SLOTimeframe{"key": []datadog.SLOTimeframe{datadog.SLOTimeframe("7d")}} // map[string][]SLOTimeframe | Delete multiple service level objective objects request body.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceLevelObjectivesApi.DeleteSLOTimeframeInBulk(ctx, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceLevelObjectivesApi.DeleteSLOTimeframeInBulk`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSLOTimeframeInBulk`: SLOBulkDeleteResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from ServiceLevelObjectivesApi.DeleteSLOTimeframeInBulk:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**body** | **map[string][]SLOTimeframe** | Delete multiple service level objective objects request body. | 


### Optional Parameters

This endpoint does not have optional parameters.


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

> SLOResponse GetSLO(ctx, sloId, datadog.GetSLOOptionalParameters{})

Get a service level objective object.

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

    sloId := "sloId_example" // string | The ID of the service level objective object.
    withConfiguredAlertIds := true // bool | Get the IDs of SLO monitors that reference this SLO. (optional)
    optionalParams := datadog.GetSLOOptionalParameters{
        WithConfiguredAlertIds: &withConfiguredAlertIds,
    }

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceLevelObjectivesApi.GetSLO(ctx, sloId, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceLevelObjectivesApi.GetSLO`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSLO`: SLOResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from ServiceLevelObjectivesApi.GetSLO:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**sloId** | **string** | The ID of the service level objective object. | 


### Optional Parameters


Other parameters are passed through a pointer to a GetSLOOptionalParameters struct.


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**withConfiguredAlertIds** | **bool** | Get the IDs of SLO monitors that reference this SLO. | 

### Return type

[**SLOResponse**](SLOResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSLOHistory

> SLOHistoryResponse GetSLOHistory(ctx, sloId, fromTs, toTs, datadog.GetSLOHistoryOptionalParameters{})

Get a specific SLO’s history, regardless of its SLO type.

The detailed history data is structured according to the source data type.
For example, metric data is included for event SLOs that use
the metric source, and monitor SLO types include the monitor transition history.

**Note:** There are different response formats for event based and time based SLOs.
Examples of both are shown.

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

    sloId := "sloId_example" // string | The ID of the service level objective object.
    fromTs := int64(789) // int64 | The `from` timestamp for the query window in epoch seconds.
    toTs := int64(789) // int64 | The `to` timestamp for the query window in epoch seconds.
    target := float64(1.2) // float64 | The SLO target. If `target` is passed in, the response will include the remaining error budget and a timeframe value of `custom`. (optional)
    optionalParams := datadog.GetSLOHistoryOptionalParameters{
        Target: &target,
    }

    configuration := datadog.NewConfiguration()
    configuration.SetUnstableOperationEnabled("GetSLOHistory", true)

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceLevelObjectivesApi.GetSLOHistory(ctx, sloId, fromTs, toTs, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceLevelObjectivesApi.GetSLOHistory`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSLOHistory`: SLOHistoryResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from ServiceLevelObjectivesApi.GetSLOHistory:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**sloId** | **string** | The ID of the service level objective object. |  |
**fromTs** | **int64** | The &#x60;from&#x60; timestamp for the query window in epoch seconds. |  |
**toTs** | **int64** | The &#x60;to&#x60; timestamp for the query window in epoch seconds. | 


### Optional Parameters


Other parameters are passed through a pointer to a GetSLOHistoryOptionalParameters struct.


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**target** | **float64** | The SLO target. If &#x60;target&#x60; is passed in, the response will include the remaining error budget and a timeframe value of &#x60;custom&#x60;. | 

### Return type

[**SLOHistoryResponse**](SLOHistoryResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSLOs

> SLOListResponse ListSLOs(ctx, datadog.ListSLOsOptionalParameters{})

Get a list of service level objective objects for your organization.

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

    ids := "id1, id2, id3" // string | A comma separated list of the IDs of the service level objectives objects. (optional)
    query := "monitor" // string | The query string to filter results based on SLO names. (optional)
    tagsQuery := "env:prod" // string | The query string to filter results based on a single SLO tag. (optional)
    metricsQuery := "aws.elb.request_count" // string | The query string to filter results based on SLO numerator and denominator. (optional)
    limit := int64(789) // int64 | The number of SLOs to return in the response. (optional)
    offset := int64(789) // int64 | The specific offset to use as the beginning of the returned response. (optional)
    optionalParams := datadog.ListSLOsOptionalParameters{
        Ids: &ids,
        Query: &query,
        TagsQuery: &tagsQuery,
        MetricsQuery: &metricsQuery,
        Limit: &limit,
        Offset: &offset,
    }

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceLevelObjectivesApi.ListSLOs(ctx, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceLevelObjectivesApi.ListSLOs`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSLOs`: SLOListResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from ServiceLevelObjectivesApi.ListSLOs:\n%s\n", responseContent)
}
```

### Required Parameters




### Optional Parameters


Other parameters are passed through a pointer to a ListSLOsOptionalParameters struct.


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ids** | **string** | A comma separated list of the IDs of the service level objectives objects. | 
**query** | **string** | The query string to filter results based on SLO names. | 
**tagsQuery** | **string** | The query string to filter results based on a single SLO tag. | 
**metricsQuery** | **string** | The query string to filter results based on SLO numerator and denominator. | 
**limit** | **int64** | The number of SLOs to return in the response. | 
**offset** | **int64** | The specific offset to use as the beginning of the returned response. | 

### Return type

[**SLOListResponse**](SLOListResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSLO

> SLOListResponse UpdateSLO(ctx, sloId, body)

Update the specified service level objective object.

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

    sloId := "sloId_example" // string | The ID of the service level objective object.
    body := *datadog.NewServiceLevelObjective("Custom Metric SLO", []datadog.SLOThreshold{*datadog.NewSLOThreshold(float64(99.9), datadog.SLOTimeframe("7d"))}, datadog.SLOType("metric")) // ServiceLevelObjective | The edited service level objective request object.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceLevelObjectivesApi.UpdateSLO(ctx, sloId, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceLevelObjectivesApi.UpdateSLO`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSLO`: SLOListResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from ServiceLevelObjectivesApi.UpdateSLO:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**sloId** | **string** | The ID of the service level objective object. |  |
**body** | [**ServiceLevelObjective**](ServiceLevelObjective.md) | The edited service level objective request object. | 


### Optional Parameters

This endpoint does not have optional parameters.


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

