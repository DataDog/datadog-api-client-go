// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// JiraIssueDataAttributesRequestMode Mode for creating the Jira issue. Can be "single" or "multiple".
type JiraIssueDataAttributesRequestMode string

// List of JiraIssueDataAttributesRequestMode.
const (
	JIRAISSUEDATAATTRIBUTESREQUESTMODE_SINGLE   JiraIssueDataAttributesRequestMode = "single"
	JIRAISSUEDATAATTRIBUTESREQUESTMODE_MULTIPLE JiraIssueDataAttributesRequestMode = "multiple"
)

var allowedJiraIssueDataAttributesRequestModeEnumValues = []JiraIssueDataAttributesRequestMode{
	JIRAISSUEDATAATTRIBUTESREQUESTMODE_SINGLE,
	JIRAISSUEDATAATTRIBUTESREQUESTMODE_MULTIPLE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *JiraIssueDataAttributesRequestMode) GetAllowedValues() []JiraIssueDataAttributesRequestMode {
	return allowedJiraIssueDataAttributesRequestModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *JiraIssueDataAttributesRequestMode) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = JiraIssueDataAttributesRequestMode(value)
	return nil
}

// NewJiraIssueDataAttributesRequestModeFromValue returns a pointer to a valid JiraIssueDataAttributesRequestMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewJiraIssueDataAttributesRequestModeFromValue(v string) (*JiraIssueDataAttributesRequestMode, error) {
	ev := JiraIssueDataAttributesRequestMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for JiraIssueDataAttributesRequestMode: valid values are %v", v, allowedJiraIssueDataAttributesRequestModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v JiraIssueDataAttributesRequestMode) IsValid() bool {
	for _, existing := range allowedJiraIssueDataAttributesRequestModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to JiraIssueDataAttributesRequestMode value.
func (v JiraIssueDataAttributesRequestMode) Ptr() *JiraIssueDataAttributesRequestMode {
	return &v
}
