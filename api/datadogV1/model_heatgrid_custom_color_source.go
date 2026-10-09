// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HeatgridCustomColorSource Use custom colors.
type HeatgridCustomColorSource string

// List of HeatgridCustomColorSource.
const (
	HEATGRIDCUSTOMCOLORSOURCE_CUSTOM HeatgridCustomColorSource = "custom"
)

var allowedHeatgridCustomColorSourceEnumValues = []HeatgridCustomColorSource{
	HEATGRIDCUSTOMCOLORSOURCE_CUSTOM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *HeatgridCustomColorSource) GetAllowedValues() []HeatgridCustomColorSource {
	return allowedHeatgridCustomColorSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *HeatgridCustomColorSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = HeatgridCustomColorSource(value)
	return nil
}

// NewHeatgridCustomColorSourceFromValue returns a pointer to a valid HeatgridCustomColorSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewHeatgridCustomColorSourceFromValue(v string) (*HeatgridCustomColorSource, error) {
	ev := HeatgridCustomColorSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for HeatgridCustomColorSource: valid values are %v", v, allowedHeatgridCustomColorSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v HeatgridCustomColorSource) IsValid() bool {
	for _, existing := range allowedHeatgridCustomColorSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HeatgridCustomColorSource value.
func (v HeatgridCustomColorSource) Ptr() *HeatgridCustomColorSource {
	return &v
}
