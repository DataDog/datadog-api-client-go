# WebhooksIntegrationApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------ | ------------ | ------------
[**CreateWebhooksIntegration**](WebhooksIntegrationApi.md#CreateWebhooksIntegration) | **Post** /api/v1/integration/webhooks/configuration/webhooks | Create a webhooks integration
[**CreateWebhooksIntegrationCustomVariable**](WebhooksIntegrationApi.md#CreateWebhooksIntegrationCustomVariable) | **Post** /api/v1/integration/webhooks/configuration/custom-variables | Create a custom variable
[**DeleteWebhooksIntegration**](WebhooksIntegrationApi.md#DeleteWebhooksIntegration) | **Delete** /api/v1/integration/webhooks/configuration/webhooks/{webhook_name} | Delete a webhook
[**DeleteWebhooksIntegrationCustomVariable**](WebhooksIntegrationApi.md#DeleteWebhooksIntegrationCustomVariable) | **Delete** /api/v1/integration/webhooks/configuration/custom-variables/{custom_variable_name} | Delete a custom variable
[**GetWebhooksIntegration**](WebhooksIntegrationApi.md#GetWebhooksIntegration) | **Get** /api/v1/integration/webhooks/configuration/webhooks/{webhook_name} | Get a webhook integration
[**GetWebhooksIntegrationCustomVariable**](WebhooksIntegrationApi.md#GetWebhooksIntegrationCustomVariable) | **Get** /api/v1/integration/webhooks/configuration/custom-variables/{custom_variable_name} | Get a custom variable
[**UpdateWebhooksIntegration**](WebhooksIntegrationApi.md#UpdateWebhooksIntegration) | **Put** /api/v1/integration/webhooks/configuration/webhooks/{webhook_name} | Update a webhook
[**UpdateWebhooksIntegrationCustomVariable**](WebhooksIntegrationApi.md#UpdateWebhooksIntegrationCustomVariable) | **Put** /api/v1/integration/webhooks/configuration/custom-variables/{custom_variable_name} | Update a custom variable



## CreateWebhooksIntegration

> WebhooksIntegration CreateWebhooksIntegration(ctx, body)

Creates an endpoint with the name `<WEBHOOK_NAME>`.

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

    body := *datadog.NewWebhooksIntegration("WEBHOOK_NAME", "https://example.com/webhook") // WebhooksIntegration | Create a webhooks integration request body.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksIntegrationApi.CreateWebhooksIntegration(ctx, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksIntegrationApi.CreateWebhooksIntegration`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWebhooksIntegration`: WebhooksIntegration
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from WebhooksIntegrationApi.CreateWebhooksIntegration:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**body** | [**WebhooksIntegration**](WebhooksIntegration.md) | Create a webhooks integration request body. | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**WebhooksIntegration**](WebhooksIntegration.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWebhooksIntegrationCustomVariable

> WebhooksIntegrationCustomVariableResponse CreateWebhooksIntegrationCustomVariable(ctx, body)

Creates an endpoint with the name `<CUSTOM_VARIABLE_NAME>`.

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

    body := *datadog.NewWebhooksIntegrationCustomVariable(true, "CUSTOM_VARIABLE_NAME", "CUSTOM_VARIABLE_VALUE") // WebhooksIntegrationCustomVariable | Define a custom variable request body.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksIntegrationApi.CreateWebhooksIntegrationCustomVariable(ctx, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksIntegrationApi.CreateWebhooksIntegrationCustomVariable`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWebhooksIntegrationCustomVariable`: WebhooksIntegrationCustomVariableResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from WebhooksIntegrationApi.CreateWebhooksIntegrationCustomVariable:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**body** | [**WebhooksIntegrationCustomVariable**](WebhooksIntegrationCustomVariable.md) | Define a custom variable request body. | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**WebhooksIntegrationCustomVariableResponse**](WebhooksIntegrationCustomVariableResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWebhooksIntegration

> DeleteWebhooksIntegration(ctx, webhookName)

Deletes the endpoint with the name `<WEBHOOK NAME>`.

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
    ctx := datadog.NewDefaultContext(context.Background())

    webhookName := "webhookName_example" // string | The name of the webhook.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    r, err := apiClient.WebhooksIntegrationApi.DeleteWebhooksIntegration(ctx, webhookName)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksIntegrationApi.DeleteWebhooksIntegration`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**webhookName** | **string** | The name of the webhook. | 


### Optional Parameters

This endpoint does not have optional parameters.


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


## DeleteWebhooksIntegrationCustomVariable

> DeleteWebhooksIntegrationCustomVariable(ctx, customVariableName)

Deletes the endpoint with the name `<CUSTOM_VARIABLE_NAME>`.

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
    ctx := datadog.NewDefaultContext(context.Background())

    customVariableName := "customVariableName_example" // string | The name of the custom variable.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    r, err := apiClient.WebhooksIntegrationApi.DeleteWebhooksIntegrationCustomVariable(ctx, customVariableName)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksIntegrationApi.DeleteWebhooksIntegrationCustomVariable`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**customVariableName** | **string** | The name of the custom variable. | 


### Optional Parameters

This endpoint does not have optional parameters.


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


## GetWebhooksIntegration

> WebhooksIntegration GetWebhooksIntegration(ctx, webhookName)

Gets the content of the webhook with the name `<WEBHOOK_NAME>`.

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

    webhookName := "webhookName_example" // string | The name of the webhook.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksIntegrationApi.GetWebhooksIntegration(ctx, webhookName)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksIntegrationApi.GetWebhooksIntegration`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWebhooksIntegration`: WebhooksIntegration
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from WebhooksIntegrationApi.GetWebhooksIntegration:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**webhookName** | **string** | The name of the webhook. | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**WebhooksIntegration**](WebhooksIntegration.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhooksIntegrationCustomVariable

> WebhooksIntegrationCustomVariableResponse GetWebhooksIntegrationCustomVariable(ctx, customVariableName)

Shows the content of the custom variable with the name `<CUSTOM_VARIABLE_NAME>`.

If the custom variable is secret, the value does not return in the
response payload.

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

    customVariableName := "customVariableName_example" // string | The name of the custom variable.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksIntegrationApi.GetWebhooksIntegrationCustomVariable(ctx, customVariableName)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksIntegrationApi.GetWebhooksIntegrationCustomVariable`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWebhooksIntegrationCustomVariable`: WebhooksIntegrationCustomVariableResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from WebhooksIntegrationApi.GetWebhooksIntegrationCustomVariable:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**customVariableName** | **string** | The name of the custom variable. | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**WebhooksIntegrationCustomVariableResponse**](WebhooksIntegrationCustomVariableResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWebhooksIntegration

> WebhooksIntegration UpdateWebhooksIntegration(ctx, webhookName, body)

Updates the endpoint with the name `<WEBHOOK_NAME>`.

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

    webhookName := "webhookName_example" // string | The name of the webhook.
    body := *datadog.NewWebhooksIntegrationUpdateRequest() // WebhooksIntegrationUpdateRequest | Update an existing Datadog-Webhooks integration.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksIntegrationApi.UpdateWebhooksIntegration(ctx, webhookName, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksIntegrationApi.UpdateWebhooksIntegration`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWebhooksIntegration`: WebhooksIntegration
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from WebhooksIntegrationApi.UpdateWebhooksIntegration:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**webhookName** | **string** | The name of the webhook. |  |
**body** | [**WebhooksIntegrationUpdateRequest**](WebhooksIntegrationUpdateRequest.md) | Update an existing Datadog-Webhooks integration. | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**WebhooksIntegration**](WebhooksIntegration.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWebhooksIntegrationCustomVariable

> WebhooksIntegrationCustomVariableResponse UpdateWebhooksIntegrationCustomVariable(ctx, customVariableName, body)

Updates the endpoint with the name `<CUSTOM_VARIABLE_NAME>`.

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

    customVariableName := "customVariableName_example" // string | The name of the custom variable.
    body := *datadog.NewWebhooksIntegrationCustomVariableUpdateRequest() // WebhooksIntegrationCustomVariableUpdateRequest | Update an existing custom variable request body.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksIntegrationApi.UpdateWebhooksIntegrationCustomVariable(ctx, customVariableName, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksIntegrationApi.UpdateWebhooksIntegrationCustomVariable`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWebhooksIntegrationCustomVariable`: WebhooksIntegrationCustomVariableResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from WebhooksIntegrationApi.UpdateWebhooksIntegrationCustomVariable:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**customVariableName** | **string** | The name of the custom variable. |  |
**body** | [**WebhooksIntegrationCustomVariableUpdateRequest**](WebhooksIntegrationCustomVariableUpdateRequest.md) | Update an existing custom variable request body. | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**WebhooksIntegrationCustomVariableResponse**](WebhooksIntegrationCustomVariableResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

