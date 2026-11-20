// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ContextMenuLinkLabel Test-only generated schema for the dashboard versioning proof of concept.
type ContextMenuLinkLabel string

// List of ContextMenuLinkLabel.
const (
	CONTEXTMENULINKLABEL_EVENTS               ContextMenuLinkLabel = "events"
	CONTEXTMENULINKLABEL_LOGS                 ContextMenuLinkLabel = "logs"
	CONTEXTMENULINKLABEL_RUM                  ContextMenuLinkLabel = "rum"
	CONTEXTMENULINKLABEL_PRODUCT_ANALYTICS    ContextMenuLinkLabel = "product_analytics"
	CONTEXTMENULINKLABEL_NETWORK              ContextMenuLinkLabel = "network"
	CONTEXTMENULINKLABEL_CLOUD_COST           ContextMenuLinkLabel = "cloud_cost"
	CONTEXTMENULINKLABEL_FUNCTIONS            ContextMenuLinkLabel = "functions"
	CONTEXTMENULINKLABEL_HOSTS                ContextMenuLinkLabel = "hosts"
	CONTEXTMENULINKLABEL_TRACES               ContextMenuLinkLabel = "traces"
	CONTEXTMENULINKLABEL_PROFILES             ContextMenuLinkLabel = "profiles"
	CONTEXTMENULINKLABEL_PROCESSES            ContextMenuLinkLabel = "processes"
	CONTEXTMENULINKLABEL_CONTAINERS           ContextMenuLinkLabel = "containers"
	CONTEXTMENULINKLABEL_CHANGE_TRACKING      ContextMenuLinkLabel = "change_tracking"
	CONTEXTMENULINKLABEL_AUDIT_TRAIL          ContextMenuLinkLabel = "audit_trail"
	CONTEXTMENULINKLABEL_NETWORK_DEVICE_FLOWS ContextMenuLinkLabel = "network_device_flows"
	CONTEXTMENULINKLABEL_DEVICES              ContextMenuLinkLabel = "devices"
	CONTEXTMENULINKLABEL_RECOMMENDATIONS      ContextMenuLinkLabel = "recommendations"
	CONTEXTMENULINKLABEL_SECURITY_FINDINGS    ContextMenuLinkLabel = "security_findings"
	CONTEXTMENULINKLABEL_SECURITY_SIGNALS     ContextMenuLinkLabel = "security_signals"
	CONTEXTMENULINKLABEL_SECURITY_TRACES      ContextMenuLinkLabel = "security_traces"
	CONTEXTMENULINKLABEL_INCIDENTS            ContextMenuLinkLabel = "incidents"
	CONTEXTMENULINKLABEL_CASES                ContextMenuLinkLabel = "cases"
	CONTEXTMENULINKLABEL_METRICS_SUMMARY      ContextMenuLinkLabel = "metrics_summary"
	CONTEXTMENULINKLABEL_PODS                 ContextMenuLinkLabel = "pods"
	CONTEXTMENULINKLABEL_CI_PIPELINES         ContextMenuLinkLabel = "ci_pipelines"
	CONTEXTMENULINKLABEL_DATABASE_QUERIES     ContextMenuLinkLabel = "database_queries"
	CONTEXTMENULINKLABEL_WATCH_SESSIONS       ContextMenuLinkLabel = "watch_sessions"
	CONTEXTMENULINKLABEL_GPU_MONITORING       ContextMenuLinkLabel = "gpu_monitoring"
	CONTEXTMENULINKLABEL_LLM_OBSERVABILITY    ContextMenuLinkLabel = "llm_observability"
)

var allowedContextMenuLinkLabelEnumValues = []ContextMenuLinkLabel{
	CONTEXTMENULINKLABEL_EVENTS,
	CONTEXTMENULINKLABEL_LOGS,
	CONTEXTMENULINKLABEL_RUM,
	CONTEXTMENULINKLABEL_PRODUCT_ANALYTICS,
	CONTEXTMENULINKLABEL_NETWORK,
	CONTEXTMENULINKLABEL_CLOUD_COST,
	CONTEXTMENULINKLABEL_FUNCTIONS,
	CONTEXTMENULINKLABEL_HOSTS,
	CONTEXTMENULINKLABEL_TRACES,
	CONTEXTMENULINKLABEL_PROFILES,
	CONTEXTMENULINKLABEL_PROCESSES,
	CONTEXTMENULINKLABEL_CONTAINERS,
	CONTEXTMENULINKLABEL_CHANGE_TRACKING,
	CONTEXTMENULINKLABEL_AUDIT_TRAIL,
	CONTEXTMENULINKLABEL_NETWORK_DEVICE_FLOWS,
	CONTEXTMENULINKLABEL_DEVICES,
	CONTEXTMENULINKLABEL_RECOMMENDATIONS,
	CONTEXTMENULINKLABEL_SECURITY_FINDINGS,
	CONTEXTMENULINKLABEL_SECURITY_SIGNALS,
	CONTEXTMENULINKLABEL_SECURITY_TRACES,
	CONTEXTMENULINKLABEL_INCIDENTS,
	CONTEXTMENULINKLABEL_CASES,
	CONTEXTMENULINKLABEL_METRICS_SUMMARY,
	CONTEXTMENULINKLABEL_PODS,
	CONTEXTMENULINKLABEL_CI_PIPELINES,
	CONTEXTMENULINKLABEL_DATABASE_QUERIES,
	CONTEXTMENULINKLABEL_WATCH_SESSIONS,
	CONTEXTMENULINKLABEL_GPU_MONITORING,
	CONTEXTMENULINKLABEL_LLM_OBSERVABILITY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ContextMenuLinkLabel) GetAllowedValues() []ContextMenuLinkLabel {
	return allowedContextMenuLinkLabelEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ContextMenuLinkLabel) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ContextMenuLinkLabel(value)
	return nil
}

// NewContextMenuLinkLabelFromValue returns a pointer to a valid ContextMenuLinkLabel
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewContextMenuLinkLabelFromValue(v string) (*ContextMenuLinkLabel, error) {
	ev := ContextMenuLinkLabel(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ContextMenuLinkLabel: valid values are %v", v, allowedContextMenuLinkLabelEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ContextMenuLinkLabel) IsValid() bool {
	for _, existing := range allowedContextMenuLinkLabelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ContextMenuLinkLabel value.
func (v ContextMenuLinkLabel) Ptr() *ContextMenuLinkLabel {
	return &v
}
