// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScatterPlotLegendType - Test-only generated schema for the dashboard versioning proof of concept.
type ScatterPlotLegendType struct {
	ScatterPlotLegendType20270101OneOf           *ScatterPlotLegendType20270101OneOf
	ScatterPlotLegendType20270101OneOf4461516160 *ScatterPlotLegendType20270101OneOf4461516160

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ScatterPlotLegendType20270101OneOfAsScatterPlotLegendType is a convenience function that returns ScatterPlotLegendType20270101OneOf wrapped in ScatterPlotLegendType.
func ScatterPlotLegendType20270101OneOfAsScatterPlotLegendType(v *ScatterPlotLegendType20270101OneOf) ScatterPlotLegendType {
	return ScatterPlotLegendType{ScatterPlotLegendType20270101OneOf: v}
}

// ScatterPlotLegendType20270101OneOf4461516160AsScatterPlotLegendType is a convenience function that returns ScatterPlotLegendType20270101OneOf4461516160 wrapped in ScatterPlotLegendType.
func ScatterPlotLegendType20270101OneOf4461516160AsScatterPlotLegendType(v *ScatterPlotLegendType20270101OneOf4461516160) ScatterPlotLegendType {
	return ScatterPlotLegendType{ScatterPlotLegendType20270101OneOf4461516160: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ScatterPlotLegendType) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ScatterPlotLegendType20270101OneOf
	err = datadog.Unmarshal(data, &obj.ScatterPlotLegendType20270101OneOf)
	if err == nil {
		if obj.ScatterPlotLegendType20270101OneOf != nil && obj.ScatterPlotLegendType20270101OneOf.UnparsedObject == nil {
			jsonScatterPlotLegendType20270101OneOf, _ := datadog.Marshal(obj.ScatterPlotLegendType20270101OneOf)
			if string(jsonScatterPlotLegendType20270101OneOf) == "{}" { // empty struct
				obj.ScatterPlotLegendType20270101OneOf = nil
			} else {
				match++
			}
		} else {
			obj.ScatterPlotLegendType20270101OneOf = nil
		}
	} else {
		obj.ScatterPlotLegendType20270101OneOf = nil
	}

	// try to unmarshal data into ScatterPlotLegendType20270101OneOf4461516160
	err = datadog.Unmarshal(data, &obj.ScatterPlotLegendType20270101OneOf4461516160)
	if err == nil {
		if obj.ScatterPlotLegendType20270101OneOf4461516160 != nil && obj.ScatterPlotLegendType20270101OneOf4461516160.UnparsedObject == nil {
			jsonScatterPlotLegendType20270101OneOf4461516160, _ := datadog.Marshal(obj.ScatterPlotLegendType20270101OneOf4461516160)
			if string(jsonScatterPlotLegendType20270101OneOf4461516160) == "{}" { // empty struct
				obj.ScatterPlotLegendType20270101OneOf4461516160 = nil
			} else {
				match++
			}
		} else {
			obj.ScatterPlotLegendType20270101OneOf4461516160 = nil
		}
	} else {
		obj.ScatterPlotLegendType20270101OneOf4461516160 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ScatterPlotLegendType20270101OneOf = nil
		obj.ScatterPlotLegendType20270101OneOf4461516160 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ScatterPlotLegendType) MarshalJSON() ([]byte, error) {
	if obj.ScatterPlotLegendType20270101OneOf != nil {
		return datadog.Marshal(&obj.ScatterPlotLegendType20270101OneOf)
	}

	if obj.ScatterPlotLegendType20270101OneOf4461516160 != nil {
		return datadog.Marshal(&obj.ScatterPlotLegendType20270101OneOf4461516160)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ScatterPlotLegendType) GetActualInstance() interface{} {
	if obj.ScatterPlotLegendType20270101OneOf != nil {
		return obj.ScatterPlotLegendType20270101OneOf
	}

	if obj.ScatterPlotLegendType20270101OneOf4461516160 != nil {
		return obj.ScatterPlotLegendType20270101OneOf4461516160
	}

	// all schemas are nil
	return nil
}
