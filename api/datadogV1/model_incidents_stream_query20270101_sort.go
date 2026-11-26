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


// IncidentsStreamQuery20270101Sort - Test-only generated schema for the dashboard versioning proof of concept.
type IncidentsStreamQuery20270101Sort struct {
	ListStreamSort_20270101 *ListStreamSort_20270101
	[]ListStreamSort_20270101 *[]ListStreamSort_20270101

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ListStreamSort_20270101AsIncidentsStreamQuery20270101Sort is a convenience function that returns ListStreamSort_20270101 wrapped in IncidentsStreamQuery20270101Sort.
func ListStreamSort_20270101AsIncidentsStreamQuery20270101Sort(v *ListStreamSort_20270101) IncidentsStreamQuery20270101Sort {
	return IncidentsStreamQuery20270101Sort{ListStreamSort_20270101: v}
}

// []ListStreamSort_20270101AsIncidentsStreamQuery20270101Sort is a convenience function that returns []ListStreamSort_20270101 wrapped in IncidentsStreamQuery20270101Sort.
func []ListStreamSort_20270101AsIncidentsStreamQuery20270101Sort(v *[]ListStreamSort_20270101) IncidentsStreamQuery20270101Sort {
	return IncidentsStreamQuery20270101Sort{[]ListStreamSort_20270101: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *IncidentsStreamQuery20270101Sort) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ListStreamSort_20270101
	err = datadog.Unmarshal(data, &obj.ListStreamSort_20270101)
	if err == nil {
		if obj.ListStreamSort_20270101 != nil && obj.ListStreamSort_20270101.UnparsedObject == nil {
			jsonListStreamSort_20270101, _ := datadog.Marshal(obj.ListStreamSort_20270101)
			if string(jsonListStreamSort_20270101) == "{}" { // empty struct
				obj.ListStreamSort_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.ListStreamSort_20270101 = nil
		}
	} else {
		obj.ListStreamSort_20270101 = nil
	}

	// try to unmarshal data into []ListStreamSort_20270101
	err = datadog.Unmarshal(data, &obj.[]ListStreamSort_20270101)
	if err == nil {
		if obj.[]ListStreamSort_20270101 != nil {
			json[]ListStreamSort_20270101, _ := datadog.Marshal(obj.[]ListStreamSort_20270101)
			if string(json[]ListStreamSort_20270101) == "{}" && string(data) != "{}"  { // empty struct
				obj.[]ListStreamSort_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.[]ListStreamSort_20270101 = nil
		}
	} else {
		obj.[]ListStreamSort_20270101 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ListStreamSort_20270101 = nil
		obj.[]ListStreamSort_20270101 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj IncidentsStreamQuery20270101Sort) MarshalJSON() ([]byte, error) {
	if obj.ListStreamSort_20270101 != nil {
		return datadog.Marshal(&obj.ListStreamSort_20270101)
	}


	if obj.[]ListStreamSort_20270101 != nil {
		return datadog.Marshal(&obj.[]ListStreamSort_20270101)
	}


	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *IncidentsStreamQuery20270101Sort) GetActualInstance() (interface{}) {
	if obj.ListStreamSort_20270101 != nil {
		return obj.ListStreamSort_20270101
	}


	if obj.[]ListStreamSort_20270101 != nil {
		return obj.[]ListStreamSort_20270101
	}


	// all schemas are nil
	return nil
}
