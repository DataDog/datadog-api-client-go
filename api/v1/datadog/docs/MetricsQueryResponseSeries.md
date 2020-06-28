# MetricsQueryResponseSeries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggr** | **string** | Aggregation type. | [optional] [readonly] 
**DisplayName** | **string** | Display name of the metric. | [optional] [readonly] 
**End** | **int64** | End of the time window, milliseconds since Unix epoch. | [optional] [readonly] 
**Expression** | **string** | Metric expression. | [optional] [readonly] 
**Interval** | **int64** | Number of seconds between data samples. | [optional] [readonly] 
**Length** | **int64** | Number of data samples. | [optional] [readonly] 
**Metric** | **string** | Metric name. | [optional] [readonly] 
**Pointlist** | [**[][]float64**](array.md) | List of points of the time series. | [optional] [readonly] 
**Scope** | **string** | Metric scope, comma separated list of tags. | [optional] [readonly] 
**Start** | **int64** | Start of the time window, milliseconds since Unix epoch. | [optional] [readonly] 
**Unit** | [**[]MetricsQueryResponseUnit**](MetricsQueryResponse_unit.md) | Detailed information about the metric unit. First element describes the \&quot;primary unit\&quot; (for example, &#x60;bytes&#x60; in &#x60;bytes per second&#x60;), second describes the \&quot;per unit\&quot; (for example, &#x60;second&#x60; in &#x60;bytes per second&#x60;). | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


