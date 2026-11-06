// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumRecoveryDiscoveryStatus Batch discovery state. QUEUED indicates that the request is waiting for processing capacity.
type RumRecoveryDiscoveryStatus string

// List of RumRecoveryDiscoveryStatus.
const (
	RUMRECOVERYDISCOVERYSTATUS_QUEUED     RumRecoveryDiscoveryStatus = "QUEUED"
	RUMRECOVERYDISCOVERYSTATUS_PENDING    RumRecoveryDiscoveryStatus = "PENDING"
	RUMRECOVERYDISCOVERYSTATUS_SLICING    RumRecoveryDiscoveryStatus = "SLICING"
	RUMRECOVERYDISCOVERYSTATUS_DISCOVERED RumRecoveryDiscoveryStatus = "DISCOVERED"
	RUMRECOVERYDISCOVERYSTATUS_PROCESSING RumRecoveryDiscoveryStatus = "PROCESSING"
	RUMRECOVERYDISCOVERYSTATUS_COMPLETED  RumRecoveryDiscoveryStatus = "COMPLETED"
	RUMRECOVERYDISCOVERYSTATUS_FAILED     RumRecoveryDiscoveryStatus = "FAILED"
	RUMRECOVERYDISCOVERYSTATUS_CANCELLED  RumRecoveryDiscoveryStatus = "CANCELLED"
)

var allowedRumRecoveryDiscoveryStatusEnumValues = []RumRecoveryDiscoveryStatus{
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
func (v *RumRecoveryDiscoveryStatus) GetAllowedValues() []RumRecoveryDiscoveryStatus {
	return allowedRumRecoveryDiscoveryStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RumRecoveryDiscoveryStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RumRecoveryDiscoveryStatus(value)
	return nil
}

// NewRumRecoveryDiscoveryStatusFromValue returns a pointer to a valid RumRecoveryDiscoveryStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRumRecoveryDiscoveryStatusFromValue(v string) (*RumRecoveryDiscoveryStatus, error) {
	ev := RumRecoveryDiscoveryStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RumRecoveryDiscoveryStatus: valid values are %v", v, allowedRumRecoveryDiscoveryStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RumRecoveryDiscoveryStatus) IsValid() bool {
	for _, existing := range allowedRumRecoveryDiscoveryStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RumRecoveryDiscoveryStatus value.
func (v RumRecoveryDiscoveryStatus) Ptr() *RumRecoveryDiscoveryStatus {
	return &v
}
