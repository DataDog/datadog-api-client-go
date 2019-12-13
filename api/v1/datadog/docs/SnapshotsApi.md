# \SnapshotsApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGraphSnapshot**](SnapshotsApi.md#GetGraphSnapshot) | **Get** /api/v1/graph/snapshot | Take graph snapshots



## GetGraphSnapshot

> GraphSnapshot GetGraphSnapshot(ctx, metricQuery, start, end, optional)

Take graph snapshots

### Overview
Take graph snapshots
### Arguments
* **`metric_query`** [*required*]: The metric query.
* **`start`** [*required*]: The POSIX timestamp of the start of the query.
* **`end`** [*required*]: The POSIX timestamp of the end of the query.
* **`event_query`** [*optional*, *default* = **None**]: A query that adds event bands to the graph.
* **`graph_def`** [*optional*, *default* = **None**]: A JSON document defining the graph.
  graph_def can be used instead of metric_query. The JSON document uses the
  [grammar defined here](https://docs.datadoghq.com/graphing/graphing_json/#grammar)
  and should be formatted to a single line then URLEncoded.

* **`title`** [*optional*, *default* = **None**]: A title for the graph.
  If no title is specified, the graph doesn’t have a title.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricQuery** | **string**| The metric query. | 
**start** | **int64**| The POSIX timestamp of the start of the query. | 
**end** | **int64**| The POSIX timestamp of the end of the query. | 
 **optional** | ***GetGraphSnapshotOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetGraphSnapshotOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **eventQuery** | **optional.String**| A query that adds event bands to the graph. | 
 **graphDef** | **optional.String**| A JSON document defining the graph. graph_def can be used instead of metric_query. The JSON document uses the [grammar defined here](https://docs.datadoghq.com/graphing/graphing_json/#grammar) and should be formatted to a single line then URLEncoded. | 
 **title** | **optional.String**| A title for the graph. If no title is specified, the graph doesn’t have a title. | 

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

