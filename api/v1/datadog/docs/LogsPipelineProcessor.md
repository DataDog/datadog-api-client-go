# LogsPipelineProcessor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**LogsFilter**](LogsFilter.md) |  | [optional] 
**Processors** | Pointer to [**[]LogsProcessor**](LogsProcessor.md) | Ordered list of processors in this pipeline. | [optional] 
**Type** | Pointer to **string** | Type of processor. | [readonly] [default to "pipeline"]
**IsEnabled** | Pointer to **bool** | Whether or not the processor is enabled. | [optional] [default to false]
**Name** | Pointer to **string** | Name of the processor. | [optional] 

## Methods

### NewLogsPipelineProcessor

`func NewLogsPipelineProcessor(type_ string, ) *LogsPipelineProcessor`

NewLogsPipelineProcessor instantiates a new LogsPipelineProcessor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogsPipelineProcessorWithDefaults

`func NewLogsPipelineProcessorWithDefaults() *LogsPipelineProcessor`

NewLogsPipelineProcessorWithDefaults instantiates a new LogsPipelineProcessor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *LogsPipelineProcessor) GetFilter() LogsFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *LogsPipelineProcessor) GetFilterOk() (*LogsFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *LogsPipelineProcessor) SetFilter(v LogsFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *LogsPipelineProcessor) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetProcessors

`func (o *LogsPipelineProcessor) GetProcessors() []LogsProcessor`

GetProcessors returns the Processors field if non-nil, zero value otherwise.

### GetProcessorsOk

`func (o *LogsPipelineProcessor) GetProcessorsOk() (*[]LogsProcessor, bool)`

GetProcessorsOk returns a tuple with the Processors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessors

`func (o *LogsPipelineProcessor) SetProcessors(v []LogsProcessor)`

SetProcessors sets Processors field to given value.

### HasProcessors

`func (o *LogsPipelineProcessor) HasProcessors() bool`

HasProcessors returns a boolean if a field has been set.

### GetType

`func (o *LogsPipelineProcessor) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LogsPipelineProcessor) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LogsPipelineProcessor) SetType(v string)`

SetType sets Type field to given value.


### GetIsEnabled

`func (o *LogsPipelineProcessor) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *LogsPipelineProcessor) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *LogsPipelineProcessor) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *LogsPipelineProcessor) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetName

`func (o *LogsPipelineProcessor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LogsPipelineProcessor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LogsPipelineProcessor) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LogsPipelineProcessor) HasName() bool`

HasName returns a boolean if a field has been set.


### AsLogsProcessor

`func (s *LogsPipelineProcessor) AsLogsProcessor() LogsProcessor`

Convenience method to wrap this instance of LogsPipelineProcessor in LogsProcessor

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


