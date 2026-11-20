// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HeatgridDiscreteConfig - Test-only generated schema for the dashboard versioning proof of concept.
type HeatgridDiscreteConfig struct {
	HeatgridDiscreteConfig20270101OneOf           *HeatgridDiscreteConfig20270101OneOf
	HeatgridDiscreteConfig20270101OneOf4457881312 *HeatgridDiscreteConfig20270101OneOf4457881312

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// HeatgridDiscreteConfig20270101OneOfAsHeatgridDiscreteConfig is a convenience function that returns HeatgridDiscreteConfig20270101OneOf wrapped in HeatgridDiscreteConfig.
func HeatgridDiscreteConfig20270101OneOfAsHeatgridDiscreteConfig(v *HeatgridDiscreteConfig20270101OneOf) HeatgridDiscreteConfig {
	return HeatgridDiscreteConfig{HeatgridDiscreteConfig20270101OneOf: v}
}

// HeatgridDiscreteConfig20270101OneOf4457881312AsHeatgridDiscreteConfig is a convenience function that returns HeatgridDiscreteConfig20270101OneOf4457881312 wrapped in HeatgridDiscreteConfig.
func HeatgridDiscreteConfig20270101OneOf4457881312AsHeatgridDiscreteConfig(v *HeatgridDiscreteConfig20270101OneOf4457881312) HeatgridDiscreteConfig {
	return HeatgridDiscreteConfig{HeatgridDiscreteConfig20270101OneOf4457881312: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *HeatgridDiscreteConfig) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into HeatgridDiscreteConfig20270101OneOf
	err = datadog.Unmarshal(data, &obj.HeatgridDiscreteConfig20270101OneOf)
	if err == nil {
		if obj.HeatgridDiscreteConfig20270101OneOf != nil && obj.HeatgridDiscreteConfig20270101OneOf.UnparsedObject == nil {
			jsonHeatgridDiscreteConfig20270101OneOf, _ := datadog.Marshal(obj.HeatgridDiscreteConfig20270101OneOf)
			if string(jsonHeatgridDiscreteConfig20270101OneOf) == "{}" { // empty struct
				obj.HeatgridDiscreteConfig20270101OneOf = nil
			} else {
				match++
			}
		} else {
			obj.HeatgridDiscreteConfig20270101OneOf = nil
		}
	} else {
		obj.HeatgridDiscreteConfig20270101OneOf = nil
	}

	// try to unmarshal data into HeatgridDiscreteConfig20270101OneOf4457881312
	err = datadog.Unmarshal(data, &obj.HeatgridDiscreteConfig20270101OneOf4457881312)
	if err == nil {
		if obj.HeatgridDiscreteConfig20270101OneOf4457881312 != nil && obj.HeatgridDiscreteConfig20270101OneOf4457881312.UnparsedObject == nil {
			jsonHeatgridDiscreteConfig20270101OneOf4457881312, _ := datadog.Marshal(obj.HeatgridDiscreteConfig20270101OneOf4457881312)
			if string(jsonHeatgridDiscreteConfig20270101OneOf4457881312) == "{}" { // empty struct
				obj.HeatgridDiscreteConfig20270101OneOf4457881312 = nil
			} else {
				match++
			}
		} else {
			obj.HeatgridDiscreteConfig20270101OneOf4457881312 = nil
		}
	} else {
		obj.HeatgridDiscreteConfig20270101OneOf4457881312 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.HeatgridDiscreteConfig20270101OneOf = nil
		obj.HeatgridDiscreteConfig20270101OneOf4457881312 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj HeatgridDiscreteConfig) MarshalJSON() ([]byte, error) {
	if obj.HeatgridDiscreteConfig20270101OneOf != nil {
		return datadog.Marshal(&obj.HeatgridDiscreteConfig20270101OneOf)
	}

	if obj.HeatgridDiscreteConfig20270101OneOf4457881312 != nil {
		return datadog.Marshal(&obj.HeatgridDiscreteConfig20270101OneOf4457881312)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *HeatgridDiscreteConfig) GetActualInstance() interface{} {
	if obj.HeatgridDiscreteConfig20270101OneOf != nil {
		return obj.HeatgridDiscreteConfig20270101OneOf
	}

	if obj.HeatgridDiscreteConfig20270101OneOf4457881312 != nil {
		return obj.HeatgridDiscreteConfig20270101OneOf4457881312
	}

	// all schemas are nil
	return nil
}
