// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// BarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemType4443916544 Test-only generated schema for the dashboard versioning proof of concept.
type BarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemType4443916544 string

// List of BarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemType4443916544.
const (
	BARCHARTBACKENDDEFINITIONSTYLEDISPLAYANYOFANYOFITEMTYPE4443916544_FLAT BarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemType4443916544 = "flat"
)

var allowedBarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemType4443916544EnumValues = []BarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemType4443916544{
	BARCHARTBACKENDDEFINITIONSTYLEDISPLAYANYOFANYOFITEMTYPE4443916544_FLAT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *BarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemType4443916544) GetAllowedValues() []BarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemType4443916544 {
	return allowedBarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemType4443916544EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *BarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemType4443916544) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = BarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemType4443916544(value)
	return nil
}

// NewBarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemType4443916544FromValue returns a pointer to a valid BarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemType4443916544
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewBarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemType4443916544FromValue(v string) (*BarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemType4443916544, error) {
	ev := BarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemType4443916544(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for BarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemType4443916544: valid values are %v", v, allowedBarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemType4443916544EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v BarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemType4443916544) IsValid() bool {
	for _, existing := range allowedBarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemType4443916544EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemType4443916544 value.
func (v BarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemType4443916544) Ptr() *BarChartBackendDefinitionStyleDisplayAnyOfAnyOfItemType4443916544 {
	return &v
}
