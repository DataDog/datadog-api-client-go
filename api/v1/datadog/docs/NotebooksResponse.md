# NotebooksResponse

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Data** | Pointer to [**[]NotebooksResponseData**](NotebooksResponseData.md) | List of notebook definitions. | [optional] 
**Meta** | Pointer to [**NotebooksResponseMeta**](NotebooksResponseMeta.md) |  | [optional] 

## Methods

### NewNotebooksResponse

`func NewNotebooksResponse() *NotebooksResponse`

NewNotebooksResponse instantiates a new NotebooksResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotebooksResponseWithDefaults

`func NewNotebooksResponseWithDefaults() *NotebooksResponse`

NewNotebooksResponseWithDefaults instantiates a new NotebooksResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *NotebooksResponse) GetData() []NotebooksResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *NotebooksResponse) GetDataOk() (*[]NotebooksResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *NotebooksResponse) SetData(v []NotebooksResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *NotebooksResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *NotebooksResponse) GetMeta() NotebooksResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *NotebooksResponse) GetMetaOk() (*NotebooksResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *NotebooksResponse) SetMeta(v NotebooksResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *NotebooksResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


