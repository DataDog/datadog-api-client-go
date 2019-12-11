# \SnapshotsApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGraphSnapshot**](SnapshotsApi.md#GetGraphSnapshot) | **Get** /api/v1/graph/snapshot | Take graph snapshots



## GetGraphSnapshot

> GraphSnapshot GetGraphSnapshot(ctx).MetricQuery(metricQuery).Start(start).End(end).EventQuery(eventQuery).GraphDef(graphDef).Title(title).Execute()

Take graph snapshots



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGraphSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **metricQuery** | **string** | The metric query. | 
 **start** | **int64** | The POSIX timestamp of the start of the query. | 
 **end** | **int64** | The POSIX timestamp of the end of the query. | 
 **eventQuery** | **string** | A query that adds event bands to the graph. | 
 **graphDef** | **string** | A JSON document defining the graph. graph_def can be used instead of metric_query. The JSON document uses the [grammar defined here](https://docs.datadoghq.com/graphing/graphing_json/#grammar) and should be formatted to a single line then URLEncoded. | 
 **title** | **string** | A title for the graph. If no title is specified, the graph doesn’t have a title. | 

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

