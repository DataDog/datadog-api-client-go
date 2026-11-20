// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SunburstKnownLegend - Test-only generated schema for the dashboard versioning proof of concept.
type SunburstKnownLegend struct {
	SunburstKnownLegend20270101OneOf           *SunburstKnownLegend20270101OneOf
	SunburstKnownLegend20270101OneOf4462427680 *SunburstKnownLegend20270101OneOf4462427680
	SunburstKnownLegend20270101OneOf4462657920 *SunburstKnownLegend20270101OneOf4462657920

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// SunburstKnownLegend20270101OneOfAsSunburstKnownLegend is a convenience function that returns SunburstKnownLegend20270101OneOf wrapped in SunburstKnownLegend.
func SunburstKnownLegend20270101OneOfAsSunburstKnownLegend(v *SunburstKnownLegend20270101OneOf) SunburstKnownLegend {
	return SunburstKnownLegend{SunburstKnownLegend20270101OneOf: v}
}

// SunburstKnownLegend20270101OneOf4462427680AsSunburstKnownLegend is a convenience function that returns SunburstKnownLegend20270101OneOf4462427680 wrapped in SunburstKnownLegend.
func SunburstKnownLegend20270101OneOf4462427680AsSunburstKnownLegend(v *SunburstKnownLegend20270101OneOf4462427680) SunburstKnownLegend {
	return SunburstKnownLegend{SunburstKnownLegend20270101OneOf4462427680: v}
}

// SunburstKnownLegend20270101OneOf4462657920AsSunburstKnownLegend is a convenience function that returns SunburstKnownLegend20270101OneOf4462657920 wrapped in SunburstKnownLegend.
func SunburstKnownLegend20270101OneOf4462657920AsSunburstKnownLegend(v *SunburstKnownLegend20270101OneOf4462657920) SunburstKnownLegend {
	return SunburstKnownLegend{SunburstKnownLegend20270101OneOf4462657920: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *SunburstKnownLegend) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SunburstKnownLegend20270101OneOf
	err = datadog.Unmarshal(data, &obj.SunburstKnownLegend20270101OneOf)
	if err == nil {
		if obj.SunburstKnownLegend20270101OneOf != nil && obj.SunburstKnownLegend20270101OneOf.UnparsedObject == nil {
			jsonSunburstKnownLegend20270101OneOf, _ := datadog.Marshal(obj.SunburstKnownLegend20270101OneOf)
			if string(jsonSunburstKnownLegend20270101OneOf) == "{}" { // empty struct
				obj.SunburstKnownLegend20270101OneOf = nil
			} else {
				match++
			}
		} else {
			obj.SunburstKnownLegend20270101OneOf = nil
		}
	} else {
		obj.SunburstKnownLegend20270101OneOf = nil
	}

	// try to unmarshal data into SunburstKnownLegend20270101OneOf4462427680
	err = datadog.Unmarshal(data, &obj.SunburstKnownLegend20270101OneOf4462427680)
	if err == nil {
		if obj.SunburstKnownLegend20270101OneOf4462427680 != nil && obj.SunburstKnownLegend20270101OneOf4462427680.UnparsedObject == nil {
			jsonSunburstKnownLegend20270101OneOf4462427680, _ := datadog.Marshal(obj.SunburstKnownLegend20270101OneOf4462427680)
			if string(jsonSunburstKnownLegend20270101OneOf4462427680) == "{}" { // empty struct
				obj.SunburstKnownLegend20270101OneOf4462427680 = nil
			} else {
				match++
			}
		} else {
			obj.SunburstKnownLegend20270101OneOf4462427680 = nil
		}
	} else {
		obj.SunburstKnownLegend20270101OneOf4462427680 = nil
	}

	// try to unmarshal data into SunburstKnownLegend20270101OneOf4462657920
	err = datadog.Unmarshal(data, &obj.SunburstKnownLegend20270101OneOf4462657920)
	if err == nil {
		if obj.SunburstKnownLegend20270101OneOf4462657920 != nil && obj.SunburstKnownLegend20270101OneOf4462657920.UnparsedObject == nil {
			jsonSunburstKnownLegend20270101OneOf4462657920, _ := datadog.Marshal(obj.SunburstKnownLegend20270101OneOf4462657920)
			if string(jsonSunburstKnownLegend20270101OneOf4462657920) == "{}" { // empty struct
				obj.SunburstKnownLegend20270101OneOf4462657920 = nil
			} else {
				match++
			}
		} else {
			obj.SunburstKnownLegend20270101OneOf4462657920 = nil
		}
	} else {
		obj.SunburstKnownLegend20270101OneOf4462657920 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.SunburstKnownLegend20270101OneOf = nil
		obj.SunburstKnownLegend20270101OneOf4462427680 = nil
		obj.SunburstKnownLegend20270101OneOf4462657920 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj SunburstKnownLegend) MarshalJSON() ([]byte, error) {
	if obj.SunburstKnownLegend20270101OneOf != nil {
		return datadog.Marshal(&obj.SunburstKnownLegend20270101OneOf)
	}

	if obj.SunburstKnownLegend20270101OneOf4462427680 != nil {
		return datadog.Marshal(&obj.SunburstKnownLegend20270101OneOf4462427680)
	}

	if obj.SunburstKnownLegend20270101OneOf4462657920 != nil {
		return datadog.Marshal(&obj.SunburstKnownLegend20270101OneOf4462657920)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *SunburstKnownLegend) GetActualInstance() interface{} {
	if obj.SunburstKnownLegend20270101OneOf != nil {
		return obj.SunburstKnownLegend20270101OneOf
	}

	if obj.SunburstKnownLegend20270101OneOf4462427680 != nil {
		return obj.SunburstKnownLegend20270101OneOf4462427680
	}

	if obj.SunburstKnownLegend20270101OneOf4462657920 != nil {
		return obj.SunburstKnownLegend20270101OneOf4462657920
	}

	// all schemas are nil
	return nil
}
