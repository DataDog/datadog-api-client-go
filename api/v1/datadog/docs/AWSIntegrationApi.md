# \AWSIntegrationApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAWSAccount**](AWSIntegrationApi.md#CreateAWSAccount) | **Post** /api/v1/integration/aws | Create an AWS integration
[**CreateAWSTagFilter**](AWSIntegrationApi.md#CreateAWSTagFilter) | **Post** /api/v1/integration/aws/filtering | Set an AWS tag filter
[**CreateNewAWSExternalID**](AWSIntegrationApi.md#CreateNewAWSExternalID) | **Put** /api/v1/integration/aws/generate_new_external_id | Generate a new external ID
[**DeleteAWSAccount**](AWSIntegrationApi.md#DeleteAWSAccount) | **Delete** /api/v1/integration/aws | Delete an AWS integration
[**DeleteAWSTagFilter**](AWSIntegrationApi.md#DeleteAWSTagFilter) | **Delete** /api/v1/integration/aws/filtering | Delete a tag filtering entry
[**ListAWSAccounts**](AWSIntegrationApi.md#ListAWSAccounts) | **Get** /api/v1/integration/aws | List all AWS integrations
[**ListAWSTagFilters**](AWSIntegrationApi.md#ListAWSTagFilters) | **Get** /api/v1/integration/aws/filtering | Get all AWS tag filters
[**ListAvailableAWSNamespaces**](AWSIntegrationApi.md#ListAvailableAWSNamespaces) | **Get** /api/v1/integration/aws/available_namespace_rules | List namespace rules
[**UpdateAWSAccount**](AWSIntegrationApi.md#UpdateAWSAccount) | **Put** /api/v1/integration/aws | Update an AWS integration



## CreateAWSAccount

> AWSAccountCreateResponse CreateAWSAccount(ctx).Body(body).Execute()

Create an AWS integration



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

    if site, ok := os.LookupEnv("DD_SITE"); ok {
        ctx = context.WithValue(
            ctx,
            datadog.ContextServerVariables,
            map[string]string{"site": site},
        )
    }

    body := *datadog.NewAWSAccount() // AWSAccount | AWS Request Object

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.AWSIntegrationApi.CreateAWSAccount(ctx).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AWSIntegrationApi.CreateAWSAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAWSAccount`: AWSAccountCreateResponse
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from AWSIntegrationApi.CreateAWSAccount:\n%s\n", response_content)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAWSAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AWSAccount**](AWSAccount.md) | AWS Request Object | 

### Return type

[**AWSAccountCreateResponse**](AWSAccountCreateResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAWSTagFilter

> interface{} CreateAWSTagFilter(ctx).Body(body).Execute()

Set an AWS tag filter



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

    if site, ok := os.LookupEnv("DD_SITE"); ok {
        ctx = context.WithValue(
            ctx,
            datadog.ContextServerVariables,
            map[string]string{"site": site},
        )
    }

    body := *datadog.NewAWSTagFilterCreateRequest() // AWSTagFilterCreateRequest | Set an AWS tag filter using an `aws_account_identifier`, `namespace`, and filtering string. Namespace options are `application_elb`, `elb`, `lambda`, `network_elb`, `rds`, `sqs`, and `custom`.

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.AWSIntegrationApi.CreateAWSTagFilter(ctx).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AWSIntegrationApi.CreateAWSTagFilter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAWSTagFilter`: interface{}
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from AWSIntegrationApi.CreateAWSTagFilter:\n%s\n", response_content)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAWSTagFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AWSTagFilterCreateRequest**](AWSTagFilterCreateRequest.md) | Set an AWS tag filter using an &#x60;aws_account_identifier&#x60;, &#x60;namespace&#x60;, and filtering string. Namespace options are &#x60;application_elb&#x60;, &#x60;elb&#x60;, &#x60;lambda&#x60;, &#x60;network_elb&#x60;, &#x60;rds&#x60;, &#x60;sqs&#x60;, and &#x60;custom&#x60;. | 

### Return type

**interface{}**

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNewAWSExternalID

> AWSAccountCreateResponse CreateNewAWSExternalID(ctx).Body(body).Execute()

Generate a new external ID



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

    if site, ok := os.LookupEnv("DD_SITE"); ok {
        ctx = context.WithValue(
            ctx,
            datadog.ContextServerVariables,
            map[string]string{"site": site},
        )
    }

    body := *datadog.NewAWSAccount() // AWSAccount | Your Datadog role delegation name. For more information about your AWS account Role name, see the [Datadog AWS integration configuration info](https://github.com/DataDog/documentation/blob/master/integrations/amazon_web_services/#installation).

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.AWSIntegrationApi.CreateNewAWSExternalID(ctx).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AWSIntegrationApi.CreateNewAWSExternalID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNewAWSExternalID`: AWSAccountCreateResponse
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from AWSIntegrationApi.CreateNewAWSExternalID:\n%s\n", response_content)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNewAWSExternalIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AWSAccount**](AWSAccount.md) | Your Datadog role delegation name. For more information about your AWS account Role name, see the [Datadog AWS integration configuration info](https://github.com/DataDog/documentation/blob/master/integrations/amazon_web_services/#installation). | 

### Return type

[**AWSAccountCreateResponse**](AWSAccountCreateResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAWSAccount

> interface{} DeleteAWSAccount(ctx).Body(body).Execute()

Delete an AWS integration



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

    if site, ok := os.LookupEnv("DD_SITE"); ok {
        ctx = context.WithValue(
            ctx,
            datadog.ContextServerVariables,
            map[string]string{"site": site},
        )
    }

    body := *datadog.NewAWSAccount() // AWSAccount | AWS request object

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.AWSIntegrationApi.DeleteAWSAccount(ctx).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AWSIntegrationApi.DeleteAWSAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAWSAccount`: interface{}
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from AWSIntegrationApi.DeleteAWSAccount:\n%s\n", response_content)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAWSAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AWSAccount**](AWSAccount.md) | AWS request object | 

### Return type

**interface{}**

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAWSTagFilter

> interface{} DeleteAWSTagFilter(ctx).Body(body).Execute()

Delete a tag filtering entry



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

    if site, ok := os.LookupEnv("DD_SITE"); ok {
        ctx = context.WithValue(
            ctx,
            datadog.ContextServerVariables,
            map[string]string{"site": site},
        )
    }

    body := *datadog.NewAWSTagFilterDeleteRequest() // AWSTagFilterDeleteRequest | Delete a tag filtering entry for a given AWS account and `dd-aws` namespace.

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.AWSIntegrationApi.DeleteAWSTagFilter(ctx).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AWSIntegrationApi.DeleteAWSTagFilter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAWSTagFilter`: interface{}
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from AWSIntegrationApi.DeleteAWSTagFilter:\n%s\n", response_content)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAWSTagFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AWSTagFilterDeleteRequest**](AWSTagFilterDeleteRequest.md) | Delete a tag filtering entry for a given AWS account and &#x60;dd-aws&#x60; namespace. | 

### Return type

**interface{}**

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAWSAccounts

> AWSAccountListResponse ListAWSAccounts(ctx).AccountId(accountId).RoleName(roleName).AccessKeyId(accessKeyId).Execute()

List all AWS integrations



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

    if site, ok := os.LookupEnv("DD_SITE"); ok {
        ctx = context.WithValue(
            ctx,
            datadog.ContextServerVariables,
            map[string]string{"site": site},
        )
    }

    accountId := "accountId_example" // string | Only return AWS accounts that matches this `account_id`. (optional)
    roleName := "roleName_example" // string | Only return AWS accounts that matches this role_name. (optional)
    accessKeyId := "accessKeyId_example" // string | Only return AWS accounts that matches this `access_key_id`. (optional)

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.AWSIntegrationApi.ListAWSAccounts(ctx).AccountId(accountId).RoleName(roleName).AccessKeyId(accessKeyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AWSIntegrationApi.ListAWSAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAWSAccounts`: AWSAccountListResponse
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from AWSIntegrationApi.ListAWSAccounts:\n%s\n", response_content)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAWSAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** | Only return AWS accounts that matches this &#x60;account_id&#x60;. | 
 **roleName** | **string** | Only return AWS accounts that matches this role_name. | 
 **accessKeyId** | **string** | Only return AWS accounts that matches this &#x60;access_key_id&#x60;. | 

### Return type

[**AWSAccountListResponse**](AWSAccountListResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAWSTagFilters

> AWSTagFilterListResponse ListAWSTagFilters(ctx).AccountId(accountId).Execute()

Get all AWS tag filters



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

    if site, ok := os.LookupEnv("DD_SITE"); ok {
        ctx = context.WithValue(
            ctx,
            datadog.ContextServerVariables,
            map[string]string{"site": site},
        )
    }

    accountId := "accountId_example" // string | Only return AWS filters that matches this `account_id`.

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.AWSIntegrationApi.ListAWSTagFilters(ctx).AccountId(accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AWSIntegrationApi.ListAWSTagFilters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAWSTagFilters`: AWSTagFilterListResponse
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from AWSIntegrationApi.ListAWSTagFilters:\n%s\n", response_content)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAWSTagFiltersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** | Only return AWS filters that matches this &#x60;account_id&#x60;. | 

### Return type

[**AWSTagFilterListResponse**](AWSTagFilterListResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAvailableAWSNamespaces

> []string ListAvailableAWSNamespaces(ctx).Execute()

List namespace rules



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

    if site, ok := os.LookupEnv("DD_SITE"); ok {
        ctx = context.WithValue(
            ctx,
            datadog.ContextServerVariables,
            map[string]string{"site": site},
        )
    }


    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.AWSIntegrationApi.ListAvailableAWSNamespaces(ctx).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AWSIntegrationApi.ListAvailableAWSNamespaces``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAvailableAWSNamespaces`: []string
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from AWSIntegrationApi.ListAvailableAWSNamespaces:\n%s\n", response_content)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAvailableAWSNamespacesRequest struct via the builder pattern


### Return type

**[]string**

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAWSAccount

> interface{} UpdateAWSAccount(ctx).Body(body).AccountId(accountId).RoleName(roleName).AccessKeyId(accessKeyId).Execute()

Update an AWS integration



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

    if site, ok := os.LookupEnv("DD_SITE"); ok {
        ctx = context.WithValue(
            ctx,
            datadog.ContextServerVariables,
            map[string]string{"site": site},
        )
    }

    body := *datadog.NewAWSAccount() // AWSAccount | AWS request object
    accountId := "accountId_example" // string | Only return AWS accounts that matches this `account_id`. (optional)
    roleName := "roleName_example" // string | Only return AWS accounts that match this `role_name`. Required if `account_id` is specified. (optional)
    accessKeyId := "accessKeyId_example" // string | Only return AWS accounts that matches this `access_key_id`. Required if none of the other two options are specified. (optional)

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.AWSIntegrationApi.UpdateAWSAccount(ctx).Body(body).AccountId(accountId).RoleName(roleName).AccessKeyId(accessKeyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AWSIntegrationApi.UpdateAWSAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAWSAccount`: interface{}
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from AWSIntegrationApi.UpdateAWSAccount:\n%s\n", response_content)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAWSAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AWSAccount**](AWSAccount.md) | AWS request object | 
 **accountId** | **string** | Only return AWS accounts that matches this &#x60;account_id&#x60;. | 
 **roleName** | **string** | Only return AWS accounts that match this &#x60;role_name&#x60;. Required if &#x60;account_id&#x60; is specified. | 
 **accessKeyId** | **string** | Only return AWS accounts that matches this &#x60;access_key_id&#x60;. Required if none of the other two options are specified. | 

### Return type

**interface{}**

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

