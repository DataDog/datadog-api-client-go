# LogsCategoryProcessor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Categories** | Pointer to [**[]LogsCategoryProcessorCategories**](LogsCategoryProcessor_categories.md) | Array of filters to match or not a log and their corresponding &#x60;name&#x60;to assign a custom value to the log. | 
**Target** | Pointer to **string** | Name of the target attribute which value is defined by the matching category. | 
**Type** | Pointer to **string** | Type of processor. | [optional] [readonly] [default to "category-processor"]
**IsEnabled** | Pointer to **bool** | Whether or not the processor is enabled. | [optional] [default to false]
**Name** | Pointer to **string** | Name of the processor. | [optional] 

## Methods

### NewLogsCategoryProcessor

`func NewLogsCategoryProcessor(categories []LogsCategoryProcessorCategories, target string, ) *LogsCategoryProcessor`

NewLogsCategoryProcessor instantiates a new LogsCategoryProcessor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogsCategoryProcessorWithDefaults

`func NewLogsCategoryProcessorWithDefaults() *LogsCategoryProcessor`

NewLogsCategoryProcessorWithDefaults instantiates a new LogsCategoryProcessor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategories

`func (o *LogsCategoryProcessor) GetCategories() []LogsCategoryProcessorCategories`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *LogsCategoryProcessor) GetCategoriesOk() (*[]LogsCategoryProcessorCategories, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *LogsCategoryProcessor) SetCategories(v []LogsCategoryProcessorCategories)`

SetCategories sets Categories field to given value.


### GetTarget

`func (o *LogsCategoryProcessor) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *LogsCategoryProcessor) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *LogsCategoryProcessor) SetTarget(v string)`

SetTarget sets Target field to given value.


### GetType

`func (o *LogsCategoryProcessor) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LogsCategoryProcessor) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LogsCategoryProcessor) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *LogsCategoryProcessor) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIsEnabled

`func (o *LogsCategoryProcessor) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *LogsCategoryProcessor) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *LogsCategoryProcessor) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *LogsCategoryProcessor) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetName

`func (o *LogsCategoryProcessor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LogsCategoryProcessor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LogsCategoryProcessor) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LogsCategoryProcessor) HasName() bool`

HasName returns a boolean if a field has been set.


### AsLogsProcessor

`func (s *LogsCategoryProcessor) AsLogsProcessor() LogsProcessor`

Convenience method to wrap this instance of LogsCategoryProcessor in LogsProcessor

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


