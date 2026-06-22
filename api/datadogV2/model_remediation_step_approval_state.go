// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RemediationStepApprovalState Approval state for a remediation step.
type RemediationStepApprovalState string

// List of RemediationStepApprovalState.
const (
	REMEDIATIONSTEPAPPROVALSTATE_PENDING_APPROVAL RemediationStepApprovalState = "pending_approval"
	REMEDIATIONSTEPAPPROVALSTATE_APPROVED         RemediationStepApprovalState = "approved"
	REMEDIATIONSTEPAPPROVALSTATE_REJECTED         RemediationStepApprovalState = "rejected"
)

var allowedRemediationStepApprovalStateEnumValues = []RemediationStepApprovalState{
	REMEDIATIONSTEPAPPROVALSTATE_PENDING_APPROVAL,
	REMEDIATIONSTEPAPPROVALSTATE_APPROVED,
	REMEDIATIONSTEPAPPROVALSTATE_REJECTED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RemediationStepApprovalState) GetAllowedValues() []RemediationStepApprovalState {
	return allowedRemediationStepApprovalStateEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RemediationStepApprovalState) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RemediationStepApprovalState(value)
	return nil
}

// NewRemediationStepApprovalStateFromValue returns a pointer to a valid RemediationStepApprovalState
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRemediationStepApprovalStateFromValue(v string) (*RemediationStepApprovalState, error) {
	ev := RemediationStepApprovalState(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RemediationStepApprovalState: valid values are %v", v, allowedRemediationStepApprovalStateEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RemediationStepApprovalState) IsValid() bool {
	for _, existing := range allowedRemediationStepApprovalStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RemediationStepApprovalState value.
func (v RemediationStepApprovalState) Ptr() *RemediationStepApprovalState {
	return &v
}
