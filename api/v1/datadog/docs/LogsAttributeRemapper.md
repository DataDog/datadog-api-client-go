# LogsAttributeRemapper

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | **bool** | Whether or not the processor is enabled. | [optional] [default to false]
**Name** | **string** | Name of the processor. | [optional] 
**OverrideOnConflict** | **bool** | Override or not the target element if already set, | [optional] [default to false]
**PreserveSource** | **bool** | Remove or preserve the remapped source element. | [optional] [default to false]
**SourceType** | **string** | Defines if the sources are from log &#x60;attribute&#x60; or &#x60;tag&#x60;. | [optional] [default to attribute]
**Sources** | **[]string** | Array of source attributes. | 
**Target** | **string** | Final attribute or tag name to remap the sources to. | 
**TargetType** | **string** | Defines if the sources are from log &#x60;attribute&#x60; or &#x60;tag&#x60;. | [optional] [default to attribute]
**Type** | [**LogsAttributeRemapperType**](LogsAttributeRemapperType.md) |  | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


