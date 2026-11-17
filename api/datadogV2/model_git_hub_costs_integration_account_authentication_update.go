// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GitHubCostsIntegrationAccountAuthenticationUpdate - Authentication for updating the GitHub Costs integration account. Exactly one method is set.
type GitHubCostsIntegrationAccountAuthenticationUpdate struct {
	GitHubCostsIntegrationAccountBearerTokenAuthUpdate *GitHubCostsIntegrationAccountBearerTokenAuthUpdate

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// GitHubCostsIntegrationAccountBearerTokenAuthUpdateAsGitHubCostsIntegrationAccountAuthenticationUpdate is a convenience function that returns GitHubCostsIntegrationAccountBearerTokenAuthUpdate wrapped in GitHubCostsIntegrationAccountAuthenticationUpdate.
func GitHubCostsIntegrationAccountBearerTokenAuthUpdateAsGitHubCostsIntegrationAccountAuthenticationUpdate(v *GitHubCostsIntegrationAccountBearerTokenAuthUpdate) GitHubCostsIntegrationAccountAuthenticationUpdate {
	return GitHubCostsIntegrationAccountAuthenticationUpdate{GitHubCostsIntegrationAccountBearerTokenAuthUpdate: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *GitHubCostsIntegrationAccountAuthenticationUpdate) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into GitHubCostsIntegrationAccountBearerTokenAuthUpdate
	err = datadog.Unmarshal(data, &obj.GitHubCostsIntegrationAccountBearerTokenAuthUpdate)
	if err == nil {
		if obj.GitHubCostsIntegrationAccountBearerTokenAuthUpdate != nil && obj.GitHubCostsIntegrationAccountBearerTokenAuthUpdate.UnparsedObject == nil {
			jsonGitHubCostsIntegrationAccountBearerTokenAuthUpdate, _ := datadog.Marshal(obj.GitHubCostsIntegrationAccountBearerTokenAuthUpdate)
			if string(jsonGitHubCostsIntegrationAccountBearerTokenAuthUpdate) == "{}" { // empty struct
				obj.GitHubCostsIntegrationAccountBearerTokenAuthUpdate = nil
			} else {
				match++
			}
		} else {
			obj.GitHubCostsIntegrationAccountBearerTokenAuthUpdate = nil
		}
	} else {
		obj.GitHubCostsIntegrationAccountBearerTokenAuthUpdate = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.GitHubCostsIntegrationAccountBearerTokenAuthUpdate = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj GitHubCostsIntegrationAccountAuthenticationUpdate) MarshalJSON() ([]byte, error) {
	if obj.GitHubCostsIntegrationAccountBearerTokenAuthUpdate != nil {
		return datadog.Marshal(&obj.GitHubCostsIntegrationAccountBearerTokenAuthUpdate)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *GitHubCostsIntegrationAccountAuthenticationUpdate) GetActualInstance() interface{} {
	if obj.GitHubCostsIntegrationAccountBearerTokenAuthUpdate != nil {
		return obj.GitHubCostsIntegrationAccountBearerTokenAuthUpdate
	}

	// all schemas are nil
	return nil
}
