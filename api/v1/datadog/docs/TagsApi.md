# TagsApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------ | ------------ | ------------
[**CreateHostTags**](TagsApi.md#CreateHostTags) | **Post** /api/v1/tags/hosts/{host_name} | Add tags to a host
[**DeleteHostTags**](TagsApi.md#DeleteHostTags) | **Delete** /api/v1/tags/hosts/{host_name} | Remove host tags
[**GetHostTags**](TagsApi.md#GetHostTags) | **Get** /api/v1/tags/hosts/{host_name} | Get host tags
[**ListHostTags**](TagsApi.md#ListHostTags) | **Get** /api/v1/tags/hosts | Get Tags
[**UpdateHostTags**](TagsApi.md#UpdateHostTags) | **Put** /api/v1/tags/hosts/{host_name} | Update host tags



## CreateHostTags

> HostTags CreateHostTags(ctx, hostName, body, datadog.CreateHostTagsOptionalParameters{})

This endpoint allows you to add new tags to a host,
optionally specifying where these tags come from.

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

    hostName := "hostName_example" // string | This endpoint allows you to add new tags to a host, optionally specifying where the tags came from.
    body := *datadog.NewHostTags() // HostTags | Update host tags request body.
    source := "chef" // string | The source of the tags. [Complete list of source attribute values](https://docs.datadoghq.com/integrations/faq/list-of-api-source-attribute-value). (optional)
    optionalParams := datadog.CreateHostTagsOptionalParameters{
        Source: &source,
    }

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.TagsApi.CreateHostTags(ctx, hostName, body, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.CreateHostTags`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateHostTags`: HostTags
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from TagsApi.CreateHostTags:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**hostName** | **string** | This endpoint allows you to add new tags to a host, optionally specifying where the tags came from. |  |
**body** | [**HostTags**](HostTags.md) | Update host tags request body. | 


### Optional Parameters


Other parameters are passed through a pointer to a CreateHostTagsOptionalParameters struct.


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**source** | **string** | The source of the tags. [Complete list of source attribute values](https://docs.datadoghq.com/integrations/faq/list-of-api-source-attribute-value). | 

### Return type

[**HostTags**](HostTags.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteHostTags

> DeleteHostTags(ctx, hostName, datadog.DeleteHostTagsOptionalParameters{})

This endpoint allows you to remove all user-assigned tags
for a single host.

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
    ctx := datadog.NewDefaultContext(context.Background())

    hostName := "hostName_example" // string | This endpoint allows you to remove all user-assigned tags for a single host.
    source := "source_example" // string | The source of the tags (e.g. chef, puppet). [Complete list of source attribute values](https://docs.datadoghq.com/integrations/faq/list-of-api-source-attribute-value). (optional)
    optionalParams := datadog.DeleteHostTagsOptionalParameters{
        Source: &source,
    }

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    r, err := apiClient.TagsApi.DeleteHostTags(ctx, hostName, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.DeleteHostTags`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**hostName** | **string** | This endpoint allows you to remove all user-assigned tags for a single host. | 


### Optional Parameters


Other parameters are passed through a pointer to a DeleteHostTagsOptionalParameters struct.


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**source** | **string** | The source of the tags (e.g. chef, puppet). [Complete list of source attribute values](https://docs.datadoghq.com/integrations/faq/list-of-api-source-attribute-value). | 

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


## GetHostTags

> HostTags GetHostTags(ctx, hostName, datadog.GetHostTagsOptionalParameters{})

Return the list of tags that apply to a given host.

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

    hostName := "hostName_example" // string | When specified, filters list of tags to those tags with the specified source.
    source := "source_example" // string | Source to filter. (optional)
    optionalParams := datadog.GetHostTagsOptionalParameters{
        Source: &source,
    }

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.TagsApi.GetHostTags(ctx, hostName, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.GetHostTags`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHostTags`: HostTags
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from TagsApi.GetHostTags:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**hostName** | **string** | When specified, filters list of tags to those tags with the specified source. | 


### Optional Parameters


Other parameters are passed through a pointer to a GetHostTagsOptionalParameters struct.


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**source** | **string** | Source to filter. | 

### Return type

[**HostTags**](HostTags.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHostTags

> TagToHosts ListHostTags(ctx, datadog.ListHostTagsOptionalParameters{})

Return a mapping of tags to hosts for your whole infrastructure.

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

    source := "source_example" // string | When specified, filters host list to those tags with the specified source. (optional)
    optionalParams := datadog.ListHostTagsOptionalParameters{
        Source: &source,
    }

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.TagsApi.ListHostTags(ctx, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.ListHostTags`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListHostTags`: TagToHosts
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from TagsApi.ListHostTags:\n%s\n", responseContent)
}
```

### Required Parameters




### Optional Parameters


Other parameters are passed through a pointer to a ListHostTagsOptionalParameters struct.


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**source** | **string** | When specified, filters host list to those tags with the specified source. | 

### Return type

[**TagToHosts**](TagToHosts.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateHostTags

> HostTags UpdateHostTags(ctx, hostName, body, datadog.UpdateHostTagsOptionalParameters{})

This endpoint allows you to update/replace all tags in
an integration source with those supplied in the request.

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

    hostName := "hostName_example" // string | This endpoint allows you to update/replace all in an integration source with those supplied in the request.
    body := *datadog.NewHostTags() // HostTags | Add tags to host
    source := "source_example" // string | The source of the tags (e.g. chef, puppet). [Complete list of source attribute values](https://docs.datadoghq.com/integrations/faq/list-of-api-source-attribute-value) (optional)
    optionalParams := datadog.UpdateHostTagsOptionalParameters{
        Source: &source,
    }

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.TagsApi.UpdateHostTags(ctx, hostName, body, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.UpdateHostTags`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateHostTags`: HostTags
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from TagsApi.UpdateHostTags:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**hostName** | **string** | This endpoint allows you to update/replace all in an integration source with those supplied in the request. |  |
**body** | [**HostTags**](HostTags.md) | Add tags to host | 


### Optional Parameters


Other parameters are passed through a pointer to a UpdateHostTagsOptionalParameters struct.


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**source** | **string** | The source of the tags (e.g. chef, puppet). [Complete list of source attribute values](https://docs.datadoghq.com/integrations/faq/list-of-api-source-attribute-value) | 

### Return type

[**HostTags**](HostTags.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

