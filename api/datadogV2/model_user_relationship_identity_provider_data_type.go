// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UserRelationshipIdentityProviderDataType The resource type for identity providers.
type UserRelationshipIdentityProviderDataType string

// List of UserRelationshipIdentityProviderDataType.
const (
	USERRELATIONSHIPIDENTITYPROVIDERDATATYPE_IDENTITY_PROVIDERS UserRelationshipIdentityProviderDataType = "identity_providers"
)

var allowedUserRelationshipIdentityProviderDataTypeEnumValues = []UserRelationshipIdentityProviderDataType{
	USERRELATIONSHIPIDENTITYPROVIDERDATATYPE_IDENTITY_PROVIDERS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *UserRelationshipIdentityProviderDataType) GetAllowedValues() []UserRelationshipIdentityProviderDataType {
	return allowedUserRelationshipIdentityProviderDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *UserRelationshipIdentityProviderDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = UserRelationshipIdentityProviderDataType(value)
	return nil
}

// NewUserRelationshipIdentityProviderDataTypeFromValue returns a pointer to a valid UserRelationshipIdentityProviderDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewUserRelationshipIdentityProviderDataTypeFromValue(v string) (*UserRelationshipIdentityProviderDataType, error) {
	ev := UserRelationshipIdentityProviderDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for UserRelationshipIdentityProviderDataType: valid values are %v", v, allowedUserRelationshipIdentityProviderDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v UserRelationshipIdentityProviderDataType) IsValid() bool {
	for _, existing := range allowedUserRelationshipIdentityProviderDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UserRelationshipIdentityProviderDataType value.
func (v UserRelationshipIdentityProviderDataType) Ptr() *UserRelationshipIdentityProviderDataType {
	return &v
}
