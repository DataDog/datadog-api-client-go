# SLOResponse

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Data** | Pointer to [**SLOResponseData**](SLOResponseData.md) |  | [optional] 
**Errors** | Pointer to **[]string** | An array of error messages. Each endpoint documents how/whether this field is used. | [optional] 

## Methods

### NewSLOResponse

`func NewSLOResponse() *SLOResponse`

NewSLOResponse instantiates a new SLOResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewSLOResponseWithDefaults

`func NewSLOResponseWithDefaults() *SLOResponse`

NewSLOResponseWithDefaults instantiates a new SLOResponse object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetData

`func (o *SLOResponse) GetData() SLOResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SLOResponse) GetDataOk() (*SLOResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SLOResponse) SetData(v SLOResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *SLOResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetErrors

`func (o *SLOResponse) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *SLOResponse) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *SLOResponse) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *SLOResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


