// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamsOwnershipMatchType How the `view_name` is matched against RUM view names.
type TeamsOwnershipMatchType string

// List of TeamsOwnershipMatchType.
const (
	TEAMSOWNERSHIPMATCHTYPE_EXACT  TeamsOwnershipMatchType = "exact"
	TEAMSOWNERSHIPMATCHTYPE_PREFIX TeamsOwnershipMatchType = "prefix"
)

var allowedTeamsOwnershipMatchTypeEnumValues = []TeamsOwnershipMatchType{
	TEAMSOWNERSHIPMATCHTYPE_EXACT,
	TEAMSOWNERSHIPMATCHTYPE_PREFIX,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TeamsOwnershipMatchType) GetAllowedValues() []TeamsOwnershipMatchType {
	return allowedTeamsOwnershipMatchTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TeamsOwnershipMatchType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TeamsOwnershipMatchType(value)
	return nil
}

// NewTeamsOwnershipMatchTypeFromValue returns a pointer to a valid TeamsOwnershipMatchType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTeamsOwnershipMatchTypeFromValue(v string) (*TeamsOwnershipMatchType, error) {
	ev := TeamsOwnershipMatchType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TeamsOwnershipMatchType: valid values are %v", v, allowedTeamsOwnershipMatchTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TeamsOwnershipMatchType) IsValid() bool {
	for _, existing := range allowedTeamsOwnershipMatchTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TeamsOwnershipMatchType value.
func (v TeamsOwnershipMatchType) Ptr() *TeamsOwnershipMatchType {
	return &v
}
