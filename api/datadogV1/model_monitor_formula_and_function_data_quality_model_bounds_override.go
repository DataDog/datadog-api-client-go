// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorFormulaAndFunctionDataQualityModelBoundsOverride Restricts which predicted bound the monitor alerts on. `UPPER_ONLY` alerts only when
// the measure rises above the upper bound, `LOWER_ONLY` only when it falls below the
// lower bound. When unset, the monitor alerts on both.
type MonitorFormulaAndFunctionDataQualityModelBoundsOverride string

// List of MonitorFormulaAndFunctionDataQualityModelBoundsOverride.
const (
	MONITORFORMULAANDFUNCTIONDATAQUALITYMODELBOUNDSOVERRIDE_UPPER_ONLY MonitorFormulaAndFunctionDataQualityModelBoundsOverride = "UPPER_ONLY"
	MONITORFORMULAANDFUNCTIONDATAQUALITYMODELBOUNDSOVERRIDE_LOWER_ONLY MonitorFormulaAndFunctionDataQualityModelBoundsOverride = "LOWER_ONLY"
)

var allowedMonitorFormulaAndFunctionDataQualityModelBoundsOverrideEnumValues = []MonitorFormulaAndFunctionDataQualityModelBoundsOverride{
	MONITORFORMULAANDFUNCTIONDATAQUALITYMODELBOUNDSOVERRIDE_UPPER_ONLY,
	MONITORFORMULAANDFUNCTIONDATAQUALITYMODELBOUNDSOVERRIDE_LOWER_ONLY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *MonitorFormulaAndFunctionDataQualityModelBoundsOverride) GetAllowedValues() []MonitorFormulaAndFunctionDataQualityModelBoundsOverride {
	return allowedMonitorFormulaAndFunctionDataQualityModelBoundsOverrideEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *MonitorFormulaAndFunctionDataQualityModelBoundsOverride) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MonitorFormulaAndFunctionDataQualityModelBoundsOverride(value)
	return nil
}

// NewMonitorFormulaAndFunctionDataQualityModelBoundsOverrideFromValue returns a pointer to a valid MonitorFormulaAndFunctionDataQualityModelBoundsOverride
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewMonitorFormulaAndFunctionDataQualityModelBoundsOverrideFromValue(v string) (*MonitorFormulaAndFunctionDataQualityModelBoundsOverride, error) {
	ev := MonitorFormulaAndFunctionDataQualityModelBoundsOverride(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for MonitorFormulaAndFunctionDataQualityModelBoundsOverride: valid values are %v", v, allowedMonitorFormulaAndFunctionDataQualityModelBoundsOverrideEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v MonitorFormulaAndFunctionDataQualityModelBoundsOverride) IsValid() bool {
	for _, existing := range allowedMonitorFormulaAndFunctionDataQualityModelBoundsOverrideEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MonitorFormulaAndFunctionDataQualityModelBoundsOverride value.
func (v MonitorFormulaAndFunctionDataQualityModelBoundsOverride) Ptr() *MonitorFormulaAndFunctionDataQualityModelBoundsOverride {
	return &v
}
