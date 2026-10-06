// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TwilioIntegrationAccountAuthenticationUpdate - Authentication for updating the Twilio integration account. Exactly one method is set.
type TwilioIntegrationAccountAuthenticationUpdate struct {
	TwilioIntegrationAccountBasicAuthUpdate *TwilioIntegrationAccountBasicAuthUpdate

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// TwilioIntegrationAccountBasicAuthUpdateAsTwilioIntegrationAccountAuthenticationUpdate is a convenience function that returns TwilioIntegrationAccountBasicAuthUpdate wrapped in TwilioIntegrationAccountAuthenticationUpdate.
func TwilioIntegrationAccountBasicAuthUpdateAsTwilioIntegrationAccountAuthenticationUpdate(v *TwilioIntegrationAccountBasicAuthUpdate) TwilioIntegrationAccountAuthenticationUpdate {
	return TwilioIntegrationAccountAuthenticationUpdate{TwilioIntegrationAccountBasicAuthUpdate: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *TwilioIntegrationAccountAuthenticationUpdate) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into TwilioIntegrationAccountBasicAuthUpdate
	err = datadog.Unmarshal(data, &obj.TwilioIntegrationAccountBasicAuthUpdate)
	if err == nil {
		if obj.TwilioIntegrationAccountBasicAuthUpdate != nil && obj.TwilioIntegrationAccountBasicAuthUpdate.UnparsedObject == nil {
			jsonTwilioIntegrationAccountBasicAuthUpdate, _ := datadog.Marshal(obj.TwilioIntegrationAccountBasicAuthUpdate)
			if string(jsonTwilioIntegrationAccountBasicAuthUpdate) == "{}" { // empty struct
				obj.TwilioIntegrationAccountBasicAuthUpdate = nil
			} else {
				match++
			}
		} else {
			obj.TwilioIntegrationAccountBasicAuthUpdate = nil
		}
	} else {
		obj.TwilioIntegrationAccountBasicAuthUpdate = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.TwilioIntegrationAccountBasicAuthUpdate = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj TwilioIntegrationAccountAuthenticationUpdate) MarshalJSON() ([]byte, error) {
	if obj.TwilioIntegrationAccountBasicAuthUpdate != nil {
		return datadog.Marshal(&obj.TwilioIntegrationAccountBasicAuthUpdate)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *TwilioIntegrationAccountAuthenticationUpdate) GetActualInstance() interface{} {
	if obj.TwilioIntegrationAccountBasicAuthUpdate != nil {
		return obj.TwilioIntegrationAccountBasicAuthUpdate
	}

	// all schemas are nil
	return nil
}
