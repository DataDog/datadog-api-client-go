# LogsProcessor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Grok** | [**LogsGrokParserRules**](LogsGrokParserRules.md) |  | 
**IsEnabled** | Pointer to **bool** | Whether or not the processor is enabled. | [optional] [default to false]
**Name** | Pointer to **string** | Name of the processor. | [optional] 
**Samples** | Pointer to **[]string** | List of sample logs to test this grok parser. | [optional] 
**Source** | **string** | Source attribute used to perform the lookup. | 
**Type** | [**LogsTraceRemapperType**](LogsTraceRemapperType.md) |  | [default to "trace-id-remapper"]
**Sources** | **[]string** | Array of source attributes. | [default to ["dd.trace_id"]]
**OverrideOnConflict** | Pointer to **bool** | Override or not the target element if already set, | [optional] [default to false]
**PreserveSource** | Pointer to **bool** | Remove or preserve the remapped source element. | [optional] [default to false]
**SourceType** | Pointer to **string** | Defines if the sources are from log &#x60;attribute&#x60; or &#x60;tag&#x60;. | [optional] [default to "attribute"]
**Target** | **string** | Name of the attribute that contains the corresponding value in the mapping list or the &#x60;default_lookup&#x60; if not found in the mapping list. | 
**TargetType** | Pointer to **string** | Defines if the final attribute or tag name is from log &#x60;attribute&#x60; or &#x60;tag&#x60;. | [optional] [default to "attribute"]
**NormalizeEndingSlashes** | Pointer to **NullableBool** | Normalize the ending slashes or not. | [optional] [default to false]
**IsEncoded** | Pointer to **bool** | Define if the source attribute is URL encoded or not. | [optional] [default to false]
**Categories** | [**[]LogsCategoryProcessorCategories**](LogsCategoryProcessor_categories.md) | Array of filters to match or not a log and their corresponding &#x60;name&#x60;to assign a custom value to the log. | 
**Expression** | **string** | Arithmetic operation between one or more log attributes. | 
**IsReplaceMissing** | Pointer to **bool** | If true, it replaces all missing attributes of &#x60;template&#x60; by an empty string. If &#x60;false&#x60; (default), skips the operation for missing attributes. | [optional] [default to false]
**Template** | **string** | A formula with one or more attributes and raw text. | 
**Filter** | Pointer to [**LogsFilter**](LogsFilter.md) |  | [optional] 
**Processors** | Pointer to [**[]LogsProcessor**](LogsProcessor.md) | Ordered list of processors in this pipeline. | [optional] 
**DefaultLookup** | Pointer to **string** | Value to set the target attribute if the source value is not found in the list. | [optional] 
**LookupTable** | **[]string** | Mapping table of values for the source attribute and their associated target attribute values, formatted as &#x60;[\&quot;source_key1,target_value1\&quot;, \&quot;source_key2,target_value2\&quot;]&#x60; | 

## Methods

### NewLogsProcessor

`func NewLogsProcessor(grok LogsGrokParserRules, source string, type_ LogsTraceRemapperType, sources []string, target string, categories []LogsCategoryProcessorCategories, expression string, template string, lookupTable []string, ) *LogsProcessor`

NewLogsProcessor instantiates a new LogsProcessor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogsProcessorWithDefaults

`func NewLogsProcessorWithDefaults() *LogsProcessor`

NewLogsProcessorWithDefaults instantiates a new LogsProcessor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrok

`func (o *LogsProcessor) GetGrok() LogsGrokParserRules`

GetGrok returns the Grok field if non-nil, zero value otherwise.

### GetGrokOk

`func (o *LogsProcessor) GetGrokOk() (*LogsGrokParserRules, bool)`

GetGrokOk returns a tuple with the Grok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrok

`func (o *LogsProcessor) SetGrok(v LogsGrokParserRules)`

SetGrok sets Grok field to given value.


### GetIsEnabled

`func (o *LogsProcessor) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *LogsProcessor) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *LogsProcessor) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *LogsProcessor) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetName

`func (o *LogsProcessor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LogsProcessor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LogsProcessor) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LogsProcessor) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSamples

