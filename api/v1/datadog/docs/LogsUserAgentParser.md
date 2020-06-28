# LogsUserAgentParser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | **bool** | Whether or not the processor is enabled. | [optional] [default to false]
**IsEncoded** | **bool** | Define if the source attribute is URL encoded or not. | [optional] [default to false]
**Name** | **string** | Name of the processor. | [optional] 
**Sources** | **[]string** | Array of source attributes. | [default to ["http.useragent"]]
**Target** | **string** | Name of the parent attribute that contains all the extracted details from the &#x60;sources&#x60;. | [default to http.useragent_details]
**Type** | [**LogsUserAgentParserType**](LogsUserAgentParserType.md) |  | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


