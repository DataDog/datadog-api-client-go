# KeyManagementApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------ | ------------ | ------------
[**CreateAPIKey**](KeyManagementApi.md#CreateAPIKey) | **Post** /api/v2/api_keys | Create an API key
[**CreateCurrentUserApplicationKey**](KeyManagementApi.md#CreateCurrentUserApplicationKey) | **Post** /api/v2/current_user/application_keys | Create an application key for current user
[**DeleteAPIKey**](KeyManagementApi.md#DeleteAPIKey) | **Delete** /api/v2/api_keys/{api_key_id} | Delete an API key
[**DeleteApplicationKey**](KeyManagementApi.md#DeleteApplicationKey) | **Delete** /api/v2/application_keys/{app_key_id} | Delete an application key
[**DeleteCurrentUserApplicationKey**](KeyManagementApi.md#DeleteCurrentUserApplicationKey) | **Delete** /api/v2/current_user/application_keys/{app_key_id} | Delete an application key owned by current user
[**GetAPIKey**](KeyManagementApi.md#GetAPIKey) | **Get** /api/v2/api_keys/{api_key_id} | Get API key
[**GetApplicationKey**](KeyManagementApi.md#GetApplicationKey) | **Get** /api/v2/application_keys/{app_key_id} | Get an application key
[**GetCurrentUserApplicationKey**](KeyManagementApi.md#GetCurrentUserApplicationKey) | **Get** /api/v2/current_user/application_keys/{app_key_id} | Get one application key owned by current user
[**ListAPIKeys**](KeyManagementApi.md#ListAPIKeys) | **Get** /api/v2/api_keys | Get all API keys
[**ListApplicationKeys**](KeyManagementApi.md#ListApplicationKeys) | **Get** /api/v2/application_keys | Get all application keys
[**ListCurrentUserApplicationKeys**](KeyManagementApi.md#ListCurrentUserApplicationKeys) | **Get** /api/v2/current_user/application_keys | Get all application keys owned by current user
[**UpdateAPIKey**](KeyManagementApi.md#UpdateAPIKey) | **Patch** /api/v2/api_keys/{api_key_id} | Edit an API key
[**UpdateApplicationKey**](KeyManagementApi.md#UpdateApplicationKey) | **Patch** /api/v2/application_keys/{app_key_id} | Edit an application key
[**UpdateCurrentUserApplicationKey**](KeyManagementApi.md#UpdateCurrentUserApplicationKey) | **Patch** /api/v2/current_user/application_keys/{app_key_id} | Edit an application key owned by current user



## CreateAPIKey

> APIKeyResponse CreateAPIKey(ctx, body)

Create an API key.

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

    body := *datadog.NewAPIKeyCreateRequest(*datadog.NewAPIKeyCreateData(*datadog.NewAPIKeyCreateAttributes("API Key for submitting metrics"), datadog.APIKeysType("api_keys"))) // APIKeyCreateRequest | 

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyManagementApi.CreateAPIKey(ctx, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyManagementApi.CreateAPIKey`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAPIKey`: APIKeyResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from KeyManagementApi.CreateAPIKey:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**body** | [**APIKeyCreateRequest**](APIKeyCreateRequest.md) |  | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**APIKeyResponse**](APIKeyResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCurrentUserApplicationKey

> ApplicationKeyResponse CreateCurrentUserApplicationKey(ctx, body)

Create an application key for current user

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

    body := *datadog.NewApplicationKeyCreateRequest(*datadog.NewApplicationKeyCreateData(*datadog.NewApplicationKeyCreateAttributes("Application Key for managing dashboards"), datadog.ApplicationKeysType("application_keys"))) // ApplicationKeyCreateRequest | 

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyManagementApi.CreateCurrentUserApplicationKey(ctx, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyManagementApi.CreateCurrentUserApplicationKey`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCurrentUserApplicationKey`: ApplicationKeyResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from KeyManagementApi.CreateCurrentUserApplicationKey:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
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


## DeleteAPIKey

> DeleteAPIKey(ctx, apiKeyId)

Delete an API key.

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

    apiKeyId := "apiKeyId_example" // string | The ID of the API key.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    r, err := apiClient.KeyManagementApi.DeleteAPIKey(ctx, apiKeyId)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyManagementApi.DeleteAPIKey`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**apiKeyId** | **string** | The ID of the API key. | 


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


## DeleteApplicationKey

> DeleteApplicationKey(ctx, appKeyId)

Delete an application key

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

    appKeyId := "appKeyId_example" // string | The ID of the application key.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    r, err := apiClient.KeyManagementApi.DeleteApplicationKey(ctx, appKeyId)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyManagementApi.DeleteApplicationKey`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
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


## DeleteCurrentUserApplicationKey

> DeleteCurrentUserApplicationKey(ctx, appKeyId)

Delete an application key owned by current user

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

    appKeyId := "appKeyId_example" // string | The ID of the application key.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    r, err := apiClient.KeyManagementApi.DeleteCurrentUserApplicationKey(ctx, appKeyId)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyManagementApi.DeleteCurrentUserApplicationKey`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
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


## GetAPIKey

> APIKeyResponse GetAPIKey(ctx, apiKeyId, datadog.GetAPIKeyOptionalParameters{})

Get an API key.

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

    apiKeyId := "apiKeyId_example" // string | The ID of the API key.
    include := "created_by,modified_by" // string | Comma separated list of resource paths for related resources to include in the response. Supported resource paths are `created_by` and `modified_by`. (optional)
    optionalParams := datadog.GetAPIKeyOptionalParameters{
        Include: &include,
    }

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyManagementApi.GetAPIKey(ctx, apiKeyId, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyManagementApi.GetAPIKey`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAPIKey`: APIKeyResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from KeyManagementApi.GetAPIKey:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**apiKeyId** | **string** | The ID of the API key. | 


### Optional Parameters


Other parameters are passed through a pointer to a GetAPIKeyOptionalParameters struct.


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**include** | **string** | Comma separated list of resource paths for related resources to include in the response. Supported resource paths are &#x60;created_by&#x60; and &#x60;modified_by&#x60;. | 

### Return type

[**APIKeyResponse**](APIKeyResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationKey

> ApplicationKeyResponse GetApplicationKey(ctx, appKeyId, datadog.GetApplicationKeyOptionalParameters{})

Get an application key for your org.

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

    appKeyId := "appKeyId_example" // string | The ID of the application key.
    include := "owned_by" // string | Resource path for related resources to include in the response. Only `owned_by` is supported. (optional)
    optionalParams := datadog.GetApplicationKeyOptionalParameters{
        Include: &include,
    }

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyManagementApi.GetApplicationKey(ctx, appKeyId, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyManagementApi.GetApplicationKey`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationKey`: ApplicationKeyResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from KeyManagementApi.GetApplicationKey:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**appKeyId** | **string** | The ID of the application key. | 


### Optional Parameters


Other parameters are passed through a pointer to a GetApplicationKeyOptionalParameters struct.


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**include** | **string** | Resource path for related resources to include in the response. Only &#x60;owned_by&#x60; is supported. | 

### Return type

[**ApplicationKeyResponse**](ApplicationKeyResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrentUserApplicationKey

> ApplicationKeyResponse GetCurrentUserApplicationKey(ctx, appKeyId)

Get an application key owned by current user

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

    appKeyId := "appKeyId_example" // string | The ID of the application key.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyManagementApi.GetCurrentUserApplicationKey(ctx, appKeyId)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyManagementApi.GetCurrentUserApplicationKey`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCurrentUserApplicationKey`: ApplicationKeyResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from KeyManagementApi.GetCurrentUserApplicationKey:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**appKeyId** | **string** | The ID of the application key. | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**ApplicationKeyResponse**](ApplicationKeyResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAPIKeys

> APIKeysResponse ListAPIKeys(ctx, datadog.ListAPIKeysOptionalParameters{})

List all API keys available for your account.

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
    sort := datadog.APIKeysSort("created_at") // APIKeysSort | API key attribute used to sort results. Sort order is ascending by default. In order to specify a descending sort, prefix the attribute with a minus sign. (optional) (default to "name")
    filter := "filter_example" // string | Filter API keys by the specified string. (optional)
    filterCreatedAtStart := "2020-11-24T18:46:21+00:00" // string | Only include API keys created on or after the specified date. (optional)
    filterCreatedAtEnd := "2020-11-24T18:46:21+00:00" // string | Only include API keys created on or before the specified date. (optional)
    filterModifiedAtStart := "2020-11-24T18:46:21+00:00" // string | Only include API keys modified on or after the specified date. (optional)
    filterModifiedAtEnd := "2020-11-24T18:46:21+00:00" // string | Only include API keys modified on or before the specified date. (optional)
    include := "created_by,modified_by" // string | Comma separated list of resource paths for related resources to include in the response. Supported resource paths are `created_by` and `modified_by`. (optional)
    optionalParams := datadog.ListAPIKeysOptionalParameters{
        PageSize: &pageSize,
        PageNumber: &pageNumber,
        Sort: &sort,
        Filter: &filter,
        FilterCreatedAtStart: &filterCreatedAtStart,
        FilterCreatedAtEnd: &filterCreatedAtEnd,
        FilterModifiedAtStart: &filterModifiedAtStart,
        FilterModifiedAtEnd: &filterModifiedAtEnd,
        Include: &include,
    }

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyManagementApi.ListAPIKeys(ctx, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyManagementApi.ListAPIKeys`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAPIKeys`: APIKeysResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from KeyManagementApi.ListAPIKeys:\n%s\n", responseContent)
}
```

### Required Parameters




### Optional Parameters


Other parameters are passed through a pointer to a ListAPIKeysOptionalParameters struct.


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**pageSize** | **int64** | Size for a given page. | [default to 10]
**pageNumber** | **int64** | Specific page number to return. | [default to 0]
**sort** | [**APIKeysSort**](APIKeysSort.md) | API key attribute used to sort results. Sort order is ascending by default. In order to specify a descending sort, prefix the attribute with a minus sign. | [default to &quot;name&quot;]
**filter** | **string** | Filter API keys by the specified string. | 
**filterCreatedAtStart** | **string** | Only include API keys created on or after the specified date. | 
**filterCreatedAtEnd** | **string** | Only include API keys created on or before the specified date. | 
**filterModifiedAtStart** | **string** | Only include API keys modified on or after the specified date. | 
**filterModifiedAtEnd** | **string** | Only include API keys modified on or before the specified date. | 
**include** | **string** | Comma separated list of resource paths for related resources to include in the response. Supported resource paths are &#x60;created_by&#x60; and &#x60;modified_by&#x60;. | 

### Return type

[**APIKeysResponse**](APIKeysResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApplicationKeys

> ListApplicationKeysResponse ListApplicationKeys(ctx, datadog.ListApplicationKeysOptionalParameters{})

List all application keys available for your org

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
    sort := datadog.ApplicationKeysSort("created_at") // ApplicationKeysSort | Application key attribute used to sort results. Sort order is ascending by default. In order to specify a descending sort, prefix the attribute with a minus sign. (optional) (default to "name")
    filter := "filter_example" // string | Filter application keys by the specified string. (optional)
    filterCreatedAtStart := "2020-11-24T18:46:21+00:00" // string | Only include application keys created on or after the specified date. (optional)
    filterCreatedAtEnd := "2020-11-24T18:46:21+00:00" // string | Only include application keys created on or before the specified date. (optional)
    optionalParams := datadog.ListApplicationKeysOptionalParameters{
        PageSize: &pageSize,
        PageNumber: &pageNumber,
        Sort: &sort,
        Filter: &filter,
        FilterCreatedAtStart: &filterCreatedAtStart,
        FilterCreatedAtEnd: &filterCreatedAtEnd,
    }

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyManagementApi.ListApplicationKeys(ctx, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyManagementApi.ListApplicationKeys`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApplicationKeys`: ListApplicationKeysResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from KeyManagementApi.ListApplicationKeys:\n%s\n", responseContent)
}
```

### Required Parameters




### Optional Parameters


Other parameters are passed through a pointer to a ListApplicationKeysOptionalParameters struct.


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


## ListCurrentUserApplicationKeys

> ListApplicationKeysResponse ListCurrentUserApplicationKeys(ctx, datadog.ListCurrentUserApplicationKeysOptionalParameters{})

List all application keys available for current user

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
    sort := datadog.ApplicationKeysSort("created_at") // ApplicationKeysSort | Application key attribute used to sort results. Sort order is ascending by default. In order to specify a descending sort, prefix the attribute with a minus sign. (optional) (default to "name")
    filter := "filter_example" // string | Filter application keys by the specified string. (optional)
    filterCreatedAtStart := "2020-11-24T18:46:21+00:00" // string | Only include application keys created on or after the specified date. (optional)
    filterCreatedAtEnd := "2020-11-24T18:46:21+00:00" // string | Only include application keys created on or before the specified date. (optional)
    optionalParams := datadog.ListCurrentUserApplicationKeysOptionalParameters{
        PageSize: &pageSize,
        PageNumber: &pageNumber,
        Sort: &sort,
        Filter: &filter,
        FilterCreatedAtStart: &filterCreatedAtStart,
        FilterCreatedAtEnd: &filterCreatedAtEnd,
    }

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyManagementApi.ListCurrentUserApplicationKeys(ctx, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyManagementApi.ListCurrentUserApplicationKeys`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCurrentUserApplicationKeys`: ListApplicationKeysResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from KeyManagementApi.ListCurrentUserApplicationKeys:\n%s\n", responseContent)
}
```

### Required Parameters




### Optional Parameters


Other parameters are passed through a pointer to a ListCurrentUserApplicationKeysOptionalParameters struct.


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


## UpdateAPIKey

> APIKeyResponse UpdateAPIKey(ctx, apiKeyId, body)

Update an API key.

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

    apiKeyId := "apiKeyId_example" // string | The ID of the API key.
    body := *datadog.NewAPIKeyUpdateRequest(*datadog.NewAPIKeyUpdateData(*datadog.NewAPIKeyUpdateAttributes("API Key for submitting metrics"), "00112233-4455-6677-8899-aabbccddeeff", datadog.APIKeysType("api_keys"))) // APIKeyUpdateRequest | 

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyManagementApi.UpdateAPIKey(ctx, apiKeyId, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyManagementApi.UpdateAPIKey`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAPIKey`: APIKeyResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from KeyManagementApi.UpdateAPIKey:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**apiKeyId** | **string** | The ID of the API key. |  |
**body** | [**APIKeyUpdateRequest**](APIKeyUpdateRequest.md) |  | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**APIKeyResponse**](APIKeyResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApplicationKey

> ApplicationKeyResponse UpdateApplicationKey(ctx, appKeyId, body)

Edit an application key

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

    appKeyId := "appKeyId_example" // string | The ID of the application key.
    body := *datadog.NewApplicationKeyUpdateRequest(*datadog.NewApplicationKeyUpdateData(*datadog.NewApplicationKeyUpdateAttributes(), "00112233-4455-6677-8899-aabbccddeeff", datadog.ApplicationKeysType("application_keys"))) // ApplicationKeyUpdateRequest | 

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyManagementApi.UpdateApplicationKey(ctx, appKeyId, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyManagementApi.UpdateApplicationKey`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApplicationKey`: ApplicationKeyResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from KeyManagementApi.UpdateApplicationKey:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**appKeyId** | **string** | The ID of the application key. |  |
**body** | [**ApplicationKeyUpdateRequest**](ApplicationKeyUpdateRequest.md) |  | 


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


## UpdateCurrentUserApplicationKey

> ApplicationKeyResponse UpdateCurrentUserApplicationKey(ctx, appKeyId, body)

Edit an application key owned by current user

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

    appKeyId := "appKeyId_example" // string | The ID of the application key.
    body := *datadog.NewApplicationKeyUpdateRequest(*datadog.NewApplicationKeyUpdateData(*datadog.NewApplicationKeyUpdateAttributes(), "00112233-4455-6677-8899-aabbccddeeff", datadog.ApplicationKeysType("application_keys"))) // ApplicationKeyUpdateRequest | 

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyManagementApi.UpdateCurrentUserApplicationKey(ctx, appKeyId, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyManagementApi.UpdateCurrentUserApplicationKey`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCurrentUserApplicationKey`: ApplicationKeyResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from KeyManagementApi.UpdateCurrentUserApplicationKey:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**appKeyId** | **string** | The ID of the application key. |  |
**body** | [**ApplicationKeyUpdateRequest**](ApplicationKeyUpdateRequest.md) |  | 


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

