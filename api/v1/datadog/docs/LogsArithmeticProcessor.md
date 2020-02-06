# LogsArithmeticProcessor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expression** | Pointer to **string** | Arithmetic operation between one or more log attributes. | 
**IsReplaceMissing** | Pointer to **bool** | If &#x60;true&#x60;, it replaces all missing attributes of expression by &#x60;0&#x60;, &#x60;false&#x60; skip the operation if an attribute is missing. | [optional] [default to false]
**Target** | Pointer to **string** | Name of the attribute that contains the result of the arithmetic operation. | 
**Type** | Pointer to **string** | Type of processor | [optional] [readonly] [default to "arithmetic-processor"]
**IsEnabled** | Pointer to **bool** | Whether or not the processor is enabled | [optional] [default to false]
**Name** | Pointer to **string** | Name of the processor | [optional] 

## Methods

### NewLogsArithmeticProcessor

`func NewLogsArithmeticProcessor(expression string, target string, ) *LogsArithmeticProcessor`

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

`func (o *LogsArithmeticProcessor) GetExpressionOk() (string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExpression

`func (o *LogsArithmeticProcessor) HasExpression() bool`

HasExpression returns a boolean if a field has been set.

### SetExpression

`func (o *LogsArithmeticProcessor) SetExpression(v string)`

SetExpression gets a reference to the given string and assigns it to the Expression field.

### GetIsReplaceMissing

`func (o *LogsArithmeticProcessor) GetIsReplaceMissing() bool`

GetIsReplaceMissing returns the IsReplaceMissing field if non-nil, zero value otherwise.

### GetIsReplaceMissingOk

`func (o *LogsArithmeticProcessor) GetIsReplaceMissingOk() (bool, bool)`

GetIsReplaceMissingOk returns a tuple with the IsReplaceMissing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsReplaceMissing

`func (o *LogsArithmeticProcessor) HasIsReplaceMissing() bool`

HasIsReplaceMissing returns a boolean if a field has been set.

### SetIsReplaceMissing

`func (o *LogsArithmeticProcessor) SetIsReplaceMissing(v bool)`

SetIsReplaceMissing gets a reference to the given bool and assigns it to the IsReplaceMissing field.

### GetTarget

`func (o *LogsArithmeticProcessor) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *LogsArithmeticProcessor) GetTargetOk() (string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTarget

`func (o *LogsArithmeticProcessor) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### SetTarget

`func (o *LogsArithmeticProcessor) SetTarget(v string)`

SetTarget gets a reference to the given string and assigns it to the Target field.

### GetType

`func (o *LogsArithmeticProcessor) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LogsArithmeticProcessor) GetTypeOk() (string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasType

`func (o *LogsArithmeticProcessor) HasType() bool`

HasType returns a boolean if a field has been set.

### SetType

`func (o *LogsArithmeticProcessor) SetType(v string)`

SetType gets a reference to the given string and assigns it to the Type field.

### GetIsEnabled

`func (o *LogsArithmeticProcessor) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *LogsArithmeticProcessor) GetIsEnabledOk() (bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsEnabled

`func (o *LogsArithmeticProcessor) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### SetIsEnabled

`func (o *LogsArithmeticProcessor) SetIsEnabled(v bool)`

SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.

### GetName

`func (o *LogsArithmeticProcessor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LogsArithmeticProcessor) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *LogsArithmeticProcessor) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *LogsArithmeticProcessor) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.


### AsLogsProcessor

`func (s *LogsArithmeticProcessor) AsLogsProcessor() LogsProcessor`

Convenience method to wrap this instance of LogsArithmeticProcessor in LogsProcessor

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


