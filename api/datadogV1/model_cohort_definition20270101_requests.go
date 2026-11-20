// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.


package datadogV1

import (
	_io "io"
	"time"
	"github.com/google/uuid"
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"

)


// CohortDefinition20270101Requests - Test-only generated schema for the dashboard versioning proof of concept.
type CohortDefinition20270101Requests struct {
	[]interface{} *[]interface{}
	[]interface{} *[]interface{}

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// []interface{}AsCohortDefinition20270101Requests is a convenience function that returns []interface{} wrapped in CohortDefinition20270101Requests.
func []interface{}AsCohortDefinition20270101Requests(v *[]interface{}) CohortDefinition20270101Requests {
	return CohortDefinition20270101Requests{[]interface{}: v}
}

// []interface{}AsCohortDefinition20270101Requests is a convenience function that returns []interface{} wrapped in CohortDefinition20270101Requests.
func []interface{}AsCohortDefinition20270101Requests(v *[]interface{}) CohortDefinition20270101Requests {
	return CohortDefinition20270101Requests{[]interface{}: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *CohortDefinition20270101Requests) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
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
		obj.[]interface{} = nil
		obj.[]interface{} = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj CohortDefinition20270101Requests) MarshalJSON() ([]byte, error) {
	if obj.[]interface{} != nil {
		return datadog.Marshal(&obj.[]interface{})
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
func (obj *CohortDefinition20270101Requests) GetActualInstance() (interface{}) {
	if obj.[]interface{} != nil {
		return obj.[]interface{}
	}


	if obj.[]interface{} != nil {
		return obj.[]interface{}
	}


	// all schemas are nil
	return nil
}
