// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TopologyMapInlineNodeColumn - Column of the nodes table.
type TopologyMapInlineNodeColumn struct {
	TopologyMapInlineNodeIdColumn     *TopologyMapInlineNodeIdColumn
	TopologyMapInlineNodeStringColumn *TopologyMapInlineNodeStringColumn
	TopologyMapInlineNodeValueColumn  *TopologyMapInlineNodeValueColumn

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// TopologyMapInlineNodeIdColumnAsTopologyMapInlineNodeColumn is a convenience function that returns TopologyMapInlineNodeIdColumn wrapped in TopologyMapInlineNodeColumn.
func TopologyMapInlineNodeIdColumnAsTopologyMapInlineNodeColumn(v *TopologyMapInlineNodeIdColumn) TopologyMapInlineNodeColumn {
	return TopologyMapInlineNodeColumn{TopologyMapInlineNodeIdColumn: v}
}

// TopologyMapInlineNodeStringColumnAsTopologyMapInlineNodeColumn is a convenience function that returns TopologyMapInlineNodeStringColumn wrapped in TopologyMapInlineNodeColumn.
func TopologyMapInlineNodeStringColumnAsTopologyMapInlineNodeColumn(v *TopologyMapInlineNodeStringColumn) TopologyMapInlineNodeColumn {
	return TopologyMapInlineNodeColumn{TopologyMapInlineNodeStringColumn: v}
}

// TopologyMapInlineNodeValueColumnAsTopologyMapInlineNodeColumn is a convenience function that returns TopologyMapInlineNodeValueColumn wrapped in TopologyMapInlineNodeColumn.
func TopologyMapInlineNodeValueColumnAsTopologyMapInlineNodeColumn(v *TopologyMapInlineNodeValueColumn) TopologyMapInlineNodeColumn {
	return TopologyMapInlineNodeColumn{TopologyMapInlineNodeValueColumn: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *TopologyMapInlineNodeColumn) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into TopologyMapInlineNodeIdColumn
	err = datadog.Unmarshal(data, &obj.TopologyMapInlineNodeIdColumn)
	if err == nil {
		if obj.TopologyMapInlineNodeIdColumn != nil && obj.TopologyMapInlineNodeIdColumn.UnparsedObject == nil {
			jsonTopologyMapInlineNodeIdColumn, _ := datadog.Marshal(obj.TopologyMapInlineNodeIdColumn)
			if string(jsonTopologyMapInlineNodeIdColumn) == "{}" { // empty struct
				obj.TopologyMapInlineNodeIdColumn = nil
			} else {
				match++
			}
		} else {
			obj.TopologyMapInlineNodeIdColumn = nil
		}
	} else {
		obj.TopologyMapInlineNodeIdColumn = nil
	}

	// try to unmarshal data into TopologyMapInlineNodeStringColumn
	err = datadog.Unmarshal(data, &obj.TopologyMapInlineNodeStringColumn)
	if err == nil {
		if obj.TopologyMapInlineNodeStringColumn != nil && obj.TopologyMapInlineNodeStringColumn.UnparsedObject == nil {
			jsonTopologyMapInlineNodeStringColumn, _ := datadog.Marshal(obj.TopologyMapInlineNodeStringColumn)
			if string(jsonTopologyMapInlineNodeStringColumn) == "{}" { // empty struct
				obj.TopologyMapInlineNodeStringColumn = nil
			} else {
				match++
			}
		} else {
			obj.TopologyMapInlineNodeStringColumn = nil
		}
	} else {
		obj.TopologyMapInlineNodeStringColumn = nil
	}

	// try to unmarshal data into TopologyMapInlineNodeValueColumn
	err = datadog.Unmarshal(data, &obj.TopologyMapInlineNodeValueColumn)
	if err == nil {
		if obj.TopologyMapInlineNodeValueColumn != nil && obj.TopologyMapInlineNodeValueColumn.UnparsedObject == nil {
			jsonTopologyMapInlineNodeValueColumn, _ := datadog.Marshal(obj.TopologyMapInlineNodeValueColumn)
			if string(jsonTopologyMapInlineNodeValueColumn) == "{}" { // empty struct
				obj.TopologyMapInlineNodeValueColumn = nil
			} else {
				match++
			}
		} else {
			obj.TopologyMapInlineNodeValueColumn = nil
		}
	} else {
		obj.TopologyMapInlineNodeValueColumn = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.TopologyMapInlineNodeIdColumn = nil
		obj.TopologyMapInlineNodeStringColumn = nil
		obj.TopologyMapInlineNodeValueColumn = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj TopologyMapInlineNodeColumn) MarshalJSON() ([]byte, error) {
	if obj.TopologyMapInlineNodeIdColumn != nil {
		return datadog.Marshal(&obj.TopologyMapInlineNodeIdColumn)
	}

	if obj.TopologyMapInlineNodeStringColumn != nil {
		return datadog.Marshal(&obj.TopologyMapInlineNodeStringColumn)
	}

	if obj.TopologyMapInlineNodeValueColumn != nil {
		return datadog.Marshal(&obj.TopologyMapInlineNodeValueColumn)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *TopologyMapInlineNodeColumn) GetActualInstance() interface{} {
	if obj.TopologyMapInlineNodeIdColumn != nil {
		return obj.TopologyMapInlineNodeIdColumn
	}

	if obj.TopologyMapInlineNodeStringColumn != nil {
		return obj.TopologyMapInlineNodeStringColumn
	}

	if obj.TopologyMapInlineNodeValueColumn != nil {
		return obj.TopologyMapInlineNodeValueColumn
	}

	// all schemas are nil
	return nil
}
