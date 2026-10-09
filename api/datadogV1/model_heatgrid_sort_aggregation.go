// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HeatgridSortAggregation Aggregation used to order rows over the displayed time range.
type HeatgridSortAggregation string

// List of HeatgridSortAggregation.
const (
	HEATGRIDSORTAGGREGATION_AVG HeatgridSortAggregation = "avg"
	HEATGRIDSORTAGGREGATION_MIN HeatgridSortAggregation = "min"
	HEATGRIDSORTAGGREGATION_MAX HeatgridSortAggregation = "max"
	HEATGRIDSORTAGGREGATION_SUM HeatgridSortAggregation = "sum"
)

var allowedHeatgridSortAggregationEnumValues = []HeatgridSortAggregation{
	HEATGRIDSORTAGGREGATION_AVG,
	HEATGRIDSORTAGGREGATION_MIN,
	HEATGRIDSORTAGGREGATION_MAX,
	HEATGRIDSORTAGGREGATION_SUM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *HeatgridSortAggregation) GetAllowedValues() []HeatgridSortAggregation {
	return allowedHeatgridSortAggregationEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *HeatgridSortAggregation) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = HeatgridSortAggregation(value)
	return nil
}

// NewHeatgridSortAggregationFromValue returns a pointer to a valid HeatgridSortAggregation
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewHeatgridSortAggregationFromValue(v string) (*HeatgridSortAggregation, error) {
	ev := HeatgridSortAggregation(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for HeatgridSortAggregation: valid values are %v", v, allowedHeatgridSortAggregationEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v HeatgridSortAggregation) IsValid() bool {
	for _, existing := range allowedHeatgridSortAggregationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HeatgridSortAggregation value.
func (v HeatgridSortAggregation) Ptr() *HeatgridSortAggregation {
	return &v
}
