// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SnowflakeIntegrationAccountAuthenticationRequest - Authentication for creating the Snowflake integration account. Exactly one method is set.
type SnowflakeIntegrationAccountAuthenticationRequest struct {
	SnowflakeIntegrationAccountPrivateKeyAuthRequest *SnowflakeIntegrationAccountPrivateKeyAuthRequest

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// SnowflakeIntegrationAccountPrivateKeyAuthRequestAsSnowflakeIntegrationAccountAuthenticationRequest is a convenience function that returns SnowflakeIntegrationAccountPrivateKeyAuthRequest wrapped in SnowflakeIntegrationAccountAuthenticationRequest.
func SnowflakeIntegrationAccountPrivateKeyAuthRequestAsSnowflakeIntegrationAccountAuthenticationRequest(v *SnowflakeIntegrationAccountPrivateKeyAuthRequest) SnowflakeIntegrationAccountAuthenticationRequest {
	return SnowflakeIntegrationAccountAuthenticationRequest{SnowflakeIntegrationAccountPrivateKeyAuthRequest: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *SnowflakeIntegrationAccountAuthenticationRequest) UnmarshalJSON(data []byte) error {
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
func (obj SnowflakeIntegrationAccountAuthenticationRequest) MarshalJSON() ([]byte, error) {
	if obj.SnowflakeIntegrationAccountPrivateKeyAuthRequest != nil {
		return datadog.Marshal(&obj.SnowflakeIntegrationAccountPrivateKeyAuthRequest)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *SnowflakeIntegrationAccountAuthenticationRequest) GetActualInstance() interface{} {
	if obj.SnowflakeIntegrationAccountPrivateKeyAuthRequest != nil {
		return obj.SnowflakeIntegrationAccountPrivateKeyAuthRequest
	}

	// all schemas are nil
	return nil
}
