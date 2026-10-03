// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GitHubCloudAuthIntakeMappingType Type identifier for GitHub cloud authentication intake mapping.
type GitHubCloudAuthIntakeMappingType string

// List of GitHubCloudAuthIntakeMappingType.
const (
	GITHUBCLOUDAUTHINTAKEMAPPINGTYPE_GITHUB_OIDC_AUTH_INTAKE_MAPPING GitHubCloudAuthIntakeMappingType = "github_oidc_auth_intake_mapping"
)

var allowedGitHubCloudAuthIntakeMappingTypeEnumValues = []GitHubCloudAuthIntakeMappingType{
	GITHUBCLOUDAUTHINTAKEMAPPINGTYPE_GITHUB_OIDC_AUTH_INTAKE_MAPPING,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GitHubCloudAuthIntakeMappingType) GetAllowedValues() []GitHubCloudAuthIntakeMappingType {
	return allowedGitHubCloudAuthIntakeMappingTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GitHubCloudAuthIntakeMappingType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GitHubCloudAuthIntakeMappingType(value)
	return nil
}

// NewGitHubCloudAuthIntakeMappingTypeFromValue returns a pointer to a valid GitHubCloudAuthIntakeMappingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGitHubCloudAuthIntakeMappingTypeFromValue(v string) (*GitHubCloudAuthIntakeMappingType, error) {
	ev := GitHubCloudAuthIntakeMappingType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GitHubCloudAuthIntakeMappingType: valid values are %v", v, allowedGitHubCloudAuthIntakeMappingTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GitHubCloudAuthIntakeMappingType) IsValid() bool {
	for _, existing := range allowedGitHubCloudAuthIntakeMappingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GitHubCloudAuthIntakeMappingType value.
func (v GitHubCloudAuthIntakeMappingType) Ptr() *GitHubCloudAuthIntakeMappingType {
	return &v
}
