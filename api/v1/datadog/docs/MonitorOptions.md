# MonitorOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregation** | Pointer to [**MonitorOptionsAggregation**](MonitorOptions_aggregation.md) |  | [optional] 
**DeviceIds** | Pointer to [**[]MonitorDeviceID**](MonitorDeviceID.md) | IDs of the device the Synthetics monitor is running on. | [optional] 
**EnableLogsSample** | Pointer to **bool** | Whether or not to send a log sample when the log monitor triggers. | [optional] 
**EscalationMessage** | Pointer to **string** | A message to include with a re-notification. Supports the &#x60;@username&#x60; notification we allow elsewhere. Not applicable if &#x60;renotify_interval&#x60; is &#x60;None&#x60;. | [optional] [default to "none"]
**EvaluationDelay** | Pointer to **NullableInt64** | Time (in seconds) to delay evaluation, as a non-negative integer. For example, if the value is set to &#x60;300&#x60; (5min), the timeframe is set to &#x60;last_5m&#x60; and the time is 7:00, the monitor evaluates data from 6:50 to 6:55. This is useful for AWS CloudWatch and other backfilled metrics to ensure the monitor always has data during evaluation. | [optional] 
**IncludeTags** | Pointer to **bool** | A Boolean indicating whether notifications from this monitor automatically inserts its triggering tags into the title.  **Examples** - If &#x60;True&#x60;, &#x60;[Triggered on {host:h1}] Monitor Title&#x60; - If &#x60;False&#x60;, &#x60;[Triggered] Monitor Title&#x60; | [optional] [default to true]
**Locked** | Pointer to **bool** | Wether or not the monitor is locked (only editable by creator and admins). | [optional] 
**MinFailureDuration** | Pointer to **NullableInt64** | How long the test should be in failure before alerting (integer, number of seconds, max 7200). | [optional] [default to 0]
**MinLocationFailed** | Pointer to **NullableInt64** | The minimum number of locations in failure at the same time during at least one moment in the &#x60;min_failure_duration&#x60; period (&#x60;min_location_failed&#x60; and &#x60;min_failure_duration&#x60; are part of the advanced alerting rules - integer, &gt;&#x3D; 1). | [optional] [default to 1]
**NewHostDelay** | Pointer to **NullableInt64** | Time (in seconds) to allow a host to boot and applications to fully start before starting the evaluation of monitor results. Should be a non negative integer. | [optional] [default to 300]
**NoDataTimeframe** | Pointer to **NullableInt64** | The number of minutes before a monitor notifies after data stops reporting. Datadog recommends at least 2x the monitor timeframe for metric alerts or 2 minutes for service checks. If omitted, 2x the evaluation timeframe is used for metric alerts, and 24 hours is used for service checks. | [optional] 
**NotifyAudit** | Pointer to **bool** | A Boolean indicating whether tagged users is notified on changes to this monitor. | [optional] [default to false]
**NotifyNoData** | Pointer to **bool** | A Boolean indicating whether this monitor notifies when data stops reporting. | [optional] [default to false]
**RenotifyInterval** | Pointer to **NullableInt64** | The number of minutes after the last notification before a monitor re-notifies on the current status. It only re-notifies if it’s not resolved. | [optional] 
**RequireFullWindow** | Pointer to **bool** | A Boolean indicating whether this monitor needs a full window of data before it’s evaluated. We highly recommend you set this to &#x60;false&#x60; for sparse metrics, otherwise some evaluations are skipped. For “on average” “at all times” and “in total” aggregation, default is true. &#x60;False&#x60; otherwise. | [optional] [default to true]
**Silenced** | Pointer to **map[string]int64** | Information about the downtime applied to the monitor. | [optional] 
**SyntheticsCheckId** | Pointer to **NullableInt64** | ID of the corresponding Synthetic check. | [optional] 
**ThresholdWindows** | Pointer to [**MonitorThresholdWindowOptions**](MonitorThresholdWindowOptions.md) |  | [optional] 
**Thresholds** | Pointer to [**MonitorThresholds**](MonitorThresholds.md) |  | [optional] 
**TimeoutH** | Pointer to **NullableInt64** | The number of hours of the monitor not reporting data before it automatically resolves from a triggered state. | [optional] 

