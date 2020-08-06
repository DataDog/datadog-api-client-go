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

> SecurityMonitoringRuleResponse CreateSecurityMonitoringRule(ctx).Body(body).Execute()

Create a detection rule



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

    body := datadog.SecurityMonitoringRuleCreatePayload{Cases: []SecurityMonitoringRuleCaseCreate{datadog.SecurityMonitoringRuleCaseCreate{Condition: "Condition_example", Name: "Name_example", Notifications: []string{"Notifications_example"), Status: datadog.SecurityMonitoringRuleSeverity{}}), IsEnabled: true, Message: "Message_example", Name: "Name_example", Options: datadog.SecurityMonitoringRuleOptions{EvaluationWindow: datadog.SecurityMonitoringRuleEvaluationWindow{}, KeepAlive: datadog.SecurityMonitoringRuleKeepAlive{}, MaxSignalDuration: datadog.SecurityMonitoringRuleMaxSignalDuration{}}, Queries: []SecurityMonitoringRuleQueryCreate{datadog.SecurityMonitoringRuleQueryCreate{DistinctFields: []string{"DistinctFields_example"), GroupByFields: []string{"GroupByFields_example"), Name: "Name_example", Query: "Query_example"}), Tags: []string{"Tags_example")} // SecurityMonitoringRuleCreatePayload |  (optional)

    configuration := datadog.NewConfiguration()
    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.SecurityMonitoringApi.CreateSecurityMonitoringRule(ctx).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.CreateSecurityMonitoringRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSecurityMonitoringRule`: SecurityMonitoringRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.CreateSecurityMonitoringRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSecurityMonitoringRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SecurityMonitoringRuleCreatePayload**](SecurityMonitoringRuleCreatePayload.md) |  | 

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

> DeleteSecurityMonitoringRule(ctx, ruleId).Execute()

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

    ruleId := "ruleId_example" // string | The ID of the rule.

    configuration := datadog.NewConfiguration()
    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.SecurityMonitoringApi.DeleteSecurityMonitoringRule(ctx, ruleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.DeleteSecurityMonitoringRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **string** | The ID of the rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSecurityMonitoringRuleRequest struct via the builder pattern


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


## GetSecurityMonitoringRule

> SecurityMonitoringRuleResponse GetSecurityMonitoringRule(ctx, ruleId).Execute()

Get a rule's details



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

    ruleId := "ruleId_example" // string | The ID of the rule.

    configuration := datadog.NewConfiguration()
    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.SecurityMonitoringApi.GetSecurityMonitoringRule(ctx, ruleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.GetSecurityMonitoringRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecurityMonitoringRule`: SecurityMonitoringRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.GetSecurityMonitoringRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **string** | The ID of the rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecurityMonitoringRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> SecurityMonitoringListRulesResponse ListSecurityMonitoringRules(ctx).PageSize(pageSize).PageNumber(pageNumber).Execute()

