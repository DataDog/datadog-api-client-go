# SloThreshold

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Target** | **float64** | The target value for the service level indicator within the corresponding timeframe. | 
**TargetDisplay** | **string** | A string representation of the target that indicates its precision. It uses trailing zeros to show significant decimal places (e.g. &#x60;98.00&#x60;).  Always included in service level objective responses. Ignored in create/update requests. | [optional] 
**Timeframe** | [**SloTimeframe**](SLOTimeframe.md) |  | 
**Warning** | **float64** | The warning value for the service level objective. | [optional] 
**WarningDisplay** | **string** | A string representation of the warning target (see the description of the &#x60;target_display&#x60; field for details).  Included in service level objective responses if a warning target exists. Ignored in create/update requests. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


