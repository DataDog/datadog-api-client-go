// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HeatgridSortOrder Sort direction.
type HeatgridSortOrder string

// List of HeatgridSortOrder.
const (
	HEATGRIDSORTORDER_ASC  HeatgridSortOrder = "asc"
	HEATGRIDSORTORDER_DESC HeatgridSortOrder = "desc"
)

var allowedHeatgridSortOrderEnumValues = []HeatgridSortOrder{
	HEATGRIDSORTORDER_ASC,
	HEATGRIDSORTORDER_DESC,
}

// GetAllowedValues reeturns the list of possible values.
func (v *HeatgridSortOrder) GetAllowedValues() []HeatgridSortOrder {
	return allowedHeatgridSortOrderEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *HeatgridSortOrder) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = HeatgridSortOrder(value)
	return nil
}

// NewHeatgridSortOrderFromValue returns a pointer to a valid HeatgridSortOrder
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewHeatgridSortOrderFromValue(v string) (*HeatgridSortOrder, error) {
	ev := HeatgridSortOrder(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for HeatgridSortOrder: valid values are %v", v, allowedHeatgridSortOrderEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v HeatgridSortOrder) IsValid() bool {
	for _, existing := range allowedHeatgridSortOrderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HeatgridSortOrder value.
func (v HeatgridSortOrder) Ptr() *HeatgridSortOrder {
	return &v
}
