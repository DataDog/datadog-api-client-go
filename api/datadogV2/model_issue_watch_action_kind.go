// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IssueWatchActionKind Kind of watch action.
type IssueWatchActionKind string

// List of IssueWatchActionKind.
const (
	ISSUEWATCHACTIONKIND_CHANGE_ISSUE_STATE IssueWatchActionKind = "CHANGE_ISSUE_STATE"
)

var allowedIssueWatchActionKindEnumValues = []IssueWatchActionKind{
	ISSUEWATCHACTIONKIND_CHANGE_ISSUE_STATE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IssueWatchActionKind) GetAllowedValues() []IssueWatchActionKind {
	return allowedIssueWatchActionKindEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IssueWatchActionKind) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IssueWatchActionKind(value)
	return nil
}

// NewIssueWatchActionKindFromValue returns a pointer to a valid IssueWatchActionKind
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIssueWatchActionKindFromValue(v string) (*IssueWatchActionKind, error) {
	ev := IssueWatchActionKind(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IssueWatchActionKind: valid values are %v", v, allowedIssueWatchActionKindEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IssueWatchActionKind) IsValid() bool {
	for _, existing := range allowedIssueWatchActionKindEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IssueWatchActionKind value.
func (v IssueWatchActionKind) Ptr() *IssueWatchActionKind {
	return &v
}
