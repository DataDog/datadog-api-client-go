// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CaseMetric Test-only generated schema for the dashboard versioning proof of concept.
type CaseMetric string

// List of CaseMetric.
const (
	CASEMETRIC_TIME_TO_CLOSE  CaseMetric = "time_to_close"
	CASEMETRIC_TIME_TO_ASSIGN CaseMetric = "time_to_assign"
	CASEMETRIC_EVENT_COUNT    CaseMetric = "event_count"
)

var allowedCaseMetricEnumValues = []CaseMetric{
	CASEMETRIC_TIME_TO_CLOSE,
	CASEMETRIC_TIME_TO_ASSIGN,
	CASEMETRIC_EVENT_COUNT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CaseMetric) GetAllowedValues() []CaseMetric {
	return allowedCaseMetricEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CaseMetric) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CaseMetric(value)
	return nil
}

// NewCaseMetricFromValue returns a pointer to a valid CaseMetric
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCaseMetricFromValue(v string) (*CaseMetric, error) {
	ev := CaseMetric(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CaseMetric: valid values are %v", v, allowedCaseMetricEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CaseMetric) IsValid() bool {
	for _, existing := range allowedCaseMetricEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CaseMetric value.
func (v CaseMetric) Ptr() *CaseMetric {
	return &v
}
