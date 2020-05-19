# \SyntheticsApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTest**](SyntheticsApi.md#CreateTest) | **Post** /api/v1/synthetics/tests | Create a test
[**DeleteTests**](SyntheticsApi.md#DeleteTests) | **Post** /api/v1/synthetics/tests/delete | Delete tests
[**GetAPITestLatestResults**](SyntheticsApi.md#GetAPITestLatestResults) | **Get** /api/v1/synthetics/tests/{public_id}/results | Get the test&#39;s latest results summaries (API)
[**GetAPITestResult**](SyntheticsApi.md#GetAPITestResult) | **Get** /api/v1/synthetics/tests/{public_id}/results/{result_id} | Get a test result (API)
[**GetBrowserTestLatestResults**](SyntheticsApi.md#GetBrowserTestLatestResults) | **Get** /api/v1/synthetics/tests/browser/{public_id}/results | Get the test&#39;s latest results summaries (browser)
[**GetBrowserTestResult**](SyntheticsApi.md#GetBrowserTestResult) | **Get** /api/v1/synthetics/tests/browser/{public_id}/results/{result_id} | Get a test result (browser)
[**GetTest**](SyntheticsApi.md#GetTest) | **Get** /api/v1/synthetics/tests/{public_id} | Get a test configuration
[**ListTests**](SyntheticsApi.md#ListTests) | **Get** /api/v1/synthetics/tests | Get a list of tests
[**UpdateTest**](SyntheticsApi.md#UpdateTest) | **Put** /api/v1/synthetics/tests/{public_id} | Edit a test
[**UpdateTestPauseStatus**](SyntheticsApi.md#UpdateTestPauseStatus) | **Put** /api/v1/synthetics/tests/{public_id}/status | Pause or start a test



## CreateTest

> SyntheticsTestDetails CreateTest(ctx).Body(body).FromTestId(fromTestId).Execute()

Create a test



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

    body := datadog.SyntheticsTestDetails{Config: datadog.SyntheticsTestConfig{Assertions: []SyntheticsAssertion{datadog.SyntheticsAssertion{Operator: datadog.SyntheticsAssertionOperator{}, Property: "Property_example", Target: 123, Type: datadog.SyntheticsAssertionType{}}), Request: datadog.SyntheticsTestRequest{BasicAuth: datadog.SyntheticsTestRequest_basicAuth{Password: "Password_example", Username: "Username_example"}, Body: "Body_example", Headers: map[string]string{ "Key" = "Value" }, Host: "Host_example", Method: datadog.HTTPMethod{}, Port: int64(123), Query: 123, Timeout: 123, Url: "Url_example"}, Variables: []SyntheticsBrowserVariable{datadog.SyntheticsBrowserVariable{Example: "Example_example", Id: "Id_example", Name: "Name_example", Pattern: "Pattern_example", Type: datadog.SyntheticsBrowserVariableType{}})}, CreatedAt: "CreatedAt_example", CreatedBy: datadog.SyntheticsTestAuthor{Email: "Email_example", Handle: "Handle_example", Id: int64(123), Name: "Name_example"}, Locations: []string{"Locations_example"), Message: "Message_example", ModifiedAt: "ModifiedAt_example", ModifiedBy: datadog.SyntheticsTestAuthor{Email: "Email_example", Handle: "Handle_example", Id: int64(123), Name: "Name_example"}, MonitorId: int64(123), Name: "Name_example", Options: datadog.SyntheticsTestOptions{AcceptSelfSigned: false, AllowInsecure: false, DeviceIds: []SyntheticsDeviceID{datadog.SyntheticsDeviceID{}), FollowRedirects: false, MinFailureDuration: int64(123), MinLocationFailed: int64(123), Retry: datadog.SyntheticsTestOptions_retry{Count: int64(123), Interval: 123}, TickEvery: datadog.SyntheticsTickInterval{}}, PublicId: "PublicId_example", Status: datadog.SyntheticsTestPauseStatus{}, Subtype: datadog.SyntheticsTestDetailsSubType{}, Tags: []string{"Tags_example"), Type: datadog.SyntheticsTestDetailsType{}} // SyntheticsTestDetails | Details of the test to create.
    fromTestId := "fromTestId_example" // string | Public ID of the test to clone, undefined if the test is newly created. (optional)

    configuration := datadog.NewConfiguration()
    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.SyntheticsApi.CreateTest(ctx, body).FromTestId(fromTestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.CreateTest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTest`: SyntheticsTestDetails
    fmt.Fprintf(os.Stdout, "Response from `SyntheticsApi.CreateTest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SyntheticsTestDetails**](SyntheticsTestDetails.md) | Details of the test to create. | 
 **fromTestId** | **string** | Public ID of the test to clone, undefined if the test is newly created. | 

### Return type

[**SyntheticsTestDetails**](SyntheticsTestDetails.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTests

> SyntheticsDeleteTestsResponse DeleteTests(ctx).Body(body).Execute()

Delete tests



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

    body := datadog.SyntheticsDeleteTestsPayload{PublicIds: []string{"PublicIds_example")} // SyntheticsDeleteTestsPayload | Public ID list of the Synthetic tests to be deleted.

    configuration := datadog.NewConfiguration()
    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.SyntheticsApi.DeleteTests(ctx, body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.DeleteTests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteTests`: SyntheticsDeleteTestsResponse
    fmt.Fprintf(os.Stdout, "Response from `SyntheticsApi.DeleteTests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SyntheticsDeleteTestsPayload**](SyntheticsDeleteTestsPayload.md) | Public ID list of the Synthetic tests to be deleted. | 

### Return type

[**SyntheticsDeleteTestsResponse**](SyntheticsDeleteTestsResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAPITestLatestResults

> SyntheticsGetAPITestLatestResultsResponse GetAPITestLatestResults(ctx, publicId).FromTs(fromTs).ToTs(toTs).ProbeDc(probeDc).Execute()

Get the test's latest results summaries (API)



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

    publicId := "publicId_example" // string | The public ID of the test for which to search results for.
    fromTs := 987 // int64 | Timestamp from which to start querying results. (optional)
    toTs := 987 // int64 | Timestamp up to which to query results. (optional)
    probeDc := []string{"Inner_example"} // []string | Locations for which to query results. (optional)

    configuration := datadog.NewConfiguration()
    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.SyntheticsApi.GetAPITestLatestResults(ctx, publicId).FromTs(fromTs).ToTs(toTs).ProbeDc(probeDc).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.GetAPITestLatestResults``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAPITestLatestResults`: SyntheticsGetAPITestLatestResultsResponse
    fmt.Fprintf(os.Stdout, "Response from `SyntheticsApi.GetAPITestLatestResults`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicId** | **string** | The public ID of the test for which to search results for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAPITestLatestResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fromTs** | **int64** | Timestamp from which to start querying results. | 
 **toTs** | **int64** | Timestamp up to which to query results. | 
 **probeDc** | [**[]string**](string.md) | Locations for which to query results. | 

### Return type

[**SyntheticsGetAPITestLatestResultsResponse**](SyntheticsGetAPITestLatestResultsResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAPITestResult

> SyntheticsAPITestResultFull GetAPITestResult(ctx, publicId, resultId).Execute()

Get a test result (API)



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

    publicId := "publicId_example" // string | The public ID of the API test to which the target result belongs.
    resultId := "resultId_example" // string | The ID of the result to get.

    configuration := datadog.NewConfiguration()
    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.SyntheticsApi.GetAPITestResult(ctx, publicId, resultId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.GetAPITestResult``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAPITestResult`: SyntheticsAPITestResultFull
    fmt.Fprintf(os.Stdout, "Response from `SyntheticsApi.GetAPITestResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicId** | **string** | The public ID of the API test to which the target result belongs. | 
**resultId** | **string** | The ID of the result to get. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAPITestResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SyntheticsAPITestResultFull**](SyntheticsAPITestResultFull.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBrowserTestLatestResults

> SyntheticsGetBrowserTestLatestResultsResponse GetBrowserTestLatestResults(ctx, publicId).FromTs(fromTs).ToTs(toTs).ProbeDc(probeDc).Execute()

Get the test's latest results summaries (browser)



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

    publicId := "publicId_example" // string | The public ID of the browser test for which to search results for.
    fromTs := 987 // int64 | Timestamp from which to start querying results. (optional)
    toTs := 987 // int64 | Timestamp up to which to query results. (optional)
    probeDc := []string{"Inner_example"} // []string | Locations for which to query results. (optional)

    configuration := datadog.NewConfiguration()
    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.SyntheticsApi.GetBrowserTestLatestResults(ctx, publicId).FromTs(fromTs).ToTs(toTs).ProbeDc(probeDc).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.GetBrowserTestLatestResults``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBrowserTestLatestResults`: SyntheticsGetBrowserTestLatestResultsResponse
    fmt.Fprintf(os.Stdout, "Response from `SyntheticsApi.GetBrowserTestLatestResults`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicId** | **string** | The public ID of the browser test for which to search results for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBrowserTestLatestResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fromTs** | **int64** | Timestamp from which to start querying results. | 
 **toTs** | **int64** | Timestamp up to which to query results. | 
 **probeDc** | [**[]string**](string.md) | Locations for which to query results. | 

### Return type

[**SyntheticsGetBrowserTestLatestResultsResponse**](SyntheticsGetBrowserTestLatestResultsResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBrowserTestResult

> SyntheticsBrowserTestResultFull GetBrowserTestResult(ctx, publicId, resultId).Execute()

Get a test result (browser)



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

    publicId := "publicId_example" // string | The public ID of the browser test to which the target result belongs.
    resultId := "resultId_example" // string | The ID of the result to get.

    configuration := datadog.NewConfiguration()
    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.SyntheticsApi.GetBrowserTestResult(ctx, publicId, resultId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.GetBrowserTestResult``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBrowserTestResult`: SyntheticsBrowserTestResultFull
    fmt.Fprintf(os.Stdout, "Response from `SyntheticsApi.GetBrowserTestResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicId** | **string** | The public ID of the browser test to which the target result belongs. | 
**resultId** | **string** | The ID of the result to get. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBrowserTestResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SyntheticsBrowserTestResultFull**](SyntheticsBrowserTestResultFull.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTest

> SyntheticsTestDetails GetTest(ctx, publicId).Execute()

Get a test configuration



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

    publicId := "publicId_example" // string | The public ID of the test to get details from.

    configuration := datadog.NewConfiguration()
    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.SyntheticsApi.GetTest(ctx, publicId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.GetTest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTest`: SyntheticsTestDetails
    fmt.Fprintf(os.Stdout, "Response from `SyntheticsApi.GetTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicId** | **string** | The public ID of the test to get details from. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SyntheticsTestDetails**](SyntheticsTestDetails.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTests

> SyntheticsListTestsResponse ListTests(ctx).CheckType(checkType).Execute()

Get a list of tests



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

    checkType := "checkType_example" // string | API or browser to filter the list by test type, undefined to get the unfiltered list. (optional)

    configuration := datadog.NewConfiguration()
    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.SyntheticsApi.ListTests(ctx).CheckType(checkType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.ListTests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTests`: SyntheticsListTestsResponse
    fmt.Fprintf(os.Stdout, "Response from `SyntheticsApi.ListTests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **checkType** | **string** | API or browser to filter the list by test type, undefined to get the unfiltered list. | 

### Return type

[**SyntheticsListTestsResponse**](SyntheticsListTestsResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTest

> SyntheticsTestDetails UpdateTest(ctx, publicId).Body(body).Execute()

Edit a test



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

    publicId := "publicId_example" // string | The public ID of the test to get details from.
    body := datadog.SyntheticsTestDetails{Config: datadog.SyntheticsTestConfig{Assertions: []SyntheticsAssertion{datadog.SyntheticsAssertion{Operator: datadog.SyntheticsAssertionOperator{}, Property: "Property_example", Target: 123, Type: datadog.SyntheticsAssertionType{}}), Request: datadog.SyntheticsTestRequest{BasicAuth: datadog.SyntheticsTestRequest_basicAuth{Password: "Password_example", Username: "Username_example"}, Body: "Body_example", Headers: map[string]string{ "Key" = "Value" }, Host: "Host_example", Method: datadog.HTTPMethod{}, Port: int64(123), Query: 123, Timeout: 123, Url: "Url_example"}, Variables: []SyntheticsBrowserVariable{datadog.SyntheticsBrowserVariable{Example: "Example_example", Id: "Id_example", Name: "Name_example", Pattern: "Pattern_example", Type: datadog.SyntheticsBrowserVariableType{}})}, CreatedAt: "CreatedAt_example", CreatedBy: , Locations: []string{"Locations_example"), Message: "Message_example", ModifiedAt: "ModifiedAt_example", ModifiedBy: , MonitorId: int64(123), Name: "Name_example", Options: datadog.SyntheticsTestOptions{AcceptSelfSigned: false, AllowInsecure: false, DeviceIds: []SyntheticsDeviceID{datadog.SyntheticsDeviceID{}), FollowRedirects: false, MinFailureDuration: int64(123), MinLocationFailed: int64(123), Retry: datadog.SyntheticsTestOptions_retry{Count: int64(123), Interval: 123}, TickEvery: datadog.SyntheticsTickInterval{}}, PublicId: "PublicId_example", Status: datadog.SyntheticsTestPauseStatus{}, Subtype: datadog.SyntheticsTestDetailsSubType{}, Tags: []string{"Tags_example"), Type: datadog.SyntheticsTestDetailsType{}} // SyntheticsTestDetails | New test details to be saved.

    configuration := datadog.NewConfiguration()
    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.SyntheticsApi.UpdateTest(ctx, publicId, body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.UpdateTest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTest`: SyntheticsTestDetails
    fmt.Fprintf(os.Stdout, "Response from `SyntheticsApi.UpdateTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicId** | **string** | The public ID of the test to get details from. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**SyntheticsTestDetails**](SyntheticsTestDetails.md) | New test details to be saved. | 

### Return type

[**SyntheticsTestDetails**](SyntheticsTestDetails.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTestPauseStatus

> bool UpdateTestPauseStatus(ctx, publicId).Body(body).Execute()

Pause or start a test



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

    publicId := "publicId_example" // string | The public ID of the Synthetic test to update.
    body := datadog.SyntheticsUpdateTestPauseStatusPayload{NewStatus: } // SyntheticsUpdateTestPauseStatusPayload | Status to set the given Synthetic test to.

    configuration := datadog.NewConfiguration()
    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.SyntheticsApi.UpdateTestPauseStatus(ctx, publicId, body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.UpdateTestPauseStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTestPauseStatus`: bool
    fmt.Fprintf(os.Stdout, "Response from `SyntheticsApi.UpdateTestPauseStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicId** | **string** | The public ID of the Synthetic test to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTestPauseStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**SyntheticsUpdateTestPauseStatusPayload**](SyntheticsUpdateTestPauseStatusPayload.md) | Status to set the given Synthetic test to. | 

### Return type

**bool**

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

