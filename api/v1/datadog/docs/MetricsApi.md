# \MetricsApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubmitMetrics**](MetricsApi.md#SubmitMetrics) | **Post** /api/v1/series | Submit metrics



## SubmitMetrics

> IntakePayloadAccepted SubmitMetrics(ctx).Series(series).Execute()

Submit metrics



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **series** | [**Series**](Series.md) |  | 

### Return type

[**IntakePayloadAccepted**](IntakePayloadAccepted.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

