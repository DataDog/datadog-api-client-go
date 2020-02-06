# LogsGrokParser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Grok** | Pointer to [**LogsGrokParserRules**](LogsGrokParserRules.md) |  | 
**Samples** | Pointer to **[]string** | List of sample logs to test this grok parser | [optional] 
**Source** | Pointer to **string** | Name of the log attribute to parse | [default to "message"]
**Type** | Pointer to **string** | Type of processor | [optional] [readonly] [default to "grok-parser"]
**IsEnabled** | Pointer to **bool** | Whether or not the processor is enabled | [optional] [default to false]
**Name** | Pointer to **string** | Name of the processor | [optional] 

## Methods

### NewLogsGrokParser

`func NewLogsGrokParser(grok LogsGrokParserRules, source string, ) *LogsGrokParser`

NewLogsGrokParser instantiates a new LogsGrokParser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogsGrokParserWithDefaults

`func NewLogsGrokParserWithDefaults() *LogsGrokParser`

NewLogsGrokParserWithDefaults instantiates a new LogsGrokParser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrok

`func (o *LogsGrokParser) GetGrok() LogsGrokParserRules`

GetGrok returns the Grok field if non-nil, zero value otherwise.

### GetGrokOk

`func (o *LogsGrokParser) GetGrokOk() (LogsGrokParserRules, bool)`

GetGrokOk returns a tuple with the Grok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasGrok

`func (o *LogsGrokParser) HasGrok() bool`

HasGrok returns a boolean if a field has been set.

### SetGrok

`func (o *LogsGrokParser) SetGrok(v LogsGrokParserRules)`

SetGrok gets a reference to the given LogsGrokParserRules and assigns it to the Grok field.

### GetSamples

`func (o *LogsGrokParser) GetSamples() []string`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *LogsGrokParser) GetSamplesOk() ([]string, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSamples

`func (o *LogsGrokParser) HasSamples() bool`

HasSamples returns a boolean if a field has been set.

### SetSamples

`func (o *LogsGrokParser) SetSamples(v []string)`

SetSamples gets a reference to the given []string and assigns it to the Samples field.

### GetSource

`func (o *LogsGrokParser) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *LogsGrokParser) GetSourceOk() (string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSource

`func (o *LogsGrokParser) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSource

`func (o *LogsGrokParser) SetSource(v string)`

SetSource gets a reference to the given string and assigns it to the Source field.

### GetType

`func (o *LogsGrokParser) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LogsGrokParser) GetTypeOk() (string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasType

`func (o *LogsGrokParser) HasType() bool`

HasType returns a boolean if a field has been set.

### SetType

`func (o *LogsGrokParser) SetType(v string)`

SetType gets a reference to the given string and assigns it to the Type field.

### GetIsEnabled

`func (o *LogsGrokParser) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *LogsGrokParser) GetIsEnabledOk() (bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsEnabled

`func (o *LogsGrokParser) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### SetIsEnabled

`func (o *LogsGrokParser) SetIsEnabled(v bool)`

SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.

### GetName

`func (o *LogsGrokParser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LogsGrokParser) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *LogsGrokParser) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *LogsGrokParser) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.


### AsLogsProcessor

`func (s *LogsGrokParser) AsLogsProcessor() LogsProcessor`

Convenience method to wrap this instance of LogsGrokParser in LogsProcessor

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


