// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HeatgridColorConfig - Test-only generated schema for the dashboard versioning proof of concept.
type HeatgridColorConfig struct {
	HeatgridGradientConfig_20270101 *HeatgridGradientConfig
	HeatgridDiscreteConfig_20270101 *HeatgridDiscreteConfig

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// HeatgridGradientConfig_20270101AsHeatgridColorConfig is a convenience function that returns HeatgridGradientConfig wrapped in HeatgridColorConfig.
func HeatgridGradientConfig_20270101AsHeatgridColorConfig(v *HeatgridGradientConfig) HeatgridColorConfig {
	return HeatgridColorConfig{HeatgridGradientConfig_20270101: v}
}

// HeatgridDiscreteConfig_20270101AsHeatgridColorConfig is a convenience function that returns HeatgridDiscreteConfig wrapped in HeatgridColorConfig.
func HeatgridDiscreteConfig_20270101AsHeatgridColorConfig(v *HeatgridDiscreteConfig) HeatgridColorConfig {
	return HeatgridColorConfig{HeatgridDiscreteConfig_20270101: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *HeatgridColorConfig) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into HeatgridGradientConfig_20270101
	err = datadog.Unmarshal(data, &obj.HeatgridGradientConfig_20270101)
	if err == nil {
		if obj.HeatgridGradientConfig_20270101 != nil && obj.HeatgridGradientConfig_20270101.UnparsedObject == nil {
			jsonHeatgridGradientConfig_20270101, _ := datadog.Marshal(obj.HeatgridGradientConfig_20270101)
			if string(jsonHeatgridGradientConfig_20270101) == "{}" && string(data) != "{}" { // empty struct
				obj.HeatgridGradientConfig_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.HeatgridGradientConfig_20270101 = nil
		}
	} else {
		obj.HeatgridGradientConfig_20270101 = nil
	}

	// try to unmarshal data into HeatgridDiscreteConfig_20270101
	err = datadog.Unmarshal(data, &obj.HeatgridDiscreteConfig_20270101)
	if err == nil {
		if obj.HeatgridDiscreteConfig_20270101 != nil && obj.HeatgridDiscreteConfig_20270101.UnparsedObject == nil {
			jsonHeatgridDiscreteConfig_20270101, _ := datadog.Marshal(obj.HeatgridDiscreteConfig_20270101)
			if string(jsonHeatgridDiscreteConfig_20270101) == "{}" && string(data) != "{}" { // empty struct
				obj.HeatgridDiscreteConfig_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.HeatgridDiscreteConfig_20270101 = nil
		}
	} else {
		obj.HeatgridDiscreteConfig_20270101 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.HeatgridGradientConfig_20270101 = nil
		obj.HeatgridDiscreteConfig_20270101 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj HeatgridColorConfig) MarshalJSON() ([]byte, error) {
	if obj.HeatgridGradientConfig_20270101 != nil {
		return datadog.Marshal(&obj.HeatgridGradientConfig_20270101)
	}

	if obj.HeatgridDiscreteConfig_20270101 != nil {
		return datadog.Marshal(&obj.HeatgridDiscreteConfig_20270101)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *HeatgridColorConfig) GetActualInstance() interface{} {
	if obj.HeatgridGradientConfig_20270101 != nil {
		return obj.HeatgridGradientConfig_20270101
	}

	if obj.HeatgridDiscreteConfig_20270101 != nil {
		return obj.HeatgridDiscreteConfig_20270101
	}

	// all schemas are nil
	return nil
}
