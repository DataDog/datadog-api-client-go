// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GuidedTableColumnFunctionArgsItem -
type GuidedTableColumnFunctionArgsItem struct {
	String *string
	Float  *float
	Bool   *bool

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// StringAsGuidedTableColumnFunctionArgsItem is a convenience function that returns string wrapped in GuidedTableColumnFunctionArgsItem.
func StringAsGuidedTableColumnFunctionArgsItem(v *string) GuidedTableColumnFunctionArgsItem {
	return GuidedTableColumnFunctionArgsItem{String: v}
}

// FloatAsGuidedTableColumnFunctionArgsItem is a convenience function that returns float wrapped in GuidedTableColumnFunctionArgsItem.
func FloatAsGuidedTableColumnFunctionArgsItem(v *float) GuidedTableColumnFunctionArgsItem {
	return GuidedTableColumnFunctionArgsItem{Float: v}
}

// BoolAsGuidedTableColumnFunctionArgsItem is a convenience function that returns bool wrapped in GuidedTableColumnFunctionArgsItem.
func BoolAsGuidedTableColumnFunctionArgsItem(v *bool) GuidedTableColumnFunctionArgsItem {
	return GuidedTableColumnFunctionArgsItem{Bool: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *GuidedTableColumnFunctionArgsItem) UnmarshalJSON(data []byte) error {
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

	// try to unmarshal data into Float
	err = datadog.Unmarshal(data, &obj.Float)
	if err == nil {
		if obj.Float != nil {
			jsonFloat, _ := datadog.Marshal(obj.Float)
			if string(jsonFloat) == "{}" { // empty struct
				obj.Float = nil
			} else {
				match++
			}
		} else {
			obj.Float = nil
		}
	} else {
		obj.Float = nil
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
		obj.Float = nil
		obj.Bool = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj GuidedTableColumnFunctionArgsItem) MarshalJSON() ([]byte, error) {
	if obj.String != nil {
		return datadog.Marshal(&obj.String)
	}

	if obj.Float != nil {
		return datadog.Marshal(&obj.Float)
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
func (obj *GuidedTableColumnFunctionArgsItem) GetActualInstance() interface{} {
	if obj.String != nil {
		return obj.String
	}

	if obj.Float != nil {
		return obj.Float
	}

	if obj.Bool != nil {
		return obj.Bool
	}

	// all schemas are nil
	return nil
}