`func (o *LogsProcessor) GetSamples() []string`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *LogsProcessor) GetSamplesOk() (*[]string, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *LogsProcessor) SetSamples(v []string)`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *LogsProcessor) HasSamples() bool`

HasSamples returns a boolean if a field has been set.

### GetSource

`func (o *LogsProcessor) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *LogsProcessor) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *LogsProcessor) SetSource(v string)`

SetSource sets Source field to given value.


### GetType

`func (o *LogsProcessor) GetType() LogsTraceRemapperType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LogsProcessor) GetTypeOk() (*LogsTraceRemapperType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LogsProcessor) SetType(v LogsTraceRemapperType)`

SetType sets Type field to given value.


### GetSources

`func (o *LogsProcessor) GetSources() []string`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *LogsProcessor) GetSourcesOk() (*[]string, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *LogsProcessor) SetSources(v []string)`

SetSources sets Sources field to given value.


### GetOverrideOnConflict

`func (o *LogsProcessor) GetOverrideOnConflict() bool`

GetOverrideOnConflict returns the OverrideOnConflict field if non-nil, zero value otherwise.

### GetOverrideOnConflictOk

`func (o *LogsProcessor) GetOverrideOnConflictOk() (*bool, bool)`

GetOverrideOnConflictOk returns a tuple with the OverrideOnConflict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideOnConflict

`func (o *LogsProcessor) SetOverrideOnConflict(v bool)`

SetOverrideOnConflict sets OverrideOnConflict field to given value.

### HasOverrideOnConflict

`func (o *LogsProcessor) HasOverrideOnConflict() bool`

HasOverrideOnConflict returns a boolean if a field has been set.

### GetPreserveSource

`func (o *LogsProcessor) GetPreserveSource() bool`

GetPreserveSource returns the PreserveSource field if non-nil, zero value otherwise.

### GetPreserveSourceOk

`func (o *LogsProcessor) GetPreserveSourceOk() (*bool, bool)`

GetPreserveSourceOk returns a tuple with the PreserveSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveSource

`func (o *LogsProcessor) SetPreserveSource(v bool)`

SetPreserveSource sets PreserveSource field to given value.

### HasPreserveSource

`func (o *LogsProcessor) HasPreserveSource() bool`

HasPreserveSource returns a boolean if a field has been set.

### GetSourceType

`func (o *LogsProcessor) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *LogsProcessor) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *LogsProcessor) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *LogsProcessor) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### GetTarget

`func (o *LogsProcessor) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *LogsProcessor) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *LogsProcessor) SetTarget(v string)`

SetTarget sets Target field to given value.


### GetTargetType

`func (o *LogsProcessor) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *LogsProcessor) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *LogsProcessor) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *LogsProcessor) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.

### GetNormalizeEndingSlashes

`func (o *LogsProcessor) GetNormalizeEndingSlashes() bool`

GetNormalizeEndingSlashes returns the NormalizeEndingSlashes field if non-nil, zero value otherwise.

### GetNormalizeEndingSlashesOk

`func (o *LogsProcessor) GetNormalizeEndingSlashesOk() (*bool, bool)`

GetNormalizeEndingSlashesOk returns a tuple with the NormalizeEndingSlashes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizeEndingSlashes

`func (o *LogsProcessor) SetNormalizeEndingSlashes(v bool)`

SetNormalizeEndingSlashes sets NormalizeEndingSlashes field to given value.

### HasNormalizeEndingSlashes

`func (o *LogsProcessor) HasNormalizeEndingSlashes() bool`

HasNormalizeEndingSlashes returns a boolean if a field has been set.

### SetNormalizeEndingSlashesNil

`func (o *LogsProcessor) SetNormalizeEndingSlashesNil(b bool)`

 SetNormalizeEndingSlashesNil sets the value for NormalizeEndingSlashes to be an explicit nil

### UnsetNormalizeEndingSlashes
`func (o *LogsProcessor) UnsetNormalizeEndingSlashes()`

UnsetNormalizeEndingSlashes ensures that no value is present for NormalizeEndingSlashes, not even an explicit nil
### GetIsEncoded

`func (o *LogsProcessor) GetIsEncoded() bool`

GetIsEncoded returns the IsEncoded field if non-nil, zero value otherwise.

### GetIsEncodedOk

`func (o *LogsProcessor) GetIsEncodedOk() (*bool, bool)`

GetIsEncodedOk returns a tuple with the IsEncoded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEncoded

`func (o *LogsProcessor) SetIsEncoded(v bool)`

SetIsEncoded sets IsEncoded field to given value.

### HasIsEncoded

`func (o *LogsProcessor) HasIsEncoded() bool`

HasIsEncoded returns a boolean if a field has been set.

### GetCategories

`func (o *LogsProcessor) GetCategories() []LogsCategoryProcessorCategories`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *LogsProcessor) GetCategoriesOk() (*[]LogsCategoryProcessorCategories, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *LogsProcessor) SetCategories(v []LogsCategoryProcessorCategories)`

