# \SnapshotsApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGraphSnapshot**](SnapshotsApi.md#GetGraphSnapshot) | **Get** /api/v1/graph/snapshot | Take graph snapshots



## GetGraphSnapshot

> GraphSnapshot GetGraphSnapshot(ctx, start, end, optional)

Take graph snapshots

Take graph snapshots. **Note**: When a snapshot is created, there is some delay before it is available.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**start** | **int64**| The POSIX timestamp of the start of the query. | 
**end** | **int64**| The POSIX timestamp of the end of the query. | 
 **optional** | ***GetGraphSnapshotOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetGraphSnapshotOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **metricQuery** | **optional.String**| The metric query. | 
 **eventQuery** | **optional.String**| A query that adds event bands to the graph. | 
 **graphDef** | **optional.String**| A JSON document defining the graph. &#x60;graph_def&#x60; can be used instead of &#x60;metric_query&#x60;. The JSON document uses the [grammar defined here](https://docs.datadoghq.com/graphing/graphing_json/#grammar) and should be formatted to a single line then URL encoded. | 
 **title** | **optional.String**| A title for the graph. If no title is specified, the graph does not have a title. | 

### Return type

[**GraphSnapshot**](GraphSnapshot.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

