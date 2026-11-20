// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EventPlatformStreamDataSource Test-only generated schema for the dashboard versioning proof of concept.
type EventPlatformStreamDataSource string

// List of EventPlatformStreamDataSource.
const (
	EVENTPLATFORMSTREAMDATASOURCE_LOGS_STREAM                         EventPlatformStreamDataSource = "logs_stream"
	EVENTPLATFORMSTREAMDATASOURCE_AUDIT_STREAM                        EventPlatformStreamDataSource = "audit_stream"
	EVENTPLATFORMSTREAMDATASOURCE_TRACE_STREAM                        EventPlatformStreamDataSource = "trace_stream"
	EVENTPLATFORMSTREAMDATASOURCE_ERRORS_STREAM                       EventPlatformStreamDataSource = "errors_stream"
	EVENTPLATFORMSTREAMDATASOURCE_EVENT_STREAM                        EventPlatformStreamDataSource = "event_stream"
	EVENTPLATFORMSTREAMDATASOURCE_RUM_STREAM                          EventPlatformStreamDataSource = "rum_stream"
	EVENTPLATFORMSTREAMDATASOURCE_CI_PIPELINE_STREAM                  EventPlatformStreamDataSource = "ci_pipeline_stream"
	EVENTPLATFORMSTREAMDATASOURCE_CI_TEST_STREAM                      EventPlatformStreamDataSource = "ci_test_stream"
	EVENTPLATFORMSTREAMDATASOURCE_CI_GATE_STREAM                      EventPlatformStreamDataSource = "ci_gate_stream"
	EVENTPLATFORMSTREAMDATASOURCE_CI_CODE_SCAN_STREAM                 EventPlatformStreamDataSource = "ci_code_scan_stream"
	EVENTPLATFORMSTREAMDATASOURCE_CI_DEPENDENCY_STREAM                EventPlatformStreamDataSource = "ci_dependency_stream"
	EVENTPLATFORMSTREAMDATASOURCE_DEPLOYMENT_STREAM                   EventPlatformStreamDataSource = "deployment_stream"
	EVENTPLATFORMSTREAMDATASOURCE_RECOMMENDATIONS_STREAM              EventPlatformStreamDataSource = "recommendations_stream"
	EVENTPLATFORMSTREAMDATASOURCE_SYNTHETICS_BATCH_STREAM             EventPlatformStreamDataSource = "synthetics_batch_stream"
	EVENTPLATFORMSTREAMDATASOURCE_SYNTHETICS_TEST_RUNS_STREAM         EventPlatformStreamDataSource = "synthetics_test_runs_stream"
	EVENTPLATFORMSTREAMDATASOURCE_OBSERVABILITY_PIPELINES_STREAM      EventPlatformStreamDataSource = "observability_pipelines_stream"
	EVENTPLATFORMSTREAMDATASOURCE_SECURITY_RUNTIME_STREAM             EventPlatformStreamDataSource = "security_runtime_stream"
	EVENTPLATFORMSTREAMDATASOURCE_SECURITY_SIGNALS_STREAM             EventPlatformStreamDataSource = "security_signals_stream"
	EVENTPLATFORMSTREAMDATASOURCE_PRODUCT_ANALYTICS_STREAM            EventPlatformStreamDataSource = "product_analytics_stream"
	EVENTPLATFORMSTREAMDATASOURCE_APP_SEC_SPAN_STREAM                 EventPlatformStreamDataSource = "app_sec_span_stream"
	EVENTPLATFORMSTREAMDATASOURCE_APP_SEC_ENTITIES_STREAM             EventPlatformStreamDataSource = "app_sec_entities_stream"
	EVENTPLATFORMSTREAMDATASOURCE_COMPLIANCE_EVENT_STREAM             EventPlatformStreamDataSource = "compliance_event_stream"
	EVENTPLATFORMSTREAMDATASOURCE_COMPLIANCE_FINDINGS_STREAM          EventPlatformStreamDataSource = "compliance_findings_stream"
	EVENTPLATFORMSTREAMDATASOURCE_SECURITY_FINDINGS_STREAM            EventPlatformStreamDataSource = "security_findings_stream"
	EVENTPLATFORMSTREAMDATASOURCE_SUPPLY_CHAIN_FIREWALL_REPORT_STREAM EventPlatformStreamDataSource = "supply_chain_firewall_report_stream"
	EVENTPLATFORMSTREAMDATASOURCE_DATA_OBSERVABILITY_STREAM           EventPlatformStreamDataSource = "data_observability_stream"
	EVENTPLATFORMSTREAMDATASOURCE_DATABASE_QUERY_STREAM               EventPlatformStreamDataSource = "database_query_stream"
	EVENTPLATFORMSTREAMDATASOURCE_DATABASE_RAW_QUERY_STREAM           EventPlatformStreamDataSource = "database_raw_query_stream"
	EVENTPLATFORMSTREAMDATASOURCE_DEPLOYMENT_GATES_STREAM             EventPlatformStreamDataSource = "deployment_gates_stream"
	EVENTPLATFORMSTREAMDATASOURCE_LLM_OBSERVABILITY_STREAM            EventPlatformStreamDataSource = "llm_observability_stream"
	EVENTPLATFORMSTREAMDATASOURCE_NETWORK_DEVICE_FLOW_STREAM          EventPlatformStreamDataSource = "network_device_flow_stream"
	EVENTPLATFORMSTREAMDATASOURCE_NETWORK_FLOW_STREAM                 EventPlatformStreamDataSource = "network_flow_stream"
	EVENTPLATFORMSTREAMDATASOURCE_NETWORK_PATH_STREAM                 EventPlatformStreamDataSource = "network_path_stream"
	EVENTPLATFORMSTREAMDATASOURCE_ON_CALL_EVENTS_STREAM               EventPlatformStreamDataSource = "on_call_events_stream"
	EVENTPLATFORMSTREAMDATASOURCE_PROFILE_STREAM                      EventPlatformStreamDataSource = "profile_stream"
	EVENTPLATFORMSTREAMDATASOURCE_SECURITY_DUMP_STREAM                EventPlatformStreamDataSource = "security_dump_stream"
	EVENTPLATFORMSTREAMDATASOURCE_SECURITY_HISTORICAL_RESULTS_STREAM  EventPlatformStreamDataSource = "security_historical_results_stream"
	EVENTPLATFORMSTREAMDATASOURCE_SESSION_REPLAY_STREAM               EventPlatformStreamDataSource = "session_replay_stream"
	EVENTPLATFORMSTREAMDATASOURCE_SOFTWARE_INVENTORY_STREAM           EventPlatformStreamDataSource = "software_inventory_stream"
	EVENTPLATFORMSTREAMDATASOURCE_CRITICAL_ASSET_STREAM               EventPlatformStreamDataSource = "critical_asset_stream"
)

