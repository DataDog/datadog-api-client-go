# MonitorStateGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastDataTs** | Pointer to **int64** |  | [optional] 
**LastNodataTs** | Pointer to **int64** |  | [optional] 
**LastNotifiedTs** | Pointer to **int64** |  | [optional] 
**LastResolvedTs** | Pointer to **int64** |  | [optional] 
**LastTriggeredTs** | Pointer to **int64** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**MonitorOverallStates**](MonitorOverallStates.md) |  | [optional] 
**TriggeringValue** | Pointer to [**MonitorStateGroupValue**](MonitorStateGroupValue.md) |  | [optional] 

## Methods

### NewMonitorStateGroup

`func NewMonitorStateGroup() *MonitorStateGroup`

NewMonitorStateGroup instantiates a new MonitorStateGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorStateGroupWithDefaults

`func NewMonitorStateGroupWithDefaults() *MonitorStateGroup`

NewMonitorStateGroupWithDefaults instantiates a new MonitorStateGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastDataTs

`func (o *MonitorStateGroup) GetLastDataTs() int64`

GetLastDataTs returns the LastDataTs field if non-nil, zero value otherwise.

### GetLastDataTsOk

`func (o *MonitorStateGroup) GetLastDataTsOk() (int64, bool)`

GetLastDataTsOk returns a tuple with the LastDataTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastDataTs

`func (o *MonitorStateGroup) HasLastDataTs() bool`

HasLastDataTs returns a boolean if a field has been set.

### SetLastDataTs

`func (o *MonitorStateGroup) SetLastDataTs(v int64)`

SetLastDataTs gets a reference to the given int64 and assigns it to the LastDataTs field.

### GetLastNodataTs

`func (o *MonitorStateGroup) GetLastNodataTs() int64`

GetLastNodataTs returns the LastNodataTs field if non-nil, zero value otherwise.

### GetLastNodataTsOk

`func (o *MonitorStateGroup) GetLastNodataTsOk() (int64, bool)`

GetLastNodataTsOk returns a tuple with the LastNodataTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastNodataTs

`func (o *MonitorStateGroup) HasLastNodataTs() bool`

HasLastNodataTs returns a boolean if a field has been set.

### SetLastNodataTs

`func (o *MonitorStateGroup) SetLastNodataTs(v int64)`

SetLastNodataTs gets a reference to the given int64 and assigns it to the LastNodataTs field.

### GetLastNotifiedTs

`func (o *MonitorStateGroup) GetLastNotifiedTs() int64`

GetLastNotifiedTs returns the LastNotifiedTs field if non-nil, zero value otherwise.

### GetLastNotifiedTsOk

`func (o *MonitorStateGroup) GetLastNotifiedTsOk() (int64, bool)`

GetLastNotifiedTsOk returns a tuple with the LastNotifiedTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastNotifiedTs

`func (o *MonitorStateGroup) HasLastNotifiedTs() bool`

HasLastNotifiedTs returns a boolean if a field has been set.

### SetLastNotifiedTs

`func (o *MonitorStateGroup) SetLastNotifiedTs(v int64)`

SetLastNotifiedTs gets a reference to the given int64 and assigns it to the LastNotifiedTs field.

### GetLastResolvedTs

`func (o *MonitorStateGroup) GetLastResolvedTs() int64`

GetLastResolvedTs returns the LastResolvedTs field if non-nil, zero value otherwise.

### GetLastResolvedTsOk

`func (o *MonitorStateGroup) GetLastResolvedTsOk() (int64, bool)`

GetLastResolvedTsOk returns a tuple with the LastResolvedTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastResolvedTs

`func (o *MonitorStateGroup) HasLastResolvedTs() bool`

HasLastResolvedTs returns a boolean if a field has been set.

### SetLastResolvedTs

`func (o *MonitorStateGroup) SetLastResolvedTs(v int64)`

SetLastResolvedTs gets a reference to the given int64 and assigns it to the LastResolvedTs field.

### GetLastTriggeredTs

`func (o *MonitorStateGroup) GetLastTriggeredTs() int64`

GetLastTriggeredTs returns the LastTriggeredTs field if non-nil, zero value otherwise.

### GetLastTriggeredTsOk

`func (o *MonitorStateGroup) GetLastTriggeredTsOk() (int64, bool)`

GetLastTriggeredTsOk returns a tuple with the LastTriggeredTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastTriggeredTs

`func (o *MonitorStateGroup) HasLastTriggeredTs() bool`

HasLastTriggeredTs returns a boolean if a field has been set.

### SetLastTriggeredTs

`func (o *MonitorStateGroup) SetLastTriggeredTs(v int64)`

SetLastTriggeredTs gets a reference to the given int64 and assigns it to the LastTriggeredTs field.

### GetMessage

`func (o *MonitorStateGroup) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MonitorStateGroup) GetMessageOk() (string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMessage

`func (o *MonitorStateGroup) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessage

`func (o *MonitorStateGroup) SetMessage(v string)`

SetMessage gets a reference to the given string and assigns it to the Message field.

### GetName

`func (o *MonitorStateGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MonitorStateGroup) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *MonitorStateGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *MonitorStateGroup) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetStatus

`func (o *MonitorStateGroup) GetStatus() MonitorOverallStates`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MonitorStateGroup) GetStatusOk() (MonitorOverallStates, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStatus

`func (o *MonitorStateGroup) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatus

`func (o *MonitorStateGroup) SetStatus(v MonitorOverallStates)`

SetStatus gets a reference to the given MonitorOverallStates and assigns it to the Status field.

### GetTriggeringValue

`func (o *MonitorStateGroup) GetTriggeringValue() MonitorStateGroupValue`

GetTriggeringValue returns the TriggeringValue field if non-nil, zero value otherwise.

### GetTriggeringValueOk

`func (o *MonitorStateGroup) GetTriggeringValueOk() (MonitorStateGroupValue, bool)`

GetTriggeringValueOk returns a tuple with the TriggeringValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTriggeringValue

`func (o *MonitorStateGroup) HasTriggeringValue() bool`

HasTriggeringValue returns a boolean if a field has been set.

### SetTriggeringValue

`func (o *MonitorStateGroup) SetTriggeringValue(v MonitorStateGroupValue)`

SetTriggeringValue gets a reference to the given MonitorStateGroupValue and assigns it to the TriggeringValue field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


