// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GitHubCostsIntegrationAccountAuthenticationResponse - Authentication configured on the GitHub Costs integration account.
type GitHubCostsIntegrationAccountAuthenticationResponse struct {
	GitHubCostsIntegrationAccountBearerTokenAuthResponse *GitHubCostsIntegrationAccountBearerTokenAuthResponse

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// GitHubCostsIntegrationAccountBearerTokenAuthResponseAsGitHubCostsIntegrationAccountAuthenticationResponse is a convenience function that returns GitHubCostsIntegrationAccountBearerTokenAuthResponse wrapped in GitHubCostsIntegrationAccountAuthenticationResponse.
func GitHubCostsIntegrationAccountBearerTokenAuthResponseAsGitHubCostsIntegrationAccountAuthenticationResponse(v *GitHubCostsIntegrationAccountBearerTokenAuthResponse) GitHubCostsIntegrationAccountAuthenticationResponse {
	return GitHubCostsIntegrationAccountAuthenticationResponse{GitHubCostsIntegrationAccountBearerTokenAuthResponse: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *GitHubCostsIntegrationAccountAuthenticationResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into GitHubCostsIntegrationAccountBearerTokenAuthResponse
	err = datadog.Unmarshal(data, &obj.GitHubCostsIntegrationAccountBearerTokenAuthResponse)
	if err == nil {
		if obj.GitHubCostsIntegrationAccountBearerTokenAuthResponse != nil && obj.GitHubCostsIntegrationAccountBearerTokenAuthResponse.UnparsedObject == nil {
			jsonGitHubCostsIntegrationAccountBearerTokenAuthResponse, _ := datadog.Marshal(obj.GitHubCostsIntegrationAccountBearerTokenAuthResponse)
			if string(jsonGitHubCostsIntegrationAccountBearerTokenAuthResponse) == "{}" { // empty struct
				obj.GitHubCostsIntegrationAccountBearerTokenAuthResponse = nil
			} else {
				match++
			}
		} else {
			obj.GitHubCostsIntegrationAccountBearerTokenAuthResponse = nil
		}
	} else {
		obj.GitHubCostsIntegrationAccountBearerTokenAuthResponse = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.GitHubCostsIntegrationAccountBearerTokenAuthResponse = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj GitHubCostsIntegrationAccountAuthenticationResponse) MarshalJSON() ([]byte, error) {
	if obj.GitHubCostsIntegrationAccountBearerTokenAuthResponse != nil {
		return datadog.Marshal(&obj.GitHubCostsIntegrationAccountBearerTokenAuthResponse)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *GitHubCostsIntegrationAccountAuthenticationResponse) GetActualInstance() interface{} {
	if obj.GitHubCostsIntegrationAccountBearerTokenAuthResponse != nil {
		return obj.GitHubCostsIntegrationAccountBearerTokenAuthResponse
	}

	// all schemas are nil
	return nil
}
