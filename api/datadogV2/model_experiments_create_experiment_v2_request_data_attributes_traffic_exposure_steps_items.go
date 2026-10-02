// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsCreateExperimentV2RequestDataAttributesTrafficExposureStepsItems Configured exposure plan rather than wall-clock history. At least two steps must have strictly increasing fractions and no gaps. Warehouse steps start at assignments_start_date and can use different durations. New Datadog plans have at most five steps and a first fraction above zero. Their nonfinal durations must be equal and exclude time paused. Datadog steps start with the experiment. Running warehouse experiments can replace step fractions, durations, and the exposure mode. After start, Datadog exposure plans cannot change through the public API. The final duration is null and its fraction holds until assignment ends.
type ExperimentsCreateExperimentV2RequestDataAttributesTrafficExposureStepsItems struct {
	// Positive step duration in milliseconds. Datadog durations exclude pauses. Send null for the final step.
	DurationMs datadog.NullableInt64 `json:"duration_ms"`
	// Fraction of traffic exposed during this step.
	Fraction float64 `json:"fraction"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsCreateExperimentV2RequestDataAttributesTrafficExposureStepsItems instantiates a new ExperimentsCreateExperimentV2RequestDataAttributesTrafficExposureStepsItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsCreateExperimentV2RequestDataAttributesTrafficExposureStepsItems(durationMs datadog.NullableInt64, fraction float64) *ExperimentsCreateExperimentV2RequestDataAttributesTrafficExposureStepsItems {
	this := ExperimentsCreateExperimentV2RequestDataAttributesTrafficExposureStepsItems{}
	this.DurationMs = durationMs
	this.Fraction = fraction
	return &this
}

// NewExperimentsCreateExperimentV2RequestDataAttributesTrafficExposureStepsItemsWithDefaults instantiates a new ExperimentsCreateExperimentV2RequestDataAttributesTrafficExposureStepsItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsCreateExperimentV2RequestDataAttributesTrafficExposureStepsItemsWithDefaults() *ExperimentsCreateExperimentV2RequestDataAttributesTrafficExposureStepsItems {
	this := ExperimentsCreateExperimentV2RequestDataAttributesTrafficExposureStepsItems{}
	return &this
}

// GetDurationMs returns the DurationMs field value.
// If the value is explicit nil, the zero value for int64 will be returned.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesTrafficExposureStepsItems) GetDurationMs() int64 {
	if o == nil || o.DurationMs.Get() == nil {
		var ret int64
		return ret
	}
	return *o.DurationMs.Get()
}

// GetDurationMsOk returns a tuple with the DurationMs field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesTrafficExposureStepsItems) GetDurationMsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DurationMs.Get(), o.DurationMs.IsSet()
}

// SetDurationMs sets field value.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesTrafficExposureStepsItems) SetDurationMs(v int64) {
	o.DurationMs.Set(&v)
}

// GetFraction returns the Fraction field value.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesTrafficExposureStepsItems) GetFraction() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Fraction
}

// GetFractionOk returns a tuple with the Fraction field value
// and a boolean to check if the value has been set.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesTrafficExposureStepsItems) GetFractionOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fraction, true
}

// SetFraction sets field value.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesTrafficExposureStepsItems) SetFraction(v float64) {
	o.Fraction = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsCreateExperimentV2RequestDataAttributesTrafficExposureStepsItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["duration_ms"] = o.DurationMs.Get()
	toSerialize["fraction"] = o.Fraction

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesTrafficExposureStepsItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DurationMs datadog.NullableInt64 `json:"duration_ms"`
		Fraction   *float64              `json:"fraction"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if !all.DurationMs.IsSet() {
		return fmt.Errorf("required field duration_ms missing")
	}
	if all.Fraction == nil {
		return fmt.Errorf("required field fraction missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"duration_ms", "fraction"})
	} else {
		return err
	}
	o.DurationMs = all.DurationMs
	o.Fraction = *all.Fraction

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
