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

> interface{} AzureUpdateHostFilters(ctx).Body(body).Execute()

Update the defined list of host filters for a given Datadog-Azure integration.



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAzureUpdateHostFiltersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AzureAccount**](AzureAccount.md) | Update a Datadog-Azure integrations host filters. | 

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

> interface{} CreateAzureIntegration(ctx).Body(body).Execute()

Add a Azure integration to your Datadog account.



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAzureIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AzureAccount**](AzureAccount.md) | Create a Datadog-Azure integration. | 

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

> interface{} DeleteAzureIntegration(ctx).Body(body).Execute()

Delete an Azure Integration from your Datadog account.



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAzureIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AzureAccount**](AzureAccount.md) | Delete a given Datadog-Azure integration. | 

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

> []AzureAccount ListAzureIntegration(ctx).Execute()

List configured Azure integrations.



### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAzureIntegrationRequest struct via the builder pattern


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

> interface{} UpdateAzureIntegration(ctx).Body(body).Execute()

Update an Azure integration to your Datadog account.



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAzureIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AzureAccount**](AzureAccount.md) | Update a Datadog-Azure integration. | 

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

