# LogsGrokParserRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MatchRules** | **string** | List of match rules for the grok parser, separated by a new line. | 
**SupportRules** | Pointer to **string** | List of support rules for the grok parser, separated by a new line. | [optional] [default to ""]

## Methods

### NewLogsGrokParserRules

`func NewLogsGrokParserRules(matchRules string, ) *LogsGrokParserRules`

NewLogsGrokParserRules instantiates a new LogsGrokParserRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogsGrokParserRulesWithDefaults

`func NewLogsGrokParserRulesWithDefaults() *LogsGrokParserRules`

NewLogsGrokParserRulesWithDefaults instantiates a new LogsGrokParserRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMatchRules

`func (o *LogsGrokParserRules) GetMatchRules() string`

GetMatchRules returns the MatchRules field if non-nil, zero value otherwise.

### GetMatchRulesOk

`func (o *LogsGrokParserRules) GetMatchRulesOk() (*string, bool)`

GetMatchRulesOk returns a tuple with the MatchRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchRules

`func (o *LogsGrokParserRules) SetMatchRules(v string)`

SetMatchRules sets MatchRules field to given value.


### GetSupportRules

`func (o *LogsGrokParserRules) GetSupportRules() string`

GetSupportRules returns the SupportRules field if non-nil, zero value otherwise.

### GetSupportRulesOk

`func (o *LogsGrokParserRules) GetSupportRulesOk() (*string, bool)`

GetSupportRulesOk returns a tuple with the SupportRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportRules

`func (o *LogsGrokParserRules) SetSupportRules(v string)`

SetSupportRules sets SupportRules field to given value.

### HasSupportRules

`func (o *LogsGrokParserRules) HasSupportRules() bool`

HasSupportRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


