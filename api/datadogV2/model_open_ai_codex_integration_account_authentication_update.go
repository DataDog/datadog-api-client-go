// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OpenAICodexIntegrationAccountAuthenticationUpdate - Authentication for updating the OpenAI Codex integration account. Exactly one method is set.
type OpenAICodexIntegrationAccountAuthenticationUpdate struct {
	OpenAICodexIntegrationAccountBearerTokenAuthUpdate *OpenAICodexIntegrationAccountBearerTokenAuthUpdate

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// OpenAICodexIntegrationAccountBearerTokenAuthUpdateAsOpenAICodexIntegrationAccountAuthenticationUpdate is a convenience function that returns OpenAICodexIntegrationAccountBearerTokenAuthUpdate wrapped in OpenAICodexIntegrationAccountAuthenticationUpdate.
func OpenAICodexIntegrationAccountBearerTokenAuthUpdateAsOpenAICodexIntegrationAccountAuthenticationUpdate(v *OpenAICodexIntegrationAccountBearerTokenAuthUpdate) OpenAICodexIntegrationAccountAuthenticationUpdate {
	return OpenAICodexIntegrationAccountAuthenticationUpdate{OpenAICodexIntegrationAccountBearerTokenAuthUpdate: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *OpenAICodexIntegrationAccountAuthenticationUpdate) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into OpenAICodexIntegrationAccountBearerTokenAuthUpdate
	err = datadog.Unmarshal(data, &obj.OpenAICodexIntegrationAccountBearerTokenAuthUpdate)
	if err == nil {
		if obj.OpenAICodexIntegrationAccountBearerTokenAuthUpdate != nil && obj.OpenAICodexIntegrationAccountBearerTokenAuthUpdate.UnparsedObject == nil {
			jsonOpenAICodexIntegrationAccountBearerTokenAuthUpdate, _ := datadog.Marshal(obj.OpenAICodexIntegrationAccountBearerTokenAuthUpdate)
			if string(jsonOpenAICodexIntegrationAccountBearerTokenAuthUpdate) == "{}" { // empty struct
				obj.OpenAICodexIntegrationAccountBearerTokenAuthUpdate = nil
			} else {
				match++
			}
		} else {
			obj.OpenAICodexIntegrationAccountBearerTokenAuthUpdate = nil
		}
	} else {
		obj.OpenAICodexIntegrationAccountBearerTokenAuthUpdate = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.OpenAICodexIntegrationAccountBearerTokenAuthUpdate = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj OpenAICodexIntegrationAccountAuthenticationUpdate) MarshalJSON() ([]byte, error) {
	if obj.OpenAICodexIntegrationAccountBearerTokenAuthUpdate != nil {
		return datadog.Marshal(&obj.OpenAICodexIntegrationAccountBearerTokenAuthUpdate)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *OpenAICodexIntegrationAccountAuthenticationUpdate) GetActualInstance() interface{} {
	if obj.OpenAICodexIntegrationAccountBearerTokenAuthUpdate != nil {
		return obj.OpenAICodexIntegrationAccountBearerTokenAuthUpdate
	}

	// all schemas are nil
	return nil
}
