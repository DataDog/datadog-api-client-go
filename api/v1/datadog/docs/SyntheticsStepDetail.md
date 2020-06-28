# SyntheticsStepDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BrowserErrors** | [**[]SyntheticsBrowserError**](SyntheticsBrowserError.md) | Array of errors collected for a browser test. | [optional] 
**CheckType** | [**SyntheticsCheckType**](SyntheticsCheckType.md) |  | [optional] 
**Description** | **string** | Description of the test. | [optional] 
**Duration** | **float64** | Total duration in millisecond of the test. | [optional] 
**Error** | **string** | Error returned by the test. | [optional] 
**PlayingTab** | [**SyntheticsPlayingTab**](SyntheticsPlayingTab.md) |  | [optional] 
**Resources** | [**[]SyntheticsResource**](SyntheticsResource.md) | Array of resources collected by the test. | [optional] 
**ScreenshotBucketKey** | **bool** | Whether or not screenshots where collected by the test. | [optional] 
**Skipped** | **bool** | Whether or not to skip this step. | [optional] 
**SnapshotBucketKey** | **bool** | Whether or not snapshots where collected by the test. | [optional] 
**StepId** | **int64** | The step ID. | [optional] 
**SubTestStepDetails** | [**[]SyntheticsStepDetail**](SyntheticsStepDetail.md) | If this steps include a sub-test. [Subtests documentation](https://docs.datadoghq.com/synthetics/browser_tests/advanced_options/#subtests). | [optional] 
**TimeToInteractive** | **float64** | Time before starting the step. | [optional] 
**Type** | [**SyntheticsStepType**](SyntheticsStepType.md) |  | [optional] 
**Url** | **string** | URL to perform the step against. | [optional] 
**Value** | [**map[string]interface{}**](.md) | Value for the step. | [optional] 
**Warnings** | [**[]SyntheticsStepDetailWarnings**](SyntheticsStepDetail_warnings.md) | Warning collected that didn&#39;t failed the step. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


