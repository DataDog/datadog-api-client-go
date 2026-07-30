// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AIGuardMessageRole The role of the message author in the conversation.
type AIGuardMessageRole string

// List of AIGuardMessageRole.
const (
	AIGUARDMESSAGEROLE_USER      AIGuardMessageRole = "user"
	AIGUARDMESSAGEROLE_ASSISTANT AIGuardMessageRole = "assistant"
	AIGUARDMESSAGEROLE_SYSTEM    AIGuardMessageRole = "system"
	AIGUARDMESSAGEROLE_TOOL      AIGuardMessageRole = "tool"
	AIGUARDMESSAGEROLE_DEVELOPER AIGuardMessageRole = "developer"
)

var allowedAIGuardMessageRoleEnumValues = []AIGuardMessageRole{
	AIGUARDMESSAGEROLE_USER,
	AIGUARDMESSAGEROLE_ASSISTANT,
	AIGUARDMESSAGEROLE_SYSTEM,
	AIGUARDMESSAGEROLE_TOOL,
	AIGUARDMESSAGEROLE_DEVELOPER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AIGuardMessageRole) GetAllowedValues() []AIGuardMessageRole {
	return allowedAIGuardMessageRoleEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AIGuardMessageRole) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AIGuardMessageRole(value)
	return nil
}

// NewAIGuardMessageRoleFromValue returns a pointer to a valid AIGuardMessageRole
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAIGuardMessageRoleFromValue(v string) (*AIGuardMessageRole, error) {
	ev := AIGuardMessageRole(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AIGuardMessageRole: valid values are %v", v, allowedAIGuardMessageRoleEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AIGuardMessageRole) IsValid() bool {
	for _, existing := range allowedAIGuardMessageRoleEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AIGuardMessageRole value.
func (v AIGuardMessageRole) Ptr() *AIGuardMessageRole {
	return &v
}
