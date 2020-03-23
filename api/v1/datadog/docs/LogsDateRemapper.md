# LogsDateRemapper

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sources** | Pointer to **[]string** | Array of source attributes. | 
**Type** | Pointer to **string** | Type of processor. | [optional] [readonly] [default to "date-remapper"]
**IsEnabled** | Pointer to **bool** | Whether or not the processor is enabled. | [optional] [default to false]
**Name** | Pointer to **string** | Name of the processor. | [optional] 

## Methods

### NewLogsDateRemapper

`func NewLogsDateRemapper(sources []string, ) *LogsDateRemapper`

NewLogsDateRemapper instantiates a new LogsDateRemapper object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogsDateRemapperWithDefaults

`func NewLogsDateRemapperWithDefaults() *LogsDateRemapper`

NewLogsDateRemapperWithDefaults instantiates a new LogsDateRemapper object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSources

`func (o *LogsDateRemapper) GetSources() []string`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *LogsDateRemapper) GetSourcesOk() ([]string, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSources

`func (o *LogsDateRemapper) HasSources() bool`

HasSources returns a boolean if a field has been set.

### SetSources

`func (o *LogsDateRemapper) SetSources(v []string)`

SetSources gets a reference to the given []string and assigns it to the Sources field.

### GetType

`func (o *LogsDateRemapper) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LogsDateRemapper) GetTypeOk() (string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasType

`func (o *LogsDateRemapper) HasType() bool`

HasType returns a boolean if a field has been set.

### SetType

`func (o *LogsDateRemapper) SetType(v string)`

SetType gets a reference to the given string and assigns it to the Type field.

### GetIsEnabled

`func (o *LogsDateRemapper) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *LogsDateRemapper) GetIsEnabledOk() (bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsEnabled

`func (o *LogsDateRemapper) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### SetIsEnabled

`func (o *LogsDateRemapper) SetIsEnabled(v bool)`

SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.

### GetName

`func (o *LogsDateRemapper) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LogsDateRemapper) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *LogsDateRemapper) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *LogsDateRemapper) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.


### AsLogsProcessor

`func (s *LogsDateRemapper) AsLogsProcessor() LogsProcessor`

Convenience method to wrap this instance of LogsDateRemapper in LogsProcessor

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


