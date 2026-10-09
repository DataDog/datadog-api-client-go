// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HeatgridColorConfig - Color configuration for continuous gradients or discrete thresholds.
type HeatgridColorConfig struct {
	HeatgridGradientCustomColor *HeatgridGradientCustomColor
	HeatgridGradientPresetColor *HeatgridGradientPresetColor
	HeatgridDiscreteCustomColor *HeatgridDiscreteCustomColor
	HeatgridDiscretePresetColor *HeatgridDiscretePresetColor

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// HeatgridGradientCustomColorAsHeatgridColorConfig is a convenience function that returns HeatgridGradientCustomColor wrapped in HeatgridColorConfig.
func HeatgridGradientCustomColorAsHeatgridColorConfig(v *HeatgridGradientCustomColor) HeatgridColorConfig {
	return HeatgridColorConfig{HeatgridGradientCustomColor: v}
}

// HeatgridGradientPresetColorAsHeatgridColorConfig is a convenience function that returns HeatgridGradientPresetColor wrapped in HeatgridColorConfig.
func HeatgridGradientPresetColorAsHeatgridColorConfig(v *HeatgridGradientPresetColor) HeatgridColorConfig {
	return HeatgridColorConfig{HeatgridGradientPresetColor: v}
}

// HeatgridDiscreteCustomColorAsHeatgridColorConfig is a convenience function that returns HeatgridDiscreteCustomColor wrapped in HeatgridColorConfig.
func HeatgridDiscreteCustomColorAsHeatgridColorConfig(v *HeatgridDiscreteCustomColor) HeatgridColorConfig {
	return HeatgridColorConfig{HeatgridDiscreteCustomColor: v}
}

// HeatgridDiscretePresetColorAsHeatgridColorConfig is a convenience function that returns HeatgridDiscretePresetColor wrapped in HeatgridColorConfig.
func HeatgridDiscretePresetColorAsHeatgridColorConfig(v *HeatgridDiscretePresetColor) HeatgridColorConfig {
	return HeatgridColorConfig{HeatgridDiscretePresetColor: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *HeatgridColorConfig) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into HeatgridGradientCustomColor
	err = datadog.Unmarshal(data, &obj.HeatgridGradientCustomColor)
	if err == nil {
		if obj.HeatgridGradientCustomColor != nil && obj.HeatgridGradientCustomColor.UnparsedObject == nil {
			jsonHeatgridGradientCustomColor, _ := datadog.Marshal(obj.HeatgridGradientCustomColor)
			if string(jsonHeatgridGradientCustomColor) == "{}" { // empty struct
				obj.HeatgridGradientCustomColor = nil
			} else {
				match++
			}
		} else {
			obj.HeatgridGradientCustomColor = nil
		}
	} else {
		obj.HeatgridGradientCustomColor = nil
	}

	// try to unmarshal data into HeatgridGradientPresetColor
	err = datadog.Unmarshal(data, &obj.HeatgridGradientPresetColor)
	if err == nil {
		if obj.HeatgridGradientPresetColor != nil && obj.HeatgridGradientPresetColor.UnparsedObject == nil {
			jsonHeatgridGradientPresetColor, _ := datadog.Marshal(obj.HeatgridGradientPresetColor)
			if string(jsonHeatgridGradientPresetColor) == "{}" { // empty struct
				obj.HeatgridGradientPresetColor = nil
			} else {
				match++
			}
		} else {
			obj.HeatgridGradientPresetColor = nil
		}
	} else {
		obj.HeatgridGradientPresetColor = nil
	}

	// try to unmarshal data into HeatgridDiscreteCustomColor
	err = datadog.Unmarshal(data, &obj.HeatgridDiscreteCustomColor)
	if err == nil {
		if obj.HeatgridDiscreteCustomColor != nil && obj.HeatgridDiscreteCustomColor.UnparsedObject == nil {
			jsonHeatgridDiscreteCustomColor, _ := datadog.Marshal(obj.HeatgridDiscreteCustomColor)
			if string(jsonHeatgridDiscreteCustomColor) == "{}" { // empty struct
				obj.HeatgridDiscreteCustomColor = nil
			} else {
				match++
			}
		} else {
			obj.HeatgridDiscreteCustomColor = nil
		}
	} else {
		obj.HeatgridDiscreteCustomColor = nil
	}

	// try to unmarshal data into HeatgridDiscretePresetColor
	err = datadog.Unmarshal(data, &obj.HeatgridDiscretePresetColor)
	if err == nil {
		if obj.HeatgridDiscretePresetColor != nil && obj.HeatgridDiscretePresetColor.UnparsedObject == nil {
			jsonHeatgridDiscretePresetColor, _ := datadog.Marshal(obj.HeatgridDiscretePresetColor)
			if string(jsonHeatgridDiscretePresetColor) == "{}" { // empty struct
				obj.HeatgridDiscretePresetColor = nil
			} else {
				match++
			}
		} else {
			obj.HeatgridDiscretePresetColor = nil
		}
	} else {
		obj.HeatgridDiscretePresetColor = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.HeatgridGradientCustomColor = nil
		obj.HeatgridGradientPresetColor = nil
		obj.HeatgridDiscreteCustomColor = nil
		obj.HeatgridDiscretePresetColor = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj HeatgridColorConfig) MarshalJSON() ([]byte, error) {
	if obj.HeatgridGradientCustomColor != nil {
		return datadog.Marshal(&obj.HeatgridGradientCustomColor)
	}

	if obj.HeatgridGradientPresetColor != nil {
		return datadog.Marshal(&obj.HeatgridGradientPresetColor)
	}

	if obj.HeatgridDiscreteCustomColor != nil {
		return datadog.Marshal(&obj.HeatgridDiscreteCustomColor)
	}

	if obj.HeatgridDiscretePresetColor != nil {
		return datadog.Marshal(&obj.HeatgridDiscretePresetColor)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *HeatgridColorConfig) GetActualInstance() interface{} {
	if obj.HeatgridGradientCustomColor != nil {
		return obj.HeatgridGradientCustomColor
	}

	if obj.HeatgridGradientPresetColor != nil {
		return obj.HeatgridGradientPresetColor
	}

	if obj.HeatgridDiscreteCustomColor != nil {
		return obj.HeatgridDiscreteCustomColor
	}

	if obj.HeatgridDiscretePresetColor != nil {
		return obj.HeatgridDiscretePresetColor
	}

	// all schemas are nil
	return nil
}
