# LogsArchiveCreateRequestDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**LogsArchiveCreateRequestAttributes**](LogsArchiveCreateRequestAttributes.md) |  | [optional] 
**Type** | **string** | The type of the resource. The value should always be archives. | [default to "archives"]

## Methods

### NewLogsArchiveCreateRequestDefinition

`func NewLogsArchiveCreateRequestDefinition(type_ string, ) *LogsArchiveCreateRequestDefinition`

NewLogsArchiveCreateRequestDefinition instantiates a new LogsArchiveCreateRequestDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogsArchiveCreateRequestDefinitionWithDefaults

`func NewLogsArchiveCreateRequestDefinitionWithDefaults() *LogsArchiveCreateRequestDefinition`

NewLogsArchiveCreateRequestDefinitionWithDefaults instantiates a new LogsArchiveCreateRequestDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *LogsArchiveCreateRequestDefinition) GetAttributes() LogsArchiveCreateRequestAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *LogsArchiveCreateRequestDefinition) GetAttributesOk() (*LogsArchiveCreateRequestAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *LogsArchiveCreateRequestDefinition) SetAttributes(v LogsArchiveCreateRequestAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *LogsArchiveCreateRequestDefinition) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetType

`func (o *LogsArchiveCreateRequestDefinition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LogsArchiveCreateRequestDefinition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LogsArchiveCreateRequestDefinition) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


