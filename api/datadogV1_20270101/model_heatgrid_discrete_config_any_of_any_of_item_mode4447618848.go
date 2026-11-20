// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HeatgridDiscreteConfigAnyOfAnyOfItemMode4447618848 Test-only generated schema for the dashboard versioning proof of concept.
type HeatgridDiscreteConfigAnyOfAnyOfItemMode4447618848 string

// List of HeatgridDiscreteConfigAnyOfAnyOfItemMode4447618848.
const (
	HEATGRIDDISCRETECONFIGANYOFANYOFITEMMODE4447618848_DISCRETE HeatgridDiscreteConfigAnyOfAnyOfItemMode4447618848 = "discrete"
)

var allowedHeatgridDiscreteConfigAnyOfAnyOfItemMode4447618848EnumValues = []HeatgridDiscreteConfigAnyOfAnyOfItemMode4447618848{
	HEATGRIDDISCRETECONFIGANYOFANYOFITEMMODE4447618848_DISCRETE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *HeatgridDiscreteConfigAnyOfAnyOfItemMode4447618848) GetAllowedValues() []HeatgridDiscreteConfigAnyOfAnyOfItemMode4447618848 {
	return allowedHeatgridDiscreteConfigAnyOfAnyOfItemMode4447618848EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *HeatgridDiscreteConfigAnyOfAnyOfItemMode4447618848) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = HeatgridDiscreteConfigAnyOfAnyOfItemMode4447618848(value)
	return nil
}

// NewHeatgridDiscreteConfigAnyOfAnyOfItemMode4447618848FromValue returns a pointer to a valid HeatgridDiscreteConfigAnyOfAnyOfItemMode4447618848
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewHeatgridDiscreteConfigAnyOfAnyOfItemMode4447618848FromValue(v string) (*HeatgridDiscreteConfigAnyOfAnyOfItemMode4447618848, error) {
	ev := HeatgridDiscreteConfigAnyOfAnyOfItemMode4447618848(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for HeatgridDiscreteConfigAnyOfAnyOfItemMode4447618848: valid values are %v", v, allowedHeatgridDiscreteConfigAnyOfAnyOfItemMode4447618848EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v HeatgridDiscreteConfigAnyOfAnyOfItemMode4447618848) IsValid() bool {
	for _, existing := range allowedHeatgridDiscreteConfigAnyOfAnyOfItemMode4447618848EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HeatgridDiscreteConfigAnyOfAnyOfItemMode4447618848 value.
func (v HeatgridDiscreteConfigAnyOfAnyOfItemMode4447618848) Ptr() *HeatgridDiscreteConfigAnyOfAnyOfItemMode4447618848 {
	return &v
}
