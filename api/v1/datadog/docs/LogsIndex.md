# LogsIndex

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DailyLimit** | Pointer to **int64** | The number of log events you can send in this index per day before you are rate-limited. | [optional] 
**ExclusionFilters** | Pointer to [**[]LogsExclusion**](LogsExclusion.md) | An array of exclusion objects. The logs are tested against the query of each filter, following the order of the array. Only the first matching active exclusion matters, others (if any) are ignored. | [optional] 
**Filter** | [**LogsFilter**](LogsFilter.md) |  | 
**IsRateLimited** | Pointer to **bool** | A boolean stating if the index is rate limited, meaning more logs than the daily limit have been sent. Rate limit is reset every-day at 2pm UTC. | [optional] [readonly] 
**Name** | **string** | The name of the index. | 
**NumRetentionDays** | Pointer to **int64** | The number of days before logs are deleted from this index. Available values depend on retention plans specified in your organization&#39;s contract/subscriptions. | [optional] 

## Methods

### NewLogsIndex

`func NewLogsIndex(filter LogsFilter, name string, ) *LogsIndex`

NewLogsIndex instantiates a new LogsIndex object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogsIndexWithDefaults

`func NewLogsIndexWithDefaults() *LogsIndex`

NewLogsIndexWithDefaults instantiates a new LogsIndex object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDailyLimit

`func (o *LogsIndex) GetDailyLimit() int64`

GetDailyLimit returns the DailyLimit field if non-nil, zero value otherwise.

### GetDailyLimitOk

`func (o *LogsIndex) GetDailyLimitOk() (*int64, bool)`

GetDailyLimitOk returns a tuple with the DailyLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyLimit

`func (o *LogsIndex) SetDailyLimit(v int64)`

SetDailyLimit sets DailyLimit field to given value.

### HasDailyLimit

`func (o *LogsIndex) HasDailyLimit() bool`

HasDailyLimit returns a boolean if a field has been set.

### GetExclusionFilters

`func (o *LogsIndex) GetExclusionFilters() []LogsExclusion`

GetExclusionFilters returns the ExclusionFilters field if non-nil, zero value otherwise.

### GetExclusionFiltersOk

`func (o *LogsIndex) GetExclusionFiltersOk() (*[]LogsExclusion, bool)`

GetExclusionFiltersOk returns a tuple with the ExclusionFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusionFilters

`func (o *LogsIndex) SetExclusionFilters(v []LogsExclusion)`

SetExclusionFilters sets ExclusionFilters field to given value.

### HasExclusionFilters

`func (o *LogsIndex) HasExclusionFilters() bool`

HasExclusionFilters returns a boolean if a field has been set.

### GetFilter

`func (o *LogsIndex) GetFilter() LogsFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *LogsIndex) GetFilterOk() (*LogsFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *LogsIndex) SetFilter(v LogsFilter)`

SetFilter sets Filter field to given value.


### GetIsRateLimited

`func (o *LogsIndex) GetIsRateLimited() bool`

GetIsRateLimited returns the IsRateLimited field if non-nil, zero value otherwise.

### GetIsRateLimitedOk

`func (o *LogsIndex) GetIsRateLimitedOk() (*bool, bool)`

GetIsRateLimitedOk returns a tuple with the IsRateLimited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRateLimited

`func (o *LogsIndex) SetIsRateLimited(v bool)`

SetIsRateLimited sets IsRateLimited field to given value.

### HasIsRateLimited

`func (o *LogsIndex) HasIsRateLimited() bool`

HasIsRateLimited returns a boolean if a field has been set.

### GetName

`func (o *LogsIndex) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LogsIndex) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LogsIndex) SetName(v string)`

SetName sets Name field to given value.


### GetNumRetentionDays

`func (o *LogsIndex) GetNumRetentionDays() int64`

GetNumRetentionDays returns the NumRetentionDays field if non-nil, zero value otherwise.

### GetNumRetentionDaysOk

`func (o *LogsIndex) GetNumRetentionDaysOk() (*int64, bool)`

GetNumRetentionDaysOk returns a tuple with the NumRetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumRetentionDays

`func (o *LogsIndex) SetNumRetentionDays(v int64)`

SetNumRetentionDays sets NumRetentionDays field to given value.

### HasNumRetentionDays

`func (o *LogsIndex) HasNumRetentionDays() bool`

HasNumRetentionDays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


