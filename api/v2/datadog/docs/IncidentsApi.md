# IncidentsApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------ | ------------ | ------------
[**CreateIncident**](IncidentsApi.md#CreateIncident) | **Post** /api/v2/incidents | Create an incident
[**DeleteIncident**](IncidentsApi.md#DeleteIncident) | **Delete** /api/v2/incidents/{incident_id} | Delete an existing incident
[**GetIncident**](IncidentsApi.md#GetIncident) | **Get** /api/v2/incidents/{incident_id} | Get the details of an incident
[**ListIncidents**](IncidentsApi.md#ListIncidents) | **Get** /api/v2/incidents | Get a list of incidents
[**UpdateIncident**](IncidentsApi.md#UpdateIncident) | **Patch** /api/v2/incidents/{incident_id} | Update an existing incident



## CreateIncident

> IncidentResponse CreateIncident(ctx, body)

Create an incident.

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

    body := *datadog.NewIncidentCreateRequest(*datadog.NewIncidentCreateData(*datadog.NewIncidentCreateAttributes(false, "A test incident title"), datadog.IncidentType("incidents"))) // IncidentCreateRequest | Incident payload.

    configuration := datadog.NewConfiguration()
    configuration.SetUnstableOperationEnabled("CreateIncident", true)

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.IncidentsApi.CreateIncident(ctx, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.CreateIncident`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIncident`: IncidentResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from IncidentsApi.CreateIncident:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**body** | [**IncidentCreateRequest**](IncidentCreateRequest.md) | Incident payload. | 


### Optional Parameters

This endpoint does not have optional parameters.


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

> DeleteIncident(ctx, incidentId)

Deletes an existing incident from the users organization.

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

    incidentId := "incidentId_example" // string | The UUID the incident.

    configuration := datadog.NewConfiguration()
    configuration.SetUnstableOperationEnabled("DeleteIncident", true)

    apiClient := datadog.NewAPIClient(configuration)
    r, err := apiClient.IncidentsApi.DeleteIncident(ctx, incidentId)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.DeleteIncident`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**incidentId** | **string** | The UUID the incident. | 


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


## GetIncident

> IncidentResponse GetIncident(ctx, incidentId, datadog.GetIncidentOptionalParameters{})

Get the details of an incident by `incident_id`.

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

    incidentId := "incidentId_example" // string | The UUID the incident.
    include := []datadog.IncidentRelatedObject{datadog.IncidentRelatedObject("users")} // []IncidentRelatedObject | Specifies which types of related objects should be included in the response. (optional)
    optionalParams := datadog.GetIncidentOptionalParameters{
        Include: &include,
    }

    configuration := datadog.NewConfiguration()
    configuration.SetUnstableOperationEnabled("GetIncident", true)

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.IncidentsApi.GetIncident(ctx, incidentId, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.GetIncident`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIncident`: IncidentResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from IncidentsApi.GetIncident:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**incidentId** | **string** | The UUID the incident. | 


### Optional Parameters


Other parameters are passed through a pointer to a GetIncidentOptionalParameters struct.


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**include** | [**[]IncidentRelatedObject**](IncidentRelatedObject.md) | Specifies which types of related objects should be included in the response. | 

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

> IncidentsResponse ListIncidents(ctx, datadog.ListIncidentsOptionalParameters{})

Get all incidents for the user's organization.

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

    include := []datadog.IncidentRelatedObject{datadog.IncidentRelatedObject("users")} // []IncidentRelatedObject | Specifies which types of related objects should be included in the response. (optional)
    pageSize := int64(789) // int64 | Size for a given page. (optional) (default to 10)
    pageOffset := int64(789) // int64 | Specific offset to use as the beginning of the returned page. (optional) (default to 0)
    optionalParams := datadog.ListIncidentsOptionalParameters{
        Include: &include,
        PageSize: &pageSize,
        PageOffset: &pageOffset,
    }

    configuration := datadog.NewConfiguration()
    configuration.SetUnstableOperationEnabled("ListIncidents", true)

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.IncidentsApi.ListIncidents(ctx, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.ListIncidents`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIncidents`: IncidentsResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from IncidentsApi.ListIncidents:\n%s\n", responseContent)
}
```

### Required Parameters




### Optional Parameters


Other parameters are passed through a pointer to a ListIncidentsOptionalParameters struct.


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**include** | [**[]IncidentRelatedObject**](IncidentRelatedObject.md) | Specifies which types of related objects should be included in the response. | 
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

> IncidentResponse UpdateIncident(ctx, incidentId, body)

Updates an incident. Provide only the attributes that should be updated as this request is a partial update.

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

    incidentId := "incidentId_example" // string | The UUID the incident.
    body := *datadog.NewIncidentUpdateRequest(*datadog.NewIncidentUpdateData("00000000-0000-0000-0000-000000000000", datadog.IncidentType("incidents"))) // IncidentUpdateRequest | Incident Payload.

    configuration := datadog.NewConfiguration()
    configuration.SetUnstableOperationEnabled("UpdateIncident", true)

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.IncidentsApi.UpdateIncident(ctx, incidentId, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.UpdateIncident`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIncident`: IncidentResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from IncidentsApi.UpdateIncident:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**incidentId** | **string** | The UUID the incident. |  |
**body** | [**IncidentUpdateRequest**](IncidentUpdateRequest.md) | Incident Payload. | 


### Optional Parameters

This endpoint does not have optional parameters.


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

