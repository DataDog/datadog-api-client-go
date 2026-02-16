// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ChatMessageRole The role of the message sender.
type ChatMessageRole string

// List of ChatMessageRole.
const (
	CHATMESSAGEROLE_USER      ChatMessageRole = "user"
	CHATMESSAGEROLE_ASSISTANT ChatMessageRole = "assistant"
	CHATMESSAGEROLE_SYSTEM    ChatMessageRole = "system"
)

var allowedChatMessageRoleEnumValues = []ChatMessageRole{
	CHATMESSAGEROLE_USER,
	CHATMESSAGEROLE_ASSISTANT,
	CHATMESSAGEROLE_SYSTEM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ChatMessageRole) GetAllowedValues() []ChatMessageRole {
	return allowedChatMessageRoleEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ChatMessageRole) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ChatMessageRole(value)
	return nil
}

// NewChatMessageRoleFromValue returns a pointer to a valid ChatMessageRole
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewChatMessageRoleFromValue(v string) (*ChatMessageRole, error) {
	ev := ChatMessageRole(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ChatMessageRole: valid values are %v", v, allowedChatMessageRoleEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ChatMessageRole) IsValid() bool {
	for _, existing := range allowedChatMessageRoleEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ChatMessageRole value.
func (v ChatMessageRole) Ptr() *ChatMessageRole {
	return &v
}
