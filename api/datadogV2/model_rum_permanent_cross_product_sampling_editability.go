// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumPermanentCrossProductSamplingEditability Flags indicating which `cross_product_sampling` rates can be updated for this filter. Read-only.
type RumPermanentCrossProductSamplingEditability struct {
	// If `true`, `cross_product_sampling.trace_sample_rate` can be updated on this filter.
	TraceSampleRate *bool `json:"trace_sample_rate,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRumPermanentCrossProductSamplingEditability instantiates a new RumPermanentCrossProductSamplingEditability object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRumPermanentCrossProductSamplingEditability() *RumPermanentCrossProductSamplingEditability {
	this := RumPermanentCrossProductSamplingEditability{}
	return &this
}

// NewRumPermanentCrossProductSamplingEditabilityWithDefaults instantiates a new RumPermanentCrossProductSamplingEditability object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRumPermanentCrossProductSamplingEditabilityWithDefaults() *RumPermanentCrossProductSamplingEditability {
	this := RumPermanentCrossProductSamplingEditability{}
	return &this
}

// GetTraceSampleRate returns the TraceSampleRate field value if set, zero value otherwise.
func (o *RumPermanentCrossProductSamplingEditability) GetTraceSampleRate() bool {
	if o == nil || o.TraceSampleRate == nil {
		var ret bool
		return ret
	}
	return *o.TraceSampleRate
}

// GetTraceSampleRateOk returns a tuple with the TraceSampleRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumPermanentCrossProductSamplingEditability) GetTraceSampleRateOk() (*bool, bool) {
	if o == nil || o.TraceSampleRate == nil {
		return nil, false
	}
	return o.TraceSampleRate, true
}

// HasTraceSampleRate returns a boolean if a field has been set.
func (o *RumPermanentCrossProductSamplingEditability) HasTraceSampleRate() bool {
	return o != nil && o.TraceSampleRate != nil
}

// SetTraceSampleRate gets a reference to the given bool and assigns it to the TraceSampleRate field.
func (o *RumPermanentCrossProductSamplingEditability) SetTraceSampleRate(v bool) {
	o.TraceSampleRate = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RumPermanentCrossProductSamplingEditability) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
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
func (o *RumPermanentCrossProductSamplingEditability) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TraceSampleRate *bool `json:"trace_sample_rate,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"trace_sample_rate"})
	} else {
		return err
	}
	o.TraceSampleRate = all.TraceSampleRate

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
