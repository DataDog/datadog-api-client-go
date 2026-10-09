// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HeatgridSortByLabelProperty Sort by label.
type HeatgridSortByLabelProperty string

// List of HeatgridSortByLabelProperty.
const (
	HEATGRIDSORTBYLABELPROPERTY_LABEL HeatgridSortByLabelProperty = "label"
)

var allowedHeatgridSortByLabelPropertyEnumValues = []HeatgridSortByLabelProperty{
	HEATGRIDSORTBYLABELPROPERTY_LABEL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *HeatgridSortByLabelProperty) GetAllowedValues() []HeatgridSortByLabelProperty {
	return allowedHeatgridSortByLabelPropertyEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *HeatgridSortByLabelProperty) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = HeatgridSortByLabelProperty(value)
	return nil
}

// NewHeatgridSortByLabelPropertyFromValue returns a pointer to a valid HeatgridSortByLabelProperty
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewHeatgridSortByLabelPropertyFromValue(v string) (*HeatgridSortByLabelProperty, error) {
	ev := HeatgridSortByLabelProperty(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for HeatgridSortByLabelProperty: valid values are %v", v, allowedHeatgridSortByLabelPropertyEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v HeatgridSortByLabelProperty) IsValid() bool {
	for _, existing := range allowedHeatgridSortByLabelPropertyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HeatgridSortByLabelProperty value.
func (v HeatgridSortByLabelProperty) Ptr() *HeatgridSortByLabelProperty {
	return &v
}
