// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AgentHealthIssueSeverity The severity of the issue.
type AgentHealthIssueSeverity string

// List of AgentHealthIssueSeverity.
const (
	AGENTHEALTHISSUESEVERITY_LOW     AgentHealthIssueSeverity = "low"
	AGENTHEALTHISSUESEVERITY_MEDIUM  AgentHealthIssueSeverity = "medium"
	AGENTHEALTHISSUESEVERITY_HIGH    AgentHealthIssueSeverity = "high"
	AGENTHEALTHISSUESEVERITY_UNKNOWN AgentHealthIssueSeverity = "unknown"
)

var allowedAgentHealthIssueSeverityEnumValues = []AgentHealthIssueSeverity{
	AGENTHEALTHISSUESEVERITY_LOW,
	AGENTHEALTHISSUESEVERITY_MEDIUM,
	AGENTHEALTHISSUESEVERITY_HIGH,
	AGENTHEALTHISSUESEVERITY_UNKNOWN,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AgentHealthIssueSeverity) GetAllowedValues() []AgentHealthIssueSeverity {
	return allowedAgentHealthIssueSeverityEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AgentHealthIssueSeverity) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AgentHealthIssueSeverity(value)
	return nil
}

// NewAgentHealthIssueSeverityFromValue returns a pointer to a valid AgentHealthIssueSeverity
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAgentHealthIssueSeverityFromValue(v string) (*AgentHealthIssueSeverity, error) {
	ev := AgentHealthIssueSeverity(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AgentHealthIssueSeverity: valid values are %v", v, allowedAgentHealthIssueSeverityEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AgentHealthIssueSeverity) IsValid() bool {
	for _, existing := range allowedAgentHealthIssueSeverityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AgentHealthIssueSeverity value.
func (v AgentHealthIssueSeverity) Ptr() *AgentHealthIssueSeverity {
	return &v
}
