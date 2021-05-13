# NotebookCreateData

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Attributes** | [**NotebookCreateDataAttributes**](NotebookCreateDataAttributes.md) |  | 
**Type** | [**NotebookResourceType**](NotebookResourceType.md) |  | [default to NOTEBOOKRESOURCETYPE_NOTEBOOKS]

## Methods

### NewNotebookCreateData

`func NewNotebookCreateData(attributes NotebookCreateDataAttributes, type_ NotebookResourceType, ) *NotebookCreateData`

NewNotebookCreateData instantiates a new NotebookCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotebookCreateDataWithDefaults

`func NewNotebookCreateDataWithDefaults() *NotebookCreateData`

NewNotebookCreateDataWithDefaults instantiates a new NotebookCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *NotebookCreateData) GetAttributes() NotebookCreateDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *NotebookCreateData) GetAttributesOk() (*NotebookCreateDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *NotebookCreateData) SetAttributes(v NotebookCreateDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetType

`func (o *NotebookCreateData) GetType() NotebookResourceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NotebookCreateData) GetTypeOk() (*NotebookResourceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NotebookCreateData) SetType(v NotebookResourceType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


