// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HeatGridSortAggregation Test-only generated schema for the dashboard versioning proof of concept.
type HeatGridSortAggregation string

// List of HeatGridSortAggregation.
const (
	HEATGRIDSORTAGGREGATION_SUM HeatGridSortAggregation = "sum"
	HEATGRIDSORTAGGREGATION_MIN HeatGridSortAggregation = "min"
	HEATGRIDSORTAGGREGATION_MAX HeatGridSortAggregation = "max"
	HEATGRIDSORTAGGREGATION_AVG HeatGridSortAggregation = "avg"
)

var allowedHeatGridSortAggregationEnumValues = []HeatGridSortAggregation{
	HEATGRIDSORTAGGREGATION_SUM,
	HEATGRIDSORTAGGREGATION_MIN,
	HEATGRIDSORTAGGREGATION_MAX,
	HEATGRIDSORTAGGREGATION_AVG,
}

// GetAllowedValues reeturns the list of possible values.
func (v *HeatGridSortAggregation) GetAllowedValues() []HeatGridSortAggregation {
	return allowedHeatGridSortAggregationEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *HeatGridSortAggregation) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = HeatGridSortAggregation(value)
	return nil
}

// NewHeatGridSortAggregationFromValue returns a pointer to a valid HeatGridSortAggregation
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewHeatGridSortAggregationFromValue(v string) (*HeatGridSortAggregation, error) {
	ev := HeatGridSortAggregation(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for HeatGridSortAggregation: valid values are %v", v, allowedHeatGridSortAggregationEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v HeatGridSortAggregation) IsValid() bool {
	for _, existing := range allowedHeatGridSortAggregationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HeatGridSortAggregation value.
func (v HeatGridSortAggregation) Ptr() *HeatGridSortAggregation {
	return &v
}