## Methods

### NewMonitorOptions

`func NewMonitorOptions() *MonitorOptions`

NewMonitorOptions instantiates a new MonitorOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorOptionsWithDefaults

`func NewMonitorOptionsWithDefaults() *MonitorOptions`

NewMonitorOptionsWithDefaults instantiates a new MonitorOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregation

`func (o *MonitorOptions) GetAggregation() MonitorOptionsAggregation`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *MonitorOptions) GetAggregationOk() (*MonitorOptionsAggregation, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *MonitorOptions) SetAggregation(v MonitorOptionsAggregation)`

SetAggregation sets Aggregation field to given value.

### HasAggregation

`func (o *MonitorOptions) HasAggregation() bool`

HasAggregation returns a boolean if a field has been set.

### GetDeviceIds

`func (o *MonitorOptions) GetDeviceIds() []MonitorDeviceID`

GetDeviceIds returns the DeviceIds field if non-nil, zero value otherwise.

### GetDeviceIdsOk

`func (o *MonitorOptions) GetDeviceIdsOk() (*[]MonitorDeviceID, bool)`

GetDeviceIdsOk returns a tuple with the DeviceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceIds

`func (o *MonitorOptions) SetDeviceIds(v []MonitorDeviceID)`

SetDeviceIds sets DeviceIds field to given value.

### HasDeviceIds

`func (o *MonitorOptions) HasDeviceIds() bool`

HasDeviceIds returns a boolean if a field has been set.

### GetEnableLogsSample

`func (o *MonitorOptions) GetEnableLogsSample() bool`

GetEnableLogsSample returns the EnableLogsSample field if non-nil, zero value otherwise.

### GetEnableLogsSampleOk

`func (o *MonitorOptions) GetEnableLogsSampleOk() (*bool, bool)`

GetEnableLogsSampleOk returns a tuple with the EnableLogsSample field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableLogsSample

`func (o *MonitorOptions) SetEnableLogsSample(v bool)`

SetEnableLogsSample sets EnableLogsSample field to given value.

### HasEnableLogsSample

`func (o *MonitorOptions) HasEnableLogsSample() bool`

HasEnableLogsSample returns a boolean if a field has been set.

### GetEscalationMessage

`func (o *MonitorOptions) GetEscalationMessage() string`

GetEscalationMessage returns the EscalationMessage field if non-nil, zero value otherwise.

### GetEscalationMessageOk

`func (o *MonitorOptions) GetEscalationMessageOk() (*string, bool)`

GetEscalationMessageOk returns a tuple with the EscalationMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEscalationMessage

`func (o *MonitorOptions) SetEscalationMessage(v string)`

SetEscalationMessage sets EscalationMessage field to given value.

### HasEscalationMessage

`func (o *MonitorOptions) HasEscalationMessage() bool`

HasEscalationMessage returns a boolean if a field has been set.

### GetEvaluationDelay

`func (o *MonitorOptions) GetEvaluationDelay() int64`

GetEvaluationDelay returns the EvaluationDelay field if non-nil, zero value otherwise.

### GetEvaluationDelayOk

`func (o *MonitorOptions) GetEvaluationDelayOk() (*int64, bool)`

GetEvaluationDelayOk returns a tuple with the EvaluationDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationDelay

`func (o *MonitorOptions) SetEvaluationDelay(v int64)`

SetEvaluationDelay sets EvaluationDelay field to given value.

### HasEvaluationDelay

`func (o *MonitorOptions) HasEvaluationDelay() bool`

HasEvaluationDelay returns a boolean if a field has been set.

### SetEvaluationDelayNil

`func (o *MonitorOptions) SetEvaluationDelayNil(b bool)`

 SetEvaluationDelayNil sets the value for EvaluationDelay to be an explicit nil

### UnsetEvaluationDelay
`func (o *MonitorOptions) UnsetEvaluationDelay()`

UnsetEvaluationDelay ensures that no value is present for EvaluationDelay, not even an explicit nil
### GetIncludeTags

`func (o *MonitorOptions) GetIncludeTags() bool`

GetIncludeTags returns the IncludeTags field if non-nil, zero value otherwise.

### GetIncludeTagsOk

`func (o *MonitorOptions) GetIncludeTagsOk() (*bool, bool)`

GetIncludeTagsOk returns a tuple with the IncludeTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeTags

`func (o *MonitorOptions) SetIncludeTags(v bool)`

SetIncludeTags sets IncludeTags field to given value.

### HasIncludeTags

`func (o *MonitorOptions) HasIncludeTags() bool`

HasIncludeTags returns a boolean if a field has been set.

### GetLocked

`func (o *MonitorOptions) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *MonitorOptions) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *MonitorOptions) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *MonitorOptions) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetMinFailureDuration

