// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GuidedTableNumberFormatPrecision - Number of digits after the decimal point. Use `*` for full precision, `integer` to show full integer values, or omit for automatic precision.
type GuidedTableNumberFormatPrecision struct {
	String *string
	Int32  *int32

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// StringAsGuidedTableNumberFormatPrecision is a convenience function that returns string wrapped in GuidedTableNumberFormatPrecision.
func StringAsGuidedTableNumberFormatPrecision(v *string) GuidedTableNumberFormatPrecision {
	return GuidedTableNumberFormatPrecision{String: v}
}

// Int32AsGuidedTableNumberFormatPrecision is a convenience function that returns int32 wrapped in GuidedTableNumberFormatPrecision.
func Int32AsGuidedTableNumberFormatPrecision(v *int32) GuidedTableNumberFormatPrecision {
	return GuidedTableNumberFormatPrecision{Int32: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *GuidedTableNumberFormatPrecision) UnmarshalJSON(data []byte) error {
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

	// try to unmarshal data into Int32
	err = datadog.Unmarshal(data, &obj.Int32)
	if err == nil {
		if obj.Int32 != nil {
			jsonInt32, _ := datadog.Marshal(obj.Int32)
			if string(jsonInt32) == "{}" { // empty struct
				obj.Int32 = nil
			} else {
				match++
			}
		} else {
			obj.Int32 = nil
		}
	} else {
		obj.Int32 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.String = nil
		obj.Int32 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj GuidedTableNumberFormatPrecision) MarshalJSON() ([]byte, error) {
	if obj.String != nil {
		return datadog.Marshal(&obj.String)
	}

	if obj.Int32 != nil {
		return datadog.Marshal(&obj.Int32)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *GuidedTableNumberFormatPrecision) GetActualInstance() interface{} {
	if obj.String != nil {
		return obj.String
	}

	if obj.Int32 != nil {
		return obj.Int32
	}

	// all schemas are nil
	return nil
}
