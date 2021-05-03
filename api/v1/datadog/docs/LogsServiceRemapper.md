# LogsServiceRemapper

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**IsEnabled** | Pointer to **bool** | Whether or not the processor is enabled. | [optional] [default to false]
**Name** | Pointer to **string** | Name of the processor. | [optional] 
**Sources** | **[]string** | Array of source attributes. | 
**Type** | [**LogsServiceRemapperType**](LogsServiceRemapperType.md) |  | [default to LOGSSERVICEREMAPPERTYPE_SERVICE_REMAPPER]

## Methods

### NewLogsServiceRemapper

`func NewLogsServiceRemapper(sources []string, type_ LogsServiceRemapperType, ) *LogsServiceRemapper`

NewLogsServiceRemapper instantiates a new LogsServiceRemapper object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogsServiceRemapperWithDefaults

`func NewLogsServiceRemapperWithDefaults() *LogsServiceRemapper`

NewLogsServiceRemapperWithDefaults instantiates a new LogsServiceRemapper object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *LogsServiceRemapper) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *LogsServiceRemapper) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *LogsServiceRemapper) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *LogsServiceRemapper) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetName

`func (o *LogsServiceRemapper) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LogsServiceRemapper) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LogsServiceRemapper) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LogsServiceRemapper) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSources

`func (o *LogsServiceRemapper) GetSources() []string`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *LogsServiceRemapper) GetSourcesOk() (*[]string, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *LogsServiceRemapper) SetSources(v []string)`

SetSources sets Sources field to given value.


### GetType

`func (o *LogsServiceRemapper) GetType() LogsServiceRemapperType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LogsServiceRemapper) GetTypeOk() (*LogsServiceRemapperType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LogsServiceRemapper) SetType(v LogsServiceRemapperType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


