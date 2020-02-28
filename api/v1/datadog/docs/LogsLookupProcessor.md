# LogsLookupProcessor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultLookup** | Pointer to **string** | Value to set the target attribute if the source value is not found in the list. | [optional] 
**LookupTable** | Pointer to **[]string** | Mapping table of values for the source attribute and their associated target attribute values, formatted as &#x60;[\&quot;source_key1,target_value1\&quot;, \&quot;source_key2,target_value2\&quot;]&#x60; | 
**Source** | Pointer to **string** | Source attribute used to perform the lookup. | 
**Target** | Pointer to **string** | Name of the attribute that contains the corresponding value in the mapping list or the &#x60;default_lookup&#x60; if not found in the mapping list. | 
**Type** | Pointer to **string** | Type of processor | [optional] [readonly] [default to "lookup-processor"]
**IsEnabled** | Pointer to **bool** | Whether or not the processor is enabled | [optional] [default to false]
**Name** | Pointer to **string** | Name of the processor | [optional] 

## Methods

### NewLogsLookupProcessor

`func NewLogsLookupProcessor(lookupTable []string, source string, target string, ) *LogsLookupProcessor`

NewLogsLookupProcessor instantiates a new LogsLookupProcessor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogsLookupProcessorWithDefaults

`func NewLogsLookupProcessorWithDefaults() *LogsLookupProcessor`

NewLogsLookupProcessorWithDefaults instantiates a new LogsLookupProcessor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultLookup

`func (o *LogsLookupProcessor) GetDefaultLookup() string`

GetDefaultLookup returns the DefaultLookup field if non-nil, zero value otherwise.

### GetDefaultLookupOk

`func (o *LogsLookupProcessor) GetDefaultLookupOk() (string, bool)`

GetDefaultLookupOk returns a tuple with the DefaultLookup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefaultLookup

`func (o *LogsLookupProcessor) HasDefaultLookup() bool`

HasDefaultLookup returns a boolean if a field has been set.

### SetDefaultLookup

`func (o *LogsLookupProcessor) SetDefaultLookup(v string)`

SetDefaultLookup gets a reference to the given string and assigns it to the DefaultLookup field.

### GetLookupTable

`func (o *LogsLookupProcessor) GetLookupTable() []string`

GetLookupTable returns the LookupTable field if non-nil, zero value otherwise.

### GetLookupTableOk

`func (o *LogsLookupProcessor) GetLookupTableOk() ([]string, bool)`

GetLookupTableOk returns a tuple with the LookupTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLookupTable

`func (o *LogsLookupProcessor) HasLookupTable() bool`

HasLookupTable returns a boolean if a field has been set.

### SetLookupTable

`func (o *LogsLookupProcessor) SetLookupTable(v []string)`

SetLookupTable gets a reference to the given []string and assigns it to the LookupTable field.

### GetSource

`func (o *LogsLookupProcessor) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *LogsLookupProcessor) GetSourceOk() (string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSource

`func (o *LogsLookupProcessor) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSource

`func (o *LogsLookupProcessor) SetSource(v string)`

SetSource gets a reference to the given string and assigns it to the Source field.

### GetTarget

`func (o *LogsLookupProcessor) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *LogsLookupProcessor) GetTargetOk() (string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTarget

`func (o *LogsLookupProcessor) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### SetTarget

`func (o *LogsLookupProcessor) SetTarget(v string)`

SetTarget gets a reference to the given string and assigns it to the Target field.

### GetType

`func (o *LogsLookupProcessor) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LogsLookupProcessor) GetTypeOk() (string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasType

`func (o *LogsLookupProcessor) HasType() bool`

HasType returns a boolean if a field has been set.

### SetType

`func (o *LogsLookupProcessor) SetType(v string)`

SetType gets a reference to the given string and assigns it to the Type field.

### GetIsEnabled

`func (o *LogsLookupProcessor) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *LogsLookupProcessor) GetIsEnabledOk() (bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsEnabled

`func (o *LogsLookupProcessor) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### SetIsEnabled

`func (o *LogsLookupProcessor) SetIsEnabled(v bool)`

SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.

### GetName

`func (o *LogsLookupProcessor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LogsLookupProcessor) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *LogsLookupProcessor) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *LogsLookupProcessor) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.


### AsLogsProcessor

`func (s *LogsLookupProcessor) AsLogsProcessor() LogsProcessor`

Convenience method to wrap this instance of LogsLookupProcessor in LogsProcessor

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


