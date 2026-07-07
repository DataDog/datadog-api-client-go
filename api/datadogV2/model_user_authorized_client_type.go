// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UserAuthorizedClientType The resource type for user authorized clients.
type UserAuthorizedClientType string

// List of UserAuthorizedClientType.
const (
	USERAUTHORIZEDCLIENTTYPE_USER_AUTHORIZED_CLIENTS UserAuthorizedClientType = "user_authorized_clients"
)

var allowedUserAuthorizedClientTypeEnumValues = []UserAuthorizedClientType{
	USERAUTHORIZEDCLIENTTYPE_USER_AUTHORIZED_CLIENTS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *UserAuthorizedClientType) GetAllowedValues() []UserAuthorizedClientType {
	return allowedUserAuthorizedClientTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *UserAuthorizedClientType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = UserAuthorizedClientType(value)
	return nil
}

// NewUserAuthorizedClientTypeFromValue returns a pointer to a valid UserAuthorizedClientType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewUserAuthorizedClientTypeFromValue(v string) (*UserAuthorizedClientType, error) {
	ev := UserAuthorizedClientType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for UserAuthorizedClientType: valid values are %v", v, allowedUserAuthorizedClientTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v UserAuthorizedClientType) IsValid() bool {
	for _, existing := range allowedUserAuthorizedClientTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UserAuthorizedClientType value.
func (v UserAuthorizedClientType) Ptr() *UserAuthorizedClientType {
	return &v
}
