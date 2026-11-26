// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HeatgridLabelColumnWidth Test-only generated schema for the dashboard versioning proof of concept.
type HeatgridLabelColumnWidth string

// List of HeatgridLabelColumnWidth.
const (
	HEATGRIDLABELCOLUMNWIDTH_S  HeatgridLabelColumnWidth = "s"
	HEATGRIDLABELCOLUMNWIDTH_L  HeatgridLabelColumnWidth = "l"
	HEATGRIDLABELCOLUMNWIDTH_XS HeatgridLabelColumnWidth = "xs"
	HEATGRIDLABELCOLUMNWIDTH_M  HeatgridLabelColumnWidth = "m"
	HEATGRIDLABELCOLUMNWIDTH_XL HeatgridLabelColumnWidth = "xl"
)

var allowedHeatgridLabelColumnWidthEnumValues = []HeatgridLabelColumnWidth{
	HEATGRIDLABELCOLUMNWIDTH_S,
	HEATGRIDLABELCOLUMNWIDTH_L,
	HEATGRIDLABELCOLUMNWIDTH_XS,
	HEATGRIDLABELCOLUMNWIDTH_M,
	HEATGRIDLABELCOLUMNWIDTH_XL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *HeatgridLabelColumnWidth) GetAllowedValues() []HeatgridLabelColumnWidth {
	return allowedHeatgridLabelColumnWidthEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *HeatgridLabelColumnWidth) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = HeatgridLabelColumnWidth(value)
	return nil
}

// NewHeatgridLabelColumnWidthFromValue returns a pointer to a valid HeatgridLabelColumnWidth
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewHeatgridLabelColumnWidthFromValue(v string) (*HeatgridLabelColumnWidth, error) {
	ev := HeatgridLabelColumnWidth(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for HeatgridLabelColumnWidth: valid values are %v", v, allowedHeatgridLabelColumnWidthEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v HeatgridLabelColumnWidth) IsValid() bool {
	for _, existing := range allowedHeatgridLabelColumnWidthEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HeatgridLabelColumnWidth value.
func (v HeatgridLabelColumnWidth) Ptr() *HeatgridLabelColumnWidth {
	return &v
}
