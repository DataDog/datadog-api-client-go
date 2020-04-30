# \OrganizationsApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateChildOrg**](OrganizationsApi.md#CreateChildOrg) | **Post** /api/v1/org | Create a child organization
[**GetOrg**](OrganizationsApi.md#GetOrg) | **Get** /api/v1/org/{public_id} | Get organization information
[**ListOrgs**](OrganizationsApi.md#ListOrgs) | **Get** /api/v1/org | List your managed organizations
[**UpdateOrg**](OrganizationsApi.md#UpdateOrg) | **Put** /api/v1/org/{public_id} | Update your organization
[**UploadIdPForOrg**](OrganizationsApi.md#UploadIdPForOrg) | **Post** /api/v1/org/{public_id}/idp_metadata | Upload IdP metadata



## CreateChildOrg

> OrganizationCreateResponse CreateChildOrg(ctx).Body(body).Execute()

Create a child organization



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateChildOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OrganizationCreateBody**](OrganizationCreateBody.md) | Create child organization body request. | 

### Return type

[**OrganizationCreateResponse**](OrganizationCreateResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrg

> OrganizationResponse GetOrg(ctx, publicId).Execute()

Get organization information



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicId** | **string** | The &#x60;public_id&#x60; of the org you are operating within. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OrganizationResponse**](OrganizationResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgs

> OrganizationListResponse ListOrgs(ctx).Execute()

List your managed organizations



### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgsRequest struct via the builder pattern


### Return type

[**OrganizationListResponse**](OrganizationListResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrg

> OrganizationResponse UpdateOrg(ctx, publicId).Body(body).Execute()

Update your organization



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicId** | **string** | The &#x60;public_id&#x60; of the org you are operating within. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**Organization**](Organization.md) |  | 

### Return type

[**OrganizationResponse**](OrganizationResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadIdPForOrg

> IdpResponse UploadIdPForOrg(ctx, publicId).IdpFile(idpFile).Execute()

Upload IdP metadata



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicId** | **string** | The &#x60;public_id&#x60; of the org you are operating within. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadIdPForOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idpFile** | ***os.File** | The path to the XML metadata file you wish to upload. | 

### Return type

[**IdpResponse**](IdpResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

