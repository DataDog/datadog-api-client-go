# LogsGeoIpParser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | **bool** | Whether or not the processor is enabled. | [optional] [default to false]
**Name** | **string** | Name of the processor. | [optional] 
**Sources** | **[]string** | Array of source attributes. | [default to ["network.client.ip"]]
**Target** | **string** | Name of the parent attribute that contains all the extracted details from the &#x60;sources&#x60;. | [default to network.client.geoip]
**Type** | [**LogsGeoIpParserType**](LogsGeoIPParserType.md) |  | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


