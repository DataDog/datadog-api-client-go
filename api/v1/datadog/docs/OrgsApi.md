# \OrgsApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateChildOrg**](OrgsApi.md#CreateChildOrg) | **Post** /api/v1/org | Create child-organization.
[**GetOrg**](OrgsApi.md#GetOrg) | **Get** /api/v1/org | Get the organization
[**UpdateOrg**](OrgsApi.md#UpdateOrg) | **Put** /api/v1/org/{public_id} | Update the organization
[**UploadIdPForOrg**](OrgsApi.md#UploadIdPForOrg) | **Post** /api/v1/org/{public_id}/idp_metadata | Upload IdP metadata



## CreateChildOrg

> OrgCreateResponse CreateChildOrg(ctx, orgCreateBody)

Create child-organization.

## Overview
This endpoint requires the [multi-org account](https://docs.datadoghq.com/account_management/multi_organization/) feature and must be enabled by [contacting support](https://docs.datadoghq.com/help/).
### ARGUMENTS
* **`name`** [*required*]: The name of the new child-organization, limited to 32 characters.
* **`subscription`** [*required*]: A JSON array of subscription type. Types available are **trial**, **free**, and **pro**.
* **`billing`** [*required*]: A JSON array of billing type. Note that only **parent_billing** is supported.

Once a new child-organization is created, you can interact with it by using the **org.public_id**, **api_key.key**, and **application_key.hash** provided in the response.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgCreateBody** | [**OrgCreateBody**](OrgCreateBody.md)| Org object that needs to be created | 

### Return type

[**OrgCreateResponse**](OrgCreateResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrg

> OrgListResponse GetOrg(ctx, )

Get the organization

## Overview
Gets information about your organization

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**OrgListResponse**](OrgListResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrg

> OrgResponse UpdateOrg(ctx, publicId, optional)

Update the organization

## Overview
Updates the organization
### ARGUMENTS
* **`name`** [*optional*]: The organization name.

* **`settings`** [*optional*]: A JSON array of settings. Settings include:

  * **`saml`**: Set the boolean property **enabled** to enable or disable single sign on with SAML. See the [SAML documentation](https://docs.datadoghq.com/account_management/saml) for more information about all SAML settings.

  * **`saml_idp_initiated_login`**: has one property **enabled** (boolean).

  * **`saml_strict_mode`**: has one property **enabled** (boolean).

  * **`saml_autocreate_users_domains`**: has two properties: **enabled** (boolean) and **domains** which is a list of domains without the @ symbol.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicId** | **string**| The public_id of the org you are operating with | 
 **optional** | ***UpdateOrgOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateOrgOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **org** | [**optional.Interface of Org**](Org.md)|  | 

### Return type

[**OrgResponse**](OrgResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadIdPForOrg

> IdpResponse UploadIdPForOrg(ctx, publicId, idpFile)

Upload IdP metadata

## Overview
There are a couple of options for updating the Identity Provider (IdP) metadata from your SAML IdP.
* **Multipart Form-Data**: Post the IdP metadata file using a form post.
### Multipart Form-Data
#### Headers
* **`Content-Type: multipart/form-data`**
#### Arguments
* **`public_id`** [*required*]: The public id of the org you want to update metadata for.
### MultiPart Form Data Body
* The encoded data for the IDP settings to upload
#### Headers
* **`Content-Type: multipart/form-data`**
#### Arguments
* The body must contain the contents of your IdP metadata XML file.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicId** | **string**| The public_id of the org you are operating with | 
**idpFile** | ***os.File*****os.File**| The path to the XML metadata file you wish to upload. | 

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

