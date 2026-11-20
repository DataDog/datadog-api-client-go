// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OpenAILLMEvaluationsIntegrationAccountAuthenticationRequest - Authentication for creating the OpenAI LLM Evaluations integration account. Exactly one method is set.
type OpenAILLMEvaluationsIntegrationAccountAuthenticationRequest struct {
	OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthRequest *OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthRequest

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthRequestAsOpenAILLMEvaluationsIntegrationAccountAuthenticationRequest is a convenience function that returns OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthRequest wrapped in OpenAILLMEvaluationsIntegrationAccountAuthenticationRequest.
func OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthRequestAsOpenAILLMEvaluationsIntegrationAccountAuthenticationRequest(v *OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthRequest) OpenAILLMEvaluationsIntegrationAccountAuthenticationRequest {
	return OpenAILLMEvaluationsIntegrationAccountAuthenticationRequest{OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthRequest: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *OpenAILLMEvaluationsIntegrationAccountAuthenticationRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthRequest
	err = datadog.Unmarshal(data, &obj.OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthRequest)
	if err == nil {
		if obj.OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthRequest != nil && obj.OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthRequest.UnparsedObject == nil {
			jsonOpenAILLMEvaluationsIntegrationAccountBearerTokenAuthRequest, _ := datadog.Marshal(obj.OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthRequest)
			if string(jsonOpenAILLMEvaluationsIntegrationAccountBearerTokenAuthRequest) == "{}" { // empty struct
				obj.OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthRequest = nil
			} else {
				match++
			}
		} else {
			obj.OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthRequest = nil
		}
	} else {
		obj.OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthRequest = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthRequest = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj OpenAILLMEvaluationsIntegrationAccountAuthenticationRequest) MarshalJSON() ([]byte, error) {
	if obj.OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthRequest != nil {
		return datadog.Marshal(&obj.OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthRequest)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *OpenAILLMEvaluationsIntegrationAccountAuthenticationRequest) GetActualInstance() interface{} {
	if obj.OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthRequest != nil {
		return obj.OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthRequest
	}

	// all schemas are nil
	return nil
}
