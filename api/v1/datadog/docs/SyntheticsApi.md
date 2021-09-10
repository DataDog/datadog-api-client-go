# SyntheticsApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------ | ------------ | ------------
[**CreateGlobalVariable**](SyntheticsApi.md#CreateGlobalVariable) | **Post** /api/v1/synthetics/variables | Create a global variable
[**CreatePrivateLocation**](SyntheticsApi.md#CreatePrivateLocation) | **Post** /api/v1/synthetics/private-locations | Create a private location
[**CreateSyntheticsAPITest**](SyntheticsApi.md#CreateSyntheticsAPITest) | **Post** /api/v1/synthetics/tests/api | Create an API test
[**CreateSyntheticsBrowserTest**](SyntheticsApi.md#CreateSyntheticsBrowserTest) | **Post** /api/v1/synthetics/tests/browser | Create a browser test
[**DeleteGlobalVariable**](SyntheticsApi.md#DeleteGlobalVariable) | **Delete** /api/v1/synthetics/variables/{variable_id} | Delete a global variable
[**DeletePrivateLocation**](SyntheticsApi.md#DeletePrivateLocation) | **Delete** /api/v1/synthetics/private-locations/{location_id} | Delete a private location
[**DeleteTests**](SyntheticsApi.md#DeleteTests) | **Post** /api/v1/synthetics/tests/delete | Delete tests
[**EditGlobalVariable**](SyntheticsApi.md#EditGlobalVariable) | **Put** /api/v1/synthetics/variables/{variable_id} | Edit a global variable
[**GetAPITest**](SyntheticsApi.md#GetAPITest) | **Get** /api/v1/synthetics/tests/api/{public_id} | Get an API test
[**GetAPITestLatestResults**](SyntheticsApi.md#GetAPITestLatestResults) | **Get** /api/v1/synthetics/tests/{public_id}/results | Get an API test&#39;s latest results summaries
[**GetAPITestResult**](SyntheticsApi.md#GetAPITestResult) | **Get** /api/v1/synthetics/tests/{public_id}/results/{result_id} | Get an API test result
[**GetBrowserTest**](SyntheticsApi.md#GetBrowserTest) | **Get** /api/v1/synthetics/tests/browser/{public_id} | Get a browser test
[**GetBrowserTestLatestResults**](SyntheticsApi.md#GetBrowserTestLatestResults) | **Get** /api/v1/synthetics/tests/browser/{public_id}/results | Get a browser test&#39;s latest results summaries
[**GetBrowserTestResult**](SyntheticsApi.md#GetBrowserTestResult) | **Get** /api/v1/synthetics/tests/browser/{public_id}/results/{result_id} | Get a browser test result
[**GetGlobalVariable**](SyntheticsApi.md#GetGlobalVariable) | **Get** /api/v1/synthetics/variables/{variable_id} | Get a global variable
[**GetPrivateLocation**](SyntheticsApi.md#GetPrivateLocation) | **Get** /api/v1/synthetics/private-locations/{location_id} | Get a private location
[**GetTest**](SyntheticsApi.md#GetTest) | **Get** /api/v1/synthetics/tests/{public_id} | Get a test configuration
[**ListGlobalVariables**](SyntheticsApi.md#ListGlobalVariables) | **Get** /api/v1/synthetics/variables | Get all global variables
[**ListLocations**](SyntheticsApi.md#ListLocations) | **Get** /api/v1/synthetics/locations | Get all locations (public and private)
[**ListTests**](SyntheticsApi.md#ListTests) | **Get** /api/v1/synthetics/tests | Get the list of all tests
[**TriggerCITests**](SyntheticsApi.md#TriggerCITests) | **Post** /api/v1/synthetics/tests/trigger/ci | Trigger tests from CI/CD pipelines
[**UpdateAPITest**](SyntheticsApi.md#UpdateAPITest) | **Put** /api/v1/synthetics/tests/api/{public_id} | Edit an API test
[**UpdateBrowserTest**](SyntheticsApi.md#UpdateBrowserTest) | **Put** /api/v1/synthetics/tests/browser/{public_id} | Edit a browser test
[**UpdatePrivateLocation**](SyntheticsApi.md#UpdatePrivateLocation) | **Put** /api/v1/synthetics/private-locations/{location_id} | Edit a private location
[**UpdateTestPauseStatus**](SyntheticsApi.md#UpdateTestPauseStatus) | **Put** /api/v1/synthetics/tests/{public_id}/status | Pause or start a test



## CreateGlobalVariable

> SyntheticsGlobalVariable CreateGlobalVariable(ctx, body)

Create a Synthetics global variable.

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

    body := *datadog.NewSyntheticsGlobalVariable("Example description", "MY_VARIABLE", []string{"Tags_example"}, *datadog.NewSyntheticsGlobalVariableValue()) // SyntheticsGlobalVariable | Details of the global variable to create.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.SyntheticsApi.CreateGlobalVariable(ctx, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.CreateGlobalVariable`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateGlobalVariable`: SyntheticsGlobalVariable
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from SyntheticsApi.CreateGlobalVariable:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**body** | [**SyntheticsGlobalVariable**](SyntheticsGlobalVariable.md) | Details of the global variable to create. | 


### Optional Parameters

This endpoint does not have optional parameters.


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

> SyntheticsPrivateLocationCreationResponse CreatePrivateLocation(ctx, body)

Create a new Synthetics private location.

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

    body := *datadog.NewSyntheticsPrivateLocation("Description of private location", "New private location", []string{"team:front"}) // SyntheticsPrivateLocation | Details of the private location to create.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.SyntheticsApi.CreatePrivateLocation(ctx, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.CreatePrivateLocation`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePrivateLocation`: SyntheticsPrivateLocationCreationResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from SyntheticsApi.CreatePrivateLocation:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**body** | [**SyntheticsPrivateLocation**](SyntheticsPrivateLocation.md) | Details of the private location to create. | 


### Optional Parameters

This endpoint does not have optional parameters.


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


## CreateSyntheticsAPITest

> SyntheticsAPITest CreateSyntheticsAPITest(ctx, body)

Create a Synthetic API test.

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

    body := *datadog.NewSyntheticsAPITest() // SyntheticsAPITest | Details of the test to create.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.SyntheticsApi.CreateSyntheticsAPITest(ctx, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.CreateSyntheticsAPITest`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSyntheticsAPITest`: SyntheticsAPITest
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from SyntheticsApi.CreateSyntheticsAPITest:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**body** | [**SyntheticsAPITest**](SyntheticsAPITest.md) | Details of the test to create. | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**SyntheticsAPITest**](SyntheticsAPITest.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSyntheticsBrowserTest

> SyntheticsBrowserTest CreateSyntheticsBrowserTest(ctx, body)

Create a Synthetic browser test.

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

    body := *datadog.NewSyntheticsBrowserTest("Message_example") // SyntheticsBrowserTest | Details of the test to create.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.SyntheticsApi.CreateSyntheticsBrowserTest(ctx, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.CreateSyntheticsBrowserTest`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSyntheticsBrowserTest`: SyntheticsBrowserTest
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from SyntheticsApi.CreateSyntheticsBrowserTest:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**body** | [**SyntheticsBrowserTest**](SyntheticsBrowserTest.md) | Details of the test to create. | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**SyntheticsBrowserTest**](SyntheticsBrowserTest.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGlobalVariable

> DeleteGlobalVariable(ctx, variableId)

Delete a Synthetics global variable.

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

    variableId := "variableId_example" // string | The ID of the global variable.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    r, err := apiClient.SyntheticsApi.DeleteGlobalVariable(ctx, variableId)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.DeleteGlobalVariable`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**variableId** | **string** | The ID of the global variable. | 


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


## DeletePrivateLocation

> DeletePrivateLocation(ctx, locationId)

Delete a Synthetics private location.

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

    locationId := "locationId_example" // string | The ID of the private location.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    r, err := apiClient.SyntheticsApi.DeletePrivateLocation(ctx, locationId)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.DeletePrivateLocation`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**locationId** | **string** | The ID of the private location. | 


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


## DeleteTests

> SyntheticsDeleteTestsResponse DeleteTests(ctx, body)

Delete multiple Synthetic tests by ID.

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

    body := *datadog.NewSyntheticsDeleteTestsPayload() // SyntheticsDeleteTestsPayload | Public ID list of the Synthetic tests to be deleted.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.SyntheticsApi.DeleteTests(ctx, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.DeleteTests`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteTests`: SyntheticsDeleteTestsResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from SyntheticsApi.DeleteTests:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**body** | [**SyntheticsDeleteTestsPayload**](SyntheticsDeleteTestsPayload.md) | Public ID list of the Synthetic tests to be deleted. | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**SyntheticsDeleteTestsResponse**](SyntheticsDeleteTestsResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditGlobalVariable

> SyntheticsGlobalVariable EditGlobalVariable(ctx, variableId, body)

Edit a Synthetics global variable.

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

    variableId := "variableId_example" // string | The ID of the global variable.
    body := *datadog.NewSyntheticsGlobalVariable("Example description", "MY_VARIABLE", []string{"Tags_example"}, *datadog.NewSyntheticsGlobalVariableValue()) // SyntheticsGlobalVariable | Details of the global variable to update.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.SyntheticsApi.EditGlobalVariable(ctx, variableId, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.EditGlobalVariable`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditGlobalVariable`: SyntheticsGlobalVariable
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from SyntheticsApi.EditGlobalVariable:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**variableId** | **string** | The ID of the global variable. |  |
**body** | [**SyntheticsGlobalVariable**](SyntheticsGlobalVariable.md) | Details of the global variable to update. | 


### Optional Parameters

This endpoint does not have optional parameters.


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


## GetAPITest

> SyntheticsAPITest GetAPITest(ctx, publicId)

Get the detailed configuration associated with
a Synthetic API test.

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

    publicId := "publicId_example" // string | The public ID of the test to get details from.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.SyntheticsApi.GetAPITest(ctx, publicId)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.GetAPITest`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAPITest`: SyntheticsAPITest
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from SyntheticsApi.GetAPITest:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**publicId** | **string** | The public ID of the test to get details from. | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**SyntheticsAPITest**](SyntheticsAPITest.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAPITestLatestResults

> SyntheticsGetAPITestLatestResultsResponse GetAPITestLatestResults(ctx, publicId, datadog.GetAPITestLatestResultsOptionalParameters{})

Get the last 50 test results summaries for a given Synthetics API test.

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

    publicId := "publicId_example" // string | The public ID of the test for which to search results for.
    fromTs := int64(789) // int64 | Timestamp in milliseconds from which to start querying results. (optional)
    toTs := int64(789) // int64 | Timestamp in milliseconds up to which to query results. (optional)
    probeDc := []string{"Inner_example"} // []string | Locations for which to query results. (optional)
    optionalParams := datadog.GetAPITestLatestResultsOptionalParameters{
        FromTs: &fromTs,
        ToTs: &toTs,
        ProbeDc: &probeDc,
    }

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.SyntheticsApi.GetAPITestLatestResults(ctx, publicId, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.GetAPITestLatestResults`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAPITestLatestResults`: SyntheticsGetAPITestLatestResultsResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from SyntheticsApi.GetAPITestLatestResults:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**publicId** | **string** | The public ID of the test for which to search results for. | 


### Optional Parameters


Other parameters are passed through a pointer to a GetAPITestLatestResultsOptionalParameters struct.


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**fromTs** | **int64** | Timestamp in milliseconds from which to start querying results. | 
**toTs** | **int64** | Timestamp in milliseconds up to which to query results. | 
**probeDc** | **[]string** | Locations for which to query results. | 

### Return type

[**SyntheticsGetAPITestLatestResultsResponse**](SyntheticsGetAPITestLatestResultsResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAPITestResult

> SyntheticsAPITestResultFull GetAPITestResult(ctx, publicId, resultId)

Get a specific full result from a given (API) Synthetic test.

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

    publicId := "publicId_example" // string | The public ID of the API test to which the target result belongs.
    resultId := "resultId_example" // string | The ID of the result to get.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.SyntheticsApi.GetAPITestResult(ctx, publicId, resultId)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.GetAPITestResult`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAPITestResult`: SyntheticsAPITestResultFull
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from SyntheticsApi.GetAPITestResult:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**publicId** | **string** | The public ID of the API test to which the target result belongs. |  |
**resultId** | **string** | The ID of the result to get. | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**SyntheticsAPITestResultFull**](SyntheticsAPITestResultFull.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBrowserTest

> SyntheticsBrowserTest GetBrowserTest(ctx, publicId)

Get the detailed configuration (including steps) associated with
a Synthetic browser test.

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

    publicId := "publicId_example" // string | The public ID of the test to get details from.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.SyntheticsApi.GetBrowserTest(ctx, publicId)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.GetBrowserTest`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBrowserTest`: SyntheticsBrowserTest
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from SyntheticsApi.GetBrowserTest:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**publicId** | **string** | The public ID of the test to get details from. | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**SyntheticsBrowserTest**](SyntheticsBrowserTest.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBrowserTestLatestResults

> SyntheticsGetBrowserTestLatestResultsResponse GetBrowserTestLatestResults(ctx, publicId, datadog.GetBrowserTestLatestResultsOptionalParameters{})

Get the last 50 test results summaries for a given Synthetics Browser test.

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

    publicId := "publicId_example" // string | The public ID of the browser test for which to search results for.
    fromTs := int64(789) // int64 | Timestamp in milliseconds from which to start querying results. (optional)
    toTs := int64(789) // int64 | Timestamp in milliseconds up to which to query results. (optional)
    probeDc := []string{"Inner_example"} // []string | Locations for which to query results. (optional)
    optionalParams := datadog.GetBrowserTestLatestResultsOptionalParameters{
        FromTs: &fromTs,
        ToTs: &toTs,
        ProbeDc: &probeDc,
    }

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.SyntheticsApi.GetBrowserTestLatestResults(ctx, publicId, optionalParams)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.GetBrowserTestLatestResults`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBrowserTestLatestResults`: SyntheticsGetBrowserTestLatestResultsResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from SyntheticsApi.GetBrowserTestLatestResults:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**publicId** | **string** | The public ID of the browser test for which to search results for. | 


### Optional Parameters


Other parameters are passed through a pointer to a GetBrowserTestLatestResultsOptionalParameters struct.


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**fromTs** | **int64** | Timestamp in milliseconds from which to start querying results. | 
**toTs** | **int64** | Timestamp in milliseconds up to which to query results. | 
**probeDc** | **[]string** | Locations for which to query results. | 

### Return type

[**SyntheticsGetBrowserTestLatestResultsResponse**](SyntheticsGetBrowserTestLatestResultsResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBrowserTestResult

> SyntheticsBrowserTestResultFull GetBrowserTestResult(ctx, publicId, resultId)

Get a specific full result from a given (browser) Synthetic test.

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

    publicId := "publicId_example" // string | The public ID of the browser test to which the target result belongs.
    resultId := "resultId_example" // string | The ID of the result to get.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.SyntheticsApi.GetBrowserTestResult(ctx, publicId, resultId)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.GetBrowserTestResult`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBrowserTestResult`: SyntheticsBrowserTestResultFull
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from SyntheticsApi.GetBrowserTestResult:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**publicId** | **string** | The public ID of the browser test to which the target result belongs. |  |
**resultId** | **string** | The ID of the result to get. | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**SyntheticsBrowserTestResultFull**](SyntheticsBrowserTestResultFull.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGlobalVariable

> SyntheticsGlobalVariable GetGlobalVariable(ctx, variableId)

Get the detailed configuration of a global variable.

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

    variableId := "variableId_example" // string | The ID of the global variable.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.SyntheticsApi.GetGlobalVariable(ctx, variableId)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.GetGlobalVariable`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGlobalVariable`: SyntheticsGlobalVariable
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from SyntheticsApi.GetGlobalVariable:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**variableId** | **string** | The ID of the global variable. | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**SyntheticsGlobalVariable**](SyntheticsGlobalVariable.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrivateLocation

> SyntheticsPrivateLocation GetPrivateLocation(ctx, locationId)

Get a Synthetics private location.

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

    locationId := "locationId_example" // string | The ID of the private location.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.SyntheticsApi.GetPrivateLocation(ctx, locationId)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.GetPrivateLocation`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPrivateLocation`: SyntheticsPrivateLocation
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from SyntheticsApi.GetPrivateLocation:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**locationId** | **string** | The ID of the private location. | 


### Optional Parameters

This endpoint does not have optional parameters.


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

> SyntheticsTestDetails GetTest(ctx, publicId)

Get the detailed configuration associated with a Synthetics test.

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

    publicId := "publicId_example" // string | The public ID of the test to get details from.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.SyntheticsApi.GetTest(ctx, publicId)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.GetTest`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTest`: SyntheticsTestDetails
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from SyntheticsApi.GetTest:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**publicId** | **string** | The public ID of the test to get details from. | 


### Optional Parameters

This endpoint does not have optional parameters.


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


## ListGlobalVariables

> SyntheticsListGlobalVariablesResponse ListGlobalVariables(ctx)

Get the list of all Synthetics global variables.

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
    resp, r, err := apiClient.SyntheticsApi.ListGlobalVariables(ctx)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.ListGlobalVariables`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGlobalVariables`: SyntheticsListGlobalVariablesResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from SyntheticsApi.ListGlobalVariables:\n%s\n", responseContent)
}
```

### Required Parameters

This endpoint does not need any parameter.


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**SyntheticsListGlobalVariablesResponse**](SyntheticsListGlobalVariablesResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLocations

> SyntheticsLocations ListLocations(ctx)

Get the list of public and private locations available for Synthetic
tests. No arguments required.

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
    resp, r, err := apiClient.SyntheticsApi.ListLocations(ctx)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.ListLocations`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLocations`: SyntheticsLocations
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from SyntheticsApi.ListLocations:\n%s\n", responseContent)
}
```

### Required Parameters

This endpoint does not need any parameter.


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**SyntheticsLocations**](SyntheticsLocations.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTests

> SyntheticsListTestsResponse ListTests(ctx)

Get the list of all Synthetic tests.

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
    resp, r, err := apiClient.SyntheticsApi.ListTests(ctx)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.ListTests`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTests`: SyntheticsListTestsResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from SyntheticsApi.ListTests:\n%s\n", responseContent)
}
```

### Required Parameters

This endpoint does not need any parameter.


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**SyntheticsListTestsResponse**](SyntheticsListTestsResponse.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TriggerCITests

> SyntheticsTriggerCITestsResponse TriggerCITests(ctx, body)

Trigger a set of Synthetics tests for continuous integration.

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

    body := *datadog.NewSyntheticsCITestBody() // SyntheticsCITestBody | Details of the test to trigger.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.SyntheticsApi.TriggerCITests(ctx, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.TriggerCITests`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TriggerCITests`: SyntheticsTriggerCITestsResponse
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from SyntheticsApi.TriggerCITests:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**body** | [**SyntheticsCITestBody**](SyntheticsCITestBody.md) | Details of the test to trigger. | 


### Optional Parameters

This endpoint does not have optional parameters.


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


## UpdateAPITest

> SyntheticsAPITest UpdateAPITest(ctx, publicId, body)

Edit the configuration of a Synthetic API test.

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

    publicId := "publicId_example" // string | The public ID of the test to get details from.
    body := *datadog.NewSyntheticsAPITest() // SyntheticsAPITest | New test details to be saved.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.SyntheticsApi.UpdateAPITest(ctx, publicId, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.UpdateAPITest`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAPITest`: SyntheticsAPITest
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from SyntheticsApi.UpdateAPITest:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**publicId** | **string** | The public ID of the test to get details from. |  |
**body** | [**SyntheticsAPITest**](SyntheticsAPITest.md) | New test details to be saved. | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**SyntheticsAPITest**](SyntheticsAPITest.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBrowserTest

> SyntheticsBrowserTest UpdateBrowserTest(ctx, publicId, body)

Edit the configuration of a Synthetic browser test.

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

    publicId := "publicId_example" // string | The public ID of the test to get details from.
    body := *datadog.NewSyntheticsBrowserTest("Message_example") // SyntheticsBrowserTest | New test details to be saved.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.SyntheticsApi.UpdateBrowserTest(ctx, publicId, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.UpdateBrowserTest`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBrowserTest`: SyntheticsBrowserTest
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from SyntheticsApi.UpdateBrowserTest:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**publicId** | **string** | The public ID of the test to get details from. |  |
**body** | [**SyntheticsBrowserTest**](SyntheticsBrowserTest.md) | New test details to be saved. | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

[**SyntheticsBrowserTest**](SyntheticsBrowserTest.md)

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePrivateLocation

> SyntheticsPrivateLocation UpdatePrivateLocation(ctx, locationId, body)

Edit a Synthetics private location.

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

    locationId := "locationId_example" // string | The ID of the private location.
    body := *datadog.NewSyntheticsPrivateLocation("Description of private location", "New private location", []string{"team:front"}) // SyntheticsPrivateLocation | Details of the private location to be updated.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.SyntheticsApi.UpdatePrivateLocation(ctx, locationId, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.UpdatePrivateLocation`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePrivateLocation`: SyntheticsPrivateLocation
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from SyntheticsApi.UpdatePrivateLocation:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**locationId** | **string** | The ID of the private location. |  |
**body** | [**SyntheticsPrivateLocation**](SyntheticsPrivateLocation.md) | Details of the private location to be updated. | 


### Optional Parameters

This endpoint does not have optional parameters.


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


## UpdateTestPauseStatus

> bool UpdateTestPauseStatus(ctx, publicId, body)

Pause or start a Synthetics test by changing the status.

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

    publicId := "publicId_example" // string | The public ID of the Synthetic test to update.
    body := *datadog.NewSyntheticsUpdateTestPauseStatusPayload() // SyntheticsUpdateTestPauseStatusPayload | Status to set the given Synthetic test to.

    configuration := datadog.NewConfiguration()

    apiClient := datadog.NewAPIClient(configuration)
    resp, r, err := apiClient.SyntheticsApi.UpdateTestPauseStatus(ctx, publicId, body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticsApi.UpdateTestPauseStatus`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTestPauseStatus`: bool
    responseContent, _ := json.MarshalIndent(resp, "", "  ")
    fmt.Fprintf(os.Stdout, "Response from SyntheticsApi.UpdateTestPauseStatus:\n%s\n", responseContent)
}
```

### Required Parameters


Name | Type | Description  | Notes
---- | ---- | ------------ | ------
**ctx** | **context.Context** | Context for authentication, logging, cancellation, deadlines, tracing, etc. |
**publicId** | **string** | The public ID of the Synthetic test to update. |  |
**body** | [**SyntheticsUpdateTestPauseStatusPayload**](SyntheticsUpdateTestPauseStatusPayload.md) | Status to set the given Synthetic test to. | 


### Optional Parameters

This endpoint does not have optional parameters.


### Return type

**bool**

### Authorization

[AuthZ](../README.md#AuthZ), [apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

