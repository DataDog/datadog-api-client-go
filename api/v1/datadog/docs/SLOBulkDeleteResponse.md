# SLOBulkDeleteResponse

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Data** | Pointer to [**SLOBulkDeleteResponseData**](SLOBulkDeleteResponseData.md) |  | [optional] 
**Errors** | Pointer to [**[]SLOBulkDeleteError**](SLOBulkDeleteError.md) | Array of errors object returned. | [optional] 

## Methods

### NewSLOBulkDeleteResponse

`func NewSLOBulkDeleteResponse() *SLOBulkDeleteResponse`

NewSLOBulkDeleteResponse instantiates a new SLOBulkDeleteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSLOBulkDeleteResponseWithDefaults

`func NewSLOBulkDeleteResponseWithDefaults() *SLOBulkDeleteResponse`

NewSLOBulkDeleteResponseWithDefaults instantiates a new SLOBulkDeleteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *SLOBulkDeleteResponse) GetData() SLOBulkDeleteResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SLOBulkDeleteResponse) GetDataOk() (*SLOBulkDeleteResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SLOBulkDeleteResponse) SetData(v SLOBulkDeleteResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *SLOBulkDeleteResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetErrors

`func (o *SLOBulkDeleteResponse) GetErrors() []SLOBulkDeleteError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *SLOBulkDeleteResponse) GetErrorsOk() (*[]SLOBulkDeleteError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *SLOBulkDeleteResponse) SetErrors(v []SLOBulkDeleteError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *SLOBulkDeleteResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


