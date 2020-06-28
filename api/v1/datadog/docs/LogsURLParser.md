# LogsUrlParser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | **bool** | Whether or not the processor is enabled. | [optional] [default to false]
**Name** | **string** | Name of the processor. | [optional] 
**NormalizeEndingSlashes** | Pointer to **bool** | Normalize the ending slashes or not. | [optional] [default to false]
**Sources** | **[]string** | Array of source attributes. | [default to ["http.url"]]
**Target** | **string** | Name of the parent attribute that contains all the extracted details from the &#x60;sources&#x60;. | [default to http.url_details]
**Type** | [**LogsUrlParserType**](LogsURLParserType.md) |  | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


