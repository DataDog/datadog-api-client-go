# \SyntheticsApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTest**](SyntheticsApi.md#CreateTest) | **Post** /api/v1/synthetics/tests | Create or clone test
[**DeleteTests**](SyntheticsApi.md#DeleteTests) | **Post** /api/v1/synthetics/tests/delete | Delete multiple tests
[**GetAPITestLatestResults**](SyntheticsApi.md#GetAPITestLatestResults) | **Get** /api/v1/synthetics/tests/{public_id}/results | Get test latest results (as summaries)
[**GetAPITestResult**](SyntheticsApi.md#GetAPITestResult) | **Get** /api/v1/synthetics/tests/{public_id}/results/{result_id} | Get test result (API)
[**GetAllDevices**](SyntheticsApi.md#GetAllDevices) | **Get** /api/v1/synthetics/browser/devices | Get all devices
[**GetAllLocations**](SyntheticsApi.md#GetAllLocations) | **Get** /api/v1/synthetics/locations | Get all locations (public and private)
[**GetAllTests**](SyntheticsApi.md#GetAllTests) | **Get** /api/v1/synthetics/tests | Get all test
[**GetBrowserTestLatestResults**](SyntheticsApi.md#GetBrowserTestLatestResults) | **Get** /api/v1/synthetics/tests/browser/{public_id}/results | Get test latest results (as summaries)
[**GetBrowserTestResult**](SyntheticsApi.md#GetBrowserTestResult) | **Get** /api/v1/synthetics/tests/browser/{public_id}/results/{result_id} | Get test result (browser)
[**GetTest**](SyntheticsApi.md#GetTest) | **Get** /api/v1/synthetics/tests/{public_id} | Get test
[**SetTestPauseStatus**](SyntheticsApi.md#SetTestPauseStatus) | **Put** /api/v1/synthetics/tests/{public_id}/status | Change test pause/live status
[**UpdateTest**](SyntheticsApi.md#UpdateTest) | **Put** /api/v1/synthetics/tests/{public_id} | Update test



## CreateTest

> SyntheticsTestDetails CreateTest(ctx).Body(body).FromTestId(fromTestId).Execute()

Create or clone test



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SyntheticsTestDetails**](SyntheticsTestDetails.md) | Details of the test to create | 
 **fromTestId** | **string** | Public id of the test to clone, undefined if the test is created ex nihilo | 

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

Delete multiple tests



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SyntheticsDeleteTestsPayload**](SyntheticsDeleteTestsPayload.md) | Public id list of the Synthetics tests to be deleted | 

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

> SyntheticsGetApiTestLatestResultsResponse GetAPITestLatestResults(ctx, publicId).Body(body).Execute()

Get test latest results (as summaries)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicId** | **string** | The public id of the test for which to search results for | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAPITestLatestResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**SyntheticsGetTestLatestResultsPayload**](SyntheticsGetTestLatestResultsPayload.md) | Pause/live status to set the given Synthetics test to | 

### Return type

[**SyntheticsGetApiTestLatestResultsResponse**](SyntheticsGetAPITestLatestResultsResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAPITestResult

> SyntheticsApiTestResultFull GetAPITestResult(ctx, publicId, resultId).Execute()

Get test result (API)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicId** | **string** | The public id of the API test to which the target result belongs | 
**resultId** | **string** | The id of the result to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAPITestResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## GetAllDevices

> SyntheticsDevices GetAllDevices(ctx).Execute()

Get all devices



### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllDevicesRequest struct via the builder pattern


### Return type

[**SyntheticsDevices**](SyntheticsDevices.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllLocations

> SyntheticsLocations GetAllLocations(ctx).Execute()

Get all locations (public and private)



### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllLocationsRequest struct via the builder pattern


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


## GetAllTests

> SyntheticsGetAllTestsResponse GetAllTests(ctx).CheckType(checkType).Execute()

Get all test



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllTestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **checkType** | **string** | api or browser to filter the list by type, undefined to get the unfiltered list | 

### Return type

[**SyntheticsGetAllTestsResponse**](SyntheticsGetAllTestsResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBrowserTestLatestResults

> SyntheticsGetBrowserTestLatestResultsResponse GetBrowserTestLatestResults(ctx, publicId).Body(body).Execute()

Get test latest results (as summaries)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicId** | **string** | The public id of the browser test for which to search results for | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBrowserTestLatestResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**SyntheticsGetTestLatestResultsPayload**](SyntheticsGetTestLatestResultsPayload.md) | Pause/live status to set the given Synthetics test to | 

### Return type

[**SyntheticsGetBrowserTestLatestResultsResponse**](SyntheticsGetBrowserTestLatestResultsResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBrowserTestResult

> SyntheticsBrowserTestResultFull GetBrowserTestResult(ctx, publicId, resultId).Execute()

Get test result (browser)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicId** | **string** | The public id of the browser test to which the target result belongs | 
**resultId** | **string** | The id of the result to get | 

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

Get test



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicId** | **string** | The public id of the test to get details from | 

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


## SetTestPauseStatus

> bool SetTestPauseStatus(ctx, publicId).Body(body).Execute()

Change test pause/live status



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicId** | **string** | The public id of the Synthetics test to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetTestPauseStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**SyntheticsSetTestPauseStatusPayload**](SyntheticsSetTestPauseStatusPayload.md) | Pause/live status to set the given Synthetics test to | 

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


## UpdateTest

> SyntheticsTestDetails UpdateTest(ctx, publicId).Body(body).Execute()

Update test



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicId** | **string** | The public id of the test to get details from | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**SyntheticsTestDetails**](SyntheticsTestDetails.md) | New test details to be saved | 

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

