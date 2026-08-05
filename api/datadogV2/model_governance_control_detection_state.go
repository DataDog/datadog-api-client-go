// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GovernanceControlDetectionState The current state of the detection. Possible values are `active`, `exception`, `mitigated`, `inactive`, `obsolete`, `resolved_externally`, and `mitigation_in_progress`.
type GovernanceControlDetectionState string

// List of GovernanceControlDetectionState.
const (
	GOVERNANCECONTROLDETECTIONSTATE_ACTIVE                 GovernanceControlDetectionState = "active"
	GOVERNANCECONTROLDETECTIONSTATE_EXCEPTION              GovernanceControlDetectionState = "exception"
	GOVERNANCECONTROLDETECTIONSTATE_MITIGATED              GovernanceControlDetectionState = "mitigated"
	GOVERNANCECONTROLDETECTIONSTATE_INACTIVE               GovernanceControlDetectionState = "inactive"
	GOVERNANCECONTROLDETECTIONSTATE_OBSOLETE               GovernanceControlDetectionState = "obsolete"
	GOVERNANCECONTROLDETECTIONSTATE_RESOLVED_EXTERNALLY    GovernanceControlDetectionState = "resolved_externally"
	GOVERNANCECONTROLDETECTIONSTATE_MITIGATION_IN_PROGRESS GovernanceControlDetectionState = "mitigation_in_progress"
)

var allowedGovernanceControlDetectionStateEnumValues = []GovernanceControlDetectionState{
	GOVERNANCECONTROLDETECTIONSTATE_ACTIVE,
	GOVERNANCECONTROLDETECTIONSTATE_EXCEPTION,
	GOVERNANCECONTROLDETECTIONSTATE_MITIGATED,
	GOVERNANCECONTROLDETECTIONSTATE_INACTIVE,
	GOVERNANCECONTROLDETECTIONSTATE_OBSOLETE,
	GOVERNANCECONTROLDETECTIONSTATE_RESOLVED_EXTERNALLY,
	GOVERNANCECONTROLDETECTIONSTATE_MITIGATION_IN_PROGRESS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GovernanceControlDetectionState) GetAllowedValues() []GovernanceControlDetectionState {
	return allowedGovernanceControlDetectionStateEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GovernanceControlDetectionState) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GovernanceControlDetectionState(value)
	return nil
}

// NewGovernanceControlDetectionStateFromValue returns a pointer to a valid GovernanceControlDetectionState
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGovernanceControlDetectionStateFromValue(v string) (*GovernanceControlDetectionState, error) {
	ev := GovernanceControlDetectionState(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GovernanceControlDetectionState: valid values are %v", v, allowedGovernanceControlDetectionStateEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GovernanceControlDetectionState) IsValid() bool {
	for _, existing := range allowedGovernanceControlDetectionStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GovernanceControlDetectionState value.
func (v GovernanceControlDetectionState) Ptr() *GovernanceControlDetectionState {
	return &v
}
