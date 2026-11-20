// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ClaudeComplianceIntegrationAccountAuthenticationUpdate - Authentication for updating the Claude Compliance integration account. Exactly one method is set.
type ClaudeComplianceIntegrationAccountAuthenticationUpdate struct {
	ClaudeComplianceIntegrationAccountBearerTokenAuthUpdate *ClaudeComplianceIntegrationAccountBearerTokenAuthUpdate

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ClaudeComplianceIntegrationAccountBearerTokenAuthUpdateAsClaudeComplianceIntegrationAccountAuthenticationUpdate is a convenience function that returns ClaudeComplianceIntegrationAccountBearerTokenAuthUpdate wrapped in ClaudeComplianceIntegrationAccountAuthenticationUpdate.
func ClaudeComplianceIntegrationAccountBearerTokenAuthUpdateAsClaudeComplianceIntegrationAccountAuthenticationUpdate(v *ClaudeComplianceIntegrationAccountBearerTokenAuthUpdate) ClaudeComplianceIntegrationAccountAuthenticationUpdate {
	return ClaudeComplianceIntegrationAccountAuthenticationUpdate{ClaudeComplianceIntegrationAccountBearerTokenAuthUpdate: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ClaudeComplianceIntegrationAccountAuthenticationUpdate) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ClaudeComplianceIntegrationAccountBearerTokenAuthUpdate
	err = datadog.Unmarshal(data, &obj.ClaudeComplianceIntegrationAccountBearerTokenAuthUpdate)
	if err == nil {
		if obj.ClaudeComplianceIntegrationAccountBearerTokenAuthUpdate != nil && obj.ClaudeComplianceIntegrationAccountBearerTokenAuthUpdate.UnparsedObject == nil {
			jsonClaudeComplianceIntegrationAccountBearerTokenAuthUpdate, _ := datadog.Marshal(obj.ClaudeComplianceIntegrationAccountBearerTokenAuthUpdate)
			if string(jsonClaudeComplianceIntegrationAccountBearerTokenAuthUpdate) == "{}" { // empty struct
				obj.ClaudeComplianceIntegrationAccountBearerTokenAuthUpdate = nil
			} else {
				match++
			}
		} else {
			obj.ClaudeComplianceIntegrationAccountBearerTokenAuthUpdate = nil
		}
	} else {
		obj.ClaudeComplianceIntegrationAccountBearerTokenAuthUpdate = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ClaudeComplianceIntegrationAccountBearerTokenAuthUpdate = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ClaudeComplianceIntegrationAccountAuthenticationUpdate) MarshalJSON() ([]byte, error) {
	if obj.ClaudeComplianceIntegrationAccountBearerTokenAuthUpdate != nil {
		return datadog.Marshal(&obj.ClaudeComplianceIntegrationAccountBearerTokenAuthUpdate)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ClaudeComplianceIntegrationAccountAuthenticationUpdate) GetActualInstance() interface{} {
	if obj.ClaudeComplianceIntegrationAccountBearerTokenAuthUpdate != nil {
		return obj.ClaudeComplianceIntegrationAccountBearerTokenAuthUpdate
	}

	// all schemas are nil
	return nil
}
