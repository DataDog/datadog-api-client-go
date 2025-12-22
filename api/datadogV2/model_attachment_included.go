// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AttachmentIncluded -
type AttachmentIncluded struct {
	User140420082644000 *User140420082644000

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// User140420082644000AsAttachmentIncluded is a convenience function that returns User140420082644000 wrapped in AttachmentIncluded.
func User140420082644000AsAttachmentIncluded(v *User140420082644000) AttachmentIncluded {
	return AttachmentIncluded{User140420082644000: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *AttachmentIncluded) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into User140420082644000
	err = datadog.Unmarshal(data, &obj.User140420082644000)
	if err == nil {
		if obj.User140420082644000 != nil && obj.User140420082644000.UnparsedObject == nil {
			jsonUser140420082644000, _ := datadog.Marshal(obj.User140420082644000)
			if string(jsonUser140420082644000) == "{}" { // empty struct
				obj.User140420082644000 = nil
			} else {
				match++
			}
		} else {
			obj.User140420082644000 = nil
		}
	} else {
		obj.User140420082644000 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.User140420082644000 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj AttachmentIncluded) MarshalJSON() ([]byte, error) {
	if obj.User140420082644000 != nil {
		return datadog.Marshal(&obj.User140420082644000)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *AttachmentIncluded) GetActualInstance() interface{} {
	if obj.User140420082644000 != nil {
		return obj.User140420082644000
	}

	// all schemas are nil
	return nil
}
