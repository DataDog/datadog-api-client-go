// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsBanditDesiredChange The direction the bandit optimizes the metric in. It's fixed when the bandit starts. Null before the first saved refresh.
type ExperimentsBanditDesiredChange string

// List of ExperimentsBanditDesiredChange.
const (
	EXPERIMENTSBANDITDESIREDCHANGE_METRIC_INCREASES ExperimentsBanditDesiredChange = "METRIC_INCREASES"
	EXPERIMENTSBANDITDESIREDCHANGE_METRIC_DECREASES ExperimentsBanditDesiredChange = "METRIC_DECREASES"
)

var allowedExperimentsBanditDesiredChangeEnumValues = []ExperimentsBanditDesiredChange{
	EXPERIMENTSBANDITDESIREDCHANGE_METRIC_INCREASES,
	EXPERIMENTSBANDITDESIREDCHANGE_METRIC_DECREASES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ExperimentsBanditDesiredChange) GetAllowedValues() []ExperimentsBanditDesiredChange {
	return allowedExperimentsBanditDesiredChangeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ExperimentsBanditDesiredChange) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ExperimentsBanditDesiredChange(value)
	return nil
}

// NewExperimentsBanditDesiredChangeFromValue returns a pointer to a valid ExperimentsBanditDesiredChange
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewExperimentsBanditDesiredChangeFromValue(v string) (*ExperimentsBanditDesiredChange, error) {
	ev := ExperimentsBanditDesiredChange(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ExperimentsBanditDesiredChange: valid values are %v", v, allowedExperimentsBanditDesiredChangeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ExperimentsBanditDesiredChange) IsValid() bool {
	for _, existing := range allowedExperimentsBanditDesiredChangeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExperimentsBanditDesiredChange value.
func (v ExperimentsBanditDesiredChange) Ptr() *ExperimentsBanditDesiredChange {
	return &v
}

// NullableExperimentsBanditDesiredChange handles when a null is used for ExperimentsBanditDesiredChange.
type NullableExperimentsBanditDesiredChange struct {
	value *ExperimentsBanditDesiredChange
	isSet bool
}

// Get returns the associated value.
func (v NullableExperimentsBanditDesiredChange) Get() *ExperimentsBanditDesiredChange {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableExperimentsBanditDesiredChange) Set(val *ExperimentsBanditDesiredChange) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableExperimentsBanditDesiredChange) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableExperimentsBanditDesiredChange) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableExperimentsBanditDesiredChange initializes the struct as if Set has been called.
func NewNullableExperimentsBanditDesiredChange(val *ExperimentsBanditDesiredChange) *NullableExperimentsBanditDesiredChange {
	return &NullableExperimentsBanditDesiredChange{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableExperimentsBanditDesiredChange) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableExperimentsBanditDesiredChange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return datadog.Unmarshal(src, &v.value)
}
