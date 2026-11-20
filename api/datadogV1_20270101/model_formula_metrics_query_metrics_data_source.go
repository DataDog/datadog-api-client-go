// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaMetricsQueryMetricsDataSource Test-only generated schema for the dashboard versioning proof of concept.
type FormulaMetricsQueryMetricsDataSource string

// List of FormulaMetricsQueryMetricsDataSource.
const (
	FORMULAMETRICSQUERYMETRICSDATASOURCE_METRICS FormulaMetricsQueryMetricsDataSource = "metrics"
)

var allowedFormulaMetricsQueryMetricsDataSourceEnumValues = []FormulaMetricsQueryMetricsDataSource{
	FORMULAMETRICSQUERYMETRICSDATASOURCE_METRICS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FormulaMetricsQueryMetricsDataSource) GetAllowedValues() []FormulaMetricsQueryMetricsDataSource {
	return allowedFormulaMetricsQueryMetricsDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FormulaMetricsQueryMetricsDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FormulaMetricsQueryMetricsDataSource(value)
	return nil
}

// NewFormulaMetricsQueryMetricsDataSourceFromValue returns a pointer to a valid FormulaMetricsQueryMetricsDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFormulaMetricsQueryMetricsDataSourceFromValue(v string) (*FormulaMetricsQueryMetricsDataSource, error) {
	ev := FormulaMetricsQueryMetricsDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FormulaMetricsQueryMetricsDataSource: valid values are %v", v, allowedFormulaMetricsQueryMetricsDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FormulaMetricsQueryMetricsDataSource) IsValid() bool {
	for _, existing := range allowedFormulaMetricsQueryMetricsDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FormulaMetricsQueryMetricsDataSource value.
func (v FormulaMetricsQueryMetricsDataSource) Ptr() *FormulaMetricsQueryMetricsDataSource {
	return &v
}
