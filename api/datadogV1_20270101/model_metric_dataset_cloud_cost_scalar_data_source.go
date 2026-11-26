// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MetricDatasetCloudCostScalarDataSource Test-only generated schema for the dashboard versioning proof of concept.
type MetricDatasetCloudCostScalarDataSource string

// List of MetricDatasetCloudCostScalarDataSource.
const (
	METRICDATASETCLOUDCOSTSCALARDATASOURCE_CLOUD_COST_SCALAR MetricDatasetCloudCostScalarDataSource = "cloud_cost_scalar"
)

var allowedMetricDatasetCloudCostScalarDataSourceEnumValues = []MetricDatasetCloudCostScalarDataSource{
	METRICDATASETCLOUDCOSTSCALARDATASOURCE_CLOUD_COST_SCALAR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *MetricDatasetCloudCostScalarDataSource) GetAllowedValues() []MetricDatasetCloudCostScalarDataSource {
	return allowedMetricDatasetCloudCostScalarDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *MetricDatasetCloudCostScalarDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MetricDatasetCloudCostScalarDataSource(value)
	return nil
}

// NewMetricDatasetCloudCostScalarDataSourceFromValue returns a pointer to a valid MetricDatasetCloudCostScalarDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewMetricDatasetCloudCostScalarDataSourceFromValue(v string) (*MetricDatasetCloudCostScalarDataSource, error) {
	ev := MetricDatasetCloudCostScalarDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for MetricDatasetCloudCostScalarDataSource: valid values are %v", v, allowedMetricDatasetCloudCostScalarDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v MetricDatasetCloudCostScalarDataSource) IsValid() bool {
	for _, existing := range allowedMetricDatasetCloudCostScalarDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MetricDatasetCloudCostScalarDataSource value.
func (v MetricDatasetCloudCostScalarDataSource) Ptr() *MetricDatasetCloudCostScalarDataSource {
	return &v
}
