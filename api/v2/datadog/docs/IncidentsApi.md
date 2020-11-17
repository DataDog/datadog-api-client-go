# \IncidentsApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIncident**](IncidentsApi.md#CreateIncident) | **Post** /api/v2/incidents | Create an incident
[**DeleteIncident**](IncidentsApi.md#DeleteIncident) | **Delete** /api/v2/incidents/{incident_id} | Delete an existing incident
[**GetIncident**](IncidentsApi.md#GetIncident) | **Get** /api/v2/incidents/{incident_id} | Get the details of an incident
[**ListIncidents**](IncidentsApi.md#ListIncidents) | **Get** /api/v2/incidents | Get a list of incidents
[**UpdateIncident**](IncidentsApi.md#UpdateIncident) | **Patch** /api/v2/incidents/{incident_id} | Update an existing incident



## CreateIncident

> IncidentResponse CreateIncident(ctx).Body(body).Execute()

Create an incident



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

    body := *datadog.NewIncidentCreateRequest(*datadog.NewIncidentCreateData(*datadog.NewIncidentCreateAttributes(false, "A test incident title"), datadog.IncidentType("incidents"))) // IncidentCreateRequest | Incident payload.

    configuration := datadog.NewConfiguration()
    configuration.SetUnstableOperationEnabled("CreateIncident", true)

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.IncidentsApi.CreateIncident(ctx).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.CreateIncident``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIncident`: IncidentResponse
    fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.CreateIncident`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIncidentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**IncidentCreateRequest**](IncidentCreateRequest.md) | Incident payload. | 

### Return type

[**IncidentResponse**](IncidentResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIncident

> DeleteIncident(ctx, incidentId).Execute()

Delete an existing incident



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

    incidentId := "incidentId_example" // string | The UUID the incident.

    configuration := datadog.NewConfiguration()
    configuration.SetUnstableOperationEnabled("DeleteIncident", true)

    api_client := datadog.NewAPIClient(configuration)
    r, err := api_client.IncidentsApi.DeleteIncident(ctx, incidentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.DeleteIncident``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**incidentId** | **string** | The UUID the incident. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIncidentRequest struct via the builder pattern


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


## GetIncident

> IncidentResponse GetIncident(ctx, incidentId).Include(include).Execute()

Get the details of an incident



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

    incidentId := "incidentId_example" // string | The UUID the incident.
    include := []string{"Include_example"} // []string | Specifies which types of related objects should be included in the response. (optional)

    configuration := datadog.NewConfiguration()
    configuration.SetUnstableOperationEnabled("GetIncident", true)

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.IncidentsApi.GetIncident(ctx, incidentId).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.GetIncident``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIncident`: IncidentResponse
    fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.GetIncident`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**incidentId** | **string** | The UUID the incident. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIncidentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Specifies which types of related objects should be included in the response. | 

### Return type

[**IncidentResponse**](IncidentResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIncidents

> IncidentsResponse ListIncidents(ctx).Include(include).PageSize(pageSize).PageOffset(pageOffset).Execute()

Get a list of incidents



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

    include := []string{"Include_example"} // []string | Specifies which types of related objects should be included in the response. (optional)
    pageSize := int64(789) // int64 | Size for a given page. (optional) (default to 10)
    pageOffset := int64(789) // int64 | Specific offset to use as the beginning of the returned page. (optional) (default to 0)

    configuration := datadog.NewConfiguration()
    configuration.SetUnstableOperationEnabled("ListIncidents", true)

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.IncidentsApi.ListIncidents(ctx).Include(include).PageSize(pageSize).PageOffset(pageOffset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.ListIncidents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIncidents`: IncidentsResponse
    fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.ListIncidents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListIncidentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | Specifies which types of related objects should be included in the response. | 
 **pageSize** | **int64** | Size for a given page. | [default to 10]
 **pageOffset** | **int64** | Specific offset to use as the beginning of the returned page. | [default to 0]

### Return type

[**IncidentsResponse**](IncidentsResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIncident

> IncidentResponse UpdateIncident(ctx, incidentId).Body(body).Execute()

Update an existing incident



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

    incidentId := "incidentId_example" // string | The UUID the incident.
    body := *datadog.NewIncidentUpdateRequest(*datadog.NewIncidentUpdateData("00000000-0000-0000-0000-000000000000", datadog.IncidentType("incidents"))) // IncidentUpdateRequest | Incident Payload.

    configuration := datadog.NewConfiguration()
    configuration.SetUnstableOperationEnabled("UpdateIncident", true)

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.IncidentsApi.UpdateIncident(ctx, incidentId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.UpdateIncident``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIncident`: IncidentResponse
    fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.UpdateIncident`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**incidentId** | **string** | The UUID the incident. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIncidentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**IncidentUpdateRequest**](IncidentUpdateRequest.md) | Incident Payload. | 

### Return type

[**IncidentResponse**](IncidentResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

