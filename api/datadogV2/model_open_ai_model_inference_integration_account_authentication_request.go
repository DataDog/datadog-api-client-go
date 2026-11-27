// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OpenAIModelInferenceIntegrationAccountAuthenticationRequest - Authentication for creating the OpenAI Model Inference integration account. Set exactly one method.
type OpenAIModelInferenceIntegrationAccountAuthenticationRequest struct {
	OpenAIModelInferenceIntegrationAccountBearerTokenAuthRequest *OpenAIModelInferenceIntegrationAccountBearerTokenAuthRequest

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// OpenAIModelInferenceIntegrationAccountBearerTokenAuthRequestAsOpenAIModelInferenceIntegrationAccountAuthenticationRequest is a convenience function that returns OpenAIModelInferenceIntegrationAccountBearerTokenAuthRequest wrapped in OpenAIModelInferenceIntegrationAccountAuthenticationRequest.
func OpenAIModelInferenceIntegrationAccountBearerTokenAuthRequestAsOpenAIModelInferenceIntegrationAccountAuthenticationRequest(v *OpenAIModelInferenceIntegrationAccountBearerTokenAuthRequest) OpenAIModelInferenceIntegrationAccountAuthenticationRequest {
	return OpenAIModelInferenceIntegrationAccountAuthenticationRequest{OpenAIModelInferenceIntegrationAccountBearerTokenAuthRequest: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *OpenAIModelInferenceIntegrationAccountAuthenticationRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into OpenAIModelInferenceIntegrationAccountBearerTokenAuthRequest
	err = datadog.Unmarshal(data, &obj.OpenAIModelInferenceIntegrationAccountBearerTokenAuthRequest)
	if err == nil {
		if obj.OpenAIModelInferenceIntegrationAccountBearerTokenAuthRequest != nil && obj.OpenAIModelInferenceIntegrationAccountBearerTokenAuthRequest.UnparsedObject == nil {
			jsonOpenAIModelInferenceIntegrationAccountBearerTokenAuthRequest, _ := datadog.Marshal(obj.OpenAIModelInferenceIntegrationAccountBearerTokenAuthRequest)
			if string(jsonOpenAIModelInferenceIntegrationAccountBearerTokenAuthRequest) == "{}" { // empty struct
				obj.OpenAIModelInferenceIntegrationAccountBearerTokenAuthRequest = nil
			} else {
				match++
			}
		} else {
			obj.OpenAIModelInferenceIntegrationAccountBearerTokenAuthRequest = nil
		}
	} else {
		obj.OpenAIModelInferenceIntegrationAccountBearerTokenAuthRequest = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.OpenAIModelInferenceIntegrationAccountBearerTokenAuthRequest = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj OpenAIModelInferenceIntegrationAccountAuthenticationRequest) MarshalJSON() ([]byte, error) {
	if obj.OpenAIModelInferenceIntegrationAccountBearerTokenAuthRequest != nil {
		return datadog.Marshal(&obj.OpenAIModelInferenceIntegrationAccountBearerTokenAuthRequest)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *OpenAIModelInferenceIntegrationAccountAuthenticationRequest) GetActualInstance() interface{} {
	if obj.OpenAIModelInferenceIntegrationAccountBearerTokenAuthRequest != nil {
		return obj.OpenAIModelInferenceIntegrationAccountBearerTokenAuthRequest
	}

	// all schemas are nil
	return nil
}
