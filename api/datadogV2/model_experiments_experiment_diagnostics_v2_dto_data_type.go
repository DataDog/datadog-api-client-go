// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsExperimentDiagnosticsV2DTODataType Experiment diagnostics resource type.
type ExperimentsExperimentDiagnosticsV2DTODataType string

// List of ExperimentsExperimentDiagnosticsV2DTODataType.
const (
	EXPERIMENTSEXPERIMENTDIAGNOSTICSV2DTODATATYPE_EXPERIMENT_DIAGNOSTICS ExperimentsExperimentDiagnosticsV2DTODataType = "experiment-diagnostics"
)

var allowedExperimentsExperimentDiagnosticsV2DTODataTypeEnumValues = []ExperimentsExperimentDiagnosticsV2DTODataType{
	EXPERIMENTSEXPERIMENTDIAGNOSTICSV2DTODATATYPE_EXPERIMENT_DIAGNOSTICS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ExperimentsExperimentDiagnosticsV2DTODataType) GetAllowedValues() []ExperimentsExperimentDiagnosticsV2DTODataType {
	return allowedExperimentsExperimentDiagnosticsV2DTODataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ExperimentsExperimentDiagnosticsV2DTODataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ExperimentsExperimentDiagnosticsV2DTODataType(value)
	return nil
}

// NewExperimentsExperimentDiagnosticsV2DTODataTypeFromValue returns a pointer to a valid ExperimentsExperimentDiagnosticsV2DTODataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewExperimentsExperimentDiagnosticsV2DTODataTypeFromValue(v string) (*ExperimentsExperimentDiagnosticsV2DTODataType, error) {
	ev := ExperimentsExperimentDiagnosticsV2DTODataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ExperimentsExperimentDiagnosticsV2DTODataType: valid values are %v", v, allowedExperimentsExperimentDiagnosticsV2DTODataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ExperimentsExperimentDiagnosticsV2DTODataType) IsValid() bool {
	for _, existing := range allowedExperimentsExperimentDiagnosticsV2DTODataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExperimentsExperimentDiagnosticsV2DTODataType value.
func (v ExperimentsExperimentDiagnosticsV2DTODataType) Ptr() *ExperimentsExperimentDiagnosticsV2DTODataType {
	return &v
}
