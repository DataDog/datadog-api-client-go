// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GithubPersonalAccessTokenType The credential type for `GithubPersonalAccessToken`.
type GithubPersonalAccessTokenType string

// List of GithubPersonalAccessTokenType.
const (
	GITHUBPERSONALACCESSTOKENTYPE_GITHUBPERSONALACCESSTOKEN GithubPersonalAccessTokenType = "GithubPersonalAccessToken"
)

var allowedGithubPersonalAccessTokenTypeEnumValues = []GithubPersonalAccessTokenType{
	GITHUBPERSONALACCESSTOKENTYPE_GITHUBPERSONALACCESSTOKEN,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GithubPersonalAccessTokenType) GetAllowedValues() []GithubPersonalAccessTokenType {
	return allowedGithubPersonalAccessTokenTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GithubPersonalAccessTokenType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GithubPersonalAccessTokenType(value)
	return nil
}

// NewGithubPersonalAccessTokenTypeFromValue returns a pointer to a valid GithubPersonalAccessTokenType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGithubPersonalAccessTokenTypeFromValue(v string) (*GithubPersonalAccessTokenType, error) {
	ev := GithubPersonalAccessTokenType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GithubPersonalAccessTokenType: valid values are %v", v, allowedGithubPersonalAccessTokenTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GithubPersonalAccessTokenType) IsValid() bool {
	for _, existing := range allowedGithubPersonalAccessTokenTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GithubPersonalAccessTokenType value.
func (v GithubPersonalAccessTokenType) Ptr() *GithubPersonalAccessTokenType {
	return &v
}
