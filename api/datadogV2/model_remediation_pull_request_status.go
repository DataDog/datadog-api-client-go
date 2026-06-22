// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RemediationPullRequestStatus Pull request status for a linked code session.
type RemediationPullRequestStatus string

// List of RemediationPullRequestStatus.
const (
	REMEDIATIONPULLREQUESTSTATUS_OPEN   RemediationPullRequestStatus = "open"
	REMEDIATIONPULLREQUESTSTATUS_CLOSED RemediationPullRequestStatus = "closed"
	REMEDIATIONPULLREQUESTSTATUS_MERGED RemediationPullRequestStatus = "merged"
)

var allowedRemediationPullRequestStatusEnumValues = []RemediationPullRequestStatus{
	REMEDIATIONPULLREQUESTSTATUS_OPEN,
	REMEDIATIONPULLREQUESTSTATUS_CLOSED,
	REMEDIATIONPULLREQUESTSTATUS_MERGED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RemediationPullRequestStatus) GetAllowedValues() []RemediationPullRequestStatus {
	return allowedRemediationPullRequestStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RemediationPullRequestStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RemediationPullRequestStatus(value)
	return nil
}

// NewRemediationPullRequestStatusFromValue returns a pointer to a valid RemediationPullRequestStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRemediationPullRequestStatusFromValue(v string) (*RemediationPullRequestStatus, error) {
	ev := RemediationPullRequestStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RemediationPullRequestStatus: valid values are %v", v, allowedRemediationPullRequestStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RemediationPullRequestStatus) IsValid() bool {
	for _, existing := range allowedRemediationPullRequestStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RemediationPullRequestStatus value.
func (v RemediationPullRequestStatus) Ptr() *RemediationPullRequestStatus {
	return &v
}
