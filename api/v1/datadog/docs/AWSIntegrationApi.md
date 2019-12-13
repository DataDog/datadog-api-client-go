# \AWSIntegrationApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAWSAccount**](AWSIntegrationApi.md#CreateAWSAccount) | **Post** /api/v1/integration/aws | Create an AWS Account
[**DeleteAWSAccount**](AWSIntegrationApi.md#DeleteAWSAccount) | **Delete** /api/v1/integration/aws | Delete an AWS Account
[**GetAllAWSAccounts**](AWSIntegrationApi.md#GetAllAWSAccounts) | **Get** /api/v1/integration/aws | Get Installed AWS Accounts
[**UpdateAWSAccount**](AWSIntegrationApi.md#UpdateAWSAccount) | **Put** /api/v1/integration/aws | Update an AWS Account



## CreateAWSAccount

> AwsAccountCreateResponse CreateAWSAccount(ctx, awsAccount)

Create an AWS Account

### Overview
Create the AWS Account with the provided values
### Arguments
* **`account_id`** [*required*]: Your AWS Account ID without dashes. Consult the Datadog AWS
  integration to learn more about your AWS account ID.

* **`role_name`** [*required*]: Your Datadog role delegation name. For more information about you
  AWS account Role name, see the Datadog AWS integration configuration info.

* **`access_key_id`** [*optional*, *default* = **None**]: If your AWS account is a GovCloud or
  China account, enter the corresponding Access Key ID.

* **`filter_tags`** [*optional*, *default* = **None**]: The array of EC2 tags (in the form key:value)
  defines a filter that Datadog uses when collecting metrics from EC2. Wildcards, such as ?
  (for single characters) and * (for multiple characters) can also be used. Only hosts that match one
  of the defined tags will be imported into Datadog. The rest will be ignored. Host matching a given
  tag can also be excluded by adding ! before the tag.
  e.x. env:production,instance-type:c1.*,!region:us-east-1 For more information on EC2 tagging,
  see the AWS tagging documentation

* **`host_tags`** [*optional*, *default* = **None**]: Array of tags (in the form key:value) to add
  to all hosts and metrics reporting through this integration.

* **`account_specific_namespace_rules`** [*optional*, *default* = **None**]: An object (in the form
  {"namespace1":true/false, "namespace2":true/false}) that enables or disables metric collection for
  specific AWS namespaces for this AWS account only. A list of namespaces can be found at the
  /v1/integration/aws/available_namespace_rules endpoint.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**awsAccount** | [**AwsAccount**](AwsAccount.md)| AWS request object | 

### Return type

[**AwsAccountCreateResponse**](AWSAccountCreateResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAWSAccount

> interface{} DeleteAWSAccount(ctx, awsAccount)

Delete an AWS Account

### Overview
Delete the AWS Account matching the specified account_id and role_name parameters
### Arguments
* **`account_id`** [*required*, *default* = **None**]: Delete the AWS account that
  matches this account_id.

* **`role_name`** [*required*, *default* = **None**]: Delete the AWS account that
  matches this role_name.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**awsAccount** | [**AwsAccount**](AwsAccount.md)| AWS request object | 

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


## GetAllAWSAccounts

> AwsAccountListResponse GetAllAWSAccounts(ctx, optional)

Get Installed AWS Accounts

### Overview
Get All Installed AWS Accounts
### Arguments
* **`account_id`** [*optional*, *default* = **None**]: Only return AWS accounts that
  matches this account_id.

* **`role_name`** [*optional*, *default* = **None**]: Only return AWS accounts that
  matches this role_name.

* **`access_key_id`** [*optional*, *default* = **None**]: Only return AWS accounts that
  matches this access_key_id.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllAWSAccountsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAllAWSAccountsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **optional.String**| Only return AWS accounts that matches this account_id. | 
 **roleName** | **optional.String**| Only return AWS accounts that matches this role_name. | 
 **accessKeyId** | **optional.String**| Only return AWS accounts that matches this access_key_id. | 

### Return type

[**AwsAccountListResponse**](AWSAccountListResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAWSAccount

> interface{} UpdateAWSAccount(ctx, awsAccount, optional)

Update an AWS Account

### Overview
Update the AWS Account based on the provided values
### Arguments
* **`account_id`** [*required if role_name is specified*, *default* = **None**]: Only return AWS accounts that
  matches this account_id.

* **`role_name`** [*required if account_id is specified*, *default* = **None**]: Only return AWS accounts that
  matches this role_name.

* **`access_key_id`** [*required if none of the other two options are specified*, *default* = **None**]: Only return AWS accounts that
  matches this access_key_id.

### Payload
* **`account_id`** [*required*]: Your AWS Account ID without dashes. Consult the Datadog AWS
  integration to learn more about your AWS account ID.

* **`role_name`** [*required*]: Your Datadog role delegation name. For more information about you
  AWS account Role name, see the Datadog AWS integration configuration info.

* **`access_key_id`** [*optional*, *default* = **None**]: If your AWS account is a GovCloud or
  China account, enter the corresponding Access Key ID.

* **`filter_tags`** [*optional*, *default* = **None**]: The array of EC2 tags (in the form key:value)
  defines a filter that Datadog uses when collecting metrics from EC2. Wildcards, such as ?
  (for single characters) and * (for multiple characters) can also be used. Only hosts that match one
  of the defined tags will be imported into Datadog. The rest will be ignored. Host matching a given
  tag can also be excluded by adding ! before the tag.
  e.g. env:production,instance-type:c1.*,!region:us-east-1 For more information on EC2 tagging,
  see the AWS tagging documentation.

* **`host_tags`** [*optional*, *default* = **None**]: Array of tags (in the form key:value) to add
  to all hosts and metrics reporting through this integration.

* **`account_specific_namespace_rules`** [*optional*, *default* = **None**]: An object (in the form
  {"namespace1":true/false, "namespace2":true/false}) that enables or disables metric collection for
  specific AWS namespaces for this AWS account only. A list of namespaces can be found at the
  /v1/integration/aws/available_namespace_rules endpoint.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**awsAccount** | [**AwsAccount**](AwsAccount.md)| AWS request object | 
 **optional** | ***UpdateAWSAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateAWSAccountOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountId** | **optional.String**| Only return AWS accounts that matches this account_id. | 
 **roleName** | **optional.String**| Only return AWS accounts that matches this role_name. *It is required if account_id is specified.* | 
 **accessKeyId** | **optional.String**| Only return AWS accounts that matches this access_key_id. *It required if none of the other two options are specified.* | 

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

