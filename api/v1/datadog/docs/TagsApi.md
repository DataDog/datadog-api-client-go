# \TagsApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateHostTags**](TagsApi.md#CreateHostTags) | **Post** /api/v1/tags/hosts/{host_name} | Add tags to a host
[**DeleteHostTags**](TagsApi.md#DeleteHostTags) | **Delete** /api/v1/tags/hosts/{host_name} | Remove host tags
[**GetHostTags**](TagsApi.md#GetHostTags) | **Get** /api/v1/tags/hosts/{host_name} | Get host tags
[**ListHostTags**](TagsApi.md#ListHostTags) | **Get** /api/v1/tags/hosts | Get Tags
[**UpdateHostTags**](TagsApi.md#UpdateHostTags) | **Put** /api/v1/tags/hosts/{host_name} | Update host tags



## CreateHostTags

> HostTags CreateHostTags(ctx, hostName).Body(body).Source(source).Execute()

Add tags to a host



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
    hostName := "hostName_example" // string | This endpoint allows you to add new tags to a host, optionally specifying where the tags came from.
    body := datadog.HostTags{Host: "Host_example", Tags: []string{"Tags_example")} // HostTags | Update host tags request body.
    source := "source_example" // string | The source of the tags. [Complete list of source attribute values](https://docs.datadoghq.com/integrations/faq/list-of-api-source-attribute-value). (optional)

    configuration := datadog.NewConfiguration()
    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.TagsApi.CreateHostTags(context.Background(), hostName, body).Source(source).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.CreateHostTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateHostTags`: HostTags
    fmt.Fprintf(os.Stdout, "Response from `TagsApi.CreateHostTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostName** | **string** | This endpoint allows you to add new tags to a host, optionally specifying where the tags came from. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateHostTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**HostTags**](HostTags.md) | Update host tags request body. | 
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

> DeleteHostTags(ctx, hostName).Source(source).Execute()

Remove host tags



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
    hostName := "hostName_example" // string | This endpoint allows you to remove all user-assigned tags for a single host.
    source := "source_example" // string | The source of the tags (e.g. chef, puppet). [Complete list of source attribute values](https://docs.datadoghq.com/integrations/faq/list-of-api-source-attribute-value). (optional)

    configuration := datadog.NewConfiguration()
    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.TagsApi.DeleteHostTags(context.Background(), hostName).Source(source).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.DeleteHostTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostName** | **string** | This endpoint allows you to remove all user-assigned tags for a single host. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteHostTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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

> HostTags GetHostTags(ctx, hostName).Source(source).Execute()

Get host tags



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
    hostName := "hostName_example" // string | When specified, filters list of tags to those tags with the specified source.
    source := "source_example" // string | Source to filter. (optional)

    configuration := datadog.NewConfiguration()
    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.TagsApi.GetHostTags(context.Background(), hostName).Source(source).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.GetHostTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHostTags`: HostTags
    fmt.Fprintf(os.Stdout, "Response from `TagsApi.GetHostTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostName** | **string** | When specified, filters list of tags to those tags with the specified source. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHostTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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

> TagToHosts ListHostTags(ctx).Source(source).Execute()

Get Tags



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
    source := "source_example" // string | When specified, filters host list to those tags with the specified source. (optional)

    configuration := datadog.NewConfiguration()
    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.TagsApi.ListHostTags(context.Background(), ).Source(source).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.ListHostTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListHostTags`: TagToHosts
    fmt.Fprintf(os.Stdout, "Response from `TagsApi.ListHostTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListHostTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **source** | **string** | When specified, filters host list to those tags with the specified source. | 

### Return type

[**TagToHosts**](TagToHosts.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateHostTags

> HostTags UpdateHostTags(ctx, hostName).Body(body).Source(source).Execute()

Update host tags



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
    hostName := "hostName_example" // string | This endpoint allows you to update/replace all in an integration source with those supplied in the request.
    body := datadog.HostTags{Host: "Host_example", Tags: []string{"Tags_example")} // HostTags | Add tags to host
    source := "source_example" // string | The source of the tags (e.g. chef, puppet). [Complete list of source attribute values](https://docs.datadoghq.com/integrations/faq/list-of-api-source-attribute-value) (optional)

    configuration := datadog.NewConfiguration()
    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.TagsApi.UpdateHostTags(context.Background(), hostName, body).Source(source).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.UpdateHostTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateHostTags`: HostTags
    fmt.Fprintf(os.Stdout, "Response from `TagsApi.UpdateHostTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostName** | **string** | This endpoint allows you to update/replace all in an integration source with those supplied in the request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateHostTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**HostTags**](HostTags.md) | Add tags to host | 
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

