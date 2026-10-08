// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GitHubCloudAuthPersonaMappingType Type identifier for GitHub cloud authentication persona mapping.
type GitHubCloudAuthPersonaMappingType string

// List of GitHubCloudAuthPersonaMappingType.
const (
	GITHUBCLOUDAUTHPERSONAMAPPINGTYPE_GITHUB_OIDC_AUTH_CONFIG GitHubCloudAuthPersonaMappingType = "github_oidc_auth_config"
)

var allowedGitHubCloudAuthPersonaMappingTypeEnumValues = []GitHubCloudAuthPersonaMappingType{
	GITHUBCLOUDAUTHPERSONAMAPPINGTYPE_GITHUB_OIDC_AUTH_CONFIG,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GitHubCloudAuthPersonaMappingType) GetAllowedValues() []GitHubCloudAuthPersonaMappingType {
	return allowedGitHubCloudAuthPersonaMappingTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GitHubCloudAuthPersonaMappingType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GitHubCloudAuthPersonaMappingType(value)
	return nil
}

// NewGitHubCloudAuthPersonaMappingTypeFromValue returns a pointer to a valid GitHubCloudAuthPersonaMappingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGitHubCloudAuthPersonaMappingTypeFromValue(v string) (*GitHubCloudAuthPersonaMappingType, error) {
	ev := GitHubCloudAuthPersonaMappingType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GitHubCloudAuthPersonaMappingType: valid values are %v", v, allowedGitHubCloudAuthPersonaMappingTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GitHubCloudAuthPersonaMappingType) IsValid() bool {
	for _, existing := range allowedGitHubCloudAuthPersonaMappingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GitHubCloudAuthPersonaMappingType value.
func (v GitHubCloudAuthPersonaMappingType) Ptr() *GitHubCloudAuthPersonaMappingType {
	return &v
}
