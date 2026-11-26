// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OpenAIModelInferenceIntegrationAccountAuthenticationUpdate - Authentication for updating the OpenAI Model Inference integration account. Set exactly one method.
type OpenAIModelInferenceIntegrationAccountAuthenticationUpdate struct {
	OpenAIModelInferenceIntegrationAccountBearerTokenAuthUpdate *OpenAIModelInferenceIntegrationAccountBearerTokenAuthUpdate

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// OpenAIModelInferenceIntegrationAccountBearerTokenAuthUpdateAsOpenAIModelInferenceIntegrationAccountAuthenticationUpdate is a convenience function that returns OpenAIModelInferenceIntegrationAccountBearerTokenAuthUpdate wrapped in OpenAIModelInferenceIntegrationAccountAuthenticationUpdate.
func OpenAIModelInferenceIntegrationAccountBearerTokenAuthUpdateAsOpenAIModelInferenceIntegrationAccountAuthenticationUpdate(v *OpenAIModelInferenceIntegrationAccountBearerTokenAuthUpdate) OpenAIModelInferenceIntegrationAccountAuthenticationUpdate {
	return OpenAIModelInferenceIntegrationAccountAuthenticationUpdate{OpenAIModelInferenceIntegrationAccountBearerTokenAuthUpdate: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *OpenAIModelInferenceIntegrationAccountAuthenticationUpdate) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into OpenAIModelInferenceIntegrationAccountBearerTokenAuthUpdate
	err = datadog.Unmarshal(data, &obj.OpenAIModelInferenceIntegrationAccountBearerTokenAuthUpdate)
	if err == nil {
		if obj.OpenAIModelInferenceIntegrationAccountBearerTokenAuthUpdate != nil && obj.OpenAIModelInferenceIntegrationAccountBearerTokenAuthUpdate.UnparsedObject == nil {
			jsonOpenAIModelInferenceIntegrationAccountBearerTokenAuthUpdate, _ := datadog.Marshal(obj.OpenAIModelInferenceIntegrationAccountBearerTokenAuthUpdate)
			if string(jsonOpenAIModelInferenceIntegrationAccountBearerTokenAuthUpdate) == "{}" { // empty struct
				obj.OpenAIModelInferenceIntegrationAccountBearerTokenAuthUpdate = nil
			} else {
				match++
			}
		} else {
			obj.OpenAIModelInferenceIntegrationAccountBearerTokenAuthUpdate = nil
		}
	} else {
		obj.OpenAIModelInferenceIntegrationAccountBearerTokenAuthUpdate = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.OpenAIModelInferenceIntegrationAccountBearerTokenAuthUpdate = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj OpenAIModelInferenceIntegrationAccountAuthenticationUpdate) MarshalJSON() ([]byte, error) {
	if obj.OpenAIModelInferenceIntegrationAccountBearerTokenAuthUpdate != nil {
		return datadog.Marshal(&obj.OpenAIModelInferenceIntegrationAccountBearerTokenAuthUpdate)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *OpenAIModelInferenceIntegrationAccountAuthenticationUpdate) GetActualInstance() interface{} {
	if obj.OpenAIModelInferenceIntegrationAccountBearerTokenAuthUpdate != nil {
		return obj.OpenAIModelInferenceIntegrationAccountBearerTokenAuthUpdate
	}

	// all schemas are nil
	return nil
}
