# UsersApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------ | ------------ | ------------
[**CreateServiceAccount**](UsersApi.md#CreateServiceAccount) | **Post** /api/v2/service_accounts | Create a service account
[**CreateUser**](UsersApi.md#CreateUser) | **Post** /api/v2/users | Create a user
[**DisableUser**](UsersApi.md#DisableUser) | **Delete** /api/v2/users/{user_id} | Disable a user
[**GetInvitation**](UsersApi.md#GetInvitation) | **Get** /api/v2/user_invitations/{user_invitation_uuid} | Get a user invitation
[**GetUser**](UsersApi.md#GetUser) | **Get** /api/v2/users/{user_id} | Get user details
[**ListUserOrganizations**](UsersApi.md#ListUserOrganizations) | **Get** /api/v2/users/{user_id}/orgs | Get a user organization
[**ListUserPermissions**](UsersApi.md#ListUserPermissions) | **Get** /api/v2/users/{user_id}/permissions | Get a user permissions
[**ListUsers**](UsersApi.md#ListUsers) | **Get** /api/v2/users | List all users
[**SendInvitations**](UsersApi.md#SendInvitations) | **Post** /api/v2/user_invitations | Send invitation emails
[**UpdateUser**](UsersApi.md#UpdateUser) | **Patch** /api/v2/users/{user_id} | Update a user



## CreateServiceAccount

> UserResponse CreateServiceAccount(ctx, body)

Create a service account for your organization.

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

    body := *datadog.NewServiceAccountCreateRequest(*datadog.NewServiceAccountCreateData(*datadog.NewServiceAccountCreateAttributes("jane.doe@example.com", true), datadog.UsersType("users"))) // ServiceAccountCreateRequest | 

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.CreateServiceAccount(ctx, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.CreateServiceAccount`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateServiceAccount`: UserResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from UsersApi.CreateServiceAccount:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**body** | [**ServiceAccountCreateRequest**](ServiceAccountCreateRequest.md) |  | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**UserResponse**](UserResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUser

> UserResponse CreateUser(ctx, body)

Create a user for your organization.

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

    body := *datadog.NewUserCreateRequest(*datadog.NewUserCreateData(*datadog.NewUserCreateAttributes("jane.doe@example.com"), datadog.UsersType("users"))) // UserCreateRequest | 

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.CreateUser(ctx, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.CreateUser`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUser`: UserResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from UsersApi.CreateUser:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**body** | [**UserCreateRequest**](UserCreateRequest.md) |  | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**UserResponse**](UserResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisableUser

> DisableUser(ctx, userId)

Disable a user. Can only be used with an application key belonging
to an administrator user.

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

    userId := "00000000-0000-0000-0000-000000000000" // string | The ID of the user.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    r, err := apiClient.UsersApi.DisableUser(ctx, userId)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.DisableUser`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**userId** | **string** | The ID of the user. | 


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


## GetInvitation

> UserInvitationResponse GetInvitation(ctx, userInvitationUuid)

Returns a single user invitation by its UUID.

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

    userInvitationUuid := "00000000-0000-0000-0000-000000000000" // string | The UUID of the user invitation.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.GetInvitation(ctx, userInvitationUuid)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.GetInvitation`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInvitation`: UserInvitationResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from UsersApi.GetInvitation:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**userInvitationUuid** | **string** | The UUID of the user invitation. | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**UserInvitationResponse**](UserInvitationResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUser

> UserResponse GetUser(ctx, userId)

Get a user in the organization specified by the user’s `user_id`.

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

    userId := "00000000-0000-0000-0000-000000000000" // string | The ID of the user.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.GetUser(ctx, userId)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.GetUser`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUser`: UserResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from UsersApi.GetUser:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**userId** | **string** | The ID of the user. | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**UserResponse**](UserResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserOrganizations

> UserResponse ListUserOrganizations(ctx, userId)

Get a user organization. Returns the user information and all organizations
joined by this user.

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

    userId := "00000000-0000-0000-0000-000000000000" // string | The ID of the user.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.ListUserOrganizations(ctx, userId)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.ListUserOrganizations`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUserOrganizations`: UserResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from UsersApi.ListUserOrganizations:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**userId** | **string** | The ID of the user. | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**UserResponse**](UserResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserPermissions

> PermissionsResponse ListUserPermissions(ctx, userId)

Get a user permission set. Returns a list of the user’s permissions
granted by the associated user's roles.

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

    userId := "00000000-0000-0000-0000-000000000000" // string | The ID of the user.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.ListUserPermissions(ctx, userId)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.ListUserPermissions`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUserPermissions`: PermissionsResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from UsersApi.ListUserPermissions:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**userId** | **string** | The ID of the user. | 


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


## ListUsers

> UsersResponse ListUsers(ctx, datadog.ListUsersOptionalParameters{})

Get the list of all users in the organization. This list includes
all users even if they are deactivated or unverified.

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
    sort := "name" // string | User attribute to order results by. Sort order is ascending by default. Sort order is descending if the field is prefixed by a negative sign, for example `sort=-name`. Options: `name`, `modified_at`, `user_count`. (optional) (default to "name")
    sortDir := datadog.QuerySortOrder("asc") // QuerySortOrder | Direction of sort. Options: `asc`, `desc`. (optional) (default to "desc")
    filter := "filter_example" // string | Filter all users by the given string. Defaults to no filtering. (optional)
    filterStatus := "Active" // string | Filter on status attribute. Comma separated list, with possible values `Active`, `Pending`, and `Disabled`. Defaults to no filtering. (optional)
    optionalParams := datadog.ListUsersOptionalParameters{
        PageSize: &pageSize,
        PageNumber: &pageNumber,
        Sort: &sort,
        SortDir: &sortDir,
        Filter: &filter,
        FilterStatus: &filterStatus,
    }

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.ListUsers(ctx, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.ListUsers`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUsers`: UsersResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from UsersApi.ListUsers:\n%s\n", responseContent)
}
```

### Required Parameters




### Optional Parameters


Other parameters are passed through a pointer to a ListUsersOptionalParameters struct.


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**pageSize** | **int64** | Size for a given page. | [default to 10]
**pageNumber** | **int64** | Specific page number to return. | [default to 0]
**sort** | **string** | User attribute to order results by. Sort order is ascending by default. Sort order is descending if the field is prefixed by a negative sign, for example &#x60;sort&#x3D;-name&#x60;. Options: &#x60;name&#x60;, &#x60;modified_at&#x60;, &#x60;user_count&#x60;. | [default to &quot;name&quot;]
**sortDir** | [**QuerySortOrder**](QuerySortOrder.md) | Direction of sort. Options: &#x60;asc&#x60;, &#x60;desc&#x60;. | [default to &quot;desc&quot;]
**filter** | **string** | Filter all users by the given string. Defaults to no filtering. | 
**filterStatus** | **string** | Filter on status attribute. Comma separated list, with possible values &#x60;Active&#x60;, &#x60;Pending&#x60;, and &#x60;Disabled&#x60;. Defaults to no filtering. | 

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


## SendInvitations

> UserInvitationsResponse SendInvitations(ctx, body)

Sends emails to one or more users inviting them to join the organization.

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

    body := *datadog.NewUserInvitationsRequest([]datadog.UserInvitationData{*datadog.NewUserInvitationData(*datadog.NewUserInvitationRelationships(*datadog.NewRelationshipToUser(*datadog.NewRelationshipToUserData("00000000-0000-0000-0000-000000000000", datadog.UsersType("users")))), datadog.UserInvitationsType("user_invitations"))}) // UserInvitationsRequest | 

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.SendInvitations(ctx, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.SendInvitations`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendInvitations`: UserInvitationsResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from UsersApi.SendInvitations:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**body** | [**UserInvitationsRequest**](UserInvitationsRequest.md) |  | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**UserInvitationsResponse**](UserInvitationsResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUser

> UserResponse UpdateUser(ctx, userId, body)

Edit a user. Can only be used with an application key belonging
to an administrator user.

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

    userId := "00000000-0000-0000-0000-000000000000" // string | The ID of the user.
    body := *datadog.NewUserUpdateRequest(*datadog.NewUserUpdateData(*datadog.NewUserUpdateAttributes(), "00000000-0000-0000-0000-000000000000", datadog.UsersType("users"))) // UserUpdateRequest | 

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UpdateUser(ctx, userId, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UpdateUser`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUser`: UserResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from UsersApi.UpdateUser:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**userId** | **string** | The ID of the user. |  |
**body** | [**UserUpdateRequest**](UserUpdateRequest.md) |  | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**UserResponse**](UserResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

