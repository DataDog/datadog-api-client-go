// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DatabricksIntegrationAccountAuthenticationResponse - Authentication configured on the Databricks integration account.
type DatabricksIntegrationAccountAuthenticationResponse struct {
	DatabricksIntegrationAccountOAuthAuthResponse     *DatabricksIntegrationAccountOAuthAuthResponse
	IntegrationAccountPrivateActionRunnerAuthResponse *IntegrationAccountPrivateActionRunnerAuthResponse
	DatabricksIntegrationAccountPatAuthResponse       *DatabricksIntegrationAccountPatAuthResponse

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// DatabricksIntegrationAccountOAuthAuthResponseAsDatabricksIntegrationAccountAuthenticationResponse is a convenience function that returns DatabricksIntegrationAccountOAuthAuthResponse wrapped in DatabricksIntegrationAccountAuthenticationResponse.
func DatabricksIntegrationAccountOAuthAuthResponseAsDatabricksIntegrationAccountAuthenticationResponse(v *DatabricksIntegrationAccountOAuthAuthResponse) DatabricksIntegrationAccountAuthenticationResponse {
	return DatabricksIntegrationAccountAuthenticationResponse{DatabricksIntegrationAccountOAuthAuthResponse: v}
}

// IntegrationAccountPrivateActionRunnerAuthResponseAsDatabricksIntegrationAccountAuthenticationResponse is a convenience function that returns IntegrationAccountPrivateActionRunnerAuthResponse wrapped in DatabricksIntegrationAccountAuthenticationResponse.
func IntegrationAccountPrivateActionRunnerAuthResponseAsDatabricksIntegrationAccountAuthenticationResponse(v *IntegrationAccountPrivateActionRunnerAuthResponse) DatabricksIntegrationAccountAuthenticationResponse {
	return DatabricksIntegrationAccountAuthenticationResponse{IntegrationAccountPrivateActionRunnerAuthResponse: v}
}

// DatabricksIntegrationAccountPatAuthResponseAsDatabricksIntegrationAccountAuthenticationResponse is a convenience function that returns DatabricksIntegrationAccountPatAuthResponse wrapped in DatabricksIntegrationAccountAuthenticationResponse.
func DatabricksIntegrationAccountPatAuthResponseAsDatabricksIntegrationAccountAuthenticationResponse(v *DatabricksIntegrationAccountPatAuthResponse) DatabricksIntegrationAccountAuthenticationResponse {
	return DatabricksIntegrationAccountAuthenticationResponse{DatabricksIntegrationAccountPatAuthResponse: v}
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

	// try to unmarshal data into IntegrationAccountPrivateActionRunnerAuthResponse
	err = datadog.Unmarshal(data, &obj.IntegrationAccountPrivateActionRunnerAuthResponse)
	if err == nil {
		if obj.IntegrationAccountPrivateActionRunnerAuthResponse != nil && obj.IntegrationAccountPrivateActionRunnerAuthResponse.UnparsedObject == nil {
			jsonIntegrationAccountPrivateActionRunnerAuthResponse, _ := datadog.Marshal(obj.IntegrationAccountPrivateActionRunnerAuthResponse)
			if string(jsonIntegrationAccountPrivateActionRunnerAuthResponse) == "{}" { // empty struct
				obj.IntegrationAccountPrivateActionRunnerAuthResponse = nil
			} else {
				match++
			}
		} else {
			obj.IntegrationAccountPrivateActionRunnerAuthResponse = nil
		}
	} else {
		obj.IntegrationAccountPrivateActionRunnerAuthResponse = nil
	}

	// try to unmarshal data into DatabricksIntegrationAccountPatAuthResponse
	err = datadog.Unmarshal(data, &obj.DatabricksIntegrationAccountPatAuthResponse)
	if err == nil {
		if obj.DatabricksIntegrationAccountPatAuthResponse != nil && obj.DatabricksIntegrationAccountPatAuthResponse.UnparsedObject == nil {
			jsonDatabricksIntegrationAccountPatAuthResponse, _ := datadog.Marshal(obj.DatabricksIntegrationAccountPatAuthResponse)
			if string(jsonDatabricksIntegrationAccountPatAuthResponse) == "{}" { // empty struct
				obj.DatabricksIntegrationAccountPatAuthResponse = nil
			} else {
				match++
			}
		} else {
			obj.DatabricksIntegrationAccountPatAuthResponse = nil
		}
	} else {
		obj.DatabricksIntegrationAccountPatAuthResponse = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.DatabricksIntegrationAccountOAuthAuthResponse = nil
		obj.IntegrationAccountPrivateActionRunnerAuthResponse = nil
		obj.DatabricksIntegrationAccountPatAuthResponse = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj DatabricksIntegrationAccountAuthenticationResponse) MarshalJSON() ([]byte, error) {
	if obj.DatabricksIntegrationAccountOAuthAuthResponse != nil {
		return datadog.Marshal(&obj.DatabricksIntegrationAccountOAuthAuthResponse)
	}

	if obj.IntegrationAccountPrivateActionRunnerAuthResponse != nil {
		return datadog.Marshal(&obj.IntegrationAccountPrivateActionRunnerAuthResponse)
	}

	if obj.DatabricksIntegrationAccountPatAuthResponse != nil {
		return datadog.Marshal(&obj.DatabricksIntegrationAccountPatAuthResponse)
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

	if obj.IntegrationAccountPrivateActionRunnerAuthResponse != nil {
		return obj.IntegrationAccountPrivateActionRunnerAuthResponse
	}

	if obj.DatabricksIntegrationAccountPatAuthResponse != nil {
		return obj.DatabricksIntegrationAccountPatAuthResponse
	}

	// all schemas are nil
	return nil
}
