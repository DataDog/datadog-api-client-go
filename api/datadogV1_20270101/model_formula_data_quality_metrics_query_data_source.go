// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaDataQualityMetricsQueryDataSource Test-only generated schema for the dashboard versioning proof of concept.
type FormulaDataQualityMetricsQueryDataSource string

// List of FormulaDataQualityMetricsQueryDataSource.
const (
	FORMULADATAQUALITYMETRICSQUERYDATASOURCE_DATA_QUALITY_METRICS FormulaDataQualityMetricsQueryDataSource = "data_quality_metrics"
)

var allowedFormulaDataQualityMetricsQueryDataSourceEnumValues = []FormulaDataQualityMetricsQueryDataSource{
	FORMULADATAQUALITYMETRICSQUERYDATASOURCE_DATA_QUALITY_METRICS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FormulaDataQualityMetricsQueryDataSource) GetAllowedValues() []FormulaDataQualityMetricsQueryDataSource {
	return allowedFormulaDataQualityMetricsQueryDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FormulaDataQualityMetricsQueryDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FormulaDataQualityMetricsQueryDataSource(value)
	return nil
}

// NewFormulaDataQualityMetricsQueryDataSourceFromValue returns a pointer to a valid FormulaDataQualityMetricsQueryDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFormulaDataQualityMetricsQueryDataSourceFromValue(v string) (*FormulaDataQualityMetricsQueryDataSource, error) {
	ev := FormulaDataQualityMetricsQueryDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FormulaDataQualityMetricsQueryDataSource: valid values are %v", v, allowedFormulaDataQualityMetricsQueryDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FormulaDataQualityMetricsQueryDataSource) IsValid() bool {
	for _, existing := range allowedFormulaDataQualityMetricsQueryDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FormulaDataQualityMetricsQueryDataSource value.
func (v FormulaDataQualityMetricsQueryDataSource) Ptr() *FormulaDataQualityMetricsQueryDataSource {
	return &v
}
