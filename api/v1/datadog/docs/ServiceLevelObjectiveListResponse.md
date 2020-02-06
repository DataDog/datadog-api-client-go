# ServiceLevelObjectiveListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ServiceLevelObjective**](ServiceLevelObjective.md) | An array of service level objective objects. | 
**Errors** | Pointer to **[]string** | An array of error messages. Each endpoint documents how/whether this field is used. | [optional] 

## Methods

### NewServiceLevelObjectiveListResponse

`func NewServiceLevelObjectiveListResponse(data []ServiceLevelObjective, ) *ServiceLevelObjectiveListResponse`

NewServiceLevelObjectiveListResponse instantiates a new ServiceLevelObjectiveListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceLevelObjectiveListResponseWithDefaults

`func NewServiceLevelObjectiveListResponseWithDefaults() *ServiceLevelObjectiveListResponse`

NewServiceLevelObjectiveListResponseWithDefaults instantiates a new ServiceLevelObjectiveListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ServiceLevelObjectiveListResponse) GetData() []ServiceLevelObjective`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ServiceLevelObjectiveListResponse) GetDataOk() ([]ServiceLevelObjective, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasData

`func (o *ServiceLevelObjectiveListResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### SetData

`func (o *ServiceLevelObjectiveListResponse) SetData(v []ServiceLevelObjective)`

SetData gets a reference to the given []ServiceLevelObjective and assigns it to the Data field.

### GetErrors

`func (o *ServiceLevelObjectiveListResponse) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ServiceLevelObjectiveListResponse) GetErrorsOk() ([]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasErrors

`func (o *ServiceLevelObjectiveListResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### SetErrors

`func (o *ServiceLevelObjectiveListResponse) SetErrors(v []string)`

SetErrors gets a reference to the given []string and assigns it to the Errors field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


