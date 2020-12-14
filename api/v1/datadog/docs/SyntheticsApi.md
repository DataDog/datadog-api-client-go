# \SyntheticsApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGlobalVariable**](SyntheticsApi.md#CreateGlobalVariable) | **Post** /api/v1/synthetics/variables | Create a global variable
[**CreatePrivateLocation**](SyntheticsApi.md#CreatePrivateLocation) | **Post** /api/v1/synthetics/private-locations | Create a private location
[**CreateTest**](SyntheticsApi.md#CreateTest) | **Post** /api/v1/synthetics/tests | Create a test
[**DeleteGlobalVariable**](SyntheticsApi.md#DeleteGlobalVariable) | **Delete** /api/v1/synthetics/variables/{variable_id} | Delete a global variable
[**DeletePrivateLocation**](SyntheticsApi.md#DeletePrivateLocation) | **Delete** /api/v1/synthetics/private-locations/{location_id} | Delete a private location
[**DeleteTests**](SyntheticsApi.md#DeleteTests) | **Post** /api/v1/synthetics/tests/delete | Delete tests
[**EditGlobalVariable**](SyntheticsApi.md#EditGlobalVariable) | **Put** /api/v1/synthetics/variables/{variable_id} | Edit a global variable
[**GetAPITestLatestResults**](SyntheticsApi.md#GetAPITestLatestResults) | **Get** /api/v1/synthetics/tests/{public_id}/results | Get the test&#39;s latest results summaries (API)
[**GetAPITestResult**](SyntheticsApi.md#GetAPITestResult) | **Get** /api/v1/synthetics/tests/{public_id}/results/{result_id} | Get a test result (API)
[**GetBrowserTest**](SyntheticsApi.md#GetBrowserTest) | **Get** /api/v1/synthetics/tests/browser/{public_id} | Get a test configuration (browser)
[**GetBrowserTestLatestResults**](SyntheticsApi.md#GetBrowserTestLatestResults) | **Get** /api/v1/synthetics/tests/browser/{public_id}/results | Get the test&#39;s latest results summaries (browser)
[**GetBrowserTestResult**](SyntheticsApi.md#GetBrowserTestResult) | **Get** /api/v1/synthetics/tests/browser/{public_id}/results/{result_id} | Get a test result (browser)
[**GetGlobalVariable**](SyntheticsApi.md#GetGlobalVariable) | **Get** /api/v1/synthetics/variables/{variable_id} | Get a global variable
[**GetPrivateLocation**](SyntheticsApi.md#GetPrivateLocation) | **Get** /api/v1/synthetics/private-locations/{location_id} | Get a private location
[**GetTest**](SyntheticsApi.md#GetTest) | **Get** /api/v1/synthetics/tests/{public_id} | Get a test configuration (API)
[**ListLocations**](SyntheticsApi.md#ListLocations) | **Get** /api/v1/synthetics/locations | Get all locations (public and private)
[**ListTests**](SyntheticsApi.md#ListTests) | **Get** /api/v1/synthetics/tests | Get the list of all tests
[**TriggerCITests**](SyntheticsApi.md#TriggerCITests) | **Post** /api/v1/synthetics/tests/trigger/ci | Trigger some Synthetics tests for CI
[**UpdatePrivateLocation**](SyntheticsApi.md#UpdatePrivateLocation) | **Put** /api/v1/synthetics/private-locations/{location_id} | Edit a private location
[**UpdateTest**](SyntheticsApi.md#UpdateTest) | **Put** /api/v1/synthetics/tests/{public_id} | Edit a test
[**UpdateTestPauseStatus**](SyntheticsApi.md#UpdateTestPauseStatus) | **Put** /api/v1/synthetics/tests/{public_id}/status | Pause or start a test



## CreateGlobalVariable

> SyntheticsGlobalVariable CreateGlobalVariable(ctx).Body(body).Execute()

Create a global variable



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

    body := *datadog.NewSyntheticsGlobalVariable("Example description", "MY_VARIABLE", []string{"Tags_example"}, *datadog.NewSyntheticsGlobalVariableValue("example-value")) // SyntheticsGlobalVariable | Details of the global variable to create.

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.SyntheticsApi.CreateGlobalVariable(ctx).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.CreateGlobalVariable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateGlobalVariable`: SyntheticsGlobalVariable
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from SyntheticsApi.CreateGlobalVariable:\n%v\n", response_content)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGlobalVariableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SyntheticsGlobalVariable**](SyntheticsGlobalVariable.md) | Details of the global variable to create. | 

### Return type

[**SyntheticsGlobalVariable**](SyntheticsGlobalVariable.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePrivateLocation

> SyntheticsPrivateLocationCreationResponse CreatePrivateLocation(ctx).Body(body).Execute()

Create a private location



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

    body := *datadog.NewSyntheticsPrivateLocation("Description of private location", "New private location", []string{"team:front"}) // SyntheticsPrivateLocation | Details of the private location to create.

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.SyntheticsApi.CreatePrivateLocation(ctx).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.CreatePrivateLocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePrivateLocation`: SyntheticsPrivateLocationCreationResponse
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from SyntheticsApi.CreatePrivateLocation:\n%v\n", response_content)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePrivateLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SyntheticsPrivateLocation**](SyntheticsPrivateLocation.md) | Details of the private location to create. | 

### Return type

[**SyntheticsPrivateLocationCreationResponse**](SyntheticsPrivateLocationCreationResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTest

> SyntheticsTestDetails CreateTest(ctx).Body(body).Execute()

Create a test



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

    body := *datadog.NewSyntheticsTestDetails() // SyntheticsTestDetails | Details of the test to create.

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.SyntheticsApi.CreateTest(ctx).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.CreateTest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTest`: SyntheticsTestDetails
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from SyntheticsApi.CreateTest:\n%v\n", response_content)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SyntheticsTestDetails**](SyntheticsTestDetails.md) | Details of the test to create. | 

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


## DeleteGlobalVariable

> DeleteGlobalVariable(ctx, variableId).Execute()

Delete a global variable



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

    variableId := "variableId_example" // string | The ID of the global variable.

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    r, err := api_client.SyntheticsApi.DeleteGlobalVariable(ctx, variableId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.DeleteGlobalVariable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**variableId** | **string** | The ID of the global variable. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGlobalVariableRequest struct via the builder pattern


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


## DeletePrivateLocation

> DeletePrivateLocation(ctx, locationId).Execute()

Delete a private location



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

    locationId := "locationId_example" // string | The ID of the private location.

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    r, err := api_client.SyntheticsApi.DeletePrivateLocation(ctx, locationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.DeletePrivateLocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locationId** | **string** | The ID of the private location. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePrivateLocationRequest struct via the builder pattern


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


## DeleteTests

> SyntheticsDeleteTestsResponse DeleteTests(ctx).Body(body).Execute()

Delete tests



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

    body := *datadog.NewSyntheticsDeleteTestsPayload() // SyntheticsDeleteTestsPayload | Public ID list of the Synthetic tests to be deleted.

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.SyntheticsApi.DeleteTests(ctx).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.DeleteTests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteTests`: SyntheticsDeleteTestsResponse
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from SyntheticsApi.DeleteTests:\n%v\n", response_content)
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


## EditGlobalVariable

> SyntheticsGlobalVariable EditGlobalVariable(ctx, variableId).Body(body).Execute()

Edit a global variable



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

    variableId := "variableId_example" // string | The ID of the global variable.
    body := *datadog.NewSyntheticsGlobalVariable("Example description", "MY_VARIABLE", []string{"Tags_example"}, *datadog.NewSyntheticsGlobalVariableValue("example-value")) // SyntheticsGlobalVariable | Details of the global variable to update.

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.SyntheticsApi.EditGlobalVariable(ctx, variableId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.EditGlobalVariable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditGlobalVariable`: SyntheticsGlobalVariable
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from SyntheticsApi.EditGlobalVariable:\n%v\n", response_content)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**variableId** | **string** | The ID of the global variable. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditGlobalVariableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**SyntheticsGlobalVariable**](SyntheticsGlobalVariable.md) | Details of the global variable to update. | 

### Return type

[**SyntheticsGlobalVariable**](SyntheticsGlobalVariable.md)

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

    publicId := "publicId_example" // string | The public ID of the test for which to search results for.
    fromTs := int64(789) // int64 | Timestamp from which to start querying results. (optional)
    toTs := int64(789) // int64 | Timestamp up to which to query results. (optional)
    probeDc := []string{"Inner_example"} // []string | Locations for which to query results. (optional)

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.SyntheticsApi.GetAPITestLatestResults(ctx, publicId).FromTs(fromTs).ToTs(toTs).ProbeDc(probeDc).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.GetAPITestLatestResults``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAPITestLatestResults`: SyntheticsGetAPITestLatestResultsResponse
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from SyntheticsApi.GetAPITestLatestResults:\n%v\n", response_content)
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
 **probeDc** | **[]string** | Locations for which to query results. | 

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
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from SyntheticsApi.GetAPITestResult:\n%v\n", response_content)
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


## GetBrowserTest

> SyntheticsTestDetails GetBrowserTest(ctx, publicId).Execute()

Get a test configuration (browser)



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

    publicId := "publicId_example" // string | The public ID of the test to get details from.

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.SyntheticsApi.GetBrowserTest(ctx, publicId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.GetBrowserTest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBrowserTest`: SyntheticsTestDetails
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from SyntheticsApi.GetBrowserTest:\n%v\n", response_content)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicId** | **string** | The public ID of the test to get details from. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBrowserTestRequest struct via the builder pattern


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


## GetBrowserTestLatestResults

> SyntheticsGetBrowserTestLatestResultsResponse GetBrowserTestLatestResults(ctx, publicId).FromTs(fromTs).ToTs(toTs).ProbeDc(probeDc).Execute()

Get the test's latest results summaries (browser)



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

    publicId := "publicId_example" // string | The public ID of the browser test for which to search results for.
    fromTs := int64(789) // int64 | Timestamp from which to start querying results. (optional)
    toTs := int64(789) // int64 | Timestamp up to which to query results. (optional)
    probeDc := []string{"Inner_example"} // []string | Locations for which to query results. (optional)

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.SyntheticsApi.GetBrowserTestLatestResults(ctx, publicId).FromTs(fromTs).ToTs(toTs).ProbeDc(probeDc).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.GetBrowserTestLatestResults``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBrowserTestLatestResults`: SyntheticsGetBrowserTestLatestResultsResponse
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from SyntheticsApi.GetBrowserTestLatestResults:\n%v\n", response_content)
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
 **probeDc** | **[]string** | Locations for which to query results. | 

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
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from SyntheticsApi.GetBrowserTestResult:\n%v\n", response_content)
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


## GetGlobalVariable

> SyntheticsGlobalVariable GetGlobalVariable(ctx, variableId).Execute()

Get a global variable



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

    variableId := "variableId_example" // string | The ID of the global variable.

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.SyntheticsApi.GetGlobalVariable(ctx, variableId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.GetGlobalVariable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGlobalVariable`: SyntheticsGlobalVariable
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from SyntheticsApi.GetGlobalVariable:\n%v\n", response_content)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**variableId** | **string** | The ID of the global variable. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGlobalVariableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SyntheticsGlobalVariable**](SyntheticsGlobalVariable.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrivateLocation

> SyntheticsPrivateLocation GetPrivateLocation(ctx, locationId).Execute()

Get a private location



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

    locationId := "locationId_example" // string | The ID of the private location.

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.SyntheticsApi.GetPrivateLocation(ctx, locationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.GetPrivateLocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPrivateLocation`: SyntheticsPrivateLocation
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from SyntheticsApi.GetPrivateLocation:\n%v\n", response_content)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locationId** | **string** | The ID of the private location. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPrivateLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SyntheticsPrivateLocation**](SyntheticsPrivateLocation.md)

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

Get a test configuration (API)



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

    publicId := "publicId_example" // string | The public ID of the test to get details from.

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.SyntheticsApi.GetTest(ctx, publicId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.GetTest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTest`: SyntheticsTestDetails
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from SyntheticsApi.GetTest:\n%v\n", response_content)
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


## ListLocations

> SyntheticsLocations ListLocations(ctx).Execute()

Get all locations (public and private)



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


    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.SyntheticsApi.ListLocations(ctx).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.ListLocations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLocations`: SyntheticsLocations
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from SyntheticsApi.ListLocations:\n%v\n", response_content)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListLocationsRequest struct via the builder pattern


### Return type

[**SyntheticsLocations**](SyntheticsLocations.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTests

> SyntheticsListTestsResponse ListTests(ctx).Execute()

Get the list of all tests



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


    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.SyntheticsApi.ListTests(ctx).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.ListTests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTests`: SyntheticsListTestsResponse
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from SyntheticsApi.ListTests:\n%v\n", response_content)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListTestsRequest struct via the builder pattern


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


## TriggerCITests

> SyntheticsTriggerCITestsResponse TriggerCITests(ctx).Body(body).Execute()

Trigger some Synthetics tests for CI



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

    body := *datadog.NewSyntheticsCITestBody() // SyntheticsCITestBody | Details of the test to trigger.

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.SyntheticsApi.TriggerCITests(ctx).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.TriggerCITests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TriggerCITests`: SyntheticsTriggerCITestsResponse
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from SyntheticsApi.TriggerCITests:\n%v\n", response_content)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTriggerCITestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SyntheticsCITestBody**](SyntheticsCITestBody.md) | Details of the test to trigger. | 

### Return type

[**SyntheticsTriggerCITestsResponse**](SyntheticsTriggerCITestsResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePrivateLocation

> SyntheticsPrivateLocation UpdatePrivateLocation(ctx, locationId).Body(body).Execute()

Edit a private location



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

    locationId := "locationId_example" // string | The ID of the private location.
    body := *datadog.NewSyntheticsPrivateLocation("Description of private location", "New private location", []string{"team:front"}) // SyntheticsPrivateLocation | Details of the private location to be updated.

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.SyntheticsApi.UpdatePrivateLocation(ctx, locationId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.UpdatePrivateLocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePrivateLocation`: SyntheticsPrivateLocation
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from SyntheticsApi.UpdatePrivateLocation:\n%v\n", response_content)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locationId** | **string** | The ID of the private location. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePrivateLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**SyntheticsPrivateLocation**](SyntheticsPrivateLocation.md) | Details of the private location to be updated. | 

### Return type

[**SyntheticsPrivateLocation**](SyntheticsPrivateLocation.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
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

    publicId := "publicId_example" // string | The public ID of the test to get details from.
    body := *datadog.NewSyntheticsTestDetails() // SyntheticsTestDetails | New test details to be saved.

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.SyntheticsApi.UpdateTest(ctx, publicId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.UpdateTest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTest`: SyntheticsTestDetails
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from SyntheticsApi.UpdateTest:\n%v\n", response_content)
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

    publicId := "publicId_example" // string | The public ID of the Synthetic test to update.
    body := *datadog.NewSyntheticsUpdateTestPauseStatusPayload() // SyntheticsUpdateTestPauseStatusPayload | Status to set the given Synthetic test to.

    configuration := datadog.NewConfiguration()

    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.SyntheticsApi.UpdateTestPauseStatus(ctx, publicId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.UpdateTestPauseStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTestPauseStatus`: bool
    response_content, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from SyntheticsApi.UpdateTestPauseStatus:\n%v\n", response_content)
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

