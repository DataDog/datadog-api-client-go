# RolesApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------ | ------------ | ------------
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

> PermissionsResponse AddPermissionToRole(ctx, roleId, body)

Adds a permission to a role.

### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
    ctx := datadog.NewDefaultContext(context.Background())

    roleId := "roleId_example" // string | The ID of the role.
    body := *datadog.NewRelationshipToPermission() // RelationshipToPermission | 

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.RolesApi.AddPermissionToRole(ctx, roleId, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.AddPermissionToRole`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddPermissionToRole`: PermissionsResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from RolesApi.AddPermissionToRole:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**roleId** | **string** | The ID of the role. |  |
**body** | [**RelationshipToPermission**](RelationshipToPermission.md) |  | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**PermissionsResponse**](PermissionsResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddUserToRole

> UsersResponse AddUserToRole(ctx, roleId, body)

Adds a user to a role.

### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
    ctx := datadog.NewDefaultContext(context.Background())

    roleId := "roleId_example" // string | The ID of the role.
    body := *datadog.NewRelationshipToUser(*datadog.NewRelationshipToUserData("00000000-0000-0000-0000-000000000000", datadog.UsersType("users"))) // RelationshipToUser | 

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.RolesApi.AddUserToRole(ctx, roleId, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.AddUserToRole`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddUserToRole`: UsersResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from RolesApi.AddUserToRole:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**roleId** | **string** | The ID of the role. |  |
**body** | [**RelationshipToUser**](RelationshipToUser.md) |  | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**UsersResponse**](UsersResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRole

> RoleCreateResponse CreateRole(ctx, body)

Create a new role for your organization.

### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
    ctx := datadog.NewDefaultContext(context.Background())

    body := *datadog.NewRoleCreateRequest(*datadog.NewRoleCreateData(*datadog.NewRoleCreateAttributes("developers"))) // RoleCreateRequest | 

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.RolesApi.CreateRole(ctx, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.CreateRole`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRole`: RoleCreateResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from RolesApi.CreateRole:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**body** | [**RoleCreateRequest**](RoleCreateRequest.md) |  | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**RoleCreateResponse**](RoleCreateResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRole

> DeleteRole(ctx, roleId)

Disables a role.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
    ctx := datadog.NewDefaultContext(context.Background())

    roleId := "roleId_example" // string | The ID of the role.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    r, err := apiClient.RolesApi.DeleteRole(ctx, roleId)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.DeleteRole`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**roleId** | **string** | The ID of the role. | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

 (empty response body)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRole

> RoleResponse GetRole(ctx, roleId)

Get a role in the organization specified by the role’s `role_id`.

### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
    ctx := datadog.NewDefaultContext(context.Background())

    roleId := "roleId_example" // string | The ID of the role.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.RolesApi.GetRole(ctx, roleId)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.GetRole`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRole`: RoleResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from RolesApi.GetRole:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**roleId** | **string** | The ID of the role. | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**RoleResponse**](RoleResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPermissions

> PermissionsResponse ListPermissions(ctx)

Returns a list of all permissions, including name, description, and ID.

### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
    ctx := datadog.NewDefaultContext(context.Background())


    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.RolesApi.ListPermissions(ctx)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.ListPermissions`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPermissions`: PermissionsResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from RolesApi.ListPermissions:\n%s\n", responseContent)
}
```

### Required Parameters

This endpoint does not need any parameter.


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**PermissionsResponse**](PermissionsResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRolePermissions

> PermissionsResponse ListRolePermissions(ctx, roleId)

Returns a list of all permissions for a single role.

### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
    ctx := datadog.NewDefaultContext(context.Background())

    roleId := "roleId_example" // string | The ID of the role.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.RolesApi.ListRolePermissions(ctx, roleId)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.ListRolePermissions`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRolePermissions`: PermissionsResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from RolesApi.ListRolePermissions:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**roleId** | **string** | The ID of the role. | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**PermissionsResponse**](PermissionsResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRoleUsers

> UsersResponse ListRoleUsers(ctx, roleId, datadog.ListRoleUsersOptionalParameters{})

Gets all users of a role.

### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
    ctx := datadog.NewDefaultContext(context.Background())

    roleId := "roleId_example" // string | The ID of the role.
    pageSize := int64(10) // int64 | Size for a given page. (optional) (default to 10)
    pageNumber := int64(0) // int64 | Specific page number to return. (optional) (default to 0)
    sort := "sort_example" // string | User attribute to order results by. Sort order is **ascending** by default. Sort order is **descending** if the field is prefixed by a negative sign, for example `sort=-name`. Options: `name`, `email`, `status`. (optional) (default to "name")
    filter := "filter_example" // string | Filter all users by the given string. Defaults to no filtering. (optional)
    optionalParams := datadog.ListRoleUsersOptionalParameters{
        PageSize: &pageSize,
        PageNumber: &pageNumber,
        Sort: &sort,
        Filter: &filter,
    }

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.RolesApi.ListRoleUsers(ctx, roleId, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.ListRoleUsers`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRoleUsers`: UsersResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from RolesApi.ListRoleUsers:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**roleId** | **string** | The ID of the role. | 


### Optional Parameters


Other parameters are passed through a pointer to a ListRoleUsersOptionalParameters struct.


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**pageSize** | **int64** | Size for a given page. | [default to 10]
**pageNumber** | **int64** | Specific page number to return. | [default to 0]
**sort** | **string** | User attribute to order results by. Sort order is **ascending** by default. Sort order is **descending** if the field is prefixed by a negative sign, for example &#x60;sort&#x3D;-name&#x60;. Options: &#x60;name&#x60;, &#x60;email&#x60;, &#x60;status&#x60;. | [default to &quot;name&quot;]
**filter** | **string** | Filter all users by the given string. Defaults to no filtering. | 

### Return type

[**UsersResponse**](UsersResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRoles

> RolesResponse ListRoles(ctx, datadog.ListRolesOptionalParameters{})

Returns all roles, including their names and IDs.

### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
    ctx := datadog.NewDefaultContext(context.Background())

    pageSize := int64(10) // int64 | Size for a given page. (optional) (default to 10)
    pageNumber := int64(0) // int64 | Specific page number to return. (optional) (default to 0)
    sort := datadog.RolesSort("name") // RolesSort | Sort roles depending on the given field. Sort order is **ascending** by default. Sort order is **descending** if the field is prefixed by a negative sign, for example: `sort=-name`. (optional) (default to "name")
    filter := "filter_example" // string | Filter all roles by the given string. (optional)
    optionalParams := datadog.ListRolesOptionalParameters{
        PageSize: &pageSize,
        PageNumber: &pageNumber,
        Sort: &sort,
        Filter: &filter,
    }

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.RolesApi.ListRoles(ctx, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.ListRoles`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRoles`: RolesResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from RolesApi.ListRoles:\n%s\n", responseContent)
}
```

### Required Parameters




### Optional Parameters


Other parameters are passed through a pointer to a ListRolesOptionalParameters struct.


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**pageSize** | **int64** | Size for a given page. | [default to 10]
**pageNumber** | **int64** | Specific page number to return. | [default to 0]
**sort** | [**RolesSort**](RolesSort.md) | Sort roles depending on the given field. Sort order is **ascending** by default. Sort order is **descending** if the field is prefixed by a negative sign, for example: &#x60;sort&#x3D;-name&#x60;. | [default to &quot;name&quot;]
**filter** | **string** | Filter all roles by the given string. | 

### Return type

[**RolesResponse**](RolesResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemovePermissionFromRole

> PermissionsResponse RemovePermissionFromRole(ctx, roleId, body)

Removes a permission from a role.

### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
    ctx := datadog.NewDefaultContext(context.Background())

    roleId := "roleId_example" // string | The ID of the role.
    body := *datadog.NewRelationshipToPermission() // RelationshipToPermission | 

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.RolesApi.RemovePermissionFromRole(ctx, roleId, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.RemovePermissionFromRole`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemovePermissionFromRole`: PermissionsResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from RolesApi.RemovePermissionFromRole:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**roleId** | **string** | The ID of the role. |  |
**body** | [**RelationshipToPermission**](RelationshipToPermission.md) |  | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**PermissionsResponse**](PermissionsResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveUserFromRole

> UsersResponse RemoveUserFromRole(ctx, roleId, body)

Removes a user from a role.

### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
    ctx := datadog.NewDefaultContext(context.Background())

    roleId := "roleId_example" // string | The ID of the role.
    body := *datadog.NewRelationshipToUser(*datadog.NewRelationshipToUserData("00000000-0000-0000-0000-000000000000", datadog.UsersType("users"))) // RelationshipToUser | 

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.RolesApi.RemoveUserFromRole(ctx, roleId, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.RemoveUserFromRole`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveUserFromRole`: UsersResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from RolesApi.RemoveUserFromRole:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**roleId** | **string** | The ID of the role. |  |
**body** | [**RelationshipToUser**](RelationshipToUser.md) |  | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**UsersResponse**](UsersResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRole

> RoleUpdateResponse UpdateRole(ctx, roleId, body)

Edit a role. Can only be used with application keys belonging to administrators.

### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
    ctx := datadog.NewDefaultContext(context.Background())

    roleId := "roleId_example" // string | The ID of the role.
    body := *datadog.NewRoleUpdateRequest(*datadog.NewRoleUpdateData(*datadog.NewRoleUpdateAttributes(), "00000000-0000-0000-0000-000000000000", datadog.RolesType("roles"))) // RoleUpdateRequest | 

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.RolesApi.UpdateRole(ctx, roleId, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.UpdateRole`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRole`: RoleUpdateResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from RolesApi.UpdateRole:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**roleId** | **string** | The ID of the role. |  |
**body** | [**RoleUpdateRequest**](RoleUpdateRequest.md) |  | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**RoleUpdateResponse**](RoleUpdateResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

