// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsAnalysisPlanV2MutationResponseDataAttributesBayesianPrior Parameters of the prior distribution used for Bayesian analysis.
type ExperimentsAnalysisPlanV2MutationResponseDataAttributesBayesianPrior struct {
	// Degrees of freedom of the prior distribution.
	DegreesOfFreedom datadog.NullableFloat64 `json:"degrees_of_freedom,omitempty"`
	// Standard deviation of the prior distribution.
	StandardDeviation *float64 `json:"standard_deviation,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsAnalysisPlanV2MutationResponseDataAttributesBayesianPrior instantiates a new ExperimentsAnalysisPlanV2MutationResponseDataAttributesBayesianPrior object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsAnalysisPlanV2MutationResponseDataAttributesBayesianPrior() *ExperimentsAnalysisPlanV2MutationResponseDataAttributesBayesianPrior {
	this := ExperimentsAnalysisPlanV2MutationResponseDataAttributesBayesianPrior{}
	return &this
}

// NewExperimentsAnalysisPlanV2MutationResponseDataAttributesBayesianPriorWithDefaults instantiates a new ExperimentsAnalysisPlanV2MutationResponseDataAttributesBayesianPrior object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsAnalysisPlanV2MutationResponseDataAttributesBayesianPriorWithDefaults() *ExperimentsAnalysisPlanV2MutationResponseDataAttributesBayesianPrior {
	this := ExperimentsAnalysisPlanV2MutationResponseDataAttributesBayesianPrior{}
	return &this
}

// GetDegreesOfFreedom returns the DegreesOfFreedom field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsAnalysisPlanV2MutationResponseDataAttributesBayesianPrior) GetDegreesOfFreedom() float64 {
	if o == nil || o.DegreesOfFreedom.Get() == nil {
		var ret float64
		return ret
	}
	return *o.DegreesOfFreedom.Get()
}

// GetDegreesOfFreedomOk returns a tuple with the DegreesOfFreedom field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsAnalysisPlanV2MutationResponseDataAttributesBayesianPrior) GetDegreesOfFreedomOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DegreesOfFreedom.Get(), o.DegreesOfFreedom.IsSet()
}

// HasDegreesOfFreedom returns a boolean if a field has been set.
func (o *ExperimentsAnalysisPlanV2MutationResponseDataAttributesBayesianPrior) HasDegreesOfFreedom() bool {
	return o != nil && o.DegreesOfFreedom.IsSet()
}

// SetDegreesOfFreedom gets a reference to the given datadog.NullableFloat64 and assigns it to the DegreesOfFreedom field.
func (o *ExperimentsAnalysisPlanV2MutationResponseDataAttributesBayesianPrior) SetDegreesOfFreedom(v float64) {
	o.DegreesOfFreedom.Set(&v)
}

// SetDegreesOfFreedomNil sets the value for DegreesOfFreedom to be an explicit nil.
func (o *ExperimentsAnalysisPlanV2MutationResponseDataAttributesBayesianPrior) SetDegreesOfFreedomNil() {
	o.DegreesOfFreedom.Set(nil)
}

// UnsetDegreesOfFreedom ensures that no value is present for DegreesOfFreedom, not even an explicit nil.
func (o *ExperimentsAnalysisPlanV2MutationResponseDataAttributesBayesianPrior) UnsetDegreesOfFreedom() {
	o.DegreesOfFreedom.Unset()
}

// GetStandardDeviation returns the StandardDeviation field value if set, zero value otherwise.
func (o *ExperimentsAnalysisPlanV2MutationResponseDataAttributesBayesianPrior) GetStandardDeviation() float64 {
	if o == nil || o.StandardDeviation == nil {
		var ret float64
		return ret
	}
	return *o.StandardDeviation
}

// GetStandardDeviationOk returns a tuple with the StandardDeviation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsAnalysisPlanV2MutationResponseDataAttributesBayesianPrior) GetStandardDeviationOk() (*float64, bool) {
	if o == nil || o.StandardDeviation == nil {
		return nil, false
	}
	return o.StandardDeviation, true
}

// HasStandardDeviation returns a boolean if a field has been set.
func (o *ExperimentsAnalysisPlanV2MutationResponseDataAttributesBayesianPrior) HasStandardDeviation() bool {
	return o != nil && o.StandardDeviation != nil
}

// SetStandardDeviation gets a reference to the given float64 and assigns it to the StandardDeviation field.
func (o *ExperimentsAnalysisPlanV2MutationResponseDataAttributesBayesianPrior) SetStandardDeviation(v float64) {
	o.StandardDeviation = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsAnalysisPlanV2MutationResponseDataAttributesBayesianPrior) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DegreesOfFreedom.IsSet() {
		toSerialize["degrees_of_freedom"] = o.DegreesOfFreedom.Get()
	}
	if o.StandardDeviation != nil {
		toSerialize["standard_deviation"] = o.StandardDeviation
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsAnalysisPlanV2MutationResponseDataAttributesBayesianPrior) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DegreesOfFreedom  datadog.NullableFloat64 `json:"degrees_of_freedom,omitempty"`
		StandardDeviation *float64                `json:"standard_deviation,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"degrees_of_freedom", "standard_deviation"})
	} else {
		return err
	}
	o.DegreesOfFreedom = all.DegreesOfFreedom
	o.StandardDeviation = all.StandardDeviation

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
