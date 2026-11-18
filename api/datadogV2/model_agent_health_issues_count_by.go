// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AgentHealthIssuesCountBy The dimension to aggregate issue counts by. Only `issue_type` is supported.
type AgentHealthIssuesCountBy string

// List of AgentHealthIssuesCountBy.
const (
	AGENTHEALTHISSUESCOUNTBY_ISSUE_TYPE AgentHealthIssuesCountBy = "issue_type"
)

var allowedAgentHealthIssuesCountByEnumValues = []AgentHealthIssuesCountBy{
	AGENTHEALTHISSUESCOUNTBY_ISSUE_TYPE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AgentHealthIssuesCountBy) GetAllowedValues() []AgentHealthIssuesCountBy {
	return allowedAgentHealthIssuesCountByEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AgentHealthIssuesCountBy) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AgentHealthIssuesCountBy(value)
	return nil
}

// NewAgentHealthIssuesCountByFromValue returns a pointer to a valid AgentHealthIssuesCountBy
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAgentHealthIssuesCountByFromValue(v string) (*AgentHealthIssuesCountBy, error) {
	ev := AgentHealthIssuesCountBy(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AgentHealthIssuesCountBy: valid values are %v", v, allowedAgentHealthIssuesCountByEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AgentHealthIssuesCountBy) IsValid() bool {
	for _, existing := range allowedAgentHealthIssuesCountByEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AgentHealthIssuesCountBy value.
func (v AgentHealthIssuesCountBy) Ptr() *AgentHealthIssuesCountBy {
	return &v
}
