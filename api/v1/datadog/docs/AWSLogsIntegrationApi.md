# \AWSLogsIntegrationApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckAWSLogsLambdaAsync**](AWSLogsIntegrationApi.md#CheckAWSLogsLambdaAsync) | **Post** /api/v1/integration/aws/logs/check_async | Check that an AWS Lambda Function exists
[**CheckAWSLogsServicesAsync**](AWSLogsIntegrationApi.md#CheckAWSLogsServicesAsync) | **Post** /api/v1/integration/aws/logs/services_async | Check permissions for Log Services
[**CreateAWSLambdaARN**](AWSLogsIntegrationApi.md#CreateAWSLambdaARN) | **Post** /api/v1/integration/aws/logs | Add AWS Log Lambda ARN
[**DeleteAWSLambdaARN**](AWSLogsIntegrationApi.md#DeleteAWSLambdaARN) | **Delete** /api/v1/integration/aws/logs | Delete an AWS Logs integration
[**EnableAWSLogServices**](AWSLogsIntegrationApi.md#EnableAWSLogServices) | **Post** /api/v1/integration/aws/logs/services | Enable an AWS Logs integration
[**ListAWSLogsIntegrations**](AWSLogsIntegrationApi.md#ListAWSLogsIntegrations) | **Get** /api/v1/integration/aws/logs | List all AWS Logs Integrations
[**ListAWSLogsServices**](AWSLogsIntegrationApi.md#ListAWSLogsServices) | **Get** /api/v1/integration/aws/logs/services | Get list of AWS log ready services



## CheckAWSLogsLambdaAsync

> AWSLogsAsyncResponse CheckAWSLogsLambdaAsync(ctx).Body(body).Execute()

Check that an AWS Lambda Function exists



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckAWSLogsLambdaAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AWSAccountAndLambdaRequest**](AWSAccountAndLambdaRequest.md) | Check AWS Log Lambda Async request body. | 

### Return type

[**AWSLogsAsyncResponse**](AWSLogsAsyncResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckAWSLogsServicesAsync

> AWSLogsAsyncResponse CheckAWSLogsServicesAsync(ctx).Body(body).Execute()

Check permissions for Log Services



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckAWSLogsServicesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AWSLogsServicesRequest**](AWSLogsServicesRequest.md) | AWS Logs Async Services check request body. | 

### Return type

[**AWSLogsAsyncResponse**](AWSLogsAsyncResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAWSLambdaARN

> interface{} CreateAWSLambdaARN(ctx).Body(body).Execute()

Add AWS Log Lambda ARN



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAWSLambdaARNRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AWSAccountAndLambdaRequest**](AWSAccountAndLambdaRequest.md) | Check AWS Log Lambda Async request body. | 

### Return type

**interface{}**

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAWSLambdaARN

> interface{} DeleteAWSLambdaARN(ctx).Body(body).Execute()

Delete an AWS Logs integration



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAWSLambdaARNRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AWSAccountAndLambdaRequest**](AWSAccountAndLambdaRequest.md) | Check AWS Log Lambda Async request body. | 

### Return type

**interface{}**

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableAWSLogServices

> interface{} EnableAWSLogServices(ctx).Body(body).Execute()

Enable an AWS Logs integration



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnableAWSLogServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AWSLogsServicesRequest**](AWSLogsServicesRequest.md) | Enable AWS Log Services request object | 

### Return type

**interface{}**

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAWSLogsIntegrations

> []AWSLogsListResponse ListAWSLogsIntegrations(ctx).Execute()

List all AWS Logs Integrations



### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAWSLogsIntegrationsRequest struct via the builder pattern


### Return type

[**[]AWSLogsListResponse**](AWSLogsListResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAWSLogsServices

> []AWSLogsListServicesResponse ListAWSLogsServices(ctx).Execute()

Get list of AWS log ready services



### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAWSLogsServicesRequest struct via the builder pattern


### Return type

[**[]AWSLogsListServicesResponse**](AWSLogsListServicesResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

