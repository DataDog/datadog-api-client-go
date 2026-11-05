// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityAgentInvestigationState State of the investigation: `investigating` while it runs, `completed` after it has a result, or `failed`
// when it produced no result before its timeout.
type SecurityAgentInvestigationState string

// List of SecurityAgentInvestigationState.
const (
	SECURITYAGENTINVESTIGATIONSTATE_INVESTIGATING SecurityAgentInvestigationState = "investigating"
	SECURITYAGENTINVESTIGATIONSTATE_COMPLETED     SecurityAgentInvestigationState = "completed"
	SECURITYAGENTINVESTIGATIONSTATE_FAILED        SecurityAgentInvestigationState = "failed"
)

var allowedSecurityAgentInvestigationStateEnumValues = []SecurityAgentInvestigationState{
	SECURITYAGENTINVESTIGATIONSTATE_INVESTIGATING,
	SECURITYAGENTINVESTIGATIONSTATE_COMPLETED,
	SECURITYAGENTINVESTIGATIONSTATE_FAILED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityAgentInvestigationState) GetAllowedValues() []SecurityAgentInvestigationState {
	return allowedSecurityAgentInvestigationStateEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityAgentInvestigationState) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityAgentInvestigationState(value)
	return nil
}

// NewSecurityAgentInvestigationStateFromValue returns a pointer to a valid SecurityAgentInvestigationState
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityAgentInvestigationStateFromValue(v string) (*SecurityAgentInvestigationState, error) {
	ev := SecurityAgentInvestigationState(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityAgentInvestigationState: valid values are %v", v, allowedSecurityAgentInvestigationStateEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityAgentInvestigationState) IsValid() bool {
	for _, existing := range allowedSecurityAgentInvestigationStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityAgentInvestigationState value.
func (v SecurityAgentInvestigationState) Ptr() *SecurityAgentInvestigationState {
	return &v
}
