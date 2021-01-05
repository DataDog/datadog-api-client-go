# \KeyManagementApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAPIKey**](KeyManagementApi.md#CreateAPIKey) | **Post** /api/v2/api_keys | Create an API key
[**CreateCurrentUserApplicationKey**](KeyManagementApi.md#CreateCurrentUserApplicationKey) | **Post** /api/v2/current_user/application_keys | Create an application key for current user
[**DeleteAPIKey**](KeyManagementApi.md#DeleteAPIKey) | **Delete** /api/v2/api_keys/{api_key_id} | Delete an API key
[**DeleteApplicationKey**](KeyManagementApi.md#DeleteApplicationKey) | **Delete** /api/v2/application_keys/{app_key_id} | Delete an application key
[**DeleteCurrentUserApplicationKey**](KeyManagementApi.md#DeleteCurrentUserApplicationKey) | **Delete** /api/v2/current_user/application_keys/{app_key_id} | Delete an application key owned by current user
[**GetAPIKey**](KeyManagementApi.md#GetAPIKey) | **Get** /api/v2/api_keys/{api_key_id} | Get API key
[**GetCurrentUserApplicationKey**](KeyManagementApi.md#GetCurrentUserApplicationKey) | **Get** /api/v2/current_user/application_keys/{app_key_id} | Get one application key owned by current user
[**ListAPIKeys**](KeyManagementApi.md#ListAPIKeys) | **Get** /api/v2/api_keys | Get all API keys
[**ListApplicationKeys**](KeyManagementApi.md#ListApplicationKeys) | **Get** /api/v2/application_keys | Get all application keys
[**ListCurrentUserApplicationKeys**](KeyManagementApi.md#ListCurrentUserApplicationKeys) | **Get** /api/v2/current_user/application_keys | Get all application keys owned by current user
[**UpdateAPIKey**](KeyManagementApi.md#UpdateAPIKey) | **Patch** /api/v2/api_keys/{api_key_id} | Edit an API key
[**UpdateApplicationKey**](KeyManagementApi.md#UpdateApplicationKey) | **Patch** /api/v2/application_keys/{app_key_id} | Edit an application key
[**UpdateCurrentUserApplicationKey**](KeyManagementApi.md#UpdateCurrentUserApplicationKey) | **Patch** /api/v2/current_user/application_keys/{app_key_id} | Edit an application key owned by current user



## CreateAPIKey

> APIKeyResponse CreateAPIKey(ctx).Body(body).Execute()

Create an API key



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

    body := *datadog.NewAPIKeyCreateRequest(*datadog.NewAPIKeyCreateData(*datadog.NewAPIKeyCreateAttributes("API Key for submitting metrics"), datadog.APIKeysType("api_keys"))) // APIKeyCreateRequest | 

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.KeyManagementApi.CreateAPIKey(ctx).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyManagementApi.CreateAPIKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAPIKey`: APIKeyResponse
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from KeyManagementApi.CreateAPIKey:\n%v\n", response_content)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAPIKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**APIKeyCreateRequest**](APIKeyCreateRequest.md) |  | 

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

> ApplicationKeyResponse CreateCurrentUserApplicationKey(ctx).Body(body).Execute()

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

    body := *datadog.NewApplicationKeyCreateRequest(*datadog.NewApplicationKeyCreateData(*datadog.NewApplicationKeyCreateAttributes("Application Key for submitting metrics"), datadog.ApplicationKeysType("application_keys"))) // ApplicationKeyCreateRequest | 

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.KeyManagementApi.CreateCurrentUserApplicationKey(ctx).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyManagementApi.CreateCurrentUserApplicationKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCurrentUserApplicationKey`: ApplicationKeyResponse
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from KeyManagementApi.CreateCurrentUserApplicationKey:\n%v\n", response_content)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCurrentUserApplicationKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ApplicationKeyCreateRequest**](ApplicationKeyCreateRequest.md) |  | 

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

> DeleteAPIKey(ctx, apiKeyId).Execute()

Delete an API key



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

    apiKeyId := "apiKeyId_example" // string | The ID of the API key.

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    r, err := api_client.KeyManagementApi.DeleteAPIKey(ctx, apiKeyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyManagementApi.DeleteAPIKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiKeyId** | **string** | The ID of the API key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAPIKeyRequest struct via the builder pattern


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


## DeleteApplicationKey

> DeleteApplicationKey(ctx, appKeyId).Execute()

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

    appKeyId := "appKeyId_example" // string | The ID of the application key.

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    r, err := api_client.KeyManagementApi.DeleteApplicationKey(ctx, appKeyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyManagementApi.DeleteApplicationKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKeyId** | **string** | The ID of the application key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationKeyRequest struct via the builder pattern


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


## DeleteCurrentUserApplicationKey

> DeleteCurrentUserApplicationKey(ctx, appKeyId).Execute()

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

    appKeyId := "appKeyId_example" // string | The ID of the application key.

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    r, err := api_client.KeyManagementApi.DeleteCurrentUserApplicationKey(ctx, appKeyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyManagementApi.DeleteCurrentUserApplicationKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKeyId** | **string** | The ID of the application key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCurrentUserApplicationKeyRequest struct via the builder pattern


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


## GetAPIKey

> APIKeyResponse GetAPIKey(ctx, apiKeyId).Include(include).Execute()

Get API key



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

    apiKeyId := "apiKeyId_example" // string | The ID of the API key.
    include := "created_by,modified_by" // string | Comma separated list of resource paths for related resources to include in the response. Supported resource paths are `created_by` and `modified_by`. (optional)

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.KeyManagementApi.GetAPIKey(ctx, apiKeyId).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyManagementApi.GetAPIKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAPIKey`: APIKeyResponse
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from KeyManagementApi.GetAPIKey:\n%v\n", response_content)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiKeyId** | **string** | The ID of the API key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAPIKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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


## GetCurrentUserApplicationKey

> ApplicationKeyResponse GetCurrentUserApplicationKey(ctx, appKeyId).Execute()

Get one application key owned by current user



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

    appKeyId := "appKeyId_example" // string | The ID of the application key.

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.KeyManagementApi.GetCurrentUserApplicationKey(ctx, appKeyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyManagementApi.GetCurrentUserApplicationKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCurrentUserApplicationKey`: ApplicationKeyResponse
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from KeyManagementApi.GetCurrentUserApplicationKey:\n%v\n", response_content)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKeyId** | **string** | The ID of the application key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentUserApplicationKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> APIKeysResponse ListAPIKeys(ctx).PageSize(pageSize).PageNumber(pageNumber).Sort(sort).Filter(filter).FilterCreatedAtStart(filterCreatedAtStart).FilterCreatedAtEnd(filterCreatedAtEnd).FilterModifiedAtStart(filterModifiedAtStart).FilterModifiedAtEnd(filterModifiedAtEnd).Include(include).Execute()

Get all API keys



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

    pageSize := int64(789) // int64 | Size for a given page. (optional) (default to 10)
    pageNumber := int64(789) // int64 | Specific page number to return. (optional) (default to 0)
    sort := "sort_example" // string | API key attribute used to sort results. Sort order is ascending by default. In order to specify a descending sort, prefix the attribute with a minus sign. (optional) (default to "name")
    filter := "filter_example" // string | Filter API keys by the specified string. (optional)
    filterCreatedAtStart := "2020-11-24T18:46:21+00:00" // string | Only include API keys created on or after the specified date. (optional)
    filterCreatedAtEnd := "2020-11-24T18:46:21+00:00" // string | Only include API keys created on or before the specified date. (optional)
    filterModifiedAtStart := "2020-11-24T18:46:21+00:00" // string | Only include API keys modified on or after the specified date. (optional)
    filterModifiedAtEnd := "2020-11-24T18:46:21+00:00" // string | Only include API keys modified on or before the specified date. (optional)
    include := "created_by,modified_by" // string | Comma separated list of resource paths for related resources to include in the response. Supported resource paths are `created_by` and `modified_by`. (optional)

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.KeyManagementApi.ListAPIKeys(ctx).PageSize(pageSize).PageNumber(pageNumber).Sort(sort).Filter(filter).FilterCreatedAtStart(filterCreatedAtStart).FilterCreatedAtEnd(filterCreatedAtEnd).FilterModifiedAtStart(filterModifiedAtStart).FilterModifiedAtEnd(filterModifiedAtEnd).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyManagementApi.ListAPIKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAPIKeys`: APIKeysResponse
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from KeyManagementApi.ListAPIKeys:\n%v\n", response_content)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAPIKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int64** | Size for a given page. | [default to 10]
 **pageNumber** | **int64** | Specific page number to return. | [default to 0]
 **sort** | **string** | API key attribute used to sort results. Sort order is ascending by default. In order to specify a descending sort, prefix the attribute with a minus sign. | [default to &quot;name&quot;]
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

> ListApplicationKeysResponse ListApplicationKeys(ctx).PageSize(pageSize).PageNumber(pageNumber).Sort(sort).Filter(filter).FilterCreatedAtStart(filterCreatedAtStart).FilterCreatedAtEnd(filterCreatedAtEnd).Execute()

Get all application keys



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

    pageSize := int64(789) // int64 | Size for a given page. (optional) (default to 10)
    pageNumber := int64(789) // int64 | Specific page number to return. (optional) (default to 0)
    sort := "sort_example" // string | Application key attribute used to sort results. Sort order is ascending by default. In order to specify a descending sort, prefix the attribute with a minus sign. (optional) (default to "name")
    filter := "filter_example" // string | Filter application keys by the specified string. (optional)
    filterCreatedAtStart := "2020-11-24T18:46:21+00:00" // string | Only include application keys created on or after the specified date. (optional)
    filterCreatedAtEnd := "2020-11-24T18:46:21+00:00" // string | Only include application keys created on or before the specified date. (optional)

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.KeyManagementApi.ListApplicationKeys(ctx).PageSize(pageSize).PageNumber(pageNumber).Sort(sort).Filter(filter).FilterCreatedAtStart(filterCreatedAtStart).FilterCreatedAtEnd(filterCreatedAtEnd).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyManagementApi.ListApplicationKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApplicationKeys`: ListApplicationKeysResponse
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from KeyManagementApi.ListApplicationKeys:\n%v\n", response_content)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListApplicationKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int64** | Size for a given page. | [default to 10]
 **pageNumber** | **int64** | Specific page number to return. | [default to 0]
 **sort** | **string** | Application key attribute used to sort results. Sort order is ascending by default. In order to specify a descending sort, prefix the attribute with a minus sign. | [default to &quot;name&quot;]
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

> ListApplicationKeysResponse ListCurrentUserApplicationKeys(ctx).PageSize(pageSize).PageNumber(pageNumber).Sort(sort).Filter(filter).FilterCreatedAtStart(filterCreatedAtStart).FilterCreatedAtEnd(filterCreatedAtEnd).Execute()

Get all application keys owned by current user



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

    pageSize := int64(789) // int64 | Size for a given page. (optional) (default to 10)
    pageNumber := int64(789) // int64 | Specific page number to return. (optional) (default to 0)
    sort := "sort_example" // string | Application key attribute used to sort results. Sort order is ascending by default. In order to specify a descending sort, prefix the attribute with a minus sign. (optional) (default to "name")
    filter := "filter_example" // string | Filter application keys by the specified string. (optional)
    filterCreatedAtStart := "2020-11-24T18:46:21+00:00" // string | Only include application keys created on or after the specified date. (optional)
    filterCreatedAtEnd := "2020-11-24T18:46:21+00:00" // string | Only include application keys created on or before the specified date. (optional)

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.KeyManagementApi.ListCurrentUserApplicationKeys(ctx).PageSize(pageSize).PageNumber(pageNumber).Sort(sort).Filter(filter).FilterCreatedAtStart(filterCreatedAtStart).FilterCreatedAtEnd(filterCreatedAtEnd).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyManagementApi.ListCurrentUserApplicationKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCurrentUserApplicationKeys`: ListApplicationKeysResponse
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from KeyManagementApi.ListCurrentUserApplicationKeys:\n%v\n", response_content)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCurrentUserApplicationKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int64** | Size for a given page. | [default to 10]
 **pageNumber** | **int64** | Specific page number to return. | [default to 0]
 **sort** | **string** | Application key attribute used to sort results. Sort order is ascending by default. In order to specify a descending sort, prefix the attribute with a minus sign. | [default to &quot;name&quot;]
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

> APIKeyResponse UpdateAPIKey(ctx, apiKeyId).Body(body).Execute()

Edit an API key



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

    apiKeyId := "apiKeyId_example" // string | The ID of the API key.
    body := *datadog.NewAPIKeyUpdateRequest(*datadog.NewAPIKeyUpdateData(*datadog.NewAPIKeyUpdateAttributes("API Key for submitting metrics"), "00112233-4455-6677-8899-aabbccddeeff", datadog.APIKeysType("api_keys"))) // APIKeyUpdateRequest | 

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.KeyManagementApi.UpdateAPIKey(ctx, apiKeyId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyManagementApi.UpdateAPIKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAPIKey`: APIKeyResponse
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from KeyManagementApi.UpdateAPIKey:\n%v\n", response_content)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiKeyId** | **string** | The ID of the API key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAPIKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**APIKeyUpdateRequest**](APIKeyUpdateRequest.md) |  | 

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

> ApplicationKeyResponse UpdateApplicationKey(ctx, appKeyId).Body(body).Execute()

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

    appKeyId := "appKeyId_example" // string | The ID of the application key.
    body := *datadog.NewApplicationKeyUpdateRequest(*datadog.NewApplicationKeyUpdateData(*datadog.NewApplicationKeyUpdateAttributes("Application Key for submitting metrics"), "00112233-4455-6677-8899-aabbccddeeff", datadog.ApplicationKeysType("application_keys"))) // ApplicationKeyUpdateRequest | 

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.KeyManagementApi.UpdateApplicationKey(ctx, appKeyId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyManagementApi.UpdateApplicationKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApplicationKey`: ApplicationKeyResponse
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from KeyManagementApi.UpdateApplicationKey:\n%v\n", response_content)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKeyId** | **string** | The ID of the application key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApplicationKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ApplicationKeyUpdateRequest**](ApplicationKeyUpdateRequest.md) |  | 

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

> ApplicationKeyResponse UpdateCurrentUserApplicationKey(ctx, appKeyId).Body(body).Execute()

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

    appKeyId := "appKeyId_example" // string | The ID of the application key.
    body := *datadog.NewApplicationKeyUpdateRequest(*datadog.NewApplicationKeyUpdateData(*datadog.NewApplicationKeyUpdateAttributes("Application Key for submitting metrics"), "00112233-4455-6677-8899-aabbccddeeff", datadog.ApplicationKeysType("application_keys"))) // ApplicationKeyUpdateRequest | 

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.KeyManagementApi.UpdateCurrentUserApplicationKey(ctx, appKeyId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyManagementApi.UpdateCurrentUserApplicationKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCurrentUserApplicationKey`: ApplicationKeyResponse
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from KeyManagementApi.UpdateCurrentUserApplicationKey:\n%v\n", response_content)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKeyId** | **string** | The ID of the application key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCurrentUserApplicationKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ApplicationKeyUpdateRequest**](ApplicationKeyUpdateRequest.md) |  | 

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

