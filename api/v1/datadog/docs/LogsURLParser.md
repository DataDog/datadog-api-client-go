# LogsURLParser

## Properties

| Name                       | Type                                          | Description                                                                                        | Notes                                     |
| -------------------------- | --------------------------------------------- | -------------------------------------------------------------------------------------------------- | ----------------------------------------- |
| **IsEnabled**              | Pointer to **bool**                           | Whether or not the processor is enabled.                                                           | [optional] [default to false]             |
| **Name**                   | Pointer to **string**                         | Name of the processor.                                                                             | [optional]                                |
| **NormalizeEndingSlashes** | Pointer to **NullableBool**                   | Normalize the ending slashes or not.                                                               | [optional] [default to false]             |
| **Sources**                | **[]string**                                  | Array of source attributes.                                                                        | [default to ["http.url"]]                 |
| **Target**                 | **string**                                    | Name of the parent attribute that contains all the extracted details from the &#x60;sources&#x60;. | [default to "http.url_details"]           |
| **Type**                   | [**LogsURLParserType**](LogsURLParserType.md) |                                                                                                    | [default to LOGSURLPARSERTYPE_URL_PARSER] |

## Methods

### NewLogsURLParser

`func NewLogsURLParser(sources []string, target string, type_ LogsURLParserType) *LogsURLParser`

NewLogsURLParser instantiates a new LogsURLParser object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewLogsURLParserWithDefaults

`func NewLogsURLParserWithDefaults() *LogsURLParser`

NewLogsURLParserWithDefaults instantiates a new LogsURLParser object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetIsEnabled

`func (o *LogsURLParser) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *LogsURLParser) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *LogsURLParser) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *LogsURLParser) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetName

`func (o *LogsURLParser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LogsURLParser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LogsURLParser) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LogsURLParser) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNormalizeEndingSlashes

`func (o *LogsURLParser) GetNormalizeEndingSlashes() bool`

GetNormalizeEndingSlashes returns the NormalizeEndingSlashes field if non-nil, zero value otherwise.

### GetNormalizeEndingSlashesOk

`func (o *LogsURLParser) GetNormalizeEndingSlashesOk() (*bool, bool)`

GetNormalizeEndingSlashesOk returns a tuple with the NormalizeEndingSlashes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizeEndingSlashes

`func (o *LogsURLParser) SetNormalizeEndingSlashes(v bool)`

SetNormalizeEndingSlashes sets NormalizeEndingSlashes field to given value.

### HasNormalizeEndingSlashes

`func (o *LogsURLParser) HasNormalizeEndingSlashes() bool`

HasNormalizeEndingSlashes returns a boolean if a field has been set.

### SetNormalizeEndingSlashesNil

`func (o *LogsURLParser) SetNormalizeEndingSlashesNil(b bool)`

SetNormalizeEndingSlashesNil sets the value for NormalizeEndingSlashes to be an explicit nil

### UnsetNormalizeEndingSlashes

`func (o *LogsURLParser) UnsetNormalizeEndingSlashes()`

UnsetNormalizeEndingSlashes ensures that no value is present for NormalizeEndingSlashes, not even an explicit nil

### GetSources

`func (o *LogsURLParser) GetSources() []string`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *LogsURLParser) GetSourcesOk() (*[]string, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *LogsURLParser) SetSources(v []string)`

SetSources sets Sources field to given value.

### GetTarget

`func (o *LogsURLParser) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *LogsURLParser) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *LogsURLParser) SetTarget(v string)`

SetTarget sets Target field to given value.

### GetType

`func (o *LogsURLParser) GetType() LogsURLParserType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LogsURLParser) GetTypeOk() (*LogsURLParserType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LogsURLParser) SetType(v LogsURLParserType)`

SetType sets Type field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
