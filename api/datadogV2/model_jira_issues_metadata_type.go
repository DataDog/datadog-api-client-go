// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// JiraIssuesMetadataType Jira issues metadata resource type.
type JiraIssuesMetadataType string

// List of JiraIssuesMetadataType.
const (
	JIRAISSUESMETADATATYPE_JIRA_ISSUES JiraIssuesMetadataType = "jira_issues"
)

var allowedJiraIssuesMetadataTypeEnumValues = []JiraIssuesMetadataType{
	JIRAISSUESMETADATATYPE_JIRA_ISSUES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *JiraIssuesMetadataType) GetAllowedValues() []JiraIssuesMetadataType {
	return allowedJiraIssuesMetadataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *JiraIssuesMetadataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = JiraIssuesMetadataType(value)
	return nil
}

// NewJiraIssuesMetadataTypeFromValue returns a pointer to a valid JiraIssuesMetadataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewJiraIssuesMetadataTypeFromValue(v string) (*JiraIssuesMetadataType, error) {
	ev := JiraIssuesMetadataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for JiraIssuesMetadataType: valid values are %v", v, allowedJiraIssuesMetadataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v JiraIssuesMetadataType) IsValid() bool {
	for _, existing := range allowedJiraIssuesMetadataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to JiraIssuesMetadataType value.
func (v JiraIssuesMetadataType) Ptr() *JiraIssuesMetadataType {
	return &v
}
