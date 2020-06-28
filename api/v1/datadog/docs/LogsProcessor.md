# LogsProcessor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Grok** | [**LogsGrokParserRules**](LogsGrokParserRules.md) |  | 
**IsEnabled** | **bool** | Whether or not the processor is enabled. | [optional] [default to false]
**Name** | **string** | Name of the processor. | [optional] 
**Samples** | **[]string** | List of sample logs to test this grok parser. | [optional] 
**Source** | **string** | Source attribute used to perform the lookup. | 
**Type** | [**LogsTraceRemapperType**](LogsTraceRemapperType.md) |  | 
**Sources** | **[]string** | Array of source attributes. | [default to ["dd.trace_id"]]
**OverrideOnConflict** | **bool** | Override or not the target element if already set, | [optional] [default to false]
**PreserveSource** | **bool** | Remove or preserve the remapped source element. | [optional] [default to false]
**SourceType** | **string** | Defines if the sources are from log &#x60;attribute&#x60; or &#x60;tag&#x60;. | [optional] [default to attribute]
**Target** | **string** | Name of the attribute that contains the corresponding value in the mapping list or the &#x60;default_lookup&#x60; if not found in the mapping list. | 
**TargetType** | **string** | Defines if the sources are from log &#x60;attribute&#x60; or &#x60;tag&#x60;. | [optional] [default to attribute]
**NormalizeEndingSlashes** | Pointer to **bool** | Normalize the ending slashes or not. | [optional] [default to false]
**IsEncoded** | **bool** | Define if the source attribute is URL encoded or not. | [optional] [default to false]
**Categories** | [**[]LogsCategoryProcessorCategories**](LogsCategoryProcessor_categories.md) | Array of filters to match or not a log and their corresponding &#x60;name&#x60;to assign a custom value to the log. | 
**Expression** | **string** | Arithmetic operation between one or more log attributes. | 
**IsReplaceMissing** | **bool** | If true, it replaces all missing attributes of &#x60;template&#x60; by an empty string. If &#x60;false&#x60; (default), skips the operation for missing attributes. | [optional] [default to false]
**Template** | **string** | A formula with one or more attributes and raw text. | 
**Filter** | [**LogsFilter**](LogsFilter.md) |  | [optional] 
**Processors** | [**[]LogsProcessor**](LogsProcessor.md) | Ordered list of processors in this pipeline. | [optional] 
**DefaultLookup** | **string** | Value to set the target attribute if the source value is not found in the list. | [optional] 
**LookupTable** | **[]string** | Mapping table of values for the source attribute and their associated target attribute values, formatted as &#x60;[\&quot;source_key1,target_value1\&quot;, \&quot;source_key2,target_value2\&quot;]&#x60; | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


