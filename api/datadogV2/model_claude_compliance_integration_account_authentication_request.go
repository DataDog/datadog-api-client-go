// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ClaudeComplianceIntegrationAccountAuthenticationRequest - Authentication for creating the Claude Compliance integration account. Exactly one method is set.
type ClaudeComplianceIntegrationAccountAuthenticationRequest struct {
	ClaudeComplianceIntegrationAccountBearerTokenAuthRequest *ClaudeComplianceIntegrationAccountBearerTokenAuthRequest

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ClaudeComplianceIntegrationAccountBearerTokenAuthRequestAsClaudeComplianceIntegrationAccountAuthenticationRequest is a convenience function that returns ClaudeComplianceIntegrationAccountBearerTokenAuthRequest wrapped in ClaudeComplianceIntegrationAccountAuthenticationRequest.
func ClaudeComplianceIntegrationAccountBearerTokenAuthRequestAsClaudeComplianceIntegrationAccountAuthenticationRequest(v *ClaudeComplianceIntegrationAccountBearerTokenAuthRequest) ClaudeComplianceIntegrationAccountAuthenticationRequest {
	return ClaudeComplianceIntegrationAccountAuthenticationRequest{ClaudeComplianceIntegrationAccountBearerTokenAuthRequest: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ClaudeComplianceIntegrationAccountAuthenticationRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ClaudeComplianceIntegrationAccountBearerTokenAuthRequest
	err = datadog.Unmarshal(data, &obj.ClaudeComplianceIntegrationAccountBearerTokenAuthRequest)
	if err == nil {
		if obj.ClaudeComplianceIntegrationAccountBearerTokenAuthRequest != nil && obj.ClaudeComplianceIntegrationAccountBearerTokenAuthRequest.UnparsedObject == nil {
			jsonClaudeComplianceIntegrationAccountBearerTokenAuthRequest, _ := datadog.Marshal(obj.ClaudeComplianceIntegrationAccountBearerTokenAuthRequest)
			if string(jsonClaudeComplianceIntegrationAccountBearerTokenAuthRequest) == "{}" { // empty struct
				obj.ClaudeComplianceIntegrationAccountBearerTokenAuthRequest = nil
			} else {
				match++
			}
		} else {
			obj.ClaudeComplianceIntegrationAccountBearerTokenAuthRequest = nil
		}
	} else {
		obj.ClaudeComplianceIntegrationAccountBearerTokenAuthRequest = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ClaudeComplianceIntegrationAccountBearerTokenAuthRequest = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ClaudeComplianceIntegrationAccountAuthenticationRequest) MarshalJSON() ([]byte, error) {
	if obj.ClaudeComplianceIntegrationAccountBearerTokenAuthRequest != nil {
		return datadog.Marshal(&obj.ClaudeComplianceIntegrationAccountBearerTokenAuthRequest)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ClaudeComplianceIntegrationAccountAuthenticationRequest) GetActualInstance() interface{} {
	if obj.ClaudeComplianceIntegrationAccountBearerTokenAuthRequest != nil {
		return obj.ClaudeComplianceIntegrationAccountBearerTokenAuthRequest
	}

	// all schemas are nil
	return nil
}
