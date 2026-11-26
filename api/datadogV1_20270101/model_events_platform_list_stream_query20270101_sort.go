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


// EventsPlatformListStreamQuery20270101Sort - Test-only generated schema for the dashboard versioning proof of concept.
type EventsPlatformListStreamQuery20270101Sort struct {
	ListStreamSort_20270101 *ListStreamSort
	[]ListStreamSort *[]ListStreamSort

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ListStreamSort_20270101AsEventsPlatformListStreamQuery20270101Sort is a convenience function that returns ListStreamSort wrapped in EventsPlatformListStreamQuery20270101Sort.
func ListStreamSort_20270101AsEventsPlatformListStreamQuery20270101Sort(v *ListStreamSort) EventsPlatformListStreamQuery20270101Sort {
	return EventsPlatformListStreamQuery20270101Sort{ListStreamSort_20270101: v}
}

// []ListStreamSortAsEventsPlatformListStreamQuery20270101Sort is a convenience function that returns []ListStreamSort wrapped in EventsPlatformListStreamQuery20270101Sort.
func []ListStreamSortAsEventsPlatformListStreamQuery20270101Sort(v *[]ListStreamSort) EventsPlatformListStreamQuery20270101Sort {
	return EventsPlatformListStreamQuery20270101Sort{[]ListStreamSort: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *EventsPlatformListStreamQuery20270101Sort) UnmarshalJSON(data []byte) error {
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

	// try to unmarshal data into []ListStreamSort
	err = datadog.Unmarshal(data, &obj.[]ListStreamSort)
	if err == nil {
		if obj.[]ListStreamSort != nil {
			json[]ListStreamSort, _ := datadog.Marshal(obj.[]ListStreamSort)
			if string(json[]ListStreamSort) == "{}" && string(data) != "{}"  { // empty struct
				obj.[]ListStreamSort = nil
			} else {
				match++
			}
		} else {
			obj.[]ListStreamSort = nil
		}
	} else {
		obj.[]ListStreamSort = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ListStreamSort_20270101 = nil
		obj.[]ListStreamSort = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj EventsPlatformListStreamQuery20270101Sort) MarshalJSON() ([]byte, error) {
	if obj.ListStreamSort_20270101 != nil {
		return datadog.Marshal(&obj.ListStreamSort_20270101)
	}


	if obj.[]ListStreamSort != nil {
		return datadog.Marshal(&obj.[]ListStreamSort)
	}


	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *EventsPlatformListStreamQuery20270101Sort) GetActualInstance() (interface{}) {
	if obj.ListStreamSort_20270101 != nil {
		return obj.ListStreamSort_20270101
	}


	if obj.[]ListStreamSort != nil {
		return obj.[]ListStreamSort
	}


	// all schemas are nil
	return nil
}
