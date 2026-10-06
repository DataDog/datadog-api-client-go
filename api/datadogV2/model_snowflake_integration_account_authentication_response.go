// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SnowflakeIntegrationAccountAuthenticationResponse - Authentication configured on the Snowflake integration account.
type SnowflakeIntegrationAccountAuthenticationResponse struct {
	SnowflakeIntegrationAccountPrivateKeyAuthResponse *SnowflakeIntegrationAccountPrivateKeyAuthResponse

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// SnowflakeIntegrationAccountPrivateKeyAuthResponseAsSnowflakeIntegrationAccountAuthenticationResponse is a convenience function that returns SnowflakeIntegrationAccountPrivateKeyAuthResponse wrapped in SnowflakeIntegrationAccountAuthenticationResponse.
func SnowflakeIntegrationAccountPrivateKeyAuthResponseAsSnowflakeIntegrationAccountAuthenticationResponse(v *SnowflakeIntegrationAccountPrivateKeyAuthResponse) SnowflakeIntegrationAccountAuthenticationResponse {
	return SnowflakeIntegrationAccountAuthenticationResponse{SnowflakeIntegrationAccountPrivateKeyAuthResponse: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *SnowflakeIntegrationAccountAuthenticationResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SnowflakeIntegrationAccountPrivateKeyAuthResponse
	err = datadog.Unmarshal(data, &obj.SnowflakeIntegrationAccountPrivateKeyAuthResponse)
	if err == nil {
		if obj.SnowflakeIntegrationAccountPrivateKeyAuthResponse != nil && obj.SnowflakeIntegrationAccountPrivateKeyAuthResponse.UnparsedObject == nil {
			jsonSnowflakeIntegrationAccountPrivateKeyAuthResponse, _ := datadog.Marshal(obj.SnowflakeIntegrationAccountPrivateKeyAuthResponse)
			if string(jsonSnowflakeIntegrationAccountPrivateKeyAuthResponse) == "{}" { // empty struct
				obj.SnowflakeIntegrationAccountPrivateKeyAuthResponse = nil
			} else {
				match++
			}
		} else {
			obj.SnowflakeIntegrationAccountPrivateKeyAuthResponse = nil
		}
	} else {
		obj.SnowflakeIntegrationAccountPrivateKeyAuthResponse = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.SnowflakeIntegrationAccountPrivateKeyAuthResponse = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj SnowflakeIntegrationAccountAuthenticationResponse) MarshalJSON() ([]byte, error) {
	if obj.SnowflakeIntegrationAccountPrivateKeyAuthResponse != nil {
		return datadog.Marshal(&obj.SnowflakeIntegrationAccountPrivateKeyAuthResponse)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *SnowflakeIntegrationAccountAuthenticationResponse) GetActualInstance() interface{} {
	if obj.SnowflakeIntegrationAccountPrivateKeyAuthResponse != nil {
		return obj.SnowflakeIntegrationAccountPrivateKeyAuthResponse
	}

	// all schemas are nil
	return nil
}
