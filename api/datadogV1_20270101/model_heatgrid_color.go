// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.


package datadogV1_20270101

import (
	_io "io"
	"time"
	"github.com/google/uuid"
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"

)


// HeatgridColor - Test-only generated schema for the dashboard versioning proof of concept.
type HeatgridColor struct {
	String *string
	[]interface{} *[]interface{}

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// StringAsHeatgridColor is a convenience function that returns string wrapped in HeatgridColor.
func StringAsHeatgridColor(v *string) HeatgridColor {
	return HeatgridColor{String: v}
}

// []interface{}AsHeatgridColor is a convenience function that returns []interface{} wrapped in HeatgridColor.
func []interface{}AsHeatgridColor(v *[]interface{}) HeatgridColor {
	return HeatgridColor{[]interface{}: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *HeatgridColor) UnmarshalJSON(data []byte) error {
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

	// try to unmarshal data into []interface{}
	err = datadog.Unmarshal(data, &obj.[]interface{})
	if err == nil {
		if obj.[]interface{} != nil {
			json[]interface{}, _ := datadog.Marshal(obj.[]interface{})
			if string(json[]interface{}) == "{}" && string(data) != "{}"  { // empty struct
				obj.[]interface{} = nil
			} else {
				match++
			}
		} else {
			obj.[]interface{} = nil
		}
	} else {
		obj.[]interface{} = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.String = nil
		obj.[]interface{} = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj HeatgridColor) MarshalJSON() ([]byte, error) {
	if obj.String != nil {
		return datadog.Marshal(&obj.String)
	}


	if obj.[]interface{} != nil {
		return datadog.Marshal(&obj.[]interface{})
	}


	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *HeatgridColor) GetActualInstance() (interface{}) {
	if obj.String != nil {
		return obj.String
	}


	if obj.[]interface{} != nil {
		return obj.[]interface{}
	}


	// all schemas are nil
	return nil
}
