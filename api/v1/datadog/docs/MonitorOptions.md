# MonitorOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregation** | Pointer to [**MonitorOptionsAggregation**](MonitorOptions_aggregation.md) |  | [optional] 
**DeviceIds** | Pointer to [**[]MonitorDeviceID**](MonitorDeviceID.md) | TODO. | [optional] 
**EnableLogsSample** | Pointer to **bool** | TODO. | [optional] 
**EscalationMessage** | Pointer to **string** | A message to include with a re-notification. Supports the ‘@username’ notification we allow elsewhere. Not applicable if &#x60;renotify_interval&#x60; is None. | [optional] [default to "none"]
**EvaluationDelay** | Pointer to **NullableInt64** | Time (in seconds) to delay evaluation, as a non-negative integer. For example, if the value is set to 300 (5min), the timeframe is set to last_5m and the time is 7:00, the monitor evaluates data from 6:50 to 6:55. This is useful for AWS CloudWatch and other backfilled metrics to ensure the monitor always has data during evaluation. | [optional] 
**IncludeTags** | Pointer to **bool** | A Boolean indicating whether notifications from this monitor automatically inserts its triggering tags into the title.  **Examples** - If &#x60;True&#x60;, &#x60;[Triggered on {host:h1}] Monitor Title&#x60; - If &#x60;False&#x60;, &#x60;[Triggered] Monitor Title&#x60; | [optional] [default to true]
**Locked** | Pointer to **bool** |  | [optional] 
**MinFailureDuration** | Pointer to **NullableInt64** | How long the test should be in failure before alerting (integer, number of seconds, max 7200). | [optional] [default to 0]
**MinLocationFailed** | Pointer to **NullableInt64** | The minimum number of locations in failure at the same time during at least one moment in the &#x60;min_failure_duration&#x60; period (&#x60;min_location_failed&#x60; and &#x60;min_failure_duration&#x60; are part of the advanced alerting rules - integer, &gt;&#x3D; 1). | [optional] [default to 1]
**NewHostDelay** | Pointer to **NullableInt64** | Time (in seconds) to allow a host to boot and applications to fully start before starting the evaluation of monitor results. Should be a non negative integer. | [optional] [default to 300]
**NoDataTimeframe** | Pointer to **NullableInt64** | The number of minutes before a monitor notifies after data stops reporting. Datadog recommends at least 2x the monitor timeframe for metric alerts or 2 minutes for service checks. If omitted, 2x the evaluation timeframe is used for metric alerts, and 24 hours is used for service checks. | [optional] 
**NotifyAudit** | Pointer to **bool** | A Boolean indicating whether tagged users is notified on changes to this monitor. | [optional] [default to false]
**NotifyNoData** | Pointer to **bool** | A Boolean indicating whether this monitor notifies when data stops reporting. | [optional] [default to false]
**RenotifyInterval** | Pointer to **NullableInt64** | The number of minutes after the last notification before a monitor re-notifies on the current status. It only re-notifies if it’s not resolved. | [optional] 
**RequireFullWindow** | Pointer to **bool** | A Boolean indicating whether this monitor needs a full window of data before it’s evaluated. We highly recommend you set this to &#x60;false&#x60; for sparse metrics, otherwise some evaluations are skipped. For “on average” “at all times” and “in total” aggregation, default is true. &#x60;False&#x60; otherwise. | [optional] [default to true]
**Silenced** | Pointer to **map[string]int64** | TODO. | [optional] 
**SyntheticsCheckId** | Pointer to **NullableInt64** | TODO. | [optional] 
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

