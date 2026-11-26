// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// BarChartBackendDefinitionType Test-only generated schema for the dashboard versioning proof of concept.
type BarChartBackendDefinitionType string

// List of BarChartBackendDefinitionType.
const (
	BARCHARTBACKENDDEFINITIONTYPE_BAR_CHART BarChartBackendDefinitionType = "bar_chart"
)

var allowedBarChartBackendDefinitionTypeEnumValues = []BarChartBackendDefinitionType{
	BARCHARTBACKENDDEFINITIONTYPE_BAR_CHART,
}

// GetAllowedValues reeturns the list of possible values.
func (v *BarChartBackendDefinitionType) GetAllowedValues() []BarChartBackendDefinitionType {
	return allowedBarChartBackendDefinitionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *BarChartBackendDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = BarChartBackendDefinitionType(value)
	return nil
}

// NewBarChartBackendDefinitionTypeFromValue returns a pointer to a valid BarChartBackendDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewBarChartBackendDefinitionTypeFromValue(v string) (*BarChartBackendDefinitionType, error) {
	ev := BarChartBackendDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for BarChartBackendDefinitionType: valid values are %v", v, allowedBarChartBackendDefinitionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v BarChartBackendDefinitionType) IsValid() bool {
	for _, existing := range allowedBarChartBackendDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BarChartBackendDefinitionType value.
func (v BarChartBackendDefinitionType) Ptr() *BarChartBackendDefinitionType {
	return &v
}
