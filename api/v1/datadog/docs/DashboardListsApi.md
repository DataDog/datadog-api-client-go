# \DashboardListsApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDashboardList**](DashboardListsApi.md#CreateDashboardList) | **Post** /api/v1/dashboard/lists/manual | Create a dashboard list
[**DeleteDashboardList**](DashboardListsApi.md#DeleteDashboardList) | **Delete** /api/v1/dashboard/lists/manual/{list_id} | Delete a dashboard list
[**GetDashboardList**](DashboardListsApi.md#GetDashboardList) | **Get** /api/v1/dashboard/lists/manual/{list_id} | Get a dashboard list
[**ListDashboardLists**](DashboardListsApi.md#ListDashboardLists) | **Get** /api/v1/dashboard/lists/manual | Get all dashboard lists
[**UpdateDashboardList**](DashboardListsApi.md#UpdateDashboardList) | **Put** /api/v1/dashboard/lists/manual/{list_id} | Update a dashboard list



## CreateDashboardList

> DashboardList CreateDashboardList(ctx, body)

Create a dashboard list

Create an empty dashboard list.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**DashboardList**](DashboardList.md)| Create a dashboard list request body. | 

### Return type

[**DashboardList**](DashboardList.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDashboardList

> DashboardListDeleteResponse DeleteDashboardList(ctx, listId)

Delete a dashboard list

Delete a dashboard list.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **int64**| ID of the dashboard list to delete. | 

### Return type

[**DashboardListDeleteResponse**](DashboardListDeleteResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDashboardList

> DashboardList GetDashboardList(ctx, listId)

Get a dashboard list

Fetch an existing dashboard list's definition.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **int64**| ID of the dashboard list to fetch. | 

### Return type

[**DashboardList**](DashboardList.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDashboardLists

> DashboardListListResponse ListDashboardLists(ctx, )

Get all dashboard lists

Fetch all of your existing dashboard list definitions.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**DashboardListListResponse**](DashboardListListResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDashboardList

> DashboardList UpdateDashboardList(ctx, listId, body)

Update a dashboard list

Update the name of a dashboard list.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **int64**| ID of the dashboard list to update. | 
**body** | [**DashboardList**](DashboardList.md)| Update a dashboard list request body. | 

### Return type

[**DashboardList**](DashboardList.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

