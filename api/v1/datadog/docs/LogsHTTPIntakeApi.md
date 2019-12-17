# \LogsHTTPIntakeApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SendLog**](LogsHTTPIntakeApi.md#SendLog) | **Post** /v1/input/{api_key} | 



## SendLog

> interface{} SendLog(ctx, apiKey, httpLog)



Send logs

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiKey** | **string**| valid API key for an organisation | 
**httpLog** | [**HttpLog**](HttpLog.md)| Log to send (JSON format) | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

