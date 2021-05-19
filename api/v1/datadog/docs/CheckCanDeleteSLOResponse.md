# CheckCanDeleteSLOResponse

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Data** | Pointer to [**CheckCanDeleteSLOResponseData**](CheckCanDeleteSLOResponseData.md) |  | [optional] 
**Errors** | Pointer to **map[string]string** | A mapping of SLO id to it&#39;s current usages. | [optional] 

## Methods

### NewCheckCanDeleteSLOResponse

`func NewCheckCanDeleteSLOResponse() *CheckCanDeleteSLOResponse`

NewCheckCanDeleteSLOResponse instantiates a new CheckCanDeleteSLOResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckCanDeleteSLOResponseWithDefaults

`func NewCheckCanDeleteSLOResponseWithDefaults() *CheckCanDeleteSLOResponse`

NewCheckCanDeleteSLOResponseWithDefaults instantiates a new CheckCanDeleteSLOResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CheckCanDeleteSLOResponse) GetData() CheckCanDeleteSLOResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CheckCanDeleteSLOResponse) GetDataOk() (*CheckCanDeleteSLOResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CheckCanDeleteSLOResponse) SetData(v CheckCanDeleteSLOResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *CheckCanDeleteSLOResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetErrors

`func (o *CheckCanDeleteSLOResponse) GetErrors() map[string]string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *CheckCanDeleteSLOResponse) GetErrorsOk() (*map[string]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *CheckCanDeleteSLOResponse) SetErrors(v map[string]string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *CheckCanDeleteSLOResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


