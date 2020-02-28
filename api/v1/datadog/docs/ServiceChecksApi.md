# \ServiceChecksApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubmitServiceCheck**](ServiceChecksApi.md#SubmitServiceCheck) | **Post** /api/v1/check_run | Submit a Service Check



## SubmitServiceCheck

> IntakePayloadAccepted SubmitServiceCheck(ctx).Body(body).Execute()

Submit a Service Check



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitServiceCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**[]ServiceCheck**](ServiceCheck.md) | Service Check body | 

### Return type

[**IntakePayloadAccepted**](IntakePayloadAccepted.md)

### Authorization

[apiKeyAuthQuery](../README.md#apiKeyAuthQuery)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

