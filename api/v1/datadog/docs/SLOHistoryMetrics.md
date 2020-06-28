# SloHistoryMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Denominator** | [**SloHistoryMetricsSeries**](SLOHistoryMetricsSeries.md) |  | 
**Interval** | **int64** | The aggregated query interval for the series data. It&#39;s implicit based on the query time window. | 
**Message** | **string** | Optional message if there are specific query issues/warnings. | [optional] 
**Numerator** | [**SloHistoryMetricsSeries**](SLOHistoryMetricsSeries.md) |  | 
**Query** | **string** | The combined numerator and denominator query CSV. | 
**ResType** | **string** | The series result type. This mimics &#x60;batch_query&#x60; response type. | 
**RespVersion** | **int64** | The series response version type. This mimics &#x60;batch_query&#x60; response type. | 
**Times** | **[]float64** | An array of query timestamps in EPOCH milliseconds | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


