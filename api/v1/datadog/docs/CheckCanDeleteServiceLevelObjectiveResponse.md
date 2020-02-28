# CheckCanDeleteServiceLevelObjectiveResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**CheckCanDeleteServiceLevelObjectiveResponseData**](CheckCanDeleteServiceLevelObjectiveResponse_data.md) |  | 
**Errors** | Pointer to **map[string]string** | A mapping of SLO id to it&#39;s current usages. | [optional] 

## Methods

### NewCheckCanDeleteServiceLevelObjectiveResponse

`func NewCheckCanDeleteServiceLevelObjectiveResponse(data CheckCanDeleteServiceLevelObjectiveResponseData, ) *CheckCanDeleteServiceLevelObjectiveResponse`

NewCheckCanDeleteServiceLevelObjectiveResponse instantiates a new CheckCanDeleteServiceLevelObjectiveResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckCanDeleteServiceLevelObjectiveResponseWithDefaults

`func NewCheckCanDeleteServiceLevelObjectiveResponseWithDefaults() *CheckCanDeleteServiceLevelObjectiveResponse`

NewCheckCanDeleteServiceLevelObjectiveResponseWithDefaults instantiates a new CheckCanDeleteServiceLevelObjectiveResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CheckCanDeleteServiceLevelObjectiveResponse) GetData() CheckCanDeleteServiceLevelObjectiveResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CheckCanDeleteServiceLevelObjectiveResponse) GetDataOk() (CheckCanDeleteServiceLevelObjectiveResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasData

`func (o *CheckCanDeleteServiceLevelObjectiveResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### SetData

`func (o *CheckCanDeleteServiceLevelObjectiveResponse) SetData(v CheckCanDeleteServiceLevelObjectiveResponseData)`

SetData gets a reference to the given CheckCanDeleteServiceLevelObjectiveResponseData and assigns it to the Data field.

### GetErrors

`func (o *CheckCanDeleteServiceLevelObjectiveResponse) GetErrors() map[string]string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *CheckCanDeleteServiceLevelObjectiveResponse) GetErrorsOk() (map[string]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasErrors

`func (o *CheckCanDeleteServiceLevelObjectiveResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### SetErrors

`func (o *CheckCanDeleteServiceLevelObjectiveResponse) SetErrors(v map[string]string)`

SetErrors gets a reference to the given map[string]string and assigns it to the Errors field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


