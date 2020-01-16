# \HostsApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllHosts**](HostsApi.md#GetAllHosts) | **Get** /api/v1/hosts | Get all hosts for your organization
[**GetHostTotals**](HostsApi.md#GetHostTotals) | **Get** /api/v1/hosts/totals | Get total number of active hosts in your Datadog Account
[**MuteHost**](HostsApi.md#MuteHost) | **Post** /api/v1/host/{host_name}/mute | Mute a Host
[**UnmuteHost**](HostsApi.md#UnmuteHost) | **Post** /api/v1/host/{host_name}/unmute | Unmute a Host



## GetAllHosts

> HostListResponse GetAllHosts(ctx).Filter(filter).SortField(sortField).SortDir(sortDir).Start(start).Count(count).From(from).Execute()

Get all hosts for your organization



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllHostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | String to filter search results | 
 **sortField** | **string** | Sort hosts by this field | 
 **sortDir** | **string** | Direction of sort | 
 **start** | **int32** | Host result to start search from | 
 **count** | **int32** | Number of hosts to return | 
 **from** | **int32** | Number of seconds from which you want to search your hosts | 

### Return type

[**HostListResponse**](HostListResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHostTotals

> HostTotals GetHostTotals(ctx).From(from).Execute()

Get total number of active hosts in your Datadog Account



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHostTotalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **int32** | Number of seconds from which you want to get total number of active hosts | 

### Return type

[**HostTotals**](HostTotals.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MuteHost

> HostMuteResponse MuteHost(ctx, hostName).Body(body).Execute()

Mute a Host



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostName** | **string** | Name of the host to mute | 

### Other Parameters

Other parameters are passed through a pointer to a apiMuteHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**HostMuteSettings**](HostMuteSettings.md) | Mute a host | 

### Return type

[**HostMuteResponse**](HostMuteResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnmuteHost

> HostMuteResponse UnmuteHost(ctx, hostName).Execute()

Unmute a Host



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostName** | **string** | Name of the host to unmute | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnmuteHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HostMuteResponse**](HostMuteResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

