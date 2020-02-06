# LogsStringBuilderProcessor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsReplaceMissing** | Pointer to **bool** | If true, it replaces all missing attributes of &#x60;template&#x60; by an empty string. If &#x60;false&#x60; (default), skips the operation for missing attributes. | [optional] [default to false]
**Target** | Pointer to **string** | The name of the attribute that contains the result of the template. | 
**Template** | Pointer to **string** | A formula with one or more attributes and raw text. | 
**Type** | Pointer to **string** | Type of processor | [optional] [readonly] [default to "string-builder-processor"]
**IsEnabled** | Pointer to **bool** | Whether or not the processor is enabled | [optional] [default to false]
**Name** | Pointer to **string** | Name of the processor | [optional] 

## Methods

### NewLogsStringBuilderProcessor

`func NewLogsStringBuilderProcessor(target string, template string, ) *LogsStringBuilderProcessor`

NewLogsStringBuilderProcessor instantiates a new LogsStringBuilderProcessor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogsStringBuilderProcessorWithDefaults

`func NewLogsStringBuilderProcessorWithDefaults() *LogsStringBuilderProcessor`

NewLogsStringBuilderProcessorWithDefaults instantiates a new LogsStringBuilderProcessor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsReplaceMissing

`func (o *LogsStringBuilderProcessor) GetIsReplaceMissing() bool`

GetIsReplaceMissing returns the IsReplaceMissing field if non-nil, zero value otherwise.

### GetIsReplaceMissingOk

`func (o *LogsStringBuilderProcessor) GetIsReplaceMissingOk() (bool, bool)`

GetIsReplaceMissingOk returns a tuple with the IsReplaceMissing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsReplaceMissing

`func (o *LogsStringBuilderProcessor) HasIsReplaceMissing() bool`

HasIsReplaceMissing returns a boolean if a field has been set.

### SetIsReplaceMissing

`func (o *LogsStringBuilderProcessor) SetIsReplaceMissing(v bool)`

SetIsReplaceMissing gets a reference to the given bool and assigns it to the IsReplaceMissing field.

### GetTarget

`func (o *LogsStringBuilderProcessor) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *LogsStringBuilderProcessor) GetTargetOk() (string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTarget

`func (o *LogsStringBuilderProcessor) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### SetTarget

`func (o *LogsStringBuilderProcessor) SetTarget(v string)`

SetTarget gets a reference to the given string and assigns it to the Target field.

### GetTemplate

`func (o *LogsStringBuilderProcessor) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *LogsStringBuilderProcessor) GetTemplateOk() (string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTemplate

`func (o *LogsStringBuilderProcessor) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### SetTemplate

`func (o *LogsStringBuilderProcessor) SetTemplate(v string)`

SetTemplate gets a reference to the given string and assigns it to the Template field.

### GetType

`func (o *LogsStringBuilderProcessor) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LogsStringBuilderProcessor) GetTypeOk() (string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasType

`func (o *LogsStringBuilderProcessor) HasType() bool`

HasType returns a boolean if a field has been set.

### SetType

`func (o *LogsStringBuilderProcessor) SetType(v string)`

SetType gets a reference to the given string and assigns it to the Type field.

### GetIsEnabled

`func (o *LogsStringBuilderProcessor) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *LogsStringBuilderProcessor) GetIsEnabledOk() (bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsEnabled

`func (o *LogsStringBuilderProcessor) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### SetIsEnabled

`func (o *LogsStringBuilderProcessor) SetIsEnabled(v bool)`

SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.

### GetName

`func (o *LogsStringBuilderProcessor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LogsStringBuilderProcessor) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *LogsStringBuilderProcessor) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *LogsStringBuilderProcessor) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.


### AsLogsProcessor

`func (s *LogsStringBuilderProcessor) AsLogsProcessor() LogsProcessor`

Convenience method to wrap this instance of LogsStringBuilderProcessor in LogsProcessor

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


