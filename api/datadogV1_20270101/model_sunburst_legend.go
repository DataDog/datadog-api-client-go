// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SunburstLegend - Test-only generated schema for the dashboard versioning proof of concept.
type SunburstLegend struct {
	SunburstLegend20270101OneOf  *SunburstLegend20270101OneOf
	SunburstKnownLegend_20270101 *SunburstKnownLegend

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// SunburstLegend20270101OneOfAsSunburstLegend is a convenience function that returns SunburstLegend20270101OneOf wrapped in SunburstLegend.
func SunburstLegend20270101OneOfAsSunburstLegend(v *SunburstLegend20270101OneOf) SunburstLegend {
	return SunburstLegend{SunburstLegend20270101OneOf: v}
}

// SunburstKnownLegend_20270101AsSunburstLegend is a convenience function that returns SunburstKnownLegend wrapped in SunburstLegend.
func SunburstKnownLegend_20270101AsSunburstLegend(v *SunburstKnownLegend) SunburstLegend {
	return SunburstLegend{SunburstKnownLegend_20270101: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *SunburstLegend) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SunburstLegend20270101OneOf
	err = datadog.Unmarshal(data, &obj.SunburstLegend20270101OneOf)
	if err == nil {
		if obj.SunburstLegend20270101OneOf != nil && obj.SunburstLegend20270101OneOf.UnparsedObject == nil {
			jsonSunburstLegend20270101OneOf, _ := datadog.Marshal(obj.SunburstLegend20270101OneOf)
			if string(jsonSunburstLegend20270101OneOf) == "{}" { // empty struct
				obj.SunburstLegend20270101OneOf = nil
			} else {
				match++
			}
		} else {
			obj.SunburstLegend20270101OneOf = nil
		}
	} else {
		obj.SunburstLegend20270101OneOf = nil
	}

	// try to unmarshal data into SunburstKnownLegend_20270101
	err = datadog.Unmarshal(data, &obj.SunburstKnownLegend_20270101)
	if err == nil {
		if obj.SunburstKnownLegend_20270101 != nil && obj.SunburstKnownLegend_20270101.UnparsedObject == nil {
			jsonSunburstKnownLegend_20270101, _ := datadog.Marshal(obj.SunburstKnownLegend_20270101)
			if string(jsonSunburstKnownLegend_20270101) == "{}" && string(data) != "{}" { // empty struct
				obj.SunburstKnownLegend_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.SunburstKnownLegend_20270101 = nil
		}
	} else {
		obj.SunburstKnownLegend_20270101 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.SunburstLegend20270101OneOf = nil
		obj.SunburstKnownLegend_20270101 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj SunburstLegend) MarshalJSON() ([]byte, error) {
	if obj.SunburstLegend20270101OneOf != nil {
		return datadog.Marshal(&obj.SunburstLegend20270101OneOf)
	}

	if obj.SunburstKnownLegend_20270101 != nil {
		return datadog.Marshal(&obj.SunburstKnownLegend_20270101)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *SunburstLegend) GetActualInstance() interface{} {
	if obj.SunburstLegend20270101OneOf != nil {
		return obj.SunburstLegend20270101OneOf
	}

	if obj.SunburstKnownLegend_20270101 != nil {
		return obj.SunburstKnownLegend_20270101
	}

	// all schemas are nil
	return nil
}
