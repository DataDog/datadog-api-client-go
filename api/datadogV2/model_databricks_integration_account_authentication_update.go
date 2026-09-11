// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DatabricksIntegrationAccountAuthenticationUpdate - Authentication for updating the Databricks integration account. Exactly one method is set. Choosing `private-action-runner` leaves the `databricks-model-serving-metrics` dataflow unable to collect data. `pat` is accepted only on accounts that already use it, so it cannot move an account onto personal access token authentication.
type DatabricksIntegrationAccountAuthenticationUpdate struct {
	DatabricksIntegrationAccountOAuthAuthRequest     *DatabricksIntegrationAccountOAuthAuthRequest
	IntegrationAccountPrivateActionRunnerAuthRequest *IntegrationAccountPrivateActionRunnerAuthRequest
	DatabricksIntegrationAccountPatAuthUpdate        *DatabricksIntegrationAccountPatAuthUpdate

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// DatabricksIntegrationAccountOAuthAuthRequestAsDatabricksIntegrationAccountAuthenticationUpdate is a convenience function that returns DatabricksIntegrationAccountOAuthAuthRequest wrapped in DatabricksIntegrationAccountAuthenticationUpdate.
func DatabricksIntegrationAccountOAuthAuthRequestAsDatabricksIntegrationAccountAuthenticationUpdate(v *DatabricksIntegrationAccountOAuthAuthRequest) DatabricksIntegrationAccountAuthenticationUpdate {
	return DatabricksIntegrationAccountAuthenticationUpdate{DatabricksIntegrationAccountOAuthAuthRequest: v}
}

// IntegrationAccountPrivateActionRunnerAuthRequestAsDatabricksIntegrationAccountAuthenticationUpdate is a convenience function that returns IntegrationAccountPrivateActionRunnerAuthRequest wrapped in DatabricksIntegrationAccountAuthenticationUpdate.
func IntegrationAccountPrivateActionRunnerAuthRequestAsDatabricksIntegrationAccountAuthenticationUpdate(v *IntegrationAccountPrivateActionRunnerAuthRequest) DatabricksIntegrationAccountAuthenticationUpdate {
	return DatabricksIntegrationAccountAuthenticationUpdate{IntegrationAccountPrivateActionRunnerAuthRequest: v}
}

// DatabricksIntegrationAccountPatAuthUpdateAsDatabricksIntegrationAccountAuthenticationUpdate is a convenience function that returns DatabricksIntegrationAccountPatAuthUpdate wrapped in DatabricksIntegrationAccountAuthenticationUpdate.
func DatabricksIntegrationAccountPatAuthUpdateAsDatabricksIntegrationAccountAuthenticationUpdate(v *DatabricksIntegrationAccountPatAuthUpdate) DatabricksIntegrationAccountAuthenticationUpdate {
	return DatabricksIntegrationAccountAuthenticationUpdate{DatabricksIntegrationAccountPatAuthUpdate: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *DatabricksIntegrationAccountAuthenticationUpdate) UnmarshalJSON(data []byte) error {
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

	// try to unmarshal data into DatabricksIntegrationAccountPatAuthUpdate
	err = datadog.Unmarshal(data, &obj.DatabricksIntegrationAccountPatAuthUpdate)
	if err == nil {
		if obj.DatabricksIntegrationAccountPatAuthUpdate != nil && obj.DatabricksIntegrationAccountPatAuthUpdate.UnparsedObject == nil {
			jsonDatabricksIntegrationAccountPatAuthUpdate, _ := datadog.Marshal(obj.DatabricksIntegrationAccountPatAuthUpdate)
			if string(jsonDatabricksIntegrationAccountPatAuthUpdate) == "{}" { // empty struct
				obj.DatabricksIntegrationAccountPatAuthUpdate = nil
			} else {
				match++
			}
		} else {
			obj.DatabricksIntegrationAccountPatAuthUpdate = nil
		}
	} else {
		obj.DatabricksIntegrationAccountPatAuthUpdate = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.DatabricksIntegrationAccountOAuthAuthRequest = nil
		obj.IntegrationAccountPrivateActionRunnerAuthRequest = nil
		obj.DatabricksIntegrationAccountPatAuthUpdate = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj DatabricksIntegrationAccountAuthenticationUpdate) MarshalJSON() ([]byte, error) {
	if obj.DatabricksIntegrationAccountOAuthAuthRequest != nil {
		return datadog.Marshal(&obj.DatabricksIntegrationAccountOAuthAuthRequest)
	}

	if obj.IntegrationAccountPrivateActionRunnerAuthRequest != nil {
		return datadog.Marshal(&obj.IntegrationAccountPrivateActionRunnerAuthRequest)
	}

	if obj.DatabricksIntegrationAccountPatAuthUpdate != nil {
		return datadog.Marshal(&obj.DatabricksIntegrationAccountPatAuthUpdate)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *DatabricksIntegrationAccountAuthenticationUpdate) GetActualInstance() interface{} {
	if obj.DatabricksIntegrationAccountOAuthAuthRequest != nil {
		return obj.DatabricksIntegrationAccountOAuthAuthRequest
	}

	if obj.IntegrationAccountPrivateActionRunnerAuthRequest != nil {
		return obj.IntegrationAccountPrivateActionRunnerAuthRequest
	}

	if obj.DatabricksIntegrationAccountPatAuthUpdate != nil {
		return obj.DatabricksIntegrationAccountPatAuthUpdate
	}

	// all schemas are nil
	return nil
}
