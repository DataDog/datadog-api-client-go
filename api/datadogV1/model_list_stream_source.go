// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ListStreamSource Source from which to query items to display in the stream. apm_issue_stream, rum_issue_stream, and logs_issue_stream are deprecated. Use issue_stream instead. apm_recommendations_stream is used to query APM recommendations, and supports filtering by environment, services, teams, recommendation types, and status.
type ListStreamSource string

// List of ListStreamSource.
const (
	LISTSTREAMSOURCE_LOGS_STREAM                ListStreamSource = "logs_stream"
	LISTSTREAMSOURCE_AUDIT_STREAM               ListStreamSource = "audit_stream"
	LISTSTREAMSOURCE_CI_PIPELINE_STREAM         ListStreamSource = "ci_pipeline_stream"
	LISTSTREAMSOURCE_CI_TEST_STREAM             ListStreamSource = "ci_test_stream"
	LISTSTREAMSOURCE_RUM_ISSUE_STREAM           ListStreamSource = "rum_issue_stream"
	LISTSTREAMSOURCE_APM_ISSUE_STREAM           ListStreamSource = "apm_issue_stream"
	LISTSTREAMSOURCE_TRACE_STREAM               ListStreamSource = "trace_stream"
	LISTSTREAMSOURCE_LOGS_ISSUE_STREAM          ListStreamSource = "logs_issue_stream"
	LISTSTREAMSOURCE_LOGS_PATTERN_STREAM        ListStreamSource = "logs_pattern_stream"
	LISTSTREAMSOURCE_LOGS_TRANSACTION_STREAM    ListStreamSource = "logs_transaction_stream"
	LISTSTREAMSOURCE_EVENT_STREAM               ListStreamSource = "event_stream"
	LISTSTREAMSOURCE_RUM_STREAM                 ListStreamSource = "rum_stream"
	LISTSTREAMSOURCE_LLM_OBSERVABILITY_STREAM   ListStreamSource = "llm_observability_stream"
	LISTSTREAMSOURCE_ISSUE_STREAM               ListStreamSource = "issue_stream"
	LISTSTREAMSOURCE_SECURITY_RUNTIME_STREAM    ListStreamSource = "security_runtime_stream"
	LISTSTREAMSOURCE_SECURITY_SIGNALS_STREAM    ListStreamSource = "security_signals_stream"
	LISTSTREAMSOURCE_INCIDENTS_STREAM           ListStreamSource = "incidents_stream"
	LISTSTREAMSOURCE_APM_RECOMMENDATIONS_STREAM ListStreamSource = "apm_recommendations_stream"
)

var allowedListStreamSourceEnumValues = []ListStreamSource{
	LISTSTREAMSOURCE_LOGS_STREAM,
	LISTSTREAMSOURCE_AUDIT_STREAM,
	LISTSTREAMSOURCE_CI_PIPELINE_STREAM,
	LISTSTREAMSOURCE_CI_TEST_STREAM,
	LISTSTREAMSOURCE_RUM_ISSUE_STREAM,
	LISTSTREAMSOURCE_APM_ISSUE_STREAM,
	LISTSTREAMSOURCE_TRACE_STREAM,
	LISTSTREAMSOURCE_LOGS_ISSUE_STREAM,
	LISTSTREAMSOURCE_LOGS_PATTERN_STREAM,
	LISTSTREAMSOURCE_LOGS_TRANSACTION_STREAM,
	LISTSTREAMSOURCE_EVENT_STREAM,
	LISTSTREAMSOURCE_RUM_STREAM,
	LISTSTREAMSOURCE_LLM_OBSERVABILITY_STREAM,
	LISTSTREAMSOURCE_ISSUE_STREAM,
	LISTSTREAMSOURCE_SECURITY_RUNTIME_STREAM,
	LISTSTREAMSOURCE_SECURITY_SIGNALS_STREAM,
	LISTSTREAMSOURCE_INCIDENTS_STREAM,
	LISTSTREAMSOURCE_APM_RECOMMENDATIONS_STREAM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ListStreamSource) GetAllowedValues() []ListStreamSource {
	return allowedListStreamSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ListStreamSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ListStreamSource(value)
	return nil
}

// NewListStreamSourceFromValue returns a pointer to a valid ListStreamSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewListStreamSourceFromValue(v string) (*ListStreamSource, error) {
	ev := ListStreamSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ListStreamSource: valid values are %v", v, allowedListStreamSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ListStreamSource) IsValid() bool {
	for _, existing := range allowedListStreamSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ListStreamSource value.
func (v ListStreamSource) Ptr() *ListStreamSource {
	return &v
}
