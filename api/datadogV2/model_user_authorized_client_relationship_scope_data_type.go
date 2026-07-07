// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UserAuthorizedClientRelationshipScopeDataType Scope resource type.
type UserAuthorizedClientRelationshipScopeDataType string

// List of UserAuthorizedClientRelationshipScopeDataType.
const (
	USERAUTHORIZEDCLIENTRELATIONSHIPSCOPEDATATYPE_SCOPES UserAuthorizedClientRelationshipScopeDataType = "scopes"
)

var allowedUserAuthorizedClientRelationshipScopeDataTypeEnumValues = []UserAuthorizedClientRelationshipScopeDataType{
	USERAUTHORIZEDCLIENTRELATIONSHIPSCOPEDATATYPE_SCOPES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *UserAuthorizedClientRelationshipScopeDataType) GetAllowedValues() []UserAuthorizedClientRelationshipScopeDataType {
	return allowedUserAuthorizedClientRelationshipScopeDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *UserAuthorizedClientRelationshipScopeDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = UserAuthorizedClientRelationshipScopeDataType(value)
	return nil
}

// NewUserAuthorizedClientRelationshipScopeDataTypeFromValue returns a pointer to a valid UserAuthorizedClientRelationshipScopeDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewUserAuthorizedClientRelationshipScopeDataTypeFromValue(v string) (*UserAuthorizedClientRelationshipScopeDataType, error) {
	ev := UserAuthorizedClientRelationshipScopeDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for UserAuthorizedClientRelationshipScopeDataType: valid values are %v", v, allowedUserAuthorizedClientRelationshipScopeDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v UserAuthorizedClientRelationshipScopeDataType) IsValid() bool {
	for _, existing := range allowedUserAuthorizedClientRelationshipScopeDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UserAuthorizedClientRelationshipScopeDataType value.
func (v UserAuthorizedClientRelationshipScopeDataType) Ptr() *UserAuthorizedClientRelationshipScopeDataType {
	return &v
}