`func (o *MonitorOptions) GetMinFailureDuration() int64`

GetMinFailureDuration returns the MinFailureDuration field if non-nil, zero value otherwise.

### GetMinFailureDurationOk

`func (o *MonitorOptions) GetMinFailureDurationOk() (*int64, bool)`

GetMinFailureDurationOk returns a tuple with the MinFailureDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinFailureDuration

`func (o *MonitorOptions) SetMinFailureDuration(v int64)`

SetMinFailureDuration sets MinFailureDuration field to given value.

### HasMinFailureDuration

`func (o *MonitorOptions) HasMinFailureDuration() bool`

HasMinFailureDuration returns a boolean if a field has been set.

### SetMinFailureDurationNil

`func (o *MonitorOptions) SetMinFailureDurationNil(b bool)`

 SetMinFailureDurationNil sets the value for MinFailureDuration to be an explicit nil

### UnsetMinFailureDuration
`func (o *MonitorOptions) UnsetMinFailureDuration()`

UnsetMinFailureDuration ensures that no value is present for MinFailureDuration, not even an explicit nil
### GetMinLocationFailed

`func (o *MonitorOptions) GetMinLocationFailed() int64`

GetMinLocationFailed returns the MinLocationFailed field if non-nil, zero value otherwise.

### GetMinLocationFailedOk

`func (o *MonitorOptions) GetMinLocationFailedOk() (*int64, bool)`

GetMinLocationFailedOk returns a tuple with the MinLocationFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLocationFailed

`func (o *MonitorOptions) SetMinLocationFailed(v int64)`

SetMinLocationFailed sets MinLocationFailed field to given value.

### HasMinLocationFailed

`func (o *MonitorOptions) HasMinLocationFailed() bool`

HasMinLocationFailed returns a boolean if a field has been set.

### SetMinLocationFailedNil

`func (o *MonitorOptions) SetMinLocationFailedNil(b bool)`

 SetMinLocationFailedNil sets the value for MinLocationFailed to be an explicit nil

### UnsetMinLocationFailed
`func (o *MonitorOptions) UnsetMinLocationFailed()`

UnsetMinLocationFailed ensures that no value is present for MinLocationFailed, not even an explicit nil
### GetNewHostDelay

`func (o *MonitorOptions) GetNewHostDelay() int64`

GetNewHostDelay returns the NewHostDelay field if non-nil, zero value otherwise.

### GetNewHostDelayOk

`func (o *MonitorOptions) GetNewHostDelayOk() (*int64, bool)`

GetNewHostDelayOk returns a tuple with the NewHostDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewHostDelay

`func (o *MonitorOptions) SetNewHostDelay(v int64)`

SetNewHostDelay sets NewHostDelay field to given value.

### HasNewHostDelay

`func (o *MonitorOptions) HasNewHostDelay() bool`

HasNewHostDelay returns a boolean if a field has been set.

### SetNewHostDelayNil

`func (o *MonitorOptions) SetNewHostDelayNil(b bool)`

 SetNewHostDelayNil sets the value for NewHostDelay to be an explicit nil

### UnsetNewHostDelay
`func (o *MonitorOptions) UnsetNewHostDelay()`

UnsetNewHostDelay ensures that no value is present for NewHostDelay, not even an explicit nil
### GetNoDataTimeframe

`func (o *MonitorOptions) GetNoDataTimeframe() int64`

GetNoDataTimeframe returns the NoDataTimeframe field if non-nil, zero value otherwise.

### GetNoDataTimeframeOk

