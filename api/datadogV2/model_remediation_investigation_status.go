// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RemediationInvestigationStatus Investigation status.
type RemediationInvestigationStatus string

// List of RemediationInvestigationStatus.
const (
	REMEDIATIONINVESTIGATIONSTATUS_OPEN              RemediationInvestigationStatus = "open"
	REMEDIATIONINVESTIGATIONSTATUS_APPROVAL_REQUIRED RemediationInvestigationStatus = "approval_required"
	REMEDIATIONINVESTIGATIONSTATUS_EXECUTING         RemediationInvestigationStatus = "executing"
	REMEDIATIONINVESTIGATIONSTATUS_SUCCEEDED         RemediationInvestigationStatus = "succeeded"
	REMEDIATIONINVESTIGATIONSTATUS_FAILED            RemediationInvestigationStatus = "failed"
)

var allowedRemediationInvestigationStatusEnumValues = []RemediationInvestigationStatus{
	REMEDIATIONINVESTIGATIONSTATUS_OPEN,
	REMEDIATIONINVESTIGATIONSTATUS_APPROVAL_REQUIRED,
	REMEDIATIONINVESTIGATIONSTATUS_EXECUTING,
	REMEDIATIONINVESTIGATIONSTATUS_SUCCEEDED,
	REMEDIATIONINVESTIGATIONSTATUS_FAILED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RemediationInvestigationStatus) GetAllowedValues() []RemediationInvestigationStatus {
	return allowedRemediationInvestigationStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RemediationInvestigationStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RemediationInvestigationStatus(value)
	return nil
}

// NewRemediationInvestigationStatusFromValue returns a pointer to a valid RemediationInvestigationStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRemediationInvestigationStatusFromValue(v string) (*RemediationInvestigationStatus, error) {
	ev := RemediationInvestigationStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RemediationInvestigationStatus: valid values are %v", v, allowedRemediationInvestigationStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RemediationInvestigationStatus) IsValid() bool {
	for _, existing := range allowedRemediationInvestigationStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RemediationInvestigationStatus value.
func (v RemediationInvestigationStatus) Ptr() *RemediationInvestigationStatus {
	return &v
}
