// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsPatchExperimentV2ResponseDataAttributesTrafficExposure Traffic exposure fraction or schedule configured for the experiment.
type ExperimentsPatchExperimentV2ResponseDataAttributesTrafficExposure struct {
	// STATIC exposure fraction. Draft experiments can change this value. After start only warehouse experiments without a Datadog flag can change a STATIC fraction through the public API.
	Fraction *float64 `json:"fraction,omitempty"`
	// Whether exposure uses a fixed fraction or a sequence of steps.
	Mode ExperimentsCreateExperimentV2RequestDataAttributesTrafficExposureMode `json:"mode"`
	// Configured exposure plan rather than wall-clock history. At least two steps must have strictly increasing fractions and no gaps. Warehouse steps start at assignments_start_date and can use different durations. New Datadog plans have at most five steps and a first fraction above zero. Their nonfinal durations must be equal and exclude time paused. Datadog steps start with the experiment. Running warehouse experiments can replace step fractions, durations, and the exposure mode. After start, Datadog exposure plans cannot change through the public API. The final duration is null and its fraction holds until assignment ends.
	Steps []ExperimentsCreateExperimentV2RequestDataAttributesTrafficExposureStepsItems `json:"steps,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsPatchExperimentV2ResponseDataAttributesTrafficExposure instantiates a new ExperimentsPatchExperimentV2ResponseDataAttributesTrafficExposure object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsPatchExperimentV2ResponseDataAttributesTrafficExposure(mode ExperimentsCreateExperimentV2RequestDataAttributesTrafficExposureMode) *ExperimentsPatchExperimentV2ResponseDataAttributesTrafficExposure {
	this := ExperimentsPatchExperimentV2ResponseDataAttributesTrafficExposure{}
	this.Mode = mode
	return &this
}

// NewExperimentsPatchExperimentV2ResponseDataAttributesTrafficExposureWithDefaults instantiates a new ExperimentsPatchExperimentV2ResponseDataAttributesTrafficExposure object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsPatchExperimentV2ResponseDataAttributesTrafficExposureWithDefaults() *ExperimentsPatchExperimentV2ResponseDataAttributesTrafficExposure {
	this := ExperimentsPatchExperimentV2ResponseDataAttributesTrafficExposure{}
	return &this
}

// GetFraction returns the Fraction field value if set, zero value otherwise.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributesTrafficExposure) GetFraction() float64 {
	if o == nil || o.Fraction == nil {
		var ret float64
		return ret
	}
	return *o.Fraction
}

// GetFractionOk returns a tuple with the Fraction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributesTrafficExposure) GetFractionOk() (*float64, bool) {
	if o == nil || o.Fraction == nil {
		return nil, false
	}
	return o.Fraction, true
}

// HasFraction returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributesTrafficExposure) HasFraction() bool {
	return o != nil && o.Fraction != nil
}

// SetFraction gets a reference to the given float64 and assigns it to the Fraction field.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributesTrafficExposure) SetFraction(v float64) {
	o.Fraction = &v
}

// GetMode returns the Mode field value.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributesTrafficExposure) GetMode() ExperimentsCreateExperimentV2RequestDataAttributesTrafficExposureMode {
	if o == nil {
		var ret ExperimentsCreateExperimentV2RequestDataAttributesTrafficExposureMode
		return ret
	}
	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributesTrafficExposure) GetModeOk() (*ExperimentsCreateExperimentV2RequestDataAttributesTrafficExposureMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributesTrafficExposure) SetMode(v ExperimentsCreateExperimentV2RequestDataAttributesTrafficExposureMode) {
	o.Mode = v
}

// GetSteps returns the Steps field value if set, zero value otherwise.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributesTrafficExposure) GetSteps() []ExperimentsCreateExperimentV2RequestDataAttributesTrafficExposureStepsItems {
	if o == nil || o.Steps == nil {
		var ret []ExperimentsCreateExperimentV2RequestDataAttributesTrafficExposureStepsItems
		return ret
	}
	return o.Steps
}

// GetStepsOk returns a tuple with the Steps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributesTrafficExposure) GetStepsOk() (*[]ExperimentsCreateExperimentV2RequestDataAttributesTrafficExposureStepsItems, bool) {
	if o == nil || o.Steps == nil {
		return nil, false
	}
	return &o.Steps, true
}

// HasSteps returns a boolean if a field has been set.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributesTrafficExposure) HasSteps() bool {
	return o != nil && o.Steps != nil
}

// SetSteps gets a reference to the given []ExperimentsCreateExperimentV2RequestDataAttributesTrafficExposureStepsItems and assigns it to the Steps field.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributesTrafficExposure) SetSteps(v []ExperimentsCreateExperimentV2RequestDataAttributesTrafficExposureStepsItems) {
	o.Steps = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsPatchExperimentV2ResponseDataAttributesTrafficExposure) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Fraction != nil {
		toSerialize["fraction"] = o.Fraction
	}
	toSerialize["mode"] = o.Mode
	if o.Steps != nil {
		toSerialize["steps"] = o.Steps
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsPatchExperimentV2ResponseDataAttributesTrafficExposure) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Fraction *float64                                                                      `json:"fraction,omitempty"`
		Mode     *ExperimentsCreateExperimentV2RequestDataAttributesTrafficExposureMode        `json:"mode"`
		Steps    []ExperimentsCreateExperimentV2RequestDataAttributesTrafficExposureStepsItems `json:"steps,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Mode == nil {
		return fmt.Errorf("required field mode missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"fraction", "mode", "steps"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Fraction = all.Fraction
	if !all.Mode.IsValid() {
		hasInvalidField = true
	} else {
		o.Mode = *all.Mode
	}
	o.Steps = all.Steps

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
