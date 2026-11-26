// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaDatasetApmMetricsQueryDataSource Test-only generated schema for the dashboard versioning proof of concept.
type FormulaDatasetApmMetricsQueryDataSource string

// List of FormulaDatasetApmMetricsQueryDataSource.
const (
	FORMULADATASETAPMMETRICSQUERYDATASOURCE_APM_METRICS FormulaDatasetApmMetricsQueryDataSource = "apm_metrics"
)

var allowedFormulaDatasetApmMetricsQueryDataSourceEnumValues = []FormulaDatasetApmMetricsQueryDataSource{
	FORMULADATASETAPMMETRICSQUERYDATASOURCE_APM_METRICS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FormulaDatasetApmMetricsQueryDataSource) GetAllowedValues() []FormulaDatasetApmMetricsQueryDataSource {
	return allowedFormulaDatasetApmMetricsQueryDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FormulaDatasetApmMetricsQueryDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FormulaDatasetApmMetricsQueryDataSource(value)
	return nil
}

// NewFormulaDatasetApmMetricsQueryDataSourceFromValue returns a pointer to a valid FormulaDatasetApmMetricsQueryDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFormulaDatasetApmMetricsQueryDataSourceFromValue(v string) (*FormulaDatasetApmMetricsQueryDataSource, error) {
	ev := FormulaDatasetApmMetricsQueryDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FormulaDatasetApmMetricsQueryDataSource: valid values are %v", v, allowedFormulaDatasetApmMetricsQueryDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FormulaDatasetApmMetricsQueryDataSource) IsValid() bool {
	for _, existing := range allowedFormulaDatasetApmMetricsQueryDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FormulaDatasetApmMetricsQueryDataSource value.
func (v FormulaDatasetApmMetricsQueryDataSource) Ptr() *FormulaDatasetApmMetricsQueryDataSource {
	return &v
}
