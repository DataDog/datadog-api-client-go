// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PointPlotBackendDefinitionType Test-only generated schema for the dashboard versioning proof of concept.
type PointPlotBackendDefinitionType string

// List of PointPlotBackendDefinitionType.
const (
	POINTPLOTBACKENDDEFINITIONTYPE_POINT_PLOT PointPlotBackendDefinitionType = "point_plot"
)

var allowedPointPlotBackendDefinitionTypeEnumValues = []PointPlotBackendDefinitionType{
	POINTPLOTBACKENDDEFINITIONTYPE_POINT_PLOT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *PointPlotBackendDefinitionType) GetAllowedValues() []PointPlotBackendDefinitionType {
	return allowedPointPlotBackendDefinitionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *PointPlotBackendDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = PointPlotBackendDefinitionType(value)
	return nil
}

// NewPointPlotBackendDefinitionTypeFromValue returns a pointer to a valid PointPlotBackendDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewPointPlotBackendDefinitionTypeFromValue(v string) (*PointPlotBackendDefinitionType, error) {
	ev := PointPlotBackendDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for PointPlotBackendDefinitionType: valid values are %v", v, allowedPointPlotBackendDefinitionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v PointPlotBackendDefinitionType) IsValid() bool {
	for _, existing := range allowedPointPlotBackendDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PointPlotBackendDefinitionType value.
func (v PointPlotBackendDefinitionType) Ptr() *PointPlotBackendDefinitionType {
	return &v
}
