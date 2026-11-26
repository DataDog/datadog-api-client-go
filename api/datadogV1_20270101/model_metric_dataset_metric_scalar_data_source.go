// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MetricDatasetMetricScalarDataSource Test-only generated schema for the dashboard versioning proof of concept.
type MetricDatasetMetricScalarDataSource string

// List of MetricDatasetMetricScalarDataSource.
const (
	METRICDATASETMETRICSCALARDATASOURCE_METRIC_SCALAR MetricDatasetMetricScalarDataSource = "metric_scalar"
)

var allowedMetricDatasetMetricScalarDataSourceEnumValues = []MetricDatasetMetricScalarDataSource{
	METRICDATASETMETRICSCALARDATASOURCE_METRIC_SCALAR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *MetricDatasetMetricScalarDataSource) GetAllowedValues() []MetricDatasetMetricScalarDataSource {
	return allowedMetricDatasetMetricScalarDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *MetricDatasetMetricScalarDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MetricDatasetMetricScalarDataSource(value)
	return nil
}

// NewMetricDatasetMetricScalarDataSourceFromValue returns a pointer to a valid MetricDatasetMetricScalarDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewMetricDatasetMetricScalarDataSourceFromValue(v string) (*MetricDatasetMetricScalarDataSource, error) {
	ev := MetricDatasetMetricScalarDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for MetricDatasetMetricScalarDataSource: valid values are %v", v, allowedMetricDatasetMetricScalarDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v MetricDatasetMetricScalarDataSource) IsValid() bool {
	for _, existing := range allowedMetricDatasetMetricScalarDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MetricDatasetMetricScalarDataSource value.
func (v MetricDatasetMetricScalarDataSource) Ptr() *MetricDatasetMetricScalarDataSource {
	return &v
}
