// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OpenAILLMEvaluationsIntegrationAccountAuthenticationResponse - Authentication configured on the OpenAI LLM Evaluations integration account.
type OpenAILLMEvaluationsIntegrationAccountAuthenticationResponse struct {
	OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthResponse *OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthResponse

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthResponseAsOpenAILLMEvaluationsIntegrationAccountAuthenticationResponse is a convenience function that returns OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthResponse wrapped in OpenAILLMEvaluationsIntegrationAccountAuthenticationResponse.
func OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthResponseAsOpenAILLMEvaluationsIntegrationAccountAuthenticationResponse(v *OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthResponse) OpenAILLMEvaluationsIntegrationAccountAuthenticationResponse {
	return OpenAILLMEvaluationsIntegrationAccountAuthenticationResponse{OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthResponse: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *OpenAILLMEvaluationsIntegrationAccountAuthenticationResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthResponse
	err = datadog.Unmarshal(data, &obj.OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthResponse)
	if err == nil {
		if obj.OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthResponse != nil && obj.OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthResponse.UnparsedObject == nil {
			jsonOpenAILLMEvaluationsIntegrationAccountBearerTokenAuthResponse, _ := datadog.Marshal(obj.OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthResponse)
			if string(jsonOpenAILLMEvaluationsIntegrationAccountBearerTokenAuthResponse) == "{}" { // empty struct
				obj.OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthResponse = nil
			} else {
				match++
			}
		} else {
			obj.OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthResponse = nil
		}
	} else {
		obj.OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthResponse = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthResponse = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj OpenAILLMEvaluationsIntegrationAccountAuthenticationResponse) MarshalJSON() ([]byte, error) {
	if obj.OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthResponse != nil {
		return datadog.Marshal(&obj.OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthResponse)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *OpenAILLMEvaluationsIntegrationAccountAuthenticationResponse) GetActualInstance() interface{} {
	if obj.OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthResponse != nil {
		return obj.OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthResponse
	}

	// all schemas are nil
	return nil
}