`func (o *MonitorOptions) GetNoDataTimeframeOk() (*int64, bool)`

GetNoDataTimeframeOk returns a tuple with the NoDataTimeframe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoDataTimeframe

`func (o *MonitorOptions) SetNoDataTimeframe(v int64)`

SetNoDataTimeframe sets NoDataTimeframe field to given value.

### HasNoDataTimeframe

`func (o *MonitorOptions) HasNoDataTimeframe() bool`

HasNoDataTimeframe returns a boolean if a field has been set.

### SetNoDataTimeframeNil

`func (o *MonitorOptions) SetNoDataTimeframeNil(b bool)`

 SetNoDataTimeframeNil sets the value for NoDataTimeframe to be an explicit nil

### UnsetNoDataTimeframe
`func (o *MonitorOptions) UnsetNoDataTimeframe()`

UnsetNoDataTimeframe ensures that no value is present for NoDataTimeframe, not even an explicit nil
### GetNotifyAudit

`func (o *MonitorOptions) GetNotifyAudit() bool`

GetNotifyAudit returns the NotifyAudit field if non-nil, zero value otherwise.

### GetNotifyAuditOk

`func (o *MonitorOptions) GetNotifyAuditOk() (*bool, bool)`

GetNotifyAuditOk returns a tuple with the NotifyAudit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyAudit

`func (o *MonitorOptions) SetNotifyAudit(v bool)`

SetNotifyAudit sets NotifyAudit field to given value.

### HasNotifyAudit

`func (o *MonitorOptions) HasNotifyAudit() bool`

HasNotifyAudit returns a boolean if a field has been set.

### GetNotifyNoData

`func (o *MonitorOptions) GetNotifyNoData() bool`

GetNotifyNoData returns the NotifyNoData field if non-nil, zero value otherwise.

### GetNotifyNoDataOk

`func (o *MonitorOptions) GetNotifyNoDataOk() (*bool, bool)`

GetNotifyNoDataOk returns a tuple with the NotifyNoData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyNoData

`func (o *MonitorOptions) SetNotifyNoData(v bool)`

SetNotifyNoData sets NotifyNoData field to given value.

### HasNotifyNoData

`func (o *MonitorOptions) HasNotifyNoData() bool`

HasNotifyNoData returns a boolean if a field has been set.

### GetRenotifyInterval

`func (o *MonitorOptions) GetRenotifyInterval() int64`

GetRenotifyInterval returns the RenotifyInterval field if non-nil, zero value otherwise.

### GetRenotifyIntervalOk

`func (o *MonitorOptions) GetRenotifyIntervalOk() (*int64, bool)`

GetRenotifyIntervalOk returns a tuple with the RenotifyInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenotifyInterval

`func (o *MonitorOptions) SetRenotifyInterval(v int64)`

SetRenotifyInterval sets RenotifyInterval field to given value.

### HasRenotifyInterval

`func (o *MonitorOptions) HasRenotifyInterval() bool`

HasRenotifyInterval returns a boolean if a field has been set.

### SetRenotifyIntervalNil

`func (o *MonitorOptions) SetRenotifyIntervalNil(b bool)`

 SetRenotifyIntervalNil sets the value for RenotifyInterval to be an explicit nil

### UnsetRenotifyInterval
`func (o *MonitorOptions) UnsetRenotifyInterval()`

UnsetRenotifyInterval ensures that no value is present for RenotifyInterval, not even an explicit nil
### GetRequireFullWindow

`func (o *MonitorOptions) GetRequireFullWindow() bool`

GetRequireFullWindow returns the RequireFullWindow field if non-nil, zero value otherwise.

### GetRequireFullWindowOk

`func (o *MonitorOptions) GetRequireFullWindowOk() (*bool, bool)`

GetRequireFullWindowOk returns a tuple with the RequireFullWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireFullWindow

`func (o *MonitorOptions) SetRequireFullWindow(v bool)`

SetRequireFullWindow sets RequireFullWindow field to given value.

### HasRequireFullWindow

`func (o *MonitorOptions) HasRequireFullWindow() bool`

HasRequireFullWindow returns a boolean if a field has been set.

### GetSilenced

