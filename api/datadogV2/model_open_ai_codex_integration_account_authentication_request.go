// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OpenAICodexIntegrationAccountAuthenticationRequest - Authentication for creating the OpenAI Codex integration account. Exactly one method is set.
type OpenAICodexIntegrationAccountAuthenticationRequest struct {
	OpenAICodexIntegrationAccountBearerTokenAuthRequest *OpenAICodexIntegrationAccountBearerTokenAuthRequest

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// OpenAICodexIntegrationAccountBearerTokenAuthRequestAsOpenAICodexIntegrationAccountAuthenticationRequest is a convenience function that returns OpenAICodexIntegrationAccountBearerTokenAuthRequest wrapped in OpenAICodexIntegrationAccountAuthenticationRequest.
func OpenAICodexIntegrationAccountBearerTokenAuthRequestAsOpenAICodexIntegrationAccountAuthenticationRequest(v *OpenAICodexIntegrationAccountBearerTokenAuthRequest) OpenAICodexIntegrationAccountAuthenticationRequest {
	return OpenAICodexIntegrationAccountAuthenticationRequest{OpenAICodexIntegrationAccountBearerTokenAuthRequest: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *OpenAICodexIntegrationAccountAuthenticationRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into OpenAICodexIntegrationAccountBearerTokenAuthRequest
	err = datadog.Unmarshal(data, &obj.OpenAICodexIntegrationAccountBearerTokenAuthRequest)
	if err == nil {
		if obj.OpenAICodexIntegrationAccountBearerTokenAuthRequest != nil && obj.OpenAICodexIntegrationAccountBearerTokenAuthRequest.UnparsedObject == nil {
			jsonOpenAICodexIntegrationAccountBearerTokenAuthRequest, _ := datadog.Marshal(obj.OpenAICodexIntegrationAccountBearerTokenAuthRequest)
			if string(jsonOpenAICodexIntegrationAccountBearerTokenAuthRequest) == "{}" { // empty struct
				obj.OpenAICodexIntegrationAccountBearerTokenAuthRequest = nil
			} else {
				match++
			}
		} else {
			obj.OpenAICodexIntegrationAccountBearerTokenAuthRequest = nil
		}
	} else {
		obj.OpenAICodexIntegrationAccountBearerTokenAuthRequest = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.OpenAICodexIntegrationAccountBearerTokenAuthRequest = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj OpenAICodexIntegrationAccountAuthenticationRequest) MarshalJSON() ([]byte, error) {
	if obj.OpenAICodexIntegrationAccountBearerTokenAuthRequest != nil {
		return datadog.Marshal(&obj.OpenAICodexIntegrationAccountBearerTokenAuthRequest)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *OpenAICodexIntegrationAccountAuthenticationRequest) GetActualInstance() interface{} {
	if obj.OpenAICodexIntegrationAccountBearerTokenAuthRequest != nil {
		return obj.OpenAICodexIntegrationAccountBearerTokenAuthRequest
	}

	// all schemas are nil
	return nil
}
