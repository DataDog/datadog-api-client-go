// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HeatmapBackendDefinitionType Test-only generated schema for the dashboard versioning proof of concept.
type HeatmapBackendDefinitionType string

// List of HeatmapBackendDefinitionType.
const (
	HEATMAPBACKENDDEFINITIONTYPE_HEATMAP HeatmapBackendDefinitionType = "heatmap"
)

var allowedHeatmapBackendDefinitionTypeEnumValues = []HeatmapBackendDefinitionType{
	HEATMAPBACKENDDEFINITIONTYPE_HEATMAP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *HeatmapBackendDefinitionType) GetAllowedValues() []HeatmapBackendDefinitionType {
	return allowedHeatmapBackendDefinitionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *HeatmapBackendDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = HeatmapBackendDefinitionType(value)
	return nil
}

// NewHeatmapBackendDefinitionTypeFromValue returns a pointer to a valid HeatmapBackendDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewHeatmapBackendDefinitionTypeFromValue(v string) (*HeatmapBackendDefinitionType, error) {
	ev := HeatmapBackendDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for HeatmapBackendDefinitionType: valid values are %v", v, allowedHeatmapBackendDefinitionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v HeatmapBackendDefinitionType) IsValid() bool {
	for _, existing := range allowedHeatmapBackendDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HeatmapBackendDefinitionType value.
func (v HeatmapBackendDefinitionType) Ptr() *HeatmapBackendDefinitionType {
	return &v
}
