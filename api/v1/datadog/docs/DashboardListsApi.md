# \DashboardListsApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDashboardList**](DashboardListsApi.md#CreateDashboardList) | **Post** /api/v1/dashboard/lists/manual | Create a dashboard list
[**DeleteDashboardList**](DashboardListsApi.md#DeleteDashboardList) | **Delete** /api/v1/dashboard/lists/manual/{list_id} | Delete a dashboard list
[**GetAllDashboardLists**](DashboardListsApi.md#GetAllDashboardLists) | **Get** /api/v1/dashboard/lists/manual | Get all dashboard lists
[**GetDashboardList**](DashboardListsApi.md#GetDashboardList) | **Get** /api/v1/dashboard/lists/manual/{list_id} | Get a dashboard list
[**UpdateDashboardList**](DashboardListsApi.md#UpdateDashboardList) | **Put** /api/v1/dashboard/lists/manual/{list_id} | Update a dashboard list



## CreateDashboardList

> DashboardList CreateDashboardList(ctx, dashboardList)

Create a dashboard list

### Overview
Create an empty dashboard list.
### Arguments
* **name** [*required*]: The name of the dashboard list.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardList** | [**DashboardList**](DashboardList.md)| DashboardList request object | 

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

### Overview
Delete a dashboard list.
### Arguments
This endpoint takes no JSON arguments.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **int64**| ID of the dashboard list to delete | 

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


## GetAllDashboardLists

> DashboardListListResponse GetAllDashboardLists(ctx, )

Get all dashboard lists

### Overview
Fetch all of your existing dashboard list definitions.
### Arguments
This endpoint takes no JSON arguments.

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


## GetDashboardList

> DashboardList GetDashboardList(ctx, listId)

Get a dashboard list

### Overview
Fetch an existing dashboard list's definition.
### Arguments
This endpoint takes no JSON arguments.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **int64**| ID of the dashboard list to fetch | 

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


## UpdateDashboardList

> DashboardList UpdateDashboardList(ctx, listId, dashboardList)

Update a dashboard list

### Overview
Update the name of a dashboard list.
### Arguments
* **name** [*required*]: The name of the dashboard list.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **int64**| ID of the dashboard list to update | 
**dashboardList** | [**DashboardList**](DashboardList.md)| DashboardList request object | 

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

