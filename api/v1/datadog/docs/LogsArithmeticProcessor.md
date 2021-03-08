# LogsArithmeticProcessor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expression** | **string** | Arithmetic operation between one or more log attributes. | 
**IsEnabled** | Pointer to **bool** | Whether or not the processor is enabled. | [optional] [default to false]
**IsReplaceMissing** | Pointer to **bool** | If &#x60;true&#x60;, it replaces all missing attributes of expression by &#x60;0&#x60;, &#x60;false&#x60; skip the operation if an attribute is missing. | [optional] [default to false]
**Name** | Pointer to **string** | Name of the processor. | [optional] 
**Target** | **string** | Name of the attribute that contains the result of the arithmetic operation. | 
**Type** | [**LogsArithmeticProcessorType**](LogsArithmeticProcessorType.md) |  | [default to LOGSARITHMETICPROCESSORTYPE_ARITHMETIC_PROCESSOR]

## Methods

### NewLogsArithmeticProcessor

`func NewLogsArithmeticProcessor(expression string, target string, type_ LogsArithmeticProcessorType, ) *LogsArithmeticProcessor`

NewLogsArithmeticProcessor instantiates a new LogsArithmeticProcessor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogsArithmeticProcessorWithDefaults

`func NewLogsArithmeticProcessorWithDefaults() *LogsArithmeticProcessor`

NewLogsArithmeticProcessorWithDefaults instantiates a new LogsArithmeticProcessor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpression

`func (o *LogsArithmeticProcessor) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *LogsArithmeticProcessor) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *LogsArithmeticProcessor) SetExpression(v string)`

SetExpression sets Expression field to given value.


### GetIsEnabled

`func (o *LogsArithmeticProcessor) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *LogsArithmeticProcessor) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *LogsArithmeticProcessor) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *LogsArithmeticProcessor) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetIsReplaceMissing

`func (o *LogsArithmeticProcessor) GetIsReplaceMissing() bool`

GetIsReplaceMissing returns the IsReplaceMissing field if non-nil, zero value otherwise.

### GetIsReplaceMissingOk

`func (o *LogsArithmeticProcessor) GetIsReplaceMissingOk() (*bool, bool)`

GetIsReplaceMissingOk returns a tuple with the IsReplaceMissing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReplaceMissing

`func (o *LogsArithmeticProcessor) SetIsReplaceMissing(v bool)`

SetIsReplaceMissing sets IsReplaceMissing field to given value.

### HasIsReplaceMissing

`func (o *LogsArithmeticProcessor) HasIsReplaceMissing() bool`

HasIsReplaceMissing returns a boolean if a field has been set.

### GetName

`func (o *LogsArithmeticProcessor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LogsArithmeticProcessor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LogsArithmeticProcessor) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LogsArithmeticProcessor) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTarget

`func (o *LogsArithmeticProcessor) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *LogsArithmeticProcessor) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *LogsArithmeticProcessor) SetTarget(v string)`

SetTarget sets Target field to given value.


### GetType

`func (o *LogsArithmeticProcessor) GetType() LogsArithmeticProcessorType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LogsArithmeticProcessor) GetTypeOk() (*LogsArithmeticProcessorType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LogsArithmeticProcessor) SetType(v LogsArithmeticProcessorType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


