# \KeysApi

All URIs are relative to *https://api.*

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

> ApiKeyResponse CreateAPIKey(ctx, optional)

Create an API key with a given name.

## Overview
Creates an API key
### ARGUMENTS
* **`name`** [*required*]: Name of your API key.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateAPIKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateAPIKeyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | [**optional.Interface of ApiKey**](ApiKey.md)|  | 

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

> ApplicationKeyResponse CreateApplicationKey(ctx, optional)

Create an application key with a given name.

## Overview
Create an application key with a given name.
### ARGUMENTS
* **`name`** [*required*]: Name of your application key.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateApplicationKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateApplicationKeyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationKey** | [**optional.Interface of ApplicationKey**](ApplicationKey.md)|  | 

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

> ApiKeyResponse DeleteAPIKey(ctx, key)

Delete a given API key.

## Overview
Delete a given API key.
### ARGUMENTS
This endpoint takes no JSON arguments.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string**| The specific API key you are working with | 

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

> ApplicationKeyResponse DeleteApplicationKey(ctx, key)

Delete a given application key.

## Overview
Delete a given application key.
### ARGUMENTS
This endpoint takes no JSON arguments.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string**| The specific APP key you are working with | 

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

> ApiKeyResponse EditAPIKey(ctx, key, optional)

Edit an API key name.

## Overview
Edit an API key name.
### ARGUMENTS
* **`name`** [*required*]: Name of your API key.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string**| The specific API key you are working with | 
 **optional** | ***EditAPIKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EditAPIKeyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiKey** | [**optional.Interface of ApiKey**](ApiKey.md)|  | 

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

> ApplicationKeyResponse EditApplicationKey(ctx, key, optional)

Edit an application key name.

## Overview
Edit an application key name.
### ARGUMENTS
* **`name`** [*required*]: Name of your application key.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string**| The specific APP key you are working with | 
 **optional** | ***EditApplicationKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EditApplicationKeyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **applicationKey** | [**optional.Interface of ApplicationKey**](ApplicationKey.md)|  | 

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

> ApiKeyResponse GetAPIKey(ctx, key)

Get a given API key.

## Overview
Get a given API key.
### ARGUMENTS
This endpoint takes no JSON arguments.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string**| The specific API key you are working with | 

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

> ApiKeyListResponse GetAllAPIKeys(ctx, )

Get all API keys available for your account.

## Overview
Get all API keys available for your account.
### ARGUMENTS
This endpoint takes no JSON arguments.

### Required Parameters

This endpoint does not need any parameter.

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

> ApplicationKeyListResponse GetAllApplicationKeys(ctx, )

Get all application keys available for your account.

## Overview
Get all application keys available for your account.
### ARGUMENTS
This endpoint takes no JSON arguments.

### Required Parameters

This endpoint does not need any parameter.

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

> ApplicationKeyResponse GetApplicationKey(ctx, key)

Get a given application key.

## Overview
Get a given application key.
### ARGUMENTS
This endpoint takes no JSON arguments.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string**| The specific APP key you are working with | 

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

