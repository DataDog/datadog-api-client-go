# \PagerDutyIntegrationApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePagerDutyIntegration**](PagerDutyIntegrationApi.md#CreatePagerDutyIntegration) | **Put** /api/v1/integration/pagerduty | Create a PagerDuty integration
[**DeletePagerDutyIntegration**](PagerDutyIntegrationApi.md#DeletePagerDutyIntegration) | **Delete** /api/v1/integration/pagerduty | Delete a PagerDuty integration
[**GetPagerDutyIntegration**](PagerDutyIntegrationApi.md#GetPagerDutyIntegration) | **Get** /api/v1/integration/pagerduty | Get a PagerDuty integration
[**UpdatePagerDutyIntegration**](PagerDutyIntegrationApi.md#UpdatePagerDutyIntegration) | **Post** /api/v1/integration/pagerduty | Add new services and schedules



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

