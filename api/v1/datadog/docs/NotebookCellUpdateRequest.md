# NotebookCellUpdateRequest

## Properties

| Name           | Type                                                                              | Description       | Notes                                                |
| -------------- | --------------------------------------------------------------------------------- | ----------------- | ---------------------------------------------------- |
| **Attributes** | [**NotebookCellUpdateRequestAttributes**](NotebookCellUpdateRequestAttributes.md) |                   |
| **Id**         | **string**                                                                        | Notebook cell ID. |
| **Type**       | [**NotebookCellResourceType**](NotebookCellResourceType.md)                       |                   | [default to NOTEBOOKCELLRESOURCETYPE_NOTEBOOK_CELLS] |

## Methods

### NewNotebookCellUpdateRequest

`func NewNotebookCellUpdateRequest(attributes NotebookCellUpdateRequestAttributes, id string, type_ NotebookCellResourceType) *NotebookCellUpdateRequest`

NewNotebookCellUpdateRequest instantiates a new NotebookCellUpdateRequest object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewNotebookCellUpdateRequestWithDefaults

`func NewNotebookCellUpdateRequestWithDefaults() *NotebookCellUpdateRequest`

NewNotebookCellUpdateRequestWithDefaults instantiates a new NotebookCellUpdateRequest object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAttributes

`func (o *NotebookCellUpdateRequest) GetAttributes() NotebookCellUpdateRequestAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *NotebookCellUpdateRequest) GetAttributesOk() (*NotebookCellUpdateRequestAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *NotebookCellUpdateRequest) SetAttributes(v NotebookCellUpdateRequestAttributes)`

SetAttributes sets Attributes field to given value.

### GetId

`func (o *NotebookCellUpdateRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotebookCellUpdateRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotebookCellUpdateRequest) SetId(v string)`

SetId sets Id field to given value.

### GetType

`func (o *NotebookCellUpdateRequest) GetType() NotebookCellResourceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NotebookCellUpdateRequest) GetTypeOk() (*NotebookCellResourceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NotebookCellUpdateRequest) SetType(v NotebookCellResourceType)`

SetType sets Type field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
