// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GovernanceControlDetectionUpdateState The new state to set for the detection. Set to `exception` to acknowledge the detection and exclude it from active counts, or `active` to reopen it.
type GovernanceControlDetectionUpdateState string

// List of GovernanceControlDetectionUpdateState.
const (
	GOVERNANCECONTROLDETECTIONUPDATESTATE_EXCEPTION GovernanceControlDetectionUpdateState = "exception"
	GOVERNANCECONTROLDETECTIONUPDATESTATE_ACTIVE    GovernanceControlDetectionUpdateState = "active"
)

var allowedGovernanceControlDetectionUpdateStateEnumValues = []GovernanceControlDetectionUpdateState{
	GOVERNANCECONTROLDETECTIONUPDATESTATE_EXCEPTION,
	GOVERNANCECONTROLDETECTIONUPDATESTATE_ACTIVE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GovernanceControlDetectionUpdateState) GetAllowedValues() []GovernanceControlDetectionUpdateState {
	return allowedGovernanceControlDetectionUpdateStateEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GovernanceControlDetectionUpdateState) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GovernanceControlDetectionUpdateState(value)
	return nil
}

// NewGovernanceControlDetectionUpdateStateFromValue returns a pointer to a valid GovernanceControlDetectionUpdateState
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGovernanceControlDetectionUpdateStateFromValue(v string) (*GovernanceControlDetectionUpdateState, error) {
	ev := GovernanceControlDetectionUpdateState(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GovernanceControlDetectionUpdateState: valid values are %v", v, allowedGovernanceControlDetectionUpdateStateEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GovernanceControlDetectionUpdateState) IsValid() bool {
	for _, existing := range allowedGovernanceControlDetectionUpdateStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GovernanceControlDetectionUpdateState value.
func (v GovernanceControlDetectionUpdateState) Ptr() *GovernanceControlDetectionUpdateState {
	return &v
}
