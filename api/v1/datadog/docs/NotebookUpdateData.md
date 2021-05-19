# NotebookUpdateData

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Attributes** | [**NotebookUpdateDataAttributes**](NotebookUpdateDataAttributes.md) |  | 
**Type** | [**NotebookResourceType**](NotebookResourceType.md) |  | [default to NOTEBOOKRESOURCETYPE_NOTEBOOKS]

## Methods

### NewNotebookUpdateData

`func NewNotebookUpdateData(attributes NotebookUpdateDataAttributes, type_ NotebookResourceType, ) *NotebookUpdateData`

NewNotebookUpdateData instantiates a new NotebookUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotebookUpdateDataWithDefaults

`func NewNotebookUpdateDataWithDefaults() *NotebookUpdateData`

NewNotebookUpdateDataWithDefaults instantiates a new NotebookUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *NotebookUpdateData) GetAttributes() NotebookUpdateDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *NotebookUpdateData) GetAttributesOk() (*NotebookUpdateDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *NotebookUpdateData) SetAttributes(v NotebookUpdateDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetType

`func (o *NotebookUpdateData) GetType() NotebookResourceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NotebookUpdateData) GetTypeOk() (*NotebookResourceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NotebookUpdateData) SetType(v NotebookResourceType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


