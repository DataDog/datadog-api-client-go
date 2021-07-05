# NotebooksResponseData

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Attributes** | [**NotebooksResponseDataAttributes**](NotebooksResponseDataAttributes.md) |  | 
**Id** | **int64** | Unique notebook ID, assigned when you create the notebook. | [readonly] 
**Type** | [**NotebookResourceType**](NotebookResourceType.md) |  | [default to NOTEBOOKRESOURCETYPE_NOTEBOOKS]

## Methods

### NewNotebooksResponseData

`func NewNotebooksResponseData(attributes NotebooksResponseDataAttributes, id int64, type_ NotebookResourceType) *NotebooksResponseData`

NewNotebooksResponseData instantiates a new NotebooksResponseData object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewNotebooksResponseDataWithDefaults

`func NewNotebooksResponseDataWithDefaults() *NotebooksResponseData`

NewNotebooksResponseDataWithDefaults instantiates a new NotebooksResponseData object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAttributes

`func (o *NotebooksResponseData) GetAttributes() NotebooksResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *NotebooksResponseData) GetAttributesOk() (*NotebooksResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *NotebooksResponseData) SetAttributes(v NotebooksResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetId

`func (o *NotebooksResponseData) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotebooksResponseData) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotebooksResponseData) SetId(v int64)`

SetId sets Id field to given value.


### GetType

`func (o *NotebooksResponseData) GetType() NotebookResourceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NotebooksResponseData) GetTypeOk() (*NotebookResourceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NotebooksResponseData) SetType(v NotebookResourceType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


