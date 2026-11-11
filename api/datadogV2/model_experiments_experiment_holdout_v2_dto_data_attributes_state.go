// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsExperimentHoldoutV2DTODataAttributesState Lifecycle phase of the holdout. PENDING holdouts have not started including experiments. INCLUDING holdouts are adding experiments. WAITING_FOR_ANALYSIS holdouts have stopped including experiments but have not started analysis. ANALYSING holdouts are in their analysis window. COMPLETE applies only to EXTERNAL holdouts that have ended their analysis window but are not archived.
type ExperimentsExperimentHoldoutV2DTODataAttributesState string

// List of ExperimentsExperimentHoldoutV2DTODataAttributesState.
const (
	EXPERIMENTSEXPERIMENTHOLDOUTV2DTODATAATTRIBUTESSTATE_PENDING              ExperimentsExperimentHoldoutV2DTODataAttributesState = "PENDING"
	EXPERIMENTSEXPERIMENTHOLDOUTV2DTODATAATTRIBUTESSTATE_INCLUDING            ExperimentsExperimentHoldoutV2DTODataAttributesState = "INCLUDING"
	EXPERIMENTSEXPERIMENTHOLDOUTV2DTODATAATTRIBUTESSTATE_WAITING_FOR_ANALYSIS ExperimentsExperimentHoldoutV2DTODataAttributesState = "WAITING_FOR_ANALYSIS"
	EXPERIMENTSEXPERIMENTHOLDOUTV2DTODATAATTRIBUTESSTATE_ANALYSING            ExperimentsExperimentHoldoutV2DTODataAttributesState = "ANALYSING"
	EXPERIMENTSEXPERIMENTHOLDOUTV2DTODATAATTRIBUTESSTATE_COMPLETE             ExperimentsExperimentHoldoutV2DTODataAttributesState = "COMPLETE"
	EXPERIMENTSEXPERIMENTHOLDOUTV2DTODATAATTRIBUTESSTATE_UNKNOWN              ExperimentsExperimentHoldoutV2DTODataAttributesState = "UNKNOWN"
)

var allowedExperimentsExperimentHoldoutV2DTODataAttributesStateEnumValues = []ExperimentsExperimentHoldoutV2DTODataAttributesState{
	EXPERIMENTSEXPERIMENTHOLDOUTV2DTODATAATTRIBUTESSTATE_PENDING,
	EXPERIMENTSEXPERIMENTHOLDOUTV2DTODATAATTRIBUTESSTATE_INCLUDING,
	EXPERIMENTSEXPERIMENTHOLDOUTV2DTODATAATTRIBUTESSTATE_WAITING_FOR_ANALYSIS,
	EXPERIMENTSEXPERIMENTHOLDOUTV2DTODATAATTRIBUTESSTATE_ANALYSING,
	EXPERIMENTSEXPERIMENTHOLDOUTV2DTODATAATTRIBUTESSTATE_COMPLETE,
	EXPERIMENTSEXPERIMENTHOLDOUTV2DTODATAATTRIBUTESSTATE_UNKNOWN,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ExperimentsExperimentHoldoutV2DTODataAttributesState) GetAllowedValues() []ExperimentsExperimentHoldoutV2DTODataAttributesState {
	return allowedExperimentsExperimentHoldoutV2DTODataAttributesStateEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ExperimentsExperimentHoldoutV2DTODataAttributesState) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ExperimentsExperimentHoldoutV2DTODataAttributesState(value)
	return nil
}

// NewExperimentsExperimentHoldoutV2DTODataAttributesStateFromValue returns a pointer to a valid ExperimentsExperimentHoldoutV2DTODataAttributesState
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewExperimentsExperimentHoldoutV2DTODataAttributesStateFromValue(v string) (*ExperimentsExperimentHoldoutV2DTODataAttributesState, error) {
	ev := ExperimentsExperimentHoldoutV2DTODataAttributesState(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ExperimentsExperimentHoldoutV2DTODataAttributesState: valid values are %v", v, allowedExperimentsExperimentHoldoutV2DTODataAttributesStateEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ExperimentsExperimentHoldoutV2DTODataAttributesState) IsValid() bool {
	for _, existing := range allowedExperimentsExperimentHoldoutV2DTODataAttributesStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExperimentsExperimentHoldoutV2DTODataAttributesState value.
func (v ExperimentsExperimentHoldoutV2DTODataAttributesState) Ptr() *ExperimentsExperimentHoldoutV2DTODataAttributesState {
	return &v
}
