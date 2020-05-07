# LogsGeoIPParser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sources** | Pointer to **[]string** | Array of source attributes. | [default to ["network.client.ip"]]
**Target** | Pointer to **string** | Name of the parent attribute that contains all the extracted details from the &#x60;sources&#x60;. | [default to "network.client.geoip"]
**Type** | Pointer to **string** | Type of processor. | [readonly] [default to "geo-ip-parser"]
**IsEnabled** | Pointer to **bool** | Whether or not the processor is enabled. | [optional] [default to false]
**Name** | Pointer to **string** | Name of the processor. | [optional] 

## Methods

### NewLogsGeoIPParser

`func NewLogsGeoIPParser(sources []string, target string, type_ string, ) *LogsGeoIPParser`

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

`func (o *LogsGeoIPParser) GetSourcesOk() (*[]string, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *LogsGeoIPParser) SetSources(v []string)`

SetSources sets Sources field to given value.


### GetTarget

`func (o *LogsGeoIPParser) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *LogsGeoIPParser) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *LogsGeoIPParser) SetTarget(v string)`

SetTarget sets Target field to given value.


### GetType

`func (o *LogsGeoIPParser) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LogsGeoIPParser) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LogsGeoIPParser) SetType(v string)`

SetType sets Type field to given value.


### GetIsEnabled

`func (o *LogsGeoIPParser) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *LogsGeoIPParser) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *LogsGeoIPParser) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *LogsGeoIPParser) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetName

`func (o *LogsGeoIPParser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LogsGeoIPParser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LogsGeoIPParser) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LogsGeoIPParser) HasName() bool`

HasName returns a boolean if a field has been set.


### AsLogsProcessor

`func (s *LogsGeoIPParser) AsLogsProcessor() LogsProcessor`

Convenience method to wrap this instance of LogsGeoIPParser in LogsProcessor

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


