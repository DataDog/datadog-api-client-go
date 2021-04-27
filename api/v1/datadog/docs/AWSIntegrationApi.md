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

> AWSAccountCreateResponse CreateAWSAccount(ctx, body)

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
    ctx := datadog.NewDefaultContext(context.Background())

    body := *datadog.NewAWSAccount() // AWSAccount | AWS Request Object

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.AWSIntegrationApi.CreateAWSAccount(ctx, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AWSIntegrationApi.CreateAWSAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAWSAccount`: AWSAccountCreateResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from AWSIntegrationApi.CreateAWSAccount:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**AWSAccount**](AWSAccount.md) | AWS Request Object | 

### Optional Parameters

This endpoint does not have optional parameters.


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

> interface{} CreateAWSTagFilter(ctx, body)

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
    ctx := datadog.NewDefaultContext(context.Background())

    body := *datadog.NewAWSTagFilterCreateRequest() // AWSTagFilterCreateRequest | Set an AWS tag filter using an `aws_account_identifier`, `namespace`, and filtering string. Namespace options are `application_elb`, `elb`, `lambda`, `network_elb`, `rds`, `sqs`, and `custom`.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.AWSIntegrationApi.CreateAWSTagFilter(ctx, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AWSIntegrationApi.CreateAWSTagFilter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAWSTagFilter`: interface{}
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from AWSIntegrationApi.CreateAWSTagFilter:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**AWSTagFilterCreateRequest**](AWSTagFilterCreateRequest.md) | Set an AWS tag filter using an &#x60;aws_account_identifier&#x60;, &#x60;namespace&#x60;, and filtering string. Namespace options are &#x60;application_elb&#x60;, &#x60;elb&#x60;, &#x60;lambda&#x60;, &#x60;network_elb&#x60;, &#x60;rds&#x60;, &#x60;sqs&#x60;, and &#x60;custom&#x60;. | 

### Optional Parameters

This endpoint does not have optional parameters.


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

> AWSAccountCreateResponse CreateNewAWSExternalID(ctx, body)

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
    ctx := datadog.NewDefaultContext(context.Background())

    body := *datadog.NewAWSAccount() // AWSAccount | Your Datadog role delegation name. For more information about your AWS account Role name, see the [Datadog AWS integration configuration info](https://github.com/DataDog/documentation/blob/master/integrations/amazon_web_services/#installation).

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.AWSIntegrationApi.CreateNewAWSExternalID(ctx, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AWSIntegrationApi.CreateNewAWSExternalID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNewAWSExternalID`: AWSAccountCreateResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from AWSIntegrationApi.CreateNewAWSExternalID:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**AWSAccount**](AWSAccount.md) | Your Datadog role delegation name. For more information about your AWS account Role name, see the [Datadog AWS integration configuration info](https://github.com/DataDog/documentation/blob/master/integrations/amazon_web_services/#installation). | 

### Optional Parameters

This endpoint does not have optional parameters.


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

> interface{} DeleteAWSAccount(ctx, body)

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
    ctx := datadog.NewDefaultContext(context.Background())

    body := *datadog.NewAWSAccount() // AWSAccount | AWS request object

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.AWSIntegrationApi.DeleteAWSAccount(ctx, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AWSIntegrationApi.DeleteAWSAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAWSAccount`: interface{}
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from AWSIntegrationApi.DeleteAWSAccount:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**AWSAccount**](AWSAccount.md) | AWS request object | 

### Optional Parameters

This endpoint does not have optional parameters.


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

> interface{} DeleteAWSTagFilter(ctx, body)

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
    ctx := datadog.NewDefaultContext(context.Background())

    body := *datadog.NewAWSTagFilterDeleteRequest() // AWSTagFilterDeleteRequest | Delete a tag filtering entry for a given AWS account and `dd-aws` namespace.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.AWSIntegrationApi.DeleteAWSTagFilter(ctx, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AWSIntegrationApi.DeleteAWSTagFilter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAWSTagFilter`: interface{}
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from AWSIntegrationApi.DeleteAWSTagFilter:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**AWSTagFilterDeleteRequest**](AWSTagFilterDeleteRequest.md) | Delete a tag filtering entry for a given AWS account and &#x60;dd-aws&#x60; namespace. | 

