# \SecurityMonitoringApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSecurityMonitoringRule**](SecurityMonitoringApi.md#CreateSecurityMonitoringRule) | **Post** /api/v2/security_monitoring/rules | Create a detection rule
[**DeleteSecurityMonitoringRule**](SecurityMonitoringApi.md#DeleteSecurityMonitoringRule) | **Delete** /api/v2/security_monitoring/rules/{rule_id} | Delete an existing rule
[**GetSecurityMonitoringRule**](SecurityMonitoringApi.md#GetSecurityMonitoringRule) | **Get** /api/v2/security_monitoring/rules/{rule_id} | Get a rule&#39;s details
[**ListSecurityMonitoringRules**](SecurityMonitoringApi.md#ListSecurityMonitoringRules) | **Get** /api/v2/security_monitoring/rules | List rules
[**ListSecurityMonitoringSignals**](SecurityMonitoringApi.md#ListSecurityMonitoringSignals) | **Get** /api/v2/security_monitoring/signals | Get a quick list of security signals
[**SearchSecurityMonitoringSignals**](SecurityMonitoringApi.md#SearchSecurityMonitoringSignals) | **Post** /api/v2/security_monitoring/signals/search | Get a list of security signals
[**UpdateSecurityMonitoringRule**](SecurityMonitoringApi.md#UpdateSecurityMonitoringRule) | **Put** /api/v2/security_monitoring/rules/{rule_id} | Update an existing rule



## CreateSecurityMonitoringRule

> SecurityMonitoringRuleResponse CreateSecurityMonitoringRule(ctx, body)

Create a detection rule



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

    body := *datadog.NewSecurityMonitoringRuleCreatePayload([]datadog.SecurityMonitoringRuleCaseCreate{*datadog.NewSecurityMonitoringRuleCaseCreate(datadog.SecurityMonitoringRuleSeverity("info"))}, true, "Message_example", "Name_example", *datadog.NewSecurityMonitoringRuleOptions(), []datadog.SecurityMonitoringRuleQueryCreate{*datadog.NewSecurityMonitoringRuleQueryCreate("a < 3")}) // SecurityMonitoringRuleCreatePayload | 

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityMonitoringApi.CreateSecurityMonitoringRule(ctx, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.CreateSecurityMonitoringRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSecurityMonitoringRule`: SecurityMonitoringRuleResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from SecurityMonitoringApi.CreateSecurityMonitoringRule:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**SecurityMonitoringRuleCreatePayload**](SecurityMonitoringRuleCreatePayload.md) |  | 

### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**SecurityMonitoringRuleResponse**](SecurityMonitoringRuleResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSecurityMonitoringRule

> DeleteSecurityMonitoringRule(ctx, ruleId)

Delete an existing rule



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

    ruleId := "ruleId_example" // string | The ID of the rule.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    r, err := apiClient.SecurityMonitoringApi.DeleteSecurityMonitoringRule(ctx, ruleId)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.DeleteSecurityMonitoringRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **string** | The ID of the rule. | 

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


## GetSecurityMonitoringRule

> SecurityMonitoringRuleResponse GetSecurityMonitoringRule(ctx, ruleId)

Get a rule's details



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

    ruleId := "ruleId_example" // string | The ID of the rule.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityMonitoringApi.GetSecurityMonitoringRule(ctx, ruleId)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.GetSecurityMonitoringRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecurityMonitoringRule`: SecurityMonitoringRuleResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from SecurityMonitoringApi.GetSecurityMonitoringRule:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **string** | The ID of the rule. | 

### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**SecurityMonitoringRuleResponse**](SecurityMonitoringRuleResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSecurityMonitoringRules

> SecurityMonitoringListRulesResponse ListSecurityMonitoringRules(ctx, datadog.ListSecurityMonitoringRulesOptionalParameters{})

List rules



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

    pageSize := int64(789) // int64 | Size for a given page. (optional) (default to 10)
    pageNumber := int64(789) // int64 | Specific page number to return. (optional) (default to 0)
    optionalParams := datadog.ListSecurityMonitoringRulesOptionalParameters{
        PageSize: &pageSize,
        PageNumber: &pageNumber,
    }

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityMonitoringApi.ListSecurityMonitoringRules(ctx, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.ListSecurityMonitoringRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSecurityMonitoringRules`: SecurityMonitoringListRulesResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from SecurityMonitoringApi.ListSecurityMonitoringRules:\n%s\n", responseContent)
}
```

### Required Parameters



### Optional Parameters


Other parameters are passed through a pointer to a ListSecurityMonitoringRulesOptionalParameters struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**pageSize** | **int64** | Size for a given page. | [default to 10]
**pageNumber** | **int64** | Specific page number to return. | [default to 0]

### Return type

[**SecurityMonitoringListRulesResponse**](SecurityMonitoringListRulesResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSecurityMonitoringSignals

> SecurityMonitoringSignalsListResponse ListSecurityMonitoringSignals(ctx, datadog.ListSecurityMonitoringSignalsOptionalParameters{})

Get a quick list of security signals



### Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    "time"
    datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
    ctx := datadog.NewDefaultContext(context.Background())

    filterQuery := "security:attack status:high" // string | The search query for security signals. (optional)
    filterFrom := time.Now() // time.Time | The minimum timestamp for requested security signals. (optional)
    filterTo := time.Now() // time.Time | The maximum timestamp for requested security signals. (optional)
    sort := datadog.SecurityMonitoringSignalsSort("timestamp") // SecurityMonitoringSignalsSort | The order of the security signals in results. (optional)
    pageCursor := "eyJzdGFydEF0IjoiQVFBQUFYS2tMS3pPbm40NGV3QUFBQUJCV0V0clRFdDZVbG8zY3pCRmNsbHJiVmxDWlEifQ==" // string | A list of results using the cursor provided in the previous query. (optional)
    pageLimit := int32(25) // int32 | The maximum number of security signals in the response. (optional) (default to 10)
    optionalParams := datadog.ListSecurityMonitoringSignalsOptionalParameters{
        FilterQuery: &filterQuery,
        FilterFrom: &filterFrom,
        FilterTo: &filterTo,
        Sort: &sort,
        PageCursor: &pageCursor,
        PageLimit: &pageLimit,
    }

    configuration := datadog.NewConfiguration()
    configuration.SetUnstableOperationEnabled("ListSecurityMonitoringSignals", true)

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityMonitoringApi.ListSecurityMonitoringSignals(ctx, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.ListSecurityMonitoringSignals``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSecurityMonitoringSignals`: SecurityMonitoringSignalsListResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from SecurityMonitoringApi.ListSecurityMonitoringSignals:\n%s\n", responseContent)
}
```

### Required Parameters



### Optional Parameters


Other parameters are passed through a pointer to a ListSecurityMonitoringSignalsOptionalParameters struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**filterQuery** | **string** | The search query for security signals. | 
**filterFrom** | **time.Time** | The minimum timestamp for requested security signals. | 
**filterTo** | **time.Time** | The maximum timestamp for requested security signals. | 
**sort** | [**SecurityMonitoringSignalsSort**](SecurityMonitoringSignalsSort.md) | The order of the security signals in results. | 
**pageCursor** | **string** | A list of results using the cursor provided in the previous query. | 
**pageLimit** | **int32** | The maximum number of security signals in the response. | [default to 10]

### Return type

[**SecurityMonitoringSignalsListResponse**](SecurityMonitoringSignalsListResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchSecurityMonitoringSignals

> SecurityMonitoringSignalsListResponse SearchSecurityMonitoringSignals(ctx, datadog.SearchSecurityMonitoringSignalsOptionalParameters{})

Get a list of security signals



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

    body := *datadog.NewSecurityMonitoringSignalListRequest() // SecurityMonitoringSignalListRequest |  (optional)
    optionalParams := datadog.SearchSecurityMonitoringSignalsOptionalParameters{
        Body: &body,
    }

    configuration := datadog.NewConfiguration()
    configuration.SetUnstableOperationEnabled("SearchSecurityMonitoringSignals", true)

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityMonitoringApi.SearchSecurityMonitoringSignals(ctx, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.SearchSecurityMonitoringSignals``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchSecurityMonitoringSignals`: SecurityMonitoringSignalsListResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from SecurityMonitoringApi.SearchSecurityMonitoringSignals:\n%s\n", responseContent)
}
```

### Required Parameters



### Optional Parameters


Other parameters are passed through a pointer to a SearchSecurityMonitoringSignalsOptionalParameters struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**SecurityMonitoringSignalListRequest**](SecurityMonitoringSignalListRequest.md) |  | 

### Return type

[**SecurityMonitoringSignalsListResponse**](SecurityMonitoringSignalsListResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSecurityMonitoringRule

> SecurityMonitoringRuleResponse UpdateSecurityMonitoringRule(ctx, ruleId, body)

Update an existing rule



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

    ruleId := "ruleId_example" // string | The ID of the rule.
    body := *datadog.NewSecurityMonitoringRuleUpdatePayload() // SecurityMonitoringRuleUpdatePayload | 

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityMonitoringApi.UpdateSecurityMonitoringRule(ctx, ruleId, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.UpdateSecurityMonitoringRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSecurityMonitoringRule`: SecurityMonitoringRuleResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from SecurityMonitoringApi.UpdateSecurityMonitoringRule:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **string** | The ID of the rule. | 
**body** | [**SecurityMonitoringRuleUpdatePayload**](SecurityMonitoringRuleUpdatePayload.md) |  | 

### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**SecurityMonitoringRuleResponse**](SecurityMonitoringRuleResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

