# CloudSIEMApi

All URIs are relative to *https://api.datadoghq.com*

| Method                                                                                 | HTTP request                                                                               | Description                          |
| -------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------ | ------------------------------------ |
| [**CreateSecurityFilter**](CloudSIEMApi.md#CreateSecurityFilter)                       | **Post** /api/v2/security_monitoring/configuration/security_filters                        | Create a security filter             |
| [**CreateSecurityMonitoringRule**](CloudSIEMApi.md#CreateSecurityMonitoringRule)       | **Post** /api/v2/security_monitoring/rules                                                 | Create a detection rule              |
| [**DeleteSecurityFilter**](CloudSIEMApi.md#DeleteSecurityFilter)                       | **Delete** /api/v2/security_monitoring/configuration/security_filters/{security_filter_id} | Delete a security filter             |
| [**DeleteSecurityMonitoringRule**](CloudSIEMApi.md#DeleteSecurityMonitoringRule)       | **Delete** /api/v2/security_monitoring/rules/{rule_id}                                     | Delete an existing rule              |
| [**GetSecurityFilter**](CloudSIEMApi.md#GetSecurityFilter)                             | **Get** /api/v2/security_monitoring/configuration/security_filters/{security_filter_id}    | Get a security filter                |
| [**GetSecurityMonitoringRule**](CloudSIEMApi.md#GetSecurityMonitoringRule)             | **Get** /api/v2/security_monitoring/rules/{rule_id}                                        | Get a rule&#39;s details             |
| [**ListSecurityFilters**](CloudSIEMApi.md#ListSecurityFilters)                         | **Get** /api/v2/security_monitoring/configuration/security_filters                         | Get all security filters             |
| [**ListSecurityMonitoringRules**](CloudSIEMApi.md#ListSecurityMonitoringRules)         | **Get** /api/v2/security_monitoring/rules                                                  | List rules                           |
| [**ListSecurityMonitoringSignals**](CloudSIEMApi.md#ListSecurityMonitoringSignals)     | **Get** /api/v2/security_monitoring/signals                                                | Get a quick list of security signals |
| [**SearchSecurityMonitoringSignals**](CloudSIEMApi.md#SearchSecurityMonitoringSignals) | **Post** /api/v2/security_monitoring/signals/search                                        | Get a list of security signals       |
| [**UpdateSecurityFilter**](CloudSIEMApi.md#UpdateSecurityFilter)                       | **Patch** /api/v2/security_monitoring/configuration/security_filters/{security_filter_id}  | Update a security filter             |
| [**UpdateSecurityMonitoringRule**](CloudSIEMApi.md#UpdateSecurityMonitoringRule)       | **Put** /api/v2/security_monitoring/rules/{rule_id}                                        | Update an existing rule              |

## CreateSecurityFilter

> SecurityFilterResponse CreateSecurityFilter(ctx, body)

Create a security filter.

See the [security filter guide](https://docs.datadoghq.com/security_platform/guide/how-to-setup-security-filters-using-security-monitoring-api/)
for more examples.

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

    body := *datadog.NewSecurityFilterCreateRequest(*datadog.NewSecurityFilterCreateData(*datadog.NewSecurityFilterCreateAttributes([]datadog.SecurityFilterExclusionFilter{*datadog.NewSecurityFilterExclusionFilter("Exclude staging", "source:staging")}, datadog.SecurityFilterFilteredDataType("logs"), true, "Custom security filter", "service:api"), datadog.SecurityFilterType("security_filters"))) // SecurityFilterCreateRequest | The definition of the new security filter.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudSIEMApi.CreateSecurityFilter(ctx, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudSIEMApi.CreateSecurityFilter`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSecurityFilter`: SecurityFilterResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from CloudSIEMApi.CreateSecurityFilter:\n%s\n", responseContent)
}
```

### Required Parameters

| Name     | Type                                                              | Description                                                                 | Notes |
| -------- | ----------------------------------------------------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**  | **context.Context**                                               | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **body** | [**SecurityFilterCreateRequest**](SecurityFilterCreateRequest.md) | The definition of the new security filter.                                  |

### Optional Parameters

This endpoint does not have optional parameters.

### Return type

[**SecurityFilterResponse**](SecurityFilterResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CreateSecurityMonitoringRule

> SecurityMonitoringRuleResponse CreateSecurityMonitoringRule(ctx, body)

Create a detection rule.

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

    body := *datadog.NewSecurityMonitoringRuleCreatePayload([]datadog.SecurityMonitoringRuleCaseCreate{*datadog.NewSecurityMonitoringRuleCaseCreate(datadog.SecurityMonitoringRuleSeverity("info"))}, true, "Message_example", "My Cloud SIEM rule.", *datadog.NewSecurityMonitoringRuleOptions(), []datadog.SecurityMonitoringRuleQueryCreate{*datadog.NewSecurityMonitoringRuleQueryCreate("a > 3")}) // SecurityMonitoringRuleCreatePayload |

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudSIEMApi.CreateSecurityMonitoringRule(ctx, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudSIEMApi.CreateSecurityMonitoringRule`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSecurityMonitoringRule`: SecurityMonitoringRuleResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from CloudSIEMApi.CreateSecurityMonitoringRule:\n%s\n", responseContent)
}
```

### Required Parameters

| Name     | Type                                                                              | Description                                                                 | Notes |
| -------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**  | **context.Context**                                                               | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **body** | [**SecurityMonitoringRuleCreatePayload**](SecurityMonitoringRuleCreatePayload.md) |                                                                             |

### Optional Parameters

This endpoint does not have optional parameters.

### Return type

[**SecurityMonitoringRuleResponse**](SecurityMonitoringRuleResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## DeleteSecurityFilter

> DeleteSecurityFilter(ctx, securityFilterId)

Delete a specific security filter.

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

    securityFilterId := "securityFilterId_example" // string | The ID of the security filter.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    r, err := apiClient.CloudSIEMApi.DeleteSecurityFilter(ctx, securityFilterId)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudSIEMApi.DeleteSecurityFilter`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Required Parameters

| Name                 | Type                | Description                                                                 | Notes |
| -------------------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**              | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **securityFilterId** | **string**          | The ID of the security filter.                                              |

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

## DeleteSecurityMonitoringRule

> DeleteSecurityMonitoringRule(ctx, ruleId)

Delete an existing rule. Default rules cannot be deleted.

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
    r, err := apiClient.CloudSIEMApi.DeleteSecurityMonitoringRule(ctx, ruleId)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudSIEMApi.DeleteSecurityMonitoringRule`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Required Parameters

| Name       | Type                | Description                                                                 | Notes |
| ---------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**    | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **ruleId** | **string**          | The ID of the rule.                                                         |

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

## GetSecurityFilter

> SecurityFilterResponse GetSecurityFilter(ctx, securityFilterId)

Get the details of a specific security filter.

See the [security filter guide](https://docs.datadoghq.com/security_platform/guide/how-to-setup-security-filters-using-security-monitoring-api/)
for more examples.

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

    securityFilterId := "securityFilterId_example" // string | The ID of the security filter.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudSIEMApi.GetSecurityFilter(ctx, securityFilterId)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudSIEMApi.GetSecurityFilter`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecurityFilter`: SecurityFilterResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from CloudSIEMApi.GetSecurityFilter:\n%s\n", responseContent)
}
```

### Required Parameters

| Name                 | Type                | Description                                                                 | Notes |
| -------------------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**              | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **securityFilterId** | **string**          | The ID of the security filter.                                              |

### Optional Parameters

This endpoint does not have optional parameters.

### Return type

[**SecurityFilterResponse**](SecurityFilterResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetSecurityMonitoringRule

> SecurityMonitoringRuleResponse GetSecurityMonitoringRule(ctx, ruleId)

Get a rule's details.

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
    resp, r, err := apiClient.CloudSIEMApi.GetSecurityMonitoringRule(ctx, ruleId)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudSIEMApi.GetSecurityMonitoringRule`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecurityMonitoringRule`: SecurityMonitoringRuleResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from CloudSIEMApi.GetSecurityMonitoringRule:\n%s\n", responseContent)
}
```

### Required Parameters

| Name       | Type                | Description                                                                 | Notes |
| ---------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**    | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **ruleId** | **string**          | The ID of the rule.                                                         |

### Optional Parameters

This endpoint does not have optional parameters.

### Return type

[**SecurityMonitoringRuleResponse**](SecurityMonitoringRuleResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListSecurityFilters

> SecurityFiltersResponse ListSecurityFilters(ctx)

Get the list of configured security filters with their definitions.

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
    resp, r, err := apiClient.CloudSIEMApi.ListSecurityFilters(ctx)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudSIEMApi.ListSecurityFilters`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSecurityFilters`: SecurityFiltersResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from CloudSIEMApi.ListSecurityFilters:\n%s\n", responseContent)
}
```

### Required Parameters

This endpoint does not need any parameter.

### Optional Parameters

This endpoint does not have optional parameters.

### Return type

[**SecurityFiltersResponse**](SecurityFiltersResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListSecurityMonitoringRules

> SecurityMonitoringListRulesResponse ListSecurityMonitoringRules(ctx, datadog.ListSecurityMonitoringRulesOptionalParameters{})

List rules.

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

    pageSize := int64(10) // int64 | Size for a given page. (optional) (default to 10)
    pageNumber := int64(0) // int64 | Specific page number to return. (optional) (default to 0)
    optionalParams := datadog.ListSecurityMonitoringRulesOptionalParameters{
        PageSize: &pageSize,
        PageNumber: &pageNumber,
    }

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudSIEMApi.ListSecurityMonitoringRules(ctx, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudSIEMApi.ListSecurityMonitoringRules`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSecurityMonitoringRules`: SecurityMonitoringListRulesResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from CloudSIEMApi.ListSecurityMonitoringRules:\n%s\n", responseContent)
}
```

### Required Parameters

### Optional Parameters

Other parameters are passed through a pointer to a ListSecurityMonitoringRulesOptionalParameters struct.

| Name           | Type      | Description                     | Notes           |
| -------------- | --------- | ------------------------------- | --------------- |
| **pageSize**   | **int64** | Size for a given page.          | [default to 10] |
| **pageNumber** | **int64** | Specific page number to return. | [default to 0]  |

### Return type

[**SecurityMonitoringListRulesResponse**](SecurityMonitoringListRulesResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListSecurityMonitoringSignals

> SecurityMonitoringSignalsListResponse ListSecurityMonitoringSignals(ctx, datadog.ListSecurityMonitoringSignalsOptionalParameters{})

The list endpoint returns security signals that match a search query.
Both this endpoint and the POST endpoint can be used interchangeably when listing
security signals.

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
    resp, r, err := apiClient.CloudSIEMApi.ListSecurityMonitoringSignals(ctx, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudSIEMApi.ListSecurityMonitoringSignals`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSecurityMonitoringSignals`: SecurityMonitoringSignalsListResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from CloudSIEMApi.ListSecurityMonitoringSignals:\n%s\n", responseContent)
}
```

### Required Parameters

### Optional Parameters

Other parameters are passed through a pointer to a ListSecurityMonitoringSignalsOptionalParameters struct.

| Name            | Type                                                                  | Description                                                        | Notes           |
| --------------- | --------------------------------------------------------------------- | ------------------------------------------------------------------ | --------------- |
| **filterQuery** | **string**                                                            | The search query for security signals.                             |
| **filterFrom**  | **time.Time**                                                         | The minimum timestamp for requested security signals.              |
| **filterTo**    | **time.Time**                                                         | The maximum timestamp for requested security signals.              |
| **sort**        | [**SecurityMonitoringSignalsSort**](SecurityMonitoringSignalsSort.md) | The order of the security signals in results.                      |
| **pageCursor**  | **string**                                                            | A list of results using the cursor provided in the previous query. |
| **pageLimit**   | **int32**                                                             | The maximum number of security signals in the response.            | [default to 10] |

### Return type

[**SecurityMonitoringSignalsListResponse**](SecurityMonitoringSignalsListResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SearchSecurityMonitoringSignals

> SecurityMonitoringSignalsListResponse SearchSecurityMonitoringSignals(ctx, datadog.SearchSecurityMonitoringSignalsOptionalParameters{})

Returns security signals that match a search query.
Both this endpoint and the GET endpoint can be used interchangeably for listing
security signals.

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
    resp, r, err := apiClient.CloudSIEMApi.SearchSecurityMonitoringSignals(ctx, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudSIEMApi.SearchSecurityMonitoringSignals`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchSecurityMonitoringSignals`: SecurityMonitoringSignalsListResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from CloudSIEMApi.SearchSecurityMonitoringSignals:\n%s\n", responseContent)
}
```

### Required Parameters

### Optional Parameters

Other parameters are passed through a pointer to a SearchSecurityMonitoringSignalsOptionalParameters struct.

| Name     | Type                                                                              | Description | Notes |
| -------- | --------------------------------------------------------------------------------- | ----------- | ----- |
| **body** | [**SecurityMonitoringSignalListRequest**](SecurityMonitoringSignalListRequest.md) |             |

### Return type

[**SecurityMonitoringSignalsListResponse**](SecurityMonitoringSignalsListResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateSecurityFilter

> SecurityFilterResponse UpdateSecurityFilter(ctx, securityFilterId, body)

Update a specific security filter.
Returns the security filter object when the request is successful.

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

    securityFilterId := "securityFilterId_example" // string | The ID of the security filter.
    body := *datadog.NewSecurityFilterUpdateRequest(*datadog.NewSecurityFilterUpdateData(*datadog.NewSecurityFilterUpdateAttributes(), datadog.SecurityFilterType("security_filters"))) // SecurityFilterUpdateRequest | New definition of the security filter.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudSIEMApi.UpdateSecurityFilter(ctx, securityFilterId, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudSIEMApi.UpdateSecurityFilter`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSecurityFilter`: SecurityFilterResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from CloudSIEMApi.UpdateSecurityFilter:\n%s\n", responseContent)
}
```

### Required Parameters

| Name                 | Type                                                              | Description                                                                 | Notes |
| -------------------- | ----------------------------------------------------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**              | **context.Context**                                               | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **securityFilterId** | **string**                                                        | The ID of the security filter.                                              |       |
| **body**             | [**SecurityFilterUpdateRequest**](SecurityFilterUpdateRequest.md) | New definition of the security filter.                                      |

### Optional Parameters

This endpoint does not have optional parameters.

### Return type

[**SecurityFilterResponse**](SecurityFilterResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateSecurityMonitoringRule

> SecurityMonitoringRuleResponse UpdateSecurityMonitoringRule(ctx, ruleId, body)

Update an existing rule. When updating `cases`, `queries` or `options`, the whole field
must be included. For example, when modifying a query all queries must be included.
Default rules can only be updated to be enabled and to change notifications.

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
    resp, r, err := apiClient.CloudSIEMApi.UpdateSecurityMonitoringRule(ctx, ruleId, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudSIEMApi.UpdateSecurityMonitoringRule`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSecurityMonitoringRule`: SecurityMonitoringRuleResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from CloudSIEMApi.UpdateSecurityMonitoringRule:\n%s\n", responseContent)
}
```

### Required Parameters

| Name       | Type                                                                              | Description                                                                 | Notes |
| ---------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**    | **context.Context**                                                               | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **ruleId** | **string**                                                                        | The ID of the rule.                                                         |       |
| **body**   | [**SecurityMonitoringRuleUpdatePayload**](SecurityMonitoringRuleUpdatePayload.md) |                                                                             |

### Optional Parameters

This endpoint does not have optional parameters.

### Return type

[**SecurityMonitoringRuleResponse**](SecurityMonitoringRuleResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
