// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EscalationPolicyActionType Indicates that the action is an escalation policy action.
type EscalationPolicyActionType string

// List of EscalationPolicyActionType.
const (
	ESCALATIONPOLICYACTIONTYPE_ESCALATION_POLICY EscalationPolicyActionType = "escalation_policy"
)

var allowedEscalationPolicyActionTypeEnumValues = []EscalationPolicyActionType{
	ESCALATIONPOLICYACTIONTYPE_ESCALATION_POLICY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EscalationPolicyActionType) GetAllowedValues() []EscalationPolicyActionType {
	return allowedEscalationPolicyActionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EscalationPolicyActionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EscalationPolicyActionType(value)
	return nil
}

// NewEscalationPolicyActionTypeFromValue returns a pointer to a valid EscalationPolicyActionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEscalationPolicyActionTypeFromValue(v string) (*EscalationPolicyActionType, error) {
	ev := EscalationPolicyActionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EscalationPolicyActionType: valid values are %v", v, allowedEscalationPolicyActionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EscalationPolicyActionType) IsValid() bool {
	for _, existing := range allowedEscalationPolicyActionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EscalationPolicyActionType value.
func (v EscalationPolicyActionType) Ptr() *EscalationPolicyActionType {
	return &v
}