List rules



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

    configuration := datadog.NewConfiguration()
    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.SecurityMonitoringApi.ListSecurityMonitoringRules(ctx).PageSize(pageSize).PageNumber(pageNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.ListSecurityMonitoringRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSecurityMonitoringRules`: SecurityMonitoringListRulesResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.ListSecurityMonitoringRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSecurityMonitoringRulesRequest struct via the builder pattern


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

> SecurityMonitoringSignalsListResponse ListSecurityMonitoringSignals(ctx).FilterQuery(filterQuery).FilterFrom(filterFrom).FilterTo(filterTo).Sort(sort).PageCursor(pageCursor).PageLimit(pageLimit).Execute()

Get a quick list of security signals



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

    filterQuery := "filterQuery_example" // string | The search query for security signals. (optional)
    filterFrom := Get-Date // time.Time | The minimum timestamp for requested security signals. (optional)
    filterTo := Get-Date // time.Time | The maximum timestamp for requested security signals. (optional)
    sort := datadog.SecurityMonitoringSignalsSort{} // SecurityMonitoringSignalsSort | The order of the security signals in results. (optional)
    pageCursor := "pageCursor_example" // string | A list of results using the cursor provided in the previous query. (optional)
    pageLimit := 987 // int32 | The maximum number of security signals in the response. (optional) (default to 10)

    configuration := datadog.NewConfiguration()
    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.SecurityMonitoringApi.ListSecurityMonitoringSignals(ctx).FilterQuery(filterQuery).FilterFrom(filterFrom).FilterTo(filterTo).Sort(sort).PageCursor(pageCursor).PageLimit(pageLimit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.ListSecurityMonitoringSignals``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSecurityMonitoringSignals`: SecurityMonitoringSignalsListResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.ListSecurityMonitoringSignals`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSecurityMonitoringSignalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterQuery** | **string** | The search query for security signals. | 
 **filterFrom** | **time.Time** | The minimum timestamp for requested security signals. | 
 **filterTo** | **time.Time** | The maximum timestamp for requested security signals. | 
 **sort** | [**SecurityMonitoringSignalsSort**](.md) | The order of the security signals in results. | 
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

> SecurityMonitoringSignalsListResponse SearchSecurityMonitoringSignals(ctx).Body(body).Execute()

Get a list of security signals



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

    body := datadog.SecurityMonitoringSignalListRequest{Filter: datadog.SecurityMonitoringSignalListRequest_filter{From: "TODO", Query: "Query_example", To: "TODO"}, Page: datadog.SecurityMonitoringSignalListRequest_page{Cursor: "Cursor_example", Limit: 123}, Sort: datadog.SecurityMonitoringSignalsSort{}} // SecurityMonitoringSignalListRequest |  (optional)

    configuration := datadog.NewConfiguration()
    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.SecurityMonitoringApi.SearchSecurityMonitoringSignals(ctx).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.SearchSecurityMonitoringSignals``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchSecurityMonitoringSignals`: SecurityMonitoringSignalsListResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.SearchSecurityMonitoringSignals`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchSecurityMonitoringSignalsRequest struct via the builder pattern


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

> SecurityMonitoringRuleResponse UpdateSecurityMonitoringRule(ctx, ruleId).Body(body).Execute()

Update an existing rule



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

    ruleId := "ruleId_example" // string | The ID of the rule.
    body := datadog.SecurityMonitoringRuleUpdatePayload{Cases: []SecurityMonitoringRuleCase{datadog.SecurityMonitoringRuleCase{Condition: "Condition_example", Name: "Name_example", Notifications: []string{"Notifications_example"), Status: datadog.SecurityMonitoringRuleSeverity{}}), IsEnabled: false, Message: "Message_example", Name: "Name_example", Options: datadog.SecurityMonitoringRuleOptions{EvaluationWindow: datadog.SecurityMonitoringRuleEvaluationWindow{}, KeepAlive: datadog.SecurityMonitoringRuleKeepAlive{}, MaxSignalDuration: datadog.SecurityMonitoringRuleMaxSignalDuration{}}, Queries: []SecurityMonitoringRuleQuery{datadog.SecurityMonitoringRuleQuery{DistinctFields: []string{"DistinctFields_example"), GroupByFields: []string{"GroupByFields_example"), Name: "Name_example", Query: "Query_example"}), Tags: []string{"Tags_example")} // SecurityMonitoringRuleUpdatePayload |  (optional)

    configuration := datadog.NewConfiguration()
    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.SecurityMonitoringApi.UpdateSecurityMonitoringRule(ctx, ruleId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.UpdateSecurityMonitoringRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSecurityMonitoringRule`: SecurityMonitoringRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.UpdateSecurityMonitoringRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **string** | The ID of the rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSecurityMonitoringRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**SecurityMonitoringRuleUpdatePayload**](SecurityMonitoringRuleUpdatePayload.md) |  | 

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

