# AuthNMappingsApi

All URIs are relative to *https://api.datadoghq.com*

| Method                                                           | HTTP request                                         | Description                  |
| ---------------------------------------------------------------- | ---------------------------------------------------- | ---------------------------- |
| [**CreateAuthNMapping**](AuthNMappingsApi.md#CreateAuthNMapping) | **Post** /api/v2/authn_mappings                      | Create an AuthN Mapping      |
| [**DeleteAuthNMapping**](AuthNMappingsApi.md#DeleteAuthNMapping) | **Delete** /api/v2/authn_mappings/{authn_mapping_id} | Delete an AuthN Mapping      |
| [**GetAuthNMapping**](AuthNMappingsApi.md#GetAuthNMapping)       | **Get** /api/v2/authn_mappings/{authn_mapping_id}    | Get an AuthN Mapping by UUID |
| [**ListAuthNMappings**](AuthNMappingsApi.md#ListAuthNMappings)   | **Get** /api/v2/authn_mappings                       | List all AuthN Mappings      |
| [**UpdateAuthNMapping**](AuthNMappingsApi.md#UpdateAuthNMapping) | **Patch** /api/v2/authn_mappings/{authn_mapping_id}  | Edit an AuthN Mapping        |

## CreateAuthNMapping

> AuthNMappingResponse CreateAuthNMapping(ctx, body)

Create an AuthN Mapping.

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

    body := *datadog.NewAuthNMappingCreateRequest(*datadog.NewAuthNMappingCreateData(datadog.AuthNMappingsType("authn_mappings"))) // AuthNMappingCreateRequest |

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthNMappingsApi.CreateAuthNMapping(ctx, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthNMappingsApi.CreateAuthNMapping`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAuthNMapping`: AuthNMappingResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from AuthNMappingsApi.CreateAuthNMapping:\n%s\n", responseContent)
}
```

### Required Parameters

| Name     | Type                                                          | Description                                                                 | Notes |
| -------- | ------------------------------------------------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**  | **context.Context**                                           | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **body** | [**AuthNMappingCreateRequest**](AuthNMappingCreateRequest.md) |                                                                             |

### Optional Parameters

This endpoint does not have optional parameters.

### Return type

[**AuthNMappingResponse**](AuthNMappingResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## DeleteAuthNMapping

> DeleteAuthNMapping(ctx, authnMappingId)

Delete an AuthN Mapping specified by AuthN Mapping UUID.

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

    authnMappingId := "authnMappingId_example" // string | The UUID of the AuthN Mapping.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    r, err := apiClient.AuthNMappingsApi.DeleteAuthNMapping(ctx, authnMappingId)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthNMappingsApi.DeleteAuthNMapping`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Required Parameters

| Name               | Type                | Description                                                                 | Notes |
| ------------------ | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**            | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **authnMappingId** | **string**          | The UUID of the AuthN Mapping.                                              |

### Optional Parameters

This endpoint does not have optional parameters.

### Return type

(empty response body)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetAuthNMapping

> AuthNMappingResponse GetAuthNMapping(ctx, authnMappingId)

Get an AuthN Mapping specified by the AuthN Mapping UUID.

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

    authnMappingId := "authnMappingId_example" // string | The UUID of the AuthN Mapping.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthNMappingsApi.GetAuthNMapping(ctx, authnMappingId)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthNMappingsApi.GetAuthNMapping`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuthNMapping`: AuthNMappingResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from AuthNMappingsApi.GetAuthNMapping:\n%s\n", responseContent)
}
```

### Required Parameters

| Name               | Type                | Description                                                                 | Notes |
| ------------------ | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**            | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **authnMappingId** | **string**          | The UUID of the AuthN Mapping.                                              |

### Optional Parameters

This endpoint does not have optional parameters.

### Return type

[**AuthNMappingResponse**](AuthNMappingResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListAuthNMappings

> AuthNMappingsResponse ListAuthNMappings(ctx, datadog.ListAuthNMappingsOptionalParameters{})

List all AuthN Mappings in the org.

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

    pageSize := int64(10) // int64 | Size for a given page. (optional) (default to 10)
    pageNumber := int64(0) // int64 | Specific page number to return. (optional) (default to 0)
    sort := datadog.AuthNMappingsSort("created_at") // AuthNMappingsSort | Sort AuthN Mappings depending on the given field. (optional)
    include := []string{"Inner_example"} // []string |  (optional)
    filter := "filter_example" // string | Filter all mappings by the given string. (optional)
    optionalParams := datadog.ListAuthNMappingsOptionalParameters{
        PageSize: &pageSize,
        PageNumber: &pageNumber,
        Sort: &sort,
        Include: &include,
        Filter: &filter,
    }

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthNMappingsApi.ListAuthNMappings(ctx, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthNMappingsApi.ListAuthNMappings`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAuthNMappings`: AuthNMappingsResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from AuthNMappingsApi.ListAuthNMappings:\n%s\n", responseContent)
}
```

### Required Parameters

### Optional Parameters

Other parameters are passed through a pointer to a ListAuthNMappingsOptionalParameters struct.

| Name           | Type                                          | Description                                       | Notes           |
| -------------- | --------------------------------------------- | ------------------------------------------------- | --------------- |
| **pageSize**   | **int64**                                     | Size for a given page.                            | [default to 10] |
| **pageNumber** | **int64**                                     | Specific page number to return.                   | [default to 0]  |
| **sort**       | [**AuthNMappingsSort**](AuthNMappingsSort.md) | Sort AuthN Mappings depending on the given field. |
| **include**    | **[]string**                                  |                                                   |
| **filter**     | **string**                                    | Filter all mappings by the given string.          |

### Return type

[**AuthNMappingsResponse**](AuthNMappingsResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateAuthNMapping

> AuthNMappingResponse UpdateAuthNMapping(ctx, authnMappingId, body)

Edit an AuthN Mapping.

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

    authnMappingId := "authnMappingId_example" // string | The UUID of the AuthN Mapping.
    body := *datadog.NewAuthNMappingUpdateRequest(*datadog.NewAuthNMappingUpdateData("3653d3c6-0c75-11ea-ad28-fb5701eabc7d", datadog.AuthNMappingsType("authn_mappings"))) // AuthNMappingUpdateRequest |

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthNMappingsApi.UpdateAuthNMapping(ctx, authnMappingId, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthNMappingsApi.UpdateAuthNMapping`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAuthNMapping`: AuthNMappingResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from AuthNMappingsApi.UpdateAuthNMapping:\n%s\n", responseContent)
}
```

### Required Parameters

| Name               | Type                                                          | Description                                                                 | Notes |
| ------------------ | ------------------------------------------------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**            | **context.Context**                                           | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **authnMappingId** | **string**                                                    | The UUID of the AuthN Mapping.                                              |       |
| **body**           | [**AuthNMappingUpdateRequest**](AuthNMappingUpdateRequest.md) |                                                                             |

### Optional Parameters

This endpoint does not have optional parameters.

### Return type

[**AuthNMappingResponse**](AuthNMappingResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
