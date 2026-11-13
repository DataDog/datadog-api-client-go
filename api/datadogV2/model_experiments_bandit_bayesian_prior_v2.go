// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsBanditBayesianPriorV2 The prior the bandit's estimates start from.
type ExperimentsBanditBayesianPriorV2 struct {
	// Prior standard deviation of each arm's mean around the shared mean.
	ArmStandardDeviation float64 `json:"arm_standard_deviation"`
	// Prior mean of the shared mean of the arms, in the metric's units.
	GrandMean float64 `json:"grand_mean"`
	// Prior standard deviation of the shared mean of the arms.
	GrandMeanStandardDeviation float64 `json:"grand_mean_standard_deviation"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsBanditBayesianPriorV2 instantiates a new ExperimentsBanditBayesianPriorV2 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsBanditBayesianPriorV2(armStandardDeviation float64, grandMean float64, grandMeanStandardDeviation float64) *ExperimentsBanditBayesianPriorV2 {
	this := ExperimentsBanditBayesianPriorV2{}
	this.ArmStandardDeviation = armStandardDeviation
	this.GrandMean = grandMean
	this.GrandMeanStandardDeviation = grandMeanStandardDeviation
	return &this
}

// NewExperimentsBanditBayesianPriorV2WithDefaults instantiates a new ExperimentsBanditBayesianPriorV2 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsBanditBayesianPriorV2WithDefaults() *ExperimentsBanditBayesianPriorV2 {
	this := ExperimentsBanditBayesianPriorV2{}
	return &this
}

// GetArmStandardDeviation returns the ArmStandardDeviation field value.
func (o *ExperimentsBanditBayesianPriorV2) GetArmStandardDeviation() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.ArmStandardDeviation
}

// GetArmStandardDeviationOk returns a tuple with the ArmStandardDeviation field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditBayesianPriorV2) GetArmStandardDeviationOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ArmStandardDeviation, true
}

// SetArmStandardDeviation sets field value.
func (o *ExperimentsBanditBayesianPriorV2) SetArmStandardDeviation(v float64) {
	o.ArmStandardDeviation = v
}

// GetGrandMean returns the GrandMean field value.
func (o *ExperimentsBanditBayesianPriorV2) GetGrandMean() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.GrandMean
}

// GetGrandMeanOk returns a tuple with the GrandMean field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditBayesianPriorV2) GetGrandMeanOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GrandMean, true
}

// SetGrandMean sets field value.
func (o *ExperimentsBanditBayesianPriorV2) SetGrandMean(v float64) {
	o.GrandMean = v
}

// GetGrandMeanStandardDeviation returns the GrandMeanStandardDeviation field value.
func (o *ExperimentsBanditBayesianPriorV2) GetGrandMeanStandardDeviation() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.GrandMeanStandardDeviation
}

// GetGrandMeanStandardDeviationOk returns a tuple with the GrandMeanStandardDeviation field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditBayesianPriorV2) GetGrandMeanStandardDeviationOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GrandMeanStandardDeviation, true
}

// SetGrandMeanStandardDeviation sets field value.
func (o *ExperimentsBanditBayesianPriorV2) SetGrandMeanStandardDeviation(v float64) {
	o.GrandMeanStandardDeviation = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsBanditBayesianPriorV2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["arm_standard_deviation"] = o.ArmStandardDeviation
	toSerialize["grand_mean"] = o.GrandMean
	toSerialize["grand_mean_standard_deviation"] = o.GrandMeanStandardDeviation

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsBanditBayesianPriorV2) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ArmStandardDeviation       *float64 `json:"arm_standard_deviation"`
		GrandMean                  *float64 `json:"grand_mean"`
		GrandMeanStandardDeviation *float64 `json:"grand_mean_standard_deviation"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ArmStandardDeviation == nil {
		return fmt.Errorf("required field arm_standard_deviation missing")
	}
	if all.GrandMean == nil {
		return fmt.Errorf("required field grand_mean missing")
	}
	if all.GrandMeanStandardDeviation == nil {
		return fmt.Errorf("required field grand_mean_standard_deviation missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"arm_standard_deviation", "grand_mean", "grand_mean_standard_deviation"})
	} else {
		return err
	}
	o.ArmStandardDeviation = *all.ArmStandardDeviation
	o.GrandMean = *all.GrandMean
	o.GrandMeanStandardDeviation = *all.GrandMeanStandardDeviation

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
