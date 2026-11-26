// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ApmResourceStat Test-only generated schema for the dashboard versioning proof of concept.
type ApmResourceStat string

// List of ApmResourceStat.
const (
	APMRESOURCESTAT_ERRORS               ApmResourceStat = "errors"
	APMRESOURCESTAT_ERROR_RATE           ApmResourceStat = "error_rate"
	APMRESOURCESTAT_LATENCY_AVG          ApmResourceStat = "latency_avg"
	APMRESOURCESTAT_LATENCY_MAX          ApmResourceStat = "latency_max"
	APMRESOURCESTAT_LATENCY_P50          ApmResourceStat = "latency_p50"
	APMRESOURCESTAT_LATENCY_P75          ApmResourceStat = "latency_p75"
	APMRESOURCESTAT_LATENCY_P90          ApmResourceStat = "latency_p90"
	APMRESOURCESTAT_LATENCY_P95          ApmResourceStat = "latency_p95"
	APMRESOURCESTAT_LATENCY_P99          ApmResourceStat = "latency_p99"
	APMRESOURCESTAT_LATENCY_DISTRIBUTION ApmResourceStat = "latency_distribution"
	APMRESOURCESTAT_HITS                 ApmResourceStat = "hits"
	APMRESOURCESTAT_TOTAL_TIME           ApmResourceStat = "total_time"
)

var allowedApmResourceStatEnumValues = []ApmResourceStat{
	APMRESOURCESTAT_ERRORS,
	APMRESOURCESTAT_ERROR_RATE,
	APMRESOURCESTAT_LATENCY_AVG,
	APMRESOURCESTAT_LATENCY_MAX,
	APMRESOURCESTAT_LATENCY_P50,
	APMRESOURCESTAT_LATENCY_P75,
	APMRESOURCESTAT_LATENCY_P90,
	APMRESOURCESTAT_LATENCY_P95,
	APMRESOURCESTAT_LATENCY_P99,
	APMRESOURCESTAT_LATENCY_DISTRIBUTION,
	APMRESOURCESTAT_HITS,
	APMRESOURCESTAT_TOTAL_TIME,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ApmResourceStat) GetAllowedValues() []ApmResourceStat {
	return allowedApmResourceStatEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ApmResourceStat) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ApmResourceStat(value)
	return nil
}

// NewApmResourceStatFromValue returns a pointer to a valid ApmResourceStat
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewApmResourceStatFromValue(v string) (*ApmResourceStat, error) {
	ev := ApmResourceStat(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ApmResourceStat: valid values are %v", v, allowedApmResourceStatEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ApmResourceStat) IsValid() bool {
	for _, existing := range allowedApmResourceStatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ApmResourceStat value.
func (v ApmResourceStat) Ptr() *ApmResourceStat {
	return &v
}
