// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OnCallPageJiraIssueType Jira issue resource type.
type OnCallPageJiraIssueType string

// List of OnCallPageJiraIssueType.
const (
	ONCALLPAGEJIRAISSUETYPE_JIRA_ISSUES OnCallPageJiraIssueType = "jira_issues"
)

var allowedOnCallPageJiraIssueTypeEnumValues = []OnCallPageJiraIssueType{
	ONCALLPAGEJIRAISSUETYPE_JIRA_ISSUES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OnCallPageJiraIssueType) GetAllowedValues() []OnCallPageJiraIssueType {
	return allowedOnCallPageJiraIssueTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OnCallPageJiraIssueType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OnCallPageJiraIssueType(value)
	return nil
}

// NewOnCallPageJiraIssueTypeFromValue returns a pointer to a valid OnCallPageJiraIssueType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOnCallPageJiraIssueTypeFromValue(v string) (*OnCallPageJiraIssueType, error) {
	ev := OnCallPageJiraIssueType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OnCallPageJiraIssueType: valid values are %v", v, allowedOnCallPageJiraIssueTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OnCallPageJiraIssueType) IsValid() bool {
	for _, existing := range allowedOnCallPageJiraIssueTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OnCallPageJiraIssueType value.
func (v OnCallPageJiraIssueType) Ptr() *OnCallPageJiraIssueType {
	return &v
}
