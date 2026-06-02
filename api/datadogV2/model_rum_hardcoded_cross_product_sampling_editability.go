// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumHardcodedCrossProductSamplingEditability Flags indicating which `cross_product_sampling` fields can be updated. Read-only.
type RumHardcodedCrossProductSamplingEditability struct {
	// If `true`, `cross_product_sampling.session_replay_sample_rate` can be updated on this filter.
	SessionReplaySampleRate *bool `json:"session_replay_sample_rate,omitempty"`
	// If `true`, `cross_product_sampling.trace_sample_rate` can be updated on this filter.
	TraceSampleRate *bool `json:"trace_sample_rate,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRumHardcodedCrossProductSamplingEditability instantiates a new RumHardcodedCrossProductSamplingEditability object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRumHardcodedCrossProductSamplingEditability() *RumHardcodedCrossProductSamplingEditability {
	this := RumHardcodedCrossProductSamplingEditability{}
	return &this
}

// NewRumHardcodedCrossProductSamplingEditabilityWithDefaults instantiates a new RumHardcodedCrossProductSamplingEditability object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRumHardcodedCrossProductSamplingEditabilityWithDefaults() *RumHardcodedCrossProductSamplingEditability {
	this := RumHardcodedCrossProductSamplingEditability{}
	return &this
}

// GetSessionReplaySampleRate returns the SessionReplaySampleRate field value if set, zero value otherwise.
func (o *RumHardcodedCrossProductSamplingEditability) GetSessionReplaySampleRate() bool {
	if o == nil || o.SessionReplaySampleRate == nil {
		var ret bool
		return ret
	}
	return *o.SessionReplaySampleRate
}

// GetSessionReplaySampleRateOk returns a tuple with the SessionReplaySampleRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumHardcodedCrossProductSamplingEditability) GetSessionReplaySampleRateOk() (*bool, bool) {
	if o == nil || o.SessionReplaySampleRate == nil {
		return nil, false
	}
	return o.SessionReplaySampleRate, true
}

// HasSessionReplaySampleRate returns a boolean if a field has been set.
func (o *RumHardcodedCrossProductSamplingEditability) HasSessionReplaySampleRate() bool {
	return o != nil && o.SessionReplaySampleRate != nil
}

// SetSessionReplaySampleRate gets a reference to the given bool and assigns it to the SessionReplaySampleRate field.
func (o *RumHardcodedCrossProductSamplingEditability) SetSessionReplaySampleRate(v bool) {
	o.SessionReplaySampleRate = &v
}

// GetTraceSampleRate returns the TraceSampleRate field value if set, zero value otherwise.
func (o *RumHardcodedCrossProductSamplingEditability) GetTraceSampleRate() bool {
	if o == nil || o.TraceSampleRate == nil {
		var ret bool
		return ret
	}
	return *o.TraceSampleRate
}

// GetTraceSampleRateOk returns a tuple with the TraceSampleRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumHardcodedCrossProductSamplingEditability) GetTraceSampleRateOk() (*bool, bool) {
	if o == nil || o.TraceSampleRate == nil {
		return nil, false
	}
	return o.TraceSampleRate, true
}

// HasTraceSampleRate returns a boolean if a field has been set.
func (o *RumHardcodedCrossProductSamplingEditability) HasTraceSampleRate() bool {
	return o != nil && o.TraceSampleRate != nil
}

// SetTraceSampleRate gets a reference to the given bool and assigns it to the TraceSampleRate field.
func (o *RumHardcodedCrossProductSamplingEditability) SetTraceSampleRate(v bool) {
	o.TraceSampleRate = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RumHardcodedCrossProductSamplingEditability) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.SessionReplaySampleRate != nil {
		toSerialize["session_replay_sample_rate"] = o.SessionReplaySampleRate
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
func (o *RumHardcodedCrossProductSamplingEditability) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		SessionReplaySampleRate *bool `json:"session_replay_sample_rate,omitempty"`
		TraceSampleRate         *bool `json:"trace_sample_rate,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"session_replay_sample_rate", "trace_sample_rate"})
	} else {
		return err
	}
	o.SessionReplaySampleRate = all.SessionReplaySampleRate
	o.TraceSampleRate = all.TraceSampleRate

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
