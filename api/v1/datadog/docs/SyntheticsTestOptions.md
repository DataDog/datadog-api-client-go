# SyntheticsTestOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptSelfSigned** | **bool** | For browser test, whether or not the test should allow self signed certificate. | [optional] 
**AllowInsecure** | **bool** | Allows loading insecure content for an HTTP request. | [optional] 
**DeviceIds** | [**[]SyntheticsDeviceId**](SyntheticsDeviceID.md) | Array with the different device IDs used to run the test. | [optional] 
**FollowRedirects** | **bool** | For API SSL test, whether or not the test should follow redirects. | [optional] 
**MinFailureDuration** | **int64** | Minimum amount of time before declaring the test has failed. | [optional] 
**MinLocationFailed** | **int64** | Minimum amount of locations that are allowed to fail for the test. | [optional] 
**Retry** | [**SyntheticsTestOptionsRetry**](SyntheticsTestOptions_retry.md) |  | [optional] 
**TickEvery** | [**SyntheticsTickInterval**](SyntheticsTickInterval.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


