# \TagsApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddToHostTags**](TagsApi.md#AddToHostTags) | **Post** /api/v1/tags/hosts/{host_name} | Add tags to a host
[**GetAllHostTags**](TagsApi.md#GetAllHostTags) | **Get** /api/v1/tags/hosts | Get Tags
[**GetHostTags**](TagsApi.md#GetHostTags) | **Get** /api/v1/tags/hosts/{host_name} | Get host tags
[**RemoveHostTags**](TagsApi.md#RemoveHostTags) | **Delete** /api/v1/tags/hosts/{host_name} | Remove host tags
[**UpdateHostTags**](TagsApi.md#UpdateHostTags) | **Put** /api/v1/tags/hosts/{host_name} | Update host tags



## AddToHostTags

> HostTags AddToHostTags(ctx, hostName).Body(body).Source(source).Execute()

Add tags to a host



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostName** | **string** | This endpoint allows you to add new tags to a host, optionally specifying where the tags came from. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddToHostTagsRequest struct via the builder pattern


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


## GetAllHostTags

> TagToHosts GetAllHostTags(ctx).Source(source).Execute()

Get Tags



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllHostTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **source** | **string** | When specified, filters host list to those tags with the specified source | 

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


## GetHostTags

> HostTags GetHostTags(ctx, hostName).Source(source).Execute()

Get host tags



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostName** | **string** | When specified, filters list of tags to those tags with the specified source. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHostTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **source** | **string** | Source to filter | 

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


## RemoveHostTags

> RemoveHostTags(ctx, hostName).Source(source).Execute()

Remove host tags



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostName** | **string** | This endpoint allows you to remove all user-assigned tags for a single host. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveHostTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **source** | **string** | The source of the tags (e.g. chef, puppet). [Complete list of source attribute values](https://docs.datadoghq.com/integrations/faq/list-of-api-source-attribute-value) | 

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


## UpdateHostTags

> HostTags UpdateHostTags(ctx, hostName).Body(body).Source(source).Execute()

Update host tags



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

