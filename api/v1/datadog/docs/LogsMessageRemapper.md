# LogsMessageRemapper

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | Pointer to **bool** | Whether or not the processor is enabled. | [optional] [default to false]
**Name** | Pointer to **string** | Name of the processor. | [optional] 
**Sources** | Pointer to **[]string** | Array of source attributes. | [default to ["msg"]]
**Type** | Pointer to [**LogsMessageRemapperType**](LogsMessageRemapperType.md) |  | [default to "message-remapper"]

## Methods

### NewLogsMessageRemapper

`func NewLogsMessageRemapper(sources []string, type_ LogsMessageRemapperType, ) *LogsMessageRemapper`

NewLogsMessageRemapper instantiates a new LogsMessageRemapper object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogsMessageRemapperWithDefaults

`func NewLogsMessageRemapperWithDefaults() *LogsMessageRemapper`

NewLogsMessageRemapperWithDefaults instantiates a new LogsMessageRemapper object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *LogsMessageRemapper) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *LogsMessageRemapper) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *LogsMessageRemapper) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *LogsMessageRemapper) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetName

`func (o *LogsMessageRemapper) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LogsMessageRemapper) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LogsMessageRemapper) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LogsMessageRemapper) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSources

`func (o *LogsMessageRemapper) GetSources() []string`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *LogsMessageRemapper) GetSourcesOk() (*[]string, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *LogsMessageRemapper) SetSources(v []string)`

SetSources sets Sources field to given value.


### GetType

`func (o *LogsMessageRemapper) GetType() LogsMessageRemapperType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LogsMessageRemapper) GetTypeOk() (*LogsMessageRemapperType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LogsMessageRemapper) SetType(v LogsMessageRemapperType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


