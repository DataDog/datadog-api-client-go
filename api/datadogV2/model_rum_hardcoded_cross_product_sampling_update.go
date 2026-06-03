// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumHardcodedCrossProductSamplingUpdate Partial update for cross-product retention of a hardcoded retention filter.
// Only fields whose matching flag in `cross_product_sampling_editability` is `true` can be updated.
type RumHardcodedCrossProductSamplingUpdate struct {
	// Controls whether Session Replay cross-product retention is active. Omit to leave unchanged.
	SessionReplayEnabled *bool `json:"session_replay_enabled,omitempty"`
	// Percentage (0–100) of retained sessions with an ingested replay whose replay data is kept.
	// Omit to leave unchanged.
	SessionReplaySampleRate *float64 `json:"session_replay_sample_rate,omitempty"`
	// Controls whether Trace cross-product retention is active. Omit to leave unchanged.
	TraceEnabled *bool `json:"trace_enabled,omitempty"`
	// Percentage (0–100) of retained sessions with ingested traces whose traces are indexed.
	// Omit to leave unchanged.
	TraceSampleRate *float64 `json:"trace_sample_rate,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRumHardcodedCrossProductSamplingUpdate instantiates a new RumHardcodedCrossProductSamplingUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRumHardcodedCrossProductSamplingUpdate() *RumHardcodedCrossProductSamplingUpdate {
	this := RumHardcodedCrossProductSamplingUpdate{}
	return &this
}

// NewRumHardcodedCrossProductSamplingUpdateWithDefaults instantiates a new RumHardcodedCrossProductSamplingUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRumHardcodedCrossProductSamplingUpdateWithDefaults() *RumHardcodedCrossProductSamplingUpdate {
	this := RumHardcodedCrossProductSamplingUpdate{}
	return &this
}

// GetSessionReplayEnabled returns the SessionReplayEnabled field value if set, zero value otherwise.
func (o *RumHardcodedCrossProductSamplingUpdate) GetSessionReplayEnabled() bool {
	if o == nil || o.SessionReplayEnabled == nil {
		var ret bool
		return ret
	}
	return *o.SessionReplayEnabled
}

// GetSessionReplayEnabledOk returns a tuple with the SessionReplayEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumHardcodedCrossProductSamplingUpdate) GetSessionReplayEnabledOk() (*bool, bool) {
	if o == nil || o.SessionReplayEnabled == nil {
		return nil, false
	}
	return o.SessionReplayEnabled, true
}

// HasSessionReplayEnabled returns a boolean if a field has been set.
func (o *RumHardcodedCrossProductSamplingUpdate) HasSessionReplayEnabled() bool {
	return o != nil && o.SessionReplayEnabled != nil
}

// SetSessionReplayEnabled gets a reference to the given bool and assigns it to the SessionReplayEnabled field.
func (o *RumHardcodedCrossProductSamplingUpdate) SetSessionReplayEnabled(v bool) {
	o.SessionReplayEnabled = &v
}

// GetSessionReplaySampleRate returns the SessionReplaySampleRate field value if set, zero value otherwise.
func (o *RumHardcodedCrossProductSamplingUpdate) GetSessionReplaySampleRate() float64 {
	if o == nil || o.SessionReplaySampleRate == nil {
		var ret float64
		return ret
	}
	return *o.SessionReplaySampleRate
}

// GetSessionReplaySampleRateOk returns a tuple with the SessionReplaySampleRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumHardcodedCrossProductSamplingUpdate) GetSessionReplaySampleRateOk() (*float64, bool) {
	if o == nil || o.SessionReplaySampleRate == nil {
		return nil, false
	}
	return o.SessionReplaySampleRate, true
}

// HasSessionReplaySampleRate returns a boolean if a field has been set.
func (o *RumHardcodedCrossProductSamplingUpdate) HasSessionReplaySampleRate() bool {
	return o != nil && o.SessionReplaySampleRate != nil
}

// SetSessionReplaySampleRate gets a reference to the given float64 and assigns it to the SessionReplaySampleRate field.
func (o *RumHardcodedCrossProductSamplingUpdate) SetSessionReplaySampleRate(v float64) {
	o.SessionReplaySampleRate = &v
}

// GetTraceEnabled returns the TraceEnabled field value if set, zero value otherwise.
func (o *RumHardcodedCrossProductSamplingUpdate) GetTraceEnabled() bool {
	if o == nil || o.TraceEnabled == nil {
		var ret bool
		return ret
	}
	return *o.TraceEnabled
}

// GetTraceEnabledOk returns a tuple with the TraceEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumHardcodedCrossProductSamplingUpdate) GetTraceEnabledOk() (*bool, bool) {
	if o == nil || o.TraceEnabled == nil {
		return nil, false
	}
	return o.TraceEnabled, true
}

// HasTraceEnabled returns a boolean if a field has been set.
func (o *RumHardcodedCrossProductSamplingUpdate) HasTraceEnabled() bool {
	return o != nil && o.TraceEnabled != nil
}

// SetTraceEnabled gets a reference to the given bool and assigns it to the TraceEnabled field.
func (o *RumHardcodedCrossProductSamplingUpdate) SetTraceEnabled(v bool) {
	o.TraceEnabled = &v
}

// GetTraceSampleRate returns the TraceSampleRate field value if set, zero value otherwise.
func (o *RumHardcodedCrossProductSamplingUpdate) GetTraceSampleRate() float64 {
	if o == nil || o.TraceSampleRate == nil {
		var ret float64
		return ret
	}
	return *o.TraceSampleRate
}

// GetTraceSampleRateOk returns a tuple with the TraceSampleRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumHardcodedCrossProductSamplingUpdate) GetTraceSampleRateOk() (*float64, bool) {
	if o == nil || o.TraceSampleRate == nil {
		return nil, false
	}
	return o.TraceSampleRate, true
}

// HasTraceSampleRate returns a boolean if a field has been set.
func (o *RumHardcodedCrossProductSamplingUpdate) HasTraceSampleRate() bool {
	return o != nil && o.TraceSampleRate != nil
}

// SetTraceSampleRate gets a reference to the given float64 and assigns it to the TraceSampleRate field.
func (o *RumHardcodedCrossProductSamplingUpdate) SetTraceSampleRate(v float64) {
	o.TraceSampleRate = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RumHardcodedCrossProductSamplingUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.SessionReplayEnabled != nil {
		toSerialize["session_replay_enabled"] = o.SessionReplayEnabled
	}
	if o.SessionReplaySampleRate != nil {
		toSerialize["session_replay_sample_rate"] = o.SessionReplaySampleRate
	}
	if o.TraceEnabled != nil {
		toSerialize["trace_enabled"] = o.TraceEnabled
	}
	if o.TraceSampleRate != nil {
		toSerialize["trace_sample_rate"] = o.TraceSampleRate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RumHardcodedCrossProductSamplingUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		SessionReplayEnabled    *bool    `json:"session_replay_enabled,omitempty"`
		SessionReplaySampleRate *float64 `json:"session_replay_sample_rate,omitempty"`
		TraceEnabled            *bool    `json:"trace_enabled,omitempty"`
		TraceSampleRate         *float64 `json:"trace_sample_rate,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"session_replay_enabled", "session_replay_sample_rate", "trace_enabled", "trace_sample_rate"})
	} else {
		return err
	}
	o.SessionReplayEnabled = all.SessionReplayEnabled
	o.SessionReplaySampleRate = all.SessionReplaySampleRate
	o.TraceEnabled = all.TraceEnabled
	o.TraceSampleRate = all.TraceSampleRate

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
