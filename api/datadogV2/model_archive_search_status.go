// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ArchiveSearchStatus Current state of an Archive Search.
type ArchiveSearchStatus string

// List of ArchiveSearchStatus.
const (
	ARCHIVESEARCHSTATUS_RUNNING       ArchiveSearchStatus = "RUNNING"
	ARCHIVESEARCHSTATUS_COMPLETED     ArchiveSearchStatus = "COMPLETED"
	ARCHIVESEARCHSTATUS_FAILED        ArchiveSearchStatus = "FAILED"
	ARCHIVESEARCHSTATUS_CANCELLED     ArchiveSearchStatus = "CANCELLED"
	ARCHIVESEARCHSTATUS_QUOTA_REACHED ArchiveSearchStatus = "QUOTA_REACHED"
	ARCHIVESEARCHSTATUS_EXPIRED       ArchiveSearchStatus = "EXPIRED"
)

var allowedArchiveSearchStatusEnumValues = []ArchiveSearchStatus{
	ARCHIVESEARCHSTATUS_RUNNING,
	ARCHIVESEARCHSTATUS_COMPLETED,
	ARCHIVESEARCHSTATUS_FAILED,
	ARCHIVESEARCHSTATUS_CANCELLED,
	ARCHIVESEARCHSTATUS_QUOTA_REACHED,
	ARCHIVESEARCHSTATUS_EXPIRED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ArchiveSearchStatus) GetAllowedValues() []ArchiveSearchStatus {
	return allowedArchiveSearchStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ArchiveSearchStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ArchiveSearchStatus(value)
	return nil
}

// NewArchiveSearchStatusFromValue returns a pointer to a valid ArchiveSearchStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewArchiveSearchStatusFromValue(v string) (*ArchiveSearchStatus, error) {
	ev := ArchiveSearchStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ArchiveSearchStatus: valid values are %v", v, allowedArchiveSearchStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ArchiveSearchStatus) IsValid() bool {
	for _, existing := range allowedArchiveSearchStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ArchiveSearchStatus value.
func (v ArchiveSearchStatus) Ptr() *ArchiveSearchStatus {
	return &v
}
