// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomTopologyMapInlineNodeColumn - Test-only generated schema for the dashboard versioning proof of concept.
type CustomTopologyMapInlineNodeColumn struct {
	CustomTopologyMapInlineNodeIdColumn_20270101     *CustomTopologyMapInlineNodeIdColumn
	CustomTopologyMapInlineNodeStringColumn_20270101 *CustomTopologyMapInlineNodeStringColumn
	CustomTopologyMapInlineNodeValueColumn_20270101  *CustomTopologyMapInlineNodeValueColumn

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// CustomTopologyMapInlineNodeIdColumn_20270101AsCustomTopologyMapInlineNodeColumn is a convenience function that returns CustomTopologyMapInlineNodeIdColumn wrapped in CustomTopologyMapInlineNodeColumn.
func CustomTopologyMapInlineNodeIdColumn_20270101AsCustomTopologyMapInlineNodeColumn(v *CustomTopologyMapInlineNodeIdColumn) CustomTopologyMapInlineNodeColumn {
	return CustomTopologyMapInlineNodeColumn{CustomTopologyMapInlineNodeIdColumn_20270101: v}
}

// CustomTopologyMapInlineNodeStringColumn_20270101AsCustomTopologyMapInlineNodeColumn is a convenience function that returns CustomTopologyMapInlineNodeStringColumn wrapped in CustomTopologyMapInlineNodeColumn.
func CustomTopologyMapInlineNodeStringColumn_20270101AsCustomTopologyMapInlineNodeColumn(v *CustomTopologyMapInlineNodeStringColumn) CustomTopologyMapInlineNodeColumn {
	return CustomTopologyMapInlineNodeColumn{CustomTopologyMapInlineNodeStringColumn_20270101: v}
}

// CustomTopologyMapInlineNodeValueColumn_20270101AsCustomTopologyMapInlineNodeColumn is a convenience function that returns CustomTopologyMapInlineNodeValueColumn wrapped in CustomTopologyMapInlineNodeColumn.
func CustomTopologyMapInlineNodeValueColumn_20270101AsCustomTopologyMapInlineNodeColumn(v *CustomTopologyMapInlineNodeValueColumn) CustomTopologyMapInlineNodeColumn {
	return CustomTopologyMapInlineNodeColumn{CustomTopologyMapInlineNodeValueColumn_20270101: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *CustomTopologyMapInlineNodeColumn) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CustomTopologyMapInlineNodeIdColumn_20270101
	err = datadog.Unmarshal(data, &obj.CustomTopologyMapInlineNodeIdColumn_20270101)
	if err == nil {
		if obj.CustomTopologyMapInlineNodeIdColumn_20270101 != nil && obj.CustomTopologyMapInlineNodeIdColumn_20270101.UnparsedObject == nil {
			jsonCustomTopologyMapInlineNodeIdColumn_20270101, _ := datadog.Marshal(obj.CustomTopologyMapInlineNodeIdColumn_20270101)
			if string(jsonCustomTopologyMapInlineNodeIdColumn_20270101) == "{}" { // empty struct
				obj.CustomTopologyMapInlineNodeIdColumn_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.CustomTopologyMapInlineNodeIdColumn_20270101 = nil
		}
	} else {
		obj.CustomTopologyMapInlineNodeIdColumn_20270101 = nil
	}

	// try to unmarshal data into CustomTopologyMapInlineNodeStringColumn_20270101
	err = datadog.Unmarshal(data, &obj.CustomTopologyMapInlineNodeStringColumn_20270101)
	if err == nil {
		if obj.CustomTopologyMapInlineNodeStringColumn_20270101 != nil && obj.CustomTopologyMapInlineNodeStringColumn_20270101.UnparsedObject == nil {
			jsonCustomTopologyMapInlineNodeStringColumn_20270101, _ := datadog.Marshal(obj.CustomTopologyMapInlineNodeStringColumn_20270101)
			if string(jsonCustomTopologyMapInlineNodeStringColumn_20270101) == "{}" { // empty struct
				obj.CustomTopologyMapInlineNodeStringColumn_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.CustomTopologyMapInlineNodeStringColumn_20270101 = nil
		}
	} else {
		obj.CustomTopologyMapInlineNodeStringColumn_20270101 = nil
	}

	// try to unmarshal data into CustomTopologyMapInlineNodeValueColumn_20270101
	err = datadog.Unmarshal(data, &obj.CustomTopologyMapInlineNodeValueColumn_20270101)
	if err == nil {
		if obj.CustomTopologyMapInlineNodeValueColumn_20270101 != nil && obj.CustomTopologyMapInlineNodeValueColumn_20270101.UnparsedObject == nil {
			jsonCustomTopologyMapInlineNodeValueColumn_20270101, _ := datadog.Marshal(obj.CustomTopologyMapInlineNodeValueColumn_20270101)
			if string(jsonCustomTopologyMapInlineNodeValueColumn_20270101) == "{}" { // empty struct
				obj.CustomTopologyMapInlineNodeValueColumn_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.CustomTopologyMapInlineNodeValueColumn_20270101 = nil
		}
	} else {
		obj.CustomTopologyMapInlineNodeValueColumn_20270101 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.CustomTopologyMapInlineNodeIdColumn_20270101 = nil
		obj.CustomTopologyMapInlineNodeStringColumn_20270101 = nil
		obj.CustomTopologyMapInlineNodeValueColumn_20270101 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj CustomTopologyMapInlineNodeColumn) MarshalJSON() ([]byte, error) {
	if obj.CustomTopologyMapInlineNodeIdColumn_20270101 != nil {
		return datadog.Marshal(&obj.CustomTopologyMapInlineNodeIdColumn_20270101)
	}

	if obj.CustomTopologyMapInlineNodeStringColumn_20270101 != nil {
		return datadog.Marshal(&obj.CustomTopologyMapInlineNodeStringColumn_20270101)
	}

	if obj.CustomTopologyMapInlineNodeValueColumn_20270101 != nil {
		return datadog.Marshal(&obj.CustomTopologyMapInlineNodeValueColumn_20270101)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *CustomTopologyMapInlineNodeColumn) GetActualInstance() interface{} {
	if obj.CustomTopologyMapInlineNodeIdColumn_20270101 != nil {
		return obj.CustomTopologyMapInlineNodeIdColumn_20270101
	}

	if obj.CustomTopologyMapInlineNodeStringColumn_20270101 != nil {
		return obj.CustomTopologyMapInlineNodeStringColumn_20270101
	}

	if obj.CustomTopologyMapInlineNodeValueColumn_20270101 != nil {
		return obj.CustomTopologyMapInlineNodeValueColumn_20270101
	}

	// all schemas are nil
	return nil
}
