// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsExperimentDiagnosticsV2DTODataAttributesState Current state of the diagnostic evaluation.
type ExperimentsExperimentDiagnosticsV2DTODataAttributesState string

// List of ExperimentsExperimentDiagnosticsV2DTODataAttributesState.
const (
	EXPERIMENTSEXPERIMENTDIAGNOSTICSV2DTODATAATTRIBUTESSTATE_NOT_STARTED ExperimentsExperimentDiagnosticsV2DTODataAttributesState = "NOT_STARTED"
	EXPERIMENTSEXPERIMENTDIAGNOSTICSV2DTODATAATTRIBUTESSTATE_RUNNING     ExperimentsExperimentDiagnosticsV2DTODataAttributesState = "RUNNING"
	EXPERIMENTSEXPERIMENTDIAGNOSTICSV2DTODATAATTRIBUTESSTATE_COMPLETED   ExperimentsExperimentDiagnosticsV2DTODataAttributesState = "COMPLETED"
	EXPERIMENTSEXPERIMENTDIAGNOSTICSV2DTODATAATTRIBUTESSTATE_FAILED      ExperimentsExperimentDiagnosticsV2DTODataAttributesState = "FAILED"
)

var allowedExperimentsExperimentDiagnosticsV2DTODataAttributesStateEnumValues = []ExperimentsExperimentDiagnosticsV2DTODataAttributesState{
	EXPERIMENTSEXPERIMENTDIAGNOSTICSV2DTODATAATTRIBUTESSTATE_NOT_STARTED,
	EXPERIMENTSEXPERIMENTDIAGNOSTICSV2DTODATAATTRIBUTESSTATE_RUNNING,
	EXPERIMENTSEXPERIMENTDIAGNOSTICSV2DTODATAATTRIBUTESSTATE_COMPLETED,
	EXPERIMENTSEXPERIMENTDIAGNOSTICSV2DTODATAATTRIBUTESSTATE_FAILED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ExperimentsExperimentDiagnosticsV2DTODataAttributesState) GetAllowedValues() []ExperimentsExperimentDiagnosticsV2DTODataAttributesState {
	return allowedExperimentsExperimentDiagnosticsV2DTODataAttributesStateEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ExperimentsExperimentDiagnosticsV2DTODataAttributesState) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ExperimentsExperimentDiagnosticsV2DTODataAttributesState(value)
	return nil
}

// NewExperimentsExperimentDiagnosticsV2DTODataAttributesStateFromValue returns a pointer to a valid ExperimentsExperimentDiagnosticsV2DTODataAttributesState
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewExperimentsExperimentDiagnosticsV2DTODataAttributesStateFromValue(v string) (*ExperimentsExperimentDiagnosticsV2DTODataAttributesState, error) {
	ev := ExperimentsExperimentDiagnosticsV2DTODataAttributesState(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ExperimentsExperimentDiagnosticsV2DTODataAttributesState: valid values are %v", v, allowedExperimentsExperimentDiagnosticsV2DTODataAttributesStateEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ExperimentsExperimentDiagnosticsV2DTODataAttributesState) IsValid() bool {
	for _, existing := range allowedExperimentsExperimentDiagnosticsV2DTODataAttributesStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExperimentsExperimentDiagnosticsV2DTODataAttributesState value.
func (v ExperimentsExperimentDiagnosticsV2DTODataAttributesState) Ptr() *ExperimentsExperimentDiagnosticsV2DTODataAttributesState {
	return &v
}
