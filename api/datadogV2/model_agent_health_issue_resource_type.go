// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AgentHealthIssueResourceType The Agent Health issue resource type.
type AgentHealthIssueResourceType string

// List of AgentHealthIssueResourceType.
const (
	AGENTHEALTHISSUERESOURCETYPE_AGENT_HEALTH_ISSUES AgentHealthIssueResourceType = "agent_health_issues"
)

var allowedAgentHealthIssueResourceTypeEnumValues = []AgentHealthIssueResourceType{
	AGENTHEALTHISSUERESOURCETYPE_AGENT_HEALTH_ISSUES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AgentHealthIssueResourceType) GetAllowedValues() []AgentHealthIssueResourceType {
	return allowedAgentHealthIssueResourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AgentHealthIssueResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AgentHealthIssueResourceType(value)
	return nil
}

// NewAgentHealthIssueResourceTypeFromValue returns a pointer to a valid AgentHealthIssueResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAgentHealthIssueResourceTypeFromValue(v string) (*AgentHealthIssueResourceType, error) {
	ev := AgentHealthIssueResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AgentHealthIssueResourceType: valid values are %v", v, allowedAgentHealthIssueResourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AgentHealthIssueResourceType) IsValid() bool {
	for _, existing := range allowedAgentHealthIssueResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AgentHealthIssueResourceType value.
func (v AgentHealthIssueResourceType) Ptr() *AgentHealthIssueResourceType {
	return &v
}
