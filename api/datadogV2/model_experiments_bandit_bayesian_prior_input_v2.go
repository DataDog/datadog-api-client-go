// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsBanditBayesianPriorInputV2 The prior the bandit's estimates start from, in the metric's units. Omitted fields keep their current values, or their defaults on create. A change takes effect on the next refresh.
type ExperimentsBanditBayesianPriorInputV2 struct {
	// Prior standard deviation of each arm's mean around the shared mean. Must be between 1e-12 and 1e12. Defaults to 100.
	ArmStandardDeviation *float64 `json:"arm_standard_deviation,omitempty"`
	// Prior mean of the shared mean of the arms. Its absolute value must be at most 1e12 times `grand_mean_standard_deviation`. Defaults to 0.
	GrandMean *float64 `json:"grand_mean,omitempty"`
	// Prior standard deviation of the shared mean of the arms. Must be between 1e-12 and 1e12, and at most 1e7 times `arm_standard_deviation`. Defaults to 10000.
	GrandMeanStandardDeviation *float64 `json:"grand_mean_standard_deviation,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsBanditBayesianPriorInputV2 instantiates a new ExperimentsBanditBayesianPriorInputV2 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsBanditBayesianPriorInputV2() *ExperimentsBanditBayesianPriorInputV2 {
	this := ExperimentsBanditBayesianPriorInputV2{}
	return &this
}

// NewExperimentsBanditBayesianPriorInputV2WithDefaults instantiates a new ExperimentsBanditBayesianPriorInputV2 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsBanditBayesianPriorInputV2WithDefaults() *ExperimentsBanditBayesianPriorInputV2 {
	this := ExperimentsBanditBayesianPriorInputV2{}
	return &this
}

// GetArmStandardDeviation returns the ArmStandardDeviation field value if set, zero value otherwise.
func (o *ExperimentsBanditBayesianPriorInputV2) GetArmStandardDeviation() float64 {
	if o == nil || o.ArmStandardDeviation == nil {
		var ret float64
		return ret
	}
	return *o.ArmStandardDeviation
}

// GetArmStandardDeviationOk returns a tuple with the ArmStandardDeviation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditBayesianPriorInputV2) GetArmStandardDeviationOk() (*float64, bool) {
	if o == nil || o.ArmStandardDeviation == nil {
		return nil, false
	}
	return o.ArmStandardDeviation, true
}

// HasArmStandardDeviation returns a boolean if a field has been set.
func (o *ExperimentsBanditBayesianPriorInputV2) HasArmStandardDeviation() bool {
	return o != nil && o.ArmStandardDeviation != nil
}

// SetArmStandardDeviation gets a reference to the given float64 and assigns it to the ArmStandardDeviation field.
func (o *ExperimentsBanditBayesianPriorInputV2) SetArmStandardDeviation(v float64) {
	o.ArmStandardDeviation = &v
}

// GetGrandMean returns the GrandMean field value if set, zero value otherwise.
func (o *ExperimentsBanditBayesianPriorInputV2) GetGrandMean() float64 {
	if o == nil || o.GrandMean == nil {
		var ret float64
		return ret
	}
	return *o.GrandMean
}

// GetGrandMeanOk returns a tuple with the GrandMean field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditBayesianPriorInputV2) GetGrandMeanOk() (*float64, bool) {
	if o == nil || o.GrandMean == nil {
		return nil, false
	}
	return o.GrandMean, true
}

// HasGrandMean returns a boolean if a field has been set.
func (o *ExperimentsBanditBayesianPriorInputV2) HasGrandMean() bool {
	return o != nil && o.GrandMean != nil
}

// SetGrandMean gets a reference to the given float64 and assigns it to the GrandMean field.
func (o *ExperimentsBanditBayesianPriorInputV2) SetGrandMean(v float64) {
	o.GrandMean = &v
}

// GetGrandMeanStandardDeviation returns the GrandMeanStandardDeviation field value if set, zero value otherwise.
func (o *ExperimentsBanditBayesianPriorInputV2) GetGrandMeanStandardDeviation() float64 {
	if o == nil || o.GrandMeanStandardDeviation == nil {
		var ret float64
		return ret
	}
	return *o.GrandMeanStandardDeviation
}

// GetGrandMeanStandardDeviationOk returns a tuple with the GrandMeanStandardDeviation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditBayesianPriorInputV2) GetGrandMeanStandardDeviationOk() (*float64, bool) {
	if o == nil || o.GrandMeanStandardDeviation == nil {
		return nil, false
	}
	return o.GrandMeanStandardDeviation, true
}

// HasGrandMeanStandardDeviation returns a boolean if a field has been set.
func (o *ExperimentsBanditBayesianPriorInputV2) HasGrandMeanStandardDeviation() bool {
	return o != nil && o.GrandMeanStandardDeviation != nil
}

// SetGrandMeanStandardDeviation gets a reference to the given float64 and assigns it to the GrandMeanStandardDeviation field.
func (o *ExperimentsBanditBayesianPriorInputV2) SetGrandMeanStandardDeviation(v float64) {
	o.GrandMeanStandardDeviation = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsBanditBayesianPriorInputV2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ArmStandardDeviation != nil {
		toSerialize["arm_standard_deviation"] = o.ArmStandardDeviation
	}
	if o.GrandMean != nil {
		toSerialize["grand_mean"] = o.GrandMean
	}
	if o.GrandMeanStandardDeviation != nil {
		toSerialize["grand_mean_standard_deviation"] = o.GrandMeanStandardDeviation
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsBanditBayesianPriorInputV2) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ArmStandardDeviation       *float64 `json:"arm_standard_deviation,omitempty"`
		GrandMean                  *float64 `json:"grand_mean,omitempty"`
		GrandMeanStandardDeviation *float64 `json:"grand_mean_standard_deviation,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"arm_standard_deviation", "grand_mean", "grand_mean_standard_deviation"})
	} else {
		return err
	}
	o.ArmStandardDeviation = all.ArmStandardDeviation
	o.GrandMean = all.GrandMean
	o.GrandMeanStandardDeviation = all.GrandMeanStandardDeviation

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
