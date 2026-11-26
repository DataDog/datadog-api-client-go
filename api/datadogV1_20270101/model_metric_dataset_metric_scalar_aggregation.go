// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MetricDatasetMetricScalarAggregation Test-only generated schema for the dashboard versioning proof of concept.
type MetricDatasetMetricScalarAggregation string

// List of MetricDatasetMetricScalarAggregation.
const (
	METRICDATASETMETRICSCALARAGGREGATION_SUM MetricDatasetMetricScalarAggregation = "sum"
	METRICDATASETMETRICSCALARAGGREGATION_MIN MetricDatasetMetricScalarAggregation = "min"
	METRICDATASETMETRICSCALARAGGREGATION_MAX MetricDatasetMetricScalarAggregation = "max"
	METRICDATASETMETRICSCALARAGGREGATION_AVG MetricDatasetMetricScalarAggregation = "avg"
)

var allowedMetricDatasetMetricScalarAggregationEnumValues = []MetricDatasetMetricScalarAggregation{
	METRICDATASETMETRICSCALARAGGREGATION_SUM,
	METRICDATASETMETRICSCALARAGGREGATION_MIN,
	METRICDATASETMETRICSCALARAGGREGATION_MAX,
	METRICDATASETMETRICSCALARAGGREGATION_AVG,
}

// GetAllowedValues reeturns the list of possible values.
func (v *MetricDatasetMetricScalarAggregation) GetAllowedValues() []MetricDatasetMetricScalarAggregation {
	return allowedMetricDatasetMetricScalarAggregationEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *MetricDatasetMetricScalarAggregation) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MetricDatasetMetricScalarAggregation(value)
	return nil
}

// NewMetricDatasetMetricScalarAggregationFromValue returns a pointer to a valid MetricDatasetMetricScalarAggregation
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewMetricDatasetMetricScalarAggregationFromValue(v string) (*MetricDatasetMetricScalarAggregation, error) {
	ev := MetricDatasetMetricScalarAggregation(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for MetricDatasetMetricScalarAggregation: valid values are %v", v, allowedMetricDatasetMetricScalarAggregationEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v MetricDatasetMetricScalarAggregation) IsValid() bool {
	for _, existing := range allowedMetricDatasetMetricScalarAggregationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MetricDatasetMetricScalarAggregation value.
func (v MetricDatasetMetricScalarAggregation) Ptr() *MetricDatasetMetricScalarAggregation {
	return &v
}
