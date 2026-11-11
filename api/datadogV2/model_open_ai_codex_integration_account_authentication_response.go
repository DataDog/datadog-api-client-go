// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OpenAICodexIntegrationAccountAuthenticationResponse - Authentication configured on the OpenAI Codex integration account.
type OpenAICodexIntegrationAccountAuthenticationResponse struct {
	OpenAICodexIntegrationAccountBearerTokenAuthResponse *OpenAICodexIntegrationAccountBearerTokenAuthResponse

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// OpenAICodexIntegrationAccountBearerTokenAuthResponseAsOpenAICodexIntegrationAccountAuthenticationResponse is a convenience function that returns OpenAICodexIntegrationAccountBearerTokenAuthResponse wrapped in OpenAICodexIntegrationAccountAuthenticationResponse.
func OpenAICodexIntegrationAccountBearerTokenAuthResponseAsOpenAICodexIntegrationAccountAuthenticationResponse(v *OpenAICodexIntegrationAccountBearerTokenAuthResponse) OpenAICodexIntegrationAccountAuthenticationResponse {
	return OpenAICodexIntegrationAccountAuthenticationResponse{OpenAICodexIntegrationAccountBearerTokenAuthResponse: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *OpenAICodexIntegrationAccountAuthenticationResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into OpenAICodexIntegrationAccountBearerTokenAuthResponse
	err = datadog.Unmarshal(data, &obj.OpenAICodexIntegrationAccountBearerTokenAuthResponse)
	if err == nil {
		if obj.OpenAICodexIntegrationAccountBearerTokenAuthResponse != nil && obj.OpenAICodexIntegrationAccountBearerTokenAuthResponse.UnparsedObject == nil {
			jsonOpenAICodexIntegrationAccountBearerTokenAuthResponse, _ := datadog.Marshal(obj.OpenAICodexIntegrationAccountBearerTokenAuthResponse)
			if string(jsonOpenAICodexIntegrationAccountBearerTokenAuthResponse) == "{}" { // empty struct
				obj.OpenAICodexIntegrationAccountBearerTokenAuthResponse = nil
			} else {
				match++
			}
		} else {
			obj.OpenAICodexIntegrationAccountBearerTokenAuthResponse = nil
		}
	} else {
		obj.OpenAICodexIntegrationAccountBearerTokenAuthResponse = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.OpenAICodexIntegrationAccountBearerTokenAuthResponse = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj OpenAICodexIntegrationAccountAuthenticationResponse) MarshalJSON() ([]byte, error) {
	if obj.OpenAICodexIntegrationAccountBearerTokenAuthResponse != nil {
		return datadog.Marshal(&obj.OpenAICodexIntegrationAccountBearerTokenAuthResponse)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *OpenAICodexIntegrationAccountAuthenticationResponse) GetActualInstance() interface{} {
	if obj.OpenAICodexIntegrationAccountBearerTokenAuthResponse != nil {
		return obj.OpenAICodexIntegrationAccountBearerTokenAuthResponse
	}

	// all schemas are nil
	return nil
}
