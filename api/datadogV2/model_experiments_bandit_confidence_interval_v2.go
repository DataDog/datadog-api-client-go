// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsBanditConfidenceIntervalV2 A central interval around point_estimate at confidence_level.
type ExperimentsBanditConfidenceIntervalV2 struct {
	// Lower bound of the interval.
	Lower float64 `json:"lower"`
	// Upper bound of the interval.
	Upper float64 `json:"upper"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsBanditConfidenceIntervalV2 instantiates a new ExperimentsBanditConfidenceIntervalV2 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsBanditConfidenceIntervalV2(lower float64, upper float64) *ExperimentsBanditConfidenceIntervalV2 {
	this := ExperimentsBanditConfidenceIntervalV2{}
	this.Lower = lower
	this.Upper = upper
	return &this
}

// NewExperimentsBanditConfidenceIntervalV2WithDefaults instantiates a new ExperimentsBanditConfidenceIntervalV2 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsBanditConfidenceIntervalV2WithDefaults() *ExperimentsBanditConfidenceIntervalV2 {
	this := ExperimentsBanditConfidenceIntervalV2{}
	return &this
}

// GetLower returns the Lower field value.
func (o *ExperimentsBanditConfidenceIntervalV2) GetLower() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Lower
}

// GetLowerOk returns a tuple with the Lower field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditConfidenceIntervalV2) GetLowerOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Lower, true
}

// SetLower sets field value.
func (o *ExperimentsBanditConfidenceIntervalV2) SetLower(v float64) {
	o.Lower = v
}

// GetUpper returns the Upper field value.
func (o *ExperimentsBanditConfidenceIntervalV2) GetUpper() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Upper
}

// GetUpperOk returns a tuple with the Upper field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditConfidenceIntervalV2) GetUpperOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Upper, true
}

// SetUpper sets field value.
func (o *ExperimentsBanditConfidenceIntervalV2) SetUpper(v float64) {
	o.Upper = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsBanditConfidenceIntervalV2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["lower"] = o.Lower
	toSerialize["upper"] = o.Upper

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsBanditConfidenceIntervalV2) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Lower *float64 `json:"lower"`
		Upper *float64 `json:"upper"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Lower == nil {
		return fmt.Errorf("required field lower missing")
	}
	if all.Upper == nil {
		return fmt.Errorf("required field upper missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"lower", "upper"})
	} else {
		return err
	}
	o.Lower = *all.Lower
	o.Upper = *all.Upper

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
