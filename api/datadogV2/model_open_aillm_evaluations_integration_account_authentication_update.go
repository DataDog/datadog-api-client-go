// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OpenAILLMEvaluationsIntegrationAccountAuthenticationUpdate - Authentication for updating the OpenAI LLM Evaluations integration account. Exactly one method is set.
type OpenAILLMEvaluationsIntegrationAccountAuthenticationUpdate struct {
	OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthUpdate *OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthUpdate

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthUpdateAsOpenAILLMEvaluationsIntegrationAccountAuthenticationUpdate is a convenience function that returns OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthUpdate wrapped in OpenAILLMEvaluationsIntegrationAccountAuthenticationUpdate.
func OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthUpdateAsOpenAILLMEvaluationsIntegrationAccountAuthenticationUpdate(v *OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthUpdate) OpenAILLMEvaluationsIntegrationAccountAuthenticationUpdate {
	return OpenAILLMEvaluationsIntegrationAccountAuthenticationUpdate{OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthUpdate: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *OpenAILLMEvaluationsIntegrationAccountAuthenticationUpdate) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthUpdate
	err = datadog.Unmarshal(data, &obj.OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthUpdate)
	if err == nil {
		if obj.OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthUpdate != nil && obj.OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthUpdate.UnparsedObject == nil {
			jsonOpenAILLMEvaluationsIntegrationAccountBearerTokenAuthUpdate, _ := datadog.Marshal(obj.OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthUpdate)
			if string(jsonOpenAILLMEvaluationsIntegrationAccountBearerTokenAuthUpdate) == "{}" { // empty struct
				obj.OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthUpdate = nil
			} else {
				match++
			}
		} else {
			obj.OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthUpdate = nil
		}
	} else {
		obj.OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthUpdate = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthUpdate = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj OpenAILLMEvaluationsIntegrationAccountAuthenticationUpdate) MarshalJSON() ([]byte, error) {
	if obj.OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthUpdate != nil {
		return datadog.Marshal(&obj.OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthUpdate)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *OpenAILLMEvaluationsIntegrationAccountAuthenticationUpdate) GetActualInstance() interface{} {
	if obj.OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthUpdate != nil {
		return obj.OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthUpdate
	}

	// all schemas are nil
	return nil
}
