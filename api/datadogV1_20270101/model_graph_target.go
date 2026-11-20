// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GraphTarget - Test-only generated schema for the dashboard versioning proof of concept.
type GraphTarget struct {
	GraphTargetNode_20270101 *GraphTargetNode
	GraphTargetPath_20270101 *GraphTargetPath

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// GraphTargetNode_20270101AsGraphTarget is a convenience function that returns GraphTargetNode wrapped in GraphTarget.
func GraphTargetNode_20270101AsGraphTarget(v *GraphTargetNode) GraphTarget {
	return GraphTarget{GraphTargetNode_20270101: v}
}

// GraphTargetPath_20270101AsGraphTarget is a convenience function that returns GraphTargetPath wrapped in GraphTarget.
func GraphTargetPath_20270101AsGraphTarget(v *GraphTargetPath) GraphTarget {
	return GraphTarget{GraphTargetPath_20270101: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *GraphTarget) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into GraphTargetNode_20270101
	err = datadog.Unmarshal(data, &obj.GraphTargetNode_20270101)
	if err == nil {
		if obj.GraphTargetNode_20270101 != nil && obj.GraphTargetNode_20270101.UnparsedObject == nil {
			jsonGraphTargetNode_20270101, _ := datadog.Marshal(obj.GraphTargetNode_20270101)
			if string(jsonGraphTargetNode_20270101) == "{}" { // empty struct
				obj.GraphTargetNode_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.GraphTargetNode_20270101 = nil
		}
	} else {
		obj.GraphTargetNode_20270101 = nil
	}

	// try to unmarshal data into GraphTargetPath_20270101
	err = datadog.Unmarshal(data, &obj.GraphTargetPath_20270101)
	if err == nil {
		if obj.GraphTargetPath_20270101 != nil && obj.GraphTargetPath_20270101.UnparsedObject == nil {
			jsonGraphTargetPath_20270101, _ := datadog.Marshal(obj.GraphTargetPath_20270101)
			if string(jsonGraphTargetPath_20270101) == "{}" { // empty struct
				obj.GraphTargetPath_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.GraphTargetPath_20270101 = nil
		}
	} else {
		obj.GraphTargetPath_20270101 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.GraphTargetNode_20270101 = nil
		obj.GraphTargetPath_20270101 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj GraphTarget) MarshalJSON() ([]byte, error) {
	if obj.GraphTargetNode_20270101 != nil {
		return datadog.Marshal(&obj.GraphTargetNode_20270101)
	}

	if obj.GraphTargetPath_20270101 != nil {
		return datadog.Marshal(&obj.GraphTargetPath_20270101)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *GraphTarget) GetActualInstance() interface{} {
	if obj.GraphTargetNode_20270101 != nil {
		return obj.GraphTargetNode_20270101
	}

	if obj.GraphTargetPath_20270101 != nil {
		return obj.GraphTargetPath_20270101
	}

	// all schemas are nil
	return nil
}
