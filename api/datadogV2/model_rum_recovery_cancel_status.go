// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMRecoveryCancelStatus Requested cancellation state.
type RUMRecoveryCancelStatus string

// List of RUMRecoveryCancelStatus.
const (
	RUMRECOVERYCANCELSTATUS_CANCELLED RUMRecoveryCancelStatus = "cancelled"
)

var allowedRUMRecoveryCancelStatusEnumValues = []RUMRecoveryCancelStatus{
	RUMRECOVERYCANCELSTATUS_CANCELLED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RUMRecoveryCancelStatus) GetAllowedValues() []RUMRecoveryCancelStatus {
	return allowedRUMRecoveryCancelStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RUMRecoveryCancelStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RUMRecoveryCancelStatus(value)
	return nil
}

// NewRUMRecoveryCancelStatusFromValue returns a pointer to a valid RUMRecoveryCancelStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRUMRecoveryCancelStatusFromValue(v string) (*RUMRecoveryCancelStatus, error) {
	ev := RUMRecoveryCancelStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RUMRecoveryCancelStatus: valid values are %v", v, allowedRUMRecoveryCancelStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RUMRecoveryCancelStatus) IsValid() bool {
	for _, existing := range allowedRUMRecoveryCancelStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RUMRecoveryCancelStatus value.
func (v RUMRecoveryCancelStatus) Ptr() *RUMRecoveryCancelStatus {
	return &v
}
