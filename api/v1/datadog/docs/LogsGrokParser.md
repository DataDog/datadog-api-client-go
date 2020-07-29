# LogsGrokParser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Grok** | [**LogsGrokParserRules**](LogsGrokParserRules.md) |  | 
**IsEnabled** | Pointer to **bool** | Whether or not the processor is enabled. | [optional] [default to false]
**Name** | Pointer to **string** | Name of the processor. | [optional] 
**Samples** | Pointer to **[]string** | List of sample logs to test this grok parser. | [optional] 
**Source** | **string** | Name of the log attribute to parse. | [default to "message"]
**Type** | [**LogsGrokParserType**](LogsGrokParserType.md) |  | [default to "grok-parser"]

## Methods

### NewLogsGrokParser

`func NewLogsGrokParser(grok LogsGrokParserRules, source string, type_ LogsGrokParserType, ) *LogsGrokParser`

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

`func (o *LogsGrokParser) GetGrokOk() (*LogsGrokParserRules, bool)`

GetGrokOk returns a tuple with the Grok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrok

`func (o *LogsGrokParser) SetGrok(v LogsGrokParserRules)`

SetGrok sets Grok field to given value.


### GetIsEnabled

`func (o *LogsGrokParser) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *LogsGrokParser) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *LogsGrokParser) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *LogsGrokParser) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetName

`func (o *LogsGrokParser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LogsGrokParser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LogsGrokParser) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LogsGrokParser) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSamples

`func (o *LogsGrokParser) GetSamples() []string`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *LogsGrokParser) GetSamplesOk() (*[]string, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *LogsGrokParser) SetSamples(v []string)`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *LogsGrokParser) HasSamples() bool`

HasSamples returns a boolean if a field has been set.

### GetSource

`func (o *LogsGrokParser) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *LogsGrokParser) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *LogsGrokParser) SetSource(v string)`

SetSource sets Source field to given value.


### GetType

`func (o *LogsGrokParser) GetType() LogsGrokParserType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LogsGrokParser) GetTypeOk() (*LogsGrokParserType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LogsGrokParser) SetType(v LogsGrokParserType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


