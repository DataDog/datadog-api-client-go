// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsExperimentDiagnosticsV2DTODataAttributesResult Overall result of the experiment diagnostic checks.
type ExperimentsExperimentDiagnosticsV2DTODataAttributesResult string

// List of ExperimentsExperimentDiagnosticsV2DTODataAttributesResult.
const (
	EXPERIMENTSEXPERIMENTDIAGNOSTICSV2DTODATAATTRIBUTESRESULT_PASS    ExperimentsExperimentDiagnosticsV2DTODataAttributesResult = "PASS"
	EXPERIMENTSEXPERIMENTDIAGNOSTICSV2DTODATAATTRIBUTESRESULT_FAIL    ExperimentsExperimentDiagnosticsV2DTODataAttributesResult = "FAIL"
	EXPERIMENTSEXPERIMENTDIAGNOSTICSV2DTODATAATTRIBUTESRESULT_WARN    ExperimentsExperimentDiagnosticsV2DTODataAttributesResult = "WARN"
	EXPERIMENTSEXPERIMENTDIAGNOSTICSV2DTODATAATTRIBUTESRESULT_NO_DATA ExperimentsExperimentDiagnosticsV2DTODataAttributesResult = "NO_DATA"
)

var allowedExperimentsExperimentDiagnosticsV2DTODataAttributesResultEnumValues = []ExperimentsExperimentDiagnosticsV2DTODataAttributesResult{
	EXPERIMENTSEXPERIMENTDIAGNOSTICSV2DTODATAATTRIBUTESRESULT_PASS,
	EXPERIMENTSEXPERIMENTDIAGNOSTICSV2DTODATAATTRIBUTESRESULT_FAIL,
	EXPERIMENTSEXPERIMENTDIAGNOSTICSV2DTODATAATTRIBUTESRESULT_WARN,
	EXPERIMENTSEXPERIMENTDIAGNOSTICSV2DTODATAATTRIBUTESRESULT_NO_DATA,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ExperimentsExperimentDiagnosticsV2DTODataAttributesResult) GetAllowedValues() []ExperimentsExperimentDiagnosticsV2DTODataAttributesResult {
	return allowedExperimentsExperimentDiagnosticsV2DTODataAttributesResultEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ExperimentsExperimentDiagnosticsV2DTODataAttributesResult) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ExperimentsExperimentDiagnosticsV2DTODataAttributesResult(value)
	return nil
}

// NewExperimentsExperimentDiagnosticsV2DTODataAttributesResultFromValue returns a pointer to a valid ExperimentsExperimentDiagnosticsV2DTODataAttributesResult
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewExperimentsExperimentDiagnosticsV2DTODataAttributesResultFromValue(v string) (*ExperimentsExperimentDiagnosticsV2DTODataAttributesResult, error) {
	ev := ExperimentsExperimentDiagnosticsV2DTODataAttributesResult(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ExperimentsExperimentDiagnosticsV2DTODataAttributesResult: valid values are %v", v, allowedExperimentsExperimentDiagnosticsV2DTODataAttributesResultEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ExperimentsExperimentDiagnosticsV2DTODataAttributesResult) IsValid() bool {
	for _, existing := range allowedExperimentsExperimentDiagnosticsV2DTODataAttributesResultEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExperimentsExperimentDiagnosticsV2DTODataAttributesResult value.
func (v ExperimentsExperimentDiagnosticsV2DTODataAttributesResult) Ptr() *ExperimentsExperimentDiagnosticsV2DTODataAttributesResult {
	return &v
}

// NullableExperimentsExperimentDiagnosticsV2DTODataAttributesResult handles when a null is used for ExperimentsExperimentDiagnosticsV2DTODataAttributesResult.
type NullableExperimentsExperimentDiagnosticsV2DTODataAttributesResult struct {
	value *ExperimentsExperimentDiagnosticsV2DTODataAttributesResult
	isSet bool
}

// Get returns the associated value.
func (v NullableExperimentsExperimentDiagnosticsV2DTODataAttributesResult) Get() *ExperimentsExperimentDiagnosticsV2DTODataAttributesResult {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableExperimentsExperimentDiagnosticsV2DTODataAttributesResult) Set(val *ExperimentsExperimentDiagnosticsV2DTODataAttributesResult) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableExperimentsExperimentDiagnosticsV2DTODataAttributesResult) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableExperimentsExperimentDiagnosticsV2DTODataAttributesResult) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableExperimentsExperimentDiagnosticsV2DTODataAttributesResult initializes the struct as if Set has been called.
func NewNullableExperimentsExperimentDiagnosticsV2DTODataAttributesResult(val *ExperimentsExperimentDiagnosticsV2DTODataAttributesResult) *NullableExperimentsExperimentDiagnosticsV2DTODataAttributesResult {
	return &NullableExperimentsExperimentDiagnosticsV2DTODataAttributesResult{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableExperimentsExperimentDiagnosticsV2DTODataAttributesResult) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableExperimentsExperimentDiagnosticsV2DTODataAttributesResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return datadog.Unmarshal(src, &v.value)
}