### Optional Parameters

This endpoint does not have optional parameters.


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

> AWSAccountListResponse ListAWSAccounts(ctx, datadog.ListAWSAccountsOptionalParameters{})

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
    ctx := datadog.NewDefaultContext(context.Background())

    accountId := "accountId_example" // string | Only return AWS accounts that matches this `account_id`. (optional)
    roleName := "roleName_example" // string | Only return AWS accounts that matches this role_name. (optional)
    accessKeyId := "accessKeyId_example" // string | Only return AWS accounts that matches this `access_key_id`. (optional)
    optionalParams := datadog.ListAWSAccountsOptionalParameters{
        AccountId: &accountId,
        RoleName: &roleName,
        AccessKeyId: &accessKeyId,
    }

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.AWSIntegrationApi.ListAWSAccounts(ctx, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AWSIntegrationApi.ListAWSAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAWSAccounts`: AWSAccountListResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from AWSIntegrationApi.ListAWSAccounts:\n%s\n", responseContent)
}
```

### Required Parameters



### Optional Parameters


Other parameters are passed through a pointer to a ListAWSAccountsOptionalParameters struct


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

> AWSTagFilterListResponse ListAWSTagFilters(ctx, accountId)

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
    ctx := datadog.NewDefaultContext(context.Background())

    accountId := "accountId_example" // string | Only return AWS filters that matches this `account_id`.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.AWSIntegrationApi.ListAWSTagFilters(ctx, accountId)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AWSIntegrationApi.ListAWSTagFilters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAWSTagFilters`: AWSTagFilterListResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from AWSIntegrationApi.ListAWSTagFilters:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Only return AWS filters that matches this &#x60;account_id&#x60;. | 

### Optional Parameters

This endpoint does not have optional parameters.


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

> []string ListAvailableAWSNamespaces(ctx)

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
    ctx := datadog.NewDefaultContext(context.Background())


    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.AWSIntegrationApi.ListAvailableAWSNamespaces(ctx)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AWSIntegrationApi.ListAvailableAWSNamespaces``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAvailableAWSNamespaces`: []string
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from AWSIntegrationApi.ListAvailableAWSNamespaces:\n%s\n", responseContent)
}
```

### Required Parameters

This endpoint does not need any parameter.

### Optional Parameters

This endpoint does not have optional parameters.


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

> interface{} UpdateAWSAccount(ctx, body, datadog.UpdateAWSAccountOptionalParameters{})

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
    ctx := datadog.NewDefaultContext(context.Background())

    body := *datadog.NewAWSAccount() // AWSAccount | AWS request object
    accountId := "accountId_example" // string | Only return AWS accounts that matches this `account_id`. (optional)
    roleName := "roleName_example" // string | Only return AWS accounts that match this `role_name`. Required if `account_id` is specified. (optional)
    accessKeyId := "accessKeyId_example" // string | Only return AWS accounts that matches this `access_key_id`. Required if none of the other two options are specified. (optional)
    optionalParams := datadog.UpdateAWSAccountOptionalParameters{
        AccountId: &accountId,
        RoleName: &roleName,
        AccessKeyId: &accessKeyId,
    }

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.AWSIntegrationApi.UpdateAWSAccount(ctx, body, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AWSIntegrationApi.UpdateAWSAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAWSAccount`: interface{}
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from AWSIntegrationApi.UpdateAWSAccount:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**AWSAccount**](AWSAccount.md) | AWS request object | 

### Optional Parameters


Other parameters are passed through a pointer to a UpdateAWSAccountOptionalParameters struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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

