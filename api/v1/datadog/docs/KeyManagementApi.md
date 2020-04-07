# \KeyManagementApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAPIKey**](KeyManagementApi.md#CreateAPIKey) | **Post** /api/v1/api_key | Create an API key
[**CreateApplicationKey**](KeyManagementApi.md#CreateApplicationKey) | **Post** /api/v1/application_key | Create an application key
[**DeleteAPIKey**](KeyManagementApi.md#DeleteAPIKey) | **Delete** /api/v1/api_key/{key} | Delete an API key
[**DeleteApplicationKey**](KeyManagementApi.md#DeleteApplicationKey) | **Delete** /api/v1/application_key/{key} | Delete an application key
[**GetAPIKey**](KeyManagementApi.md#GetAPIKey) | **Get** /api/v1/api_key/{key} | Get API key
[**GetApplicationKey**](KeyManagementApi.md#GetApplicationKey) | **Get** /api/v1/application_key/{key} | Get an application key
[**ListAPIKeys**](KeyManagementApi.md#ListAPIKeys) | **Get** /api/v1/api_key | Get all API keys
[**ListApplicationKeys**](KeyManagementApi.md#ListApplicationKeys) | **Get** /api/v1/application_key | Get all application keys
[**UpdateAPIKey**](KeyManagementApi.md#UpdateAPIKey) | **Put** /api/v1/api_key/{key} | Edit an API key
[**UpdateApplicationKey**](KeyManagementApi.md#UpdateApplicationKey) | **Put** /api/v1/application_key/{key} | Edit an application key



## CreateAPIKey

> ApiKeyResponse CreateAPIKey(ctx).Body(body).Execute()

Create an API key



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAPIKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ApiKey**](ApiKey.md) |  | 

### Return type

[**ApiKeyResponse**](ApiKeyResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateApplicationKey

> ApplicationKeyResponse CreateApplicationKey(ctx).Body(body).Execute()

Create an application key



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicationKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ApplicationKey**](ApplicationKey.md) |  | 

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

> ApiKeyResponse DeleteAPIKey(ctx, key).Execute()

Delete an API key



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | The specific API key you are working with. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAPIKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiKeyResponse**](ApiKeyResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApplicationKey

> ApplicationKeyResponse DeleteApplicationKey(ctx, key).Execute()

Delete an application key



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | The specific APP key you are working with. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationKeyRequest struct via the builder pattern


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


## GetAPIKey

> ApiKeyResponse GetAPIKey(ctx, key).Execute()

Get API key



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | The specific API key you are working with. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAPIKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiKeyResponse**](ApiKeyResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationKey

> ApplicationKeyResponse GetApplicationKey(ctx, key).Execute()

Get an application key



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | The specific APP key you are working with. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationKeyRequest struct via the builder pattern


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

> ApiKeyListResponse ListAPIKeys(ctx).Execute()

Get all API keys



### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAPIKeysRequest struct via the builder pattern


### Return type

[**ApiKeyListResponse**](ApiKeyListResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApplicationKeys

> ApplicationKeyListResponse ListApplicationKeys(ctx).Execute()

Get all application keys



### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListApplicationKeysRequest struct via the builder pattern


### Return type

[**ApplicationKeyListResponse**](ApplicationKeyListResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAPIKey

> ApiKeyResponse UpdateAPIKey(ctx, key).Body(body).Execute()

Edit an API key



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | The specific API key you are working with. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAPIKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ApiKey**](ApiKey.md) |  | 

### Return type

[**ApiKeyResponse**](ApiKeyResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApplicationKey

> ApplicationKeyResponse UpdateApplicationKey(ctx, key).Body(body).Execute()

Edit an application key



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | The specific APP key you are working with. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApplicationKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ApplicationKey**](ApplicationKey.md) |  | 

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

