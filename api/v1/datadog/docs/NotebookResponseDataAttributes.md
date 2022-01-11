# NotebookResponseDataAttributes

## Properties

| Name         | Type                                                   | Description                                             | Notes                                            |
| ------------ | ------------------------------------------------------ | ------------------------------------------------------- | ------------------------------------------------ |
| **Author**   | Pointer to [**NotebookAuthor**](NotebookAuthor.md)     |                                                         | [optional]                                       |
| **Cells**    | [**[]NotebookCellResponse**](NotebookCellResponse.md)  | List of cells to display in the notebook.               |
| **Created**  | Pointer to **time.Time**                               | UTC time stamp for when the notebook was created.       | [optional] [readonly]                            |
| **Metadata** | Pointer to [**NotebookMetadata**](NotebookMetadata.md) |                                                         | [optional]                                       |
| **Modified** | Pointer to **time.Time**                               | UTC time stamp for when the notebook was last modified. | [optional] [readonly]                            |
| **Name**     | **string**                                             | The name of the notebook.                               |
| **Status**   | Pointer to [**NotebookStatus**](NotebookStatus.md)     |                                                         | [optional] [default to NOTEBOOKSTATUS_PUBLISHED] |
| **Time**     | [**NotebookGlobalTime**](NotebookGlobalTime.md)        |                                                         |

## Methods

### NewNotebookResponseDataAttributes

`func NewNotebookResponseDataAttributes(cells []NotebookCellResponse, name string, time NotebookGlobalTime) *NotebookResponseDataAttributes`

NewNotebookResponseDataAttributes instantiates a new NotebookResponseDataAttributes object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewNotebookResponseDataAttributesWithDefaults

`func NewNotebookResponseDataAttributesWithDefaults() *NotebookResponseDataAttributes`

NewNotebookResponseDataAttributesWithDefaults instantiates a new NotebookResponseDataAttributes object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAuthor

`func (o *NotebookResponseDataAttributes) GetAuthor() NotebookAuthor`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *NotebookResponseDataAttributes) GetAuthorOk() (*NotebookAuthor, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *NotebookResponseDataAttributes) SetAuthor(v NotebookAuthor)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *NotebookResponseDataAttributes) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetCells

`func (o *NotebookResponseDataAttributes) GetCells() []NotebookCellResponse`

GetCells returns the Cells field if non-nil, zero value otherwise.

### GetCellsOk

`func (o *NotebookResponseDataAttributes) GetCellsOk() (*[]NotebookCellResponse, bool)`

GetCellsOk returns a tuple with the Cells field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCells

`func (o *NotebookResponseDataAttributes) SetCells(v []NotebookCellResponse)`

SetCells sets Cells field to given value.

### GetCreated

`func (o *NotebookResponseDataAttributes) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *NotebookResponseDataAttributes) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *NotebookResponseDataAttributes) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *NotebookResponseDataAttributes) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetMetadata

`func (o *NotebookResponseDataAttributes) GetMetadata() NotebookMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NotebookResponseDataAttributes) GetMetadataOk() (*NotebookMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NotebookResponseDataAttributes) SetMetadata(v NotebookMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *NotebookResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetModified

`func (o *NotebookResponseDataAttributes) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *NotebookResponseDataAttributes) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *NotebookResponseDataAttributes) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *NotebookResponseDataAttributes) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetName

`func (o *NotebookResponseDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NotebookResponseDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NotebookResponseDataAttributes) SetName(v string)`

SetName sets Name field to given value.

### GetStatus

`func (o *NotebookResponseDataAttributes) GetStatus() NotebookStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NotebookResponseDataAttributes) GetStatusOk() (*NotebookStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NotebookResponseDataAttributes) SetStatus(v NotebookStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NotebookResponseDataAttributes) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTime

`func (o *NotebookResponseDataAttributes) GetTime() NotebookGlobalTime`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *NotebookResponseDataAttributes) GetTimeOk() (*NotebookGlobalTime, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *NotebookResponseDataAttributes) SetTime(v NotebookGlobalTime)`

SetTime sets Time field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