`func (o *MonitorOptions) GetAggregationOk() (MonitorOptionsAggregation, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAggregation

`func (o *MonitorOptions) HasAggregation() bool`

HasAggregation returns a boolean if a field has been set.

### SetAggregation

`func (o *MonitorOptions) SetAggregation(v MonitorOptionsAggregation)`

SetAggregation gets a reference to the given MonitorOptionsAggregation and assigns it to the Aggregation field.

### GetDeviceIds

`func (o *MonitorOptions) GetDeviceIds() []MonitorDeviceID`

GetDeviceIds returns the DeviceIds field if non-nil, zero value otherwise.

### GetDeviceIdsOk

`func (o *MonitorOptions) GetDeviceIdsOk() ([]MonitorDeviceID, bool)`

GetDeviceIdsOk returns a tuple with the DeviceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceIds

`func (o *MonitorOptions) HasDeviceIds() bool`

HasDeviceIds returns a boolean if a field has been set.

### SetDeviceIds

`func (o *MonitorOptions) SetDeviceIds(v []MonitorDeviceID)`

SetDeviceIds gets a reference to the given []MonitorDeviceID and assigns it to the DeviceIds field.

### GetEnableLogsSample

`func (o *MonitorOptions) GetEnableLogsSample() bool`

GetEnableLogsSample returns the EnableLogsSample field if non-nil, zero value otherwise.

### GetEnableLogsSampleOk

`func (o *MonitorOptions) GetEnableLogsSampleOk() (bool, bool)`

GetEnableLogsSampleOk returns a tuple with the EnableLogsSample field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnableLogsSample

`func (o *MonitorOptions) HasEnableLogsSample() bool`

HasEnableLogsSample returns a boolean if a field has been set.

### SetEnableLogsSample

`func (o *MonitorOptions) SetEnableLogsSample(v bool)`

SetEnableLogsSample gets a reference to the given bool and assigns it to the EnableLogsSample field.

### GetEscalationMessage

`func (o *MonitorOptions) GetEscalationMessage() string`

GetEscalationMessage returns the EscalationMessage field if non-nil, zero value otherwise.

### GetEscalationMessageOk

`func (o *MonitorOptions) GetEscalationMessageOk() (string, bool)`

GetEscalationMessageOk returns a tuple with the EscalationMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEscalationMessage

`func (o *MonitorOptions) HasEscalationMessage() bool`

HasEscalationMessage returns a boolean if a field has been set.

### SetEscalationMessage

`func (o *MonitorOptions) SetEscalationMessage(v string)`

SetEscalationMessage gets a reference to the given string and assigns it to the EscalationMessage field.

### GetEvaluationDelay

`func (o *MonitorOptions) GetEvaluationDelay() NullableInt64`

GetEvaluationDelay returns the EvaluationDelay field if non-nil, zero value otherwise.

### GetEvaluationDelayOk

`func (o *MonitorOptions) GetEvaluationDelayOk() (NullableInt64, bool)`

GetEvaluationDelayOk returns a tuple with the EvaluationDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEvaluationDelay

`func (o *MonitorOptions) HasEvaluationDelay() bool`

HasEvaluationDelay returns a boolean if a field has been set.

### SetEvaluationDelay

`func (o *MonitorOptions) SetEvaluationDelay(v NullableInt64)`

SetEvaluationDelay gets a reference to the given NullableInt64 and assigns it to the EvaluationDelay field.

### SetEvaluationDelayExplicitNull

`func (o *MonitorOptions) SetEvaluationDelayExplicitNull(b bool)`

SetEvaluationDelayExplicitNull (un)sets EvaluationDelay to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The EvaluationDelay value is set to nil even if false is passed
### GetIncludeTags

`func (o *MonitorOptions) GetIncludeTags() bool`

GetIncludeTags returns the IncludeTags field if non-nil, zero value otherwise.

### GetIncludeTagsOk

`func (o *MonitorOptions) GetIncludeTagsOk() (bool, bool)`

GetIncludeTagsOk returns a tuple with the IncludeTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIncludeTags

`func (o *MonitorOptions) HasIncludeTags() bool`

HasIncludeTags returns a boolean if a field has been set.

### SetIncludeTags

`func (o *MonitorOptions) SetIncludeTags(v bool)`

SetIncludeTags gets a reference to the given bool and assigns it to the IncludeTags field.

### GetLocked

`func (o *MonitorOptions) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *MonitorOptions) GetLockedOk() (bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLocked

`func (o *MonitorOptions) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### SetLocked

`func (o *MonitorOptions) SetLocked(v bool)`

SetLocked gets a reference to the given bool and assigns it to the Locked field.

### GetMinFailureDuration

`func (o *MonitorOptions) GetMinFailureDuration() NullableInt64`

GetMinFailureDuration returns the MinFailureDuration field if non-nil, zero value otherwise.

### GetMinFailureDurationOk

`func (o *MonitorOptions) GetMinFailureDurationOk() (NullableInt64, bool)`

GetMinFailureDurationOk returns a tuple with the MinFailureDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMinFailureDuration

`func (o *MonitorOptions) HasMinFailureDuration() bool`

HasMinFailureDuration returns a boolean if a field has been set.

### SetMinFailureDuration

`func (o *MonitorOptions) SetMinFailureDuration(v NullableInt64)`

SetMinFailureDuration gets a reference to the given NullableInt64 and assigns it to the MinFailureDuration field.

### SetMinFailureDurationExplicitNull

`func (o *MonitorOptions) SetMinFailureDurationExplicitNull(b bool)`

SetMinFailureDurationExplicitNull (un)sets MinFailureDuration to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MinFailureDuration value is set to nil even if false is passed
### GetMinLocationFailed

`func (o *MonitorOptions) GetMinLocationFailed() NullableInt64`

GetMinLocationFailed returns the MinLocationFailed field if non-nil, zero value otherwise.

### GetMinLocationFailedOk

`func (o *MonitorOptions) GetMinLocationFailedOk() (NullableInt64, bool)`

GetMinLocationFailedOk returns a tuple with the MinLocationFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMinLocationFailed

`func (o *MonitorOptions) HasMinLocationFailed() bool`

HasMinLocationFailed returns a boolean if a field has been set.

### SetMinLocationFailed

`func (o *MonitorOptions) SetMinLocationFailed(v NullableInt64)`

SetMinLocationFailed gets a reference to the given NullableInt64 and assigns it to the MinLocationFailed field.

### SetMinLocationFailedExplicitNull

`func (o *MonitorOptions) SetMinLocationFailedExplicitNull(b bool)`

SetMinLocationFailedExplicitNull (un)sets MinLocationFailed to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MinLocationFailed value is set to nil even if false is passed
### GetNewHostDelay

`func (o *MonitorOptions) GetNewHostDelay() NullableInt64`

GetNewHostDelay returns the NewHostDelay field if non-nil, zero value otherwise.

### GetNewHostDelayOk

`func (o *MonitorOptions) GetNewHostDelayOk() (NullableInt64, bool)`

GetNewHostDelayOk returns a tuple with the NewHostDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNewHostDelay

`func (o *MonitorOptions) HasNewHostDelay() bool`

HasNewHostDelay returns a boolean if a field has been set.

### SetNewHostDelay

`func (o *MonitorOptions) SetNewHostDelay(v NullableInt64)`

SetNewHostDelay gets a reference to the given NullableInt64 and assigns it to the NewHostDelay field.

### SetNewHostDelayExplicitNull

`func (o *MonitorOptions) SetNewHostDelayExplicitNull(b bool)`

SetNewHostDelayExplicitNull (un)sets NewHostDelay to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The NewHostDelay value is set to nil even if false is passed
### GetNoDataTimeframe

`func (o *MonitorOptions) GetNoDataTimeframe() NullableInt64`

GetNoDataTimeframe returns the NoDataTimeframe field if non-nil, zero value otherwise.

### GetNoDataTimeframeOk

`func (o *MonitorOptions) GetNoDataTimeframeOk() (NullableInt64, bool)`

GetNoDataTimeframeOk returns a tuple with the NoDataTimeframe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNoDataTimeframe

`func (o *MonitorOptions) HasNoDataTimeframe() bool`

HasNoDataTimeframe returns a boolean if a field has been set.

### SetNoDataTimeframe

`func (o *MonitorOptions) SetNoDataTimeframe(v NullableInt64)`

SetNoDataTimeframe gets a reference to the given NullableInt64 and assigns it to the NoDataTimeframe field.

### SetNoDataTimeframeExplicitNull

`func (o *MonitorOptions) SetNoDataTimeframeExplicitNull(b bool)`

SetNoDataTimeframeExplicitNull (un)sets NoDataTimeframe to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The NoDataTimeframe value is set to nil even if false is passed
### GetNotifyAudit

`func (o *MonitorOptions) GetNotifyAudit() bool`

GetNotifyAudit returns the NotifyAudit field if non-nil, zero value otherwise.

### GetNotifyAuditOk

`func (o *MonitorOptions) GetNotifyAuditOk() (bool, bool)`

GetNotifyAuditOk returns a tuple with the NotifyAudit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNotifyAudit

`func (o *MonitorOptions) HasNotifyAudit() bool`

HasNotifyAudit returns a boolean if a field has been set.

### SetNotifyAudit

`func (o *MonitorOptions) SetNotifyAudit(v bool)`

SetNotifyAudit gets a reference to the given bool and assigns it to the NotifyAudit field.

### GetNotifyNoData

`func (o *MonitorOptions) GetNotifyNoData() bool`

GetNotifyNoData returns the NotifyNoData field if non-nil, zero value otherwise.

### GetNotifyNoDataOk

`func (o *MonitorOptions) GetNotifyNoDataOk() (bool, bool)`

GetNotifyNoDataOk returns a tuple with the NotifyNoData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNotifyNoData

`func (o *MonitorOptions) HasNotifyNoData() bool`

HasNotifyNoData returns a boolean if a field has been set.

### SetNotifyNoData

`func (o *MonitorOptions) SetNotifyNoData(v bool)`

SetNotifyNoData gets a reference to the given bool and assigns it to the NotifyNoData field.

### GetRenotifyInterval

`func (o *MonitorOptions) GetRenotifyInterval() NullableInt64`

GetRenotifyInterval returns the RenotifyInterval field if non-nil, zero value otherwise.

### GetRenotifyIntervalOk

`func (o *MonitorOptions) GetRenotifyIntervalOk() (NullableInt64, bool)`

GetRenotifyIntervalOk returns a tuple with the RenotifyInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRenotifyInterval

`func (o *MonitorOptions) HasRenotifyInterval() bool`

HasRenotifyInterval returns a boolean if a field has been set.

### SetRenotifyInterval

`func (o *MonitorOptions) SetRenotifyInterval(v NullableInt64)`

SetRenotifyInterval gets a reference to the given NullableInt64 and assigns it to the RenotifyInterval field.

### SetRenotifyIntervalExplicitNull

`func (o *MonitorOptions) SetRenotifyIntervalExplicitNull(b bool)`

SetRenotifyIntervalExplicitNull (un)sets RenotifyInterval to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The RenotifyInterval value is set to nil even if false is passed
### GetRequireFullWindow

`func (o *MonitorOptions) GetRequireFullWindow() bool`

GetRequireFullWindow returns the RequireFullWindow field if non-nil, zero value otherwise.

### GetRequireFullWindowOk

`func (o *MonitorOptions) GetRequireFullWindowOk() (bool, bool)`

GetRequireFullWindowOk returns a tuple with the RequireFullWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRequireFullWindow

`func (o *MonitorOptions) HasRequireFullWindow() bool`

HasRequireFullWindow returns a boolean if a field has been set.

### SetRequireFullWindow

`func (o *MonitorOptions) SetRequireFullWindow(v bool)`

SetRequireFullWindow gets a reference to the given bool and assigns it to the RequireFullWindow field.

### GetSilenced

`func (o *MonitorOptions) GetSilenced() map[string]int64`

GetSilenced returns the Silenced field if non-nil, zero value otherwise.

### GetSilencedOk

`func (o *MonitorOptions) GetSilencedOk() (map[string]int64, bool)`

GetSilencedOk returns a tuple with the Silenced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSilenced

`func (o *MonitorOptions) HasSilenced() bool`

HasSilenced returns a boolean if a field has been set.

### SetSilenced

`func (o *MonitorOptions) SetSilenced(v map[string]int64)`

SetSilenced gets a reference to the given map[string]int64 and assigns it to the Silenced field.

### GetSyntheticsCheckId

`func (o *MonitorOptions) GetSyntheticsCheckId() NullableInt64`

GetSyntheticsCheckId returns the SyntheticsCheckId field if non-nil, zero value otherwise.

### GetSyntheticsCheckIdOk

`func (o *MonitorOptions) GetSyntheticsCheckIdOk() (NullableInt64, bool)`

GetSyntheticsCheckIdOk returns a tuple with the SyntheticsCheckId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSyntheticsCheckId

`func (o *MonitorOptions) HasSyntheticsCheckId() bool`

HasSyntheticsCheckId returns a boolean if a field has been set.

### SetSyntheticsCheckId

`func (o *MonitorOptions) SetSyntheticsCheckId(v NullableInt64)`

SetSyntheticsCheckId gets a reference to the given NullableInt64 and assigns it to the SyntheticsCheckId field.

### SetSyntheticsCheckIdExplicitNull

`func (o *MonitorOptions) SetSyntheticsCheckIdExplicitNull(b bool)`

SetSyntheticsCheckIdExplicitNull (un)sets SyntheticsCheckId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SyntheticsCheckId value is set to nil even if false is passed
### GetThresholdWindows

`func (o *MonitorOptions) GetThresholdWindows() MonitorThresholdWindowOptions`

GetThresholdWindows returns the ThresholdWindows field if non-nil, zero value otherwise.

### GetThresholdWindowsOk

`func (o *MonitorOptions) GetThresholdWindowsOk() (MonitorThresholdWindowOptions, bool)`

GetThresholdWindowsOk returns a tuple with the ThresholdWindows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasThresholdWindows

`func (o *MonitorOptions) HasThresholdWindows() bool`

HasThresholdWindows returns a boolean if a field has been set.

### SetThresholdWindows

`func (o *MonitorOptions) SetThresholdWindows(v MonitorThresholdWindowOptions)`

SetThresholdWindows gets a reference to the given MonitorThresholdWindowOptions and assigns it to the ThresholdWindows field.

### GetThresholds

`func (o *MonitorOptions) GetThresholds() MonitorThresholds`

GetThresholds returns the Thresholds field if non-nil, zero value otherwise.

### GetThresholdsOk

`func (o *MonitorOptions) GetThresholdsOk() (MonitorThresholds, bool)`

GetThresholdsOk returns a tuple with the Thresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasThresholds

`func (o *MonitorOptions) HasThresholds() bool`

HasThresholds returns a boolean if a field has been set.

### SetThresholds

`func (o *MonitorOptions) SetThresholds(v MonitorThresholds)`

SetThresholds gets a reference to the given MonitorThresholds and assigns it to the Thresholds field.

### GetTimeoutH

`func (o *MonitorOptions) GetTimeoutH() NullableInt64`

GetTimeoutH returns the TimeoutH field if non-nil, zero value otherwise.

### GetTimeoutHOk

`func (o *MonitorOptions) GetTimeoutHOk() (NullableInt64, bool)`

GetTimeoutHOk returns a tuple with the TimeoutH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTimeoutH

`func (o *MonitorOptions) HasTimeoutH() bool`

HasTimeoutH returns a boolean if a field has been set.

### SetTimeoutH

`func (o *MonitorOptions) SetTimeoutH(v NullableInt64)`

SetTimeoutH gets a reference to the given NullableInt64 and assigns it to the TimeoutH field.

### SetTimeoutHExplicitNull

`func (o *MonitorOptions) SetTimeoutHExplicitNull(b bool)`

SetTimeoutHExplicitNull (un)sets TimeoutH to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The TimeoutH value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