var allowedEventPlatformStreamDataSourceEnumValues = []EventPlatformStreamDataSource{
	EVENTPLATFORMSTREAMDATASOURCE_LOGS_STREAM,
	EVENTPLATFORMSTREAMDATASOURCE_AUDIT_STREAM,
	EVENTPLATFORMSTREAMDATASOURCE_TRACE_STREAM,
	EVENTPLATFORMSTREAMDATASOURCE_ERRORS_STREAM,
	EVENTPLATFORMSTREAMDATASOURCE_EVENT_STREAM,
	EVENTPLATFORMSTREAMDATASOURCE_RUM_STREAM,
	EVENTPLATFORMSTREAMDATASOURCE_CI_PIPELINE_STREAM,
	EVENTPLATFORMSTREAMDATASOURCE_CI_TEST_STREAM,
	EVENTPLATFORMSTREAMDATASOURCE_CI_GATE_STREAM,
	EVENTPLATFORMSTREAMDATASOURCE_CI_CODE_SCAN_STREAM,
	EVENTPLATFORMSTREAMDATASOURCE_CI_DEPENDENCY_STREAM,
	EVENTPLATFORMSTREAMDATASOURCE_DEPLOYMENT_STREAM,
	EVENTPLATFORMSTREAMDATASOURCE_RECOMMENDATIONS_STREAM,
	EVENTPLATFORMSTREAMDATASOURCE_SYNTHETICS_BATCH_STREAM,
	EVENTPLATFORMSTREAMDATASOURCE_SYNTHETICS_TEST_RUNS_STREAM,
	EVENTPLATFORMSTREAMDATASOURCE_OBSERVABILITY_PIPELINES_STREAM,
	EVENTPLATFORMSTREAMDATASOURCE_SECURITY_RUNTIME_STREAM,
	EVENTPLATFORMSTREAMDATASOURCE_SECURITY_SIGNALS_STREAM,
	EVENTPLATFORMSTREAMDATASOURCE_PRODUCT_ANALYTICS_STREAM,
	EVENTPLATFORMSTREAMDATASOURCE_APP_SEC_SPAN_STREAM,
	EVENTPLATFORMSTREAMDATASOURCE_APP_SEC_ENTITIES_STREAM,
	EVENTPLATFORMSTREAMDATASOURCE_COMPLIANCE_EVENT_STREAM,
	EVENTPLATFORMSTREAMDATASOURCE_COMPLIANCE_FINDINGS_STREAM,
	EVENTPLATFORMSTREAMDATASOURCE_SECURITY_FINDINGS_STREAM,
	EVENTPLATFORMSTREAMDATASOURCE_SUPPLY_CHAIN_FIREWALL_REPORT_STREAM,
	EVENTPLATFORMSTREAMDATASOURCE_DATA_OBSERVABILITY_STREAM,
	EVENTPLATFORMSTREAMDATASOURCE_DATABASE_QUERY_STREAM,
	EVENTPLATFORMSTREAMDATASOURCE_DATABASE_RAW_QUERY_STREAM,
	EVENTPLATFORMSTREAMDATASOURCE_DEPLOYMENT_GATES_STREAM,
	EVENTPLATFORMSTREAMDATASOURCE_LLM_OBSERVABILITY_STREAM,
	EVENTPLATFORMSTREAMDATASOURCE_NETWORK_DEVICE_FLOW_STREAM,
	EVENTPLATFORMSTREAMDATASOURCE_NETWORK_FLOW_STREAM,
	EVENTPLATFORMSTREAMDATASOURCE_NETWORK_PATH_STREAM,
	EVENTPLATFORMSTREAMDATASOURCE_ON_CALL_EVENTS_STREAM,
	EVENTPLATFORMSTREAMDATASOURCE_PROFILE_STREAM,
	EVENTPLATFORMSTREAMDATASOURCE_SECURITY_DUMP_STREAM,
	EVENTPLATFORMSTREAMDATASOURCE_SECURITY_HISTORICAL_RESULTS_STREAM,
	EVENTPLATFORMSTREAMDATASOURCE_SESSION_REPLAY_STREAM,
	EVENTPLATFORMSTREAMDATASOURCE_SOFTWARE_INVENTORY_STREAM,
	EVENTPLATFORMSTREAMDATASOURCE_CRITICAL_ASSET_STREAM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EventPlatformStreamDataSource) GetAllowedValues() []EventPlatformStreamDataSource {
	return allowedEventPlatformStreamDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EventPlatformStreamDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EventPlatformStreamDataSource(value)
	return nil
}

// NewEventPlatformStreamDataSourceFromValue returns a pointer to a valid EventPlatformStreamDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEventPlatformStreamDataSourceFromValue(v string) (*EventPlatformStreamDataSource, error) {
	ev := EventPlatformStreamDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EventPlatformStreamDataSource: valid values are %v", v, allowedEventPlatformStreamDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EventPlatformStreamDataSource) IsValid() bool {
	for _, existing := range allowedEventPlatformStreamDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EventPlatformStreamDataSource value.
func (v EventPlatformStreamDataSource) Ptr() *EventPlatformStreamDataSource {
	return &v
}
