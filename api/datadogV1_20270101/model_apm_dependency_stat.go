// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ApmDependencyStat Test-only generated schema for the dashboard versioning proof of concept.
type ApmDependencyStat string

// List of ApmDependencyStat.
const (
	APMDEPENDENCYSTAT_ERROR_RATE          ApmDependencyStat = "error_rate"
	APMDEPENDENCYSTAT_AVG_DURATION        ApmDependencyStat = "avg_duration"
	APMDEPENDENCYSTAT_AVG_ROOT_DURATION   ApmDependencyStat = "avg_root_duration"
	APMDEPENDENCYSTAT_AVG_SPANS_PER_TRACE ApmDependencyStat = "avg_spans_per_trace"
	APMDEPENDENCYSTAT_PCT_EXEC_TIME       ApmDependencyStat = "pct_exec_time"
	APMDEPENDENCYSTAT_PCT_OF_TRACES       ApmDependencyStat = "pct_of_traces"
	APMDEPENDENCYSTAT_TOTAL_TRACES_COUNT  ApmDependencyStat = "total_traces_count"
)

var allowedApmDependencyStatEnumValues = []ApmDependencyStat{
	APMDEPENDENCYSTAT_ERROR_RATE,
	APMDEPENDENCYSTAT_AVG_DURATION,
	APMDEPENDENCYSTAT_AVG_ROOT_DURATION,
	APMDEPENDENCYSTAT_AVG_SPANS_PER_TRACE,
	APMDEPENDENCYSTAT_PCT_EXEC_TIME,
	APMDEPENDENCYSTAT_PCT_OF_TRACES,
	APMDEPENDENCYSTAT_TOTAL_TRACES_COUNT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ApmDependencyStat) GetAllowedValues() []ApmDependencyStat {
	return allowedApmDependencyStatEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ApmDependencyStat) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ApmDependencyStat(value)
	return nil
}

// NewApmDependencyStatFromValue returns a pointer to a valid ApmDependencyStat
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewApmDependencyStatFromValue(v string) (*ApmDependencyStat, error) {
	ev := ApmDependencyStat(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ApmDependencyStat: valid values are %v", v, allowedApmDependencyStatEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ApmDependencyStat) IsValid() bool {
	for _, existing := range allowedApmDependencyStatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ApmDependencyStat value.
func (v ApmDependencyStat) Ptr() *ApmDependencyStat {
	return &v
}
