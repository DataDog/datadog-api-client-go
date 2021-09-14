# AWSLogsIntegrationApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------ | ------------ | ------------
[**CheckAWSLogsLambdaAsync**](AWSLogsIntegrationApi.md#CheckAWSLogsLambdaAsync) | **Post** /api/v1/integration/aws/logs/check_async | Check that an AWS Lambda Function exists
[**CheckAWSLogsServicesAsync**](AWSLogsIntegrationApi.md#CheckAWSLogsServicesAsync) | **Post** /api/v1/integration/aws/logs/services_async | Check permissions for log services
[**CreateAWSLambdaARN**](AWSLogsIntegrationApi.md#CreateAWSLambdaARN) | **Post** /api/v1/integration/aws/logs | Add AWS Log Lambda ARN
[**DeleteAWSLambdaARN**](AWSLogsIntegrationApi.md#DeleteAWSLambdaARN) | **Delete** /api/v1/integration/aws/logs | Delete an AWS Logs integration
[**EnableAWSLogServices**](AWSLogsIntegrationApi.md#EnableAWSLogServices) | **Post** /api/v1/integration/aws/logs/services | Enable an AWS Logs integration
[**ListAWSLogsIntegrations**](AWSLogsIntegrationApi.md#ListAWSLogsIntegrations) | **Get** /api/v1/integration/aws/logs | List all AWS Logs integrations
[**ListAWSLogsServices**](AWSLogsIntegrationApi.md#ListAWSLogsServices) | **Get** /api/v1/integration/aws/logs/services | Get list of AWS log ready services



## CheckAWSLogsLambdaAsync

> AWSLogsAsyncResponse CheckAWSLogsLambdaAsync(ctx, body)

Test if permissions are present to add a log-forwarding triggers for the given services and AWS account. The input
is the same as for Enable an AWS service log collection. Subsequent requests will always repeat the above, so this
endpoint can be polled intermittently instead of blocking.

- Returns a status of 'created' when it's checking if the Lambda exists in the account.
- Returns a status of 'waiting' while checking.
- Returns a status of 'checked and ok' if the Lambda exists.
- Returns a status of 'error' if the Lambda does not exist.

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

    body := *datadog.NewAWSAccountAndLambdaRequest("1234567", "arn:aws:lambda:us-east-1:1234567:function:LogsCollectionAPITest") // AWSAccountAndLambdaRequest | Check AWS Log Lambda Async request body.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.AWSLogsIntegrationApi.CheckAWSLogsLambdaAsync(ctx, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AWSLogsIntegrationApi.CheckAWSLogsLambdaAsync`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckAWSLogsLambdaAsync`: AWSLogsAsyncResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from AWSLogsIntegrationApi.CheckAWSLogsLambdaAsync:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**body** | [**AWSAccountAndLambdaRequest**](AWSAccountAndLambdaRequest.md) | Check AWS Log Lambda Async request body. | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**AWSLogsAsyncResponse**](AWSLogsAsyncResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckAWSLogsServicesAsync

> AWSLogsAsyncResponse CheckAWSLogsServicesAsync(ctx, body)

Test if permissions are present to add log-forwarding triggers for the
given services and AWS account. Input is the same as for `EnableAWSLogServices`.
Done async, so can be repeatedly polled in a non-blocking fashion until
the async request completes.

- Returns a status of `created` when it's checking if the permissions exists
  in the AWS account.
- Returns a status of `waiting` while checking.
- Returns a status of `checked and ok` if the Lambda exists.
- Returns a status of `error` if the Lambda does not exist.

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

    body := *datadog.NewAWSLogsServicesRequest("1234567", []string{"Services_example"}) // AWSLogsServicesRequest | Check AWS Logs Async Services request body.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.AWSLogsIntegrationApi.CheckAWSLogsServicesAsync(ctx, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AWSLogsIntegrationApi.CheckAWSLogsServicesAsync`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckAWSLogsServicesAsync`: AWSLogsAsyncResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from AWSLogsIntegrationApi.CheckAWSLogsServicesAsync:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**body** | [**AWSLogsServicesRequest**](AWSLogsServicesRequest.md) | Check AWS Logs Async Services request body. | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**AWSLogsAsyncResponse**](AWSLogsAsyncResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAWSLambdaARN

> interface{} CreateAWSLambdaARN(ctx, body)

Attach the Lambda ARN of the Lambda created for the Datadog-AWS log collection to your AWS account ID to enable log collection.

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

    body := *datadog.NewAWSAccountAndLambdaRequest("1234567", "arn:aws:lambda:us-east-1:1234567:function:LogsCollectionAPITest") // AWSAccountAndLambdaRequest | AWS Log Lambda Async request body.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.AWSLogsIntegrationApi.CreateAWSLambdaARN(ctx, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AWSLogsIntegrationApi.CreateAWSLambdaARN`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAWSLambdaARN`: interface{}
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from AWSLogsIntegrationApi.CreateAWSLambdaARN:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**body** | [**AWSAccountAndLambdaRequest**](AWSAccountAndLambdaRequest.md) | AWS Log Lambda Async request body. | 


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


## DeleteAWSLambdaARN

> interface{} DeleteAWSLambdaARN(ctx, body)

Delete a Datadog-AWS logs configuration by removing the specific Lambda ARN associated with a given AWS account.

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

    body := *datadog.NewAWSAccountAndLambdaRequest("1234567", "arn:aws:lambda:us-east-1:1234567:function:LogsCollectionAPITest") // AWSAccountAndLambdaRequest | Delete AWS Lambda ARN request body.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.AWSLogsIntegrationApi.DeleteAWSLambdaARN(ctx, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AWSLogsIntegrationApi.DeleteAWSLambdaARN`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAWSLambdaARN`: interface{}
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from AWSLogsIntegrationApi.DeleteAWSLambdaARN:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**body** | [**AWSAccountAndLambdaRequest**](AWSAccountAndLambdaRequest.md) | Delete AWS Lambda ARN request body. | 


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


## EnableAWSLogServices

> interface{} EnableAWSLogServices(ctx, body)

Enable automatic log collection for a list of services. This should be run after running `CreateAWSLambdaARN` to save the configuration.

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

    body := *datadog.NewAWSLogsServicesRequest("1234567", []string{"Services_example"}) // AWSLogsServicesRequest | Enable AWS Log Services request body.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.AWSLogsIntegrationApi.EnableAWSLogServices(ctx, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AWSLogsIntegrationApi.EnableAWSLogServices`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnableAWSLogServices`: interface{}
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from AWSLogsIntegrationApi.EnableAWSLogServices:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**body** | [**AWSLogsServicesRequest**](AWSLogsServicesRequest.md) | Enable AWS Log Services request body. | 


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


## ListAWSLogsIntegrations

> []AWSLogsListResponse ListAWSLogsIntegrations(ctx)

List all Datadog-AWS Logs integrations configured in your Datadog account.

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
    resp, r, err := apiClient.AWSLogsIntegrationApi.ListAWSLogsIntegrations(ctx)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AWSLogsIntegrationApi.ListAWSLogsIntegrations`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAWSLogsIntegrations`: []AWSLogsListResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from AWSLogsIntegrationApi.ListAWSLogsIntegrations:\n%s\n", responseContent)
}
```

### Required Parameters

This endpoint does not need any parameter.


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**[]AWSLogsListResponse**](AWSLogsListResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAWSLogsServices

> []AWSLogsListServicesResponse ListAWSLogsServices(ctx)

Get the list of current AWS services that Datadog offers automatic log collection. Use returned service IDs with the services parameter for the Enable an AWS service log collection API endpoint.

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
    resp, r, err := apiClient.AWSLogsIntegrationApi.ListAWSLogsServices(ctx)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AWSLogsIntegrationApi.ListAWSLogsServices`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAWSLogsServices`: []AWSLogsListServicesResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from AWSLogsIntegrationApi.ListAWSLogsServices:\n%s\n", responseContent)
}
```

### Required Parameters

This endpoint does not need any parameter.


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**[]AWSLogsListServicesResponse**](AWSLogsListServicesResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

