# SLOHistoryResponse

## Properties

| Name       | Type                                                                   | Description                                                                       | Notes      |
| ---------- | ---------------------------------------------------------------------- | --------------------------------------------------------------------------------- | ---------- |
| **Data**   | Pointer to [**SLOHistoryResponseData**](SLOHistoryResponseData.md)     |                                                                                   | [optional] |
| **Errors** | Pointer to [**[]SLOHistoryResponseError**](SLOHistoryResponseError.md) | A list of errors while querying the history data for the service level objective. | [optional] |

## Methods

### NewSLOHistoryResponse

`func NewSLOHistoryResponse() *SLOHistoryResponse`

NewSLOHistoryResponse instantiates a new SLOHistoryResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewSLOHistoryResponseWithDefaults

`func NewSLOHistoryResponseWithDefaults() *SLOHistoryResponse`

NewSLOHistoryResponseWithDefaults instantiates a new SLOHistoryResponse object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetData

`func (o *SLOHistoryResponse) GetData() SLOHistoryResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SLOHistoryResponse) GetDataOk() (*SLOHistoryResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SLOHistoryResponse) SetData(v SLOHistoryResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *SLOHistoryResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetErrors

`func (o *SLOHistoryResponse) GetErrors() []SLOHistoryResponseError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *SLOHistoryResponse) GetErrorsOk() (*[]SLOHistoryResponseError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *SLOHistoryResponse) SetErrors(v []SLOHistoryResponseError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *SLOHistoryResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
