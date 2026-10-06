// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DatabricksIntegrationAccountAuthenticationRequest - Authentication for creating the Databricks integration account. Exactly one method is set. Choosing `private_action_runner` leaves the `databricks-model-serving-metrics` dataflow unable to collect data. `bearer_token` is not available on creation: Databricks accepts it only on accounts that already use it, so it can only appear on read and update.
type DatabricksIntegrationAccountAuthenticationRequest struct {
	DatabricksIntegrationAccountOAuthAuthRequest               *DatabricksIntegrationAccountOAuthAuthRequest
	DatabricksIntegrationAccountPrivateActionRunnerAuthRequest *DatabricksIntegrationAccountPrivateActionRunnerAuthRequest

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// DatabricksIntegrationAccountOAuthAuthRequestAsDatabricksIntegrationAccountAuthenticationRequest is a convenience function that returns DatabricksIntegrationAccountOAuthAuthRequest wrapped in DatabricksIntegrationAccountAuthenticationRequest.
func DatabricksIntegrationAccountOAuthAuthRequestAsDatabricksIntegrationAccountAuthenticationRequest(v *DatabricksIntegrationAccountOAuthAuthRequest) DatabricksIntegrationAccountAuthenticationRequest {
	return DatabricksIntegrationAccountAuthenticationRequest{DatabricksIntegrationAccountOAuthAuthRequest: v}
}

// DatabricksIntegrationAccountPrivateActionRunnerAuthRequestAsDatabricksIntegrationAccountAuthenticationRequest is a convenience function that returns DatabricksIntegrationAccountPrivateActionRunnerAuthRequest wrapped in DatabricksIntegrationAccountAuthenticationRequest.
func DatabricksIntegrationAccountPrivateActionRunnerAuthRequestAsDatabricksIntegrationAccountAuthenticationRequest(v *DatabricksIntegrationAccountPrivateActionRunnerAuthRequest) DatabricksIntegrationAccountAuthenticationRequest {
	return DatabricksIntegrationAccountAuthenticationRequest{DatabricksIntegrationAccountPrivateActionRunnerAuthRequest: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *DatabricksIntegrationAccountAuthenticationRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into DatabricksIntegrationAccountOAuthAuthRequest
	err = datadog.Unmarshal(data, &obj.DatabricksIntegrationAccountOAuthAuthRequest)
	if err == nil {
		if obj.DatabricksIntegrationAccountOAuthAuthRequest != nil && obj.DatabricksIntegrationAccountOAuthAuthRequest.UnparsedObject == nil {
			jsonDatabricksIntegrationAccountOAuthAuthRequest, _ := datadog.Marshal(obj.DatabricksIntegrationAccountOAuthAuthRequest)
			if string(jsonDatabricksIntegrationAccountOAuthAuthRequest) == "{}" { // empty struct
				obj.DatabricksIntegrationAccountOAuthAuthRequest = nil
			} else {
				match++
			}
		} else {
			obj.DatabricksIntegrationAccountOAuthAuthRequest = nil
		}
	} else {
		obj.DatabricksIntegrationAccountOAuthAuthRequest = nil
	}

	// try to unmarshal data into DatabricksIntegrationAccountPrivateActionRunnerAuthRequest
	err = datadog.Unmarshal(data, &obj.DatabricksIntegrationAccountPrivateActionRunnerAuthRequest)
	if err == nil {
		if obj.DatabricksIntegrationAccountPrivateActionRunnerAuthRequest != nil && obj.DatabricksIntegrationAccountPrivateActionRunnerAuthRequest.UnparsedObject == nil {
			jsonDatabricksIntegrationAccountPrivateActionRunnerAuthRequest, _ := datadog.Marshal(obj.DatabricksIntegrationAccountPrivateActionRunnerAuthRequest)
			if string(jsonDatabricksIntegrationAccountPrivateActionRunnerAuthRequest) == "{}" { // empty struct
				obj.DatabricksIntegrationAccountPrivateActionRunnerAuthRequest = nil
			} else {
				match++
			}
		} else {
			obj.DatabricksIntegrationAccountPrivateActionRunnerAuthRequest = nil
		}
	} else {
		obj.DatabricksIntegrationAccountPrivateActionRunnerAuthRequest = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.DatabricksIntegrationAccountOAuthAuthRequest = nil
		obj.DatabricksIntegrationAccountPrivateActionRunnerAuthRequest = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj DatabricksIntegrationAccountAuthenticationRequest) MarshalJSON() ([]byte, error) {
	if obj.DatabricksIntegrationAccountOAuthAuthRequest != nil {
		return datadog.Marshal(&obj.DatabricksIntegrationAccountOAuthAuthRequest)
	}

	if obj.DatabricksIntegrationAccountPrivateActionRunnerAuthRequest != nil {
		return datadog.Marshal(&obj.DatabricksIntegrationAccountPrivateActionRunnerAuthRequest)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *DatabricksIntegrationAccountAuthenticationRequest) GetActualInstance() interface{} {
	if obj.DatabricksIntegrationAccountOAuthAuthRequest != nil {
		return obj.DatabricksIntegrationAccountOAuthAuthRequest
	}

	if obj.DatabricksIntegrationAccountPrivateActionRunnerAuthRequest != nil {
		return obj.DatabricksIntegrationAccountPrivateActionRunnerAuthRequest
	}

	// all schemas are nil
	return nil
}
