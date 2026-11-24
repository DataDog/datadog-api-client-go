// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AgentHealthIssueResourceKind The kind of affected resource. The meaning of `resource_id` depends on this: hostname for `host`, DaemonSet UID for `deployment` (node Agent), cluster UUID for `cluster`, hostname of the host running the Cluster Agent for `cluster_agent`, or OpenTelemetry collector ID for `otel`.
type AgentHealthIssueResourceKind string

// List of AgentHealthIssueResourceKind.
const (
	AGENTHEALTHISSUERESOURCEKIND_HOST          AgentHealthIssueResourceKind = "host"
	AGENTHEALTHISSUERESOURCEKIND_DEPLOYMENT    AgentHealthIssueResourceKind = "deployment"
	AGENTHEALTHISSUERESOURCEKIND_CLUSTER       AgentHealthIssueResourceKind = "cluster"
	AGENTHEALTHISSUERESOURCEKIND_CLUSTER_AGENT AgentHealthIssueResourceKind = "cluster_agent"
	AGENTHEALTHISSUERESOURCEKIND_OTEL          AgentHealthIssueResourceKind = "otel"
)

var allowedAgentHealthIssueResourceKindEnumValues = []AgentHealthIssueResourceKind{
	AGENTHEALTHISSUERESOURCEKIND_HOST,
	AGENTHEALTHISSUERESOURCEKIND_DEPLOYMENT,
	AGENTHEALTHISSUERESOURCEKIND_CLUSTER,
	AGENTHEALTHISSUERESOURCEKIND_CLUSTER_AGENT,
	AGENTHEALTHISSUERESOURCEKIND_OTEL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AgentHealthIssueResourceKind) GetAllowedValues() []AgentHealthIssueResourceKind {
	return allowedAgentHealthIssueResourceKindEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AgentHealthIssueResourceKind) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AgentHealthIssueResourceKind(value)
	return nil
}

// NewAgentHealthIssueResourceKindFromValue returns a pointer to a valid AgentHealthIssueResourceKind
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAgentHealthIssueResourceKindFromValue(v string) (*AgentHealthIssueResourceKind, error) {
	ev := AgentHealthIssueResourceKind(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AgentHealthIssueResourceKind: valid values are %v", v, allowedAgentHealthIssueResourceKindEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AgentHealthIssueResourceKind) IsValid() bool {
	for _, existing := range allowedAgentHealthIssueResourceKindEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AgentHealthIssueResourceKind value.
func (v AgentHealthIssueResourceKind) Ptr() *AgentHealthIssueResourceKind {
	return &v
}
