// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HeatgridWidgetDefinitionType Type of the heatgrid widget.
type HeatgridWidgetDefinitionType string

// List of HeatgridWidgetDefinitionType.
const (
	HEATGRIDWIDGETDEFINITIONTYPE_HEATGRID HeatgridWidgetDefinitionType = "heatgrid"
)

var allowedHeatgridWidgetDefinitionTypeEnumValues = []HeatgridWidgetDefinitionType{
	HEATGRIDWIDGETDEFINITIONTYPE_HEATGRID,
}

// GetAllowedValues reeturns the list of possible values.
func (v *HeatgridWidgetDefinitionType) GetAllowedValues() []HeatgridWidgetDefinitionType {
	return allowedHeatgridWidgetDefinitionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *HeatgridWidgetDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = HeatgridWidgetDefinitionType(value)
	return nil
}

// NewHeatgridWidgetDefinitionTypeFromValue returns a pointer to a valid HeatgridWidgetDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewHeatgridWidgetDefinitionTypeFromValue(v string) (*HeatgridWidgetDefinitionType, error) {
	ev := HeatgridWidgetDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for HeatgridWidgetDefinitionType: valid values are %v", v, allowedHeatgridWidgetDefinitionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v HeatgridWidgetDefinitionType) IsValid() bool {
	for _, existing := range allowedHeatgridWidgetDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HeatgridWidgetDefinitionType value.
func (v HeatgridWidgetDefinitionType) Ptr() *HeatgridWidgetDefinitionType {
	return &v
}
