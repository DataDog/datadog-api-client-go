// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsPatchProtocolAnalysisPlanInput Partial statistical settings. Omitted fields remain unchanged. Null clears the whole plan. Null clears nullable fields; null or an empty string clears confidence_interval_method.
type ExperimentsPatchProtocolAnalysisPlanInput struct {
	// Whether to use pre-experiment data to reduce variance with CUPED.
	ComputeCuped datadog.NullableBool `json:"compute_cuped,omitempty"`
	// Confidence interval method: `Sequential`, `FixedSample`, `Bayesian`, or `SequentialFixedHybrid`. An empty string leaves the method unset.
	ConfidenceIntervalMethod datadog.NullableString `json:"confidence_interval_method,omitempty"`
	// Confidence level used by the statistical analysis.
	ConfidenceLevel datadog.NullableFloat64 `json:"confidence_level,omitempty"`
	// Method used to adjust for testing multiple metrics.
	MultipleTestingCorrectionMethod datadog.NullableString `json:"multiple_testing_correction_method,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsPatchProtocolAnalysisPlanInput instantiates a new ExperimentsPatchProtocolAnalysisPlanInput object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsPatchProtocolAnalysisPlanInput() *ExperimentsPatchProtocolAnalysisPlanInput {
	this := ExperimentsPatchProtocolAnalysisPlanInput{}
	return &this
}

// NewExperimentsPatchProtocolAnalysisPlanInputWithDefaults instantiates a new ExperimentsPatchProtocolAnalysisPlanInput object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsPatchProtocolAnalysisPlanInputWithDefaults() *ExperimentsPatchProtocolAnalysisPlanInput {
	this := ExperimentsPatchProtocolAnalysisPlanInput{}
	return &this
}

// GetComputeCuped returns the ComputeCuped field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPatchProtocolAnalysisPlanInput) GetComputeCuped() bool {
	if o == nil || o.ComputeCuped.Get() == nil {
		var ret bool
		return ret
	}
	return *o.ComputeCuped.Get()
}

// GetComputeCupedOk returns a tuple with the ComputeCuped field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPatchProtocolAnalysisPlanInput) GetComputeCupedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ComputeCuped.Get(), o.ComputeCuped.IsSet()
}

// HasComputeCuped returns a boolean if a field has been set.
func (o *ExperimentsPatchProtocolAnalysisPlanInput) HasComputeCuped() bool {
	return o != nil && o.ComputeCuped.IsSet()
}

// SetComputeCuped gets a reference to the given datadog.NullableBool and assigns it to the ComputeCuped field.
func (o *ExperimentsPatchProtocolAnalysisPlanInput) SetComputeCuped(v bool) {
	o.ComputeCuped.Set(&v)
}

// SetComputeCupedNil sets the value for ComputeCuped to be an explicit nil.
func (o *ExperimentsPatchProtocolAnalysisPlanInput) SetComputeCupedNil() {
	o.ComputeCuped.Set(nil)
}

// UnsetComputeCuped ensures that no value is present for ComputeCuped, not even an explicit nil.
func (o *ExperimentsPatchProtocolAnalysisPlanInput) UnsetComputeCuped() {
	o.ComputeCuped.Unset()
}

// GetConfidenceIntervalMethod returns the ConfidenceIntervalMethod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPatchProtocolAnalysisPlanInput) GetConfidenceIntervalMethod() string {
	if o == nil || o.ConfidenceIntervalMethod.Get() == nil {
		var ret string
		return ret
	}
	return *o.ConfidenceIntervalMethod.Get()
}

// GetConfidenceIntervalMethodOk returns a tuple with the ConfidenceIntervalMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPatchProtocolAnalysisPlanInput) GetConfidenceIntervalMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfidenceIntervalMethod.Get(), o.ConfidenceIntervalMethod.IsSet()
}

// HasConfidenceIntervalMethod returns a boolean if a field has been set.
func (o *ExperimentsPatchProtocolAnalysisPlanInput) HasConfidenceIntervalMethod() bool {
	return o != nil && o.ConfidenceIntervalMethod.IsSet()
}

// SetConfidenceIntervalMethod gets a reference to the given datadog.NullableString and assigns it to the ConfidenceIntervalMethod field.
func (o *ExperimentsPatchProtocolAnalysisPlanInput) SetConfidenceIntervalMethod(v string) {
	o.ConfidenceIntervalMethod.Set(&v)
}

// SetConfidenceIntervalMethodNil sets the value for ConfidenceIntervalMethod to be an explicit nil.
func (o *ExperimentsPatchProtocolAnalysisPlanInput) SetConfidenceIntervalMethodNil() {
	o.ConfidenceIntervalMethod.Set(nil)
}

// UnsetConfidenceIntervalMethod ensures that no value is present for ConfidenceIntervalMethod, not even an explicit nil.
func (o *ExperimentsPatchProtocolAnalysisPlanInput) UnsetConfidenceIntervalMethod() {
	o.ConfidenceIntervalMethod.Unset()
}

// GetConfidenceLevel returns the ConfidenceLevel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPatchProtocolAnalysisPlanInput) GetConfidenceLevel() float64 {
	if o == nil || o.ConfidenceLevel.Get() == nil {
		var ret float64
		return ret
	}
	return *o.ConfidenceLevel.Get()
}

// GetConfidenceLevelOk returns a tuple with the ConfidenceLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPatchProtocolAnalysisPlanInput) GetConfidenceLevelOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfidenceLevel.Get(), o.ConfidenceLevel.IsSet()
}

// HasConfidenceLevel returns a boolean if a field has been set.
func (o *ExperimentsPatchProtocolAnalysisPlanInput) HasConfidenceLevel() bool {
	return o != nil && o.ConfidenceLevel.IsSet()
}

// SetConfidenceLevel gets a reference to the given datadog.NullableFloat64 and assigns it to the ConfidenceLevel field.
func (o *ExperimentsPatchProtocolAnalysisPlanInput) SetConfidenceLevel(v float64) {
	o.ConfidenceLevel.Set(&v)
}

// SetConfidenceLevelNil sets the value for ConfidenceLevel to be an explicit nil.
func (o *ExperimentsPatchProtocolAnalysisPlanInput) SetConfidenceLevelNil() {
	o.ConfidenceLevel.Set(nil)
}

// UnsetConfidenceLevel ensures that no value is present for ConfidenceLevel, not even an explicit nil.
func (o *ExperimentsPatchProtocolAnalysisPlanInput) UnsetConfidenceLevel() {
	o.ConfidenceLevel.Unset()
}

// GetMultipleTestingCorrectionMethod returns the MultipleTestingCorrectionMethod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPatchProtocolAnalysisPlanInput) GetMultipleTestingCorrectionMethod() string {
	if o == nil || o.MultipleTestingCorrectionMethod.Get() == nil {
		var ret string
		return ret
	}
	return *o.MultipleTestingCorrectionMethod.Get()
}

// GetMultipleTestingCorrectionMethodOk returns a tuple with the MultipleTestingCorrectionMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPatchProtocolAnalysisPlanInput) GetMultipleTestingCorrectionMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MultipleTestingCorrectionMethod.Get(), o.MultipleTestingCorrectionMethod.IsSet()
}

// HasMultipleTestingCorrectionMethod returns a boolean if a field has been set.
func (o *ExperimentsPatchProtocolAnalysisPlanInput) HasMultipleTestingCorrectionMethod() bool {
	return o != nil && o.MultipleTestingCorrectionMethod.IsSet()
}

// SetMultipleTestingCorrectionMethod gets a reference to the given datadog.NullableString and assigns it to the MultipleTestingCorrectionMethod field.
func (o *ExperimentsPatchProtocolAnalysisPlanInput) SetMultipleTestingCorrectionMethod(v string) {
	o.MultipleTestingCorrectionMethod.Set(&v)
}

// SetMultipleTestingCorrectionMethodNil sets the value for MultipleTestingCorrectionMethod to be an explicit nil.
func (o *ExperimentsPatchProtocolAnalysisPlanInput) SetMultipleTestingCorrectionMethodNil() {
	o.MultipleTestingCorrectionMethod.Set(nil)
}

// UnsetMultipleTestingCorrectionMethod ensures that no value is present for MultipleTestingCorrectionMethod, not even an explicit nil.
func (o *ExperimentsPatchProtocolAnalysisPlanInput) UnsetMultipleTestingCorrectionMethod() {
	o.MultipleTestingCorrectionMethod.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsPatchProtocolAnalysisPlanInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ComputeCuped.IsSet() {
		toSerialize["compute_cuped"] = o.ComputeCuped.Get()
	}
	if o.ConfidenceIntervalMethod.IsSet() {
		toSerialize["confidence_interval_method"] = o.ConfidenceIntervalMethod.Get()
	}
	if o.ConfidenceLevel.IsSet() {
		toSerialize["confidence_level"] = o.ConfidenceLevel.Get()
	}
	if o.MultipleTestingCorrectionMethod.IsSet() {
		toSerialize["multiple_testing_correction_method"] = o.MultipleTestingCorrectionMethod.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsPatchProtocolAnalysisPlanInput) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ComputeCuped                    datadog.NullableBool    `json:"compute_cuped,omitempty"`
		ConfidenceIntervalMethod        datadog.NullableString  `json:"confidence_interval_method,omitempty"`
		ConfidenceLevel                 datadog.NullableFloat64 `json:"confidence_level,omitempty"`
		MultipleTestingCorrectionMethod datadog.NullableString  `json:"multiple_testing_correction_method,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"compute_cuped", "confidence_interval_method", "confidence_level", "multiple_testing_correction_method"})
	} else {
		return err
	}
	o.ComputeCuped = all.ComputeCuped
	o.ConfidenceIntervalMethod = all.ConfidenceIntervalMethod
	o.ConfidenceLevel = all.ConfidenceLevel
	o.MultipleTestingCorrectionMethod = all.MultipleTestingCorrectionMethod

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}

