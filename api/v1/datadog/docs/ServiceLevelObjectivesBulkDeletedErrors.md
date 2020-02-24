# ServiceLevelObjectivesBulkDeletedErrors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the service level objective object associated with this error. | 
**Message** | Pointer to **string** | The error message | 
**Timeframe** | Pointer to [**SLOErrorTimeframe**](SLOErrorTimeframe.md) |  | 

## Methods

### NewServiceLevelObjectivesBulkDeletedErrors

`func NewServiceLevelObjectivesBulkDeletedErrors(id string, message string, timeframe SLOErrorTimeframe, ) *ServiceLevelObjectivesBulkDeletedErrors`

NewServiceLevelObjectivesBulkDeletedErrors instantiates a new ServiceLevelObjectivesBulkDeletedErrors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceLevelObjectivesBulkDeletedErrorsWithDefaults

`func NewServiceLevelObjectivesBulkDeletedErrorsWithDefaults() *ServiceLevelObjectivesBulkDeletedErrors`

NewServiceLevelObjectivesBulkDeletedErrorsWithDefaults instantiates a new ServiceLevelObjectivesBulkDeletedErrors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServiceLevelObjectivesBulkDeletedErrors) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceLevelObjectivesBulkDeletedErrors) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *ServiceLevelObjectivesBulkDeletedErrors) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *ServiceLevelObjectivesBulkDeletedErrors) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetMessage

`func (o *ServiceLevelObjectivesBulkDeletedErrors) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ServiceLevelObjectivesBulkDeletedErrors) GetMessageOk() (string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMessage

`func (o *ServiceLevelObjectivesBulkDeletedErrors) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessage

`func (o *ServiceLevelObjectivesBulkDeletedErrors) SetMessage(v string)`

SetMessage gets a reference to the given string and assigns it to the Message field.

### GetTimeframe

`func (o *ServiceLevelObjectivesBulkDeletedErrors) GetTimeframe() SLOErrorTimeframe`

GetTimeframe returns the Timeframe field if non-nil, zero value otherwise.

### GetTimeframeOk

`func (o *ServiceLevelObjectivesBulkDeletedErrors) GetTimeframeOk() (SLOErrorTimeframe, bool)`

GetTimeframeOk returns a tuple with the Timeframe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTimeframe

`func (o *ServiceLevelObjectivesBulkDeletedErrors) HasTimeframe() bool`

HasTimeframe returns a boolean if a field has been set.

### SetTimeframe

`func (o *ServiceLevelObjectivesBulkDeletedErrors) SetTimeframe(v SLOErrorTimeframe)`

SetTimeframe gets a reference to the given SLOErrorTimeframe and assigns it to the Timeframe field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


