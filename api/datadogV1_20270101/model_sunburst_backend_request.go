// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SunburstBackendRequest - Test-only generated schema for the dashboard versioning proof of concept.
type SunburstBackendRequest struct {
	SunburstBackendRequest20270101OneOf           *SunburstBackendRequest20270101OneOf
	SunburstBackendRequest20270101OneOf4462422560 *SunburstBackendRequest20270101OneOf4462422560
	SunburstBackendRequest20270101OneOf4462423840 *SunburstBackendRequest20270101OneOf4462423840
	SunburstBackendRequest20270101OneOf4462424800 *SunburstBackendRequest20270101OneOf4462424800

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// SunburstBackendRequest20270101OneOfAsSunburstBackendRequest is a convenience function that returns SunburstBackendRequest20270101OneOf wrapped in SunburstBackendRequest.
func SunburstBackendRequest20270101OneOfAsSunburstBackendRequest(v *SunburstBackendRequest20270101OneOf) SunburstBackendRequest {
	return SunburstBackendRequest{SunburstBackendRequest20270101OneOf: v}
}

// SunburstBackendRequest20270101OneOf4462422560AsSunburstBackendRequest is a convenience function that returns SunburstBackendRequest20270101OneOf4462422560 wrapped in SunburstBackendRequest.
func SunburstBackendRequest20270101OneOf4462422560AsSunburstBackendRequest(v *SunburstBackendRequest20270101OneOf4462422560) SunburstBackendRequest {
	return SunburstBackendRequest{SunburstBackendRequest20270101OneOf4462422560: v}
}

// SunburstBackendRequest20270101OneOf4462423840AsSunburstBackendRequest is a convenience function that returns SunburstBackendRequest20270101OneOf4462423840 wrapped in SunburstBackendRequest.
func SunburstBackendRequest20270101OneOf4462423840AsSunburstBackendRequest(v *SunburstBackendRequest20270101OneOf4462423840) SunburstBackendRequest {
	return SunburstBackendRequest{SunburstBackendRequest20270101OneOf4462423840: v}
}

// SunburstBackendRequest20270101OneOf4462424800AsSunburstBackendRequest is a convenience function that returns SunburstBackendRequest20270101OneOf4462424800 wrapped in SunburstBackendRequest.
func SunburstBackendRequest20270101OneOf4462424800AsSunburstBackendRequest(v *SunburstBackendRequest20270101OneOf4462424800) SunburstBackendRequest {
	return SunburstBackendRequest{SunburstBackendRequest20270101OneOf4462424800: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *SunburstBackendRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SunburstBackendRequest20270101OneOf
	err = datadog.Unmarshal(data, &obj.SunburstBackendRequest20270101OneOf)
	if err == nil {
		if obj.SunburstBackendRequest20270101OneOf != nil && obj.SunburstBackendRequest20270101OneOf.UnparsedObject == nil {
			jsonSunburstBackendRequest20270101OneOf, _ := datadog.Marshal(obj.SunburstBackendRequest20270101OneOf)
			if string(jsonSunburstBackendRequest20270101OneOf) == "{}" { // empty struct
				obj.SunburstBackendRequest20270101OneOf = nil
			} else {
				match++
			}
		} else {
			obj.SunburstBackendRequest20270101OneOf = nil
		}
	} else {
		obj.SunburstBackendRequest20270101OneOf = nil
	}

	// try to unmarshal data into SunburstBackendRequest20270101OneOf4462422560
	err = datadog.Unmarshal(data, &obj.SunburstBackendRequest20270101OneOf4462422560)
	if err == nil {
		if obj.SunburstBackendRequest20270101OneOf4462422560 != nil && obj.SunburstBackendRequest20270101OneOf4462422560.UnparsedObject == nil {
			jsonSunburstBackendRequest20270101OneOf4462422560, _ := datadog.Marshal(obj.SunburstBackendRequest20270101OneOf4462422560)
			if string(jsonSunburstBackendRequest20270101OneOf4462422560) == "{}" { // empty struct
				obj.SunburstBackendRequest20270101OneOf4462422560 = nil
			} else {
				match++
			}
		} else {
			obj.SunburstBackendRequest20270101OneOf4462422560 = nil
		}
	} else {
		obj.SunburstBackendRequest20270101OneOf4462422560 = nil
	}

	// try to unmarshal data into SunburstBackendRequest20270101OneOf4462423840
	err = datadog.Unmarshal(data, &obj.SunburstBackendRequest20270101OneOf4462423840)
	if err == nil {
		if obj.SunburstBackendRequest20270101OneOf4462423840 != nil && obj.SunburstBackendRequest20270101OneOf4462423840.UnparsedObject == nil {
			jsonSunburstBackendRequest20270101OneOf4462423840, _ := datadog.Marshal(obj.SunburstBackendRequest20270101OneOf4462423840)
			if string(jsonSunburstBackendRequest20270101OneOf4462423840) == "{}" { // empty struct
				obj.SunburstBackendRequest20270101OneOf4462423840 = nil
			} else {
				match++
			}
		} else {
			obj.SunburstBackendRequest20270101OneOf4462423840 = nil
		}
	} else {
		obj.SunburstBackendRequest20270101OneOf4462423840 = nil
	}

	// try to unmarshal data into SunburstBackendRequest20270101OneOf4462424800
	err = datadog.Unmarshal(data, &obj.SunburstBackendRequest20270101OneOf4462424800)
	if err == nil {
		if obj.SunburstBackendRequest20270101OneOf4462424800 != nil && obj.SunburstBackendRequest20270101OneOf4462424800.UnparsedObject == nil {
			jsonSunburstBackendRequest20270101OneOf4462424800, _ := datadog.Marshal(obj.SunburstBackendRequest20270101OneOf4462424800)
			if string(jsonSunburstBackendRequest20270101OneOf4462424800) == "{}" { // empty struct
				obj.SunburstBackendRequest20270101OneOf4462424800 = nil
			} else {
				match++
			}
		} else {
			obj.SunburstBackendRequest20270101OneOf4462424800 = nil
		}
	} else {
		obj.SunburstBackendRequest20270101OneOf4462424800 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.SunburstBackendRequest20270101OneOf = nil
		obj.SunburstBackendRequest20270101OneOf4462422560 = nil
		obj.SunburstBackendRequest20270101OneOf4462423840 = nil
		obj.SunburstBackendRequest20270101OneOf4462424800 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj SunburstBackendRequest) MarshalJSON() ([]byte, error) {
	if obj.SunburstBackendRequest20270101OneOf != nil {
		return datadog.Marshal(&obj.SunburstBackendRequest20270101OneOf)
	}

	if obj.SunburstBackendRequest20270101OneOf4462422560 != nil {
		return datadog.Marshal(&obj.SunburstBackendRequest20270101OneOf4462422560)
	}

	if obj.SunburstBackendRequest20270101OneOf4462423840 != nil {
		return datadog.Marshal(&obj.SunburstBackendRequest20270101OneOf4462423840)
	}

	if obj.SunburstBackendRequest20270101OneOf4462424800 != nil {
		return datadog.Marshal(&obj.SunburstBackendRequest20270101OneOf4462424800)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *SunburstBackendRequest) GetActualInstance() interface{} {
	if obj.SunburstBackendRequest20270101OneOf != nil {
		return obj.SunburstBackendRequest20270101OneOf
	}

	if obj.SunburstBackendRequest20270101OneOf4462422560 != nil {
		return obj.SunburstBackendRequest20270101OneOf4462422560
	}

	if obj.SunburstBackendRequest20270101OneOf4462423840 != nil {
		return obj.SunburstBackendRequest20270101OneOf4462423840
	}

	if obj.SunburstBackendRequest20270101OneOf4462424800 != nil {
		return obj.SunburstBackendRequest20270101OneOf4462424800
	}

	// all schemas are nil
	return nil
}
