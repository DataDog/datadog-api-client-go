# \DashboardListsApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddDashboardListItems**](DashboardListsApi.md#AddDashboardListItems) | **Post** /api/v2/dashboard/lists/manual/{dashboard_list_id}/dashboards | Add Items to a Dashboard List
[**DeleteDashboardListItems**](DashboardListsApi.md#DeleteDashboardListItems) | **Delete** /api/v2/dashboard/lists/manual/{dashboard_list_id}/dashboards | Delete Items from a Dashboard List
[**GetDashboardListItems**](DashboardListsApi.md#GetDashboardListItems) | **Get** /api/v2/dashboard/lists/manual/{dashboard_list_id}/dashboards | Get a Dashboard List
[**UpdateDashboardListItems**](DashboardListsApi.md#UpdateDashboardListItems) | **Put** /api/v2/dashboard/lists/manual/{dashboard_list_id}/dashboards | Update Items of a Dashboard List



## AddDashboardListItems

> DashboardListAddItemsResponse AddDashboardListItems(ctx, dashboardListId).Body(body).Execute()

Add Items to a Dashboard List



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardListId** | **int64** | ID of the dashboard list to add items to | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddDashboardListItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DashboardListItems**](DashboardListItems.md) | Dashboards to add to the dashboard list | 

### Return type

[**DashboardListAddItemsResponse**](DashboardListAddItemsResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDashboardListItems

> DashboardListDeleteItemsResponse DeleteDashboardListItems(ctx, dashboardListId).Body(body).Execute()

Delete Items from a Dashboard List



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardListId** | **int64** | ID of the dashboard list to delete items from | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDashboardListItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DashboardListItems**](DashboardListItems.md) | Dashboards to delete from the dashboard list | 

### Return type

[**DashboardListDeleteItemsResponse**](DashboardListDeleteItemsResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDashboardListItems

> DashboardListItems GetDashboardListItems(ctx, dashboardListId).Execute()

Get a Dashboard List



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardListId** | **int64** | ID of the dashboard list to get items from | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDashboardListItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DashboardListItems**](DashboardListItems.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDashboardListItems

> DashboardListItems UpdateDashboardListItems(ctx, dashboardListId).Body(body).Execute()

Update Items of a Dashboard List



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardListId** | **int64** | ID of the dashboard list to update items from. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDashboardListItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DashboardListItems**](DashboardListItems.md) | New dashboards of the dashboard list. | 

### Return type

[**DashboardListItems**](DashboardListItems.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

