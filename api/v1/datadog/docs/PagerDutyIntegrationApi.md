# \PagerDutyIntegrationApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePagerDutyIntegration**](PagerDutyIntegrationApi.md#CreatePagerDutyIntegration) | **Put** /api/v1/integration/pagerduty | Create a PagerDuty integration
[**CreatePagerDutyIntegrationService**](PagerDutyIntegrationApi.md#CreatePagerDutyIntegrationService) | **Post** /api/v1/integration/pagerduty/configuration/services | Create a new service object
[**DeletePagerDutyIntegration**](PagerDutyIntegrationApi.md#DeletePagerDutyIntegration) | **Delete** /api/v1/integration/pagerduty | Delete a PagerDuty integration
[**DeletePagerDutyIntegrationService**](PagerDutyIntegrationApi.md#DeletePagerDutyIntegrationService) | **Delete** /api/v1/integration/pagerduty/configuration/services/{service_name} | Delete a single service object
[**GetPagerDutyIntegration**](PagerDutyIntegrationApi.md#GetPagerDutyIntegration) | **Get** /api/v1/integration/pagerduty | Get a PagerDuty integration
[**GetPagerDutyIntegrationService**](PagerDutyIntegrationApi.md#GetPagerDutyIntegrationService) | **Get** /api/v1/integration/pagerduty/configuration/services/{service_name} | Get a single service object
[**UpdatePagerDutyIntegration**](PagerDutyIntegrationApi.md#UpdatePagerDutyIntegration) | **Post** /api/v1/integration/pagerduty | Add new services and schedules
[**UpdatePagerDutyIntegrationService**](PagerDutyIntegrationApi.md#UpdatePagerDutyIntegrationService) | **Put** /api/v1/integration/pagerduty/configuration/services/{service_name} | Update a single service object



## CreatePagerDutyIntegration

> CreatePagerDutyIntegration(ctx).Body(body).Execute()

Create a PagerDuty integration



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePagerDutyIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PagerDutyIntegration**](PagerDutyIntegration.md) | Create Datadog-PagerDuty integration. | 

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


## CreatePagerDutyIntegrationService

> PagerDutyServiceName CreatePagerDutyIntegrationService(ctx).Body(body).Execute()

Create a new service object



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePagerDutyIntegrationServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PagerDutyService**](PagerDutyService.md) | Create a new service object in the Datadog-PagerDuty integration. | 

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


## DeletePagerDutyIntegration

> DeletePagerDutyIntegration(ctx).Execute()

Delete a PagerDuty integration



### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePagerDutyIntegrationRequest struct via the builder pattern


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


## DeletePagerDutyIntegrationService

> DeletePagerDutyIntegrationService(ctx, serviceName).Execute()

Delete a single service object



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceName** | **string** | The service name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePagerDutyIntegrationServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetPagerDutyIntegration

> PagerDutyIntegration GetPagerDutyIntegration(ctx).Execute()

Get a PagerDuty integration



### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPagerDutyIntegrationRequest struct via the builder pattern


### Return type

[**PagerDutyIntegration**](PagerDutyIntegration.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPagerDutyIntegrationService

> PagerDutyServiceName GetPagerDutyIntegrationService(ctx, serviceName).Execute()

Get a single service object



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceName** | **string** | The service name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPagerDutyIntegrationServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## UpdatePagerDutyIntegration

> UpdatePagerDutyIntegration(ctx).Body(body).Execute()

Add new services and schedules



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePagerDutyIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PagerDutyServicesAndSchedules**](PagerDutyServicesAndSchedules.md) | Update an existing Datadog-PagerDuty integration. | 

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


## UpdatePagerDutyIntegrationService

> UpdatePagerDutyIntegrationService(ctx, serviceName).Body(body).Execute()

Update a single service object



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceName** | **string** | The service name | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePagerDutyIntegrationServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**PagerDutyServiceKey**](PagerDutyServiceKey.md) | Update an existing service object in the Datadog-PagerDuty integration. | 

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

