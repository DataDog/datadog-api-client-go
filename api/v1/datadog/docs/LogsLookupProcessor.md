# LogsLookupProcessor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultLookup** | **string** | Value to set the target attribute if the source value is not found in the list. | [optional] 
**IsEnabled** | **bool** | Whether or not the processor is enabled. | [optional] [default to false]
**LookupTable** | **[]string** | Mapping table of values for the source attribute and their associated target attribute values, formatted as &#x60;[\&quot;source_key1,target_value1\&quot;, \&quot;source_key2,target_value2\&quot;]&#x60; | 
**Name** | **string** | Name of the processor. | [optional] 
**Source** | **string** | Source attribute used to perform the lookup. | 
**Target** | **string** | Name of the attribute that contains the corresponding value in the mapping list or the &#x60;default_lookup&#x60; if not found in the mapping list. | 
**Type** | [**LogsLookupProcessorType**](LogsLookupProcessorType.md) |  | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


