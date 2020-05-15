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



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
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

    body := datadog.OrganizationCreateBody{Billing: datadog.OrganizationBilling{Type: "Type_example"}, Name: "Name_example", Subscription: datadog.OrganizationSubscription{Type: "Type_example"}} // OrganizationCreateBody | Organization object that needs to be created

    configuration := datadog.NewConfiguration()
    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.OrganizationsApi.CreateChildOrg(ctx, body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.CreateChildOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateChildOrg`: OrganizationCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.CreateChildOrg`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateChildOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OrganizationCreateBody**](OrganizationCreateBody.md) | Organization object that needs to be created | 

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



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
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

    publicId := "publicId_example" // string | The `public_id` of the organization you are operating within.

    configuration := datadog.NewConfiguration()
    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.OrganizationsApi.GetOrg(ctx, publicId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrg`: OrganizationResponse
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicId** | **string** | The &#x60;public_id&#x60; of the organization you are operating within. | 

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



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
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


    configuration := datadog.NewConfiguration()
    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.OrganizationsApi.ListOrgs(ctx).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.ListOrgs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOrgs`: OrganizationListResponse
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.ListOrgs`: %v\n", resp)
}
```

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



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
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

    publicId := "publicId_example" // string | The `public_id` of the organization you are operating within.
    body := datadog.Organization{Billing: datadog.OrganizationBilling{Type: "Type_example"}, Created: "Created_example", Description: "Description_example", Name: "Name_example", PublicId: "PublicId_example", Settings: datadog.Organization_settings{PrivateWidgetShare: false, Saml: datadog.Organization_settings_saml{Enabled: false}, SamlAutocreateAccessRole: datadog.AccessRole{}, SamlAutocreateUsersDomains: datadog.Organization_settings_saml_autocreate_users_domains{Domains: []string{"Domains_example"), Enabled: false}, SamlCanBeEnabled: false, SamlIdpEndpoint: "SamlIdpEndpoint_example", SamlIdpInitiatedLogin: datadog.Organization_settings_saml_idp_initiated_login{Enabled: false}, SamlIdpMetadataUploaded: false, SamlLoginUrl: "SamlLoginUrl_example", SamlStrictMode: datadog.Organization_settings_saml_strict_mode{Enabled: false}}, Subscription: datadog.OrganizationSubscription{Type: "Type_example"}} // Organization | 

    configuration := datadog.NewConfiguration()
    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.OrganizationsApi.UpdateOrg(ctx, publicId, body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.UpdateOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrg`: OrganizationResponse
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.UpdateOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicId** | **string** | The &#x60;public_id&#x60; of the organization you are operating within. | 

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



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
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

    publicId := "publicId_example" // string | The `public_id` of the organization you are operating with
    idpFile := 987 // *os.File | The path to the XML metadata file you wish to upload.

    configuration := datadog.NewConfiguration()
    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.OrganizationsApi.UploadIdPForOrg(ctx, publicId, idpFile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.UploadIdPForOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadIdPForOrg`: IdpResponse
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.UploadIdPForOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicId** | **string** | The &#x60;public_id&#x60; of the organization you are operating with | 

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

