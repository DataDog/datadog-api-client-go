// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumRecoveryRequestStatus Recovery request state.
type RumRecoveryRequestStatus string

// List of RumRecoveryRequestStatus.
const (
	RUMRECOVERYREQUESTSTATUS_PENDING    RumRecoveryRequestStatus = "PENDING"
	RUMRECOVERYREQUESTSTATUS_PROCESSING RumRecoveryRequestStatus = "PROCESSING"
	RUMRECOVERYREQUESTSTATUS_COMPLETED  RumRecoveryRequestStatus = "COMPLETED"
	RUMRECOVERYREQUESTSTATUS_FAILED     RumRecoveryRequestStatus = "FAILED"
	RUMRECOVERYREQUESTSTATUS_CANCELLED  RumRecoveryRequestStatus = "CANCELLED"
	RUMRECOVERYREQUESTSTATUS_EXPIRED    RumRecoveryRequestStatus = "EXPIRED"
)

var allowedRumRecoveryRequestStatusEnumValues = []RumRecoveryRequestStatus{
	RUMRECOVERYREQUESTSTATUS_PENDING,
	RUMRECOVERYREQUESTSTATUS_PROCESSING,
	RUMRECOVERYREQUESTSTATUS_COMPLETED,
	RUMRECOVERYREQUESTSTATUS_FAILED,
	RUMRECOVERYREQUESTSTATUS_CANCELLED,
	RUMRECOVERYREQUESTSTATUS_EXPIRED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RumRecoveryRequestStatus) GetAllowedValues() []RumRecoveryRequestStatus {
	return allowedRumRecoveryRequestStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RumRecoveryRequestStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RumRecoveryRequestStatus(value)
	return nil
}

// NewRumRecoveryRequestStatusFromValue returns a pointer to a valid RumRecoveryRequestStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRumRecoveryRequestStatusFromValue(v string) (*RumRecoveryRequestStatus, error) {
	ev := RumRecoveryRequestStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RumRecoveryRequestStatus: valid values are %v", v, allowedRumRecoveryRequestStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RumRecoveryRequestStatus) IsValid() bool {
	for _, existing := range allowedRumRecoveryRequestStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RumRecoveryRequestStatus value.
func (v RumRecoveryRequestStatus) Ptr() *RumRecoveryRequestStatus {
	return &v
}
