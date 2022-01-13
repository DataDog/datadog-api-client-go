# LogsIndexUpdateRequest

## Properties

| Name                  | Type                                               | Description                                                                                                                                                                                                                                                                                                       | Notes      |
| --------------------- | -------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------- |
| **DailyLimit**        | Pointer to **int64**                               | The number of log events you can send in this index per day before you are rate-limited.                                                                                                                                                                                                                          | [optional] |
| **DisableDailyLimit** | Pointer to **bool**                                | If true, sets the &#x60;daily_limit&#x60; value to null and the index is not limited on a daily basis (any specified &#x60;daily_limit&#x60; value in the request is ignored). If false or omitted, the index&#39;s current &#x60;daily_limit&#x60; is maintained.                                                | [optional] |
| **ExclusionFilters**  | Pointer to [**[]LogsExclusion**](LogsExclusion.md) | An array of exclusion objects. The logs are tested against the query of each filter, following the order of the array. Only the first matching active exclusion matters, others (if any) are ignored.                                                                                                             | [optional] |
| **Filter**            | [**LogsFilter**](LogsFilter.md)                    |                                                                                                                                                                                                                                                                                                                   |
| **NumRetentionDays**  | Pointer to **int64**                               | The number of days before logs are deleted from this index. Available values depend on retention plans specified in your organization&#39;s contract/subscriptions. **Note:** Changing the retention for an index adjusts the length of retention for all logs already in this index. It may also affect billing. | [optional] |

## Methods

### NewLogsIndexUpdateRequest

`func NewLogsIndexUpdateRequest(filter LogsFilter) *LogsIndexUpdateRequest`

NewLogsIndexUpdateRequest instantiates a new LogsIndexUpdateRequest object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewLogsIndexUpdateRequestWithDefaults

`func NewLogsIndexUpdateRequestWithDefaults() *LogsIndexUpdateRequest`

NewLogsIndexUpdateRequestWithDefaults instantiates a new LogsIndexUpdateRequest object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetDailyLimit

`func (o *LogsIndexUpdateRequest) GetDailyLimit() int64`

GetDailyLimit returns the DailyLimit field if non-nil, zero value otherwise.

### GetDailyLimitOk

`func (o *LogsIndexUpdateRequest) GetDailyLimitOk() (*int64, bool)`

GetDailyLimitOk returns a tuple with the DailyLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyLimit

`func (o *LogsIndexUpdateRequest) SetDailyLimit(v int64)`

SetDailyLimit sets DailyLimit field to given value.

### HasDailyLimit

`func (o *LogsIndexUpdateRequest) HasDailyLimit() bool`

HasDailyLimit returns a boolean if a field has been set.

### GetDisableDailyLimit

`func (o *LogsIndexUpdateRequest) GetDisableDailyLimit() bool`

GetDisableDailyLimit returns the DisableDailyLimit field if non-nil, zero value otherwise.

### GetDisableDailyLimitOk

`func (o *LogsIndexUpdateRequest) GetDisableDailyLimitOk() (*bool, bool)`

GetDisableDailyLimitOk returns a tuple with the DisableDailyLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableDailyLimit

`func (o *LogsIndexUpdateRequest) SetDisableDailyLimit(v bool)`

SetDisableDailyLimit sets DisableDailyLimit field to given value.

### HasDisableDailyLimit

`func (o *LogsIndexUpdateRequest) HasDisableDailyLimit() bool`

HasDisableDailyLimit returns a boolean if a field has been set.

### GetExclusionFilters

`func (o *LogsIndexUpdateRequest) GetExclusionFilters() []LogsExclusion`

GetExclusionFilters returns the ExclusionFilters field if non-nil, zero value otherwise.

### GetExclusionFiltersOk

`func (o *LogsIndexUpdateRequest) GetExclusionFiltersOk() (*[]LogsExclusion, bool)`

GetExclusionFiltersOk returns a tuple with the ExclusionFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusionFilters

`func (o *LogsIndexUpdateRequest) SetExclusionFilters(v []LogsExclusion)`

SetExclusionFilters sets ExclusionFilters field to given value.

### HasExclusionFilters

`func (o *LogsIndexUpdateRequest) HasExclusionFilters() bool`

HasExclusionFilters returns a boolean if a field has been set.

### GetFilter

`func (o *LogsIndexUpdateRequest) GetFilter() LogsFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *LogsIndexUpdateRequest) GetFilterOk() (*LogsFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *LogsIndexUpdateRequest) SetFilter(v LogsFilter)`

SetFilter sets Filter field to given value.

### GetNumRetentionDays

`func (o *LogsIndexUpdateRequest) GetNumRetentionDays() int64`

GetNumRetentionDays returns the NumRetentionDays field if non-nil, zero value otherwise.

### GetNumRetentionDaysOk

`func (o *LogsIndexUpdateRequest) GetNumRetentionDaysOk() (*int64, bool)`

GetNumRetentionDaysOk returns a tuple with the NumRetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumRetentionDays

`func (o *LogsIndexUpdateRequest) SetNumRetentionDays(v int64)`

SetNumRetentionDays sets NumRetentionDays field to given value.

### HasNumRetentionDays

`func (o *LogsIndexUpdateRequest) HasNumRetentionDays() bool`

HasNumRetentionDays returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
