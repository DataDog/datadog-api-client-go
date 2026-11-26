// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// BarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemType Test-only generated schema for the dashboard versioning proof of concept.
type BarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemType string

// List of BarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemType.
const (
	BARCHARTBACKENDDEFINITIONSTYLEDISPLAYANYOFANYOFITEMTYPE_STACKED BarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemType = "stacked"
)

var allowedBarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemTypeEnumValues = []BarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemType{
	BARCHARTBACKENDDEFINITIONSTYLEDISPLAYANYOFANYOFITEMTYPE_STACKED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *BarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemType) GetAllowedValues() []BarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemType {
	return allowedBarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *BarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = BarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemType(value)
	return nil
}

// NewBarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemTypeFromValue returns a pointer to a valid BarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewBarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemTypeFromValue(v string) (*BarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemType, error) {
	ev := BarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for BarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemType: valid values are %v", v, allowedBarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v BarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemType) IsValid() bool {
	for _, existing := range allowedBarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemType value.
func (v BarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemType) Ptr() *BarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemType {
	return &v
}
