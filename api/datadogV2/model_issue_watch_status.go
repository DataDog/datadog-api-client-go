// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IssueWatchStatus Status of the watch.
type IssueWatchStatus string

// List of IssueWatchStatus.
const (
	ISSUEWATCHSTATUS_ACTIVE IssueWatchStatus = "ACTIVE"
)

var allowedIssueWatchStatusEnumValues = []IssueWatchStatus{
	ISSUEWATCHSTATUS_ACTIVE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IssueWatchStatus) GetAllowedValues() []IssueWatchStatus {
	return allowedIssueWatchStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IssueWatchStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IssueWatchStatus(value)
	return nil
}

// NewIssueWatchStatusFromValue returns a pointer to a valid IssueWatchStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIssueWatchStatusFromValue(v string) (*IssueWatchStatus, error) {
	ev := IssueWatchStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IssueWatchStatus: valid values are %v", v, allowedIssueWatchStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IssueWatchStatus) IsValid() bool {
	for _, existing := range allowedIssueWatchStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IssueWatchStatus value.
func (v IssueWatchStatus) Ptr() *IssueWatchStatus {
	return &v
}
