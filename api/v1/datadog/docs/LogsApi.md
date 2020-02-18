# \LogsApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListLogs**](LogsApi.md#ListLogs) | **Post** /api/v1/logs-queries/list | Get a list of logs
[**SendLog**](LogsApi.md#SendLog) | **Post** /v1/input | 



## ListLogs

> LogsListResponse ListLogs(ctx).Body(body).Execute()

Get a list of logs



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**LogsListRequest**](LogsListRequest.md) | Logs filter | 

### Return type

[**LogsListResponse**](LogsListResponse.md)

### Authorization

[apiKeyAuthHeader](../README.md#apiKeyAuthHeader), [appKeyAuthHeader](../README.md#appKeyAuthHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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

