// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentJiraIssueType Incident Jira issue resource type.
type IncidentJiraIssueType string

// List of IncidentJiraIssueType.
const (
	INCIDENTJIRAISSUETYPE_INCIDENT_JIRA_ISSUES IncidentJiraIssueType = "incident_jira_issues"
)

var allowedIncidentJiraIssueTypeEnumValues = []IncidentJiraIssueType{
	INCIDENTJIRAISSUETYPE_INCIDENT_JIRA_ISSUES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentJiraIssueType) GetAllowedValues() []IncidentJiraIssueType {
	return allowedIncidentJiraIssueTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentJiraIssueType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentJiraIssueType(value)
	return nil
}

// NewIncidentJiraIssueTypeFromValue returns a pointer to a valid IncidentJiraIssueType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentJiraIssueTypeFromValue(v string) (*IncidentJiraIssueType, error) {
	ev := IncidentJiraIssueType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentJiraIssueType: valid values are %v", v, allowedIncidentJiraIssueTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentJiraIssueType) IsValid() bool {
	for _, existing := range allowedIncidentJiraIssueTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentJiraIssueType value.
func (v IncidentJiraIssueType) Ptr() *IncidentJiraIssueType {
	return &v
}
