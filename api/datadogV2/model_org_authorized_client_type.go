// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OrgAuthorizedClientType The resource type for org authorized clients.
type OrgAuthorizedClientType string

// List of OrgAuthorizedClientType.
const (
	ORGAUTHORIZEDCLIENTTYPE_ORG_AUTHORIZED_CLIENTS OrgAuthorizedClientType = "org_authorized_clients"
)

var allowedOrgAuthorizedClientTypeEnumValues = []OrgAuthorizedClientType{
	ORGAUTHORIZEDCLIENTTYPE_ORG_AUTHORIZED_CLIENTS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OrgAuthorizedClientType) GetAllowedValues() []OrgAuthorizedClientType {
	return allowedOrgAuthorizedClientTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OrgAuthorizedClientType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OrgAuthorizedClientType(value)
	return nil
}

// NewOrgAuthorizedClientTypeFromValue returns a pointer to a valid OrgAuthorizedClientType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOrgAuthorizedClientTypeFromValue(v string) (*OrgAuthorizedClientType, error) {
	ev := OrgAuthorizedClientType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OrgAuthorizedClientType: valid values are %v", v, allowedOrgAuthorizedClientTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OrgAuthorizedClientType) IsValid() bool {
	for _, existing := range allowedOrgAuthorizedClientTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OrgAuthorizedClientType value.
func (v OrgAuthorizedClientType) Ptr() *OrgAuthorizedClientType {
	return &v
}
