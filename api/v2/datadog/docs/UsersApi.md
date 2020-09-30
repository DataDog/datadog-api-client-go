# \UsersApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUser**](UsersApi.md#CreateUser) | **Post** /api/v2/users | Create a user
[**DisableUser**](UsersApi.md#DisableUser) | **Delete** /api/v2/users/{user_id} | Disable a user
[**GetInvitation**](UsersApi.md#GetInvitation) | **Get** /api/v2/user_invitations/{user_invitation_uuid} | Get a user invitation
[**GetUser**](UsersApi.md#GetUser) | **Get** /api/v2/users/{user_id} | Get user details
[**ListUserOrganizations**](UsersApi.md#ListUserOrganizations) | **Get** /api/v2/users/{user_id}/orgs | Get a user organization
[**ListUserPermissions**](UsersApi.md#ListUserPermissions) | **Get** /api/v2/users/{user_id}/permissions | Get a user permissions
[**ListUsers**](UsersApi.md#ListUsers) | **Get** /api/v2/users | List all users
[**SendInvitations**](UsersApi.md#SendInvitations) | **Post** /api/v2/user_invitations | Send invitation emails
[**UpdateUser**](UsersApi.md#UpdateUser) | **Patch** /api/v2/users/{user_id} | Update a user



## CreateUser

> UserResponse CreateUser(ctx).Body(body).Execute()

Create a user



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
    ctx := context.WithValue(
        context.Background(),
        datadog.ContextAPIKeys,
        map[string]datadog.APIKey{
            "apiKeyAuth": {
                Key: os.Getenv("DD_CLIENT_API_KEY"),
            },
            "appKeyAuth": {
                Key: os.Getenv("DD_CLIENT_APP_KEY"),
            },
        },
    )

    body := datadog.UserCreateRequest{Data: datadog.UserCreateData{Attributes: datadog.UserCreateAttributes{Email: "Email_example", Name: "Name_example", Title: "Title_example"}, Relationships: datadog.UserRelationships{Roles: datadog.RelationshipToRoles{Data: []RelationshipToRoleData{datadog.RelationshipToRoleData{Id: "Id_example", Type: datadog.RolesType{}})}}, Type: datadog.UsersType{}}} // UserCreateRequest |  (optional)

    configuration := datadog.NewConfiguration()
    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.CreateUser(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.CreateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUser`: UserResponse
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.CreateUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UserCreateRequest**](UserCreateRequest.md) |  | 

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


## DisableUser

> DisableUser(ctx, userId).Execute()

Disable a user



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
    ctx := context.WithValue(
        context.Background(),
        datadog.ContextAPIKeys,
        map[string]datadog.APIKey{
            "apiKeyAuth": {
                Key: os.Getenv("DD_CLIENT_API_KEY"),
            },
            "appKeyAuth": {
                Key: os.Getenv("DD_CLIENT_APP_KEY"),
            },
        },
    )

    userId := "userId_example" // string | The ID of the user.

    configuration := datadog.NewConfiguration()
    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.DisableUser(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.DisableUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The ID of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableUserRequest struct via the builder pattern


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


## GetInvitation

> UserInvitationResponse GetInvitation(ctx, userInvitationUuid).Execute()

Get a user invitation



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
    ctx := context.WithValue(
        context.Background(),
        datadog.ContextAPIKeys,
        map[string]datadog.APIKey{
            "apiKeyAuth": {
                Key: os.Getenv("DD_CLIENT_API_KEY"),
            },
            "appKeyAuth": {
                Key: os.Getenv("DD_CLIENT_APP_KEY"),
            },
        },
    )

    userInvitationUuid := "userInvitationUuid_example" // string | The UUID of the user invitation.

    configuration := datadog.NewConfiguration()
    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.GetInvitation(context.Background(), userInvitationUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.GetInvitation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInvitation`: UserInvitationResponse
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.GetInvitation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userInvitationUuid** | **string** | The UUID of the user invitation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInvitationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserInvitationResponse**](UserInvitationResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUser

> UserResponse GetUser(ctx, userId).Execute()

Get user details



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
    ctx := context.WithValue(
        context.Background(),
        datadog.ContextAPIKeys,
        map[string]datadog.APIKey{
            "apiKeyAuth": {
                Key: os.Getenv("DD_CLIENT_API_KEY"),
            },
            "appKeyAuth": {
                Key: os.Getenv("DD_CLIENT_APP_KEY"),
            },
        },
    )

    userId := "userId_example" // string | The ID of the user.

    configuration := datadog.NewConfiguration()
    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.GetUser(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.GetUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUser`: UserResponse
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.GetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The ID of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserResponse**](UserResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserOrganizations

> UserResponse ListUserOrganizations(ctx, userId).Execute()

Get a user organization



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
    ctx := context.WithValue(
        context.Background(),
        datadog.ContextAPIKeys,
        map[string]datadog.APIKey{
            "apiKeyAuth": {
                Key: os.Getenv("DD_CLIENT_API_KEY"),
            },
            "appKeyAuth": {
                Key: os.Getenv("DD_CLIENT_APP_KEY"),
            },
        },
    )

    userId := "userId_example" // string | The ID of the user.

    configuration := datadog.NewConfiguration()
    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.ListUserOrganizations(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.ListUserOrganizations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUserOrganizations`: UserResponse
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.ListUserOrganizations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The ID of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUserOrganizationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserResponse**](UserResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserPermissions

> PermissionsResponse ListUserPermissions(ctx, userId).Execute()

Get a user permissions



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
    ctx := context.WithValue(
        context.Background(),
        datadog.ContextAPIKeys,
        map[string]datadog.APIKey{
            "apiKeyAuth": {
                Key: os.Getenv("DD_CLIENT_API_KEY"),
            },
            "appKeyAuth": {
                Key: os.Getenv("DD_CLIENT_APP_KEY"),
            },
        },
    )

    userId := "userId_example" // string | The ID of the user.

    configuration := datadog.NewConfiguration()
    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.ListUserPermissions(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.ListUserPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUserPermissions`: PermissionsResponse
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.ListUserPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The ID of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUserPermissionsRequest struct via the builder pattern


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


## ListUsers

> UsersResponse ListUsers(ctx).PageSize(pageSize).PageNumber(pageNumber).Sort(sort).SortDir(sortDir).Filter(filter).FilterStatus(filterStatus).Execute()

List all users



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
    ctx := context.WithValue(
        context.Background(),
        datadog.ContextAPIKeys,
        map[string]datadog.APIKey{
            "apiKeyAuth": {
                Key: os.Getenv("DD_CLIENT_API_KEY"),
            },
            "appKeyAuth": {
                Key: os.Getenv("DD_CLIENT_APP_KEY"),
            },
        },
    )

    pageSize := 987 // int64 | Size for a given page. (optional) (default to 10)
    pageNumber := 987 // int64 | Specific page number to return. (optional) (default to 0)
    sort := "sort_example" // string | User attribute to order results by. Sort order is ascending by default. Sort order is descending if the field is prefixed by a negative sign, for example `sort=-name`. Options: `name`, `modified_at`, `user_count`. (optional) (default to "name")
    sortDir := datadog.QuerySortOrder{} // QuerySortOrder | Direction of sort. Options: `asc`, `desc`. (optional) (default to "desc")
    filter := "filter_example" // string | Filter all users by the given string. Defaults to no filtering. (optional)
    filterStatus := "filterStatus_example" // string | Filter on status attribute. Comma separated list, with possible values `Active`, `Pending`, and `Disabled`. Defaults to no filtering. (optional)

    configuration := datadog.NewConfiguration()
    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.ListUsers(context.Background()).PageSize(pageSize).PageNumber(pageNumber).Sort(sort).SortDir(sortDir).Filter(filter).FilterStatus(filterStatus).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.ListUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUsers`: UsersResponse
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.ListUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int64** | Size for a given page. | [default to 10]
 **pageNumber** | **int64** | Specific page number to return. | [default to 0]
 **sort** | **string** | User attribute to order results by. Sort order is ascending by default. Sort order is descending if the field is prefixed by a negative sign, for example &#x60;sort&#x3D;-name&#x60;. Options: &#x60;name&#x60;, &#x60;modified_at&#x60;, &#x60;user_count&#x60;. | [default to &quot;name&quot;]
 **sortDir** | [**QuerySortOrder**](.md) | Direction of sort. Options: &#x60;asc&#x60;, &#x60;desc&#x60;. | [default to &quot;desc&quot;]
 **filter** | **string** | Filter all users by the given string. Defaults to no filtering. | 
 **filterStatus** | **string** | Filter on status attribute. Comma separated list, with possible values &#x60;Active&#x60;, &#x60;Pending&#x60;, and &#x60;Disabled&#x60;. Defaults to no filtering. | 

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


## SendInvitations

> UserInvitationsResponse SendInvitations(ctx).Body(body).Execute()

Send invitation emails



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
    ctx := context.WithValue(
        context.Background(),
        datadog.ContextAPIKeys,
        map[string]datadog.APIKey{
            "apiKeyAuth": {
                Key: os.Getenv("DD_CLIENT_API_KEY"),
            },
            "appKeyAuth": {
                Key: os.Getenv("DD_CLIENT_APP_KEY"),
            },
        },
    )

    body := datadog.UserInvitationsRequest{Data: []UserInvitationData{datadog.UserInvitationData{Relationships: datadog.UserInvitationRelationships{User: datadog.RelationshipToUser{Data: datadog.RelationshipToUserData{Id: "Id_example", Type: datadog.UsersType{}}}}, Type: datadog.UserInvitationsType{}})} // UserInvitationsRequest |  (optional)

    configuration := datadog.NewConfiguration()
    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.SendInvitations(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.SendInvitations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendInvitations`: UserInvitationsResponse
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.SendInvitations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendInvitationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UserInvitationsRequest**](UserInvitationsRequest.md) |  | 

### Return type

[**UserInvitationsResponse**](UserInvitationsResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUser

> UserResponse UpdateUser(ctx, userId).Body(body).Execute()

Update a user



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
    ctx := context.WithValue(
        context.Background(),
        datadog.ContextAPIKeys,
        map[string]datadog.APIKey{
            "apiKeyAuth": {
                Key: os.Getenv("DD_CLIENT_API_KEY"),
            },
            "appKeyAuth": {
                Key: os.Getenv("DD_CLIENT_APP_KEY"),
            },
        },
    )

    userId := "userId_example" // string | The ID of the user.
    body := datadog.UserUpdateRequest{Data: datadog.UserUpdateData{Attributes: datadog.UserUpdateAttributes{Disabled: false, Email: "Email_example", Name: "Name_example"}, Id: "Id_example", Type: }} // UserUpdateRequest |  (optional)

    configuration := datadog.NewConfiguration()
    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.UpdateUser(context.Background(), userId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UpdateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUser`: UserResponse
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UpdateUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The ID of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UserUpdateRequest**](UserUpdateRequest.md) |  | 

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

