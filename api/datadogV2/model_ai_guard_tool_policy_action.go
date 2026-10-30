// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AIGuardToolPolicyAction The policy action for a tool named by the property. Tool policies override
// the recommendation after evaluation; they do not skip evaluation.
// Any BLOCK takes precedence over PASS and EVALUATE. The Supply Chain Firewall
// is applied afterward and can still return DENY.
type AIGuardToolPolicyAction string

// List of AIGuardToolPolicyAction.
const (
	AIGUARDTOOLPOLICYACTION_PASS     AIGuardToolPolicyAction = "PASS"
	AIGUARDTOOLPOLICYACTION_BLOCK    AIGuardToolPolicyAction = "BLOCK"
	AIGUARDTOOLPOLICYACTION_EVALUATE AIGuardToolPolicyAction = "EVALUATE"
)

var allowedAIGuardToolPolicyActionEnumValues = []AIGuardToolPolicyAction{
	AIGUARDTOOLPOLICYACTION_PASS,
	AIGUARDTOOLPOLICYACTION_BLOCK,
	AIGUARDTOOLPOLICYACTION_EVALUATE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AIGuardToolPolicyAction) GetAllowedValues() []AIGuardToolPolicyAction {
	return allowedAIGuardToolPolicyActionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AIGuardToolPolicyAction) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AIGuardToolPolicyAction(value)
	return nil
}

// NewAIGuardToolPolicyActionFromValue returns a pointer to a valid AIGuardToolPolicyAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAIGuardToolPolicyActionFromValue(v string) (*AIGuardToolPolicyAction, error) {
	ev := AIGuardToolPolicyAction(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AIGuardToolPolicyAction: valid values are %v", v, allowedAIGuardToolPolicyActionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AIGuardToolPolicyAction) IsValid() bool {
	for _, existing := range allowedAIGuardToolPolicyActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AIGuardToolPolicyAction value.
func (v AIGuardToolPolicyAction) Ptr() *AIGuardToolPolicyAction {
	return &v
}
