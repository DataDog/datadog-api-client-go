// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TwilioAuthentication - Authentication methods supported by the Twilio interface. Exactly one is set, selected by its `type`.
type TwilioAuthentication struct {
	TwilioBasicAuth *TwilioBasicAuth

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// TwilioBasicAuthAsTwilioAuthentication is a convenience function that returns TwilioBasicAuth wrapped in TwilioAuthentication.
func TwilioBasicAuthAsTwilioAuthentication(v *TwilioBasicAuth) TwilioAuthentication {
	return TwilioAuthentication{TwilioBasicAuth: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *TwilioAuthentication) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into TwilioBasicAuth
	err = datadog.Unmarshal(data, &obj.TwilioBasicAuth)
	if err == nil {
		if obj.TwilioBasicAuth != nil && obj.TwilioBasicAuth.UnparsedObject == nil {
			jsonTwilioBasicAuth, _ := datadog.Marshal(obj.TwilioBasicAuth)
			if string(jsonTwilioBasicAuth) == "{}" { // empty struct
				obj.TwilioBasicAuth = nil
			} else {
				match++
			}
		} else {
			obj.TwilioBasicAuth = nil
		}
	} else {
		obj.TwilioBasicAuth = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.TwilioBasicAuth = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj TwilioAuthentication) MarshalJSON() ([]byte, error) {
	if obj.TwilioBasicAuth != nil {
		return datadog.Marshal(&obj.TwilioBasicAuth)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *TwilioAuthentication) GetActualInstance() interface{} {
	if obj.TwilioBasicAuth != nil {
		return obj.TwilioBasicAuth
	}

	// all schemas are nil
	return nil
}
