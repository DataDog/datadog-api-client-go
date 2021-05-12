# NotebookUpdateCell

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Attributes** | [**NotebookCellUpdateRequestAttributes**](NotebookCellUpdateRequestAttributes.md) |  | 
**Type** | [**NotebookCellResourceType**](NotebookCellResourceType.md) |  | [default to NOTEBOOKCELLRESOURCETYPE_NOTEBOOK_CELLS]
**Id** | **string** | Notebook cell ID. | 

## Methods

### NewNotebookUpdateCell

`func NewNotebookUpdateCell(attributes NotebookCellUpdateRequestAttributes, type_ NotebookCellResourceType, id string, ) *NotebookUpdateCell`

NewNotebookUpdateCell instantiates a new NotebookUpdateCell object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotebookUpdateCellWithDefaults

`func NewNotebookUpdateCellWithDefaults() *NotebookUpdateCell`

NewNotebookUpdateCellWithDefaults instantiates a new NotebookUpdateCell object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *NotebookUpdateCell) GetAttributes() NotebookCellUpdateRequestAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *NotebookUpdateCell) GetAttributesOk() (*NotebookCellUpdateRequestAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *NotebookUpdateCell) SetAttributes(v NotebookCellUpdateRequestAttributes)`

SetAttributes sets Attributes field to given value.


### GetType

`func (o *NotebookUpdateCell) GetType() NotebookCellResourceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NotebookUpdateCell) GetTypeOk() (*NotebookCellResourceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NotebookUpdateCell) SetType(v NotebookCellResourceType)`

SetType sets Type field to given value.


### GetId

`func (o *NotebookUpdateCell) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotebookUpdateCell) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotebookUpdateCell) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


