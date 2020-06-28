# SyntheticsApiTestResultData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cert** | [**SyntheticsSslCertificate**](SyntheticsSSLCertificate.md) |  | [optional] 
**ErrorCode** | [**SyntheticsErrorCode**](SyntheticsErrorCode.md) |  | [optional] 
**ErrorMessage** | **string** | The API test error message. | [optional] 
**EventType** | [**SyntheticsTestProcessStatus**](SyntheticsTestProcessStatus.md) |  | [optional] 
**HttpStatusCode** | **int64** | The API test HTTP status code. | [optional] 
**RequestHeaders** | **map[string]map[string]interface{}** | Request header object used for the API test. | [optional] 
**ResponseBody** | **string** | Response body returned for the API test. | [optional] 
**ResponseHeaders** | **map[string]map[string]interface{}** | Response headers returned for the API test. | [optional] 
**ResponseSize** | **int64** | Global size in byte of the API test response. | [optional] 
**Timings** | [**SyntheticsTiming**](SyntheticsTiming.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


