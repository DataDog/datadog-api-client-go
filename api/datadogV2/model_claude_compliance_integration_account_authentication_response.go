// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ClaudeComplianceIntegrationAccountAuthenticationResponse - Authentication configured on the Claude Compliance integration account.
type ClaudeComplianceIntegrationAccountAuthenticationResponse struct {
	ClaudeComplianceIntegrationAccountBearerTokenAuthResponse *ClaudeComplianceIntegrationAccountBearerTokenAuthResponse

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ClaudeComplianceIntegrationAccountBearerTokenAuthResponseAsClaudeComplianceIntegrationAccountAuthenticationResponse is a convenience function that returns ClaudeComplianceIntegrationAccountBearerTokenAuthResponse wrapped in ClaudeComplianceIntegrationAccountAuthenticationResponse.
func ClaudeComplianceIntegrationAccountBearerTokenAuthResponseAsClaudeComplianceIntegrationAccountAuthenticationResponse(v *ClaudeComplianceIntegrationAccountBearerTokenAuthResponse) ClaudeComplianceIntegrationAccountAuthenticationResponse {
	return ClaudeComplianceIntegrationAccountAuthenticationResponse{ClaudeComplianceIntegrationAccountBearerTokenAuthResponse: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ClaudeComplianceIntegrationAccountAuthenticationResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ClaudeComplianceIntegrationAccountBearerTokenAuthResponse
	err = datadog.Unmarshal(data, &obj.ClaudeComplianceIntegrationAccountBearerTokenAuthResponse)
	if err == nil {
		if obj.ClaudeComplianceIntegrationAccountBearerTokenAuthResponse != nil && obj.ClaudeComplianceIntegrationAccountBearerTokenAuthResponse.UnparsedObject == nil {
			jsonClaudeComplianceIntegrationAccountBearerTokenAuthResponse, _ := datadog.Marshal(obj.ClaudeComplianceIntegrationAccountBearerTokenAuthResponse)
			if string(jsonClaudeComplianceIntegrationAccountBearerTokenAuthResponse) == "{}" { // empty struct
				obj.ClaudeComplianceIntegrationAccountBearerTokenAuthResponse = nil
			} else {
				match++
			}
		} else {
			obj.ClaudeComplianceIntegrationAccountBearerTokenAuthResponse = nil
		}
	} else {
		obj.ClaudeComplianceIntegrationAccountBearerTokenAuthResponse = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ClaudeComplianceIntegrationAccountBearerTokenAuthResponse = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ClaudeComplianceIntegrationAccountAuthenticationResponse) MarshalJSON() ([]byte, error) {
	if obj.ClaudeComplianceIntegrationAccountBearerTokenAuthResponse != nil {
		return datadog.Marshal(&obj.ClaudeComplianceIntegrationAccountBearerTokenAuthResponse)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ClaudeComplianceIntegrationAccountAuthenticationResponse) GetActualInstance() interface{} {
	if obj.ClaudeComplianceIntegrationAccountBearerTokenAuthResponse != nil {
		return obj.ClaudeComplianceIntegrationAccountBearerTokenAuthResponse
	}

	// all schemas are nil
	return nil
}
