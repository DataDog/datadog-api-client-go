// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CILogAttributeValue - A flat additional log attribute. Objects and arrays are not accepted.
type CILogAttributeValue struct {
	String  *string
	Float64 *float64
	Bool    *bool

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// StringAsCILogAttributeValue is a convenience function that returns string wrapped in CILogAttributeValue.
func StringAsCILogAttributeValue(v *string) CILogAttributeValue {
	return CILogAttributeValue{String: v}
}

// Float64AsCILogAttributeValue is a convenience function that returns float64 wrapped in CILogAttributeValue.
func Float64AsCILogAttributeValue(v *float64) CILogAttributeValue {
	return CILogAttributeValue{Float64: v}
}

// BoolAsCILogAttributeValue is a convenience function that returns bool wrapped in CILogAttributeValue.
func BoolAsCILogAttributeValue(v *bool) CILogAttributeValue {
	return CILogAttributeValue{Bool: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *CILogAttributeValue) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into String
	err = datadog.Unmarshal(data, &obj.String)
	if err == nil {
		if obj.String != nil {
			jsonString, _ := datadog.Marshal(obj.String)
			if string(jsonString) == "{}" { // empty struct
				obj.String = nil
			} else {
				match++
			}
		} else {
			obj.String = nil
		}
	} else {
		obj.String = nil
	}

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

	// try to unmarshal data into Bool
	err = datadog.Unmarshal(data, &obj.Bool)
	if err == nil {
		if obj.Bool != nil {
			jsonBool, _ := datadog.Marshal(obj.Bool)
			if string(jsonBool) == "{}" { // empty struct
				obj.Bool = nil
			} else {
				match++
			}
		} else {
			obj.Bool = nil
		}
	} else {
		obj.Bool = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.String = nil
		obj.Float64 = nil
		obj.Bool = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj CILogAttributeValue) MarshalJSON() ([]byte, error) {
	if obj.String != nil {
		return datadog.Marshal(&obj.String)
	}

	if obj.Float64 != nil {
		return datadog.Marshal(&obj.Float64)
	}

	if obj.Bool != nil {
		return datadog.Marshal(&obj.Bool)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *CILogAttributeValue) GetActualInstance() interface{} {
	if obj.String != nil {
		return obj.String
	}

	if obj.Float64 != nil {
		return obj.Float64
	}

	if obj.Bool != nil {
		return obj.Bool
	}

	// all schemas are nil
	return nil
}
