// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// BarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemLegend Test-only generated schema for the dashboard versioning proof of concept.
type BarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemLegend string

// List of BarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemLegend.
const (
	BARCHARTBACKENDDEFINITIONSTYLEDISPLAYANYOFANYOFITEMLEGEND_NONE      BarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemLegend = "none"
	BARCHARTBACKENDDEFINITIONSTYLEDISPLAYANYOFANYOFITEMLEGEND_INLINE    BarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemLegend = "inline"
	BARCHARTBACKENDDEFINITIONSTYLEDISPLAYANYOFANYOFITEMLEGEND_AUTOMATIC BarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemLegend = "automatic"
)

var allowedBarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemLegendEnumValues = []BarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemLegend{
	BARCHARTBACKENDDEFINITIONSTYLEDISPLAYANYOFANYOFITEMLEGEND_NONE,
	BARCHARTBACKENDDEFINITIONSTYLEDISPLAYANYOFANYOFITEMLEGEND_INLINE,
	BARCHARTBACKENDDEFINITIONSTYLEDISPLAYANYOFANYOFITEMLEGEND_AUTOMATIC,
}

// GetAllowedValues reeturns the list of possible values.
func (v *BarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemLegend) GetAllowedValues() []BarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemLegend {
	return allowedBarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemLegendEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *BarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemLegend) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = BarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemLegend(value)
	return nil
}

// NewBarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemLegendFromValue returns a pointer to a valid BarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemLegend
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewBarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemLegendFromValue(v string) (*BarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemLegend, error) {
	ev := BarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemLegend(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for BarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemLegend: valid values are %v", v, allowedBarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemLegendEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v BarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemLegend) IsValid() bool {
	for _, existing := range allowedBarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemLegendEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemLegend value.
func (v BarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemLegend) Ptr() *BarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemLegend {
	return &v
}
