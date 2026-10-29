// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TemporaryGrantRequestStatus The durable request decision, independent of whether its grant is still active.
type TemporaryGrantRequestStatus string

// List of TemporaryGrantRequestStatus.
const (
	TEMPORARYGRANTREQUESTSTATUS_PENDING   TemporaryGrantRequestStatus = "pending"
	TEMPORARYGRANTREQUESTSTATUS_APPROVED  TemporaryGrantRequestStatus = "approved"
	TEMPORARYGRANTREQUESTSTATUS_DENIED    TemporaryGrantRequestStatus = "denied"
	TEMPORARYGRANTREQUESTSTATUS_CANCELLED TemporaryGrantRequestStatus = "cancelled"
)

var allowedTemporaryGrantRequestStatusEnumValues = []TemporaryGrantRequestStatus{
	TEMPORARYGRANTREQUESTSTATUS_PENDING,
	TEMPORARYGRANTREQUESTSTATUS_APPROVED,
	TEMPORARYGRANTREQUESTSTATUS_DENIED,
	TEMPORARYGRANTREQUESTSTATUS_CANCELLED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TemporaryGrantRequestStatus) GetAllowedValues() []TemporaryGrantRequestStatus {
	return allowedTemporaryGrantRequestStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TemporaryGrantRequestStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TemporaryGrantRequestStatus(value)
	return nil
}

// NewTemporaryGrantRequestStatusFromValue returns a pointer to a valid TemporaryGrantRequestStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTemporaryGrantRequestStatusFromValue(v string) (*TemporaryGrantRequestStatus, error) {
	ev := TemporaryGrantRequestStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TemporaryGrantRequestStatus: valid values are %v", v, allowedTemporaryGrantRequestStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TemporaryGrantRequestStatus) IsValid() bool {
	for _, existing := range allowedTemporaryGrantRequestStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TemporaryGrantRequestStatus value.
func (v TemporaryGrantRequestStatus) Ptr() *TemporaryGrantRequestStatus {
	return &v
}
