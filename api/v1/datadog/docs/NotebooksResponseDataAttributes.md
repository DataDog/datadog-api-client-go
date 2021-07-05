# NotebooksResponseDataAttributes

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Author** | Pointer to [**NotebookAuthor**](NotebookAuthor.md) |  | [optional] 
**Cells** | Pointer to [**[]NotebookCellResponse**](NotebookCellResponse.md) | List of cells to display in the notebook. | [optional] 
**Created** | Pointer to **time.Time** | UTC time stamp for when the notebook was created. | [optional] [readonly] 
**Modified** | Pointer to **time.Time** | UTC time stamp for when the notebook was last modified. | [optional] [readonly] 
**Name** | **string** | The name of the notebook. | 
**Status** | Pointer to [**NotebookStatus**](NotebookStatus.md) |  | [optional] [default to NOTEBOOKSTATUS_PUBLISHED]
**Time** | Pointer to [**NotebookGlobalTime**](NotebookGlobalTime.md) |  | [optional] 

## Methods

### NewNotebooksResponseDataAttributes

`func NewNotebooksResponseDataAttributes(name string) *NotebooksResponseDataAttributes`

NewNotebooksResponseDataAttributes instantiates a new NotebooksResponseDataAttributes object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewNotebooksResponseDataAttributesWithDefaults

`func NewNotebooksResponseDataAttributesWithDefaults() *NotebooksResponseDataAttributes`

NewNotebooksResponseDataAttributesWithDefaults instantiates a new NotebooksResponseDataAttributes object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAuthor

`func (o *NotebooksResponseDataAttributes) GetAuthor() NotebookAuthor`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *NotebooksResponseDataAttributes) GetAuthorOk() (*NotebookAuthor, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *NotebooksResponseDataAttributes) SetAuthor(v NotebookAuthor)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *NotebooksResponseDataAttributes) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetCells

`func (o *NotebooksResponseDataAttributes) GetCells() []NotebookCellResponse`

GetCells returns the Cells field if non-nil, zero value otherwise.

### GetCellsOk

`func (o *NotebooksResponseDataAttributes) GetCellsOk() (*[]NotebookCellResponse, bool)`

GetCellsOk returns a tuple with the Cells field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCells

`func (o *NotebooksResponseDataAttributes) SetCells(v []NotebookCellResponse)`

SetCells sets Cells field to given value.

### HasCells

`func (o *NotebooksResponseDataAttributes) HasCells() bool`

HasCells returns a boolean if a field has been set.

### GetCreated

`func (o *NotebooksResponseDataAttributes) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *NotebooksResponseDataAttributes) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *NotebooksResponseDataAttributes) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *NotebooksResponseDataAttributes) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetModified

`func (o *NotebooksResponseDataAttributes) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *NotebooksResponseDataAttributes) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *NotebooksResponseDataAttributes) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *NotebooksResponseDataAttributes) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetName

`func (o *NotebooksResponseDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NotebooksResponseDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NotebooksResponseDataAttributes) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *NotebooksResponseDataAttributes) GetStatus() NotebookStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NotebooksResponseDataAttributes) GetStatusOk() (*NotebookStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NotebooksResponseDataAttributes) SetStatus(v NotebookStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NotebooksResponseDataAttributes) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTime

`func (o *NotebooksResponseDataAttributes) GetTime() NotebookGlobalTime`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *NotebooksResponseDataAttributes) GetTimeOk() (*NotebookGlobalTime, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *NotebooksResponseDataAttributes) SetTime(v NotebookGlobalTime)`

SetTime sets Time field to given value.

### HasTime

`func (o *NotebooksResponseDataAttributes) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


