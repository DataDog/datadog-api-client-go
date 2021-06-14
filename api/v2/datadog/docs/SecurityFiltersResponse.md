# SecurityFiltersResponse

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Data** | Pointer to [**[]SecurityFilter**](SecurityFilter.md) | A list of security filters objects. | [optional] 
**Meta** | Pointer to [**SecurityFilterMeta**](SecurityFilterMeta.md) |  | [optional] 

## Methods

### NewSecurityFiltersResponse

`func NewSecurityFiltersResponse() *SecurityFiltersResponse`

NewSecurityFiltersResponse instantiates a new SecurityFiltersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityFiltersResponseWithDefaults

`func NewSecurityFiltersResponseWithDefaults() *SecurityFiltersResponse`

NewSecurityFiltersResponseWithDefaults instantiates a new SecurityFiltersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *SecurityFiltersResponse) GetData() []SecurityFilter`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SecurityFiltersResponse) GetDataOk() (*[]SecurityFilter, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SecurityFiltersResponse) SetData(v []SecurityFilter)`

SetData sets Data field to given value.

### HasData

`func (o *SecurityFiltersResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *SecurityFiltersResponse) GetMeta() SecurityFilterMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *SecurityFiltersResponse) GetMetaOk() (*SecurityFilterMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *SecurityFiltersResponse) SetMeta(v SecurityFilterMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *SecurityFiltersResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


