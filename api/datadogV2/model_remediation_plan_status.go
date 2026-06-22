// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RemediationPlanStatus Plan status.
type RemediationPlanStatus string

// List of RemediationPlanStatus.
const (
	REMEDIATIONPLANSTATUS_PENDING     RemediationPlanStatus = "pending"
	REMEDIATIONPLANSTATUS_IN_PROGRESS RemediationPlanStatus = "in_progress"
	REMEDIATIONPLANSTATUS_COMPLETED   RemediationPlanStatus = "completed"
	REMEDIATIONPLANSTATUS_FAILED      RemediationPlanStatus = "failed"
)

var allowedRemediationPlanStatusEnumValues = []RemediationPlanStatus{
	REMEDIATIONPLANSTATUS_PENDING,
	REMEDIATIONPLANSTATUS_IN_PROGRESS,
	REMEDIATIONPLANSTATUS_COMPLETED,
	REMEDIATIONPLANSTATUS_FAILED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RemediationPlanStatus) GetAllowedValues() []RemediationPlanStatus {
	return allowedRemediationPlanStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RemediationPlanStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RemediationPlanStatus(value)
	return nil
}

// NewRemediationPlanStatusFromValue returns a pointer to a valid RemediationPlanStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRemediationPlanStatusFromValue(v string) (*RemediationPlanStatus, error) {
	ev := RemediationPlanStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RemediationPlanStatus: valid values are %v", v, allowedRemediationPlanStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RemediationPlanStatus) IsValid() bool {
	for _, existing := range allowedRemediationPlanStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RemediationPlanStatus value.
func (v RemediationPlanStatus) Ptr() *RemediationPlanStatus {
	return &v
}
