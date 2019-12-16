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

> AwsLogsAsyncResponse AWSLogsCheckLambdaAsync(ctx, awsAccountAndLambdaInput)

Check function to see if a lambda_arn exists within an account.

### Overview
Check function to see if a lambda_arn exists within an account. This sends a job on our side if it does not exist, then immediately returns the status of that job. Subsequent requests will always repeat the above, so this endpoint can be polled intermittently instead of blocking.
- Returns a status of 'created' when it's checking if the Lambda exists in the account.
- Returns a status of 'waiting' while checking.
- Returns a status of 'checked and ok' if the Lambda exists.
- Returns a status of 'error' if the Lambda does not exist.
### Arguments
* **`account_id`** [*required*, *default* = **None**]: Your AWS Account ID without dashes.
* **`lambda_arn`** [*required*, *default* = **None**]: ARN of the Lambda to be checked.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**awsAccountAndLambdaInput** | [**AwsAccountAndLambdaInput**](AwsAccountAndLambdaInput.md)| Check AWS Log Lambda Async request body. | 

### Return type

[**AwsLogsAsyncResponse**](AWSLogsAsyncResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AWSLogsCheckServicesAsync

> AwsLogsAsyncResponse AWSLogsCheckServicesAsync(ctx, awsLogsServicesInput)

Asynchronous check for permissions for AWS log lambda config.

### Overview
Test if permissions are present to add log-forwarding triggers for the given services + AWS account. Input is the same as for EnableAWSLogServices. Done async, so can be repeatedly polled in a non-blocking fashion until the async request completes
- Returns a status of 'created' when it's checking if the permissions exists in the AWS account.
- Returns a status of 'waiting' while checking.
- Returns a status of 'checked and ok' if the Lambda exists.
- Returns a status of 'error' if the Lambda does not exist.
### Arguments
* **`account_id`** [*required*, *default* = **None**]: Your AWS Account ID without dashes.
* **`services`** [*required*, *default* = **None**]: Array of services IDs set to enable automatic log collection. Discover the list of available services with the Get list of AWS log ready services API endpoint

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**awsLogsServicesInput** | [**AwsLogsServicesInput**](AwsLogsServicesInput.md)| AWS Logs Async Services check request body | 

### Return type

[**AwsLogsAsyncResponse**](AWSLogsAsyncResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AWSLogsList

> []AwsLogsListOutput AWSLogsList(ctx, )

List configured AWS log integrations.

### Overview
List all Datadog-AWS Logs integrations configured in your Datadog account.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]AwsLogsListOutput**](AWSLogsListOutput.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AWSLogsServicesList

> []AwsLogsListServicesOutput AWSLogsServicesList(ctx, )

Get list of AWS log ready services.

### Overview
Get the list of current AWS services that Datadog offers automatic log collection.  Use returned service IDs with the services parameter for the Enable an AWS service log collection API endpoint.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]AwsLogsListServicesOutput**](AWSLogsListServicesOutput.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddAWSLambdaARN

> interface{} AddAWSLambdaARN(ctx, awsAccountAndLambdaInput)

Add a AWS Lambda ARN to your Datadog account.

### Overview
Attach the Lambda ARN of the Lambda created for the Datadog-AWS log collection to your AWS account ID to enable log collection.
### Arguments
* **`account_id`** [*required*, *default* = **None**]: Your AWS Account ID without dashes.
* **`lambda_arn`** [*required*, *default* = **None**]: ARN of the Datadog Lambda created during the Datadog-Amazon Web services Log collection setup.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**awsAccountAndLambdaInput** | [**AwsAccountAndLambdaInput**](AwsAccountAndLambdaInput.md)| Check AWS Log Lambda Async request body. | 

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

> interface{} DeleteAWSLambdaARN(ctx, awsAccountAndLambdaInput)

Delete a AWS Lambda ARN from your Datadog account.

### Overview
Delete a Lambda ARN of a Lambda created for the Datadog-AWS log collection in your Datadog account.
### Arguments
* **`account_id`** [*required*, *default* = **None**]: Your AWS Account ID without dashes.
* **`lambda_arn`** [*required*, *default* = **None**]: ARN of the Lambda to be deleted.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**awsAccountAndLambdaInput** | [**AwsAccountAndLambdaInput**](AwsAccountAndLambdaInput.md)| Check AWS Log Lambda Async request body. | 

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

> interface{} EnableAWSLogServices(ctx, awsLogsServicesInput)

Enable Automatic Log collection for your AWS services.

### Overview
Enable automatic log collection for a list of services. This should be run after running 'AddAWSLambdaARN' to save the config.
### Arguments
* **`account_id`** [*required*, *default* = **None**]: Your AWS Account ID without dashes.
* **`services`** [*required*, *default* = **None**]: Array of services IDs set to enable automatic log collection. Discover the list of available services with the Get list of AWS log ready services API endpoint

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**awsLogsServicesInput** | [**AwsLogsServicesInput**](AwsLogsServicesInput.md)| Enable AWS Log Services request object | 

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

