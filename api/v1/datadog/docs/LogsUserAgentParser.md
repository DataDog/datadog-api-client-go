# LogsUserAgentParser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEncoded** | Pointer to **bool** | Define if the source attribute is url encoded or not. | [optional] [default to false]
**Sources** | Pointer to **[]string** | Array of source attributes. | [default to ["http.useragent"]]
**Target** | Pointer to **string** | Name of the parent attribute that contains all the extracted details from the &#x60;sources&#x60;. | [default to "http.useragent_details"]
**Type** | Pointer to **string** | Type of processor. | [optional] [readonly] [default to "user-agent-parser"]
**IsEnabled** | Pointer to **bool** | Whether or not the processor is enabled. | [optional] [default to false]
**Name** | Pointer to **string** | Name of the processor. | [optional] 

## Methods

### NewLogsUserAgentParser

`func NewLogsUserAgentParser(sources []string, target string, ) *LogsUserAgentParser`

NewLogsUserAgentParser instantiates a new LogsUserAgentParser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogsUserAgentParserWithDefaults

`func NewLogsUserAgentParserWithDefaults() *LogsUserAgentParser`

NewLogsUserAgentParserWithDefaults instantiates a new LogsUserAgentParser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEncoded

`func (o *LogsUserAgentParser) GetIsEncoded() bool`

GetIsEncoded returns the IsEncoded field if non-nil, zero value otherwise.

### GetIsEncodedOk

`func (o *LogsUserAgentParser) GetIsEncodedOk() (bool, bool)`

GetIsEncodedOk returns a tuple with the IsEncoded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsEncoded

`func (o *LogsUserAgentParser) HasIsEncoded() bool`

HasIsEncoded returns a boolean if a field has been set.

### SetIsEncoded

`func (o *LogsUserAgentParser) SetIsEncoded(v bool)`

SetIsEncoded gets a reference to the given bool and assigns it to the IsEncoded field.

### GetSources

`func (o *LogsUserAgentParser) GetSources() []string`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *LogsUserAgentParser) GetSourcesOk() ([]string, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSources

`func (o *LogsUserAgentParser) HasSources() bool`

HasSources returns a boolean if a field has been set.

### SetSources

`func (o *LogsUserAgentParser) SetSources(v []string)`

SetSources gets a reference to the given []string and assigns it to the Sources field.

### GetTarget

`func (o *LogsUserAgentParser) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *LogsUserAgentParser) GetTargetOk() (string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTarget

`func (o *LogsUserAgentParser) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### SetTarget

`func (o *LogsUserAgentParser) SetTarget(v string)`

SetTarget gets a reference to the given string and assigns it to the Target field.

### GetType

`func (o *LogsUserAgentParser) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LogsUserAgentParser) GetTypeOk() (string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasType

`func (o *LogsUserAgentParser) HasType() bool`

HasType returns a boolean if a field has been set.

### SetType

`func (o *LogsUserAgentParser) SetType(v string)`

SetType gets a reference to the given string and assigns it to the Type field.

### GetIsEnabled

`func (o *LogsUserAgentParser) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *LogsUserAgentParser) GetIsEnabledOk() (bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsEnabled

`func (o *LogsUserAgentParser) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### SetIsEnabled

`func (o *LogsUserAgentParser) SetIsEnabled(v bool)`

SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.

### GetName

`func (o *LogsUserAgentParser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LogsUserAgentParser) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *LogsUserAgentParser) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *LogsUserAgentParser) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.


### AsLogsProcessor

`func (s *LogsUserAgentParser) AsLogsProcessor() LogsProcessor`

Convenience method to wrap this instance of LogsUserAgentParser in LogsProcessor

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


