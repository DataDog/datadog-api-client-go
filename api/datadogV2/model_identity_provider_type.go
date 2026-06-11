// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IdentityProviderType The resource type for identity providers.
type IdentityProviderType string

// List of IdentityProviderType.
const (
	IDENTITYPROVIDERTYPE_IDENTITY_PROVIDERS IdentityProviderType = "identity_providers"
)

var allowedIdentityProviderTypeEnumValues = []IdentityProviderType{
	IDENTITYPROVIDERTYPE_IDENTITY_PROVIDERS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IdentityProviderType) GetAllowedValues() []IdentityProviderType {
	return allowedIdentityProviderTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IdentityProviderType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IdentityProviderType(value)
	return nil
}

// NewIdentityProviderTypeFromValue returns a pointer to a valid IdentityProviderType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIdentityProviderTypeFromValue(v string) (*IdentityProviderType, error) {
	ev := IdentityProviderType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IdentityProviderType: valid values are %v", v, allowedIdentityProviderTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IdentityProviderType) IsValid() bool {
	for _, existing := range allowedIdentityProviderTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IdentityProviderType value.
func (v IdentityProviderType) Ptr() *IdentityProviderType {
	return &v
}
