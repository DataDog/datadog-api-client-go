// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"encoding/json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsRetentionAggSumUsage Object containing indexed logs usage aggregated across organizations and months for a retention period.
type LogsRetentionAggSumUsage struct {
	// Total indexed logs for this retention period.
	LogsIndexedLogsUsageAggSum datadog.NullableInt64 `json:"logs_indexed_logs_usage_agg_sum,omitempty"`
	// Live indexed logs for this retention period.
	LogsLiveIndexedLogsUsageAggSum datadog.NullableInt64 `json:"logs_live_indexed_logs_usage_agg_sum,omitempty"`
	// Rehydrated indexed logs for this retention period.
	LogsRehydratedIndexedLogsUsageAggSum datadog.NullableInt64 `json:"logs_rehydrated_indexed_logs_usage_agg_sum,omitempty"`
	// The retention period in days or "custom" for all custom retention periods.
	Retention datadog.NullableString `json:"retention,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewLogsRetentionAggSumUsage instantiates a new LogsRetentionAggSumUsage object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLogsRetentionAggSumUsage() *LogsRetentionAggSumUsage {
	this := LogsRetentionAggSumUsage{}
	return &this
}

// NewLogsRetentionAggSumUsageWithDefaults instantiates a new LogsRetentionAggSumUsage object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLogsRetentionAggSumUsageWithDefaults() *LogsRetentionAggSumUsage {
	this := LogsRetentionAggSumUsage{}
	return &this
}

// GetLogsIndexedLogsUsageAggSum returns the LogsIndexedLogsUsageAggSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LogsRetentionAggSumUsage) GetLogsIndexedLogsUsageAggSum() int64 {
	if o == nil || o.LogsIndexedLogsUsageAggSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.LogsIndexedLogsUsageAggSum.Get()
}

// GetLogsIndexedLogsUsageAggSumOk returns a tuple with the LogsIndexedLogsUsageAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LogsRetentionAggSumUsage) GetLogsIndexedLogsUsageAggSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LogsIndexedLogsUsageAggSum.Get(), o.LogsIndexedLogsUsageAggSum.IsSet()
}

// HasLogsIndexedLogsUsageAggSum returns a boolean if a field has been set.
func (o *LogsRetentionAggSumUsage) HasLogsIndexedLogsUsageAggSum() bool {
	return o != nil && o.LogsIndexedLogsUsageAggSum.IsSet()
}

// SetLogsIndexedLogsUsageAggSum gets a reference to the given datadog.NullableInt64 and assigns it to the LogsIndexedLogsUsageAggSum field.
func (o *LogsRetentionAggSumUsage) SetLogsIndexedLogsUsageAggSum(v int64) {
	o.LogsIndexedLogsUsageAggSum.Set(&v)
}

// SetLogsIndexedLogsUsageAggSumNil sets the value for LogsIndexedLogsUsageAggSum to be an explicit nil.
func (o *LogsRetentionAggSumUsage) SetLogsIndexedLogsUsageAggSumNil() {
	o.LogsIndexedLogsUsageAggSum.Set(nil)
}

// UnsetLogsIndexedLogsUsageAggSum ensures that no value is present for LogsIndexedLogsUsageAggSum, not even an explicit nil.
func (o *LogsRetentionAggSumUsage) UnsetLogsIndexedLogsUsageAggSum() {
	o.LogsIndexedLogsUsageAggSum.Unset()
}

// GetLogsLiveIndexedLogsUsageAggSum returns the LogsLiveIndexedLogsUsageAggSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LogsRetentionAggSumUsage) GetLogsLiveIndexedLogsUsageAggSum() int64 {
	if o == nil || o.LogsLiveIndexedLogsUsageAggSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.LogsLiveIndexedLogsUsageAggSum.Get()
}

// GetLogsLiveIndexedLogsUsageAggSumOk returns a tuple with the LogsLiveIndexedLogsUsageAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LogsRetentionAggSumUsage) GetLogsLiveIndexedLogsUsageAggSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LogsLiveIndexedLogsUsageAggSum.Get(), o.LogsLiveIndexedLogsUsageAggSum.IsSet()
}

// HasLogsLiveIndexedLogsUsageAggSum returns a boolean if a field has been set.
func (o *LogsRetentionAggSumUsage) HasLogsLiveIndexedLogsUsageAggSum() bool {
	return o != nil && o.LogsLiveIndexedLogsUsageAggSum.IsSet()
}

// SetLogsLiveIndexedLogsUsageAggSum gets a reference to the given datadog.NullableInt64 and assigns it to the LogsLiveIndexedLogsUsageAggSum field.
func (o *LogsRetentionAggSumUsage) SetLogsLiveIndexedLogsUsageAggSum(v int64) {
	o.LogsLiveIndexedLogsUsageAggSum.Set(&v)
}

// SetLogsLiveIndexedLogsUsageAggSumNil sets the value for LogsLiveIndexedLogsUsageAggSum to be an explicit nil.
func (o *LogsRetentionAggSumUsage) SetLogsLiveIndexedLogsUsageAggSumNil() {
	o.LogsLiveIndexedLogsUsageAggSum.Set(nil)
}

// UnsetLogsLiveIndexedLogsUsageAggSum ensures that no value is present for LogsLiveIndexedLogsUsageAggSum, not even an explicit nil.
func (o *LogsRetentionAggSumUsage) UnsetLogsLiveIndexedLogsUsageAggSum() {
	o.LogsLiveIndexedLogsUsageAggSum.Unset()
}

// GetLogsRehydratedIndexedLogsUsageAggSum returns the LogsRehydratedIndexedLogsUsageAggSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LogsRetentionAggSumUsage) GetLogsRehydratedIndexedLogsUsageAggSum() int64 {
	if o == nil || o.LogsRehydratedIndexedLogsUsageAggSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.LogsRehydratedIndexedLogsUsageAggSum.Get()
}

// GetLogsRehydratedIndexedLogsUsageAggSumOk returns a tuple with the LogsRehydratedIndexedLogsUsageAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LogsRetentionAggSumUsage) GetLogsRehydratedIndexedLogsUsageAggSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LogsRehydratedIndexedLogsUsageAggSum.Get(), o.LogsRehydratedIndexedLogsUsageAggSum.IsSet()
}

// HasLogsRehydratedIndexedLogsUsageAggSum returns a boolean if a field has been set.
func (o *LogsRetentionAggSumUsage) HasLogsRehydratedIndexedLogsUsageAggSum() bool {
	return o != nil && o.LogsRehydratedIndexedLogsUsageAggSum.IsSet()
}

// SetLogsRehydratedIndexedLogsUsageAggSum gets a reference to the given datadog.NullableInt64 and assigns it to the LogsRehydratedIndexedLogsUsageAggSum field.
func (o *LogsRetentionAggSumUsage) SetLogsRehydratedIndexedLogsUsageAggSum(v int64) {
	o.LogsRehydratedIndexedLogsUsageAggSum.Set(&v)
}

// SetLogsRehydratedIndexedLogsUsageAggSumNil sets the value for LogsRehydratedIndexedLogsUsageAggSum to be an explicit nil.
func (o *LogsRetentionAggSumUsage) SetLogsRehydratedIndexedLogsUsageAggSumNil() {
	o.LogsRehydratedIndexedLogsUsageAggSum.Set(nil)
}

// UnsetLogsRehydratedIndexedLogsUsageAggSum ensures that no value is present for LogsRehydratedIndexedLogsUsageAggSum, not even an explicit nil.
func (o *LogsRetentionAggSumUsage) UnsetLogsRehydratedIndexedLogsUsageAggSum() {
	o.LogsRehydratedIndexedLogsUsageAggSum.Unset()
}

// GetRetention returns the Retention field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LogsRetentionAggSumUsage) GetRetention() string {
	if o == nil || o.Retention.Get() == nil {
		var ret string
		return ret
	}
	return *o.Retention.Get()
}

// GetRetentionOk returns a tuple with the Retention field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LogsRetentionAggSumUsage) GetRetentionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Retention.Get(), o.Retention.IsSet()
}

// HasRetention returns a boolean if a field has been set.
func (o *LogsRetentionAggSumUsage) HasRetention() bool {
	return o != nil && o.Retention.IsSet()
}

// SetRetention gets a reference to the given datadog.NullableString and assigns it to the Retention field.
func (o *LogsRetentionAggSumUsage) SetRetention(v string) {
	o.Retention.Set(&v)
}

// SetRetentionNil sets the value for Retention to be an explicit nil.
func (o *LogsRetentionAggSumUsage) SetRetentionNil() {
	o.Retention.Set(nil)
}

// UnsetRetention ensures that no value is present for Retention, not even an explicit nil.
func (o *LogsRetentionAggSumUsage) UnsetRetention() {
	o.Retention.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o LogsRetentionAggSumUsage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.LogsIndexedLogsUsageAggSum.IsSet() {
		toSerialize["logs_indexed_logs_usage_agg_sum"] = o.LogsIndexedLogsUsageAggSum.Get()
	}
	if o.LogsLiveIndexedLogsUsageAggSum.IsSet() {
		toSerialize["logs_live_indexed_logs_usage_agg_sum"] = o.LogsLiveIndexedLogsUsageAggSum.Get()
	}
	if o.LogsRehydratedIndexedLogsUsageAggSum.IsSet() {
		toSerialize["logs_rehydrated_indexed_logs_usage_agg_sum"] = o.LogsRehydratedIndexedLogsUsageAggSum.Get()
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
func (o *LogsRetentionAggSumUsage) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		LogsIndexedLogsUsageAggSum           datadog.NullableInt64  `json:"logs_indexed_logs_usage_agg_sum,omitempty"`
		LogsLiveIndexedLogsUsageAggSum       datadog.NullableInt64  `json:"logs_live_indexed_logs_usage_agg_sum,omitempty"`
		LogsRehydratedIndexedLogsUsageAggSum datadog.NullableInt64  `json:"logs_rehydrated_indexed_logs_usage_agg_sum,omitempty"`
		Retention                            datadog.NullableString `json:"retention,omitempty"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"logs_indexed_logs_usage_agg_sum", "logs_live_indexed_logs_usage_agg_sum", "logs_rehydrated_indexed_logs_usage_agg_sum", "retention"})
	} else {
		return err
	}
	o.LogsIndexedLogsUsageAggSum = all.LogsIndexedLogsUsageAggSum
	o.LogsLiveIndexedLogsUsageAggSum = all.LogsLiveIndexedLogsUsageAggSum
	o.LogsRehydratedIndexedLogsUsageAggSum = all.LogsRehydratedIndexedLogsUsageAggSum
	o.Retention = all.Retention
	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
