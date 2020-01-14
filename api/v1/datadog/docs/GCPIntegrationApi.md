# \GCPIntegrationApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGCPIntegration**](GCPIntegrationApi.md#CreateGCPIntegration) | **Post** /api/v1/integration/gcp | Add a GCP integration to your Datadog account.
[**DeleteGCPIntegration**](GCPIntegrationApi.md#DeleteGCPIntegration) | **Delete** /api/v1/integration/gcp | Delete a GCP Integration from your Datadog account.
[**ListGCPIntegration**](GCPIntegrationApi.md#ListGCPIntegration) | **Get** /api/v1/integration/gcp | List configured GCP integrations.
[**UpdateGCPIntegration**](GCPIntegrationApi.md#UpdateGCPIntegration) | **Put** /api/v1/integration/gcp | Update a GCP integration in your Datadog account.



## CreateGCPIntegration

> interface{} CreateGCPIntegration(ctx).Body(body).Execute()

Add a GCP integration to your Datadog account.



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGCPIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GcpAccount**](GcpAccount.md) | Create a Datadog-Azure integration. | 

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


## DeleteGCPIntegration

> interface{} DeleteGCPIntegration(ctx).Body(body).Execute()

Delete a GCP Integration from your Datadog account.



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGCPIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GcpAccount**](GcpAccount.md) | Delete a given Datadog-GCP integration. | 

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


## ListGCPIntegration

> []GcpAccount ListGCPIntegration(ctx).Execute()

List configured GCP integrations.



### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListGCPIntegrationRequest struct via the builder pattern


### Return type

[**[]GcpAccount**](GCPAccount.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGCPIntegration

> interface{} UpdateGCPIntegration(ctx).Body(body).Execute()

Update a GCP integration in your Datadog account.



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGCPIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GcpAccount**](GcpAccount.md) | Update a Datadog-GCP integration. | 

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

