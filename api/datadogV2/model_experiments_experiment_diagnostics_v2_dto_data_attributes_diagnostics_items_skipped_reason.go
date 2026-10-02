// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsSkippedReason Reason the diagnostic check could not be evaluated.
type ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsSkippedReason string

// List of ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsSkippedReason.
const (
	EXPERIMENTSEXPERIMENTDIAGNOSTICSV2DTODATAATTRIBUTESDIAGNOSTICSITEMSSKIPPEDREASON_NO_ASSIGNMENTS      ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsSkippedReason = "NO_ASSIGNMENTS"
	EXPERIMENTSEXPERIMENTDIAGNOSTICSV2DTODATAATTRIBUTESDIAGNOSTICSITEMSSKIPPEDREASON_NO_DIMENSIONAL_DATA ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsSkippedReason = "NO_DIMENSIONAL_DATA"
	EXPERIMENTSEXPERIMENTDIAGNOSTICSV2DTODATAATTRIBUTESDIAGNOSTICSITEMSSKIPPEDREASON_NO_METRIC_DATA      ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsSkippedReason = "NO_METRIC_DATA"
	EXPERIMENTSEXPERIMENTDIAGNOSTICSV2DTODATAATTRIBUTESDIAGNOSTICSITEMSSKIPPEDREASON_ZERO_VARIANCE       ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsSkippedReason = "ZERO_VARIANCE"
)

var allowedExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsSkippedReasonEnumValues = []ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsSkippedReason{
	EXPERIMENTSEXPERIMENTDIAGNOSTICSV2DTODATAATTRIBUTESDIAGNOSTICSITEMSSKIPPEDREASON_NO_ASSIGNMENTS,
	EXPERIMENTSEXPERIMENTDIAGNOSTICSV2DTODATAATTRIBUTESDIAGNOSTICSITEMSSKIPPEDREASON_NO_DIMENSIONAL_DATA,
	EXPERIMENTSEXPERIMENTDIAGNOSTICSV2DTODATAATTRIBUTESDIAGNOSTICSITEMSSKIPPEDREASON_NO_METRIC_DATA,
	EXPERIMENTSEXPERIMENTDIAGNOSTICSV2DTODATAATTRIBUTESDIAGNOSTICSITEMSSKIPPEDREASON_ZERO_VARIANCE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsSkippedReason) GetAllowedValues() []ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsSkippedReason {
	return allowedExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsSkippedReasonEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsSkippedReason) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsSkippedReason(value)
	return nil
}

// NewExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsSkippedReasonFromValue returns a pointer to a valid ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsSkippedReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsSkippedReasonFromValue(v string) (*ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsSkippedReason, error) {
	ev := ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsSkippedReason(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsSkippedReason: valid values are %v", v, allowedExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsSkippedReasonEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsSkippedReason) IsValid() bool {
	for _, existing := range allowedExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsSkippedReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsSkippedReason value.
func (v ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsSkippedReason) Ptr() *ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsSkippedReason {
	return &v
}

// NullableExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsSkippedReason handles when a null is used for ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsSkippedReason.
type NullableExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsSkippedReason struct {
	value *ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsSkippedReason
	isSet bool
}

// Get returns the associated value.
func (v NullableExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsSkippedReason) Get() *ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsSkippedReason {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsSkippedReason) Set(val *ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsSkippedReason) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsSkippedReason) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsSkippedReason) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsSkippedReason initializes the struct as if Set has been called.
func NewNullableExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsSkippedReason(val *ExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsSkippedReason) *NullableExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsSkippedReason {
	return &NullableExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsSkippedReason{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsSkippedReason) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableExperimentsExperimentDiagnosticsV2DTODataAttributesDiagnosticsItemsSkippedReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return datadog.Unmarshal(src, &v.value)
}
