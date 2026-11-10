// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IssueWatchConditionKind Kind of watch condition.
type IssueWatchConditionKind string

// List of IssueWatchConditionKind.
const (
	ISSUEWATCHCONDITIONKIND_OCCURRENCES  IssueWatchConditionKind = "OCCURRENCES"
	ISSUEWATCHCONDITIONKIND_CARDINALITY  IssueWatchConditionKind = "CARDINALITY"
	ISSUEWATCHCONDITIONKIND_ELAPSED_TIME IssueWatchConditionKind = "ELAPSED_TIME"
	ISSUEWATCHCONDITIONKIND_ANY          IssueWatchConditionKind = "ANY"
	ISSUEWATCHCONDITIONKIND_ALL          IssueWatchConditionKind = "ALL"
)

var allowedIssueWatchConditionKindEnumValues = []IssueWatchConditionKind{
	ISSUEWATCHCONDITIONKIND_OCCURRENCES,
	ISSUEWATCHCONDITIONKIND_CARDINALITY,
	ISSUEWATCHCONDITIONKIND_ELAPSED_TIME,
	ISSUEWATCHCONDITIONKIND_ANY,
	ISSUEWATCHCONDITIONKIND_ALL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IssueWatchConditionKind) GetAllowedValues() []IssueWatchConditionKind {
	return allowedIssueWatchConditionKindEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IssueWatchConditionKind) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IssueWatchConditionKind(value)
	return nil
}

// NewIssueWatchConditionKindFromValue returns a pointer to a valid IssueWatchConditionKind
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIssueWatchConditionKindFromValue(v string) (*IssueWatchConditionKind, error) {
	ev := IssueWatchConditionKind(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IssueWatchConditionKind: valid values are %v", v, allowedIssueWatchConditionKindEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IssueWatchConditionKind) IsValid() bool {
	for _, existing := range allowedIssueWatchConditionKindEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IssueWatchConditionKind value.
func (v IssueWatchConditionKind) Ptr() *IssueWatchConditionKind {
	return &v
}
