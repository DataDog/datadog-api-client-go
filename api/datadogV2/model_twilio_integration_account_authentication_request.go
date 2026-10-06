// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TwilioIntegrationAccountAuthenticationRequest - Authentication for creating the Twilio integration account. Exactly one method is set.
type TwilioIntegrationAccountAuthenticationRequest struct {
	TwilioIntegrationAccountBasicAuthRequest *TwilioIntegrationAccountBasicAuthRequest

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// TwilioIntegrationAccountBasicAuthRequestAsTwilioIntegrationAccountAuthenticationRequest is a convenience function that returns TwilioIntegrationAccountBasicAuthRequest wrapped in TwilioIntegrationAccountAuthenticationRequest.
func TwilioIntegrationAccountBasicAuthRequestAsTwilioIntegrationAccountAuthenticationRequest(v *TwilioIntegrationAccountBasicAuthRequest) TwilioIntegrationAccountAuthenticationRequest {
	return TwilioIntegrationAccountAuthenticationRequest{TwilioIntegrationAccountBasicAuthRequest: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *TwilioIntegrationAccountAuthenticationRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into TwilioIntegrationAccountBasicAuthRequest
	err = datadog.Unmarshal(data, &obj.TwilioIntegrationAccountBasicAuthRequest)
	if err == nil {
		if obj.TwilioIntegrationAccountBasicAuthRequest != nil && obj.TwilioIntegrationAccountBasicAuthRequest.UnparsedObject == nil {
			jsonTwilioIntegrationAccountBasicAuthRequest, _ := datadog.Marshal(obj.TwilioIntegrationAccountBasicAuthRequest)
			if string(jsonTwilioIntegrationAccountBasicAuthRequest) == "{}" { // empty struct
				obj.TwilioIntegrationAccountBasicAuthRequest = nil
			} else {
				match++
			}
		} else {
			obj.TwilioIntegrationAccountBasicAuthRequest = nil
		}
	} else {
		obj.TwilioIntegrationAccountBasicAuthRequest = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.TwilioIntegrationAccountBasicAuthRequest = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj TwilioIntegrationAccountAuthenticationRequest) MarshalJSON() ([]byte, error) {
	if obj.TwilioIntegrationAccountBasicAuthRequest != nil {
		return datadog.Marshal(&obj.TwilioIntegrationAccountBasicAuthRequest)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *TwilioIntegrationAccountAuthenticationRequest) GetActualInstance() interface{} {
	if obj.TwilioIntegrationAccountBasicAuthRequest != nil {
		return obj.TwilioIntegrationAccountBasicAuthRequest
	}

	// all schemas are nil
	return nil
}
