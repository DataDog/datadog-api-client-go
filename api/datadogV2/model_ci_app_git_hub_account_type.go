// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CIAppGitHubAccountType JSON:API type for the GitHub account resource.
// The value must always be `ci_github_account`.
type CIAppGitHubAccountType string

// List of CIAppGitHubAccountType.
const (
	CIAPPGITHUBACCOUNTTYPE_CI_GITHUB_ACCOUNT CIAppGitHubAccountType = "ci_github_account"
)

var allowedCIAppGitHubAccountTypeEnumValues = []CIAppGitHubAccountType{
	CIAPPGITHUBACCOUNTTYPE_CI_GITHUB_ACCOUNT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CIAppGitHubAccountType) GetAllowedValues() []CIAppGitHubAccountType {
	return allowedCIAppGitHubAccountTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CIAppGitHubAccountType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CIAppGitHubAccountType(value)
	return nil
}

// NewCIAppGitHubAccountTypeFromValue returns a pointer to a valid CIAppGitHubAccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCIAppGitHubAccountTypeFromValue(v string) (*CIAppGitHubAccountType, error) {
	ev := CIAppGitHubAccountType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CIAppGitHubAccountType: valid values are %v", v, allowedCIAppGitHubAccountTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CIAppGitHubAccountType) IsValid() bool {
	for _, existing := range allowedCIAppGitHubAccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CIAppGitHubAccountType value.
func (v CIAppGitHubAccountType) Ptr() *CIAppGitHubAccountType {
	return &v
}
