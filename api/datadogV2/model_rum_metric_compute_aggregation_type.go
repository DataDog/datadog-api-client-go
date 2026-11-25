// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMMetricComputeAggregationType The type of aggregation to use.
// `count` counts the matching events. `distribution` aggregates the values found at `path`.
type RUMMetricComputeAggregationType string

// List of RUMMetricComputeAggregationType.
const (
	RUMMETRICCOMPUTEAGGREGATIONTYPE_COUNT        RUMMetricComputeAggregationType = "count"
	RUMMETRICCOMPUTEAGGREGATIONTYPE_DISTRIBUTION RUMMetricComputeAggregationType = "distribution"
)

var allowedRUMMetricComputeAggregationTypeEnumValues = []RUMMetricComputeAggregationType{
	RUMMETRICCOMPUTEAGGREGATIONTYPE_COUNT,
	RUMMETRICCOMPUTEAGGREGATIONTYPE_DISTRIBUTION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RUMMetricComputeAggregationType) GetAllowedValues() []RUMMetricComputeAggregationType {
	return allowedRUMMetricComputeAggregationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RUMMetricComputeAggregationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RUMMetricComputeAggregationType(value)
	return nil
}

// NewRUMMetricComputeAggregationTypeFromValue returns a pointer to a valid RUMMetricComputeAggregationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRUMMetricComputeAggregationTypeFromValue(v string) (*RUMMetricComputeAggregationType, error) {
	ev := RUMMetricComputeAggregationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RUMMetricComputeAggregationType: valid values are %v", v, allowedRUMMetricComputeAggregationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RUMMetricComputeAggregationType) IsValid() bool {
	for _, existing := range allowedRUMMetricComputeAggregationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RUMMetricComputeAggregationType value.
func (v RUMMetricComputeAggregationType) Ptr() *RUMMetricComputeAggregationType {
	return &v
}
