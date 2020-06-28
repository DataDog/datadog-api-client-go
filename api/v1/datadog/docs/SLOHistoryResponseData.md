# SloHistoryResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromTs** | **int64** | The &#x60;from&#x60; timestamp in epoch seconds. | [optional] 
**Groups** | [**SloHistorySliData**](SLOHistorySLIData.md) |  | [optional] 
**Overall** | [**SloHistorySliData**](SLOHistorySLIData.md) |  | [optional] 
**Series** | [**SloHistoryMetrics**](SLOHistoryMetrics.md) |  | [optional] 
**Thresholds** | [**map[string]SloThreshold**](SLOThreshold.md) | mapping of string timeframe to the SLO threshold. | [optional] 
**ToTs** | **int64** | The &#x60;to&#x60; timestamp in epoch seconds. | [optional] 
**Type** | [**SloType**](SLOType.md) |  | [optional] 
**TypeId** | [**SloTypeNumeric**](SLOTypeNumeric.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