`func (o *MonitorOptions) GetSilenced() map[string]int64`

GetSilenced returns the Silenced field if non-nil, zero value otherwise.

### GetSilencedOk

`func (o *MonitorOptions) GetSilencedOk() (*map[string]int64, bool)`

GetSilencedOk returns a tuple with the Silenced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSilenced

`func (o *MonitorOptions) SetSilenced(v map[string]int64)`

SetSilenced sets Silenced field to given value.

### HasSilenced

`func (o *MonitorOptions) HasSilenced() bool`

HasSilenced returns a boolean if a field has been set.

### GetSyntheticsCheckId

`func (o *MonitorOptions) GetSyntheticsCheckId() int64`

GetSyntheticsCheckId returns the SyntheticsCheckId field if non-nil, zero value otherwise.

### GetSyntheticsCheckIdOk

`func (o *MonitorOptions) GetSyntheticsCheckIdOk() (*int64, bool)`

GetSyntheticsCheckIdOk returns a tuple with the SyntheticsCheckId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyntheticsCheckId

`func (o *MonitorOptions) SetSyntheticsCheckId(v int64)`

SetSyntheticsCheckId sets SyntheticsCheckId field to given value.

### HasSyntheticsCheckId

`func (o *MonitorOptions) HasSyntheticsCheckId() bool`

HasSyntheticsCheckId returns a boolean if a field has been set.

### SetSyntheticsCheckIdNil

`func (o *MonitorOptions) SetSyntheticsCheckIdNil(b bool)`

 SetSyntheticsCheckIdNil sets the value for SyntheticsCheckId to be an explicit nil

### UnsetSyntheticsCheckId
`func (o *MonitorOptions) UnsetSyntheticsCheckId()`

UnsetSyntheticsCheckId ensures that no value is present for SyntheticsCheckId, not even an explicit nil
### GetThresholdWindows

`func (o *MonitorOptions) GetThresholdWindows() MonitorThresholdWindowOptions`

GetThresholdWindows returns the ThresholdWindows field if non-nil, zero value otherwise.

### GetThresholdWindowsOk

`func (o *MonitorOptions) GetThresholdWindowsOk() (*MonitorThresholdWindowOptions, bool)`

GetThresholdWindowsOk returns a tuple with the ThresholdWindows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdWindows

`func (o *MonitorOptions) SetThresholdWindows(v MonitorThresholdWindowOptions)`

SetThresholdWindows sets ThresholdWindows field to given value.

### HasThresholdWindows

`func (o *MonitorOptions) HasThresholdWindows() bool`

HasThresholdWindows returns a boolean if a field has been set.

### GetThresholds

`func (o *MonitorOptions) GetThresholds() MonitorThresholds`

GetThresholds returns the Thresholds field if non-nil, zero value otherwise.

### GetThresholdsOk

`func (o *MonitorOptions) GetThresholdsOk() (*MonitorThresholds, bool)`

GetThresholdsOk returns a tuple with the Thresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholds

`func (o *MonitorOptions) SetThresholds(v MonitorThresholds)`

SetThresholds sets Thresholds field to given value.

### HasThresholds

`func (o *MonitorOptions) HasThresholds() bool`

HasThresholds returns a boolean if a field has been set.

### GetTimeoutH

`func (o *MonitorOptions) GetTimeoutH() int64`

GetTimeoutH returns the TimeoutH field if non-nil, zero value otherwise.

### GetTimeoutHOk

`func (o *MonitorOptions) GetTimeoutHOk() (*int64, bool)`

GetTimeoutHOk returns a tuple with the TimeoutH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutH

`func (o *MonitorOptions) SetTimeoutH(v int64)`

SetTimeoutH sets TimeoutH field to given value.

### HasTimeoutH

`func (o *MonitorOptions) HasTimeoutH() bool`

HasTimeoutH returns a boolean if a field has been set.

### SetTimeoutHNil

`func (o *MonitorOptions) SetTimeoutHNil(b bool)`

 SetTimeoutHNil sets the value for TimeoutH to be an explicit nil

### UnsetTimeoutH
`func (o *MonitorOptions) UnsetTimeoutH()`

UnsetTimeoutH ensures that no value is present for TimeoutH, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


