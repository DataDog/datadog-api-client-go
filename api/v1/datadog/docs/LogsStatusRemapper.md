# LogsStatusRemapper

## Properties

| Name          | Type                                                    | Description                              | Notes                                               |
| ------------- | ------------------------------------------------------- | ---------------------------------------- | --------------------------------------------------- |
| **IsEnabled** | Pointer to **bool**                                     | Whether or not the processor is enabled. | [optional] [default to false]                       |
| **Name**      | Pointer to **string**                                   | Name of the processor.                   | [optional]                                          |
| **Sources**   | **[]string**                                            | Array of source attributes.              |
| **Type**      | [**LogsStatusRemapperType**](LogsStatusRemapperType.md) |                                          | [default to LOGSSTATUSREMAPPERTYPE_STATUS_REMAPPER] |

## Methods

### NewLogsStatusRemapper

`func NewLogsStatusRemapper(sources []string, type_ LogsStatusRemapperType) *LogsStatusRemapper`

NewLogsStatusRemapper instantiates a new LogsStatusRemapper object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewLogsStatusRemapperWithDefaults

`func NewLogsStatusRemapperWithDefaults() *LogsStatusRemapper`

NewLogsStatusRemapperWithDefaults instantiates a new LogsStatusRemapper object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetIsEnabled

`func (o *LogsStatusRemapper) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *LogsStatusRemapper) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *LogsStatusRemapper) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *LogsStatusRemapper) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetName

`func (o *LogsStatusRemapper) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LogsStatusRemapper) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LogsStatusRemapper) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LogsStatusRemapper) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSources

`func (o *LogsStatusRemapper) GetSources() []string`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *LogsStatusRemapper) GetSourcesOk() (*[]string, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *LogsStatusRemapper) SetSources(v []string)`

SetSources sets Sources field to given value.

### GetType

`func (o *LogsStatusRemapper) GetType() LogsStatusRemapperType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LogsStatusRemapper) GetTypeOk() (*LogsStatusRemapperType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LogsStatusRemapper) SetType(v LogsStatusRemapperType)`

SetType sets Type field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
