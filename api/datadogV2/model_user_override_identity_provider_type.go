// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UserOverrideIdentityProviderType The definition of `UserOverrideIdentityProviderType` object.
type UserOverrideIdentityProviderType string

// List of UserOverrideIdentityProviderType.
const (
	USEROVERRIDEIDENTITYPROVIDERTYPE_IDENTITY_PROVIDERS UserOverrideIdentityProviderType = "identity_providers"
)

var allowedUserOverrideIdentityProviderTypeEnumValues = []UserOverrideIdentityProviderType{
	USEROVERRIDEIDENTITYPROVIDERTYPE_IDENTITY_PROVIDERS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *UserOverrideIdentityProviderType) GetAllowedValues() []UserOverrideIdentityProviderType {
	return allowedUserOverrideIdentityProviderTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *UserOverrideIdentityProviderType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = UserOverrideIdentityProviderType(value)
	return nil
}

// NewUserOverrideIdentityProviderTypeFromValue returns a pointer to a valid UserOverrideIdentityProviderType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewUserOverrideIdentityProviderTypeFromValue(v string) (*UserOverrideIdentityProviderType, error) {
	ev := UserOverrideIdentityProviderType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for UserOverrideIdentityProviderType: valid values are %v", v, allowedUserOverrideIdentityProviderTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v UserOverrideIdentityProviderType) IsValid() bool {
	for _, existing := range allowedUserOverrideIdentityProviderTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UserOverrideIdentityProviderType value.
func (v UserOverrideIdentityProviderType) Ptr() *UserOverrideIdentityProviderType {
	return &v
}
