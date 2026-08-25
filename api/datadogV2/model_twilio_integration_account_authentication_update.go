// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TwilioIntegrationAccountAuthenticationUpdate - Authentication for updating the Twilio integration account. Exactly one method is set.
type TwilioIntegrationAccountAuthenticationUpdate struct {
	IntegrationAccountBasicAuthUpdate *IntegrationAccountBasicAuthUpdate

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// IntegrationAccountBasicAuthUpdateAsTwilioIntegrationAccountAuthenticationUpdate is a convenience function that returns IntegrationAccountBasicAuthUpdate wrapped in TwilioIntegrationAccountAuthenticationUpdate.
func IntegrationAccountBasicAuthUpdateAsTwilioIntegrationAccountAuthenticationUpdate(v *IntegrationAccountBasicAuthUpdate) TwilioIntegrationAccountAuthenticationUpdate {
	return TwilioIntegrationAccountAuthenticationUpdate{IntegrationAccountBasicAuthUpdate: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *TwilioIntegrationAccountAuthenticationUpdate) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into IntegrationAccountBasicAuthUpdate
	err = datadog.Unmarshal(data, &obj.IntegrationAccountBasicAuthUpdate)
	if err == nil {
		if obj.IntegrationAccountBasicAuthUpdate != nil && obj.IntegrationAccountBasicAuthUpdate.UnparsedObject == nil {
			jsonIntegrationAccountBasicAuthUpdate, _ := datadog.Marshal(obj.IntegrationAccountBasicAuthUpdate)
			if string(jsonIntegrationAccountBasicAuthUpdate) == "{}" { // empty struct
				obj.IntegrationAccountBasicAuthUpdate = nil
			} else {
				match++
			}
		} else {
			obj.IntegrationAccountBasicAuthUpdate = nil
		}
	} else {
		obj.IntegrationAccountBasicAuthUpdate = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.IntegrationAccountBasicAuthUpdate = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj TwilioIntegrationAccountAuthenticationUpdate) MarshalJSON() ([]byte, error) {
	if obj.IntegrationAccountBasicAuthUpdate != nil {
		return datadog.Marshal(&obj.IntegrationAccountBasicAuthUpdate)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *TwilioIntegrationAccountAuthenticationUpdate) GetActualInstance() interface{} {
	if obj.IntegrationAccountBasicAuthUpdate != nil {
		return obj.IntegrationAccountBasicAuthUpdate
	}

	// all schemas are nil
	return nil
}
