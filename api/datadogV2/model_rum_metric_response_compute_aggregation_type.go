// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumMetricResponseComputeAggregationType The type of aggregation to use.
type RumMetricResponseComputeAggregationType string

// List of RumMetricResponseComputeAggregationType.
const (
	RUMMETRICRESPONSECOMPUTEAGGREGATIONTYPE_COUNT        RumMetricResponseComputeAggregationType = "count"
	RUMMETRICRESPONSECOMPUTEAGGREGATIONTYPE_DISTRIBUTION RumMetricResponseComputeAggregationType = "distribution"
)

var allowedRumMetricResponseComputeAggregationTypeEnumValues = []RumMetricResponseComputeAggregationType{
	RUMMETRICRESPONSECOMPUTEAGGREGATIONTYPE_COUNT,
	RUMMETRICRESPONSECOMPUTEAGGREGATIONTYPE_DISTRIBUTION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RumMetricResponseComputeAggregationType) GetAllowedValues() []RumMetricResponseComputeAggregationType {
	return allowedRumMetricResponseComputeAggregationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RumMetricResponseComputeAggregationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RumMetricResponseComputeAggregationType(value)
	return nil
}

// NewRumMetricResponseComputeAggregationTypeFromValue returns a pointer to a valid RumMetricResponseComputeAggregationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRumMetricResponseComputeAggregationTypeFromValue(v string) (*RumMetricResponseComputeAggregationType, error) {
	ev := RumMetricResponseComputeAggregationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RumMetricResponseComputeAggregationType: valid values are %v", v, allowedRumMetricResponseComputeAggregationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RumMetricResponseComputeAggregationType) IsValid() bool {
	for _, existing := range allowedRumMetricResponseComputeAggregationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RumMetricResponseComputeAggregationType value.
func (v RumMetricResponseComputeAggregationType) Ptr() *RumMetricResponseComputeAggregationType {
	return &v
}
