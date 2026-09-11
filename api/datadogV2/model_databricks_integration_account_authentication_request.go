// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DatabricksIntegrationAccountAuthenticationRequest - Authentication for creating the Databricks integration account. Exactly one method is set. Choosing `private-action-runner` leaves the `databricks-model-serving-metrics` dataflow unable to collect data.
type DatabricksIntegrationAccountAuthenticationRequest struct {
	DatabricksIntegrationAccountOAuthAuthRequest     *DatabricksIntegrationAccountOAuthAuthRequest
	IntegrationAccountPrivateActionRunnerAuthRequest *IntegrationAccountPrivateActionRunnerAuthRequest

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// DatabricksIntegrationAccountOAuthAuthRequestAsDatabricksIntegrationAccountAuthenticationRequest is a convenience function that returns DatabricksIntegrationAccountOAuthAuthRequest wrapped in DatabricksIntegrationAccountAuthenticationRequest.
func DatabricksIntegrationAccountOAuthAuthRequestAsDatabricksIntegrationAccountAuthenticationRequest(v *DatabricksIntegrationAccountOAuthAuthRequest) DatabricksIntegrationAccountAuthenticationRequest {
	return DatabricksIntegrationAccountAuthenticationRequest{DatabricksIntegrationAccountOAuthAuthRequest: v}
}

// IntegrationAccountPrivateActionRunnerAuthRequestAsDatabricksIntegrationAccountAuthenticationRequest is a convenience function that returns IntegrationAccountPrivateActionRunnerAuthRequest wrapped in DatabricksIntegrationAccountAuthenticationRequest.
func IntegrationAccountPrivateActionRunnerAuthRequestAsDatabricksIntegrationAccountAuthenticationRequest(v *IntegrationAccountPrivateActionRunnerAuthRequest) DatabricksIntegrationAccountAuthenticationRequest {
	return DatabricksIntegrationAccountAuthenticationRequest{IntegrationAccountPrivateActionRunnerAuthRequest: v}
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

	// try to unmarshal data into IntegrationAccountPrivateActionRunnerAuthRequest
	err = datadog.Unmarshal(data, &obj.IntegrationAccountPrivateActionRunnerAuthRequest)
	if err == nil {
		if obj.IntegrationAccountPrivateActionRunnerAuthRequest != nil && obj.IntegrationAccountPrivateActionRunnerAuthRequest.UnparsedObject == nil {
			jsonIntegrationAccountPrivateActionRunnerAuthRequest, _ := datadog.Marshal(obj.IntegrationAccountPrivateActionRunnerAuthRequest)
			if string(jsonIntegrationAccountPrivateActionRunnerAuthRequest) == "{}" { // empty struct
				obj.IntegrationAccountPrivateActionRunnerAuthRequest = nil
			} else {
				match++
			}
		} else {
			obj.IntegrationAccountPrivateActionRunnerAuthRequest = nil
		}
	} else {
		obj.IntegrationAccountPrivateActionRunnerAuthRequest = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.DatabricksIntegrationAccountOAuthAuthRequest = nil
		obj.IntegrationAccountPrivateActionRunnerAuthRequest = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj DatabricksIntegrationAccountAuthenticationRequest) MarshalJSON() ([]byte, error) {
	if obj.DatabricksIntegrationAccountOAuthAuthRequest != nil {
		return datadog.Marshal(&obj.DatabricksIntegrationAccountOAuthAuthRequest)
	}

	if obj.IntegrationAccountPrivateActionRunnerAuthRequest != nil {
		return datadog.Marshal(&obj.IntegrationAccountPrivateActionRunnerAuthRequest)
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

	if obj.IntegrationAccountPrivateActionRunnerAuthRequest != nil {
		return obj.IntegrationAccountPrivateActionRunnerAuthRequest
	}

	// all schemas are nil
	return nil
}
