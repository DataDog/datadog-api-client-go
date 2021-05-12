# NotebookResponseData

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Attributes** | [**NotebookResponseDataAttributes**](NotebookResponseDataAttributes.md) |  | 
**Id** | **int64** | Unique notebook ID, assigned when you create the notebook. | [readonly] 
**Type** | [**NotebookResourceType**](NotebookResourceType.md) |  | [default to NOTEBOOKRESOURCETYPE_NOTEBOOKS]

## Methods

### NewNotebookResponseData

`func NewNotebookResponseData(attributes NotebookResponseDataAttributes, id int64, type_ NotebookResourceType, ) *NotebookResponseData`

NewNotebookResponseData instantiates a new NotebookResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotebookResponseDataWithDefaults

`func NewNotebookResponseDataWithDefaults() *NotebookResponseData`

NewNotebookResponseDataWithDefaults instantiates a new NotebookResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *NotebookResponseData) GetAttributes() NotebookResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *NotebookResponseData) GetAttributesOk() (*NotebookResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *NotebookResponseData) SetAttributes(v NotebookResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetId

`func (o *NotebookResponseData) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotebookResponseData) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotebookResponseData) SetId(v int64)`

SetId sets Id field to given value.


### GetType

`func (o *NotebookResponseData) GetType() NotebookResourceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NotebookResponseData) GetTypeOk() (*NotebookResourceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NotebookResponseData) SetType(v NotebookResourceType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


