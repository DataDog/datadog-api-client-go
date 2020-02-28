# \AWSLogsIntegrationApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AWSLogsCheckLambdaAsync**](AWSLogsIntegrationApi.md#AWSLogsCheckLambdaAsync) | **Post** /api/v1/integration/aws/logs/check_async | Check function to see if a lambda_arn exists within an account.
[**AWSLogsCheckServicesAsync**](AWSLogsIntegrationApi.md#AWSLogsCheckServicesAsync) | **Post** /api/v1/integration/aws/logs/services_async | Asynchronous check for permissions for AWS log lambda config.
[**AWSLogsList**](AWSLogsIntegrationApi.md#AWSLogsList) | **Get** /api/v1/integration/aws/logs | List configured AWS log integrations.
[**AWSLogsServicesList**](AWSLogsIntegrationApi.md#AWSLogsServicesList) | **Get** /api/v1/integration/aws/logs/services | Get list of AWS log ready services.
[**AddAWSLambdaARN**](AWSLogsIntegrationApi.md#AddAWSLambdaARN) | **Post** /api/v1/integration/aws/logs | Add a AWS Lambda ARN to your Datadog account.
[**DeleteAWSLambdaARN**](AWSLogsIntegrationApi.md#DeleteAWSLambdaARN) | **Delete** /api/v1/integration/aws/logs | Delete a AWS Lambda ARN from your Datadog account.
[**EnableAWSLogServices**](AWSLogsIntegrationApi.md#EnableAWSLogServices) | **Post** /api/v1/integration/aws/logs/services | Enable Automatic Log collection for your AWS services.



## AWSLogsCheckLambdaAsync

> AWSLogsAsyncResponse AWSLogsCheckLambdaAsync(ctx).Body(body).Execute()

Check function to see if a lambda_arn exists within an account.



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAWSLogsCheckLambdaAsyncRequest struct via the builder pattern


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


## AWSLogsCheckServicesAsync

> AWSLogsAsyncResponse AWSLogsCheckServicesAsync(ctx).Body(body).Execute()

Asynchronous check for permissions for AWS log lambda config.



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAWSLogsCheckServicesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AWSLogsServicesRequest**](AWSLogsServicesRequest.md) | AWS Logs Async Services check request body | 

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


## AWSLogsList

> []AWSLogsListResponse AWSLogsList(ctx).Execute()

List configured AWS log integrations.



### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAWSLogsListRequest struct via the builder pattern


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


## AWSLogsServicesList

> []AWSLogsListServicesResponse AWSLogsServicesList(ctx).Execute()

Get list of AWS log ready services.



### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAWSLogsServicesListRequest struct via the builder pattern


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


## AddAWSLambdaARN

> interface{} AddAWSLambdaARN(ctx).Body(body).Execute()

Add a AWS Lambda ARN to your Datadog account.



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddAWSLambdaARNRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AWSAccountAndLambdaRequest**](AWSAccountAndLambdaRequest.md) | Check AWS Log Lambda Async request body. | 

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


## DeleteAWSLambdaARN

> interface{} DeleteAWSLambdaARN(ctx).Body(body).Execute()

Delete a AWS Lambda ARN from your Datadog account.



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAWSLambdaARNRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AWSAccountAndLambdaRequest**](AWSAccountAndLambdaRequest.md) | Check AWS Log Lambda Async request body. | 

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


## EnableAWSLogServices

> interface{} EnableAWSLogServices(ctx).Body(body).Execute()

Enable Automatic Log collection for your AWS services.



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnableAWSLogServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AWSLogsServicesRequest**](AWSLogsServicesRequest.md) | Enable AWS Log Services request object | 

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

