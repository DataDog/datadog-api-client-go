// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EnvironmentSubjectAuthorizationState Whether assignment requests for this environment require subject authorization. Allowed values are `disabled` and `enabled`. Enabling requires a provider ID formatted as a nonzero UUID. If omitted on create, defaults to `disabled`. If omitted on update, the stored state is preserved. `null` is not accepted. This setting applies only to the current Datadog site and is not copied to other sites. Before you enable it, configure an eligible Custom OIDC provider in the same organization and site, and verify that the provider's keys are published. Assignment requests fail when the required verification configuration is unavailable.
type EnvironmentSubjectAuthorizationState string

// List of EnvironmentSubjectAuthorizationState.
const (
	ENVIRONMENTSUBJECTAUTHORIZATIONSTATE_DISABLED EnvironmentSubjectAuthorizationState = "disabled"
	ENVIRONMENTSUBJECTAUTHORIZATIONSTATE_ENABLED  EnvironmentSubjectAuthorizationState = "enabled"
)

var allowedEnvironmentSubjectAuthorizationStateEnumValues = []EnvironmentSubjectAuthorizationState{
	ENVIRONMENTSUBJECTAUTHORIZATIONSTATE_DISABLED,
	ENVIRONMENTSUBJECTAUTHORIZATIONSTATE_ENABLED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EnvironmentSubjectAuthorizationState) GetAllowedValues() []EnvironmentSubjectAuthorizationState {
	return allowedEnvironmentSubjectAuthorizationStateEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EnvironmentSubjectAuthorizationState) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EnvironmentSubjectAuthorizationState(value)
	return nil
}

// NewEnvironmentSubjectAuthorizationStateFromValue returns a pointer to a valid EnvironmentSubjectAuthorizationState
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEnvironmentSubjectAuthorizationStateFromValue(v string) (*EnvironmentSubjectAuthorizationState, error) {
	ev := EnvironmentSubjectAuthorizationState(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EnvironmentSubjectAuthorizationState: valid values are %v", v, allowedEnvironmentSubjectAuthorizationStateEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EnvironmentSubjectAuthorizationState) IsValid() bool {
	for _, existing := range allowedEnvironmentSubjectAuthorizationStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnvironmentSubjectAuthorizationState value.
func (v EnvironmentSubjectAuthorizationState) Ptr() *EnvironmentSubjectAuthorizationState {
	return &v
}
