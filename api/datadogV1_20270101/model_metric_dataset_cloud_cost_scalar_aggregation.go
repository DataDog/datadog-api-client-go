// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MetricDatasetCloudCostScalarAggregation Test-only generated schema for the dashboard versioning proof of concept.
type MetricDatasetCloudCostScalarAggregation string

// List of MetricDatasetCloudCostScalarAggregation.
const (
	METRICDATASETCLOUDCOSTSCALARAGGREGATION_SUM MetricDatasetCloudCostScalarAggregation = "sum"
	METRICDATASETCLOUDCOSTSCALARAGGREGATION_MIN MetricDatasetCloudCostScalarAggregation = "min"
	METRICDATASETCLOUDCOSTSCALARAGGREGATION_MAX MetricDatasetCloudCostScalarAggregation = "max"
	METRICDATASETCLOUDCOSTSCALARAGGREGATION_AVG MetricDatasetCloudCostScalarAggregation = "avg"
)

var allowedMetricDatasetCloudCostScalarAggregationEnumValues = []MetricDatasetCloudCostScalarAggregation{
	METRICDATASETCLOUDCOSTSCALARAGGREGATION_SUM,
	METRICDATASETCLOUDCOSTSCALARAGGREGATION_MIN,
	METRICDATASETCLOUDCOSTSCALARAGGREGATION_MAX,
	METRICDATASETCLOUDCOSTSCALARAGGREGATION_AVG,
}

// GetAllowedValues reeturns the list of possible values.
func (v *MetricDatasetCloudCostScalarAggregation) GetAllowedValues() []MetricDatasetCloudCostScalarAggregation {
	return allowedMetricDatasetCloudCostScalarAggregationEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *MetricDatasetCloudCostScalarAggregation) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MetricDatasetCloudCostScalarAggregation(value)
	return nil
}

// NewMetricDatasetCloudCostScalarAggregationFromValue returns a pointer to a valid MetricDatasetCloudCostScalarAggregation
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewMetricDatasetCloudCostScalarAggregationFromValue(v string) (*MetricDatasetCloudCostScalarAggregation, error) {
	ev := MetricDatasetCloudCostScalarAggregation(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for MetricDatasetCloudCostScalarAggregation: valid values are %v", v, allowedMetricDatasetCloudCostScalarAggregationEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v MetricDatasetCloudCostScalarAggregation) IsValid() bool {
	for _, existing := range allowedMetricDatasetCloudCostScalarAggregationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MetricDatasetCloudCostScalarAggregation value.
func (v MetricDatasetCloudCostScalarAggregation) Ptr() *MetricDatasetCloudCostScalarAggregation {
	return &v
}
