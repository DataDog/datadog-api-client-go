// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaAndFunctionEventsDataSource Test-only generated schema for the dashboard versioning proof of concept.
type FormulaAndFunctionEventsDataSource string

// List of FormulaAndFunctionEventsDataSource.
const (
	FORMULAANDFUNCTIONEVENTSDATASOURCE_EVENTS                        FormulaAndFunctionEventsDataSource = "events"
	FORMULAANDFUNCTIONEVENTSDATASOURCE_LOGS                          FormulaAndFunctionEventsDataSource = "logs"
	FORMULAANDFUNCTIONEVENTSDATASOURCE_RUM                           FormulaAndFunctionEventsDataSource = "rum"
	FORMULAANDFUNCTIONEVENTSDATASOURCE_SYNTHETICS_TEST_RUNS          FormulaAndFunctionEventsDataSource = "synthetics_test_runs"
	FORMULAANDFUNCTIONEVENTSDATASOURCE_PRODUCT_ANALYTICS             FormulaAndFunctionEventsDataSource = "product_analytics"
	FORMULAANDFUNCTIONEVENTSDATASOURCE_NETWORK                       FormulaAndFunctionEventsDataSource = "network"
	FORMULAANDFUNCTIONEVENTSDATASOURCE_PROFILES                      FormulaAndFunctionEventsDataSource = "profiles"
	FORMULAANDFUNCTIONEVENTSDATASOURCE_NETWORK_DEVICE_FLOWS          FormulaAndFunctionEventsDataSource = "network_device_flows"
	FORMULAANDFUNCTIONEVENTSDATASOURCE_RECOMMENDATIONS               FormulaAndFunctionEventsDataSource = "recommendations"
	FORMULAANDFUNCTIONEVENTSDATASOURCE_SECURITY_FINDINGS             FormulaAndFunctionEventsDataSource = "security_findings"
	FORMULAANDFUNCTIONEVENTSDATASOURCE_SECURITY_SIGNALS              FormulaAndFunctionEventsDataSource = "security_signals"
	FORMULAANDFUNCTIONEVENTSDATASOURCE_CI_PIPELINES                  FormulaAndFunctionEventsDataSource = "ci_pipelines"
	FORMULAANDFUNCTIONEVENTSDATASOURCE_DATABASE_QUERIES              FormulaAndFunctionEventsDataSource = "database_queries"
	FORMULAANDFUNCTIONEVENTSDATASOURCE_LLM_OBSERVABILITY             FormulaAndFunctionEventsDataSource = "llm_observability"
	FORMULAANDFUNCTIONEVENTSDATASOURCE_AI_USAGE                      FormulaAndFunctionEventsDataSource = "ai_usage"
	FORMULAANDFUNCTIONEVENTSDATASOURCE_AUDIT                         FormulaAndFunctionEventsDataSource = "audit"
	FORMULAANDFUNCTIONEVENTSDATASOURCE_CI_TESTS                      FormulaAndFunctionEventsDataSource = "ci_tests"
	FORMULAANDFUNCTIONEVENTSDATASOURCE_ERRORS                        FormulaAndFunctionEventsDataSource = "errors"
	FORMULAANDFUNCTIONEVENTSDATASOURCE_SPANS                         FormulaAndFunctionEventsDataSource = "spans"
	FORMULAANDFUNCTIONEVENTSDATASOURCE_SECURITY_RUNTIME              FormulaAndFunctionEventsDataSource = "security_runtime"
	FORMULAANDFUNCTIONEVENTSDATASOURCE_APP_SEC_SPANS                 FormulaAndFunctionEventsDataSource = "app_sec_spans"
	FORMULAANDFUNCTIONEVENTSDATASOURCE_APP_SEC_ENTITIES              FormulaAndFunctionEventsDataSource = "app_sec_entities"
	FORMULAANDFUNCTIONEVENTSDATASOURCE_BITS_AI_INVESTIGATIONS        FormulaAndFunctionEventsDataSource = "bits_ai_investigations"
	FORMULAANDFUNCTIONEVENTSDATASOURCE_CI_CODE_SCAN                  FormulaAndFunctionEventsDataSource = "ci_code_scan"
	FORMULAANDFUNCTIONEVENTSDATASOURCE_CI_DEPENDENCIES               FormulaAndFunctionEventsDataSource = "ci_dependencies"
	FORMULAANDFUNCTIONEVENTSDATASOURCE_CI_GATES                      FormulaAndFunctionEventsDataSource = "ci_gates"
	FORMULAANDFUNCTIONEVENTSDATASOURCE_CI_COVERAGE_REPORTS           FormulaAndFunctionEventsDataSource = "ci_coverage_reports"
	FORMULAANDFUNCTIONEVENTSDATASOURCE_DATA_OBSERVABILITY_EVENTS     FormulaAndFunctionEventsDataSource = "data_observability_events"
	FORMULAANDFUNCTIONEVENTSDATASOURCE_DORA_METRICS                  FormulaAndFunctionEventsDataSource = "dora_metrics"
	FORMULAANDFUNCTIONEVENTSDATASOURCE_COMPLIANCE_FINDINGS           FormulaAndFunctionEventsDataSource = "compliance_findings"
	FORMULAANDFUNCTIONEVENTSDATASOURCE_DEPLOYMENTS                   FormulaAndFunctionEventsDataSource = "deployments"
	FORMULAANDFUNCTIONEVENTSDATASOURCE_DEPLOYMENT_GATES              FormulaAndFunctionEventsDataSource = "deployment_gates"
	FORMULAANDFUNCTIONEVENTSDATASOURCE_EXPOSURES                     FormulaAndFunctionEventsDataSource = "exposures"
	FORMULAANDFUNCTIONEVENTSDATASOURCE_FLAG_EVALUATIONS              FormulaAndFunctionEventsDataSource = "flag_evaluations"
	FORMULAANDFUNCTIONEVENTSDATASOURCE_NETWORK_PATH                  FormulaAndFunctionEventsDataSource = "network_path"
	FORMULAANDFUNCTIONEVENTSDATASOURCE_OBSERVABILITY_PIPELINES       FormulaAndFunctionEventsDataSource = "observability_pipelines"
	FORMULAANDFUNCTIONEVENTSDATASOURCE_ON_CALL_EVENTS                FormulaAndFunctionEventsDataSource = "on_call_events"
	FORMULAANDFUNCTIONEVENTSDATASOURCE_ORCHESTRATION_TROUBLESHOOTING FormulaAndFunctionEventsDataSource = "orchestration_troubleshooting"
	FORMULAANDFUNCTIONEVENTSDATASOURCE_SECURITY_HISTORICAL_RESULTS   FormulaAndFunctionEventsDataSource = "security_historical_results"
	FORMULAANDFUNCTIONEVENTSDATASOURCE_SECURITY_INFO                 FormulaAndFunctionEventsDataSource = "security_info"
	FORMULAANDFUNCTIONEVENTSDATASOURCE_SOFTWARE_INVENTORY            FormulaAndFunctionEventsDataSource = "software_inventory"
	FORMULAANDFUNCTIONEVENTSDATASOURCE_SYNTHETICS_BATCHES            FormulaAndFunctionEventsDataSource = "synthetics_batches"
	FORMULAANDFUNCTIONEVENTSDATASOURCE_SYNTHETICS_MAP                FormulaAndFunctionEventsDataSource = "synthetics_map"
	FORMULAANDFUNCTIONEVENTSDATASOURCE_REPLAY                        FormulaAndFunctionEventsDataSource = "replay"
	FORMULAANDFUNCTIONEVENTSDATASOURCE_BEHAVIOR                      FormulaAndFunctionEventsDataSource = "behavior"
)

