# SyntheticsBrowserTestResultData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BrowserType** | **string** | Type of browser device used for the browser test. | [optional] 
**BrowserVersion** | **string** | Browser version used for the browser test. | [optional] 
**Device** | [**SyntheticsDevice**](SyntheticsDevice.md) |  | [optional] 
**Duration** | **float64** | Global duration in second of the browser test. | [optional] 
**Error** | **string** | Error returned for the browser test. | [optional] 
**Passed** | **bool** | Whether or not the browser test was conducted. | [optional] 
**ReceivedEmailCount** | **int64** | The amount of email received during the browser test. | [optional] 
**StartUrl** | **string** | Starting URL for the browser test. | [optional] 
**StepDetails** | [**[]SyntheticsStepDetail**](SyntheticsStepDetail.md) | Array containing the different browser test steps. | [optional] 
**ThumbnailsBucketKey** | **bool** | Whether or not a thumbnail is associated with the browser test. | [optional] 
**TimeToInteractive** | **float64** | Time in second to wait before the browser test starts after reaching the start URL. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


