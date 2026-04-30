// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumPermanentCrossProductSamplingUpdate Partial update of the cross-product sample rates for a permanent retention filter.
// Only rates with a matching `cross_product_sampling_editability` flag set to `true` can be updated.
type RumPermanentCrossProductSamplingUpdate struct {
	// Percentage (0–100) of retained sessions (with ingested traces) whose traces are indexed.
	// Omit to leave unchanged. Rejected if the filter's `cross_product_sampling_editability.trace_sample_rate` is `false`.
	TraceSampleRate *float64 `json:"trace_sample_rate,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRumPermanentCrossProductSamplingUpdate instantiates a new RumPermanentCrossProductSamplingUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRumPermanentCrossProductSamplingUpdate() *RumPermanentCrossProductSamplingUpdate {
	this := RumPermanentCrossProductSamplingUpdate{}
	return &this
}

// NewRumPermanentCrossProductSamplingUpdateWithDefaults instantiates a new RumPermanentCrossProductSamplingUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRumPermanentCrossProductSamplingUpdateWithDefaults() *RumPermanentCrossProductSamplingUpdate {
	this := RumPermanentCrossProductSamplingUpdate{}
	return &this
}

// GetTraceSampleRate returns the TraceSampleRate field value if set, zero value otherwise.
func (o *RumPermanentCrossProductSamplingUpdate) GetTraceSampleRate() float64 {
	if o == nil || o.TraceSampleRate == nil {
		var ret float64
		return ret
	}
	return *o.TraceSampleRate
}

// GetTraceSampleRateOk returns a tuple with the TraceSampleRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumPermanentCrossProductSamplingUpdate) GetTraceSampleRateOk() (*float64, bool) {
	if o == nil || o.TraceSampleRate == nil {
		return nil, false
	}
	return o.TraceSampleRate, true
}

// HasTraceSampleRate returns a boolean if a field has been set.
func (o *RumPermanentCrossProductSamplingUpdate) HasTraceSampleRate() bool {
	return o != nil && o.TraceSampleRate != nil
}

// SetTraceSampleRate gets a reference to the given float64 and assigns it to the TraceSampleRate field.
func (o *RumPermanentCrossProductSamplingUpdate) SetTraceSampleRate(v float64) {
	o.TraceSampleRate = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RumPermanentCrossProductSamplingUpdate) MarshalJSON() ([]byte, error) {
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
func (o *RumPermanentCrossProductSamplingUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TraceSampleRate *float64 `json:"trace_sample_rate,omitempty"`
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
