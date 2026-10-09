// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HeatgridDiscreteMode Use discrete color thresholds.
type HeatgridDiscreteMode string

// List of HeatgridDiscreteMode.
const (
	HEATGRIDDISCRETEMODE_DISCRETE HeatgridDiscreteMode = "discrete"
)

var allowedHeatgridDiscreteModeEnumValues = []HeatgridDiscreteMode{
	HEATGRIDDISCRETEMODE_DISCRETE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *HeatgridDiscreteMode) GetAllowedValues() []HeatgridDiscreteMode {
	return allowedHeatgridDiscreteModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *HeatgridDiscreteMode) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = HeatgridDiscreteMode(value)
	return nil
}

// NewHeatgridDiscreteModeFromValue returns a pointer to a valid HeatgridDiscreteMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewHeatgridDiscreteModeFromValue(v string) (*HeatgridDiscreteMode, error) {
	ev := HeatgridDiscreteMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for HeatgridDiscreteMode: valid values are %v", v, allowedHeatgridDiscreteModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v HeatgridDiscreteMode) IsValid() bool {
	for _, existing := range allowedHeatgridDiscreteModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HeatgridDiscreteMode value.
func (v HeatgridDiscreteMode) Ptr() *HeatgridDiscreteMode {
	return &v
}
