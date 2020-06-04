# LogsArchiveAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destination** | Pointer to [**NullableLogsArchiveDestination**](LogsArchiveDestination.md) |  | 
**Name** | Pointer to **string** | The archive name. | 
**Query** | Pointer to **string** | The archive query/filter. Logs matching this query are included in the archive. | 
**State** | Pointer to [**LogsArchiveState**](LogsArchiveState.md) |  | [optional] 

## Methods

### NewLogsArchiveAttributes

`func NewLogsArchiveAttributes(destination NullableLogsArchiveDestination, name string, query string, ) *LogsArchiveAttributes`

NewLogsArchiveAttributes instantiates a new LogsArchiveAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogsArchiveAttributesWithDefaults

`func NewLogsArchiveAttributesWithDefaults() *LogsArchiveAttributes`

NewLogsArchiveAttributesWithDefaults instantiates a new LogsArchiveAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestination

`func (o *LogsArchiveAttributes) GetDestination() LogsArchiveDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *LogsArchiveAttributes) GetDestinationOk() (*LogsArchiveDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *LogsArchiveAttributes) SetDestination(v LogsArchiveDestination)`

SetDestination sets Destination field to given value.


### SetDestinationNil

`func (o *LogsArchiveAttributes) SetDestinationNil(b bool)`

 SetDestinationNil sets the value for Destination to be an explicit nil

### UnsetDestination
`func (o *LogsArchiveAttributes) UnsetDestination()`

UnsetDestination ensures that no value is present for Destination, not even an explicit nil
### GetName

`func (o *LogsArchiveAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LogsArchiveAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LogsArchiveAttributes) SetName(v string)`

SetName sets Name field to given value.


### GetQuery

`func (o *LogsArchiveAttributes) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *LogsArchiveAttributes) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *LogsArchiveAttributes) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetState

`func (o *LogsArchiveAttributes) GetState() LogsArchiveState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *LogsArchiveAttributes) GetStateOk() (*LogsArchiveState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *LogsArchiveAttributes) SetState(v LogsArchiveState)`

SetState sets State field to given value.

### HasState

`func (o *LogsArchiveAttributes) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


