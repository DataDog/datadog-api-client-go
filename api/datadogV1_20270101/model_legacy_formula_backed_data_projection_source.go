// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LegacyFormulaBackedDataProjectionSource Test-only generated schema for the dashboard versioning proof of concept.
type LegacyFormulaBackedDataProjectionSource string

// List of LegacyFormulaBackedDataProjectionSource.
const (
	LEGACYFORMULABACKEDDATAPROJECTIONSOURCE_EVENTS               LegacyFormulaBackedDataProjectionSource = "events"
	LEGACYFORMULABACKEDDATAPROJECTIONSOURCE_LOGS                 LegacyFormulaBackedDataProjectionSource = "logs"
	LEGACYFORMULABACKEDDATAPROJECTIONSOURCE_RUM                  LegacyFormulaBackedDataProjectionSource = "rum"
	LEGACYFORMULABACKEDDATAPROJECTIONSOURCE_SYNTHETICS_TEST_RUNS LegacyFormulaBackedDataProjectionSource = "synthetics_test_runs"
	LEGACYFORMULABACKEDDATAPROJECTIONSOURCE_PRODUCT_ANALYTICS    LegacyFormulaBackedDataProjectionSource = "product_analytics"
	LEGACYFORMULABACKEDDATAPROJECTIONSOURCE_NETWORK              LegacyFormulaBackedDataProjectionSource = "network"
	LEGACYFORMULABACKEDDATAPROJECTIONSOURCE_NETWORK_DEVICE_FLOWS LegacyFormulaBackedDataProjectionSource = "network_device_flows"
	LEGACYFORMULABACKEDDATAPROJECTIONSOURCE_SECURITY_FINDINGS    LegacyFormulaBackedDataProjectionSource = "security_findings"
	LEGACYFORMULABACKEDDATAPROJECTIONSOURCE_SECURITY_SIGNALS     LegacyFormulaBackedDataProjectionSource = "security_signals"
	LEGACYFORMULABACKEDDATAPROJECTIONSOURCE_CI_PIPELINES         LegacyFormulaBackedDataProjectionSource = "ci_pipelines"
	LEGACYFORMULABACKEDDATAPROJECTIONSOURCE_DATABASE_QUERIES     LegacyFormulaBackedDataProjectionSource = "database_queries"
	LEGACYFORMULABACKEDDATAPROJECTIONSOURCE_LLM_OBSERVABILITY    LegacyFormulaBackedDataProjectionSource = "llm_observability"
	LEGACYFORMULABACKEDDATAPROJECTIONSOURCE_AI_USAGE             LegacyFormulaBackedDataProjectionSource = "ai_usage"
	LEGACYFORMULABACKEDDATAPROJECTIONSOURCE_CI_TESTS             LegacyFormulaBackedDataProjectionSource = "ci_tests"
	LEGACYFORMULABACKEDDATAPROJECTIONSOURCE_ERRORS               LegacyFormulaBackedDataProjectionSource = "errors"
	LEGACYFORMULABACKEDDATAPROJECTIONSOURCE_SPANS                LegacyFormulaBackedDataProjectionSource = "spans"
	LEGACYFORMULABACKEDDATAPROJECTIONSOURCE_SECURITY_RUNTIME     LegacyFormulaBackedDataProjectionSource = "security_runtime"
)

var allowedLegacyFormulaBackedDataProjectionSourceEnumValues = []LegacyFormulaBackedDataProjectionSource{
	LEGACYFORMULABACKEDDATAPROJECTIONSOURCE_EVENTS,
	LEGACYFORMULABACKEDDATAPROJECTIONSOURCE_LOGS,
	LEGACYFORMULABACKEDDATAPROJECTIONSOURCE_RUM,
	LEGACYFORMULABACKEDDATAPROJECTIONSOURCE_SYNTHETICS_TEST_RUNS,
	LEGACYFORMULABACKEDDATAPROJECTIONSOURCE_PRODUCT_ANALYTICS,
	LEGACYFORMULABACKEDDATAPROJECTIONSOURCE_NETWORK,
	LEGACYFORMULABACKEDDATAPROJECTIONSOURCE_NETWORK_DEVICE_FLOWS,
	LEGACYFORMULABACKEDDATAPROJECTIONSOURCE_SECURITY_FINDINGS,
	LEGACYFORMULABACKEDDATAPROJECTIONSOURCE_SECURITY_SIGNALS,
	LEGACYFORMULABACKEDDATAPROJECTIONSOURCE_CI_PIPELINES,
	LEGACYFORMULABACKEDDATAPROJECTIONSOURCE_DATABASE_QUERIES,
	LEGACYFORMULABACKEDDATAPROJECTIONSOURCE_LLM_OBSERVABILITY,
	LEGACYFORMULABACKEDDATAPROJECTIONSOURCE_AI_USAGE,
	LEGACYFORMULABACKEDDATAPROJECTIONSOURCE_CI_TESTS,
	LEGACYFORMULABACKEDDATAPROJECTIONSOURCE_ERRORS,
	LEGACYFORMULABACKEDDATAPROJECTIONSOURCE_SPANS,
	LEGACYFORMULABACKEDDATAPROJECTIONSOURCE_SECURITY_RUNTIME,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LegacyFormulaBackedDataProjectionSource) GetAllowedValues() []LegacyFormulaBackedDataProjectionSource {
	return allowedLegacyFormulaBackedDataProjectionSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LegacyFormulaBackedDataProjectionSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LegacyFormulaBackedDataProjectionSource(value)
	return nil
}

// NewLegacyFormulaBackedDataProjectionSourceFromValue returns a pointer to a valid LegacyFormulaBackedDataProjectionSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLegacyFormulaBackedDataProjectionSourceFromValue(v string) (*LegacyFormulaBackedDataProjectionSource, error) {
	ev := LegacyFormulaBackedDataProjectionSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LegacyFormulaBackedDataProjectionSource: valid values are %v", v, allowedLegacyFormulaBackedDataProjectionSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LegacyFormulaBackedDataProjectionSource) IsValid() bool {
	for _, existing := range allowedLegacyFormulaBackedDataProjectionSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LegacyFormulaBackedDataProjectionSource value.
func (v LegacyFormulaBackedDataProjectionSource) Ptr() *LegacyFormulaBackedDataProjectionSource {
	return &v
}
