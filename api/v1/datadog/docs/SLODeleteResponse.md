# SLODeleteResponse

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Data** | Pointer to **[]string** | An array containing the ID of the deleted service level objective object. | [optional] 
**Errors** | Pointer to **map[string]string** | An dictionary containing the ID of the SLO as key and a deletion error as value. | [optional] 

## Methods

### NewSLODeleteResponse

`func NewSLODeleteResponse() *SLODeleteResponse`

NewSLODeleteResponse instantiates a new SLODeleteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSLODeleteResponseWithDefaults

`func NewSLODeleteResponseWithDefaults() *SLODeleteResponse`

NewSLODeleteResponseWithDefaults instantiates a new SLODeleteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *SLODeleteResponse) GetData() []string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SLODeleteResponse) GetDataOk() (*[]string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SLODeleteResponse) SetData(v []string)`

SetData sets Data field to given value.

### HasData

`func (o *SLODeleteResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetErrors

`func (o *SLODeleteResponse) GetErrors() map[string]string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *SLODeleteResponse) GetErrorsOk() (*map[string]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *SLODeleteResponse) SetErrors(v map[string]string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *SLODeleteResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


