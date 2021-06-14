# MonitorGroupSearchResult

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Group** | Pointer to **string** | The name of the group. | [optional] [readonly] 
**GroupTags** | Pointer to **[]string** | The list of tags of the monitor group. | [optional] [readonly] 
**LastNodataTs** | Pointer to **int64** | Latest timestamp the monitor group was in NO_DATA state. | [optional] [readonly] 
**LastTriggeredTs** | Pointer to **NullableInt64** | Latest timestamp the monitor group triggered. | [optional] [readonly] 
**MonitorId** | Pointer to **int64** | The ID of the monitor. | [optional] [readonly] 
**MonitorName** | Pointer to **string** | The name of the monitor. | [optional] [readonly] 
**Status** | Pointer to [**MonitorOverallStates**](MonitorOverallStates.md) |  | [optional] 

## Methods

### NewMonitorGroupSearchResult

`func NewMonitorGroupSearchResult() *MonitorGroupSearchResult`

NewMonitorGroupSearchResult instantiates a new MonitorGroupSearchResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorGroupSearchResultWithDefaults

`func NewMonitorGroupSearchResultWithDefaults() *MonitorGroupSearchResult`

NewMonitorGroupSearchResultWithDefaults instantiates a new MonitorGroupSearchResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *MonitorGroupSearchResult) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *MonitorGroupSearchResult) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *MonitorGroupSearchResult) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *MonitorGroupSearchResult) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetGroupTags

`func (o *MonitorGroupSearchResult) GetGroupTags() []string`

GetGroupTags returns the GroupTags field if non-nil, zero value otherwise.

### GetGroupTagsOk

`func (o *MonitorGroupSearchResult) GetGroupTagsOk() (*[]string, bool)`

GetGroupTagsOk returns a tuple with the GroupTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupTags

`func (o *MonitorGroupSearchResult) SetGroupTags(v []string)`

SetGroupTags sets GroupTags field to given value.

### HasGroupTags

`func (o *MonitorGroupSearchResult) HasGroupTags() bool`

HasGroupTags returns a boolean if a field has been set.

### GetLastNodataTs

`func (o *MonitorGroupSearchResult) GetLastNodataTs() int64`

GetLastNodataTs returns the LastNodataTs field if non-nil, zero value otherwise.

### GetLastNodataTsOk

`func (o *MonitorGroupSearchResult) GetLastNodataTsOk() (*int64, bool)`

GetLastNodataTsOk returns a tuple with the LastNodataTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastNodataTs

`func (o *MonitorGroupSearchResult) SetLastNodataTs(v int64)`

SetLastNodataTs sets LastNodataTs field to given value.

### HasLastNodataTs

`func (o *MonitorGroupSearchResult) HasLastNodataTs() bool`

HasLastNodataTs returns a boolean if a field has been set.

### GetLastTriggeredTs

`func (o *MonitorGroupSearchResult) GetLastTriggeredTs() int64`

GetLastTriggeredTs returns the LastTriggeredTs field if non-nil, zero value otherwise.

### GetLastTriggeredTsOk

`func (o *MonitorGroupSearchResult) GetLastTriggeredTsOk() (*int64, bool)`

GetLastTriggeredTsOk returns a tuple with the LastTriggeredTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTriggeredTs

`func (o *MonitorGroupSearchResult) SetLastTriggeredTs(v int64)`

SetLastTriggeredTs sets LastTriggeredTs field to given value.

### HasLastTriggeredTs

`func (o *MonitorGroupSearchResult) HasLastTriggeredTs() bool`

HasLastTriggeredTs returns a boolean if a field has been set.

### SetLastTriggeredTsNil

`func (o *MonitorGroupSearchResult) SetLastTriggeredTsNil(b bool)`

 SetLastTriggeredTsNil sets the value for LastTriggeredTs to be an explicit nil

### UnsetLastTriggeredTs
`func (o *MonitorGroupSearchResult) UnsetLastTriggeredTs()`

UnsetLastTriggeredTs ensures that no value is present for LastTriggeredTs, not even an explicit nil
### GetMonitorId

`func (o *MonitorGroupSearchResult) GetMonitorId() int64`

GetMonitorId returns the MonitorId field if non-nil, zero value otherwise.

### GetMonitorIdOk

`func (o *MonitorGroupSearchResult) GetMonitorIdOk() (*int64, bool)`

GetMonitorIdOk returns a tuple with the MonitorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorId

`func (o *MonitorGroupSearchResult) SetMonitorId(v int64)`

SetMonitorId sets MonitorId field to given value.

### HasMonitorId

`func (o *MonitorGroupSearchResult) HasMonitorId() bool`

HasMonitorId returns a boolean if a field has been set.

### GetMonitorName

`func (o *MonitorGroupSearchResult) GetMonitorName() string`

GetMonitorName returns the MonitorName field if non-nil, zero value otherwise.

### GetMonitorNameOk

`func (o *MonitorGroupSearchResult) GetMonitorNameOk() (*string, bool)`

GetMonitorNameOk returns a tuple with the MonitorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorName

`func (o *MonitorGroupSearchResult) SetMonitorName(v string)`

SetMonitorName sets MonitorName field to given value.

### HasMonitorName

`func (o *MonitorGroupSearchResult) HasMonitorName() bool`

HasMonitorName returns a boolean if a field has been set.

### GetStatus

`func (o *MonitorGroupSearchResult) GetStatus() MonitorOverallStates`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MonitorGroupSearchResult) GetStatusOk() (*MonitorOverallStates, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MonitorGroupSearchResult) SetStatus(v MonitorOverallStates)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MonitorGroupSearchResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