var allowedFormulaAndFunctionEventsDataSourceEnumValues = []FormulaAndFunctionEventsDataSource{
	FORMULAANDFUNCTIONEVENTSDATASOURCE_EVENTS,
	FORMULAANDFUNCTIONEVENTSDATASOURCE_LOGS,
	FORMULAANDFUNCTIONEVENTSDATASOURCE_RUM,
	FORMULAANDFUNCTIONEVENTSDATASOURCE_SYNTHETICS_TEST_RUNS,
	FORMULAANDFUNCTIONEVENTSDATASOURCE_PRODUCT_ANALYTICS,
	FORMULAANDFUNCTIONEVENTSDATASOURCE_NETWORK,
	FORMULAANDFUNCTIONEVENTSDATASOURCE_PROFILES,
	FORMULAANDFUNCTIONEVENTSDATASOURCE_NETWORK_DEVICE_FLOWS,
	FORMULAANDFUNCTIONEVENTSDATASOURCE_RECOMMENDATIONS,
	FORMULAANDFUNCTIONEVENTSDATASOURCE_SECURITY_FINDINGS,
	FORMULAANDFUNCTIONEVENTSDATASOURCE_SECURITY_SIGNALS,
	FORMULAANDFUNCTIONEVENTSDATASOURCE_CI_PIPELINES,
	FORMULAANDFUNCTIONEVENTSDATASOURCE_DATABASE_QUERIES,
	FORMULAANDFUNCTIONEVENTSDATASOURCE_LLM_OBSERVABILITY,
	FORMULAANDFUNCTIONEVENTSDATASOURCE_AI_USAGE,
	FORMULAANDFUNCTIONEVENTSDATASOURCE_AUDIT,
	FORMULAANDFUNCTIONEVENTSDATASOURCE_CI_TESTS,
	FORMULAANDFUNCTIONEVENTSDATASOURCE_ERRORS,
	FORMULAANDFUNCTIONEVENTSDATASOURCE_SPANS,
	FORMULAANDFUNCTIONEVENTSDATASOURCE_SECURITY_RUNTIME,
	FORMULAANDFUNCTIONEVENTSDATASOURCE_APP_SEC_SPANS,
	FORMULAANDFUNCTIONEVENTSDATASOURCE_APP_SEC_ENTITIES,
	FORMULAANDFUNCTIONEVENTSDATASOURCE_BITS_AI_INVESTIGATIONS,
	FORMULAANDFUNCTIONEVENTSDATASOURCE_CI_CODE_SCAN,
	FORMULAANDFUNCTIONEVENTSDATASOURCE_CI_DEPENDENCIES,
	FORMULAANDFUNCTIONEVENTSDATASOURCE_CI_GATES,
	FORMULAANDFUNCTIONEVENTSDATASOURCE_CI_COVERAGE_REPORTS,
	FORMULAANDFUNCTIONEVENTSDATASOURCE_DATA_OBSERVABILITY_EVENTS,
	FORMULAANDFUNCTIONEVENTSDATASOURCE_DORA_METRICS,
	FORMULAANDFUNCTIONEVENTSDATASOURCE_COMPLIANCE_FINDINGS,
	FORMULAANDFUNCTIONEVENTSDATASOURCE_DEPLOYMENTS,
	FORMULAANDFUNCTIONEVENTSDATASOURCE_DEPLOYMENT_GATES,
	FORMULAANDFUNCTIONEVENTSDATASOURCE_EXPOSURES,
	FORMULAANDFUNCTIONEVENTSDATASOURCE_FLAG_EVALUATIONS,
	FORMULAANDFUNCTIONEVENTSDATASOURCE_NETWORK_PATH,
	FORMULAANDFUNCTIONEVENTSDATASOURCE_OBSERVABILITY_PIPELINES,
	FORMULAANDFUNCTIONEVENTSDATASOURCE_ON_CALL_EVENTS,
	FORMULAANDFUNCTIONEVENTSDATASOURCE_ORCHESTRATION_TROUBLESHOOTING,
	FORMULAANDFUNCTIONEVENTSDATASOURCE_SECURITY_HISTORICAL_RESULTS,
	FORMULAANDFUNCTIONEVENTSDATASOURCE_SECURITY_INFO,
	FORMULAANDFUNCTIONEVENTSDATASOURCE_SOFTWARE_INVENTORY,
	FORMULAANDFUNCTIONEVENTSDATASOURCE_SYNTHETICS_BATCHES,
	FORMULAANDFUNCTIONEVENTSDATASOURCE_SYNTHETICS_MAP,
	FORMULAANDFUNCTIONEVENTSDATASOURCE_REPLAY,
	FORMULAANDFUNCTIONEVENTSDATASOURCE_BEHAVIOR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FormulaAndFunctionEventsDataSource) GetAllowedValues() []FormulaAndFunctionEventsDataSource {
	return allowedFormulaAndFunctionEventsDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FormulaAndFunctionEventsDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FormulaAndFunctionEventsDataSource(value)
	return nil
}

// NewFormulaAndFunctionEventsDataSourceFromValue returns a pointer to a valid FormulaAndFunctionEventsDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFormulaAndFunctionEventsDataSourceFromValue(v string) (*FormulaAndFunctionEventsDataSource, error) {
	ev := FormulaAndFunctionEventsDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FormulaAndFunctionEventsDataSource: valid values are %v", v, allowedFormulaAndFunctionEventsDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FormulaAndFunctionEventsDataSource) IsValid() bool {
	for _, existing := range allowedFormulaAndFunctionEventsDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FormulaAndFunctionEventsDataSource value.
func (v FormulaAndFunctionEventsDataSource) Ptr() *FormulaAndFunctionEventsDataSource {
	return &v
}
