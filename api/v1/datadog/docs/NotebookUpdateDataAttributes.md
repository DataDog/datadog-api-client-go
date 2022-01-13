# NotebookUpdateDataAttributes

## Properties

| Name         | Type                                                   | Description                               | Notes                                            |
| ------------ | ------------------------------------------------------ | ----------------------------------------- | ------------------------------------------------ |
| **Cells**    | [**[]NotebookUpdateCell**](NotebookUpdateCell.md)      | List of cells to display in the notebook. |
| **Metadata** | Pointer to [**NotebookMetadata**](NotebookMetadata.md) |                                           | [optional]                                       |
| **Name**     | **string**                                             | The name of the notebook.                 |
| **Status**   | Pointer to [**NotebookStatus**](NotebookStatus.md)     |                                           | [optional] [default to NOTEBOOKSTATUS_PUBLISHED] |
| **Time**     | [**NotebookGlobalTime**](NotebookGlobalTime.md)        |                                           |

## Methods

### NewNotebookUpdateDataAttributes

`func NewNotebookUpdateDataAttributes(cells []NotebookUpdateCell, name string, time NotebookGlobalTime) *NotebookUpdateDataAttributes`

NewNotebookUpdateDataAttributes instantiates a new NotebookUpdateDataAttributes object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewNotebookUpdateDataAttributesWithDefaults

`func NewNotebookUpdateDataAttributesWithDefaults() *NotebookUpdateDataAttributes`

NewNotebookUpdateDataAttributesWithDefaults instantiates a new NotebookUpdateDataAttributes object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetCells

`func (o *NotebookUpdateDataAttributes) GetCells() []NotebookUpdateCell`

GetCells returns the Cells field if non-nil, zero value otherwise.

### GetCellsOk

`func (o *NotebookUpdateDataAttributes) GetCellsOk() (*[]NotebookUpdateCell, bool)`

GetCellsOk returns a tuple with the Cells field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCells

`func (o *NotebookUpdateDataAttributes) SetCells(v []NotebookUpdateCell)`

SetCells sets Cells field to given value.

### GetMetadata

`func (o *NotebookUpdateDataAttributes) GetMetadata() NotebookMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NotebookUpdateDataAttributes) GetMetadataOk() (*NotebookMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NotebookUpdateDataAttributes) SetMetadata(v NotebookMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *NotebookUpdateDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *NotebookUpdateDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NotebookUpdateDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NotebookUpdateDataAttributes) SetName(v string)`

SetName sets Name field to given value.

### GetStatus

`func (o *NotebookUpdateDataAttributes) GetStatus() NotebookStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NotebookUpdateDataAttributes) GetStatusOk() (*NotebookStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NotebookUpdateDataAttributes) SetStatus(v NotebookStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NotebookUpdateDataAttributes) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTime

`func (o *NotebookUpdateDataAttributes) GetTime() NotebookGlobalTime`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *NotebookUpdateDataAttributes) GetTimeOk() (*NotebookGlobalTime, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *NotebookUpdateDataAttributes) SetTime(v NotebookGlobalTime)`

SetTime sets Time field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
