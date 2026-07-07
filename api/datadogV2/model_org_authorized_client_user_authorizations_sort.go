// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OrgAuthorizedClientUserAuthorizationsSort Field to sort user authorizations by.
type OrgAuthorizedClientUserAuthorizationsSort string

// List of OrgAuthorizedClientUserAuthorizationsSort.
const (
	ORGAUTHORIZEDCLIENTUSERAUTHORIZATIONSSORT_USER_NAME          OrgAuthorizedClientUserAuthorizationsSort = "user.name"
	ORGAUTHORIZEDCLIENTUSERAUTHORIZATIONSSORT_USER_EMAIL         OrgAuthorizedClientUserAuthorizationsSort = "user.email"
	ORGAUTHORIZEDCLIENTUSERAUTHORIZATIONSSORT_OAUTH2_CLIENT_NAME OrgAuthorizedClientUserAuthorizationsSort = "oauth2_client.name"
)

var allowedOrgAuthorizedClientUserAuthorizationsSortEnumValues = []OrgAuthorizedClientUserAuthorizationsSort{
	ORGAUTHORIZEDCLIENTUSERAUTHORIZATIONSSORT_USER_NAME,
	ORGAUTHORIZEDCLIENTUSERAUTHORIZATIONSSORT_USER_EMAIL,
	ORGAUTHORIZEDCLIENTUSERAUTHORIZATIONSSORT_OAUTH2_CLIENT_NAME,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OrgAuthorizedClientUserAuthorizationsSort) GetAllowedValues() []OrgAuthorizedClientUserAuthorizationsSort {
	return allowedOrgAuthorizedClientUserAuthorizationsSortEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OrgAuthorizedClientUserAuthorizationsSort) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OrgAuthorizedClientUserAuthorizationsSort(value)
	return nil
}

// NewOrgAuthorizedClientUserAuthorizationsSortFromValue returns a pointer to a valid OrgAuthorizedClientUserAuthorizationsSort
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOrgAuthorizedClientUserAuthorizationsSortFromValue(v string) (*OrgAuthorizedClientUserAuthorizationsSort, error) {
	ev := OrgAuthorizedClientUserAuthorizationsSort(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OrgAuthorizedClientUserAuthorizationsSort: valid values are %v", v, allowedOrgAuthorizedClientUserAuthorizationsSortEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OrgAuthorizedClientUserAuthorizationsSort) IsValid() bool {
	for _, existing := range allowedOrgAuthorizedClientUserAuthorizationsSortEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OrgAuthorizedClientUserAuthorizationsSort value.
func (v OrgAuthorizedClientUserAuthorizationsSort) Ptr() *OrgAuthorizedClientUserAuthorizationsSort {
	return &v
}
