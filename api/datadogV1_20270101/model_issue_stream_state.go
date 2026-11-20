// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IssueStreamState Test-only generated schema for the dashboard versioning proof of concept.
type IssueStreamState string

// List of IssueStreamState.
const (
	ISSUESTREAMSTATE_OPEN         IssueStreamState = "OPEN"
	ISSUESTREAMSTATE_RESOLVED     IssueStreamState = "RESOLVED"
	ISSUESTREAMSTATE_IGNORED      IssueStreamState = "IGNORED"
	ISSUESTREAMSTATE_EXCLUDED     IssueStreamState = "EXCLUDED"
	ISSUESTREAMSTATE_IN_PROGRESS  IssueStreamState = "IN_PROGRESS"
	ISSUESTREAMSTATE_ACKNOWLEDGED IssueStreamState = "ACKNOWLEDGED"
)

var allowedIssueStreamStateEnumValues = []IssueStreamState{
	ISSUESTREAMSTATE_OPEN,
	ISSUESTREAMSTATE_RESOLVED,
	ISSUESTREAMSTATE_IGNORED,
	ISSUESTREAMSTATE_EXCLUDED,
	ISSUESTREAMSTATE_IN_PROGRESS,
	ISSUESTREAMSTATE_ACKNOWLEDGED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IssueStreamState) GetAllowedValues() []IssueStreamState {
	return allowedIssueStreamStateEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IssueStreamState) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IssueStreamState(value)
	return nil
}

// NewIssueStreamStateFromValue returns a pointer to a valid IssueStreamState
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIssueStreamStateFromValue(v string) (*IssueStreamState, error) {
	ev := IssueStreamState(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IssueStreamState: valid values are %v", v, allowedIssueStreamStateEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IssueStreamState) IsValid() bool {
	for _, existing := range allowedIssueStreamStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IssueStreamState value.
func (v IssueStreamState) Ptr() *IssueStreamState {
	return &v
}
