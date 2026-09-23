// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DatabricksIntegrationAccountAuthenticationUpdate - Authentication for updating the Databricks integration account. Exactly one method is set. Choosing `private_action_runner` leaves the `databricks-model-serving-metrics` dataflow unable to collect data. `bearer_token` is deprecated on Databricks: it is accepted only on accounts that already use it and never on creation, so it cannot move an account onto token authentication. Migrate those accounts to `databricks_oauth` or `private_action_runner`.
type DatabricksIntegrationAccountAuthenticationUpdate struct {
	DatabricksIntegrationAccountOAuthAuthUpdate               *DatabricksIntegrationAccountOAuthAuthUpdate
	DatabricksIntegrationAccountPrivateActionRunnerAuthUpdate *DatabricksIntegrationAccountPrivateActionRunnerAuthUpdate
	DatabricksIntegrationAccountBearerTokenAuthUpdate         *DatabricksIntegrationAccountBearerTokenAuthUpdate

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// DatabricksIntegrationAccountOAuthAuthUpdateAsDatabricksIntegrationAccountAuthenticationUpdate is a convenience function that returns DatabricksIntegrationAccountOAuthAuthUpdate wrapped in DatabricksIntegrationAccountAuthenticationUpdate.
func DatabricksIntegrationAccountOAuthAuthUpdateAsDatabricksIntegrationAccountAuthenticationUpdate(v *DatabricksIntegrationAccountOAuthAuthUpdate) DatabricksIntegrationAccountAuthenticationUpdate {
	return DatabricksIntegrationAccountAuthenticationUpdate{DatabricksIntegrationAccountOAuthAuthUpdate: v}
}

// DatabricksIntegrationAccountPrivateActionRunnerAuthUpdateAsDatabricksIntegrationAccountAuthenticationUpdate is a convenience function that returns DatabricksIntegrationAccountPrivateActionRunnerAuthUpdate wrapped in DatabricksIntegrationAccountAuthenticationUpdate.
func DatabricksIntegrationAccountPrivateActionRunnerAuthUpdateAsDatabricksIntegrationAccountAuthenticationUpdate(v *DatabricksIntegrationAccountPrivateActionRunnerAuthUpdate) DatabricksIntegrationAccountAuthenticationUpdate {
	return DatabricksIntegrationAccountAuthenticationUpdate{DatabricksIntegrationAccountPrivateActionRunnerAuthUpdate: v}
}

// DatabricksIntegrationAccountBearerTokenAuthUpdateAsDatabricksIntegrationAccountAuthenticationUpdate is a convenience function that returns DatabricksIntegrationAccountBearerTokenAuthUpdate wrapped in DatabricksIntegrationAccountAuthenticationUpdate.
func DatabricksIntegrationAccountBearerTokenAuthUpdateAsDatabricksIntegrationAccountAuthenticationUpdate(v *DatabricksIntegrationAccountBearerTokenAuthUpdate) DatabricksIntegrationAccountAuthenticationUpdate {
	return DatabricksIntegrationAccountAuthenticationUpdate{DatabricksIntegrationAccountBearerTokenAuthUpdate: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *DatabricksIntegrationAccountAuthenticationUpdate) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into DatabricksIntegrationAccountOAuthAuthUpdate
	err = datadog.Unmarshal(data, &obj.DatabricksIntegrationAccountOAuthAuthUpdate)
	if err == nil {
		if obj.DatabricksIntegrationAccountOAuthAuthUpdate != nil && obj.DatabricksIntegrationAccountOAuthAuthUpdate.UnparsedObject == nil {
			jsonDatabricksIntegrationAccountOAuthAuthUpdate, _ := datadog.Marshal(obj.DatabricksIntegrationAccountOAuthAuthUpdate)
			if string(jsonDatabricksIntegrationAccountOAuthAuthUpdate) == "{}" { // empty struct
				obj.DatabricksIntegrationAccountOAuthAuthUpdate = nil
			} else {
				match++
			}
		} else {
			obj.DatabricksIntegrationAccountOAuthAuthUpdate = nil
		}
	} else {
		obj.DatabricksIntegrationAccountOAuthAuthUpdate = nil
	}

	// try to unmarshal data into DatabricksIntegrationAccountPrivateActionRunnerAuthUpdate
	err = datadog.Unmarshal(data, &obj.DatabricksIntegrationAccountPrivateActionRunnerAuthUpdate)
	if err == nil {
		if obj.DatabricksIntegrationAccountPrivateActionRunnerAuthUpdate != nil && obj.DatabricksIntegrationAccountPrivateActionRunnerAuthUpdate.UnparsedObject == nil {
			jsonDatabricksIntegrationAccountPrivateActionRunnerAuthUpdate, _ := datadog.Marshal(obj.DatabricksIntegrationAccountPrivateActionRunnerAuthUpdate)
			if string(jsonDatabricksIntegrationAccountPrivateActionRunnerAuthUpdate) == "{}" { // empty struct
				obj.DatabricksIntegrationAccountPrivateActionRunnerAuthUpdate = nil
			} else {
				match++
			}
		} else {
			obj.DatabricksIntegrationAccountPrivateActionRunnerAuthUpdate = nil
		}
	} else {
		obj.DatabricksIntegrationAccountPrivateActionRunnerAuthUpdate = nil
	}

	// try to unmarshal data into DatabricksIntegrationAccountBearerTokenAuthUpdate
	err = datadog.Unmarshal(data, &obj.DatabricksIntegrationAccountBearerTokenAuthUpdate)
	if err == nil {
		if obj.DatabricksIntegrationAccountBearerTokenAuthUpdate != nil && obj.DatabricksIntegrationAccountBearerTokenAuthUpdate.UnparsedObject == nil {
			jsonDatabricksIntegrationAccountBearerTokenAuthUpdate, _ := datadog.Marshal(obj.DatabricksIntegrationAccountBearerTokenAuthUpdate)
			if string(jsonDatabricksIntegrationAccountBearerTokenAuthUpdate) == "{}" { // empty struct
				obj.DatabricksIntegrationAccountBearerTokenAuthUpdate = nil
			} else {
				match++
			}
		} else {
			obj.DatabricksIntegrationAccountBearerTokenAuthUpdate = nil
		}
	} else {
		obj.DatabricksIntegrationAccountBearerTokenAuthUpdate = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.DatabricksIntegrationAccountOAuthAuthUpdate = nil
		obj.DatabricksIntegrationAccountPrivateActionRunnerAuthUpdate = nil
		obj.DatabricksIntegrationAccountBearerTokenAuthUpdate = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj DatabricksIntegrationAccountAuthenticationUpdate) MarshalJSON() ([]byte, error) {
	if obj.DatabricksIntegrationAccountOAuthAuthUpdate != nil {
		return datadog.Marshal(&obj.DatabricksIntegrationAccountOAuthAuthUpdate)
	}

	if obj.DatabricksIntegrationAccountPrivateActionRunnerAuthUpdate != nil {
		return datadog.Marshal(&obj.DatabricksIntegrationAccountPrivateActionRunnerAuthUpdate)
	}

	if obj.DatabricksIntegrationAccountBearerTokenAuthUpdate != nil {
		return datadog.Marshal(&obj.DatabricksIntegrationAccountBearerTokenAuthUpdate)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *DatabricksIntegrationAccountAuthenticationUpdate) GetActualInstance() interface{} {
	if obj.DatabricksIntegrationAccountOAuthAuthUpdate != nil {
		return obj.DatabricksIntegrationAccountOAuthAuthUpdate
	}

	if obj.DatabricksIntegrationAccountPrivateActionRunnerAuthUpdate != nil {
		return obj.DatabricksIntegrationAccountPrivateActionRunnerAuthUpdate
	}

	if obj.DatabricksIntegrationAccountBearerTokenAuthUpdate != nil {
		return obj.DatabricksIntegrationAccountBearerTokenAuthUpdate
	}

	// all schemas are nil
	return nil
}
