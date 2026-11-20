// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// Hour - Test-only generated schema for the dashboard versioning proof of concept.
type Hour struct {
	Hour12Format_20270101 *Hour12Format
	Hour24Format_20270101 *Hour24Format

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// Hour12Format_20270101AsHour is a convenience function that returns Hour12Format wrapped in Hour.
func Hour12Format_20270101AsHour(v *Hour12Format) Hour {
	return Hour{Hour12Format_20270101: v}
}

// Hour24Format_20270101AsHour is a convenience function that returns Hour24Format wrapped in Hour.
func Hour24Format_20270101AsHour(v *Hour24Format) Hour {
	return Hour{Hour24Format_20270101: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *Hour) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Hour12Format_20270101
	err = datadog.Unmarshal(data, &obj.Hour12Format_20270101)
	if err == nil {
		if obj.Hour12Format_20270101 != nil {
			jsonHour12Format_20270101, _ := datadog.Marshal(obj.Hour12Format_20270101)
			if string(jsonHour12Format_20270101) == "{}" && string(data) != "{}" { // empty struct
				obj.Hour12Format_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.Hour12Format_20270101 = nil
		}
	} else {
		obj.Hour12Format_20270101 = nil
	}

	// try to unmarshal data into Hour24Format_20270101
	err = datadog.Unmarshal(data, &obj.Hour24Format_20270101)
	if err == nil {
		if obj.Hour24Format_20270101 != nil {
			jsonHour24Format_20270101, _ := datadog.Marshal(obj.Hour24Format_20270101)
			if string(jsonHour24Format_20270101) == "{}" && string(data) != "{}" { // empty struct
				obj.Hour24Format_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.Hour24Format_20270101 = nil
		}
	} else {
		obj.Hour24Format_20270101 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.Hour12Format_20270101 = nil
		obj.Hour24Format_20270101 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj Hour) MarshalJSON() ([]byte, error) {
	if obj.Hour12Format_20270101 != nil {
		return datadog.Marshal(&obj.Hour12Format_20270101)
	}

	if obj.Hour24Format_20270101 != nil {
		return datadog.Marshal(&obj.Hour24Format_20270101)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *Hour) GetActualInstance() interface{} {
	if obj.Hour12Format_20270101 != nil {
		return obj.Hour12Format_20270101
	}

	if obj.Hour24Format_20270101 != nil {
		return obj.Hour24Format_20270101
	}

	// all schemas are nil
	return nil
}
