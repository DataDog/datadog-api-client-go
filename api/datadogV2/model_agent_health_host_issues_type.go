// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AgentHealthHostIssuesType The Agent Health issue resource type.
type AgentHealthHostIssuesType string

// List of AgentHealthHostIssuesType.
const (
	AGENTHEALTHHOSTISSUESTYPE_AGENT_HEALTH_ISSUES AgentHealthHostIssuesType = "agent_health_issues"
)

var allowedAgentHealthHostIssuesTypeEnumValues = []AgentHealthHostIssuesType{
	AGENTHEALTHHOSTISSUESTYPE_AGENT_HEALTH_ISSUES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AgentHealthHostIssuesType) GetAllowedValues() []AgentHealthHostIssuesType {
	return allowedAgentHealthHostIssuesTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AgentHealthHostIssuesType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AgentHealthHostIssuesType(value)
	return nil
}

// NewAgentHealthHostIssuesTypeFromValue returns a pointer to a valid AgentHealthHostIssuesType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAgentHealthHostIssuesTypeFromValue(v string) (*AgentHealthHostIssuesType, error) {
	ev := AgentHealthHostIssuesType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AgentHealthHostIssuesType: valid values are %v", v, allowedAgentHealthHostIssuesTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AgentHealthHostIssuesType) IsValid() bool {
	for _, existing := range allowedAgentHealthHostIssuesTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AgentHealthHostIssuesType value.
func (v AgentHealthHostIssuesType) Ptr() *AgentHealthHostIssuesType {
	return &v
}
