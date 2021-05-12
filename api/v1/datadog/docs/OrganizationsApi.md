# OrganizationsApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------ | ------------ | ------------
[**CreateChildOrg**](OrganizationsApi.md#CreateChildOrg) | **Post** /api/v1/org | Create a child organization
[**GetOrg**](OrganizationsApi.md#GetOrg) | **Get** /api/v1/org/{public_id} | Get organization information
[**ListOrgs**](OrganizationsApi.md#ListOrgs) | **Get** /api/v1/org | List your managed organizations
[**UpdateOrg**](OrganizationsApi.md#UpdateOrg) | **Put** /api/v1/org/{public_id} | Update your organization
[**UploadIdPForOrg**](OrganizationsApi.md#UploadIdPForOrg) | **Post** /api/v1/org/{public_id}/idp_metadata | Upload IdP metadata



## CreateChildOrg

> OrganizationCreateResponse CreateChildOrg(ctx, body)

Create a child organization.

This endpoint requires the
[multi-organization account](https://docs.datadoghq.com/account_management/multi_organization/)
feature and must be enabled by
[contacting support](https://docs.datadoghq.com/help/).

Once a new child organization is created, you can interact with it
by using the `org.public_id`, `pi_key.key`, and
`application_key.hash` provided in the response.

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

    body := *datadog.NewOrganizationCreateBody("New child org") // OrganizationCreateBody | Organization object that needs to be created

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.CreateChildOrg(ctx, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.CreateChildOrg`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateChildOrg`: OrganizationCreateResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from OrganizationsApi.CreateChildOrg:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**body** | [**OrganizationCreateBody**](OrganizationCreateBody.md) | Organization object that needs to be created | 


### Optional Parameters

This endpoint does not have optional parameters.


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

> OrganizationResponse GetOrg(ctx, publicId)

Get organization information.

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

    publicId := "abc123" // string | The `public_id` of the organization you are operating within.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetOrg(ctx, publicId)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrg`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrg`: OrganizationResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from OrganizationsApi.GetOrg:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**publicId** | **string** | The &#x60;public_id&#x60; of the organization you are operating within. | 


### Optional Parameters

This endpoint does not have optional parameters.


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

> OrganizationListResponse ListOrgs(ctx)

List your managed organizations.

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
    resp, r, err := apiClient.OrganizationsApi.ListOrgs(ctx)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.ListOrgs`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOrgs`: OrganizationListResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from OrganizationsApi.ListOrgs:\n%s\n", responseContent)
}
```

### Required Parameters

This endpoint does not need any parameter.


### Optional Parameters

This endpoint does not have optional parameters.


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

> OrganizationResponse UpdateOrg(ctx, publicId, body)

Update your organization.

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

    publicId := "abc123" // string | The `public_id` of the organization you are operating within.
    body := *datadog.NewOrganization() // Organization | 

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.UpdateOrg(ctx, publicId, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.UpdateOrg`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrg`: OrganizationResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from OrganizationsApi.UpdateOrg:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**publicId** | **string** | The &#x60;public_id&#x60; of the organization you are operating within. |  |
**body** | [**Organization**](Organization.md) |  | 


### Optional Parameters

This endpoint does not have optional parameters.


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

> IdpResponse UploadIdPForOrg(ctx, publicId, idpFile)

There are a couple of options for updating the Identity Provider (IdP)
metadata from your SAML IdP.

* **Multipart Form-Data**: Post the IdP metadata file using a form post.

* **XML Body:** Post the IdP metadata file as the body of the request.

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

    publicId := "abc123" // string | The `public_id` of the organization you are operating with
    idpFile := os.NewFile(1234, "some_file") // *os.File | The path to the XML metadata file you wish to upload.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.UploadIdPForOrg(ctx, publicId, idpFile)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.UploadIdPForOrg`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadIdPForOrg`: IdpResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from OrganizationsApi.UploadIdPForOrg:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**publicId** | **string** | The &#x60;public_id&#x60; of the organization you are operating with |  |
**idpFile** | ***os.File** | The path to the XML metadata file you wish to upload. | 


### Optional Parameters

This endpoint does not have optional parameters.


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

