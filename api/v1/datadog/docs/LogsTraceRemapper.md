# LogsTraceRemapper

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | Pointer to **bool** | Whether or not the processor is enabled. | [optional] [default to false]
**Name** | Pointer to **string** | Name of the processor. | [optional] 
**Sources** | Pointer to **[]string** | Array of source attributes. | [optional] [default to ["dd.trace_id"]]
**Type** | Pointer to [**LogsTraceRemapperType**](LogsTraceRemapperType.md) |  | [default to "trace-id-remapper"]

## Methods

### NewLogsTraceRemapper

`func NewLogsTraceRemapper(type_ LogsTraceRemapperType, ) *LogsTraceRemapper`

NewLogsTraceRemapper instantiates a new LogsTraceRemapper object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogsTraceRemapperWithDefaults

`func NewLogsTraceRemapperWithDefaults() *LogsTraceRemapper`

NewLogsTraceRemapperWithDefaults instantiates a new LogsTraceRemapper object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *LogsTraceRemapper) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *LogsTraceRemapper) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *LogsTraceRemapper) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *LogsTraceRemapper) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetName

`func (o *LogsTraceRemapper) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LogsTraceRemapper) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LogsTraceRemapper) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LogsTraceRemapper) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSources

`func (o *LogsTraceRemapper) GetSources() []string`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *LogsTraceRemapper) GetSourcesOk() (*[]string, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *LogsTraceRemapper) SetSources(v []string)`

SetSources sets Sources field to given value.

### HasSources

`func (o *LogsTraceRemapper) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetType

`func (o *LogsTraceRemapper) GetType() LogsTraceRemapperType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LogsTraceRemapper) GetTypeOk() (*LogsTraceRemapperType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LogsTraceRemapper) SetType(v LogsTraceRemapperType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


