// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ImageBackendDefinition20270101Sizing - Test-only generated schema for the dashboard versioning proof of concept.
type ImageBackendDefinition20270101Sizing struct {
	Float64              *float64
	ImageSizing_20270101 *ImageSizing_20270101

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// Float64AsImageBackendDefinition20270101Sizing is a convenience function that returns float64 wrapped in ImageBackendDefinition20270101Sizing.
func Float64AsImageBackendDefinition20270101Sizing(v *float64) ImageBackendDefinition20270101Sizing {
	return ImageBackendDefinition20270101Sizing{Float64: v}
}

// ImageSizing_20270101AsImageBackendDefinition20270101Sizing is a convenience function that returns ImageSizing_20270101 wrapped in ImageBackendDefinition20270101Sizing.
func ImageSizing_20270101AsImageBackendDefinition20270101Sizing(v *ImageSizing_20270101) ImageBackendDefinition20270101Sizing {
	return ImageBackendDefinition20270101Sizing{ImageSizing_20270101: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ImageBackendDefinition20270101Sizing) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Float64
	err = datadog.Unmarshal(data, &obj.Float64)
	if err == nil {
		if obj.Float64 != nil {
			jsonFloat64, _ := datadog.Marshal(obj.Float64)
			if string(jsonFloat64) == "{}" { // empty struct
				obj.Float64 = nil
			} else {
				match++
			}
		} else {
			obj.Float64 = nil
		}
	} else {
		obj.Float64 = nil
	}

	// try to unmarshal data into ImageSizing_20270101
	err = datadog.Unmarshal(data, &obj.ImageSizing_20270101)
	if err == nil {
		if obj.ImageSizing_20270101 != nil {
			jsonImageSizing_20270101, _ := datadog.Marshal(obj.ImageSizing_20270101)
			if string(jsonImageSizing_20270101) == "{}" && string(data) != "{}" { // empty struct
				obj.ImageSizing_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.ImageSizing_20270101 = nil
		}
	} else {
		obj.ImageSizing_20270101 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.Float64 = nil
		obj.ImageSizing_20270101 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ImageBackendDefinition20270101Sizing) MarshalJSON() ([]byte, error) {
	if obj.Float64 != nil {
		return datadog.Marshal(&obj.Float64)
	}

	if obj.ImageSizing_20270101 != nil {
		return datadog.Marshal(&obj.ImageSizing_20270101)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ImageBackendDefinition20270101Sizing) GetActualInstance() interface{} {
	if obj.Float64 != nil {
		return obj.Float64
	}

	if obj.ImageSizing_20270101 != nil {
		return obj.ImageSizing_20270101
	}

	// all schemas are nil
	return nil
}
