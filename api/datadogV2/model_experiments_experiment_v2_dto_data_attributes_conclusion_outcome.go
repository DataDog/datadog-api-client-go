// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsExperimentV2DTODataAttributesConclusionOutcome Recorded experiment outcome.
type ExperimentsExperimentV2DTODataAttributesConclusionOutcome string

// List of ExperimentsExperimentV2DTODataAttributesConclusionOutcome.
const (
	EXPERIMENTSEXPERIMENTV2DTODATAATTRIBUTESCONCLUSIONOUTCOME_POSITIVE      ExperimentsExperimentV2DTODataAttributesConclusionOutcome = "POSITIVE"
	EXPERIMENTSEXPERIMENTV2DTODATAATTRIBUTESCONCLUSIONOUTCOME_NEGATIVE      ExperimentsExperimentV2DTODataAttributesConclusionOutcome = "NEGATIVE"
	EXPERIMENTSEXPERIMENTV2DTODATAATTRIBUTESCONCLUSIONOUTCOME_NEUTRAL       ExperimentsExperimentV2DTODataAttributesConclusionOutcome = "NEUTRAL"
	EXPERIMENTSEXPERIMENTV2DTODATAATTRIBUTESCONCLUSIONOUTCOME_INCONCLUSIVE  ExperimentsExperimentV2DTODataAttributesConclusionOutcome = "INCONCLUSIVE"
	EXPERIMENTSEXPERIMENTV2DTODATAATTRIBUTESCONCLUSIONOUTCOME_MISCONFIGURED ExperimentsExperimentV2DTODataAttributesConclusionOutcome = "MISCONFIGURED"
	EXPERIMENTSEXPERIMENTV2DTODATAATTRIBUTESCONCLUSIONOUTCOME_UNKNOWN       ExperimentsExperimentV2DTODataAttributesConclusionOutcome = "UNKNOWN"
)

var allowedExperimentsExperimentV2DTODataAttributesConclusionOutcomeEnumValues = []ExperimentsExperimentV2DTODataAttributesConclusionOutcome{
	EXPERIMENTSEXPERIMENTV2DTODATAATTRIBUTESCONCLUSIONOUTCOME_POSITIVE,
	EXPERIMENTSEXPERIMENTV2DTODATAATTRIBUTESCONCLUSIONOUTCOME_NEGATIVE,
	EXPERIMENTSEXPERIMENTV2DTODATAATTRIBUTESCONCLUSIONOUTCOME_NEUTRAL,
	EXPERIMENTSEXPERIMENTV2DTODATAATTRIBUTESCONCLUSIONOUTCOME_INCONCLUSIVE,
	EXPERIMENTSEXPERIMENTV2DTODATAATTRIBUTESCONCLUSIONOUTCOME_MISCONFIGURED,
	EXPERIMENTSEXPERIMENTV2DTODATAATTRIBUTESCONCLUSIONOUTCOME_UNKNOWN,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ExperimentsExperimentV2DTODataAttributesConclusionOutcome) GetAllowedValues() []ExperimentsExperimentV2DTODataAttributesConclusionOutcome {
	return allowedExperimentsExperimentV2DTODataAttributesConclusionOutcomeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ExperimentsExperimentV2DTODataAttributesConclusionOutcome) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ExperimentsExperimentV2DTODataAttributesConclusionOutcome(value)
	return nil
}

// NewExperimentsExperimentV2DTODataAttributesConclusionOutcomeFromValue returns a pointer to a valid ExperimentsExperimentV2DTODataAttributesConclusionOutcome
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewExperimentsExperimentV2DTODataAttributesConclusionOutcomeFromValue(v string) (*ExperimentsExperimentV2DTODataAttributesConclusionOutcome, error) {
	ev := ExperimentsExperimentV2DTODataAttributesConclusionOutcome(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ExperimentsExperimentV2DTODataAttributesConclusionOutcome: valid values are %v", v, allowedExperimentsExperimentV2DTODataAttributesConclusionOutcomeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ExperimentsExperimentV2DTODataAttributesConclusionOutcome) IsValid() bool {
	for _, existing := range allowedExperimentsExperimentV2DTODataAttributesConclusionOutcomeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExperimentsExperimentV2DTODataAttributesConclusionOutcome value.
func (v ExperimentsExperimentV2DTODataAttributesConclusionOutcome) Ptr() *ExperimentsExperimentV2DTODataAttributesConclusionOutcome {
	return &v
}

// NullableExperimentsExperimentV2DTODataAttributesConclusionOutcome handles when a null is used for ExperimentsExperimentV2DTODataAttributesConclusionOutcome.
type NullableExperimentsExperimentV2DTODataAttributesConclusionOutcome struct {
	value *ExperimentsExperimentV2DTODataAttributesConclusionOutcome
	isSet bool
}

// Get returns the associated value.
func (v NullableExperimentsExperimentV2DTODataAttributesConclusionOutcome) Get() *ExperimentsExperimentV2DTODataAttributesConclusionOutcome {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableExperimentsExperimentV2DTODataAttributesConclusionOutcome) Set(val *ExperimentsExperimentV2DTODataAttributesConclusionOutcome) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableExperimentsExperimentV2DTODataAttributesConclusionOutcome) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableExperimentsExperimentV2DTODataAttributesConclusionOutcome) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableExperimentsExperimentV2DTODataAttributesConclusionOutcome initializes the struct as if Set has been called.
func NewNullableExperimentsExperimentV2DTODataAttributesConclusionOutcome(val *ExperimentsExperimentV2DTODataAttributesConclusionOutcome) *NullableExperimentsExperimentV2DTODataAttributesConclusionOutcome {
	return &NullableExperimentsExperimentV2DTODataAttributesConclusionOutcome{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableExperimentsExperimentV2DTODataAttributesConclusionOutcome) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableExperimentsExperimentV2DTODataAttributesConclusionOutcome) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return datadog.Unmarshal(src, &v.value)
}
