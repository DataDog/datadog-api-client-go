// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SnowflakeIntegrationAccountAuthenticationUpdate - Authentication for updating the Snowflake integration account. Exactly one method is set. An update replaces the authentication object entirely, so it requires the same fields as creating an account.
type SnowflakeIntegrationAccountAuthenticationUpdate struct {
	SnowflakeIntegrationAccountPrivateKeyAuthRequest *SnowflakeIntegrationAccountPrivateKeyAuthRequest

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// SnowflakeIntegrationAccountPrivateKeyAuthRequestAsSnowflakeIntegrationAccountAuthenticationUpdate is a convenience function that returns SnowflakeIntegrationAccountPrivateKeyAuthRequest wrapped in SnowflakeIntegrationAccountAuthenticationUpdate.
func SnowflakeIntegrationAccountPrivateKeyAuthRequestAsSnowflakeIntegrationAccountAuthenticationUpdate(v *SnowflakeIntegrationAccountPrivateKeyAuthRequest) SnowflakeIntegrationAccountAuthenticationUpdate {
	return SnowflakeIntegrationAccountAuthenticationUpdate{SnowflakeIntegrationAccountPrivateKeyAuthRequest: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *SnowflakeIntegrationAccountAuthenticationUpdate) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SnowflakeIntegrationAccountPrivateKeyAuthRequest
	err = datadog.Unmarshal(data, &obj.SnowflakeIntegrationAccountPrivateKeyAuthRequest)
	if err == nil {
		if obj.SnowflakeIntegrationAccountPrivateKeyAuthRequest != nil && obj.SnowflakeIntegrationAccountPrivateKeyAuthRequest.UnparsedObject == nil {
			jsonSnowflakeIntegrationAccountPrivateKeyAuthRequest, _ := datadog.Marshal(obj.SnowflakeIntegrationAccountPrivateKeyAuthRequest)
			if string(jsonSnowflakeIntegrationAccountPrivateKeyAuthRequest) == "{}" { // empty struct
				obj.SnowflakeIntegrationAccountPrivateKeyAuthRequest = nil
			} else {
				match++
			}
		} else {
			obj.SnowflakeIntegrationAccountPrivateKeyAuthRequest = nil
		}
	} else {
		obj.SnowflakeIntegrationAccountPrivateKeyAuthRequest = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.SnowflakeIntegrationAccountPrivateKeyAuthRequest = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj SnowflakeIntegrationAccountAuthenticationUpdate) MarshalJSON() ([]byte, error) {
	if obj.SnowflakeIntegrationAccountPrivateKeyAuthRequest != nil {
		return datadog.Marshal(&obj.SnowflakeIntegrationAccountPrivateKeyAuthRequest)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *SnowflakeIntegrationAccountAuthenticationUpdate) GetActualInstance() interface{} {
	if obj.SnowflakeIntegrationAccountPrivateKeyAuthRequest != nil {
		return obj.SnowflakeIntegrationAccountPrivateKeyAuthRequest
	}

	// all schemas are nil
	return nil
}
