# \AWSIntegrationApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAWSAccount**](AWSIntegrationApi.md#CreateAWSAccount) | **Post** /api/v1/integration/aws | Create an AWS integration
[**CreateNewAWSExternalID**](AWSIntegrationApi.md#CreateNewAWSExternalID) | **Put** /api/v1/integration/aws/generate_new_external_id | Generate a new external ID
[**DeleteAWSAccount**](AWSIntegrationApi.md#DeleteAWSAccount) | **Delete** /api/v1/integration/aws | Delete an AWS integration
[**ListAWSAccounts**](AWSIntegrationApi.md#ListAWSAccounts) | **Get** /api/v1/integration/aws | List all AWS integrations
[**ListAvailableAWSNamespaces**](AWSIntegrationApi.md#ListAvailableAWSNamespaces) | **Get** /api/v1/integration/aws/available_namespace_rules | List namespace rules
[**UpdateAWSAccount**](AWSIntegrationApi.md#UpdateAWSAccount) | **Put** /api/v1/integration/aws | Update an AWS integration



## CreateAWSAccount

> AWSAccountCreateResponse CreateAWSAccount(ctx).Body(body).Execute()

Create an AWS integration



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAWSAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AWSAccount**](AWSAccount.md) | AWS Request Object | 

### Return type

[**AWSAccountCreateResponse**](AWSAccountCreateResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNewAWSExternalID

> AWSAccountCreateResponse CreateNewAWSExternalID(ctx).Body(body).Execute()

Generate a new external ID



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNewAWSExternalIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AWSAccount**](AWSAccount.md) | Your Datadog role delegation name. For more information about your AWS account Role name, see the [Datadog AWS integration configuration info](https://github.com/DataDog/documentation/blob/master/integrations/amazon_web_services/#installation). | 

### Return type

[**AWSAccountCreateResponse**](AWSAccountCreateResponse.md)

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

Delete an AWS integration



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAWSAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AWSAccount**](AWSAccount.md) | AWS request object | 

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


## ListAWSAccounts

> AWSAccountListResponse ListAWSAccounts(ctx).AccountId(accountId).RoleName(roleName).AccessKeyId(accessKeyId).Execute()

List all AWS integrations



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAWSAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** | Only return AWS accounts that matches this &#x60;account_id&#x60;. | 
 **roleName** | **string** | Only return AWS accounts that matches this role_name. | 
 **accessKeyId** | **string** | Only return AWS accounts that matches this &#x60;access_key_id&#x60;. | 

### Return type

[**AWSAccountListResponse**](AWSAccountListResponse.md)

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

List namespace rules



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

Update an AWS integration



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAWSAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AWSAccount**](AWSAccount.md) | AWS request object | 
 **accountId** | **string** | Only return AWS accounts that matches this &#x60;account_id&#x60;. | 
 **roleName** | **string** | Only return AWS accounts that match this &#x60;role_name&#x60;. Required if &#x60;account_id&#x60; is specified. | 
 **accessKeyId** | **string** | Only return AWS accounts that matches this &#x60;access_key_id&#x60;. Required if none of the other two options are specified. | 

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

