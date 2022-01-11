# NotebookCreateDataAttributes

## Properties

| Name         | Type                                                            | Description                               | Notes                                            |
| ------------ | --------------------------------------------------------------- | ----------------------------------------- | ------------------------------------------------ |
| **Cells**    | [**[]NotebookCellCreateRequest**](NotebookCellCreateRequest.md) | List of cells to display in the notebook. |
| **Metadata** | Pointer to [**NotebookMetadata**](NotebookMetadata.md)          |                                           | [optional]                                       |
| **Name**     | **string**                                                      | The name of the notebook.                 |
| **Status**   | Pointer to [**NotebookStatus**](NotebookStatus.md)              |                                           | [optional] [default to NOTEBOOKSTATUS_PUBLISHED] |
| **Time**     | [**NotebookGlobalTime**](NotebookGlobalTime.md)                 |                                           |

## Methods

### NewNotebookCreateDataAttributes

`func NewNotebookCreateDataAttributes(cells []NotebookCellCreateRequest, name string, time NotebookGlobalTime) *NotebookCreateDataAttributes`

NewNotebookCreateDataAttributes instantiates a new NotebookCreateDataAttributes object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewNotebookCreateDataAttributesWithDefaults

`func NewNotebookCreateDataAttributesWithDefaults() *NotebookCreateDataAttributes`

NewNotebookCreateDataAttributesWithDefaults instantiates a new NotebookCreateDataAttributes object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetCells

`func (o *NotebookCreateDataAttributes) GetCells() []NotebookCellCreateRequest`

GetCells returns the Cells field if non-nil, zero value otherwise.

### GetCellsOk

`func (o *NotebookCreateDataAttributes) GetCellsOk() (*[]NotebookCellCreateRequest, bool)`

GetCellsOk returns a tuple with the Cells field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCells

`func (o *NotebookCreateDataAttributes) SetCells(v []NotebookCellCreateRequest)`

SetCells sets Cells field to given value.

### GetMetadata

`func (o *NotebookCreateDataAttributes) GetMetadata() NotebookMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NotebookCreateDataAttributes) GetMetadataOk() (*NotebookMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NotebookCreateDataAttributes) SetMetadata(v NotebookMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *NotebookCreateDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *NotebookCreateDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NotebookCreateDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NotebookCreateDataAttributes) SetName(v string)`

SetName sets Name field to given value.

### GetStatus

`func (o *NotebookCreateDataAttributes) GetStatus() NotebookStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NotebookCreateDataAttributes) GetStatusOk() (*NotebookStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NotebookCreateDataAttributes) SetStatus(v NotebookStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NotebookCreateDataAttributes) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTime

`func (o *NotebookCreateDataAttributes) GetTime() NotebookGlobalTime`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *NotebookCreateDataAttributes) GetTimeOk() (*NotebookGlobalTime, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *NotebookCreateDataAttributes) SetTime(v NotebookGlobalTime)`

SetTime sets Time field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
