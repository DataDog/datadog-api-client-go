// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HeatgridDefinitionType Test-only generated schema for the dashboard versioning proof of concept.
type HeatgridDefinitionType string

// List of HeatgridDefinitionType.
const (
	HEATGRIDDEFINITIONTYPE_HEATGRID HeatgridDefinitionType = "heatgrid"
)

var allowedHeatgridDefinitionTypeEnumValues = []HeatgridDefinitionType{
	HEATGRIDDEFINITIONTYPE_HEATGRID,
}

// GetAllowedValues reeturns the list of possible values.
func (v *HeatgridDefinitionType) GetAllowedValues() []HeatgridDefinitionType {
	return allowedHeatgridDefinitionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *HeatgridDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = HeatgridDefinitionType(value)
	return nil
}

// NewHeatgridDefinitionTypeFromValue returns a pointer to a valid HeatgridDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewHeatgridDefinitionTypeFromValue(v string) (*HeatgridDefinitionType, error) {
	ev := HeatgridDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for HeatgridDefinitionType: valid values are %v", v, allowedHeatgridDefinitionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v HeatgridDefinitionType) IsValid() bool {
	for _, existing := range allowedHeatgridDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HeatgridDefinitionType value.
func (v HeatgridDefinitionType) Ptr() *HeatgridDefinitionType {
	return &v
}
