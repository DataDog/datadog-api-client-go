// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMRecoveryRequestStatus Recovery request state.
type RUMRecoveryRequestStatus string

// List of RUMRecoveryRequestStatus.
const (
	RUMRECOVERYREQUESTSTATUS_PENDING    RUMRecoveryRequestStatus = "PENDING"
	RUMRECOVERYREQUESTSTATUS_PROCESSING RUMRecoveryRequestStatus = "PROCESSING"
	RUMRECOVERYREQUESTSTATUS_COMPLETED  RUMRecoveryRequestStatus = "COMPLETED"
	RUMRECOVERYREQUESTSTATUS_FAILED     RUMRecoveryRequestStatus = "FAILED"
	RUMRECOVERYREQUESTSTATUS_CANCELLED  RUMRecoveryRequestStatus = "CANCELLED"
	RUMRECOVERYREQUESTSTATUS_EXPIRED    RUMRecoveryRequestStatus = "EXPIRED"
)

var allowedRUMRecoveryRequestStatusEnumValues = []RUMRecoveryRequestStatus{
	RUMRECOVERYREQUESTSTATUS_PENDING,
	RUMRECOVERYREQUESTSTATUS_PROCESSING,
	RUMRECOVERYREQUESTSTATUS_COMPLETED,
	RUMRECOVERYREQUESTSTATUS_FAILED,
	RUMRECOVERYREQUESTSTATUS_CANCELLED,
	RUMRECOVERYREQUESTSTATUS_EXPIRED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RUMRecoveryRequestStatus) GetAllowedValues() []RUMRecoveryRequestStatus {
	return allowedRUMRecoveryRequestStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RUMRecoveryRequestStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RUMRecoveryRequestStatus(value)
	return nil
}

// NewRUMRecoveryRequestStatusFromValue returns a pointer to a valid RUMRecoveryRequestStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRUMRecoveryRequestStatusFromValue(v string) (*RUMRecoveryRequestStatus, error) {
	ev := RUMRecoveryRequestStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RUMRecoveryRequestStatus: valid values are %v", v, allowedRUMRecoveryRequestStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RUMRecoveryRequestStatus) IsValid() bool {
	for _, existing := range allowedRUMRecoveryRequestStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RUMRecoveryRequestStatus value.
func (v RUMRecoveryRequestStatus) Ptr() *RUMRecoveryRequestStatus {
	return &v
}
