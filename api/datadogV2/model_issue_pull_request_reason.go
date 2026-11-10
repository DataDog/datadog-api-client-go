// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IssuePullRequestReason Why the pull request is linked to the issue. `fix` means the pull request fixes the issue, and `trigger` means the pull request introduced the issue.
type IssuePullRequestReason string

// List of IssuePullRequestReason.
const (
	ISSUEPULLREQUESTREASON_UNKNOWN IssuePullRequestReason = "unknown"
	ISSUEPULLREQUESTREASON_FIX     IssuePullRequestReason = "fix"
	ISSUEPULLREQUESTREASON_TRIGGER IssuePullRequestReason = "trigger"
)

var allowedIssuePullRequestReasonEnumValues = []IssuePullRequestReason{
	ISSUEPULLREQUESTREASON_UNKNOWN,
	ISSUEPULLREQUESTREASON_FIX,
	ISSUEPULLREQUESTREASON_TRIGGER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IssuePullRequestReason) GetAllowedValues() []IssuePullRequestReason {
	return allowedIssuePullRequestReasonEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IssuePullRequestReason) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IssuePullRequestReason(value)
	return nil
}

// NewIssuePullRequestReasonFromValue returns a pointer to a valid IssuePullRequestReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIssuePullRequestReasonFromValue(v string) (*IssuePullRequestReason, error) {
	ev := IssuePullRequestReason(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IssuePullRequestReason: valid values are %v", v, allowedIssuePullRequestReasonEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IssuePullRequestReason) IsValid() bool {
	for _, existing := range allowedIssuePullRequestReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IssuePullRequestReason value.
func (v IssuePullRequestReason) Ptr() *IssuePullRequestReason {
	return &v
}
