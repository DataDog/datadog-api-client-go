# LogsAttributeRemapper

## Properties

| Name                   | Type                                                          | Description                                                                                      | Notes                                                     |
| ---------------------- | ------------------------------------------------------------- | ------------------------------------------------------------------------------------------------ | --------------------------------------------------------- |
| **IsEnabled**          | Pointer to **bool**                                           | Whether or not the processor is enabled.                                                         | [optional] [default to false]                             |
| **Name**               | Pointer to **string**                                         | Name of the processor.                                                                           | [optional]                                                |
| **OverrideOnConflict** | Pointer to **bool**                                           | Override or not the target element if already set,                                               | [optional] [default to false]                             |
| **PreserveSource**     | Pointer to **bool**                                           | Remove or preserve the remapped source element.                                                  | [optional] [default to false]                             |
| **SourceType**         | Pointer to **string**                                         | Defines if the sources are from log &#x60;attribute&#x60; or &#x60;tag&#x60;.                    | [optional] [default to "attribute"]                       |
| **Sources**            | **[]string**                                                  | Array of source attributes.                                                                      |
| **Target**             | **string**                                                    | Final attribute or tag name to remap the sources to.                                             |
| **TargetFormat**       | Pointer to [**TargetFormatType**](TargetFormatType.md)        |                                                                                                  | [optional]                                                |
| **TargetType**         | Pointer to **string**                                         | Defines if the final attribute or tag name is from log &#x60;attribute&#x60; or &#x60;tag&#x60;. | [optional] [default to "attribute"]                       |
| **Type**               | [**LogsAttributeRemapperType**](LogsAttributeRemapperType.md) |                                                                                                  | [default to LOGSATTRIBUTEREMAPPERTYPE_ATTRIBUTE_REMAPPER] |

## Methods

### NewLogsAttributeRemapper

`func NewLogsAttributeRemapper(sources []string, target string, type_ LogsAttributeRemapperType) *LogsAttributeRemapper`

NewLogsAttributeRemapper instantiates a new LogsAttributeRemapper object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewLogsAttributeRemapperWithDefaults

`func NewLogsAttributeRemapperWithDefaults() *LogsAttributeRemapper`

NewLogsAttributeRemapperWithDefaults instantiates a new LogsAttributeRemapper object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetIsEnabled

`func (o *LogsAttributeRemapper) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *LogsAttributeRemapper) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *LogsAttributeRemapper) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *LogsAttributeRemapper) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetName

`func (o *LogsAttributeRemapper) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LogsAttributeRemapper) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LogsAttributeRemapper) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LogsAttributeRemapper) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOverrideOnConflict

`func (o *LogsAttributeRemapper) GetOverrideOnConflict() bool`

GetOverrideOnConflict returns the OverrideOnConflict field if non-nil, zero value otherwise.

### GetOverrideOnConflictOk

`func (o *LogsAttributeRemapper) GetOverrideOnConflictOk() (*bool, bool)`

GetOverrideOnConflictOk returns a tuple with the OverrideOnConflict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideOnConflict

`func (o *LogsAttributeRemapper) SetOverrideOnConflict(v bool)`

SetOverrideOnConflict sets OverrideOnConflict field to given value.

### HasOverrideOnConflict

`func (o *LogsAttributeRemapper) HasOverrideOnConflict() bool`

HasOverrideOnConflict returns a boolean if a field has been set.

### GetPreserveSource

`func (o *LogsAttributeRemapper) GetPreserveSource() bool`

GetPreserveSource returns the PreserveSource field if non-nil, zero value otherwise.

### GetPreserveSourceOk

`func (o *LogsAttributeRemapper) GetPreserveSourceOk() (*bool, bool)`

GetPreserveSourceOk returns a tuple with the PreserveSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveSource

`func (o *LogsAttributeRemapper) SetPreserveSource(v bool)`

SetPreserveSource sets PreserveSource field to given value.

### HasPreserveSource

`func (o *LogsAttributeRemapper) HasPreserveSource() bool`

HasPreserveSource returns a boolean if a field has been set.

### GetSourceType

`func (o *LogsAttributeRemapper) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *LogsAttributeRemapper) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *LogsAttributeRemapper) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *LogsAttributeRemapper) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### GetSources

`func (o *LogsAttributeRemapper) GetSources() []string`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *LogsAttributeRemapper) GetSourcesOk() (*[]string, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *LogsAttributeRemapper) SetSources(v []string)`

SetSources sets Sources field to given value.

### GetTarget

`func (o *LogsAttributeRemapper) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *LogsAttributeRemapper) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *LogsAttributeRemapper) SetTarget(v string)`

SetTarget sets Target field to given value.

### GetTargetFormat

`func (o *LogsAttributeRemapper) GetTargetFormat() TargetFormatType`

GetTargetFormat returns the TargetFormat field if non-nil, zero value otherwise.

### GetTargetFormatOk

`func (o *LogsAttributeRemapper) GetTargetFormatOk() (*TargetFormatType, bool)`

GetTargetFormatOk returns a tuple with the TargetFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetFormat

`func (o *LogsAttributeRemapper) SetTargetFormat(v TargetFormatType)`

SetTargetFormat sets TargetFormat field to given value.

### HasTargetFormat

`func (o *LogsAttributeRemapper) HasTargetFormat() bool`

HasTargetFormat returns a boolean if a field has been set.

### GetTargetType

`func (o *LogsAttributeRemapper) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *LogsAttributeRemapper) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *LogsAttributeRemapper) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *LogsAttributeRemapper) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.

### GetType

`func (o *LogsAttributeRemapper) GetType() LogsAttributeRemapperType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LogsAttributeRemapper) GetTypeOk() (*LogsAttributeRemapperType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LogsAttributeRemapper) SetType(v LogsAttributeRemapperType)`

SetType sets Type field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
