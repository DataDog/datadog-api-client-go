// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RemediationDeploymentRolloutState ECS deployment state, populated only for deployment failures.
type RemediationDeploymentRolloutState string

// List of RemediationDeploymentRolloutState.
const (
	REMEDIATIONDEPLOYMENTROLLOUTSTATE_IN_PROGRESS RemediationDeploymentRolloutState = "IN_PROGRESS"
	REMEDIATIONDEPLOYMENTROLLOUTSTATE_COMPLETED   RemediationDeploymentRolloutState = "COMPLETED"
	REMEDIATIONDEPLOYMENTROLLOUTSTATE_FAILED      RemediationDeploymentRolloutState = "FAILED"
)

var allowedRemediationDeploymentRolloutStateEnumValues = []RemediationDeploymentRolloutState{
	REMEDIATIONDEPLOYMENTROLLOUTSTATE_IN_PROGRESS,
	REMEDIATIONDEPLOYMENTROLLOUTSTATE_COMPLETED,
	REMEDIATIONDEPLOYMENTROLLOUTSTATE_FAILED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RemediationDeploymentRolloutState) GetAllowedValues() []RemediationDeploymentRolloutState {
	return allowedRemediationDeploymentRolloutStateEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RemediationDeploymentRolloutState) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RemediationDeploymentRolloutState(value)
	return nil
}

// NewRemediationDeploymentRolloutStateFromValue returns a pointer to a valid RemediationDeploymentRolloutState
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRemediationDeploymentRolloutStateFromValue(v string) (*RemediationDeploymentRolloutState, error) {
	ev := RemediationDeploymentRolloutState(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RemediationDeploymentRolloutState: valid values are %v", v, allowedRemediationDeploymentRolloutStateEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RemediationDeploymentRolloutState) IsValid() bool {
	for _, existing := range allowedRemediationDeploymentRolloutStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RemediationDeploymentRolloutState value.
func (v RemediationDeploymentRolloutState) Ptr() *RemediationDeploymentRolloutState {
	return &v
}
