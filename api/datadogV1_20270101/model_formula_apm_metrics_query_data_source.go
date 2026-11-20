// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaApmMetricsQueryDataSource Test-only generated schema for the dashboard versioning proof of concept.
type FormulaApmMetricsQueryDataSource string

// List of FormulaApmMetricsQueryDataSource.
const (
	FORMULAAPMMETRICSQUERYDATASOURCE_APM_METRICS FormulaApmMetricsQueryDataSource = "apm_metrics"
)

var allowedFormulaApmMetricsQueryDataSourceEnumValues = []FormulaApmMetricsQueryDataSource{
	FORMULAAPMMETRICSQUERYDATASOURCE_APM_METRICS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FormulaApmMetricsQueryDataSource) GetAllowedValues() []FormulaApmMetricsQueryDataSource {
	return allowedFormulaApmMetricsQueryDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FormulaApmMetricsQueryDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FormulaApmMetricsQueryDataSource(value)
	return nil
}

// NewFormulaApmMetricsQueryDataSourceFromValue returns a pointer to a valid FormulaApmMetricsQueryDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFormulaApmMetricsQueryDataSourceFromValue(v string) (*FormulaApmMetricsQueryDataSource, error) {
	ev := FormulaApmMetricsQueryDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FormulaApmMetricsQueryDataSource: valid values are %v", v, allowedFormulaApmMetricsQueryDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FormulaApmMetricsQueryDataSource) IsValid() bool {
	for _, existing := range allowedFormulaApmMetricsQueryDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FormulaApmMetricsQueryDataSource value.
func (v FormulaApmMetricsQueryDataSource) Ptr() *FormulaApmMetricsQueryDataSource {
	return &v
}
