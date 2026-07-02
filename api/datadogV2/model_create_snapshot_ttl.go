// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateSnapshotTTL The time-to-live for the snapshot. This value corresponds to storage lifecycle policies that automatically delete the snapshot after the specified period.
type CreateSnapshotTTL string

// List of CreateSnapshotTTL.
const (
	CREATESNAPSHOTTTL_THIRTY_DAYS CreateSnapshotTTL = "30d"
	CREATESNAPSHOTTTL_SIXTY_DAYS  CreateSnapshotTTL = "60d"
	CREATESNAPSHOTTTL_NINETY_DAYS CreateSnapshotTTL = "90d"
	CREATESNAPSHOTTTL_ONE_YEAR    CreateSnapshotTTL = "1y"
	CREATESNAPSHOTTTL_TWO_YEARS   CreateSnapshotTTL = "2y"
	CREATESNAPSHOTTTL_INFINITE    CreateSnapshotTTL = "inf"
)

var allowedCreateSnapshotTTLEnumValues = []CreateSnapshotTTL{
	CREATESNAPSHOTTTL_THIRTY_DAYS,
	CREATESNAPSHOTTTL_SIXTY_DAYS,
	CREATESNAPSHOTTTL_NINETY_DAYS,
	CREATESNAPSHOTTTL_ONE_YEAR,
	CREATESNAPSHOTTTL_TWO_YEARS,
	CREATESNAPSHOTTTL_INFINITE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CreateSnapshotTTL) GetAllowedValues() []CreateSnapshotTTL {
	return allowedCreateSnapshotTTLEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CreateSnapshotTTL) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CreateSnapshotTTL(value)
	return nil
}

// NewCreateSnapshotTTLFromValue returns a pointer to a valid CreateSnapshotTTL
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCreateSnapshotTTLFromValue(v string) (*CreateSnapshotTTL, error) {
	ev := CreateSnapshotTTL(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CreateSnapshotTTL: valid values are %v", v, allowedCreateSnapshotTTLEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CreateSnapshotTTL) IsValid() bool {
	for _, existing := range allowedCreateSnapshotTTLEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CreateSnapshotTTL value.
func (v CreateSnapshotTTL) Ptr() *CreateSnapshotTTL {
	return &v
}
