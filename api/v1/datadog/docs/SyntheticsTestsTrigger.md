# SyntheticsTestsTrigger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublicId** | **string** | The test public ID. | [optional] 
**AllowInsecureCertificates** | **bool** | Allows loading insecure content for an HTTP request. | [optional] 
**BasicAuth** | [**SyntheticsTestRequestBasicAuth**](SyntheticsTestRequest_basicAuth.md) |  | [optional] 
**Body** | **string** | Body to include in the test. | [optional] 
**BodyType** | **string** | Body to include in the test. | [optional] 
**DeviceIds** | [**[]SyntheticsDeviceId**](SyntheticsDeviceID.md) | Array with the different device IDs used to run the test. | [optional] 
**FollowRedirects** | **bool** | For API SSL test, whether or not the test should follow redirects. | [optional] 
**Headers** | **map[string]string** | Headers to include when performing the test. | [optional] 
**Locations** | **[]string** | Array of locations used to run the test. | [optional] 
**Retry** | [**SyntheticsTestOptionsRetry**](SyntheticsTestOptions_retry.md) |  | [optional] 
**StartUrl** | **string** | Starting URL for the browser test. | [optional] 
**Variables** | [**[]SyntheticsBrowserVariable**](SyntheticsBrowserVariable.md) | Array of variables used for the test. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


