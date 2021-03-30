# \IncidentTeamsApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIncidentTeam**](IncidentTeamsApi.md#CreateIncidentTeam) | **Post** /api/v2/teams | Create a new incident team
[**DeleteIncidentTeam**](IncidentTeamsApi.md#DeleteIncidentTeam) | **Delete** /api/v2/teams/{team_id} | Delete an existing incident team
[**GetIncidentTeam**](IncidentTeamsApi.md#GetIncidentTeam) | **Get** /api/v2/teams/{team_id} | Get details of an incident team
[**ListIncidentTeams**](IncidentTeamsApi.md#ListIncidentTeams) | **Get** /api/v2/teams | Get a list of all incident teams
[**UpdateIncidentTeam**](IncidentTeamsApi.md#UpdateIncidentTeam) | **Patch** /api/v2/teams/{team_id} | Update an existing incident team



## CreateIncidentTeam

> IncidentTeamResponse CreateIncidentTeam(ctx).Body(body).Execute()

Create a new incident team



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

    body := *datadog.NewIncidentTeamCreateRequest(*datadog.NewIncidentTeamCreateData(datadog.IncidentTeamType("teams"))) // IncidentTeamCreateRequest | Incident Team Payload.

    configuration := datadog.NewConfiguration()
    configuration.SetUnstableOperationEnabled("CreateIncidentTeam", true)

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.IncidentTeamsApi.CreateIncidentTeam(ctx).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IncidentTeamsApi.CreateIncidentTeam``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIncidentTeam`: IncidentTeamResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from IncidentTeamsApi.CreateIncidentTeam:\n%s\n", responseContent)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIncidentTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**IncidentTeamCreateRequest**](IncidentTeamCreateRequest.md) | Incident Team Payload. | 

### Return type

[**IncidentTeamResponse**](IncidentTeamResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIncidentTeam

> DeleteIncidentTeam(ctx, teamId).Execute()

Delete an existing incident team



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

    teamId := "teamId_example" // string | The ID of the incident team.

    configuration := datadog.NewConfiguration()
    configuration.SetUnstableOperationEnabled("DeleteIncidentTeam", true)

    apiClient := datadog.NewAPIClient(configuration)
    r, err := apiClient.IncidentTeamsApi.DeleteIncidentTeam(ctx, teamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IncidentTeamsApi.DeleteIncidentTeam``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | The ID of the incident team. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIncidentTeamRequest struct via the builder pattern


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


## GetIncidentTeam

> IncidentTeamResponse GetIncidentTeam(ctx, teamId).Include(include).Execute()

Get details of an incident team



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

    teamId := "teamId_example" // string | The ID of the incident team.
    include := datadog.IncidentRelatedObject("users") // IncidentRelatedObject | Specifies which types of related objects should be included in the response. (optional)

    configuration := datadog.NewConfiguration()
    configuration.SetUnstableOperationEnabled("GetIncidentTeam", true)

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.IncidentTeamsApi.GetIncidentTeam(ctx, teamId).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IncidentTeamsApi.GetIncidentTeam``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIncidentTeam`: IncidentTeamResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from IncidentTeamsApi.GetIncidentTeam:\n%s\n", responseContent)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | The ID of the incident team. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIncidentTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | [**IncidentRelatedObject**](IncidentRelatedObject.md) | Specifies which types of related objects should be included in the response. | 

### Return type

[**IncidentTeamResponse**](IncidentTeamResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIncidentTeams

> IncidentTeamsResponse ListIncidentTeams(ctx).Include(include).PageSize(pageSize).PageOffset(pageOffset).Filter(filter).Execute()

Get a list of all incident teams



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
    pageSize := int64(789) // int64 | Size for a given page. (optional) (default to 10)
    pageOffset := int64(789) // int64 | Specific offset to use as the beginning of the returned page. (optional) (default to 0)
    filter := "ExampleTeamName" // string | A search query that filters teams by name. (optional)

    configuration := datadog.NewConfiguration()
    configuration.SetUnstableOperationEnabled("ListIncidentTeams", true)

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.IncidentTeamsApi.ListIncidentTeams(ctx).Include(include).PageSize(pageSize).PageOffset(pageOffset).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IncidentTeamsApi.ListIncidentTeams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIncidentTeams`: IncidentTeamsResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from IncidentTeamsApi.ListIncidentTeams:\n%s\n", responseContent)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListIncidentTeamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | [**IncidentRelatedObject**](IncidentRelatedObject.md) | Specifies which types of related objects should be included in the response. | 
 **pageSize** | **int64** | Size for a given page. | [default to 10]
 **pageOffset** | **int64** | Specific offset to use as the beginning of the returned page. | [default to 0]
 **filter** | **string** | A search query that filters teams by name. | 

### Return type

[**IncidentTeamsResponse**](IncidentTeamsResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIncidentTeam

> IncidentTeamResponse UpdateIncidentTeam(ctx, teamId).Body(body).Execute()

Update an existing incident team



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

    teamId := "teamId_example" // string | The ID of the incident team.
    body := *datadog.NewIncidentTeamUpdateRequest(*datadog.NewIncidentTeamUpdateData("00000000-0000-0000-0000-000000000000", datadog.IncidentTeamType("teams"))) // IncidentTeamUpdateRequest | Incident Team Payload.

    configuration := datadog.NewConfiguration()
    configuration.SetUnstableOperationEnabled("UpdateIncidentTeam", true)

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.IncidentTeamsApi.UpdateIncidentTeam(ctx, teamId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IncidentTeamsApi.UpdateIncidentTeam``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIncidentTeam`: IncidentTeamResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from IncidentTeamsApi.UpdateIncidentTeam:\n%s\n", responseContent)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | The ID of the incident team. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIncidentTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**IncidentTeamUpdateRequest**](IncidentTeamUpdateRequest.md) | Incident Team Payload. | 

### Return type

[**IncidentTeamResponse**](IncidentTeamResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

