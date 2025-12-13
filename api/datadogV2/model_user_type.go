// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UserType Users resource type.
type UserType string

// List of UserType.
const (
	USERTYPE_USERS UserType = "users"
)

var allowedUserTypeEnumValues = []UserType{
	USERTYPE_USERS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *UserType) GetAllowedValues() []UserType {
	return allowedUserTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *UserType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = UserType(value)
	return nil
}

// NewUserTypeFromValue returns a pointer to a valid UserType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewUserTypeFromValue(v string) (*UserType, error) {
	ev := UserType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for UserType: valid values are %v", v, allowedUserTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v UserType) IsValid() bool {
	for _, existing := range allowedUserTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UserType value.
func (v UserType) Ptr() *UserType {
	return &v
}
