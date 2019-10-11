# \UsersApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUser**](UsersApi.md#CreateUser) | **Post** /api/v1/user | Create user
[**DisableUser**](UsersApi.md#DisableUser) | **Delete** /api/v1/user/{user_handle} | Disable user
[**GetAllUsers**](UsersApi.md#GetAllUsers) | **Get** /api/v1/user | Get all users
[**GetUser**](UsersApi.md#GetUser) | **Get** /api/v1/user/{user_handle} | Get user
[**UpdateUser**](UsersApi.md#UpdateUser) | **Put** /api/v1/user/{user_handle} | Update user



## CreateUser

> UserCreateResponse CreateUser(ctx, userCreatePayload)
Create user

### Overview
Create a user for your organization.
### Arguments
* **`handle`** [*required*]: The user handle, must be a valid email.
* **`name`** [*optional*, *default*=**None**]: The name of the user.
* **`access_role`** [*optional*, *default*=**st**]: The access role of the user. Choose from:

  *  **st** (standard user),

  *  **adm** (admin user),

  *  **ro** (read-only user). *Note: users can be created with admin access role
     only with application keys belonging to administrators.*

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userCreatePayload** | [**UserCreatePayload**](UserCreatePayload.md)| User object that needs to be created | 

### Return type

[**UserCreateResponse**](UserCreateResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisableUser

> UserDisableResponse DisableUser(ctx, userHandle)
Disable user

### Overview
Delete a user from an organization.

**Note**: This endpoint can only be used with application keys belonging to administrators.
### Arguments
* **`id`** [*required*]: The handle of the user.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | [**string**](.md)| The handle of the user | 

### Return type

[**UserDisableResponse**](UserDisableResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: applcation/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllUsers

> UserGetAllResponse GetAllUsers(ctx, )
Get all users

### Overview
Get all users for your organization.
### Arguments
This endpoint takes no JSON argument.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**UserGetAllResponse**](UserGetAllResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUser

> UserGetResponse GetUser(ctx, userHandle)
Get user

### Overview
Get a user details.
### Arguments
* **`user_handle`** [*required*]: The handle of the user.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | [**string**](.md)| The id of the user | 

### Return type

[**UserGetResponse**](UserGetResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUser

> UserUpdateResponse UpdateUser(ctx, userHandle, userUpdatePayload)
Update user

### Overview
Update a user informations.

**Note**: It can only be used with application keys belonging to administrators.
### Arguments
* **`id`** [*required*]: The handle of the user.
* **`name`** [*optional*, *default*=**None**]: The new name of the user.
* **`email`** [*optional*, *default*=**None**]: The new email of the user.
* **`disabled`** [*optional*, *default*=**None**]: The new disabled status of the user.
* **`access_role`** [*optional*, *default*=**st**]: The access role of the user. Choose from:

  *  **st** (standard user)

  *  **adm** (admin user)

  *  **ro** (read-only user)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | [**string**](.md)| The id of the user | 
**userUpdatePayload** | [**UserUpdatePayload**](UserUpdatePayload.md)| Description of the update | 

### Return type

[**UserUpdateResponse**](UserUpdateResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

