# \LogsHTTPIntakeApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SendLog**](LogsHTTPIntakeApi.md#SendLog) | **Post** /v1/input | 



## SendLog

> interface{} SendLog(ctx).Body(body).Execute()





### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**HttpLog**](HttpLog.md) | Log to send (JSON format) | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[apiKeyAuthHeader](../README.md#apiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

