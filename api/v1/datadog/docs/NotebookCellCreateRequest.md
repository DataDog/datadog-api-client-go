# NotebookCellCreateRequest

## Properties

| Name           | Type                                                                              | Description | Notes                                                |
| -------------- | --------------------------------------------------------------------------------- | ----------- | ---------------------------------------------------- |
| **Attributes** | [**NotebookCellCreateRequestAttributes**](NotebookCellCreateRequestAttributes.md) |             |
| **Type**       | [**NotebookCellResourceType**](NotebookCellResourceType.md)                       |             | [default to NOTEBOOKCELLRESOURCETYPE_NOTEBOOK_CELLS] |

## Methods

### NewNotebookCellCreateRequest

`func NewNotebookCellCreateRequest(attributes NotebookCellCreateRequestAttributes, type_ NotebookCellResourceType) *NotebookCellCreateRequest`

NewNotebookCellCreateRequest instantiates a new NotebookCellCreateRequest object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewNotebookCellCreateRequestWithDefaults

`func NewNotebookCellCreateRequestWithDefaults() *NotebookCellCreateRequest`

NewNotebookCellCreateRequestWithDefaults instantiates a new NotebookCellCreateRequest object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAttributes

`func (o *NotebookCellCreateRequest) GetAttributes() NotebookCellCreateRequestAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *NotebookCellCreateRequest) GetAttributesOk() (*NotebookCellCreateRequestAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *NotebookCellCreateRequest) SetAttributes(v NotebookCellCreateRequestAttributes)`

SetAttributes sets Attributes field to given value.

### GetType

`func (o *NotebookCellCreateRequest) GetType() NotebookCellResourceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NotebookCellCreateRequest) GetTypeOk() (*NotebookCellResourceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NotebookCellCreateRequest) SetType(v NotebookCellResourceType)`

SetType sets Type field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
