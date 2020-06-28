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

> HostTags CreateHostTags(ctx, hostName, body, optional)

Add tags to a host

This endpoint allows you to add new tags to a host, optionally specifying where these tags come from.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostName** | **string**| This endpoint allows you to add new tags to a host, optionally specifying where the tags came from. | 
**body** | [**HostTags**](HostTags.md)| Update host tags request body. | 
 **optional** | ***CreateHostTagsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateHostTagsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **source** | **optional.String**| The source of the tags. [Complete list of source attribute values](https://docs.datadoghq.com/integrations/faq/list-of-api-source-attribute-value). | 

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

> DeleteHostTags(ctx, hostName, optional)

Remove host tags

This endpoint allows you to remove all user-assigned tags for a single host.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostName** | **string**| This endpoint allows you to remove all user-assigned tags for a single host. | 
 **optional** | ***DeleteHostTagsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteHostTagsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **source** | **optional.String**| The source of the tags (e.g. chef, puppet). [Complete list of source attribute values](https://docs.datadoghq.com/integrations/faq/list-of-api-source-attribute-value). | 

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

> HostTags GetHostTags(ctx, hostName, optional)

Get host tags

Return the list of tags that apply to a given host.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostName** | **string**| When specified, filters list of tags to those tags with the specified source. | 
 **optional** | ***GetHostTagsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetHostTagsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **source** | **optional.String**| Source to filter. | 

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

> TagToHosts ListHostTags(ctx, optional)

Get Tags

Return a mapping of tags to hosts for your whole infrastructure.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListHostTagsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListHostTagsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **source** | **optional.String**| When specified, filters host list to those tags with the specified source. | 

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

> HostTags UpdateHostTags(ctx, hostName, body, optional)

Update host tags

This endpoint allows you to update/replace all tags in an integration source with those supplied in the request.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostName** | **string**| This endpoint allows you to update/replace all in an integration source with those supplied in the request. | 
**body** | [**HostTags**](HostTags.md)| Add tags to host | 
 **optional** | ***UpdateHostTagsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateHostTagsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **source** | **optional.String**| The source of the tags (e.g. chef, puppet). [Complete list of source attribute values](https://docs.datadoghq.com/integrations/faq/list-of-api-source-attribute-value) | 

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

