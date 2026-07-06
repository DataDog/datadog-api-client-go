// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UserOverrideIdentityProviderDataType The resource type for identity providers.
type UserOverrideIdentityProviderDataType string

// List of UserOverrideIdentityProviderDataType.
const (
	USEROVERRIDEIDENTITYPROVIDERDATATYPE_IDENTITY_PROVIDERS UserOverrideIdentityProviderDataType = "identity_providers"
)

var allowedUserOverrideIdentityProviderDataTypeEnumValues = []UserOverrideIdentityProviderDataType{
	USEROVERRIDEIDENTITYPROVIDERDATATYPE_IDENTITY_PROVIDERS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *UserOverrideIdentityProviderDataType) GetAllowedValues() []UserOverrideIdentityProviderDataType {
	return allowedUserOverrideIdentityProviderDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *UserOverrideIdentityProviderDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = UserOverrideIdentityProviderDataType(value)
	return nil
}

// NewUserOverrideIdentityProviderDataTypeFromValue returns a pointer to a valid UserOverrideIdentityProviderDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewUserOverrideIdentityProviderDataTypeFromValue(v string) (*UserOverrideIdentityProviderDataType, error) {
	ev := UserOverrideIdentityProviderDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for UserOverrideIdentityProviderDataType: valid values are %v", v, allowedUserOverrideIdentityProviderDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v UserOverrideIdentityProviderDataType) IsValid() bool {
	for _, existing := range allowedUserOverrideIdentityProviderDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UserOverrideIdentityProviderDataType value.
func (v UserOverrideIdentityProviderDataType) Ptr() *UserOverrideIdentityProviderDataType {
	return &v
}
