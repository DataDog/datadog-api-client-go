# LogsArchiveOrderDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | [**LogsArchiveOrderAttributes**](LogsArchiveOrderAttributes.md) |  | 
**Type** | [**LogsArchiveOrderDefinitionType**](LogsArchiveOrderDefinitionType.md) |  | [default to LOGSARCHIVEORDERDEFINITIONTYPE_ARCHIVE_ORDER]

## Methods

### NewLogsArchiveOrderDefinition

`func NewLogsArchiveOrderDefinition(attributes LogsArchiveOrderAttributes, type_ LogsArchiveOrderDefinitionType, ) *LogsArchiveOrderDefinition`

NewLogsArchiveOrderDefinition instantiates a new LogsArchiveOrderDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogsArchiveOrderDefinitionWithDefaults

`func NewLogsArchiveOrderDefinitionWithDefaults() *LogsArchiveOrderDefinition`

NewLogsArchiveOrderDefinitionWithDefaults instantiates a new LogsArchiveOrderDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *LogsArchiveOrderDefinition) GetAttributes() LogsArchiveOrderAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *LogsArchiveOrderDefinition) GetAttributesOk() (*LogsArchiveOrderAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *LogsArchiveOrderDefinition) SetAttributes(v LogsArchiveOrderAttributes)`

SetAttributes sets Attributes field to given value.


### GetType

`func (o *LogsArchiveOrderDefinition) GetType() LogsArchiveOrderDefinitionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LogsArchiveOrderDefinition) GetTypeOk() (*LogsArchiveOrderDefinitionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LogsArchiveOrderDefinition) SetType(v LogsArchiveOrderDefinitionType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


