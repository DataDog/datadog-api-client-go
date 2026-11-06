// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GithubAppType The credential type for `GithubApp`.
type GithubAppType string

// List of GithubAppType.
const (
	GITHUBAPPTYPE_GITHUBAPP GithubAppType = "GithubApp"
)

var allowedGithubAppTypeEnumValues = []GithubAppType{
	GITHUBAPPTYPE_GITHUBAPP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GithubAppType) GetAllowedValues() []GithubAppType {
	return allowedGithubAppTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GithubAppType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GithubAppType(value)
	return nil
}

// NewGithubAppTypeFromValue returns a pointer to a valid GithubAppType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGithubAppTypeFromValue(v string) (*GithubAppType, error) {
	ev := GithubAppType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GithubAppType: valid values are %v", v, allowedGithubAppTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GithubAppType) IsValid() bool {
	for _, existing := range allowedGithubAppTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GithubAppType value.
func (v GithubAppType) Ptr() *GithubAppType {
	return &v
}
