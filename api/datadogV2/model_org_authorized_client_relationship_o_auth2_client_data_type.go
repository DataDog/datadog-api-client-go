// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OrgAuthorizedClientRelationshipOAuth2ClientDataType OAuth2 client resource type.
type OrgAuthorizedClientRelationshipOAuth2ClientDataType string

// List of OrgAuthorizedClientRelationshipOAuth2ClientDataType.
const (
	ORGAUTHORIZEDCLIENTRELATIONSHIPOAUTH2CLIENTDATATYPE_OAUTH2_CLIENTS OrgAuthorizedClientRelationshipOAuth2ClientDataType = "oauth2_clients"
)

var allowedOrgAuthorizedClientRelationshipOAuth2ClientDataTypeEnumValues = []OrgAuthorizedClientRelationshipOAuth2ClientDataType{
	ORGAUTHORIZEDCLIENTRELATIONSHIPOAUTH2CLIENTDATATYPE_OAUTH2_CLIENTS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OrgAuthorizedClientRelationshipOAuth2ClientDataType) GetAllowedValues() []OrgAuthorizedClientRelationshipOAuth2ClientDataType {
	return allowedOrgAuthorizedClientRelationshipOAuth2ClientDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OrgAuthorizedClientRelationshipOAuth2ClientDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OrgAuthorizedClientRelationshipOAuth2ClientDataType(value)
	return nil
}

// NewOrgAuthorizedClientRelationshipOAuth2ClientDataTypeFromValue returns a pointer to a valid OrgAuthorizedClientRelationshipOAuth2ClientDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOrgAuthorizedClientRelationshipOAuth2ClientDataTypeFromValue(v string) (*OrgAuthorizedClientRelationshipOAuth2ClientDataType, error) {
	ev := OrgAuthorizedClientRelationshipOAuth2ClientDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OrgAuthorizedClientRelationshipOAuth2ClientDataType: valid values are %v", v, allowedOrgAuthorizedClientRelationshipOAuth2ClientDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OrgAuthorizedClientRelationshipOAuth2ClientDataType) IsValid() bool {
	for _, existing := range allowedOrgAuthorizedClientRelationshipOAuth2ClientDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OrgAuthorizedClientRelationshipOAuth2ClientDataType value.
func (v OrgAuthorizedClientRelationshipOAuth2ClientDataType) Ptr() *OrgAuthorizedClientRelationshipOAuth2ClientDataType {
	return &v
}
