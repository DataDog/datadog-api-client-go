// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AdvancedEventPlatformSource Test-only generated schema for the dashboard versioning proof of concept.
type AdvancedEventPlatformSource string

// List of AdvancedEventPlatformSource.
const (
	ADVANCEDEVENTPLATFORMSOURCE_EVENTS               AdvancedEventPlatformSource = "events"
	ADVANCEDEVENTPLATFORMSOURCE_LOGS                 AdvancedEventPlatformSource = "logs"
	ADVANCEDEVENTPLATFORMSOURCE_RUM                  AdvancedEventPlatformSource = "rum"
	ADVANCEDEVENTPLATFORMSOURCE_SYNTHETICS_TEST_RUNS AdvancedEventPlatformSource = "synthetics_test_runs"
	ADVANCEDEVENTPLATFORMSOURCE_PRODUCT_ANALYTICS    AdvancedEventPlatformSource = "product_analytics"
	ADVANCEDEVENTPLATFORMSOURCE_NETWORK              AdvancedEventPlatformSource = "network"
	ADVANCEDEVENTPLATFORMSOURCE_NETWORK_DEVICE_FLOWS AdvancedEventPlatformSource = "network_device_flows"
	ADVANCEDEVENTPLATFORMSOURCE_SECURITY_FINDINGS    AdvancedEventPlatformSource = "security_findings"
	ADVANCEDEVENTPLATFORMSOURCE_SECURITY_SIGNALS     AdvancedEventPlatformSource = "security_signals"
	ADVANCEDEVENTPLATFORMSOURCE_CI_PIPELINES         AdvancedEventPlatformSource = "ci_pipelines"
	ADVANCEDEVENTPLATFORMSOURCE_DATABASE_QUERIES     AdvancedEventPlatformSource = "database_queries"
	ADVANCEDEVENTPLATFORMSOURCE_LLM_OBSERVABILITY    AdvancedEventPlatformSource = "llm_observability"
	ADVANCEDEVENTPLATFORMSOURCE_AI_USAGE             AdvancedEventPlatformSource = "ai_usage"
	ADVANCEDEVENTPLATFORMSOURCE_AUDIT                AdvancedEventPlatformSource = "audit"
	ADVANCEDEVENTPLATFORMSOURCE_CI_TESTS             AdvancedEventPlatformSource = "ci_tests"
	ADVANCEDEVENTPLATFORMSOURCE_ERRORS               AdvancedEventPlatformSource = "errors"
	ADVANCEDEVENTPLATFORMSOURCE_MONITORS             AdvancedEventPlatformSource = "monitors"
	ADVANCEDEVENTPLATFORMSOURCE_MONITOR_GROUPS       AdvancedEventPlatformSource = "monitor_groups"
	ADVANCEDEVENTPLATFORMSOURCE_SPANS                AdvancedEventPlatformSource = "spans"
	ADVANCEDEVENTPLATFORMSOURCE_SECURITY_RUNTIME     AdvancedEventPlatformSource = "security_runtime"
	ADVANCEDEVENTPLATFORMSOURCE_SYNTHETICS_MATCHES   AdvancedEventPlatformSource = "synthetics_matches"
)

var allowedAdvancedEventPlatformSourceEnumValues = []AdvancedEventPlatformSource{
	ADVANCEDEVENTPLATFORMSOURCE_EVENTS,
	ADVANCEDEVENTPLATFORMSOURCE_LOGS,
	ADVANCEDEVENTPLATFORMSOURCE_RUM,
	ADVANCEDEVENTPLATFORMSOURCE_SYNTHETICS_TEST_RUNS,
	ADVANCEDEVENTPLATFORMSOURCE_PRODUCT_ANALYTICS,
	ADVANCEDEVENTPLATFORMSOURCE_NETWORK,
	ADVANCEDEVENTPLATFORMSOURCE_NETWORK_DEVICE_FLOWS,
	ADVANCEDEVENTPLATFORMSOURCE_SECURITY_FINDINGS,
	ADVANCEDEVENTPLATFORMSOURCE_SECURITY_SIGNALS,
	ADVANCEDEVENTPLATFORMSOURCE_CI_PIPELINES,
	ADVANCEDEVENTPLATFORMSOURCE_DATABASE_QUERIES,
	ADVANCEDEVENTPLATFORMSOURCE_LLM_OBSERVABILITY,
	ADVANCEDEVENTPLATFORMSOURCE_AI_USAGE,
	ADVANCEDEVENTPLATFORMSOURCE_AUDIT,
	ADVANCEDEVENTPLATFORMSOURCE_CI_TESTS,
	ADVANCEDEVENTPLATFORMSOURCE_ERRORS,
	ADVANCEDEVENTPLATFORMSOURCE_MONITORS,
	ADVANCEDEVENTPLATFORMSOURCE_MONITOR_GROUPS,
	ADVANCEDEVENTPLATFORMSOURCE_SPANS,
	ADVANCEDEVENTPLATFORMSOURCE_SECURITY_RUNTIME,
	ADVANCEDEVENTPLATFORMSOURCE_SYNTHETICS_MATCHES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AdvancedEventPlatformSource) GetAllowedValues() []AdvancedEventPlatformSource {
	return allowedAdvancedEventPlatformSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AdvancedEventPlatformSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AdvancedEventPlatformSource(value)
	return nil
}

// NewAdvancedEventPlatformSourceFromValue returns a pointer to a valid AdvancedEventPlatformSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAdvancedEventPlatformSourceFromValue(v string) (*AdvancedEventPlatformSource, error) {
	ev := AdvancedEventPlatformSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AdvancedEventPlatformSource: valid values are %v", v, allowedAdvancedEventPlatformSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AdvancedEventPlatformSource) IsValid() bool {
	for _, existing := range allowedAdvancedEventPlatformSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AdvancedEventPlatformSource value.
func (v AdvancedEventPlatformSource) Ptr() *AdvancedEventPlatformSource {
	return &v
}
