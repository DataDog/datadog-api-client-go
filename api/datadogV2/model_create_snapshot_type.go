// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateSnapshotType The type identifier for snapshot creation resources.
type CreateSnapshotType string

// List of CreateSnapshotType.
const (
	CREATESNAPSHOTTYPE_CREATE_SNAPSHOT CreateSnapshotType = "create_snapshot"
)

var allowedCreateSnapshotTypeEnumValues = []CreateSnapshotType{
	CREATESNAPSHOTTYPE_CREATE_SNAPSHOT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CreateSnapshotType) GetAllowedValues() []CreateSnapshotType {
	return allowedCreateSnapshotTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CreateSnapshotType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CreateSnapshotType(value)
	return nil
}

// NewCreateSnapshotTypeFromValue returns a pointer to a valid CreateSnapshotType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCreateSnapshotTypeFromValue(v string) (*CreateSnapshotType, error) {
	ev := CreateSnapshotType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CreateSnapshotType: valid values are %v", v, allowedCreateSnapshotTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CreateSnapshotType) IsValid() bool {
	for _, existing := range allowedCreateSnapshotTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CreateSnapshotType value.
func (v CreateSnapshotType) Ptr() *CreateSnapshotType {
	return &v
}
