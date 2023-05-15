// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"encoding/json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsRetentionSumUsage Object containing indexed logs usage grouped by retention period and summed.
type LogsRetentionSumUsage struct {
	// Total indexed logs for this retention period.
	LogsIndexedLogsUsageSum datadog.NullableInt64 `json:"logs_indexed_logs_usage_sum,omitempty"`
	// Live indexed logs for this retention period.
	LogsLiveIndexedLogsUsageSum datadog.NullableInt64 `json:"logs_live_indexed_logs_usage_sum,omitempty"`
	// Rehydrated indexed logs for this retention period.
	LogsRehydratedIndexedLogsUsageSum datadog.NullableInt64 `json:"logs_rehydrated_indexed_logs_usage_sum,omitempty"`
	// The retention period in days or "custom" for all custom retention periods.
	Retention datadog.NullableString `json:"retention,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewLogsRetentionSumUsage instantiates a new LogsRetentionSumUsage object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLogsRetentionSumUsage() *LogsRetentionSumUsage {
	this := LogsRetentionSumUsage{}
	return &this
}

// NewLogsRetentionSumUsageWithDefaults instantiates a new LogsRetentionSumUsage object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLogsRetentionSumUsageWithDefaults() *LogsRetentionSumUsage {
	this := LogsRetentionSumUsage{}
	return &this
}

// GetLogsIndexedLogsUsageSum returns the LogsIndexedLogsUsageSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LogsRetentionSumUsage) GetLogsIndexedLogsUsageSum() int64 {
	if o == nil || o.LogsIndexedLogsUsageSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.LogsIndexedLogsUsageSum.Get()
}

// GetLogsIndexedLogsUsageSumOk returns a tuple with the LogsIndexedLogsUsageSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LogsRetentionSumUsage) GetLogsIndexedLogsUsageSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LogsIndexedLogsUsageSum.Get(), o.LogsIndexedLogsUsageSum.IsSet()
}

// HasLogsIndexedLogsUsageSum returns a boolean if a field has been set.
func (o *LogsRetentionSumUsage) HasLogsIndexedLogsUsageSum() bool {
	return o != nil && o.LogsIndexedLogsUsageSum.IsSet()
}

// SetLogsIndexedLogsUsageSum gets a reference to the given datadog.NullableInt64 and assigns it to the LogsIndexedLogsUsageSum field.
func (o *LogsRetentionSumUsage) SetLogsIndexedLogsUsageSum(v int64) {
	o.LogsIndexedLogsUsageSum.Set(&v)
}

// SetLogsIndexedLogsUsageSumNil sets the value for LogsIndexedLogsUsageSum to be an explicit nil.
func (o *LogsRetentionSumUsage) SetLogsIndexedLogsUsageSumNil() {
	o.LogsIndexedLogsUsageSum.Set(nil)
}

// UnsetLogsIndexedLogsUsageSum ensures that no value is present for LogsIndexedLogsUsageSum, not even an explicit nil.
func (o *LogsRetentionSumUsage) UnsetLogsIndexedLogsUsageSum() {
	o.LogsIndexedLogsUsageSum.Unset()
}

// GetLogsLiveIndexedLogsUsageSum returns the LogsLiveIndexedLogsUsageSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LogsRetentionSumUsage) GetLogsLiveIndexedLogsUsageSum() int64 {
	if o == nil || o.LogsLiveIndexedLogsUsageSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.LogsLiveIndexedLogsUsageSum.Get()
}

// GetLogsLiveIndexedLogsUsageSumOk returns a tuple with the LogsLiveIndexedLogsUsageSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LogsRetentionSumUsage) GetLogsLiveIndexedLogsUsageSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LogsLiveIndexedLogsUsageSum.Get(), o.LogsLiveIndexedLogsUsageSum.IsSet()
}

// HasLogsLiveIndexedLogsUsageSum returns a boolean if a field has been set.
func (o *LogsRetentionSumUsage) HasLogsLiveIndexedLogsUsageSum() bool {
	return o != nil && o.LogsLiveIndexedLogsUsageSum.IsSet()
}

// SetLogsLiveIndexedLogsUsageSum gets a reference to the given datadog.NullableInt64 and assigns it to the LogsLiveIndexedLogsUsageSum field.
func (o *LogsRetentionSumUsage) SetLogsLiveIndexedLogsUsageSum(v int64) {
	o.LogsLiveIndexedLogsUsageSum.Set(&v)
}

// SetLogsLiveIndexedLogsUsageSumNil sets the value for LogsLiveIndexedLogsUsageSum to be an explicit nil.
func (o *LogsRetentionSumUsage) SetLogsLiveIndexedLogsUsageSumNil() {
	o.LogsLiveIndexedLogsUsageSum.Set(nil)
}

// UnsetLogsLiveIndexedLogsUsageSum ensures that no value is present for LogsLiveIndexedLogsUsageSum, not even an explicit nil.
func (o *LogsRetentionSumUsage) UnsetLogsLiveIndexedLogsUsageSum() {
	o.LogsLiveIndexedLogsUsageSum.Unset()
}

// GetLogsRehydratedIndexedLogsUsageSum returns the LogsRehydratedIndexedLogsUsageSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LogsRetentionSumUsage) GetLogsRehydratedIndexedLogsUsageSum() int64 {
	if o == nil || o.LogsRehydratedIndexedLogsUsageSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.LogsRehydratedIndexedLogsUsageSum.Get()
}

// GetLogsRehydratedIndexedLogsUsageSumOk returns a tuple with the LogsRehydratedIndexedLogsUsageSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LogsRetentionSumUsage) GetLogsRehydratedIndexedLogsUsageSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LogsRehydratedIndexedLogsUsageSum.Get(), o.LogsRehydratedIndexedLogsUsageSum.IsSet()
}

// HasLogsRehydratedIndexedLogsUsageSum returns a boolean if a field has been set.
func (o *LogsRetentionSumUsage) HasLogsRehydratedIndexedLogsUsageSum() bool {
	return o != nil && o.LogsRehydratedIndexedLogsUsageSum.IsSet()
}

// SetLogsRehydratedIndexedLogsUsageSum gets a reference to the given datadog.NullableInt64 and assigns it to the LogsRehydratedIndexedLogsUsageSum field.
func (o *LogsRetentionSumUsage) SetLogsRehydratedIndexedLogsUsageSum(v int64) {
	o.LogsRehydratedIndexedLogsUsageSum.Set(&v)
}

// SetLogsRehydratedIndexedLogsUsageSumNil sets the value for LogsRehydratedIndexedLogsUsageSum to be an explicit nil.
func (o *LogsRetentionSumUsage) SetLogsRehydratedIndexedLogsUsageSumNil() {
	o.LogsRehydratedIndexedLogsUsageSum.Set(nil)
}

// UnsetLogsRehydratedIndexedLogsUsageSum ensures that no value is present for LogsRehydratedIndexedLogsUsageSum, not even an explicit nil.
func (o *LogsRetentionSumUsage) UnsetLogsRehydratedIndexedLogsUsageSum() {
	o.LogsRehydratedIndexedLogsUsageSum.Unset()
}

// GetRetention returns the Retention field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LogsRetentionSumUsage) GetRetention() string {
	if o == nil || o.Retention.Get() == nil {
		var ret string
		return ret
	}
	return *o.Retention.Get()
}

// GetRetentionOk returns a tuple with the Retention field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LogsRetentionSumUsage) GetRetentionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Retention.Get(), o.Retention.IsSet()
}

// HasRetention returns a boolean if a field has been set.
func (o *LogsRetentionSumUsage) HasRetention() bool {
	return o != nil && o.Retention.IsSet()
}

// SetRetention gets a reference to the given datadog.NullableString and assigns it to the Retention field.
func (o *LogsRetentionSumUsage) SetRetention(v string) {
	o.Retention.Set(&v)
}

// SetRetentionNil sets the value for Retention to be an explicit nil.
func (o *LogsRetentionSumUsage) SetRetentionNil() {
	o.Retention.Set(nil)
}

// UnsetRetention ensures that no value is present for Retention, not even an explicit nil.
func (o *LogsRetentionSumUsage) UnsetRetention() {
	o.Retention.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o LogsRetentionSumUsage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.LogsIndexedLogsUsageSum.IsSet() {
		toSerialize["logs_indexed_logs_usage_sum"] = o.LogsIndexedLogsUsageSum.Get()
	}
	if o.LogsLiveIndexedLogsUsageSum.IsSet() {
		toSerialize["logs_live_indexed_logs_usage_sum"] = o.LogsLiveIndexedLogsUsageSum.Get()
	}
	if o.LogsRehydratedIndexedLogsUsageSum.IsSet() {
		toSerialize["logs_rehydrated_indexed_logs_usage_sum"] = o.LogsRehydratedIndexedLogsUsageSum.Get()
	}
	if o.Retention.IsSet() {
		toSerialize["retention"] = o.Retention.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LogsRetentionSumUsage) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		LogsIndexedLogsUsageSum           datadog.NullableInt64  `json:"logs_indexed_logs_usage_sum,omitempty"`
		LogsLiveIndexedLogsUsageSum       datadog.NullableInt64  `json:"logs_live_indexed_logs_usage_sum,omitempty"`
		LogsRehydratedIndexedLogsUsageSum datadog.NullableInt64  `json:"logs_rehydrated_indexed_logs_usage_sum,omitempty"`
		Retention                         datadog.NullableString `json:"retention,omitempty"`
	}{}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"logs_indexed_logs_usage_sum", "logs_live_indexed_logs_usage_sum", "logs_rehydrated_indexed_logs_usage_sum", "retention"})
	} else {
		return err
	}
	o.LogsIndexedLogsUsageSum = all.LogsIndexedLogsUsageSum
	o.LogsLiveIndexedLogsUsageSum = all.LogsLiveIndexedLogsUsageSum
	o.LogsRehydratedIndexedLogsUsageSum = all.LogsRehydratedIndexedLogsUsageSum
	o.Retention = all.Retention
	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
