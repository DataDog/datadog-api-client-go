// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OrgGroupPolicySuggestionStatus The status of the policy suggestion.
type OrgGroupPolicySuggestionStatus string

// List of OrgGroupPolicySuggestionStatus.
const (
	ORGGROUPPOLICYSUGGESTIONSTATUS_PENDING   OrgGroupPolicySuggestionStatus = "pending"
	ORGGROUPPOLICYSUGGESTIONSTATUS_ACCEPTED  OrgGroupPolicySuggestionStatus = "accepted"
	ORGGROUPPOLICYSUGGESTIONSTATUS_DISMISSED OrgGroupPolicySuggestionStatus = "dismissed"
)

var allowedOrgGroupPolicySuggestionStatusEnumValues = []OrgGroupPolicySuggestionStatus{
	ORGGROUPPOLICYSUGGESTIONSTATUS_PENDING,
	ORGGROUPPOLICYSUGGESTIONSTATUS_ACCEPTED,
	ORGGROUPPOLICYSUGGESTIONSTATUS_DISMISSED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OrgGroupPolicySuggestionStatus) GetAllowedValues() []OrgGroupPolicySuggestionStatus {
	return allowedOrgGroupPolicySuggestionStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OrgGroupPolicySuggestionStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OrgGroupPolicySuggestionStatus(value)
	return nil
}

// NewOrgGroupPolicySuggestionStatusFromValue returns a pointer to a valid OrgGroupPolicySuggestionStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOrgGroupPolicySuggestionStatusFromValue(v string) (*OrgGroupPolicySuggestionStatus, error) {
	ev := OrgGroupPolicySuggestionStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OrgGroupPolicySuggestionStatus: valid values are %v", v, allowedOrgGroupPolicySuggestionStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OrgGroupPolicySuggestionStatus) IsValid() bool {
	for _, existing := range allowedOrgGroupPolicySuggestionStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OrgGroupPolicySuggestionStatus value.
func (v OrgGroupPolicySuggestionStatus) Ptr() *OrgGroupPolicySuggestionStatus {
	return &v
}
