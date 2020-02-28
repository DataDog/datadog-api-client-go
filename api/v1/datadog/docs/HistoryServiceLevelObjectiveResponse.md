# HistoryServiceLevelObjectiveResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**HistoryServiceLevelObjectiveResponseData**](HistoryServiceLevelObjectiveResponse_data.md) |  | 
**Errors** | Pointer to [**[]HistoryServiceLevelObjectiveResponseError**](HistoryServiceLevelObjectiveResponseError.md) | A list of errors while querying the history data for the service level obective. | [optional] 

## Methods

### NewHistoryServiceLevelObjectiveResponse

`func NewHistoryServiceLevelObjectiveResponse(data HistoryServiceLevelObjectiveResponseData, ) *HistoryServiceLevelObjectiveResponse`

NewHistoryServiceLevelObjectiveResponse instantiates a new HistoryServiceLevelObjectiveResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoryServiceLevelObjectiveResponseWithDefaults

`func NewHistoryServiceLevelObjectiveResponseWithDefaults() *HistoryServiceLevelObjectiveResponse`

NewHistoryServiceLevelObjectiveResponseWithDefaults instantiates a new HistoryServiceLevelObjectiveResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *HistoryServiceLevelObjectiveResponse) GetData() HistoryServiceLevelObjectiveResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *HistoryServiceLevelObjectiveResponse) GetDataOk() (HistoryServiceLevelObjectiveResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasData

`func (o *HistoryServiceLevelObjectiveResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### SetData

`func (o *HistoryServiceLevelObjectiveResponse) SetData(v HistoryServiceLevelObjectiveResponseData)`

SetData gets a reference to the given HistoryServiceLevelObjectiveResponseData and assigns it to the Data field.

### GetErrors

`func (o *HistoryServiceLevelObjectiveResponse) GetErrors() []HistoryServiceLevelObjectiveResponseError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *HistoryServiceLevelObjectiveResponse) GetErrorsOk() ([]HistoryServiceLevelObjectiveResponseError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasErrors

`func (o *HistoryServiceLevelObjectiveResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### SetErrors

`func (o *HistoryServiceLevelObjectiveResponse) SetErrors(v []HistoryServiceLevelObjectiveResponseError)`

SetErrors gets a reference to the given []HistoryServiceLevelObjectiveResponseError and assigns it to the Errors field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


