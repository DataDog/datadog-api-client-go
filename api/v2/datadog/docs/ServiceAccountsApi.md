# ServiceAccountsApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------ | ------------ | ------------
[**CreateServiceAccountApplicationKey**](ServiceAccountsApi.md#CreateServiceAccountApplicationKey) | **Post** /api/v2/service_accounts/{service_account_id}/application_keys | Create an application key for this service account
[**DeleteServiceAccountApplicationKey**](ServiceAccountsApi.md#DeleteServiceAccountApplicationKey) | **Delete** /api/v2/service_accounts/{service_account_id}/application_keys/{app_key_id} | Delete an application key for this service account
[**GetServiceAccountApplicationKey**](ServiceAccountsApi.md#GetServiceAccountApplicationKey) | **Get** /api/v2/service_accounts/{service_account_id}/application_keys/{app_key_id} | Get one application key for this service account
[**ListServiceAccountApplicationKeys**](ServiceAccountsApi.md#ListServiceAccountApplicationKeys) | **Get** /api/v2/service_accounts/{service_account_id}/application_keys | List application keys for this service account
[**UpdateServiceAccountApplicationKey**](ServiceAccountsApi.md#UpdateServiceAccountApplicationKey) | **Patch** /api/v2/service_accounts/{service_account_id}/application_keys/{app_key_id} | Edit an application key for this service account



## CreateServiceAccountApplicationKey

> ApplicationKeyResponse CreateServiceAccountApplicationKey(ctx, serviceAccountId, body)

Create an application key for this service account.

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

    serviceAccountId := "00000000-0000-0000-0000-000000000000" // string | The ID of the service account.
    body := *datadog.NewApplicationKeyCreateRequest(*datadog.NewApplicationKeyCreateData(*datadog.NewApplicationKeyCreateAttributes("Application Key for managing dashboards"), datadog.ApplicationKeysType("application_keys"))) // ApplicationKeyCreateRequest | 

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceAccountsApi.CreateServiceAccountApplicationKey(ctx, serviceAccountId, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsApi.CreateServiceAccountApplicationKey`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateServiceAccountApplicationKey`: ApplicationKeyResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from ServiceAccountsApi.CreateServiceAccountApplicationKey:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**serviceAccountId** | **string** | The ID of the service account. |  |
**body** | [**ApplicationKeyCreateRequest**](ApplicationKeyCreateRequest.md) |  | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**ApplicationKeyResponse**](ApplicationKeyResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteServiceAccountApplicationKey

> DeleteServiceAccountApplicationKey(ctx, serviceAccountId, appKeyId)

Delete an application key owned by this service account.

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

    serviceAccountId := "00000000-0000-0000-0000-000000000000" // string | The ID of the service account.
    appKeyId := "appKeyId_example" // string | The ID of the application key.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    r, err := apiClient.ServiceAccountsApi.DeleteServiceAccountApplicationKey(ctx, serviceAccountId, appKeyId)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsApi.DeleteServiceAccountApplicationKey`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**serviceAccountId** | **string** | The ID of the service account. |  |
**appKeyId** | **string** | The ID of the application key. | 


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


## GetServiceAccountApplicationKey

> PartialApplicationKeyResponse GetServiceAccountApplicationKey(ctx, serviceAccountId, appKeyId)

Get an application key owned by this service account.

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

    serviceAccountId := "00000000-0000-0000-0000-000000000000" // string | The ID of the service account.
    appKeyId := "appKeyId_example" // string | The ID of the application key.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceAccountsApi.GetServiceAccountApplicationKey(ctx, serviceAccountId, appKeyId)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsApi.GetServiceAccountApplicationKey`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServiceAccountApplicationKey`: PartialApplicationKeyResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from ServiceAccountsApi.GetServiceAccountApplicationKey:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**serviceAccountId** | **string** | The ID of the service account. |  |
**appKeyId** | **string** | The ID of the application key. | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**PartialApplicationKeyResponse**](PartialApplicationKeyResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServiceAccountApplicationKeys

> ListApplicationKeysResponse ListServiceAccountApplicationKeys(ctx, serviceAccountId, datadog.ListServiceAccountApplicationKeysOptionalParameters{})

List all application keys available for this service account.

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

    serviceAccountId := "00000000-0000-0000-0000-000000000000" // string | The ID of the service account.
    pageSize := int64(10) // int64 | Size for a given page. (optional) (default to 10)
    pageNumber := int64(0) // int64 | Specific page number to return. (optional) (default to 0)
    sort := datadog.ApplicationKeysSort("created_at") // ApplicationKeysSort | Application key attribute used to sort results. Sort order is ascending by default. In order to specify a descending sort, prefix the attribute with a minus sign. (optional) (default to "name")
    filter := "filter_example" // string | Filter application keys by the specified string. (optional)
    filterCreatedAtStart := "2020-11-24T18:46:21+00:00" // string | Only include application keys created on or after the specified date. (optional)
    filterCreatedAtEnd := "2020-11-24T18:46:21+00:00" // string | Only include application keys created on or before the specified date. (optional)
    optionalParams := datadog.ListServiceAccountApplicationKeysOptionalParameters{
        PageSize: &pageSize,
        PageNumber: &pageNumber,
        Sort: &sort,
        Filter: &filter,
        FilterCreatedAtStart: &filterCreatedAtStart,
        FilterCreatedAtEnd: &filterCreatedAtEnd,
    }

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceAccountsApi.ListServiceAccountApplicationKeys(ctx, serviceAccountId, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsApi.ListServiceAccountApplicationKeys`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListServiceAccountApplicationKeys`: ListApplicationKeysResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from ServiceAccountsApi.ListServiceAccountApplicationKeys:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**serviceAccountId** | **string** | The ID of the service account. | 


### Optional Parameters


Other parameters are passed through a pointer to a ListServiceAccountApplicationKeysOptionalParameters struct.


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**pageSize** | **int64** | Size for a given page. | [default to 10]
**pageNumber** | **int64** | Specific page number to return. | [default to 0]
**sort** | [**ApplicationKeysSort**](ApplicationKeysSort.md) | Application key attribute used to sort results. Sort order is ascending by default. In order to specify a descending sort, prefix the attribute with a minus sign. | [default to &quot;name&quot;]
**filter** | **string** | Filter application keys by the specified string. | 
**filterCreatedAtStart** | **string** | Only include application keys created on or after the specified date. | 
**filterCreatedAtEnd** | **string** | Only include application keys created on or before the specified date. | 

### Return type

[**ListApplicationKeysResponse**](ListApplicationKeysResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServiceAccountApplicationKey

> PartialApplicationKeyResponse UpdateServiceAccountApplicationKey(ctx, serviceAccountId, appKeyId, body)

Edit an application key owned by this service account.

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

    serviceAccountId := "00000000-0000-0000-0000-000000000000" // string | The ID of the service account.
    appKeyId := "appKeyId_example" // string | The ID of the application key.
    body := *datadog.NewApplicationKeyUpdateRequest(*datadog.NewApplicationKeyUpdateData(*datadog.NewApplicationKeyUpdateAttributes(), "00112233-4455-6677-8899-aabbccddeeff", datadog.ApplicationKeysType("application_keys"))) // ApplicationKeyUpdateRequest | 

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceAccountsApi.UpdateServiceAccountApplicationKey(ctx, serviceAccountId, appKeyId, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsApi.UpdateServiceAccountApplicationKey`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateServiceAccountApplicationKey`: PartialApplicationKeyResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from ServiceAccountsApi.UpdateServiceAccountApplicationKey:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**serviceAccountId** | **string** | The ID of the service account. |  |
**appKeyId** | **string** | The ID of the application key. |  |
**body** | [**ApplicationKeyUpdateRequest**](ApplicationKeyUpdateRequest.md) |  | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**PartialApplicationKeyResponse**](PartialApplicationKeyResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

