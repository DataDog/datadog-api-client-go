// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HeatgridGradientConfig - Test-only generated schema for the dashboard versioning proof of concept.
type HeatgridGradientConfig struct {
	HeatgridGradientConfig20270101OneOf           *HeatgridGradientConfig20270101OneOf
	HeatgridGradientConfig20270101OneOf4457883232 *HeatgridGradientConfig20270101OneOf4457883232

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// HeatgridGradientConfig20270101OneOfAsHeatgridGradientConfig is a convenience function that returns HeatgridGradientConfig20270101OneOf wrapped in HeatgridGradientConfig.
func HeatgridGradientConfig20270101OneOfAsHeatgridGradientConfig(v *HeatgridGradientConfig20270101OneOf) HeatgridGradientConfig {
	return HeatgridGradientConfig{HeatgridGradientConfig20270101OneOf: v}
}

// HeatgridGradientConfig20270101OneOf4457883232AsHeatgridGradientConfig is a convenience function that returns HeatgridGradientConfig20270101OneOf4457883232 wrapped in HeatgridGradientConfig.
func HeatgridGradientConfig20270101OneOf4457883232AsHeatgridGradientConfig(v *HeatgridGradientConfig20270101OneOf4457883232) HeatgridGradientConfig {
	return HeatgridGradientConfig{HeatgridGradientConfig20270101OneOf4457883232: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *HeatgridGradientConfig) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into HeatgridGradientConfig20270101OneOf
	err = datadog.Unmarshal(data, &obj.HeatgridGradientConfig20270101OneOf)
	if err == nil {
		if obj.HeatgridGradientConfig20270101OneOf != nil && obj.HeatgridGradientConfig20270101OneOf.UnparsedObject == nil {
			jsonHeatgridGradientConfig20270101OneOf, _ := datadog.Marshal(obj.HeatgridGradientConfig20270101OneOf)
			if string(jsonHeatgridGradientConfig20270101OneOf) == "{}" { // empty struct
				obj.HeatgridGradientConfig20270101OneOf = nil
			} else {
				match++
			}
		} else {
			obj.HeatgridGradientConfig20270101OneOf = nil
		}
	} else {
		obj.HeatgridGradientConfig20270101OneOf = nil
	}

	// try to unmarshal data into HeatgridGradientConfig20270101OneOf4457883232
	err = datadog.Unmarshal(data, &obj.HeatgridGradientConfig20270101OneOf4457883232)
	if err == nil {
		if obj.HeatgridGradientConfig20270101OneOf4457883232 != nil && obj.HeatgridGradientConfig20270101OneOf4457883232.UnparsedObject == nil {
			jsonHeatgridGradientConfig20270101OneOf4457883232, _ := datadog.Marshal(obj.HeatgridGradientConfig20270101OneOf4457883232)
			if string(jsonHeatgridGradientConfig20270101OneOf4457883232) == "{}" { // empty struct
				obj.HeatgridGradientConfig20270101OneOf4457883232 = nil
			} else {
				match++
			}
		} else {
			obj.HeatgridGradientConfig20270101OneOf4457883232 = nil
		}
	} else {
		obj.HeatgridGradientConfig20270101OneOf4457883232 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.HeatgridGradientConfig20270101OneOf = nil
		obj.HeatgridGradientConfig20270101OneOf4457883232 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj HeatgridGradientConfig) MarshalJSON() ([]byte, error) {
	if obj.HeatgridGradientConfig20270101OneOf != nil {
		return datadog.Marshal(&obj.HeatgridGradientConfig20270101OneOf)
	}

	if obj.HeatgridGradientConfig20270101OneOf4457883232 != nil {
		return datadog.Marshal(&obj.HeatgridGradientConfig20270101OneOf4457883232)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *HeatgridGradientConfig) GetActualInstance() interface{} {
	if obj.HeatgridGradientConfig20270101OneOf != nil {
		return obj.HeatgridGradientConfig20270101OneOf
	}

	if obj.HeatgridGradientConfig20270101OneOf4457883232 != nil {
		return obj.HeatgridGradientConfig20270101OneOf4457883232
	}

	// all schemas are nil
	return nil
}
