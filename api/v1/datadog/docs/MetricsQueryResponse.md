# MetricsQueryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | **string** | Message indicating the errors if status is not &#x60;ok&#x60;. | [optional] [readonly] 
**FromDate** | **int64** | Start of requested time window, milliseconds since Unix epoch. | [optional] [readonly] 
**GroupBy** | **[]string** | List of tag keys on which to group. | [optional] [readonly] 
**Message** | **string** | Message indicating &#x60;success&#x60; if status is &#x60;ok&#x60;. | [optional] [readonly] 
**Query** | **string** | Query string | [optional] [readonly] 
**ResType** | **string** | Type of response. | [optional] [readonly] 
**Series** | [**[]MetricsQueryResponseSeries**](MetricsQueryResponse_series.md) | List of timeseries queried. | [optional] [readonly] 
**Status** | **string** | Status of the query. | [optional] [readonly] 
**ToDate** | **int64** | End of requested time window, milliseconds since Unix epoch. | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


