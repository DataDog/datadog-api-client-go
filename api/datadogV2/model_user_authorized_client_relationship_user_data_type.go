// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UserAuthorizedClientRelationshipUserDataType User resource type.
type UserAuthorizedClientRelationshipUserDataType string

// List of UserAuthorizedClientRelationshipUserDataType.
const (
	USERAUTHORIZEDCLIENTRELATIONSHIPUSERDATATYPE_USERS UserAuthorizedClientRelationshipUserDataType = "users"
)

var allowedUserAuthorizedClientRelationshipUserDataTypeEnumValues = []UserAuthorizedClientRelationshipUserDataType{
	USERAUTHORIZEDCLIENTRELATIONSHIPUSERDATATYPE_USERS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *UserAuthorizedClientRelationshipUserDataType) GetAllowedValues() []UserAuthorizedClientRelationshipUserDataType {
	return allowedUserAuthorizedClientRelationshipUserDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *UserAuthorizedClientRelationshipUserDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = UserAuthorizedClientRelationshipUserDataType(value)
	return nil
}

// NewUserAuthorizedClientRelationshipUserDataTypeFromValue returns a pointer to a valid UserAuthorizedClientRelationshipUserDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewUserAuthorizedClientRelationshipUserDataTypeFromValue(v string) (*UserAuthorizedClientRelationshipUserDataType, error) {
	ev := UserAuthorizedClientRelationshipUserDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for UserAuthorizedClientRelationshipUserDataType: valid values are %v", v, allowedUserAuthorizedClientRelationshipUserDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v UserAuthorizedClientRelationshipUserDataType) IsValid() bool {
	for _, existing := range allowedUserAuthorizedClientRelationshipUserDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UserAuthorizedClientRelationshipUserDataType value.
func (v UserAuthorizedClientRelationshipUserDataType) Ptr() *UserAuthorizedClientRelationshipUserDataType {
	return &v
}
