# LogsRetentionAggSumUsage

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**LogsIndexedLogsUsageAggSum** | Pointer to **int64** | Total indexed logs for this retention period. | [optional] 
**LogsLiveIndexedLogsUsageAggSum** | Pointer to **int64** | Live indexed logs for this retention period. | [optional] 
**LogsRehydratedIndexedLogsUsageAggSum** | Pointer to **int64** | Rehydrated indexed logs for this retention period. | [optional] 
**Retention** | Pointer to **string** | The retention period in days or \&quot;custom\&quot; for all custom retention periods. | [optional] 

## Methods

### NewLogsRetentionAggSumUsage

`func NewLogsRetentionAggSumUsage() *LogsRetentionAggSumUsage`

NewLogsRetentionAggSumUsage instantiates a new LogsRetentionAggSumUsage object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewLogsRetentionAggSumUsageWithDefaults

`func NewLogsRetentionAggSumUsageWithDefaults() *LogsRetentionAggSumUsage`

NewLogsRetentionAggSumUsageWithDefaults instantiates a new LogsRetentionAggSumUsage object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetLogsIndexedLogsUsageAggSum

`func (o *LogsRetentionAggSumUsage) GetLogsIndexedLogsUsageAggSum() int64`

GetLogsIndexedLogsUsageAggSum returns the LogsIndexedLogsUsageAggSum field if non-nil, zero value otherwise.

### GetLogsIndexedLogsUsageAggSumOk

`func (o *LogsRetentionAggSumUsage) GetLogsIndexedLogsUsageAggSumOk() (*int64, bool)`

GetLogsIndexedLogsUsageAggSumOk returns a tuple with the LogsIndexedLogsUsageAggSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogsIndexedLogsUsageAggSum

`func (o *LogsRetentionAggSumUsage) SetLogsIndexedLogsUsageAggSum(v int64)`

SetLogsIndexedLogsUsageAggSum sets LogsIndexedLogsUsageAggSum field to given value.

### HasLogsIndexedLogsUsageAggSum

`func (o *LogsRetentionAggSumUsage) HasLogsIndexedLogsUsageAggSum() bool`

HasLogsIndexedLogsUsageAggSum returns a boolean if a field has been set.

### GetLogsLiveIndexedLogsUsageAggSum

`func (o *LogsRetentionAggSumUsage) GetLogsLiveIndexedLogsUsageAggSum() int64`

GetLogsLiveIndexedLogsUsageAggSum returns the LogsLiveIndexedLogsUsageAggSum field if non-nil, zero value otherwise.

### GetLogsLiveIndexedLogsUsageAggSumOk

`func (o *LogsRetentionAggSumUsage) GetLogsLiveIndexedLogsUsageAggSumOk() (*int64, bool)`

GetLogsLiveIndexedLogsUsageAggSumOk returns a tuple with the LogsLiveIndexedLogsUsageAggSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogsLiveIndexedLogsUsageAggSum

`func (o *LogsRetentionAggSumUsage) SetLogsLiveIndexedLogsUsageAggSum(v int64)`

SetLogsLiveIndexedLogsUsageAggSum sets LogsLiveIndexedLogsUsageAggSum field to given value.

### HasLogsLiveIndexedLogsUsageAggSum

`func (o *LogsRetentionAggSumUsage) HasLogsLiveIndexedLogsUsageAggSum() bool`

HasLogsLiveIndexedLogsUsageAggSum returns a boolean if a field has been set.

### GetLogsRehydratedIndexedLogsUsageAggSum

`func (o *LogsRetentionAggSumUsage) GetLogsRehydratedIndexedLogsUsageAggSum() int64`

GetLogsRehydratedIndexedLogsUsageAggSum returns the LogsRehydratedIndexedLogsUsageAggSum field if non-nil, zero value otherwise.

### GetLogsRehydratedIndexedLogsUsageAggSumOk

`func (o *LogsRetentionAggSumUsage) GetLogsRehydratedIndexedLogsUsageAggSumOk() (*int64, bool)`

GetLogsRehydratedIndexedLogsUsageAggSumOk returns a tuple with the LogsRehydratedIndexedLogsUsageAggSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogsRehydratedIndexedLogsUsageAggSum

`func (o *LogsRetentionAggSumUsage) SetLogsRehydratedIndexedLogsUsageAggSum(v int64)`

SetLogsRehydratedIndexedLogsUsageAggSum sets LogsRehydratedIndexedLogsUsageAggSum field to given value.

### HasLogsRehydratedIndexedLogsUsageAggSum

`func (o *LogsRetentionAggSumUsage) HasLogsRehydratedIndexedLogsUsageAggSum() bool`

HasLogsRehydratedIndexedLogsUsageAggSum returns a boolean if a field has been set.

### GetRetention

`func (o *LogsRetentionAggSumUsage) GetRetention() string`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *LogsRetentionAggSumUsage) GetRetentionOk() (*string, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *LogsRetentionAggSumUsage) SetRetention(v string)`

SetRetention sets Retention field to given value.

### HasRetention

`func (o *LogsRetentionAggSumUsage) HasRetention() bool`

HasRetention returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


