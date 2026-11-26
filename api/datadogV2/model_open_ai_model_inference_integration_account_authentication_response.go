// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OpenAIModelInferenceIntegrationAccountAuthenticationResponse - Authentication configured on the OpenAI Model Inference integration account.
type OpenAIModelInferenceIntegrationAccountAuthenticationResponse struct {
	OpenAIModelInferenceIntegrationAccountBearerTokenAuthResponse *OpenAIModelInferenceIntegrationAccountBearerTokenAuthResponse

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// OpenAIModelInferenceIntegrationAccountBearerTokenAuthResponseAsOpenAIModelInferenceIntegrationAccountAuthenticationResponse is a convenience function that returns OpenAIModelInferenceIntegrationAccountBearerTokenAuthResponse wrapped in OpenAIModelInferenceIntegrationAccountAuthenticationResponse.
func OpenAIModelInferenceIntegrationAccountBearerTokenAuthResponseAsOpenAIModelInferenceIntegrationAccountAuthenticationResponse(v *OpenAIModelInferenceIntegrationAccountBearerTokenAuthResponse) OpenAIModelInferenceIntegrationAccountAuthenticationResponse {
	return OpenAIModelInferenceIntegrationAccountAuthenticationResponse{OpenAIModelInferenceIntegrationAccountBearerTokenAuthResponse: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *OpenAIModelInferenceIntegrationAccountAuthenticationResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into OpenAIModelInferenceIntegrationAccountBearerTokenAuthResponse
	err = datadog.Unmarshal(data, &obj.OpenAIModelInferenceIntegrationAccountBearerTokenAuthResponse)
	if err == nil {
		if obj.OpenAIModelInferenceIntegrationAccountBearerTokenAuthResponse != nil && obj.OpenAIModelInferenceIntegrationAccountBearerTokenAuthResponse.UnparsedObject == nil {
			jsonOpenAIModelInferenceIntegrationAccountBearerTokenAuthResponse, _ := datadog.Marshal(obj.OpenAIModelInferenceIntegrationAccountBearerTokenAuthResponse)
			if string(jsonOpenAIModelInferenceIntegrationAccountBearerTokenAuthResponse) == "{}" { // empty struct
				obj.OpenAIModelInferenceIntegrationAccountBearerTokenAuthResponse = nil
			} else {
				match++
			}
		} else {
			obj.OpenAIModelInferenceIntegrationAccountBearerTokenAuthResponse = nil
		}
	} else {
		obj.OpenAIModelInferenceIntegrationAccountBearerTokenAuthResponse = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.OpenAIModelInferenceIntegrationAccountBearerTokenAuthResponse = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj OpenAIModelInferenceIntegrationAccountAuthenticationResponse) MarshalJSON() ([]byte, error) {
	if obj.OpenAIModelInferenceIntegrationAccountBearerTokenAuthResponse != nil {
		return datadog.Marshal(&obj.OpenAIModelInferenceIntegrationAccountBearerTokenAuthResponse)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *OpenAIModelInferenceIntegrationAccountAuthenticationResponse) GetActualInstance() interface{} {
	if obj.OpenAIModelInferenceIntegrationAccountBearerTokenAuthResponse != nil {
		return obj.OpenAIModelInferenceIntegrationAccountBearerTokenAuthResponse
	}

	// all schemas are nil
	return nil
}
