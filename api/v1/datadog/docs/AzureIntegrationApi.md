# \AzureIntegrationApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAzureIntegration**](AzureIntegrationApi.md#CreateAzureIntegration) | **Post** /api/v1/integration/azure | Create an Azure integration
[**DeleteAzureIntegration**](AzureIntegrationApi.md#DeleteAzureIntegration) | **Delete** /api/v1/integration/azure | Delete an Azure integration
[**ListAzureIntegration**](AzureIntegrationApi.md#ListAzureIntegration) | **Get** /api/v1/integration/azure | List all Azure integrations
[**UpdateAzureHostFilters**](AzureIntegrationApi.md#UpdateAzureHostFilters) | **Post** /api/v1/integration/azure/host_filters | Update Azure integration host filters
[**UpdateAzureIntegration**](AzureIntegrationApi.md#UpdateAzureIntegration) | **Put** /api/v1/integration/azure | Update an Azure integration



## CreateAzureIntegration

> map[string]interface{} CreateAzureIntegration(ctx, body)

Create an Azure integration

Create a Datadog-Azure integration.  Using the `POST` method updates your integration configuration by adding your new configuration to the existing one in your Datadog organization.  Using the `PUT` method updates your integration configuration by replacing your current configuration with the new one sent to your Datadog organization.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**AzureAccount**](AzureAccount.md)| Create a Datadog-Azure integration for your Datadog account request body. | 

### Return type

**map[string]interface{}**

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAzureIntegration

> map[string]interface{} DeleteAzureIntegration(ctx, body)

Delete an Azure integration

Delete a given Datadog-Azure integration from your Datadog account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**AzureAccount**](AzureAccount.md)| Delete a given Datadog-Azure integration request body. | 

### Return type

**map[string]interface{}**

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAzureIntegration

> []AzureAccount ListAzureIntegration(ctx, )

List all Azure integrations

List all Datadog-Azure integrations configured in your Datadog account.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]AzureAccount**](AzureAccount.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAzureHostFilters

> map[string]interface{} UpdateAzureHostFilters(ctx, body)

Update Azure integration host filters

Update the defined list of host filters for a given Datadog-Azure integration.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**AzureAccount**](AzureAccount.md)| Update a Datadog-Azure integration&#39;s host filters request body. | 

### Return type

**map[string]interface{}**

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAzureIntegration

> map[string]interface{} UpdateAzureIntegration(ctx, body)

Update an Azure integration

Update a Datadog-Azure integration. Requires an existing `tenant_name` and `client_id`. Any other fields supplied will overwrite existing values. To overwrite `tenant_name` or `client_id`, use `new_tenant_name` and `new_client_id`. To leave a field unchanged, do not supply that field in the payload.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**AzureAccount**](AzureAccount.md)| Update a Datadog-Azure integration request body. | 

### Return type

**map[string]interface{}**

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

