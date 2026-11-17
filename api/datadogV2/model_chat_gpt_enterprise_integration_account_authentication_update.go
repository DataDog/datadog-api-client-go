// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ChatGPTEnterpriseIntegrationAccountAuthenticationUpdate - Authentication for updating the ChatGPT Enterprise integration account. Exactly one method is set.
type ChatGPTEnterpriseIntegrationAccountAuthenticationUpdate struct {
	ChatGPTEnterpriseIntegrationAccountBearerTokenAuthUpdate *ChatGPTEnterpriseIntegrationAccountBearerTokenAuthUpdate

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ChatGPTEnterpriseIntegrationAccountBearerTokenAuthUpdateAsChatGPTEnterpriseIntegrationAccountAuthenticationUpdate is a convenience function that returns ChatGPTEnterpriseIntegrationAccountBearerTokenAuthUpdate wrapped in ChatGPTEnterpriseIntegrationAccountAuthenticationUpdate.
func ChatGPTEnterpriseIntegrationAccountBearerTokenAuthUpdateAsChatGPTEnterpriseIntegrationAccountAuthenticationUpdate(v *ChatGPTEnterpriseIntegrationAccountBearerTokenAuthUpdate) ChatGPTEnterpriseIntegrationAccountAuthenticationUpdate {
	return ChatGPTEnterpriseIntegrationAccountAuthenticationUpdate{ChatGPTEnterpriseIntegrationAccountBearerTokenAuthUpdate: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ChatGPTEnterpriseIntegrationAccountAuthenticationUpdate) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ChatGPTEnterpriseIntegrationAccountBearerTokenAuthUpdate
	err = datadog.Unmarshal(data, &obj.ChatGPTEnterpriseIntegrationAccountBearerTokenAuthUpdate)
	if err == nil {
		if obj.ChatGPTEnterpriseIntegrationAccountBearerTokenAuthUpdate != nil && obj.ChatGPTEnterpriseIntegrationAccountBearerTokenAuthUpdate.UnparsedObject == nil {
			jsonChatGPTEnterpriseIntegrationAccountBearerTokenAuthUpdate, _ := datadog.Marshal(obj.ChatGPTEnterpriseIntegrationAccountBearerTokenAuthUpdate)
			if string(jsonChatGPTEnterpriseIntegrationAccountBearerTokenAuthUpdate) == "{}" { // empty struct
				obj.ChatGPTEnterpriseIntegrationAccountBearerTokenAuthUpdate = nil
			} else {
				match++
			}
		} else {
			obj.ChatGPTEnterpriseIntegrationAccountBearerTokenAuthUpdate = nil
		}
	} else {
		obj.ChatGPTEnterpriseIntegrationAccountBearerTokenAuthUpdate = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ChatGPTEnterpriseIntegrationAccountBearerTokenAuthUpdate = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ChatGPTEnterpriseIntegrationAccountAuthenticationUpdate) MarshalJSON() ([]byte, error) {
	if obj.ChatGPTEnterpriseIntegrationAccountBearerTokenAuthUpdate != nil {
		return datadog.Marshal(&obj.ChatGPTEnterpriseIntegrationAccountBearerTokenAuthUpdate)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ChatGPTEnterpriseIntegrationAccountAuthenticationUpdate) GetActualInstance() interface{} {
	if obj.ChatGPTEnterpriseIntegrationAccountBearerTokenAuthUpdate != nil {
		return obj.ChatGPTEnterpriseIntegrationAccountBearerTokenAuthUpdate
	}

	// all schemas are nil
	return nil
}
