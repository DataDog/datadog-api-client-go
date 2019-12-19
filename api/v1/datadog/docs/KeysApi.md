# \KeysApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAPIKey**](KeysApi.md#CreateAPIKey) | **Post** /api/v1/api_key | Create an API key with a given name.
[**CreateApplicationKey**](KeysApi.md#CreateApplicationKey) | **Post** /api/v1/application_key | Create an application key with a given name.
[**DeleteAPIKey**](KeysApi.md#DeleteAPIKey) | **Delete** /api/v1/api_key/{key} | Delete a given API key.
[**DeleteApplicationKey**](KeysApi.md#DeleteApplicationKey) | **Delete** /api/v1/application_key/{key} | Delete a given application key.
[**EditAPIKey**](KeysApi.md#EditAPIKey) | **Put** /api/v1/api_key/{key} | Edit an API key name.
[**EditApplicationKey**](KeysApi.md#EditApplicationKey) | **Put** /api/v1/application_key/{key} | Edit an application key name.
[**GetAPIKey**](KeysApi.md#GetAPIKey) | **Get** /api/v1/api_key/{key} | Get a given API key.
[**GetAllAPIKeys**](KeysApi.md#GetAllAPIKeys) | **Get** /api/v1/api_key | Get all API keys available for your account.
[**GetAllApplicationKeys**](KeysApi.md#GetAllApplicationKeys) | **Get** /api/v1/application_key | Get all application keys available for your account.
[**GetApplicationKey**](KeysApi.md#GetApplicationKey) | **Get** /api/v1/application_key/{key} | Get a given application key.



## CreateAPIKey

> ApiKeyResponse CreateAPIKey(ctx).ApiKey(apiKey).Execute()

Create an API key with a given name.



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAPIKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | [**ApiKey**](ApiKey.md) |  | 

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

> ApplicationKeyResponse CreateApplicationKey(ctx).ApplicationKey(applicationKey).Execute()

Create an application key with a given name.



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicationKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationKey** | [**ApplicationKey**](ApplicationKey.md) |  | 

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

Delete a given API key.



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | The specific API key you are working with | 

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

Delete a given application key.



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | The specific APP key you are working with | 

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


## EditAPIKey

> ApiKeyResponse EditAPIKey(ctx, key).ApiKey(apiKey).Execute()

Edit an API key name.



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | The specific API key you are working with | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditAPIKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiKey** | [**ApiKey**](ApiKey.md) |  | 

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


## EditApplicationKey

> ApplicationKeyResponse EditApplicationKey(ctx, key).ApplicationKey(applicationKey).Execute()

Edit an application key name.



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | The specific APP key you are working with | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditApplicationKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **applicationKey** | [**ApplicationKey**](ApplicationKey.md) |  | 

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


## GetAPIKey

> ApiKeyResponse GetAPIKey(ctx, key).Execute()

Get a given API key.



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | The specific API key you are working with | 

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


## GetAllAPIKeys

> ApiKeyListResponse GetAllAPIKeys(ctx).Execute()

Get all API keys available for your account.



### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllAPIKeysRequest struct via the builder pattern


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


## GetAllApplicationKeys

> ApplicationKeyListResponse GetAllApplicationKeys(ctx).Execute()

Get all application keys available for your account.



### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllApplicationKeysRequest struct via the builder pattern


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


## GetApplicationKey

> ApplicationKeyResponse GetApplicationKey(ctx, key).Execute()

Get a given application key.



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | The specific APP key you are working with | 

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

