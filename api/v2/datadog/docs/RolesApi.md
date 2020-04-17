# \RolesApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPermissionToRole**](RolesApi.md#AddPermissionToRole) | **Post** /api/v2/roles/{role_id}/permissions | Grant permission to a role
[**AddUserToRole**](RolesApi.md#AddUserToRole) | **Post** /api/v2/roles/{role_id}/users | Add a user to a role
[**CreateRole**](RolesApi.md#CreateRole) | **Post** /api/v2/roles | Create role
[**DeleteRole**](RolesApi.md#DeleteRole) | **Delete** /api/v2/roles/{role_id} | Delete role
[**GetRole**](RolesApi.md#GetRole) | **Get** /api/v2/roles/{role_id} | Get a role
[**ListPermissions**](RolesApi.md#ListPermissions) | **Get** /api/v2/permissions | List permissions
[**ListRolePermissions**](RolesApi.md#ListRolePermissions) | **Get** /api/v2/roles/{role_id}/permissions | List permissions for a role
[**ListRoleUsers**](RolesApi.md#ListRoleUsers) | **Get** /api/v2/roles/{role_id}/users | Get all users of a role
[**ListRoles**](RolesApi.md#ListRoles) | **Get** /api/v2/roles | List roles
[**RemovePermissionFromRole**](RolesApi.md#RemovePermissionFromRole) | **Delete** /api/v2/roles/{role_id}/permissions | Revoke permission
[**RemoveUserFromRole**](RolesApi.md#RemoveUserFromRole) | **Delete** /api/v2/roles/{role_id}/users | Remove a user from a role
[**UpdateRole**](RolesApi.md#UpdateRole) | **Patch** /api/v2/roles/{role_id} | Update a role



## AddPermissionToRole

> PermissionsResponse AddPermissionToRole(ctx, roleId).Body(body).Execute()

Grant permission to a role



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **string** | The ID of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddPermissionToRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**RelationshipToPermission**](RelationshipToPermission.md) |  | 

### Return type

[**PermissionsResponse**](PermissionsResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddUserToRole

> UsersResponse AddUserToRole(ctx, roleId).Body(body).Execute()

Add a user to a role



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **string** | The ID of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddUserToRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**RelationshipToUser**](RelationshipToUser.md) |  | 

### Return type

[**UsersResponse**](UsersResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRole

> RoleResponse CreateRole(ctx).Body(body).Execute()

Create role



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RoleCreatePayload**](RoleCreatePayload.md) |  | 

### Return type

[**RoleResponse**](RoleResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRole

> DeleteRole(ctx, roleId).Execute()

Delete role



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **string** | The ID of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetRole

> RoleResponse GetRole(ctx, roleId).Execute()

Get a role



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **string** | The ID of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RoleResponse**](RoleResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPermissions

> PermissionsResponse ListPermissions(ctx).Execute()

List permissions



### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListPermissionsRequest struct via the builder pattern


### Return type

[**PermissionsResponse**](PermissionsResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, applcation/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRolePermissions

> PermissionsResponse ListRolePermissions(ctx, roleId).Execute()

List permissions for a role



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **string** | The ID of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRolePermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PermissionsResponse**](PermissionsResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRoleUsers

> UsersResponse ListRoleUsers(ctx, roleId).PageSize(pageSize).PageNumber(pageNumber).Sort(sort).Filter(filter).Execute()

Get all users of a role



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **string** | The ID of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRoleUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int64** | Number of users to return for a given page. | [default to 10]
 **pageNumber** | **int64** | Specific page number to return. | [default to 0]
 **sort** | **string** | User attribute to order results by. Sort order is **ascending** by default. Sort order is **descending** if the field is prefixed by a negative sign, for example &#x60;sort&#x3D;-name&#x60;. Options: &#x60;name&#x60;, &#x60;email&#x60;, &#x60;status&#x60;. | [default to &quot;name&quot;]
 **filter** | **string** | Filter all users by the given string. Defaults to no filtering. | 

### Return type

[**UsersResponse**](UsersResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRoles

> RolesResponse ListRoles(ctx).PageSize(pageSize).PageNumber(pageNumber).Sort(sort).Filter(filter).Execute()

List roles



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int64** | Number of users to return for a given page. | [default to 10]
 **pageNumber** | **int64** | Specific page number to return. | [default to 0]
 **sort** | [**RolesSort**](.md) | Sort roles depending on the given field. Sort order is **ascending** by default. Sort order is **descending** if the field is prefixed by a negative sign, for example: &#x60;sort&#x3D;-name&#x60;. | [default to &quot;name&quot;]
 **filter** | **string** | Filter all roles by the given string. | 

### Return type

[**RolesResponse**](RolesResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, applcation/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemovePermissionFromRole

> PermissionsResponse RemovePermissionFromRole(ctx, roleId).Body(body).Execute()

Revoke permission



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **string** | The ID of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemovePermissionFromRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**RelationshipToPermission**](RelationshipToPermission.md) |  | 

### Return type

[**PermissionsResponse**](PermissionsResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveUserFromRole

> UsersResponse RemoveUserFromRole(ctx, roleId).Body(body).Execute()

Remove a user from a role



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **string** | The ID of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveUserFromRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**RelationshipToUser**](RelationshipToUser.md) |  | 

### Return type

[**UsersResponse**](UsersResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRole

> RoleResponse UpdateRole(ctx, roleId).Body(body).Execute()

Update a role



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **string** | The ID of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**RoleUpdatePayload**](RoleUpdatePayload.md) |  | 

### Return type

[**RoleResponse**](RoleResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

