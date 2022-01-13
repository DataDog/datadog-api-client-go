# CloudWorkloadSecurityApi

All URIs are relative to *https://api.datadoghq.com*

| Method                                                                                                       | HTTP request                                                                               | Description                                   |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | --------------------------------------------- |
| [**CreateCloudWorkloadSecurityAgentRule**](CloudWorkloadSecurityApi.md#CreateCloudWorkloadSecurityAgentRule) | **Post** /api/v2/security_monitoring/cloud_workload_security/agent_rules                   | Create a Cloud Workload Security Agent rule   |
| [**DeleteCloudWorkloadSecurityAgentRule**](CloudWorkloadSecurityApi.md#DeleteCloudWorkloadSecurityAgentRule) | **Delete** /api/v2/security_monitoring/cloud_workload_security/agent_rules/{agent_rule_id} | Delete a Cloud Workload Security Agent rule   |
| [**DownloadCloudWorkloadPolicyFile**](CloudWorkloadSecurityApi.md#DownloadCloudWorkloadPolicyFile)           | **Get** /api/v2/security/cloud_workload/policy/download                                    | Get the latest Cloud Workload Security policy |
| [**GetCloudWorkloadSecurityAgentRule**](CloudWorkloadSecurityApi.md#GetCloudWorkloadSecurityAgentRule)       | **Get** /api/v2/security_monitoring/cloud_workload_security/agent_rules/{agent_rule_id}    | Get a Cloud Workload Security Agent rule      |
| [**ListCloudWorkloadSecurityAgentRules**](CloudWorkloadSecurityApi.md#ListCloudWorkloadSecurityAgentRules)   | **Get** /api/v2/security_monitoring/cloud_workload_security/agent_rules                    | Get all Cloud Workload Security Agent rules   |
| [**UpdateCloudWorkloadSecurityAgentRule**](CloudWorkloadSecurityApi.md#UpdateCloudWorkloadSecurityAgentRule) | **Patch** /api/v2/security_monitoring/cloud_workload_security/agent_rules/{agent_rule_id}  | Update a Cloud Workload Security Agent rule   |

## CreateCloudWorkloadSecurityAgentRule

> CloudWorkloadSecurityAgentRuleResponse CreateCloudWorkloadSecurityAgentRule(ctx, body)

Create a new Agent rule with the given parameters.

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

    body := *datadog.NewCloudWorkloadSecurityAgentRuleCreateRequest(*datadog.NewCloudWorkloadSecurityAgentRuleCreateData(*datadog.NewCloudWorkloadSecurityAgentRuleCreateAttributes("exec.file.name == \"sh\"", "my_agent_rule"), datadog.CloudWorkloadSecurityAgentRuleType("agent_rule"))) // CloudWorkloadSecurityAgentRuleCreateRequest | The definition of the new Agent rule.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudWorkloadSecurityApi.CreateCloudWorkloadSecurityAgentRule(ctx, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudWorkloadSecurityApi.CreateCloudWorkloadSecurityAgentRule`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCloudWorkloadSecurityAgentRule`: CloudWorkloadSecurityAgentRuleResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from CloudWorkloadSecurityApi.CreateCloudWorkloadSecurityAgentRule:\n%s\n", responseContent)
}
```

### Required Parameters

| Name     | Type                                                                                              | Description                                                                 | Notes |
| -------- | ------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**  | **context.Context**                                                                               | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **body** | [**CloudWorkloadSecurityAgentRuleCreateRequest**](CloudWorkloadSecurityAgentRuleCreateRequest.md) | The definition of the new Agent rule.                                       |

### Optional Parameters

This endpoint does not have optional parameters.

### Return type

[**CloudWorkloadSecurityAgentRuleResponse**](CloudWorkloadSecurityAgentRuleResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## DeleteCloudWorkloadSecurityAgentRule

> DeleteCloudWorkloadSecurityAgentRule(ctx, agentRuleId)

Delete a specific Agent rule.

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

    agentRuleId := "3b5-v82-ns6" // string | The ID of the Agent rule.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    r, err := apiClient.CloudWorkloadSecurityApi.DeleteCloudWorkloadSecurityAgentRule(ctx, agentRuleId)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudWorkloadSecurityApi.DeleteCloudWorkloadSecurityAgentRule`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Required Parameters

| Name            | Type                | Description                                                                 | Notes |
| --------------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**         | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **agentRuleId** | **string**          | The ID of the Agent rule.                                                   |

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

## DownloadCloudWorkloadPolicyFile

> \*os.File DownloadCloudWorkloadPolicyFile(ctx)

The download endpoint generates a Cloud Workload Security policy file from your currently active
Cloud Workload Security rules, and downloads them as a .policy file. This file can then be deployed to
your agents to update the policy running in your environment.

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
    resp, r, err := apiClient.CloudWorkloadSecurityApi.DownloadCloudWorkloadPolicyFile(ctx)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudWorkloadSecurityApi.DownloadCloudWorkloadPolicyFile`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadCloudWorkloadPolicyFile`: *os.File
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from CloudWorkloadSecurityApi.DownloadCloudWorkloadPolicyFile:\n%s\n", responseContent)
}
```

### Required Parameters

This endpoint does not need any parameter.

### Optional Parameters

This endpoint does not have optional parameters.

### Return type

[**\*os.File**](*os.File.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/yaml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetCloudWorkloadSecurityAgentRule

> CloudWorkloadSecurityAgentRuleResponse GetCloudWorkloadSecurityAgentRule(ctx, agentRuleId)

Get the details of a specific Agent rule.

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

    agentRuleId := "3b5-v82-ns6" // string | The ID of the Agent rule.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudWorkloadSecurityApi.GetCloudWorkloadSecurityAgentRule(ctx, agentRuleId)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudWorkloadSecurityApi.GetCloudWorkloadSecurityAgentRule`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCloudWorkloadSecurityAgentRule`: CloudWorkloadSecurityAgentRuleResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from CloudWorkloadSecurityApi.GetCloudWorkloadSecurityAgentRule:\n%s\n", responseContent)
}
```

### Required Parameters

| Name            | Type                | Description                                                                 | Notes |
| --------------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**         | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **agentRuleId** | **string**          | The ID of the Agent rule.                                                   |

### Optional Parameters

This endpoint does not have optional parameters.

### Return type

[**CloudWorkloadSecurityAgentRuleResponse**](CloudWorkloadSecurityAgentRuleResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListCloudWorkloadSecurityAgentRules

> CloudWorkloadSecurityAgentRulesListResponse ListCloudWorkloadSecurityAgentRules(ctx)

Get the list of Agent rules.

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
    resp, r, err := apiClient.CloudWorkloadSecurityApi.ListCloudWorkloadSecurityAgentRules(ctx)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudWorkloadSecurityApi.ListCloudWorkloadSecurityAgentRules`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCloudWorkloadSecurityAgentRules`: CloudWorkloadSecurityAgentRulesListResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from CloudWorkloadSecurityApi.ListCloudWorkloadSecurityAgentRules:\n%s\n", responseContent)
}
```

### Required Parameters

This endpoint does not need any parameter.

### Optional Parameters

This endpoint does not have optional parameters.

### Return type

[**CloudWorkloadSecurityAgentRulesListResponse**](CloudWorkloadSecurityAgentRulesListResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateCloudWorkloadSecurityAgentRule

> CloudWorkloadSecurityAgentRuleResponse UpdateCloudWorkloadSecurityAgentRule(ctx, agentRuleId, body)

Update a specific Agent rule.
Returns the Agent rule object when the request is successful.

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

    agentRuleId := "3b5-v82-ns6" // string | The ID of the Agent rule.
    body := *datadog.NewCloudWorkloadSecurityAgentRuleUpdateRequest(*datadog.NewCloudWorkloadSecurityAgentRuleUpdateData(*datadog.NewCloudWorkloadSecurityAgentRuleUpdateAttributes(), datadog.CloudWorkloadSecurityAgentRuleType("agent_rule"))) // CloudWorkloadSecurityAgentRuleUpdateRequest | New definition of the Agent rule.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudWorkloadSecurityApi.UpdateCloudWorkloadSecurityAgentRule(ctx, agentRuleId, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudWorkloadSecurityApi.UpdateCloudWorkloadSecurityAgentRule`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCloudWorkloadSecurityAgentRule`: CloudWorkloadSecurityAgentRuleResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from CloudWorkloadSecurityApi.UpdateCloudWorkloadSecurityAgentRule:\n%s\n", responseContent)
}
```

### Required Parameters

| Name            | Type                                                                                              | Description                                                                 | Notes |
| --------------- | ------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**         | **context.Context**                                                                               | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **agentRuleId** | **string**                                                                                        | The ID of the Agent rule.                                                   |       |
| **body**        | [**CloudWorkloadSecurityAgentRuleUpdateRequest**](CloudWorkloadSecurityAgentRuleUpdateRequest.md) | New definition of the Agent rule.                                           |

### Optional Parameters

This endpoint does not have optional parameters.

### Return type

[**CloudWorkloadSecurityAgentRuleResponse**](CloudWorkloadSecurityAgentRuleResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
