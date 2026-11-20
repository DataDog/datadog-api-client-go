// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ReferenceTableColumn - Test-only generated schema for the dashboard versioning proof of concept.
type ReferenceTableColumn struct {
	StringReferenceTableColumn_20270101 *StringReferenceTableColumn
	NumberReferenceTableColumn_20270101 *NumberReferenceTableColumn

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// StringReferenceTableColumn_20270101AsReferenceTableColumn is a convenience function that returns StringReferenceTableColumn wrapped in ReferenceTableColumn.
func StringReferenceTableColumn_20270101AsReferenceTableColumn(v *StringReferenceTableColumn) ReferenceTableColumn {
	return ReferenceTableColumn{StringReferenceTableColumn_20270101: v}
}

// NumberReferenceTableColumn_20270101AsReferenceTableColumn is a convenience function that returns NumberReferenceTableColumn wrapped in ReferenceTableColumn.
func NumberReferenceTableColumn_20270101AsReferenceTableColumn(v *NumberReferenceTableColumn) ReferenceTableColumn {
	return ReferenceTableColumn{NumberReferenceTableColumn_20270101: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ReferenceTableColumn) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into StringReferenceTableColumn_20270101
	err = datadog.Unmarshal(data, &obj.StringReferenceTableColumn_20270101)
	if err == nil {
		if obj.StringReferenceTableColumn_20270101 != nil && obj.StringReferenceTableColumn_20270101.UnparsedObject == nil {
			jsonStringReferenceTableColumn_20270101, _ := datadog.Marshal(obj.StringReferenceTableColumn_20270101)
			if string(jsonStringReferenceTableColumn_20270101) == "{}" { // empty struct
				obj.StringReferenceTableColumn_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.StringReferenceTableColumn_20270101 = nil
		}
	} else {
		obj.StringReferenceTableColumn_20270101 = nil
	}

	// try to unmarshal data into NumberReferenceTableColumn_20270101
	err = datadog.Unmarshal(data, &obj.NumberReferenceTableColumn_20270101)
	if err == nil {
		if obj.NumberReferenceTableColumn_20270101 != nil && obj.NumberReferenceTableColumn_20270101.UnparsedObject == nil {
			jsonNumberReferenceTableColumn_20270101, _ := datadog.Marshal(obj.NumberReferenceTableColumn_20270101)
			if string(jsonNumberReferenceTableColumn_20270101) == "{}" { // empty struct
				obj.NumberReferenceTableColumn_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.NumberReferenceTableColumn_20270101 = nil
		}
	} else {
		obj.NumberReferenceTableColumn_20270101 = nil
	}

	if match > 1 {
		// more than one variant matched, so tell them apart by the items of the arrays they hold
		if obj.StringReferenceTableColumn_20270101 != nil {
			allFunctionsUnparsed := len(obj.StringReferenceTableColumn_20270101.Functions) > 0
			for _, item := range obj.StringReferenceTableColumn_20270101.Functions {
				if item.UnparsedObject == nil {
					allFunctionsUnparsed = false
					break
				}
			}
			if allFunctionsUnparsed {
				obj.StringReferenceTableColumn_20270101 = nil
				match--
			}
		}
		if obj.NumberReferenceTableColumn_20270101 != nil {
			allFunctionsUnparsed := len(obj.NumberReferenceTableColumn_20270101.Functions) > 0
			for _, item := range obj.NumberReferenceTableColumn_20270101.Functions {
				if item.UnparsedObject == nil {
					allFunctionsUnparsed = false
					break
				}
			}
			if allFunctionsUnparsed {
				obj.NumberReferenceTableColumn_20270101 = nil
				match--
			}
		}
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.StringReferenceTableColumn_20270101 = nil
		obj.NumberReferenceTableColumn_20270101 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ReferenceTableColumn) MarshalJSON() ([]byte, error) {
	if obj.StringReferenceTableColumn_20270101 != nil {
		return datadog.Marshal(&obj.StringReferenceTableColumn_20270101)
	}

	if obj.NumberReferenceTableColumn_20270101 != nil {
		return datadog.Marshal(&obj.NumberReferenceTableColumn_20270101)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ReferenceTableColumn) GetActualInstance() interface{} {
	if obj.StringReferenceTableColumn_20270101 != nil {
		return obj.StringReferenceTableColumn_20270101
	}

	if obj.NumberReferenceTableColumn_20270101 != nil {
		return obj.NumberReferenceTableColumn_20270101
	}

	// all schemas are nil
	return nil
}
