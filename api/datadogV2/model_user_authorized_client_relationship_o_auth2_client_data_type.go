// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UserAuthorizedClientRelationshipOAuth2ClientDataType OAuth2 client resource type.
type UserAuthorizedClientRelationshipOAuth2ClientDataType string

// List of UserAuthorizedClientRelationshipOAuth2ClientDataType.
const (
	USERAUTHORIZEDCLIENTRELATIONSHIPOAUTH2CLIENTDATATYPE_OAUTH2_CLIENTS UserAuthorizedClientRelationshipOAuth2ClientDataType = "oauth2_clients"
)

var allowedUserAuthorizedClientRelationshipOAuth2ClientDataTypeEnumValues = []UserAuthorizedClientRelationshipOAuth2ClientDataType{
	USERAUTHORIZEDCLIENTRELATIONSHIPOAUTH2CLIENTDATATYPE_OAUTH2_CLIENTS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *UserAuthorizedClientRelationshipOAuth2ClientDataType) GetAllowedValues() []UserAuthorizedClientRelationshipOAuth2ClientDataType {
	return allowedUserAuthorizedClientRelationshipOAuth2ClientDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *UserAuthorizedClientRelationshipOAuth2ClientDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = UserAuthorizedClientRelationshipOAuth2ClientDataType(value)
	return nil
}

// NewUserAuthorizedClientRelationshipOAuth2ClientDataTypeFromValue returns a pointer to a valid UserAuthorizedClientRelationshipOAuth2ClientDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewUserAuthorizedClientRelationshipOAuth2ClientDataTypeFromValue(v string) (*UserAuthorizedClientRelationshipOAuth2ClientDataType, error) {
	ev := UserAuthorizedClientRelationshipOAuth2ClientDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for UserAuthorizedClientRelationshipOAuth2ClientDataType: valid values are %v", v, allowedUserAuthorizedClientRelationshipOAuth2ClientDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v UserAuthorizedClientRelationshipOAuth2ClientDataType) IsValid() bool {
	for _, existing := range allowedUserAuthorizedClientRelationshipOAuth2ClientDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UserAuthorizedClientRelationshipOAuth2ClientDataType value.
func (v UserAuthorizedClientRelationshipOAuth2ClientDataType) Ptr() *UserAuthorizedClientRelationshipOAuth2ClientDataType {
	return &v
}
