# LogsDateRemapper

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**IsEnabled** | Pointer to **bool** | Whether or not the processor is enabled. | [optional] [default to false]
**Name** | Pointer to **string** | Name of the processor. | [optional] 
**Sources** | **[]string** | Array of source attributes. | 
**Type** | [**LogsDateRemapperType**](LogsDateRemapperType.md) |  | [default to LOGSDATEREMAPPERTYPE_DATE_REMAPPER]

## Methods

### NewLogsDateRemapper

`func NewLogsDateRemapper(sources []string, type_ LogsDateRemapperType) *LogsDateRemapper`

NewLogsDateRemapper instantiates a new LogsDateRemapper object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewLogsDateRemapperWithDefaults

`func NewLogsDateRemapperWithDefaults() *LogsDateRemapper`

NewLogsDateRemapperWithDefaults instantiates a new LogsDateRemapper object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetIsEnabled

`func (o *LogsDateRemapper) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *LogsDateRemapper) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *LogsDateRemapper) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *LogsDateRemapper) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetName

`func (o *LogsDateRemapper) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LogsDateRemapper) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LogsDateRemapper) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LogsDateRemapper) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSources

`func (o *LogsDateRemapper) GetSources() []string`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *LogsDateRemapper) GetSourcesOk() (*[]string, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *LogsDateRemapper) SetSources(v []string)`

SetSources sets Sources field to given value.


### GetType

`func (o *LogsDateRemapper) GetType() LogsDateRemapperType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LogsDateRemapper) GetTypeOk() (*LogsDateRemapperType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LogsDateRemapper) SetType(v LogsDateRemapperType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


