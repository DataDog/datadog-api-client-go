// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaDatasetMetricQueryDataSource Test-only generated schema for the dashboard versioning proof of concept.
type FormulaDatasetMetricQueryDataSource string

// List of FormulaDatasetMetricQueryDataSource.
const (
	FORMULADATASETMETRICQUERYDATASOURCE_METRICS    FormulaDatasetMetricQueryDataSource = "metrics"
	FORMULADATASETMETRICQUERYDATASOURCE_CLOUD_COST FormulaDatasetMetricQueryDataSource = "cloud_cost"
)

var allowedFormulaDatasetMetricQueryDataSourceEnumValues = []FormulaDatasetMetricQueryDataSource{
	FORMULADATASETMETRICQUERYDATASOURCE_METRICS,
	FORMULADATASETMETRICQUERYDATASOURCE_CLOUD_COST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FormulaDatasetMetricQueryDataSource) GetAllowedValues() []FormulaDatasetMetricQueryDataSource {
	return allowedFormulaDatasetMetricQueryDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FormulaDatasetMetricQueryDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FormulaDatasetMetricQueryDataSource(value)
	return nil
}

// NewFormulaDatasetMetricQueryDataSourceFromValue returns a pointer to a valid FormulaDatasetMetricQueryDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFormulaDatasetMetricQueryDataSourceFromValue(v string) (*FormulaDatasetMetricQueryDataSource, error) {
	ev := FormulaDatasetMetricQueryDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FormulaDatasetMetricQueryDataSource: valid values are %v", v, allowedFormulaDatasetMetricQueryDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FormulaDatasetMetricQueryDataSource) IsValid() bool {
	for _, existing := range allowedFormulaDatasetMetricQueryDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FormulaDatasetMetricQueryDataSource value.
func (v FormulaDatasetMetricQueryDataSource) Ptr() *FormulaDatasetMetricQueryDataSource {
	return &v
}
