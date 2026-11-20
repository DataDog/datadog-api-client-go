// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaMetricsQueryDatadogUsageDataSource Test-only generated schema for the dashboard versioning proof of concept.
type FormulaMetricsQueryDatadogUsageDataSource string

// List of FormulaMetricsQueryDatadogUsageDataSource.
const (
	FORMULAMETRICSQUERYDATADOGUSAGEDATASOURCE_DATADOG_USAGE FormulaMetricsQueryDatadogUsageDataSource = "datadog_usage"
)

var allowedFormulaMetricsQueryDatadogUsageDataSourceEnumValues = []FormulaMetricsQueryDatadogUsageDataSource{
	FORMULAMETRICSQUERYDATADOGUSAGEDATASOURCE_DATADOG_USAGE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FormulaMetricsQueryDatadogUsageDataSource) GetAllowedValues() []FormulaMetricsQueryDatadogUsageDataSource {
	return allowedFormulaMetricsQueryDatadogUsageDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FormulaMetricsQueryDatadogUsageDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FormulaMetricsQueryDatadogUsageDataSource(value)
	return nil
}

// NewFormulaMetricsQueryDatadogUsageDataSourceFromValue returns a pointer to a valid FormulaMetricsQueryDatadogUsageDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFormulaMetricsQueryDatadogUsageDataSourceFromValue(v string) (*FormulaMetricsQueryDatadogUsageDataSource, error) {
	ev := FormulaMetricsQueryDatadogUsageDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FormulaMetricsQueryDatadogUsageDataSource: valid values are %v", v, allowedFormulaMetricsQueryDatadogUsageDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FormulaMetricsQueryDatadogUsageDataSource) IsValid() bool {
	for _, existing := range allowedFormulaMetricsQueryDatadogUsageDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FormulaMetricsQueryDatadogUsageDataSource value.
func (v FormulaMetricsQueryDatadogUsageDataSource) Ptr() *FormulaMetricsQueryDatadogUsageDataSource {
	return &v
}
