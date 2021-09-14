# UsersApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------ | ------------ | ------------
[**CreateUser**](UsersApi.md#CreateUser) | **Post** /api/v1/user | Create a user
[**DisableUser**](UsersApi.md#DisableUser) | **Delete** /api/v1/user/{user_handle} | Disable a user
[**GetUser**](UsersApi.md#GetUser) | **Get** /api/v1/user/{user_handle} | Get user details
[**ListUsers**](UsersApi.md#ListUsers) | **Get** /api/v1/user | List all users
[**UpdateUser**](UsersApi.md#UpdateUser) | **Put** /api/v1/user/{user_handle} | Update a user



## CreateUser

> UserResponse CreateUser(ctx, body)

Create a user for your organization.

**Note**: Users can only be created with the admin access role
if application keys belong to administrators.

### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
    ctx := datadog.NewDefaultContext(context.Background())

    body := *datadog.NewUser() // User | User object that needs to be created.

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
**body** | [**User**](User.md) | User object that needs to be created. | 


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

> UserDisableResponse DisableUser(ctx, userHandle)

Delete a user from an organization.

**Note**: This endpoint can only be used with application keys belonging to
administrators.

### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
    ctx := datadog.NewDefaultContext(context.Background())

    userHandle := "test@datadoghq.com" // string | The handle of the user.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.DisableUser(ctx, userHandle)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.DisableUser`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DisableUser`: UserDisableResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from UsersApi.DisableUser:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**userHandle** | **string** | The handle of the user. | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**UserDisableResponse**](UserDisableResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUser

> UserResponse GetUser(ctx, userHandle)

Get a user's details.

### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
    ctx := datadog.NewDefaultContext(context.Background())

    userHandle := "test@datadoghq.com" // string | The ID of the user.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.GetUser(ctx, userHandle)
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
**userHandle** | **string** | The ID of the user. | 


### Optional Parameters

This endpoint does not have optional parameters.


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


## ListUsers

> UserListResponse ListUsers(ctx)

List all users for your organization.

### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
    ctx := datadog.NewDefaultContext(context.Background())


    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.ListUsers(ctx)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.ListUsers`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUsers`: UserListResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from UsersApi.ListUsers:\n%s\n", responseContent)
}
```

### Required Parameters

This endpoint does not need any parameter.


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**UserListResponse**](UserListResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUser

> UserResponse UpdateUser(ctx, userHandle, body)

Update a user information.

**Note**: It can only be used with application keys belonging to administrators.

### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
    ctx := datadog.NewDefaultContext(context.Background())

    userHandle := "test@datadoghq.com" // string | The ID of the user.
    body := *datadog.NewUser() // User | Description of the update.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UpdateUser(ctx, userHandle, body)
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
**userHandle** | **string** | The ID of the user. |  |
**body** | [**User**](User.md) | Description of the update. | 


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

