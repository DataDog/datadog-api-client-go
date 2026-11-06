// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumRecoverySessionStatus Session availability or recovery state. `RECOVERED` indicates that the session is ready to query.
type RumRecoverySessionStatus string

// List of RumRecoverySessionStatus.
const (
	RUMRECOVERYSESSIONSTATUS_AVAILABLE  RumRecoverySessionStatus = "AVAILABLE"
	RUMRECOVERYSESSIONSTATUS_PENDING    RumRecoverySessionStatus = "PENDING"
	RUMRECOVERYSESSIONSTATUS_PROCESSING RumRecoverySessionStatus = "PROCESSING"
	RUMRECOVERYSESSIONSTATUS_RECOVERED  RumRecoverySessionStatus = "RECOVERED"
	RUMRECOVERYSESSIONSTATUS_FAILED     RumRecoverySessionStatus = "FAILED"
	RUMRECOVERYSESSIONSTATUS_RETAINED   RumRecoverySessionStatus = "RETAINED"
	RUMRECOVERYSESSIONSTATUS_EXPIRED    RumRecoverySessionStatus = "EXPIRED"
	RUMRECOVERYSESSIONSTATUS_CANCELLED  RumRecoverySessionStatus = "CANCELLED"
)

var allowedRumRecoverySessionStatusEnumValues = []RumRecoverySessionStatus{
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
func (v *RumRecoverySessionStatus) GetAllowedValues() []RumRecoverySessionStatus {
	return allowedRumRecoverySessionStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RumRecoverySessionStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RumRecoverySessionStatus(value)
	return nil
}

// NewRumRecoverySessionStatusFromValue returns a pointer to a valid RumRecoverySessionStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRumRecoverySessionStatusFromValue(v string) (*RumRecoverySessionStatus, error) {
	ev := RumRecoverySessionStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RumRecoverySessionStatus: valid values are %v", v, allowedRumRecoverySessionStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RumRecoverySessionStatus) IsValid() bool {
	for _, existing := range allowedRumRecoverySessionStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RumRecoverySessionStatus value.
func (v RumRecoverySessionStatus) Ptr() *RumRecoverySessionStatus {
	return &v
}
