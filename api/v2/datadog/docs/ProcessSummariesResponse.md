# ProcessSummariesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ProcessSummary**](ProcessSummary.md) | Array of process summary objects. | [optional] 
**Meta** | Pointer to [**ProcessSummariesMeta**](ProcessSummariesMeta.md) |  | [optional] 

## Methods

### NewProcessSummariesResponse

`func NewProcessSummariesResponse() *ProcessSummariesResponse`

NewProcessSummariesResponse instantiates a new ProcessSummariesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessSummariesResponseWithDefaults

`func NewProcessSummariesResponseWithDefaults() *ProcessSummariesResponse`

NewProcessSummariesResponseWithDefaults instantiates a new ProcessSummariesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ProcessSummariesResponse) GetData() []ProcessSummary`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ProcessSummariesResponse) GetDataOk() (*[]ProcessSummary, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ProcessSummariesResponse) SetData(v []ProcessSummary)`

SetData sets Data field to given value.

### HasData

`func (o *ProcessSummariesResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *ProcessSummariesResponse) GetMeta() ProcessSummariesMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ProcessSummariesResponse) GetMetaOk() (*ProcessSummariesMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ProcessSummariesResponse) SetMeta(v ProcessSummariesMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ProcessSummariesResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


