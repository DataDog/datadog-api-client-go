# LogsArithmeticProcessor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expression** | **string** | Arithmetic operation between one or more log attributes. | 
**IsEnabled** | **bool** | Whether or not the processor is enabled. | [optional] [default to false]
**IsReplaceMissing** | **bool** | If &#x60;true&#x60;, it replaces all missing attributes of expression by &#x60;0&#x60;, &#x60;false&#x60; skip the operation if an attribute is missing. | [optional] [default to false]
**Name** | **string** | Name of the processor. | [optional] 
**Target** | **string** | Name of the attribute that contains the result of the arithmetic operation. | 
**Type** | [**LogsArithmeticProcessorType**](LogsArithmeticProcessorType.md) |  | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


