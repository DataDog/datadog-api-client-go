# LogsArchiveDefinition

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Attributes** | Pointer to [**LogsArchiveAttributes**](LogsArchiveAttributes.md) |  | [optional] 
**Id** | Pointer to **string** | The archive ID. | [optional] [readonly] 
**Type** | **string** | The type of the resource. The value should always be archives. | [readonly] [default to "archives"]

## Methods

### NewLogsArchiveDefinition

`func NewLogsArchiveDefinition(type_ string) *LogsArchiveDefinition`

NewLogsArchiveDefinition instantiates a new LogsArchiveDefinition object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewLogsArchiveDefinitionWithDefaults

`func NewLogsArchiveDefinitionWithDefaults() *LogsArchiveDefinition`

NewLogsArchiveDefinitionWithDefaults instantiates a new LogsArchiveDefinition object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAttributes

`func (o *LogsArchiveDefinition) GetAttributes() LogsArchiveAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *LogsArchiveDefinition) GetAttributesOk() (*LogsArchiveAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *LogsArchiveDefinition) SetAttributes(v LogsArchiveAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *LogsArchiveDefinition) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *LogsArchiveDefinition) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LogsArchiveDefinition) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LogsArchiveDefinition) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LogsArchiveDefinition) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *LogsArchiveDefinition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LogsArchiveDefinition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LogsArchiveDefinition) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


