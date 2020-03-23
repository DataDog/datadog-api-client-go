# LogsRemapper

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OverrideOnConflict** | Pointer to **bool** | Override or not the target element if already set, | [optional] [default to false]
**PreserveSource** | Pointer to **bool** | Remove or preserve the remapped source element. | [optional] [default to false]
**SourceType** | Pointer to **string** | Defines if the sources are from log &#x60;attribute&#x60; or &#x60;tag&#x60;. | [optional] [default to "attribute"]
**Sources** | Pointer to **[]string** | Array of source attributes. | 
**Target** | Pointer to **string** | Final attribute or tag name to remap the sources to. | 
**TargetType** | Pointer to **string** | Defines if the sources are from log &#x60;attribute&#x60; or &#x60;tag&#x60;. | [optional] [default to "attribute"]
**Type** | Pointer to **string** | Type of processor. | [optional] [readonly] [default to "attribute-remapper"]
**IsEnabled** | Pointer to **bool** | Whether or not the processor is enabled. | [optional] [default to false]
**Name** | Pointer to **string** | Name of the processor. | [optional] 

## Methods

### NewLogsRemapper

`func NewLogsRemapper(sources []string, target string, ) *LogsRemapper`

NewLogsRemapper instantiates a new LogsRemapper object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogsRemapperWithDefaults

`func NewLogsRemapperWithDefaults() *LogsRemapper`

NewLogsRemapperWithDefaults instantiates a new LogsRemapper object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOverrideOnConflict

`func (o *LogsRemapper) GetOverrideOnConflict() bool`

GetOverrideOnConflict returns the OverrideOnConflict field if non-nil, zero value otherwise.

### GetOverrideOnConflictOk

`func (o *LogsRemapper) GetOverrideOnConflictOk() (bool, bool)`

GetOverrideOnConflictOk returns a tuple with the OverrideOnConflict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOverrideOnConflict

`func (o *LogsRemapper) HasOverrideOnConflict() bool`

HasOverrideOnConflict returns a boolean if a field has been set.

### SetOverrideOnConflict

`func (o *LogsRemapper) SetOverrideOnConflict(v bool)`

SetOverrideOnConflict gets a reference to the given bool and assigns it to the OverrideOnConflict field.

### GetPreserveSource

`func (o *LogsRemapper) GetPreserveSource() bool`

GetPreserveSource returns the PreserveSource field if non-nil, zero value otherwise.

### GetPreserveSourceOk

`func (o *LogsRemapper) GetPreserveSourceOk() (bool, bool)`

GetPreserveSourceOk returns a tuple with the PreserveSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPreserveSource

`func (o *LogsRemapper) HasPreserveSource() bool`

HasPreserveSource returns a boolean if a field has been set.

### SetPreserveSource

`func (o *LogsRemapper) SetPreserveSource(v bool)`

SetPreserveSource gets a reference to the given bool and assigns it to the PreserveSource field.

### GetSourceType

`func (o *LogsRemapper) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *LogsRemapper) GetSourceTypeOk() (string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSourceType

`func (o *LogsRemapper) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### SetSourceType

`func (o *LogsRemapper) SetSourceType(v string)`

SetSourceType gets a reference to the given string and assigns it to the SourceType field.

### GetSources

`func (o *LogsRemapper) GetSources() []string`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *LogsRemapper) GetSourcesOk() ([]string, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSources

`func (o *LogsRemapper) HasSources() bool`

HasSources returns a boolean if a field has been set.

### SetSources

`func (o *LogsRemapper) SetSources(v []string)`

SetSources gets a reference to the given []string and assigns it to the Sources field.

### GetTarget

`func (o *LogsRemapper) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *LogsRemapper) GetTargetOk() (string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTarget

`func (o *LogsRemapper) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### SetTarget

`func (o *LogsRemapper) SetTarget(v string)`

SetTarget gets a reference to the given string and assigns it to the Target field.

### GetTargetType

`func (o *LogsRemapper) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *LogsRemapper) GetTargetTypeOk() (string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTargetType

`func (o *LogsRemapper) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.

### SetTargetType

`func (o *LogsRemapper) SetTargetType(v string)`

SetTargetType gets a reference to the given string and assigns it to the TargetType field.

### GetType

`func (o *LogsRemapper) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LogsRemapper) GetTypeOk() (string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasType

`func (o *LogsRemapper) HasType() bool`

HasType returns a boolean if a field has been set.

### SetType

`func (o *LogsRemapper) SetType(v string)`

SetType gets a reference to the given string and assigns it to the Type field.

### GetIsEnabled

`func (o *LogsRemapper) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *LogsRemapper) GetIsEnabledOk() (bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsEnabled

`func (o *LogsRemapper) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### SetIsEnabled

`func (o *LogsRemapper) SetIsEnabled(v bool)`

SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.

### GetName

`func (o *LogsRemapper) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LogsRemapper) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *LogsRemapper) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *LogsRemapper) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.


### AsLogsProcessor

`func (s *LogsRemapper) AsLogsProcessor() LogsProcessor`

Convenience method to wrap this instance of LogsRemapper in LogsProcessor

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


