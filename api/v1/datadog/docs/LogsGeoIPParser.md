# LogsGeoIPParser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sources** | Pointer to **[]string** | Array of source attributes. | [default to ["network.client.ip"]]
**Target** | Pointer to **string** | Name of the parent attribute that contains all the extracted details from the &#x60;sources&#x60;. | [default to "network.client.geoip"]
**Type** | Pointer to **string** | Type of processor. | [optional] [readonly] [default to "geo-ip-parser"]
**IsEnabled** | Pointer to **bool** | Whether or not the processor is enabled. | [optional] [default to false]
**Name** | Pointer to **string** | Name of the processor. | [optional] 

## Methods

### NewLogsGeoIPParser

`func NewLogsGeoIPParser(sources []string, target string, ) *LogsGeoIPParser`

NewLogsGeoIPParser instantiates a new LogsGeoIPParser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogsGeoIPParserWithDefaults

`func NewLogsGeoIPParserWithDefaults() *LogsGeoIPParser`

NewLogsGeoIPParserWithDefaults instantiates a new LogsGeoIPParser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSources

`func (o *LogsGeoIPParser) GetSources() []string`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *LogsGeoIPParser) GetSourcesOk() ([]string, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSources

`func (o *LogsGeoIPParser) HasSources() bool`

HasSources returns a boolean if a field has been set.

### SetSources

`func (o *LogsGeoIPParser) SetSources(v []string)`

SetSources gets a reference to the given []string and assigns it to the Sources field.

### GetTarget

`func (o *LogsGeoIPParser) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *LogsGeoIPParser) GetTargetOk() (string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTarget

`func (o *LogsGeoIPParser) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### SetTarget

`func (o *LogsGeoIPParser) SetTarget(v string)`

SetTarget gets a reference to the given string and assigns it to the Target field.

### GetType

`func (o *LogsGeoIPParser) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LogsGeoIPParser) GetTypeOk() (string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasType

`func (o *LogsGeoIPParser) HasType() bool`

HasType returns a boolean if a field has been set.

### SetType

`func (o *LogsGeoIPParser) SetType(v string)`

SetType gets a reference to the given string and assigns it to the Type field.

### GetIsEnabled

`func (o *LogsGeoIPParser) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *LogsGeoIPParser) GetIsEnabledOk() (bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsEnabled

`func (o *LogsGeoIPParser) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### SetIsEnabled

`func (o *LogsGeoIPParser) SetIsEnabled(v bool)`

SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.

### GetName

`func (o *LogsGeoIPParser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LogsGeoIPParser) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *LogsGeoIPParser) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *LogsGeoIPParser) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.


### AsLogsProcessor

`func (s *LogsGeoIPParser) AsLogsProcessor() LogsProcessor`

Convenience method to wrap this instance of LogsGeoIPParser in LogsProcessor

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


