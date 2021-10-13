# NotebookMetadata

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**IsTemplate** | Pointer to **bool** | Whether or not the notebook is a template. | [optional] [default to false]
**TakeSnapshots** | Pointer to **bool** | Whether or not the notebook takes snapshot image backups of the notebook&#39;s fixed-time graphs. | [optional] [default to false]
**Type** | Pointer to [**NullableNotebookMetadataType**](NotebookMetadataType.md) |  | [optional] [default to "null"]

## Methods

### NewNotebookMetadata

`func NewNotebookMetadata() *NotebookMetadata`

NewNotebookMetadata instantiates a new NotebookMetadata object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewNotebookMetadataWithDefaults

`func NewNotebookMetadataWithDefaults() *NotebookMetadata`

NewNotebookMetadataWithDefaults instantiates a new NotebookMetadata object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetIsTemplate

`func (o *NotebookMetadata) GetIsTemplate() bool`

GetIsTemplate returns the IsTemplate field if non-nil, zero value otherwise.

### GetIsTemplateOk

`func (o *NotebookMetadata) GetIsTemplateOk() (*bool, bool)`

GetIsTemplateOk returns a tuple with the IsTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTemplate

`func (o *NotebookMetadata) SetIsTemplate(v bool)`

SetIsTemplate sets IsTemplate field to given value.

### HasIsTemplate

`func (o *NotebookMetadata) HasIsTemplate() bool`

HasIsTemplate returns a boolean if a field has been set.

### GetTakeSnapshots

`func (o *NotebookMetadata) GetTakeSnapshots() bool`

GetTakeSnapshots returns the TakeSnapshots field if non-nil, zero value otherwise.

### GetTakeSnapshotsOk

`func (o *NotebookMetadata) GetTakeSnapshotsOk() (*bool, bool)`

GetTakeSnapshotsOk returns a tuple with the TakeSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTakeSnapshots

`func (o *NotebookMetadata) SetTakeSnapshots(v bool)`

SetTakeSnapshots sets TakeSnapshots field to given value.

### HasTakeSnapshots

`func (o *NotebookMetadata) HasTakeSnapshots() bool`

HasTakeSnapshots returns a boolean if a field has been set.

### GetType

`func (o *NotebookMetadata) GetType() NotebookMetadataType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NotebookMetadata) GetTypeOk() (*NotebookMetadataType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NotebookMetadata) SetType(v NotebookMetadataType)`

SetType sets Type field to given value.

### HasType

`func (o *NotebookMetadata) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *NotebookMetadata) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *NotebookMetadata) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


