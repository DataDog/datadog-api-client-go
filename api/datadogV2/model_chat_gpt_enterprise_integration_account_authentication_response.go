// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ChatGPTEnterpriseIntegrationAccountAuthenticationResponse - Authentication configured on the ChatGPT Enterprise integration account.
type ChatGPTEnterpriseIntegrationAccountAuthenticationResponse struct {
	ChatGPTEnterpriseIntegrationAccountBearerTokenAuthResponse *ChatGPTEnterpriseIntegrationAccountBearerTokenAuthResponse

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ChatGPTEnterpriseIntegrationAccountBearerTokenAuthResponseAsChatGPTEnterpriseIntegrationAccountAuthenticationResponse is a convenience function that returns ChatGPTEnterpriseIntegrationAccountBearerTokenAuthResponse wrapped in ChatGPTEnterpriseIntegrationAccountAuthenticationResponse.
func ChatGPTEnterpriseIntegrationAccountBearerTokenAuthResponseAsChatGPTEnterpriseIntegrationAccountAuthenticationResponse(v *ChatGPTEnterpriseIntegrationAccountBearerTokenAuthResponse) ChatGPTEnterpriseIntegrationAccountAuthenticationResponse {
	return ChatGPTEnterpriseIntegrationAccountAuthenticationResponse{ChatGPTEnterpriseIntegrationAccountBearerTokenAuthResponse: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ChatGPTEnterpriseIntegrationAccountAuthenticationResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ChatGPTEnterpriseIntegrationAccountBearerTokenAuthResponse
	err = datadog.Unmarshal(data, &obj.ChatGPTEnterpriseIntegrationAccountBearerTokenAuthResponse)
	if err == nil {
		if obj.ChatGPTEnterpriseIntegrationAccountBearerTokenAuthResponse != nil && obj.ChatGPTEnterpriseIntegrationAccountBearerTokenAuthResponse.UnparsedObject == nil {
			jsonChatGPTEnterpriseIntegrationAccountBearerTokenAuthResponse, _ := datadog.Marshal(obj.ChatGPTEnterpriseIntegrationAccountBearerTokenAuthResponse)
			if string(jsonChatGPTEnterpriseIntegrationAccountBearerTokenAuthResponse) == "{}" { // empty struct
				obj.ChatGPTEnterpriseIntegrationAccountBearerTokenAuthResponse = nil
			} else {
				match++
			}
		} else {
			obj.ChatGPTEnterpriseIntegrationAccountBearerTokenAuthResponse = nil
		}
	} else {
		obj.ChatGPTEnterpriseIntegrationAccountBearerTokenAuthResponse = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ChatGPTEnterpriseIntegrationAccountBearerTokenAuthResponse = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ChatGPTEnterpriseIntegrationAccountAuthenticationResponse) MarshalJSON() ([]byte, error) {
	if obj.ChatGPTEnterpriseIntegrationAccountBearerTokenAuthResponse != nil {
		return datadog.Marshal(&obj.ChatGPTEnterpriseIntegrationAccountBearerTokenAuthResponse)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ChatGPTEnterpriseIntegrationAccountAuthenticationResponse) GetActualInstance() interface{} {
	if obj.ChatGPTEnterpriseIntegrationAccountBearerTokenAuthResponse != nil {
		return obj.ChatGPTEnterpriseIntegrationAccountBearerTokenAuthResponse
	}

	// all schemas are nil
	return nil
}
