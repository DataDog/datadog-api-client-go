// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HeatgridGradientMode Use a continuous color gradient.
type HeatgridGradientMode string

// List of HeatgridGradientMode.
const (
	HEATGRIDGRADIENTMODE_GRADIENT HeatgridGradientMode = "gradient"
)

var allowedHeatgridGradientModeEnumValues = []HeatgridGradientMode{
	HEATGRIDGRADIENTMODE_GRADIENT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *HeatgridGradientMode) GetAllowedValues() []HeatgridGradientMode {
	return allowedHeatgridGradientModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *HeatgridGradientMode) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = HeatgridGradientMode(value)
	return nil
}

// NewHeatgridGradientModeFromValue returns a pointer to a valid HeatgridGradientMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewHeatgridGradientModeFromValue(v string) (*HeatgridGradientMode, error) {
	ev := HeatgridGradientMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for HeatgridGradientMode: valid values are %v", v, allowedHeatgridGradientModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v HeatgridGradientMode) IsValid() bool {
	for _, existing := range allowedHeatgridGradientModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HeatgridGradientMode value.
func (v HeatgridGradientMode) Ptr() *HeatgridGradientMode {
	return &v
}
