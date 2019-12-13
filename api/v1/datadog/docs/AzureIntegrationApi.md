# \AzureIntegrationApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AzureUpdateHostFilters**](AzureIntegrationApi.md#AzureUpdateHostFilters) | **Post** /api/v1/integration/azure/host_filters | Update the defined list of host filters for a given Datadog-Azure integration.
[**CreateAzureIntegration**](AzureIntegrationApi.md#CreateAzureIntegration) | **Post** /api/v1/integration/azure | Add a Azure integration to your Datadog account.
[**DeleteAzureIntegration**](AzureIntegrationApi.md#DeleteAzureIntegration) | **Delete** /api/v1/integration/azure | Delete an Azure Integration from your Datadog account.
[**ListAzureIntegration**](AzureIntegrationApi.md#ListAzureIntegration) | **Get** /api/v1/integration/azure | List configured Azure integrations.
[**UpdateAzureIntegration**](AzureIntegrationApi.md#UpdateAzureIntegration) | **Put** /api/v1/integration/azure | Update an Azure integration to your Datadog account.



## AzureUpdateHostFilters

> interface{} AzureUpdateHostFilters(ctx, azureAccount)

Update the defined list of host filters for a given Datadog-Azure integration.

### Overview
Update the defined list of host filters for a given Datadog-Azure integration.
### Arguments
* **`tenant_name`** [*required*, *default* = **None**]: Your Azure Active Directory ID.
* **`client_id`** [*required*, *default* = **None**]: Your Azure web application ID.
* **`host_filters`** [*required*, *default* = **None**]: Limit the Azure instances that are pulled into Datadog by using tags. Only hosts that match one of the defined tags are imported into Datadog.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**azureAccount** | [**AzureAccount**](AzureAccount.md)| Update a Datadog-Azure integrations host filters. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAzureIntegration

> interface{} CreateAzureIntegration(ctx, azureAccount)

Add a Azure integration to your Datadog account.

### Overview
Create a Datadog-Azure integration.
### Arguments
* **`tenant_name`** [*required*, *default* = **None**]: Your Azure Active Directory ID.
* **`client_id`** [*required*, *default* = **None**]: Your Azure web application ID.
* **`client_secret`** [*required*, *default* = **None**]: Your Azure web application secret key.
* **`host_filters`** [*optional*, *default* = **None**]: Limit the Azure instances that are pulled into Datadog by using tags. Only hosts that match one of the defined tags are imported into Datadog.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**azureAccount** | [**AzureAccount**](AzureAccount.md)| Create a Datadog-Azure integration. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAzureIntegration

> interface{} DeleteAzureIntegration(ctx, azureAccount)

Delete an Azure Integration from your Datadog account.

### Overview
Delete a given Datadog-Azure integration.
### Arguments
* **`tenant_name`** [*required*, *default* = **None**]: Your Azure Active Directory ID.
* **`client_id`** [*required*, *default* = **None**]: Your Azure web application ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**azureAccount** | [**AzureAccount**](AzureAccount.md)| Delete a given Datadog-Azure integration. | 

### Return type

[**interface{}**](interface{}.md)

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

List configured Azure integrations.

### Overview
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


## UpdateAzureIntegration

> interface{} UpdateAzureIntegration(ctx, azureAccount)

Update an Azure integration to your Datadog account.

### Overview
Update an Datadog-Azure integration. Requires an existing tenant_name and client_id. Any other fields supplied will overwrite existing values. To overwrite tenant_name or client_id, use new_tenant_name and new_client_id. To leave a field unchanged, do not supply that field in the payload.
### Arguments
* **`tenant_name`** [*required*, *default* = **None**]: Your existing Azure Active Directory ID.
* **`new_tenant_name`** [*optional*, *default* = **None**]: Your new Azure Active Directory ID.
* **`client_id`** [*required*, *default* = **None**]: Your existing Azure web application ID.
* **`new_client_id`** [*optional*, *default* = **None**]: Your new Azure web application ID.
* **`client_secret`** [*optional*, *default* = **None**]: Your Azure web application secret key.
* **`host_filters`** [*optional*, *default* = **None**]: Limit the Azure instances that are pulled into Datadog by using tags. Only hosts that match one of the defined tags are imported into Datadog.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**azureAccount** | [**AzureAccount**](AzureAccount.md)| Update a Datadog-Azure integration. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

