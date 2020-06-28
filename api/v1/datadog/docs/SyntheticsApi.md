# \SyntheticsApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTest**](SyntheticsApi.md#CreateTest) | **Post** /api/v1/synthetics/tests | Create a test
[**DeleteTests**](SyntheticsApi.md#DeleteTests) | **Post** /api/v1/synthetics/tests/delete | Delete tests
[**GetAPITestLatestResults**](SyntheticsApi.md#GetAPITestLatestResults) | **Get** /api/v1/synthetics/tests/{public_id}/results | Get the test&#39;s latest results summaries (API)
[**GetAPITestResult**](SyntheticsApi.md#GetAPITestResult) | **Get** /api/v1/synthetics/tests/{public_id}/results/{result_id} | Get a test result (API)
[**GetBrowserTest**](SyntheticsApi.md#GetBrowserTest) | **Get** /api/v1/synthetics/tests/browser/{public_id} | Get a test configuration (browser)
[**GetBrowserTestLatestResults**](SyntheticsApi.md#GetBrowserTestLatestResults) | **Get** /api/v1/synthetics/tests/browser/{public_id}/results | Get the test&#39;s latest results summaries (browser)
[**GetBrowserTestResult**](SyntheticsApi.md#GetBrowserTestResult) | **Get** /api/v1/synthetics/tests/browser/{public_id}/results/{result_id} | Get a test result (browser)
[**GetTest**](SyntheticsApi.md#GetTest) | **Get** /api/v1/synthetics/tests/{public_id} | Get a test configuration (API)
[**ListLocations**](SyntheticsApi.md#ListLocations) | **Get** /api/v1/synthetics/locations | Get all locations (public and private)
[**ListTests**](SyntheticsApi.md#ListTests) | **Get** /api/v1/synthetics/tests | Get a list of tests
[**TriggerTests**](SyntheticsApi.md#TriggerTests) | **Post** /api/v1/synthetics/tests/trigger/ci | Create a test
[**UpdateTest**](SyntheticsApi.md#UpdateTest) | **Put** /api/v1/synthetics/tests/{public_id} | Edit a test
[**UpdateTestPauseStatus**](SyntheticsApi.md#UpdateTestPauseStatus) | **Put** /api/v1/synthetics/tests/{public_id}/status | Pause or start a test



## CreateTest

> SyntheticsTestDetails CreateTest(ctx, body)

Create a test

Create a Synthetic test.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**SyntheticsTestDetails**](SyntheticsTestDetails.md)| Details of the test to create. | 

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

> SyntheticsDeleteTestsResponse DeleteTests(ctx, body)

Delete tests

Delete multiple Synthetic tests by ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**SyntheticsDeleteTestsPayload**](SyntheticsDeleteTestsPayload.md)| Public ID list of the Synthetic tests to be deleted. | 

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

> SyntheticsGetApiTestLatestResultsResponse GetAPITestLatestResults(ctx, publicId, optional)

Get the test's latest results summaries (API)

Get the last 50 test results summaries for a given Synthetics API test.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicId** | **string**| The public ID of the test for which to search results for. | 
 **optional** | ***GetAPITestLatestResultsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAPITestLatestResultsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fromTs** | **optional.Int64**| Timestamp from which to start querying results. | 
 **toTs** | **optional.Int64**| Timestamp up to which to query results. | 
 **probeDc** | [**optional.Interface of []string**](string.md)| Locations for which to query results. | 

### Return type

[**SyntheticsGetApiTestLatestResultsResponse**](SyntheticsGetAPITestLatestResultsResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAPITestResult

> SyntheticsApiTestResultFull GetAPITestResult(ctx, publicId, resultId)

Get a test result (API)

Get a specific full result from a given (API) Synthetic test.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicId** | **string**| The public ID of the API test to which the target result belongs. | 
**resultId** | **string**| The ID of the result to get. | 

### Return type

[**SyntheticsApiTestResultFull**](SyntheticsAPITestResultFull.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBrowserTest

> SyntheticsTestDetails GetBrowserTest(ctx, publicId)

Get a test configuration (browser)

Get the detailed configuration (including steps) associated with a Synthetics browser test.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicId** | **string**| The public ID of the test to get details from. | 

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

> SyntheticsGetBrowserTestLatestResultsResponse GetBrowserTestLatestResults(ctx, publicId, optional)

Get the test's latest results summaries (browser)

Get the last 50 test results summaries for a given Synthetics Browser test.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicId** | **string**| The public ID of the browser test for which to search results for. | 
 **optional** | ***GetBrowserTestLatestResultsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetBrowserTestLatestResultsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fromTs** | **optional.Int64**| Timestamp from which to start querying results. | 
 **toTs** | **optional.Int64**| Timestamp up to which to query results. | 
 **probeDc** | [**optional.Interface of []string**](string.md)| Locations for which to query results. | 

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

> SyntheticsBrowserTestResultFull GetBrowserTestResult(ctx, publicId, resultId)

Get a test result (browser)

Get a specific full result from a given (browser) Synthetic test.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicId** | **string**| The public ID of the browser test to which the target result belongs. | 
**resultId** | **string**| The ID of the result to get. | 

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

> SyntheticsTestDetails GetTest(ctx, publicId)

Get a test configuration (API)

Get the detailed configuration associated with a Synthetics test.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicId** | **string**| The public ID of the test to get details from. | 

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

> SyntheticsLocations ListLocations(ctx, )

Get all locations (public and private)

Get the list of public and private locations available for Synthetics tests. No arguments required.

### Required Parameters

This endpoint does not need any parameter.

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

> SyntheticsListTestsResponse ListTests(ctx, optional)

Get a list of tests

Get the list of all Synthetic tests (can be filtered by type).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListTestsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListTestsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **checkType** | **optional.String**| API or browser to filter the list by test type, undefined to get the unfiltered list. | 

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


## TriggerTests

> SyntheticsTestsTriggerResults TriggerTests(ctx, body)

Create a test

Trigger Synthetic CI tests.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**SyntheticsTestsTrigger**](SyntheticsTestsTrigger.md)| Details of the test to create. | 

### Return type

[**SyntheticsTestsTriggerResults**](SyntheticsTestsTriggerResults.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTest

> SyntheticsTestDetails UpdateTest(ctx, publicId, body)

Edit a test

Edit the configuration of a Synthetic test.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicId** | **string**| The public ID of the test to get details from. | 
**body** | [**SyntheticsTestDetails**](SyntheticsTestDetails.md)| New test details to be saved. | 

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

> bool UpdateTestPauseStatus(ctx, publicId, body)

Pause or start a test

Pause or start a Synthetics test by changing the status.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicId** | **string**| The public ID of the Synthetic test to update. | 
**body** | [**SyntheticsUpdateTestPauseStatusPayload**](SyntheticsUpdateTestPauseStatusPayload.md)| Status to set the given Synthetic test to. | 

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

