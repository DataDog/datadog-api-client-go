// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GithubIntegrationType The GitHub integration type.
type GithubIntegrationType string

// List of GithubIntegrationType.
const (
	GITHUBINTEGRATIONTYPE_GITHUB GithubIntegrationType = "Github"
)

var allowedGithubIntegrationTypeEnumValues = []GithubIntegrationType{
	GITHUBINTEGRATIONTYPE_GITHUB,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GithubIntegrationType) GetAllowedValues() []GithubIntegrationType {
	return allowedGithubIntegrationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GithubIntegrationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GithubIntegrationType(value)
	return nil
}

// NewGithubIntegrationTypeFromValue returns a pointer to a valid GithubIntegrationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGithubIntegrationTypeFromValue(v string) (*GithubIntegrationType, error) {
	ev := GithubIntegrationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GithubIntegrationType: valid values are %v", v, allowedGithubIntegrationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GithubIntegrationType) IsValid() bool {
	for _, existing := range allowedGithubIntegrationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GithubIntegrationType value.
func (v GithubIntegrationType) Ptr() *GithubIntegrationType {
	return &v
}
