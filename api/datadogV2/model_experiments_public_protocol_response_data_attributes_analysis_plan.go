// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsPublicProtocolResponseDataAttributesAnalysisPlan Default statistical settings supplied by the protocol.
type ExperimentsPublicProtocolResponseDataAttributesAnalysisPlan struct {
	// Whether to use pre-experiment data to reduce variance with CUPED.
	ComputeCuped *bool `json:"compute_cuped,omitempty"`
	// Statistical method used to calculate confidence intervals.
	ConfidenceIntervalMethod *string `json:"confidence_interval_method,omitempty"`
	// Confidence level used by the statistical analysis.
	ConfidenceLevel *float64 `json:"confidence_level,omitempty"`
	// Method used to adjust for testing multiple metrics.
	MultipleTestingCorrectionMethod *string `json:"multiple_testing_correction_method,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsPublicProtocolResponseDataAttributesAnalysisPlan instantiates a new ExperimentsPublicProtocolResponseDataAttributesAnalysisPlan object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsPublicProtocolResponseDataAttributesAnalysisPlan() *ExperimentsPublicProtocolResponseDataAttributesAnalysisPlan {
	this := ExperimentsPublicProtocolResponseDataAttributesAnalysisPlan{}
	return &this
}

// NewExperimentsPublicProtocolResponseDataAttributesAnalysisPlanWithDefaults instantiates a new ExperimentsPublicProtocolResponseDataAttributesAnalysisPlan object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsPublicProtocolResponseDataAttributesAnalysisPlanWithDefaults() *ExperimentsPublicProtocolResponseDataAttributesAnalysisPlan {
	this := ExperimentsPublicProtocolResponseDataAttributesAnalysisPlan{}
	return &this
}

// GetComputeCuped returns the ComputeCuped field value if set, zero value otherwise.
func (o *ExperimentsPublicProtocolResponseDataAttributesAnalysisPlan) GetComputeCuped() bool {
	if o == nil || o.ComputeCuped == nil {
		var ret bool
		return ret
	}
	return *o.ComputeCuped
}

// GetComputeCupedOk returns a tuple with the ComputeCuped field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesAnalysisPlan) GetComputeCupedOk() (*bool, bool) {
	if o == nil || o.ComputeCuped == nil {
		return nil, false
	}
	return o.ComputeCuped, true
}

// HasComputeCuped returns a boolean if a field has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesAnalysisPlan) HasComputeCuped() bool {
	return o != nil && o.ComputeCuped != nil
}

// SetComputeCuped gets a reference to the given bool and assigns it to the ComputeCuped field.
func (o *ExperimentsPublicProtocolResponseDataAttributesAnalysisPlan) SetComputeCuped(v bool) {
	o.ComputeCuped = &v
}

// GetConfidenceIntervalMethod returns the ConfidenceIntervalMethod field value if set, zero value otherwise.
func (o *ExperimentsPublicProtocolResponseDataAttributesAnalysisPlan) GetConfidenceIntervalMethod() string {
	if o == nil || o.ConfidenceIntervalMethod == nil {
		var ret string
		return ret
	}
	return *o.ConfidenceIntervalMethod
}

// GetConfidenceIntervalMethodOk returns a tuple with the ConfidenceIntervalMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesAnalysisPlan) GetConfidenceIntervalMethodOk() (*string, bool) {
	if o == nil || o.ConfidenceIntervalMethod == nil {
		return nil, false
	}
	return o.ConfidenceIntervalMethod, true
}

// HasConfidenceIntervalMethod returns a boolean if a field has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesAnalysisPlan) HasConfidenceIntervalMethod() bool {
	return o != nil && o.ConfidenceIntervalMethod != nil
}

// SetConfidenceIntervalMethod gets a reference to the given string and assigns it to the ConfidenceIntervalMethod field.
func (o *ExperimentsPublicProtocolResponseDataAttributesAnalysisPlan) SetConfidenceIntervalMethod(v string) {
	o.ConfidenceIntervalMethod = &v
}

// GetConfidenceLevel returns the ConfidenceLevel field value if set, zero value otherwise.
func (o *ExperimentsPublicProtocolResponseDataAttributesAnalysisPlan) GetConfidenceLevel() float64 {
	if o == nil || o.ConfidenceLevel == nil {
		var ret float64
		return ret
	}
	return *o.ConfidenceLevel
}

// GetConfidenceLevelOk returns a tuple with the ConfidenceLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesAnalysisPlan) GetConfidenceLevelOk() (*float64, bool) {
	if o == nil || o.ConfidenceLevel == nil {
		return nil, false
	}
	return o.ConfidenceLevel, true
}

// HasConfidenceLevel returns a boolean if a field has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesAnalysisPlan) HasConfidenceLevel() bool {
	return o != nil && o.ConfidenceLevel != nil
}

// SetConfidenceLevel gets a reference to the given float64 and assigns it to the ConfidenceLevel field.
func (o *ExperimentsPublicProtocolResponseDataAttributesAnalysisPlan) SetConfidenceLevel(v float64) {
	o.ConfidenceLevel = &v
}

// GetMultipleTestingCorrectionMethod returns the MultipleTestingCorrectionMethod field value if set, zero value otherwise.
func (o *ExperimentsPublicProtocolResponseDataAttributesAnalysisPlan) GetMultipleTestingCorrectionMethod() string {
	if o == nil || o.MultipleTestingCorrectionMethod == nil {
		var ret string
		return ret
	}
	return *o.MultipleTestingCorrectionMethod
}

// GetMultipleTestingCorrectionMethodOk returns a tuple with the MultipleTestingCorrectionMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesAnalysisPlan) GetMultipleTestingCorrectionMethodOk() (*string, bool) {
	if o == nil || o.MultipleTestingCorrectionMethod == nil {
		return nil, false
	}
	return o.MultipleTestingCorrectionMethod, true
}

// HasMultipleTestingCorrectionMethod returns a boolean if a field has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesAnalysisPlan) HasMultipleTestingCorrectionMethod() bool {
	return o != nil && o.MultipleTestingCorrectionMethod != nil
}

// SetMultipleTestingCorrectionMethod gets a reference to the given string and assigns it to the MultipleTestingCorrectionMethod field.
func (o *ExperimentsPublicProtocolResponseDataAttributesAnalysisPlan) SetMultipleTestingCorrectionMethod(v string) {
	o.MultipleTestingCorrectionMethod = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsPublicProtocolResponseDataAttributesAnalysisPlan) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ComputeCuped != nil {
		toSerialize["compute_cuped"] = o.ComputeCuped
	}
	if o.ConfidenceIntervalMethod != nil {
		toSerialize["confidence_interval_method"] = o.ConfidenceIntervalMethod
	}
	if o.ConfidenceLevel != nil {
		toSerialize["confidence_level"] = o.ConfidenceLevel
	}
	if o.MultipleTestingCorrectionMethod != nil {
		toSerialize["multiple_testing_correction_method"] = o.MultipleTestingCorrectionMethod
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsPublicProtocolResponseDataAttributesAnalysisPlan) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ComputeCuped                    *bool    `json:"compute_cuped,omitempty"`
		ConfidenceIntervalMethod        *string  `json:"confidence_interval_method,omitempty"`
		ConfidenceLevel                 *float64 `json:"confidence_level,omitempty"`
		MultipleTestingCorrectionMethod *string  `json:"multiple_testing_correction_method,omitempty"`
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