// NullableExperimentsPatchProtocolAnalysisPlanInput handles when a null is used for ExperimentsPatchProtocolAnalysisPlanInput.
type NullableExperimentsPatchProtocolAnalysisPlanInput struct {
	value *ExperimentsPatchProtocolAnalysisPlanInput
	isSet bool
}

// Get returns the associated value.
func (v NullableExperimentsPatchProtocolAnalysisPlanInput) Get() *ExperimentsPatchProtocolAnalysisPlanInput {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableExperimentsPatchProtocolAnalysisPlanInput) Set(val *ExperimentsPatchProtocolAnalysisPlanInput) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableExperimentsPatchProtocolAnalysisPlanInput) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableExperimentsPatchProtocolAnalysisPlanInput) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableExperimentsPatchProtocolAnalysisPlanInput initializes the struct as if Set has been called.
func NewNullableExperimentsPatchProtocolAnalysisPlanInput(val *ExperimentsPatchProtocolAnalysisPlanInput) *NullableExperimentsPatchProtocolAnalysisPlanInput {
	return &NullableExperimentsPatchProtocolAnalysisPlanInput{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableExperimentsPatchProtocolAnalysisPlanInput) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableExperimentsPatchProtocolAnalysisPlanInput) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return datadog.Unmarshal(src, &v.value)
}
