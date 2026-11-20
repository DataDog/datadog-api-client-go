// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaMetricsDataSource Test-only generated schema for the dashboard versioning proof of concept.
type FormulaMetricsDataSource string

// List of FormulaMetricsDataSource.
const (
	FORMULAMETRICSDATASOURCE_METRICS       FormulaMetricsDataSource = "metrics"
	FORMULAMETRICSDATASOURCE_CLOUD_COST    FormulaMetricsDataSource = "cloud_cost"
	FORMULAMETRICSDATASOURCE_DATADOG_USAGE FormulaMetricsDataSource = "datadog_usage"
)

var allowedFormulaMetricsDataSourceEnumValues = []FormulaMetricsDataSource{
	FORMULAMETRICSDATASOURCE_METRICS,
	FORMULAMETRICSDATASOURCE_CLOUD_COST,
	FORMULAMETRICSDATASOURCE_DATADOG_USAGE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FormulaMetricsDataSource) GetAllowedValues() []FormulaMetricsDataSource {
	return allowedFormulaMetricsDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FormulaMetricsDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FormulaMetricsDataSource(value)
	return nil
}

// NewFormulaMetricsDataSourceFromValue returns a pointer to a valid FormulaMetricsDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFormulaMetricsDataSourceFromValue(v string) (*FormulaMetricsDataSource, error) {
	ev := FormulaMetricsDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FormulaMetricsDataSource: valid values are %v", v, allowedFormulaMetricsDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FormulaMetricsDataSource) IsValid() bool {
	for _, existing := range allowedFormulaMetricsDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FormulaMetricsDataSource value.
func (v FormulaMetricsDataSource) Ptr() *FormulaMetricsDataSource {
	return &v
}
