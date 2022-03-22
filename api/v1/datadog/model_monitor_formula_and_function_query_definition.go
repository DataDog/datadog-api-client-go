/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package datadog

import (
	"encoding/json"
)

// MonitorFormulaAndFunctionQueryDefinition - A formula and function query.
type MonitorFormulaAndFunctionQueryDefinition struct {
	MonitorFormulaAndFunctionEventQueryDefinition *MonitorFormulaAndFunctionEventQueryDefinition

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// MonitorFormulaAndFunctionEventQueryDefinitionAsMonitorFormulaAndFunctionQueryDefinition is a convenience function that returns MonitorFormulaAndFunctionEventQueryDefinition wrapped in MonitorFormulaAndFunctionQueryDefinition
func MonitorFormulaAndFunctionEventQueryDefinitionAsMonitorFormulaAndFunctionQueryDefinition(v *MonitorFormulaAndFunctionEventQueryDefinition) MonitorFormulaAndFunctionQueryDefinition {
	return MonitorFormulaAndFunctionQueryDefinition{MonitorFormulaAndFunctionEventQueryDefinition: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *MonitorFormulaAndFunctionQueryDefinition) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into MonitorFormulaAndFunctionEventQueryDefinition
	err = json.Unmarshal(data, &dst.MonitorFormulaAndFunctionEventQueryDefinition)
	if err == nil {
		if dst.MonitorFormulaAndFunctionEventQueryDefinition != nil && dst.MonitorFormulaAndFunctionEventQueryDefinition.UnparsedObject == nil {
			jsonMonitorFormulaAndFunctionEventQueryDefinition, _ := json.Marshal(dst.MonitorFormulaAndFunctionEventQueryDefinition)
			if string(jsonMonitorFormulaAndFunctionEventQueryDefinition) == "{}" { // empty struct
				dst.MonitorFormulaAndFunctionEventQueryDefinition = nil
			} else {
				match++
			}
		} else {
			dst.MonitorFormulaAndFunctionEventQueryDefinition = nil
		}
	} else {
		dst.MonitorFormulaAndFunctionEventQueryDefinition = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		dst.MonitorFormulaAndFunctionEventQueryDefinition = nil
		return json.Unmarshal(data, &dst.UnparsedObject)
	} else {
		return nil // exactly one match
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src MonitorFormulaAndFunctionQueryDefinition) MarshalJSON() ([]byte, error) {
	if src.MonitorFormulaAndFunctionEventQueryDefinition != nil {
		return json.Marshal(&src.MonitorFormulaAndFunctionEventQueryDefinition)
	}

	if src.UnparsedObject != nil {
		return json.Marshal(src.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *MonitorFormulaAndFunctionQueryDefinition) GetActualInstance() interface{} {
	if obj.MonitorFormulaAndFunctionEventQueryDefinition != nil {
		return obj.MonitorFormulaAndFunctionEventQueryDefinition
	}

	// all schemas are nil
	return nil
}

type NullableMonitorFormulaAndFunctionQueryDefinition struct {
	value *MonitorFormulaAndFunctionQueryDefinition
	isSet bool
}

func (v NullableMonitorFormulaAndFunctionQueryDefinition) Get() *MonitorFormulaAndFunctionQueryDefinition {
	return v.value
}

func (v *NullableMonitorFormulaAndFunctionQueryDefinition) Set(val *MonitorFormulaAndFunctionQueryDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitorFormulaAndFunctionQueryDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitorFormulaAndFunctionQueryDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitorFormulaAndFunctionQueryDefinition(val *MonitorFormulaAndFunctionQueryDefinition) *NullableMonitorFormulaAndFunctionQueryDefinition {
	return &NullableMonitorFormulaAndFunctionQueryDefinition{value: val, isSet: true}
}

func (v NullableMonitorFormulaAndFunctionQueryDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitorFormulaAndFunctionQueryDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return json.Unmarshal(src, &v.value)
}
