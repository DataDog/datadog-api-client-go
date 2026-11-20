// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMRecoverySessionStatus Session availability or recovery state. `RECOVERED` indicates that the session is ready to query.
type RUMRecoverySessionStatus string

// List of RUMRecoverySessionStatus.
const (
	RUMRECOVERYSESSIONSTATUS_AVAILABLE  RUMRecoverySessionStatus = "AVAILABLE"
	RUMRECOVERYSESSIONSTATUS_PENDING    RUMRecoverySessionStatus = "PENDING"
	RUMRECOVERYSESSIONSTATUS_PROCESSING RUMRecoverySessionStatus = "PROCESSING"
	RUMRECOVERYSESSIONSTATUS_RECOVERED  RUMRecoverySessionStatus = "RECOVERED"
	RUMRECOVERYSESSIONSTATUS_FAILED     RUMRecoverySessionStatus = "FAILED"
	RUMRECOVERYSESSIONSTATUS_RETAINED   RUMRecoverySessionStatus = "RETAINED"
	RUMRECOVERYSESSIONSTATUS_EXPIRED    RUMRecoverySessionStatus = "EXPIRED"
	RUMRECOVERYSESSIONSTATUS_CANCELLED  RUMRecoverySessionStatus = "CANCELLED"
)

var allowedRUMRecoverySessionStatusEnumValues = []RUMRecoverySessionStatus{
	RUMRECOVERYSESSIONSTATUS_AVAILABLE,
	RUMRECOVERYSESSIONSTATUS_PENDING,
	RUMRECOVERYSESSIONSTATUS_PROCESSING,
	RUMRECOVERYSESSIONSTATUS_RECOVERED,
	RUMRECOVERYSESSIONSTATUS_FAILED,
	RUMRECOVERYSESSIONSTATUS_RETAINED,
	RUMRECOVERYSESSIONSTATUS_EXPIRED,
	RUMRECOVERYSESSIONSTATUS_CANCELLED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RUMRecoverySessionStatus) GetAllowedValues() []RUMRecoverySessionStatus {
	return allowedRUMRecoverySessionStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RUMRecoverySessionStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RUMRecoverySessionStatus(value)
	return nil
}

// NewRUMRecoverySessionStatusFromValue returns a pointer to a valid RUMRecoverySessionStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRUMRecoverySessionStatusFromValue(v string) (*RUMRecoverySessionStatus, error) {
	ev := RUMRecoverySessionStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RUMRecoverySessionStatus: valid values are %v", v, allowedRUMRecoverySessionStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RUMRecoverySessionStatus) IsValid() bool {
	for _, existing := range allowedRUMRecoverySessionStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RUMRecoverySessionStatus value.
func (v RUMRecoverySessionStatus) Ptr() *RUMRecoverySessionStatus {
	return &v
}
