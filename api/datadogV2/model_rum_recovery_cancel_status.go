// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumRecoveryCancelStatus Requested cancellation state.
type RumRecoveryCancelStatus string

// List of RumRecoveryCancelStatus.
const (
	RUMRECOVERYCANCELSTATUS_CANCELLED RumRecoveryCancelStatus = "cancelled"
)

var allowedRumRecoveryCancelStatusEnumValues = []RumRecoveryCancelStatus{
	RUMRECOVERYCANCELSTATUS_CANCELLED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RumRecoveryCancelStatus) GetAllowedValues() []RumRecoveryCancelStatus {
	return allowedRumRecoveryCancelStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RumRecoveryCancelStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RumRecoveryCancelStatus(value)
	return nil
}

// NewRumRecoveryCancelStatusFromValue returns a pointer to a valid RumRecoveryCancelStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRumRecoveryCancelStatusFromValue(v string) (*RumRecoveryCancelStatus, error) {
	ev := RumRecoveryCancelStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RumRecoveryCancelStatus: valid values are %v", v, allowedRumRecoveryCancelStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RumRecoveryCancelStatus) IsValid() bool {
	for _, existing := range allowedRumRecoveryCancelStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RumRecoveryCancelStatus value.
func (v RumRecoveryCancelStatus) Ptr() *RumRecoveryCancelStatus {
	return &v
}
