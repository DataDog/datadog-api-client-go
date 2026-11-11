// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ChatGPTEnterpriseIntegrationAccountAuthenticationRequest - Authentication for creating the ChatGPT Enterprise integration account. Exactly one method is set.
type ChatGPTEnterpriseIntegrationAccountAuthenticationRequest struct {
	ChatGPTEnterpriseIntegrationAccountBearerTokenAuthRequest *ChatGPTEnterpriseIntegrationAccountBearerTokenAuthRequest

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ChatGPTEnterpriseIntegrationAccountBearerTokenAuthRequestAsChatGPTEnterpriseIntegrationAccountAuthenticationRequest is a convenience function that returns ChatGPTEnterpriseIntegrationAccountBearerTokenAuthRequest wrapped in ChatGPTEnterpriseIntegrationAccountAuthenticationRequest.
func ChatGPTEnterpriseIntegrationAccountBearerTokenAuthRequestAsChatGPTEnterpriseIntegrationAccountAuthenticationRequest(v *ChatGPTEnterpriseIntegrationAccountBearerTokenAuthRequest) ChatGPTEnterpriseIntegrationAccountAuthenticationRequest {
	return ChatGPTEnterpriseIntegrationAccountAuthenticationRequest{ChatGPTEnterpriseIntegrationAccountBearerTokenAuthRequest: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ChatGPTEnterpriseIntegrationAccountAuthenticationRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ChatGPTEnterpriseIntegrationAccountBearerTokenAuthRequest
	err = datadog.Unmarshal(data, &obj.ChatGPTEnterpriseIntegrationAccountBearerTokenAuthRequest)
	if err == nil {
		if obj.ChatGPTEnterpriseIntegrationAccountBearerTokenAuthRequest != nil && obj.ChatGPTEnterpriseIntegrationAccountBearerTokenAuthRequest.UnparsedObject == nil {
			jsonChatGPTEnterpriseIntegrationAccountBearerTokenAuthRequest, _ := datadog.Marshal(obj.ChatGPTEnterpriseIntegrationAccountBearerTokenAuthRequest)
			if string(jsonChatGPTEnterpriseIntegrationAccountBearerTokenAuthRequest) == "{}" { // empty struct
				obj.ChatGPTEnterpriseIntegrationAccountBearerTokenAuthRequest = nil
			} else {
				match++
			}
		} else {
			obj.ChatGPTEnterpriseIntegrationAccountBearerTokenAuthRequest = nil
		}
	} else {
		obj.ChatGPTEnterpriseIntegrationAccountBearerTokenAuthRequest = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ChatGPTEnterpriseIntegrationAccountBearerTokenAuthRequest = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ChatGPTEnterpriseIntegrationAccountAuthenticationRequest) MarshalJSON() ([]byte, error) {
	if obj.ChatGPTEnterpriseIntegrationAccountBearerTokenAuthRequest != nil {
		return datadog.Marshal(&obj.ChatGPTEnterpriseIntegrationAccountBearerTokenAuthRequest)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ChatGPTEnterpriseIntegrationAccountAuthenticationRequest) GetActualInstance() interface{} {
	if obj.ChatGPTEnterpriseIntegrationAccountBearerTokenAuthRequest != nil {
		return obj.ChatGPTEnterpriseIntegrationAccountBearerTokenAuthRequest
	}

	// all schemas are nil
	return nil
}
