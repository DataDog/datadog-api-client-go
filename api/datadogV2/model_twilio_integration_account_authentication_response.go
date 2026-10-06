// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TwilioIntegrationAccountAuthenticationResponse - Authentication configured on the Twilio integration account.
type TwilioIntegrationAccountAuthenticationResponse struct {
	TwilioIntegrationAccountBasicAuthResponse *TwilioIntegrationAccountBasicAuthResponse

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// TwilioIntegrationAccountBasicAuthResponseAsTwilioIntegrationAccountAuthenticationResponse is a convenience function that returns TwilioIntegrationAccountBasicAuthResponse wrapped in TwilioIntegrationAccountAuthenticationResponse.
func TwilioIntegrationAccountBasicAuthResponseAsTwilioIntegrationAccountAuthenticationResponse(v *TwilioIntegrationAccountBasicAuthResponse) TwilioIntegrationAccountAuthenticationResponse {
	return TwilioIntegrationAccountAuthenticationResponse{TwilioIntegrationAccountBasicAuthResponse: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *TwilioIntegrationAccountAuthenticationResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into TwilioIntegrationAccountBasicAuthResponse
	err = datadog.Unmarshal(data, &obj.TwilioIntegrationAccountBasicAuthResponse)
	if err == nil {
		if obj.TwilioIntegrationAccountBasicAuthResponse != nil && obj.TwilioIntegrationAccountBasicAuthResponse.UnparsedObject == nil {
			jsonTwilioIntegrationAccountBasicAuthResponse, _ := datadog.Marshal(obj.TwilioIntegrationAccountBasicAuthResponse)
			if string(jsonTwilioIntegrationAccountBasicAuthResponse) == "{}" { // empty struct
				obj.TwilioIntegrationAccountBasicAuthResponse = nil
			} else {
				match++
			}
		} else {
			obj.TwilioIntegrationAccountBasicAuthResponse = nil
		}
	} else {
		obj.TwilioIntegrationAccountBasicAuthResponse = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.TwilioIntegrationAccountBasicAuthResponse = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj TwilioIntegrationAccountAuthenticationResponse) MarshalJSON() ([]byte, error) {
	if obj.TwilioIntegrationAccountBasicAuthResponse != nil {
		return datadog.Marshal(&obj.TwilioIntegrationAccountBasicAuthResponse)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *TwilioIntegrationAccountAuthenticationResponse) GetActualInstance() interface{} {
	if obj.TwilioIntegrationAccountBasicAuthResponse != nil {
		return obj.TwilioIntegrationAccountBasicAuthResponse
	}

	// all schemas are nil
	return nil
}
