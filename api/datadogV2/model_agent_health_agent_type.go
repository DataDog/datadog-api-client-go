// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AgentHealthAgentType Agent type associated with the resource.
type AgentHealthAgentType string

// List of AgentHealthAgentType.
const (
	AGENTHEALTHAGENTTYPE_NODE    AgentHealthAgentType = "node"
	AGENTHEALTHAGENTTYPE_CLUSTER AgentHealthAgentType = "cluster"
)

var allowedAgentHealthAgentTypeEnumValues = []AgentHealthAgentType{
	AGENTHEALTHAGENTTYPE_NODE,
	AGENTHEALTHAGENTTYPE_CLUSTER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AgentHealthAgentType) GetAllowedValues() []AgentHealthAgentType {
	return allowedAgentHealthAgentTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AgentHealthAgentType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AgentHealthAgentType(value)
	return nil
}

// NewAgentHealthAgentTypeFromValue returns a pointer to a valid AgentHealthAgentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAgentHealthAgentTypeFromValue(v string) (*AgentHealthAgentType, error) {
	ev := AgentHealthAgentType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AgentHealthAgentType: valid values are %v", v, allowedAgentHealthAgentTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AgentHealthAgentType) IsValid() bool {
	for _, existing := range allowedAgentHealthAgentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AgentHealthAgentType value.
func (v AgentHealthAgentType) Ptr() *AgentHealthAgentType {
	return &v
}
