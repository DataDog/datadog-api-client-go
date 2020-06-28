# \PagerDutyIntegrationApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePagerDutyIntegrationService**](PagerDutyIntegrationApi.md#CreatePagerDutyIntegrationService) | **Post** /api/v1/integration/pagerduty/configuration/services | Create a new service object
[**DeletePagerDutyIntegrationService**](PagerDutyIntegrationApi.md#DeletePagerDutyIntegrationService) | **Delete** /api/v1/integration/pagerduty/configuration/services/{service_name} | Delete a single service object
[**GetPagerDutyIntegrationService**](PagerDutyIntegrationApi.md#GetPagerDutyIntegrationService) | **Get** /api/v1/integration/pagerduty/configuration/services/{service_name} | Get a single service object
[**UpdatePagerDutyIntegrationService**](PagerDutyIntegrationApi.md#UpdatePagerDutyIntegrationService) | **Put** /api/v1/integration/pagerduty/configuration/services/{service_name} | Update a single service object



## CreatePagerDutyIntegrationService

> PagerDutyServiceName CreatePagerDutyIntegrationService(ctx, body)

Create a new service object

Create a new service object in the PagerDuty integration.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**PagerDutyService**](PagerDutyService.md)| Create a new service object request body. | 

### Return type

[**PagerDutyServiceName**](PagerDutyServiceName.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePagerDutyIntegrationService

> DeletePagerDutyIntegrationService(ctx, serviceName)

Delete a single service object

Delete a single service object in the Datadog-PagerDuty integration.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceName** | **string**| The service name | 

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


## GetPagerDutyIntegrationService

> PagerDutyServiceName GetPagerDutyIntegrationService(ctx, serviceName)

Get a single service object

Get service name in the Datadog-PagerDuty integration.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceName** | **string**| The service name. | 

### Return type

[**PagerDutyServiceName**](PagerDutyServiceName.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePagerDutyIntegrationService

> UpdatePagerDutyIntegrationService(ctx, serviceName, body)

Update a single service object

Update a single service object in the Datadog-PagerDuty integration.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceName** | **string**| The service name | 
**body** | [**PagerDutyServiceKey**](PagerDutyServiceKey.md)| Update an existing service object request body. | 

### Return type

 (empty response body)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

