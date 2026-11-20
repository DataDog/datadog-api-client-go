// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ApmMetricStat Test-only generated schema for the dashboard versioning proof of concept.
type ApmMetricStat string

// List of ApmMetricStat.
const (
	APMMETRICSTAT_ERRORS               ApmMetricStat = "errors"
	APMMETRICSTAT_ERROR_RATE           ApmMetricStat = "error_rate"
	APMMETRICSTAT_ERRORS_PER_SECOND    ApmMetricStat = "errors_per_second"
	APMMETRICSTAT_LATENCY_AVG          ApmMetricStat = "latency_avg"
	APMMETRICSTAT_LATENCY_MAX          ApmMetricStat = "latency_max"
	APMMETRICSTAT_LATENCY_P50          ApmMetricStat = "latency_p50"
	APMMETRICSTAT_LATENCY_P75          ApmMetricStat = "latency_p75"
	APMMETRICSTAT_LATENCY_P90          ApmMetricStat = "latency_p90"
	APMMETRICSTAT_LATENCY_P95          ApmMetricStat = "latency_p95"
	APMMETRICSTAT_LATENCY_P99          ApmMetricStat = "latency_p99"
	APMMETRICSTAT_LATENCY_P999         ApmMetricStat = "latency_p999"
	APMMETRICSTAT_LATENCY_DISTRIBUTION ApmMetricStat = "latency_distribution"
	APMMETRICSTAT_HITS                 ApmMetricStat = "hits"
	APMMETRICSTAT_HITS_PER_SECOND      ApmMetricStat = "hits_per_second"
	APMMETRICSTAT_TOTAL_TIME           ApmMetricStat = "total_time"
	APMMETRICSTAT_APDEX                ApmMetricStat = "apdex"
)

var allowedApmMetricStatEnumValues = []ApmMetricStat{
	APMMETRICSTAT_ERRORS,
	APMMETRICSTAT_ERROR_RATE,
	APMMETRICSTAT_ERRORS_PER_SECOND,
	APMMETRICSTAT_LATENCY_AVG,
	APMMETRICSTAT_LATENCY_MAX,
	APMMETRICSTAT_LATENCY_P50,
	APMMETRICSTAT_LATENCY_P75,
	APMMETRICSTAT_LATENCY_P90,
	APMMETRICSTAT_LATENCY_P95,
	APMMETRICSTAT_LATENCY_P99,
	APMMETRICSTAT_LATENCY_P999,
	APMMETRICSTAT_LATENCY_DISTRIBUTION,
	APMMETRICSTAT_HITS,
	APMMETRICSTAT_HITS_PER_SECOND,
	APMMETRICSTAT_TOTAL_TIME,
	APMMETRICSTAT_APDEX,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ApmMetricStat) GetAllowedValues() []ApmMetricStat {
	return allowedApmMetricStatEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ApmMetricStat) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ApmMetricStat(value)
	return nil
}

// NewApmMetricStatFromValue returns a pointer to a valid ApmMetricStat
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewApmMetricStatFromValue(v string) (*ApmMetricStat, error) {
	ev := ApmMetricStat(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ApmMetricStat: valid values are %v", v, allowedApmMetricStatEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ApmMetricStat) IsValid() bool {
	for _, existing := range allowedApmMetricStatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ApmMetricStat value.
func (v ApmMetricStat) Ptr() *ApmMetricStat {
	return &v
}
