# MonitorStateGroup

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**LastNodataTs** | Pointer to **int64** | Latest timestamp the monitor was in NO_DATA state. | [optional] 
**LastNotifiedTs** | Pointer to **int64** | Latest timestamp of the notification sent for this monitor group. | [optional] 
**LastResolvedTs** | Pointer to **int64** | Latest timestamp the monitor group was resolved. | [optional] 
**LastTriggeredTs** | Pointer to **int64** | Latest timestamp the monitor group triggered. | [optional] 
**Name** | Pointer to **string** | The name of the monitor. | [optional] 
**Status** | Pointer to [**MonitorOverallStates**](MonitorOverallStates.md) |  | [optional] 

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

### GetLastNodataTs

`func (o *MonitorStateGroup) GetLastNodataTs() int64`

GetLastNodataTs returns the LastNodataTs field if non-nil, zero value otherwise.

### GetLastNodataTsOk

`func (o *MonitorStateGroup) GetLastNodataTsOk() (*int64, bool)`

GetLastNodataTsOk returns a tuple with the LastNodataTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastNodataTs

`func (o *MonitorStateGroup) SetLastNodataTs(v int64)`

SetLastNodataTs sets LastNodataTs field to given value.

### HasLastNodataTs

`func (o *MonitorStateGroup) HasLastNodataTs() bool`

HasLastNodataTs returns a boolean if a field has been set.

### GetLastNotifiedTs

`func (o *MonitorStateGroup) GetLastNotifiedTs() int64`

GetLastNotifiedTs returns the LastNotifiedTs field if non-nil, zero value otherwise.

### GetLastNotifiedTsOk

`func (o *MonitorStateGroup) GetLastNotifiedTsOk() (*int64, bool)`

GetLastNotifiedTsOk returns a tuple with the LastNotifiedTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastNotifiedTs

`func (o *MonitorStateGroup) SetLastNotifiedTs(v int64)`

SetLastNotifiedTs sets LastNotifiedTs field to given value.

### HasLastNotifiedTs

`func (o *MonitorStateGroup) HasLastNotifiedTs() bool`

HasLastNotifiedTs returns a boolean if a field has been set.

### GetLastResolvedTs

`func (o *MonitorStateGroup) GetLastResolvedTs() int64`

GetLastResolvedTs returns the LastResolvedTs field if non-nil, zero value otherwise.

### GetLastResolvedTsOk

`func (o *MonitorStateGroup) GetLastResolvedTsOk() (*int64, bool)`

GetLastResolvedTsOk returns a tuple with the LastResolvedTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastResolvedTs

`func (o *MonitorStateGroup) SetLastResolvedTs(v int64)`

SetLastResolvedTs sets LastResolvedTs field to given value.

### HasLastResolvedTs

`func (o *MonitorStateGroup) HasLastResolvedTs() bool`

HasLastResolvedTs returns a boolean if a field has been set.

### GetLastTriggeredTs

`func (o *MonitorStateGroup) GetLastTriggeredTs() int64`

GetLastTriggeredTs returns the LastTriggeredTs field if non-nil, zero value otherwise.

### GetLastTriggeredTsOk

`func (o *MonitorStateGroup) GetLastTriggeredTsOk() (*int64, bool)`

GetLastTriggeredTsOk returns a tuple with the LastTriggeredTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTriggeredTs

`func (o *MonitorStateGroup) SetLastTriggeredTs(v int64)`

SetLastTriggeredTs sets LastTriggeredTs field to given value.

### HasLastTriggeredTs

`func (o *MonitorStateGroup) HasLastTriggeredTs() bool`

HasLastTriggeredTs returns a boolean if a field has been set.

### GetName

`func (o *MonitorStateGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MonitorStateGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MonitorStateGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MonitorStateGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *MonitorStateGroup) GetStatus() MonitorOverallStates`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MonitorStateGroup) GetStatusOk() (*MonitorOverallStates, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MonitorStateGroup) SetStatus(v MonitorOverallStates)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MonitorStateGroup) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


