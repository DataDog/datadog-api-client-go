# NotebookCellResponse

## Properties

| Name           | Type                                                                    | Description       | Notes                                                |
| -------------- | ----------------------------------------------------------------------- | ----------------- | ---------------------------------------------------- |
| **Attributes** | [**NotebookCellResponseAttributes**](NotebookCellResponseAttributes.md) |                   |
| **Id**         | **string**                                                              | Notebook cell ID. |
| **Type**       | [**NotebookCellResourceType**](NotebookCellResourceType.md)             |                   | [default to NOTEBOOKCELLRESOURCETYPE_NOTEBOOK_CELLS] |

## Methods

### NewNotebookCellResponse

`func NewNotebookCellResponse(attributes NotebookCellResponseAttributes, id string, type_ NotebookCellResourceType) *NotebookCellResponse`

NewNotebookCellResponse instantiates a new NotebookCellResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewNotebookCellResponseWithDefaults

`func NewNotebookCellResponseWithDefaults() *NotebookCellResponse`

NewNotebookCellResponseWithDefaults instantiates a new NotebookCellResponse object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAttributes

`func (o *NotebookCellResponse) GetAttributes() NotebookCellResponseAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *NotebookCellResponse) GetAttributesOk() (*NotebookCellResponseAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *NotebookCellResponse) SetAttributes(v NotebookCellResponseAttributes)`

SetAttributes sets Attributes field to given value.

### GetId

`func (o *NotebookCellResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotebookCellResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotebookCellResponse) SetId(v string)`

SetId sets Id field to given value.

### GetType

`func (o *NotebookCellResponse) GetType() NotebookCellResourceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NotebookCellResponse) GetTypeOk() (*NotebookCellResourceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NotebookCellResponse) SetType(v NotebookCellResourceType)`

SetType sets Type field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
