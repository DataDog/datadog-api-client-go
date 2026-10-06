// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DatabricksIntegrationAccountAuthenticationResponse - Authentication configured on the Databricks integration account. A `bearer_token` method indicates an account still on token authentication, which Databricks accepts only on accounts that already use it.
type DatabricksIntegrationAccountAuthenticationResponse struct {
	DatabricksIntegrationAccountOAuthAuthResponse               *DatabricksIntegrationAccountOAuthAuthResponse
	DatabricksIntegrationAccountPrivateActionRunnerAuthResponse *DatabricksIntegrationAccountPrivateActionRunnerAuthResponse
	DatabricksIntegrationAccountBearerTokenAuthResponse         *DatabricksIntegrationAccountBearerTokenAuthResponse

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// DatabricksIntegrationAccountOAuthAuthResponseAsDatabricksIntegrationAccountAuthenticationResponse is a convenience function that returns DatabricksIntegrationAccountOAuthAuthResponse wrapped in DatabricksIntegrationAccountAuthenticationResponse.
func DatabricksIntegrationAccountOAuthAuthResponseAsDatabricksIntegrationAccountAuthenticationResponse(v *DatabricksIntegrationAccountOAuthAuthResponse) DatabricksIntegrationAccountAuthenticationResponse {
	return DatabricksIntegrationAccountAuthenticationResponse{DatabricksIntegrationAccountOAuthAuthResponse: v}
}

// DatabricksIntegrationAccountPrivateActionRunnerAuthResponseAsDatabricksIntegrationAccountAuthenticationResponse is a convenience function that returns DatabricksIntegrationAccountPrivateActionRunnerAuthResponse wrapped in DatabricksIntegrationAccountAuthenticationResponse.
func DatabricksIntegrationAccountPrivateActionRunnerAuthResponseAsDatabricksIntegrationAccountAuthenticationResponse(v *DatabricksIntegrationAccountPrivateActionRunnerAuthResponse) DatabricksIntegrationAccountAuthenticationResponse {
	return DatabricksIntegrationAccountAuthenticationResponse{DatabricksIntegrationAccountPrivateActionRunnerAuthResponse: v}
}

// DatabricksIntegrationAccountBearerTokenAuthResponseAsDatabricksIntegrationAccountAuthenticationResponse is a convenience function that returns DatabricksIntegrationAccountBearerTokenAuthResponse wrapped in DatabricksIntegrationAccountAuthenticationResponse.
func DatabricksIntegrationAccountBearerTokenAuthResponseAsDatabricksIntegrationAccountAuthenticationResponse(v *DatabricksIntegrationAccountBearerTokenAuthResponse) DatabricksIntegrationAccountAuthenticationResponse {
	return DatabricksIntegrationAccountAuthenticationResponse{DatabricksIntegrationAccountBearerTokenAuthResponse: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *DatabricksIntegrationAccountAuthenticationResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into DatabricksIntegrationAccountOAuthAuthResponse
	err = datadog.Unmarshal(data, &obj.DatabricksIntegrationAccountOAuthAuthResponse)
	if err == nil {
		if obj.DatabricksIntegrationAccountOAuthAuthResponse != nil && obj.DatabricksIntegrationAccountOAuthAuthResponse.UnparsedObject == nil {
			jsonDatabricksIntegrationAccountOAuthAuthResponse, _ := datadog.Marshal(obj.DatabricksIntegrationAccountOAuthAuthResponse)
			if string(jsonDatabricksIntegrationAccountOAuthAuthResponse) == "{}" { // empty struct
				obj.DatabricksIntegrationAccountOAuthAuthResponse = nil
			} else {
				match++
			}
		} else {
			obj.DatabricksIntegrationAccountOAuthAuthResponse = nil
		}
	} else {
		obj.DatabricksIntegrationAccountOAuthAuthResponse = nil
	}

	// try to unmarshal data into DatabricksIntegrationAccountPrivateActionRunnerAuthResponse
	err = datadog.Unmarshal(data, &obj.DatabricksIntegrationAccountPrivateActionRunnerAuthResponse)
	if err == nil {
		if obj.DatabricksIntegrationAccountPrivateActionRunnerAuthResponse != nil && obj.DatabricksIntegrationAccountPrivateActionRunnerAuthResponse.UnparsedObject == nil {
			jsonDatabricksIntegrationAccountPrivateActionRunnerAuthResponse, _ := datadog.Marshal(obj.DatabricksIntegrationAccountPrivateActionRunnerAuthResponse)
			if string(jsonDatabricksIntegrationAccountPrivateActionRunnerAuthResponse) == "{}" { // empty struct
				obj.DatabricksIntegrationAccountPrivateActionRunnerAuthResponse = nil
			} else {
				match++
			}
		} else {
			obj.DatabricksIntegrationAccountPrivateActionRunnerAuthResponse = nil
		}
	} else {
		obj.DatabricksIntegrationAccountPrivateActionRunnerAuthResponse = nil
	}

	// try to unmarshal data into DatabricksIntegrationAccountBearerTokenAuthResponse
	err = datadog.Unmarshal(data, &obj.DatabricksIntegrationAccountBearerTokenAuthResponse)
	if err == nil {
		if obj.DatabricksIntegrationAccountBearerTokenAuthResponse != nil && obj.DatabricksIntegrationAccountBearerTokenAuthResponse.UnparsedObject == nil {
			jsonDatabricksIntegrationAccountBearerTokenAuthResponse, _ := datadog.Marshal(obj.DatabricksIntegrationAccountBearerTokenAuthResponse)
			if string(jsonDatabricksIntegrationAccountBearerTokenAuthResponse) == "{}" { // empty struct
				obj.DatabricksIntegrationAccountBearerTokenAuthResponse = nil
			} else {
				match++
			}
		} else {
			obj.DatabricksIntegrationAccountBearerTokenAuthResponse = nil
		}
	} else {
		obj.DatabricksIntegrationAccountBearerTokenAuthResponse = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.DatabricksIntegrationAccountOAuthAuthResponse = nil
		obj.DatabricksIntegrationAccountPrivateActionRunnerAuthResponse = nil
		obj.DatabricksIntegrationAccountBearerTokenAuthResponse = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj DatabricksIntegrationAccountAuthenticationResponse) MarshalJSON() ([]byte, error) {
	if obj.DatabricksIntegrationAccountOAuthAuthResponse != nil {
		return datadog.Marshal(&obj.DatabricksIntegrationAccountOAuthAuthResponse)
	}

	if obj.DatabricksIntegrationAccountPrivateActionRunnerAuthResponse != nil {
		return datadog.Marshal(&obj.DatabricksIntegrationAccountPrivateActionRunnerAuthResponse)
	}

	if obj.DatabricksIntegrationAccountBearerTokenAuthResponse != nil {
		return datadog.Marshal(&obj.DatabricksIntegrationAccountBearerTokenAuthResponse)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *DatabricksIntegrationAccountAuthenticationResponse) GetActualInstance() interface{} {
	if obj.DatabricksIntegrationAccountOAuthAuthResponse != nil {
		return obj.DatabricksIntegrationAccountOAuthAuthResponse
	}

	if obj.DatabricksIntegrationAccountPrivateActionRunnerAuthResponse != nil {
		return obj.DatabricksIntegrationAccountPrivateActionRunnerAuthResponse
	}

	if obj.DatabricksIntegrationAccountBearerTokenAuthResponse != nil {
		return obj.DatabricksIntegrationAccountBearerTokenAuthResponse
	}

	// all schemas are nil
	return nil
}
