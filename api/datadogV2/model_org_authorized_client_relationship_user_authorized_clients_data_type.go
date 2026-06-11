// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OrgAuthorizedClientRelationshipUserAuthorizedClientsDataType User authorized client resource type.
type OrgAuthorizedClientRelationshipUserAuthorizedClientsDataType string

// List of OrgAuthorizedClientRelationshipUserAuthorizedClientsDataType.
const (
	ORGAUTHORIZEDCLIENTRELATIONSHIPUSERAUTHORIZEDCLIENTSDATATYPE_USER_AUTHORIZED_CLIENTS OrgAuthorizedClientRelationshipUserAuthorizedClientsDataType = "user_authorized_clients"
)

var allowedOrgAuthorizedClientRelationshipUserAuthorizedClientsDataTypeEnumValues = []OrgAuthorizedClientRelationshipUserAuthorizedClientsDataType{
	ORGAUTHORIZEDCLIENTRELATIONSHIPUSERAUTHORIZEDCLIENTSDATATYPE_USER_AUTHORIZED_CLIENTS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OrgAuthorizedClientRelationshipUserAuthorizedClientsDataType) GetAllowedValues() []OrgAuthorizedClientRelationshipUserAuthorizedClientsDataType {
	return allowedOrgAuthorizedClientRelationshipUserAuthorizedClientsDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OrgAuthorizedClientRelationshipUserAuthorizedClientsDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OrgAuthorizedClientRelationshipUserAuthorizedClientsDataType(value)
	return nil
}

// NewOrgAuthorizedClientRelationshipUserAuthorizedClientsDataTypeFromValue returns a pointer to a valid OrgAuthorizedClientRelationshipUserAuthorizedClientsDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOrgAuthorizedClientRelationshipUserAuthorizedClientsDataTypeFromValue(v string) (*OrgAuthorizedClientRelationshipUserAuthorizedClientsDataType, error) {
	ev := OrgAuthorizedClientRelationshipUserAuthorizedClientsDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OrgAuthorizedClientRelationshipUserAuthorizedClientsDataType: valid values are %v", v, allowedOrgAuthorizedClientRelationshipUserAuthorizedClientsDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OrgAuthorizedClientRelationshipUserAuthorizedClientsDataType) IsValid() bool {
	for _, existing := range allowedOrgAuthorizedClientRelationshipUserAuthorizedClientsDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OrgAuthorizedClientRelationshipUserAuthorizedClientsDataType value.
func (v OrgAuthorizedClientRelationshipUserAuthorizedClientsDataType) Ptr() *OrgAuthorizedClientRelationshipUserAuthorizedClientsDataType {
	return &v
}
