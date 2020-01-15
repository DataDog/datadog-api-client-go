# \AWSIntegrationApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAWSAccount**](AWSIntegrationApi.md#CreateAWSAccount) | **Post** /api/v1/integration/aws | Create an AWS Account
[**DeleteAWSAccount**](AWSIntegrationApi.md#DeleteAWSAccount) | **Delete** /api/v1/integration/aws | Delete an AWS Account
[**GenerateNewAWSExternalID**](AWSIntegrationApi.md#GenerateNewAWSExternalID) | **Put** /api/v1/integration/aws/generate_new_external_id | Generate New External ID
[**GetAllAWSAccounts**](AWSIntegrationApi.md#GetAllAWSAccounts) | **Get** /api/v1/integration/aws | Get Installed AWS Accounts
[**ListAvailableAWSNamespaces**](AWSIntegrationApi.md#ListAvailableAWSNamespaces) | **Get** /api/v1/integration/aws/available_namespace_rules | List available namespaces.
[**UpdateAWSAccount**](AWSIntegrationApi.md#UpdateAWSAccount) | **Put** /api/v1/integration/aws | Update an AWS Account



## CreateAWSAccount

> AwsAccountCreateResponse CreateAWSAccount(ctx).Body(body).Execute()

Create an AWS Account



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAWSAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AwsAccount**](AwsAccount.md) | AWS request object | 

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

> interface{} DeleteAWSAccount(ctx).Body(body).Execute()

Delete an AWS Account



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAWSAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AwsAccount**](AwsAccount.md) | AWS request object | 

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


## GenerateNewAWSExternalID

> AwsAccountCreateResponse GenerateNewAWSExternalID(ctx).Body(body).Execute()

Generate New External ID



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateNewAWSExternalIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AwsAccount**](AwsAccount.md) | Generate New AWS External ID request object | 

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


## GetAllAWSAccounts

> AwsAccountListResponse GetAllAWSAccounts(ctx).AccountId(accountId).RoleName(roleName).AccessKeyId(accessKeyId).Execute()

Get Installed AWS Accounts



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllAWSAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** | Only return AWS accounts that matches this account_id. | 
 **roleName** | **string** | Only return AWS accounts that matches this role_name. | 
 **accessKeyId** | **string** | Only return AWS accounts that matches this access_key_id. | 

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


## ListAvailableAWSNamespaces

> []string ListAvailableAWSNamespaces(ctx).Execute()

List available namespaces.



### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAvailableAWSNamespacesRequest struct via the builder pattern


### Return type

**[]string**

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAWSAccount

> interface{} UpdateAWSAccount(ctx).Body(body).AccountId(accountId).RoleName(roleName).AccessKeyId(accessKeyId).Execute()

Update an AWS Account



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAWSAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AwsAccount**](AwsAccount.md) | AWS request object | 
 **accountId** | **string** | Only return AWS accounts that matches this account_id. | 
 **roleName** | **string** | Only return AWS accounts that matches this role_name. *It is required if account_id is specified.* | 
 **accessKeyId** | **string** | Only return AWS accounts that matches this access_key_id. *It required if none of the other two options are specified.* | 

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

