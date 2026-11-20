// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMRecoveryDiscoveryStatus Batch discovery state. QUEUED indicates that the request is waiting for processing capacity.
type RUMRecoveryDiscoveryStatus string

// List of RUMRecoveryDiscoveryStatus.
const (
	RUMRECOVERYDISCOVERYSTATUS_QUEUED     RUMRecoveryDiscoveryStatus = "QUEUED"
	RUMRECOVERYDISCOVERYSTATUS_PENDING    RUMRecoveryDiscoveryStatus = "PENDING"
	RUMRECOVERYDISCOVERYSTATUS_SLICING    RUMRecoveryDiscoveryStatus = "SLICING"
	RUMRECOVERYDISCOVERYSTATUS_DISCOVERED RUMRecoveryDiscoveryStatus = "DISCOVERED"
	RUMRECOVERYDISCOVERYSTATUS_PROCESSING RUMRecoveryDiscoveryStatus = "PROCESSING"
	RUMRECOVERYDISCOVERYSTATUS_COMPLETED  RUMRecoveryDiscoveryStatus = "COMPLETED"
	RUMRECOVERYDISCOVERYSTATUS_FAILED     RUMRecoveryDiscoveryStatus = "FAILED"
	RUMRECOVERYDISCOVERYSTATUS_CANCELLED  RUMRecoveryDiscoveryStatus = "CANCELLED"
)

var allowedRUMRecoveryDiscoveryStatusEnumValues = []RUMRecoveryDiscoveryStatus{
	RUMRECOVERYDISCOVERYSTATUS_QUEUED,
	RUMRECOVERYDISCOVERYSTATUS_PENDING,
	RUMRECOVERYDISCOVERYSTATUS_SLICING,
	RUMRECOVERYDISCOVERYSTATUS_DISCOVERED,
	RUMRECOVERYDISCOVERYSTATUS_PROCESSING,
	RUMRECOVERYDISCOVERYSTATUS_COMPLETED,
	RUMRECOVERYDISCOVERYSTATUS_FAILED,
	RUMRECOVERYDISCOVERYSTATUS_CANCELLED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RUMRecoveryDiscoveryStatus) GetAllowedValues() []RUMRecoveryDiscoveryStatus {
	return allowedRUMRecoveryDiscoveryStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RUMRecoveryDiscoveryStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RUMRecoveryDiscoveryStatus(value)
	return nil
}

// NewRUMRecoveryDiscoveryStatusFromValue returns a pointer to a valid RUMRecoveryDiscoveryStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRUMRecoveryDiscoveryStatusFromValue(v string) (*RUMRecoveryDiscoveryStatus, error) {
	ev := RUMRecoveryDiscoveryStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RUMRecoveryDiscoveryStatus: valid values are %v", v, allowedRUMRecoveryDiscoveryStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RUMRecoveryDiscoveryStatus) IsValid() bool {
	for _, existing := range allowedRUMRecoveryDiscoveryStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RUMRecoveryDiscoveryStatus value.
func (v RUMRecoveryDiscoveryStatus) Ptr() *RUMRecoveryDiscoveryStatus {
	return &v
}
