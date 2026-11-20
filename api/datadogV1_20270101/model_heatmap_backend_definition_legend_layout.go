// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HeatmapBackendDefinitionLegendLayout Test-only generated schema for the dashboard versioning proof of concept.
type HeatmapBackendDefinitionLegendLayout string

// List of HeatmapBackendDefinitionLegendLayout.
const (
	HEATMAPBACKENDDEFINITIONLEGENDLAYOUT_VERTICAL   HeatmapBackendDefinitionLegendLayout = "vertical"
	HEATMAPBACKENDDEFINITIONLEGENDLAYOUT_HORIZONTAL HeatmapBackendDefinitionLegendLayout = "horizontal"
	HEATMAPBACKENDDEFINITIONLEGENDLAYOUT_AUTO       HeatmapBackendDefinitionLegendLayout = "auto"
)

var allowedHeatmapBackendDefinitionLegendLayoutEnumValues = []HeatmapBackendDefinitionLegendLayout{
	HEATMAPBACKENDDEFINITIONLEGENDLAYOUT_VERTICAL,
	HEATMAPBACKENDDEFINITIONLEGENDLAYOUT_HORIZONTAL,
	HEATMAPBACKENDDEFINITIONLEGENDLAYOUT_AUTO,
}

// GetAllowedValues reeturns the list of possible values.
func (v *HeatmapBackendDefinitionLegendLayout) GetAllowedValues() []HeatmapBackendDefinitionLegendLayout {
	return allowedHeatmapBackendDefinitionLegendLayoutEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *HeatmapBackendDefinitionLegendLayout) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = HeatmapBackendDefinitionLegendLayout(value)
	return nil
}

// NewHeatmapBackendDefinitionLegendLayoutFromValue returns a pointer to a valid HeatmapBackendDefinitionLegendLayout
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewHeatmapBackendDefinitionLegendLayoutFromValue(v string) (*HeatmapBackendDefinitionLegendLayout, error) {
	ev := HeatmapBackendDefinitionLegendLayout(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for HeatmapBackendDefinitionLegendLayout: valid values are %v", v, allowedHeatmapBackendDefinitionLegendLayoutEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v HeatmapBackendDefinitionLegendLayout) IsValid() bool {
	for _, existing := range allowedHeatmapBackendDefinitionLegendLayoutEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HeatmapBackendDefinitionLegendLayout value.
func (v HeatmapBackendDefinitionLegendLayout) Ptr() *HeatmapBackendDefinitionLegendLayout {
	return &v
}
