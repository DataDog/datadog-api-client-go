// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/goccy/go-json"
)

// HttpDestinationAuth - The authentication method used for HTTP destinations.
type HttpDestinationAuth struct {
	HttpDestinationBasicAuth        *HttpDestinationBasicAuth
	HttpDestinationCustomHeaderAuth *HttpDestinationCustomHeaderAuth

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// HttpDestinationBasicAuthAsHttpDestinationAuth is a convenience function that returns HttpDestinationBasicAuth wrapped in HttpDestinationAuth.
func HttpDestinationBasicAuthAsHttpDestinationAuth(v *HttpDestinationBasicAuth) HttpDestinationAuth {
	return HttpDestinationAuth{HttpDestinationBasicAuth: v}
}

// HttpDestinationCustomHeaderAuthAsHttpDestinationAuth is a convenience function that returns HttpDestinationCustomHeaderAuth wrapped in HttpDestinationAuth.
func HttpDestinationCustomHeaderAuthAsHttpDestinationAuth(v *HttpDestinationCustomHeaderAuth) HttpDestinationAuth {
	return HttpDestinationAuth{HttpDestinationCustomHeaderAuth: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *HttpDestinationAuth) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into HttpDestinationBasicAuth
	err = json.Unmarshal(data, &obj.HttpDestinationBasicAuth)
	if err == nil {
		if obj.HttpDestinationBasicAuth != nil && obj.HttpDestinationBasicAuth.UnparsedObject == nil {
			jsonHttpDestinationBasicAuth, _ := json.Marshal(obj.HttpDestinationBasicAuth)
			if string(jsonHttpDestinationBasicAuth) == "{}" { // empty struct
				obj.HttpDestinationBasicAuth = nil
			} else {
				match++
			}
		} else {
			obj.HttpDestinationBasicAuth = nil
		}
	} else {
		obj.HttpDestinationBasicAuth = nil
	}

	// try to unmarshal data into HttpDestinationCustomHeaderAuth
	err = json.Unmarshal(data, &obj.HttpDestinationCustomHeaderAuth)
	if err == nil {
		if obj.HttpDestinationCustomHeaderAuth != nil && obj.HttpDestinationCustomHeaderAuth.UnparsedObject == nil {
			jsonHttpDestinationCustomHeaderAuth, _ := json.Marshal(obj.HttpDestinationCustomHeaderAuth)
			if string(jsonHttpDestinationCustomHeaderAuth) == "{}" { // empty struct
				obj.HttpDestinationCustomHeaderAuth = nil
			} else {
				match++
			}
		} else {
			obj.HttpDestinationCustomHeaderAuth = nil
		}
	} else {
		obj.HttpDestinationCustomHeaderAuth = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.HttpDestinationBasicAuth = nil
		obj.HttpDestinationCustomHeaderAuth = nil
		return json.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj HttpDestinationAuth) MarshalJSON() ([]byte, error) {
	if obj.HttpDestinationBasicAuth != nil {
		return json.Marshal(&obj.HttpDestinationBasicAuth)
	}

	if obj.HttpDestinationCustomHeaderAuth != nil {
		return json.Marshal(&obj.HttpDestinationCustomHeaderAuth)
	}

	if obj.UnparsedObject != nil {
		return json.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *HttpDestinationAuth) GetActualInstance() interface{} {
	if obj.HttpDestinationBasicAuth != nil {
		return obj.HttpDestinationBasicAuth
	}

	if obj.HttpDestinationCustomHeaderAuth != nil {
		return obj.HttpDestinationCustomHeaderAuth
	}

	// all schemas are nil
	return nil
}