SetCategories sets Categories field to given value.


### GetExpression

`func (o *LogsProcessor) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *LogsProcessor) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *LogsProcessor) SetExpression(v string)`

SetExpression sets Expression field to given value.


### GetIsReplaceMissing

`func (o *LogsProcessor) GetIsReplaceMissing() bool`

GetIsReplaceMissing returns the IsReplaceMissing field if non-nil, zero value otherwise.

### GetIsReplaceMissingOk

`func (o *LogsProcessor) GetIsReplaceMissingOk() (*bool, bool)`

GetIsReplaceMissingOk returns a tuple with the IsReplaceMissing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReplaceMissing

`func (o *LogsProcessor) SetIsReplaceMissing(v bool)`

SetIsReplaceMissing sets IsReplaceMissing field to given value.

### HasIsReplaceMissing

`func (o *LogsProcessor) HasIsReplaceMissing() bool`

HasIsReplaceMissing returns a boolean if a field has been set.

### GetTemplate

`func (o *LogsProcessor) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *LogsProcessor) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *LogsProcessor) SetTemplate(v string)`

SetTemplate sets Template field to given value.


### GetFilter

`func (o *LogsProcessor) GetFilter() LogsFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *LogsProcessor) GetFilterOk() (*LogsFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *LogsProcessor) SetFilter(v LogsFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *LogsProcessor) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetProcessors

`func (o *LogsProcessor) GetProcessors() []LogsProcessor`

GetProcessors returns the Processors field if non-nil, zero value otherwise.

### GetProcessorsOk

`func (o *LogsProcessor) GetProcessorsOk() (*[]LogsProcessor, bool)`

GetProcessorsOk returns a tuple with the Processors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessors

`func (o *LogsProcessor) SetProcessors(v []LogsProcessor)`

SetProcessors sets Processors field to given value.

### HasProcessors

`func (o *LogsProcessor) HasProcessors() bool`

HasProcessors returns a boolean if a field has been set.

### GetDefaultLookup

`func (o *LogsProcessor) GetDefaultLookup() string`

GetDefaultLookup returns the DefaultLookup field if non-nil, zero value otherwise.

### GetDefaultLookupOk

`func (o *LogsProcessor) GetDefaultLookupOk() (*string, bool)`

GetDefaultLookupOk returns a tuple with the DefaultLookup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultLookup

`func (o *LogsProcessor) SetDefaultLookup(v string)`

SetDefaultLookup sets DefaultLookup field to given value.

### HasDefaultLookup

`func (o *LogsProcessor) HasDefaultLookup() bool`

HasDefaultLookup returns a boolean if a field has been set.

### GetLookupTable

`func (o *LogsProcessor) GetLookupTable() []string`

GetLookupTable returns the LookupTable field if non-nil, zero value otherwise.

### GetLookupTableOk

`func (o *LogsProcessor) GetLookupTableOk() (*[]string, bool)`

GetLookupTableOk returns a tuple with the LookupTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookupTable

`func (o *LogsProcessor) SetLookupTable(v []string)`

SetLookupTable sets LookupTable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


