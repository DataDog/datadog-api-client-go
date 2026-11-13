// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsBanditNullableConfidenceIntervalV2 A central interval around point_estimate at confidence_level. Null when the estimate is unavailable.
type ExperimentsBanditNullableConfidenceIntervalV2 struct {
	// Lower bound of the interval.
	Lower float64 `json:"lower"`
	// Upper bound of the interval.
	Upper float64 `json:"upper"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsBanditNullableConfidenceIntervalV2 instantiates a new ExperimentsBanditNullableConfidenceIntervalV2 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsBanditNullableConfidenceIntervalV2(lower float64, upper float64) *ExperimentsBanditNullableConfidenceIntervalV2 {
	this := ExperimentsBanditNullableConfidenceIntervalV2{}
	this.Lower = lower
	this.Upper = upper
	return &this
}

// NewExperimentsBanditNullableConfidenceIntervalV2WithDefaults instantiates a new ExperimentsBanditNullableConfidenceIntervalV2 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsBanditNullableConfidenceIntervalV2WithDefaults() *ExperimentsBanditNullableConfidenceIntervalV2 {
	this := ExperimentsBanditNullableConfidenceIntervalV2{}
	return &this
}

// GetLower returns the Lower field value.
func (o *ExperimentsBanditNullableConfidenceIntervalV2) GetLower() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Lower
}

// GetLowerOk returns a tuple with the Lower field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditNullableConfidenceIntervalV2) GetLowerOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Lower, true
}

// SetLower sets field value.
func (o *ExperimentsBanditNullableConfidenceIntervalV2) SetLower(v float64) {
	o.Lower = v
}

// GetUpper returns the Upper field value.
func (o *ExperimentsBanditNullableConfidenceIntervalV2) GetUpper() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Upper
}

// GetUpperOk returns a tuple with the Upper field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditNullableConfidenceIntervalV2) GetUpperOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Upper, true
}

// SetUpper sets field value.
func (o *ExperimentsBanditNullableConfidenceIntervalV2) SetUpper(v float64) {
	o.Upper = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsBanditNullableConfidenceIntervalV2) MarshalJSON() ([]byte, error) {
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
func (o *ExperimentsBanditNullableConfidenceIntervalV2) UnmarshalJSON(bytes []byte) (err error) {
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

// NullableExperimentsBanditNullableConfidenceIntervalV2 handles when a null is used for ExperimentsBanditNullableConfidenceIntervalV2.
type NullableExperimentsBanditNullableConfidenceIntervalV2 struct {
	value *ExperimentsBanditNullableConfidenceIntervalV2
	isSet bool
}

// Get returns the associated value.
func (v NullableExperimentsBanditNullableConfidenceIntervalV2) Get() *ExperimentsBanditNullableConfidenceIntervalV2 {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableExperimentsBanditNullableConfidenceIntervalV2) Set(val *ExperimentsBanditNullableConfidenceIntervalV2) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableExperimentsBanditNullableConfidenceIntervalV2) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableExperimentsBanditNullableConfidenceIntervalV2) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableExperimentsBanditNullableConfidenceIntervalV2 initializes the struct as if Set has been called.
func NewNullableExperimentsBanditNullableConfidenceIntervalV2(val *ExperimentsBanditNullableConfidenceIntervalV2) *NullableExperimentsBanditNullableConfidenceIntervalV2 {
	return &NullableExperimentsBanditNullableConfidenceIntervalV2{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableExperimentsBanditNullableConfidenceIntervalV2) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableExperimentsBanditNullableConfidenceIntervalV2) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return datadog.Unmarshal(src, &v.value)
}
