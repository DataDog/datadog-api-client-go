// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaMetricsQueryCloudCostDataSource Test-only generated schema for the dashboard versioning proof of concept.
type FormulaMetricsQueryCloudCostDataSource string

// List of FormulaMetricsQueryCloudCostDataSource.
const (
	FORMULAMETRICSQUERYCLOUDCOSTDATASOURCE_CLOUD_COST FormulaMetricsQueryCloudCostDataSource = "cloud_cost"
)

var allowedFormulaMetricsQueryCloudCostDataSourceEnumValues = []FormulaMetricsQueryCloudCostDataSource{
	FORMULAMETRICSQUERYCLOUDCOSTDATASOURCE_CLOUD_COST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FormulaMetricsQueryCloudCostDataSource) GetAllowedValues() []FormulaMetricsQueryCloudCostDataSource {
	return allowedFormulaMetricsQueryCloudCostDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FormulaMetricsQueryCloudCostDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FormulaMetricsQueryCloudCostDataSource(value)
	return nil
}

// NewFormulaMetricsQueryCloudCostDataSourceFromValue returns a pointer to a valid FormulaMetricsQueryCloudCostDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFormulaMetricsQueryCloudCostDataSourceFromValue(v string) (*FormulaMetricsQueryCloudCostDataSource, error) {
	ev := FormulaMetricsQueryCloudCostDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FormulaMetricsQueryCloudCostDataSource: valid values are %v", v, allowedFormulaMetricsQueryCloudCostDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FormulaMetricsQueryCloudCostDataSource) IsValid() bool {
	for _, existing := range allowedFormulaMetricsQueryCloudCostDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FormulaMetricsQueryCloudCostDataSource value.
func (v FormulaMetricsQueryCloudCostDataSource) Ptr() *FormulaMetricsQueryCloudCostDataSource {
	return &v
}
