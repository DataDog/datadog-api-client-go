# LogsRetentionSumUsage

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**LogsIndexedLogsUsageSum** | Pointer to **int64** | Total indexed logs for this retention period. | [optional] 
**LogsLiveIndexedLogsUsageSum** | Pointer to **int64** | Live indexed logs for this retention period. | [optional] 
**LogsRehydratedIndexedLogsUsageSum** | Pointer to **int64** | Rehydrated indexed logs for this retention period. | [optional] 
**Retention** | Pointer to **string** | The retention period in days or \&quot;custom\&quot; for all custom retention periods. | [optional] 

## Methods

### NewLogsRetentionSumUsage

`func NewLogsRetentionSumUsage() *LogsRetentionSumUsage`

NewLogsRetentionSumUsage instantiates a new LogsRetentionSumUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogsRetentionSumUsageWithDefaults

`func NewLogsRetentionSumUsageWithDefaults() *LogsRetentionSumUsage`

NewLogsRetentionSumUsageWithDefaults instantiates a new LogsRetentionSumUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogsIndexedLogsUsageSum

`func (o *LogsRetentionSumUsage) GetLogsIndexedLogsUsageSum() int64`

GetLogsIndexedLogsUsageSum returns the LogsIndexedLogsUsageSum field if non-nil, zero value otherwise.

### GetLogsIndexedLogsUsageSumOk

`func (o *LogsRetentionSumUsage) GetLogsIndexedLogsUsageSumOk() (*int64, bool)`

GetLogsIndexedLogsUsageSumOk returns a tuple with the LogsIndexedLogsUsageSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogsIndexedLogsUsageSum

`func (o *LogsRetentionSumUsage) SetLogsIndexedLogsUsageSum(v int64)`

SetLogsIndexedLogsUsageSum sets LogsIndexedLogsUsageSum field to given value.

### HasLogsIndexedLogsUsageSum

`func (o *LogsRetentionSumUsage) HasLogsIndexedLogsUsageSum() bool`

HasLogsIndexedLogsUsageSum returns a boolean if a field has been set.

### GetLogsLiveIndexedLogsUsageSum

`func (o *LogsRetentionSumUsage) GetLogsLiveIndexedLogsUsageSum() int64`

GetLogsLiveIndexedLogsUsageSum returns the LogsLiveIndexedLogsUsageSum field if non-nil, zero value otherwise.

### GetLogsLiveIndexedLogsUsageSumOk

`func (o *LogsRetentionSumUsage) GetLogsLiveIndexedLogsUsageSumOk() (*int64, bool)`

GetLogsLiveIndexedLogsUsageSumOk returns a tuple with the LogsLiveIndexedLogsUsageSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogsLiveIndexedLogsUsageSum

`func (o *LogsRetentionSumUsage) SetLogsLiveIndexedLogsUsageSum(v int64)`

SetLogsLiveIndexedLogsUsageSum sets LogsLiveIndexedLogsUsageSum field to given value.

### HasLogsLiveIndexedLogsUsageSum

`func (o *LogsRetentionSumUsage) HasLogsLiveIndexedLogsUsageSum() bool`

HasLogsLiveIndexedLogsUsageSum returns a boolean if a field has been set.

### GetLogsRehydratedIndexedLogsUsageSum

`func (o *LogsRetentionSumUsage) GetLogsRehydratedIndexedLogsUsageSum() int64`

GetLogsRehydratedIndexedLogsUsageSum returns the LogsRehydratedIndexedLogsUsageSum field if non-nil, zero value otherwise.

### GetLogsRehydratedIndexedLogsUsageSumOk

`func (o *LogsRetentionSumUsage) GetLogsRehydratedIndexedLogsUsageSumOk() (*int64, bool)`

GetLogsRehydratedIndexedLogsUsageSumOk returns a tuple with the LogsRehydratedIndexedLogsUsageSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogsRehydratedIndexedLogsUsageSum

`func (o *LogsRetentionSumUsage) SetLogsRehydratedIndexedLogsUsageSum(v int64)`

SetLogsRehydratedIndexedLogsUsageSum sets LogsRehydratedIndexedLogsUsageSum field to given value.

### HasLogsRehydratedIndexedLogsUsageSum

`func (o *LogsRetentionSumUsage) HasLogsRehydratedIndexedLogsUsageSum() bool`

HasLogsRehydratedIndexedLogsUsageSum returns a boolean if a field has been set.

### GetRetention

`func (o *LogsRetentionSumUsage) GetRetention() string`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *LogsRetentionSumUsage) GetRetentionOk() (*string, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *LogsRetentionSumUsage) SetRetention(v string)`

SetRetention sets Retention field to given value.

### HasRetention

`func (o *LogsRetentionSumUsage) HasRetention() bool`

HasRetention returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


