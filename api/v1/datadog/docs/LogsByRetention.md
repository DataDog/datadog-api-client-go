# LogsByRetention

## Properties

| Name             | Type                                                                         | Description                                                       | Notes      |
| ---------------- | ---------------------------------------------------------------------------- | ----------------------------------------------------------------- | ---------- |
| **Orgs**         | Pointer to [**LogsByRetentionOrgs**](LogsByRetentionOrgs.md)                 |                                                                   | [optional] |
| **Usage**        | Pointer to [**[]LogsRetentionAggSumUsage**](LogsRetentionAggSumUsage.md)     | Aggregated index logs usage for each retention period with usage. | [optional] |
| **UsageByMonth** | Pointer to [**LogsByRetentionMonthlyUsage**](LogsByRetentionMonthlyUsage.md) |                                                                   | [optional] |

## Methods

### NewLogsByRetention

`func NewLogsByRetention() *LogsByRetention`

NewLogsByRetention instantiates a new LogsByRetention object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewLogsByRetentionWithDefaults

`func NewLogsByRetentionWithDefaults() *LogsByRetention`

NewLogsByRetentionWithDefaults instantiates a new LogsByRetention object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetOrgs

`func (o *LogsByRetention) GetOrgs() LogsByRetentionOrgs`

GetOrgs returns the Orgs field if non-nil, zero value otherwise.

### GetOrgsOk

`func (o *LogsByRetention) GetOrgsOk() (*LogsByRetentionOrgs, bool)`

GetOrgsOk returns a tuple with the Orgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgs

`func (o *LogsByRetention) SetOrgs(v LogsByRetentionOrgs)`

SetOrgs sets Orgs field to given value.

### HasOrgs

`func (o *LogsByRetention) HasOrgs() bool`

HasOrgs returns a boolean if a field has been set.

### GetUsage

`func (o *LogsByRetention) GetUsage() []LogsRetentionAggSumUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *LogsByRetention) GetUsageOk() (*[]LogsRetentionAggSumUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *LogsByRetention) SetUsage(v []LogsRetentionAggSumUsage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *LogsByRetention) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetUsageByMonth

`func (o *LogsByRetention) GetUsageByMonth() LogsByRetentionMonthlyUsage`

GetUsageByMonth returns the UsageByMonth field if non-nil, zero value otherwise.

### GetUsageByMonthOk

`func (o *LogsByRetention) GetUsageByMonthOk() (*LogsByRetentionMonthlyUsage, bool)`

GetUsageByMonthOk returns a tuple with the UsageByMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageByMonth

`func (o *LogsByRetention) SetUsageByMonth(v LogsByRetentionMonthlyUsage)`

SetUsageByMonth sets UsageByMonth field to given value.

### HasUsageByMonth

`func (o *LogsByRetention) HasUsageByMonth() bool`

HasUsageByMonth returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
