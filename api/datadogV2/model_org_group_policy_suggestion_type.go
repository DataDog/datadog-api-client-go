// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OrgGroupPolicySuggestionType Org group policy suggestions resource type.
type OrgGroupPolicySuggestionType string

// List of OrgGroupPolicySuggestionType.
const (
	ORGGROUPPOLICYSUGGESTIONTYPE_ORG_GROUP_POLICY_SUGGESTIONS OrgGroupPolicySuggestionType = "org_group_policy_suggestions"
)

var allowedOrgGroupPolicySuggestionTypeEnumValues = []OrgGroupPolicySuggestionType{
	ORGGROUPPOLICYSUGGESTIONTYPE_ORG_GROUP_POLICY_SUGGESTIONS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OrgGroupPolicySuggestionType) GetAllowedValues() []OrgGroupPolicySuggestionType {
	return allowedOrgGroupPolicySuggestionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OrgGroupPolicySuggestionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OrgGroupPolicySuggestionType(value)
	return nil
}

// NewOrgGroupPolicySuggestionTypeFromValue returns a pointer to a valid OrgGroupPolicySuggestionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOrgGroupPolicySuggestionTypeFromValue(v string) (*OrgGroupPolicySuggestionType, error) {
	ev := OrgGroupPolicySuggestionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OrgGroupPolicySuggestionType: valid values are %v", v, allowedOrgGroupPolicySuggestionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OrgGroupPolicySuggestionType) IsValid() bool {
	for _, existing := range allowedOrgGroupPolicySuggestionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OrgGroupPolicySuggestionType value.
func (v OrgGroupPolicySuggestionType) Ptr() *OrgGroupPolicySuggestionType {
	return &v
}
