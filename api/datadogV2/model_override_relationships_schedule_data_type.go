// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OverrideRelationshipsScheduleDataType Indicates that the related resource is of type 'schedules'.
type OverrideRelationshipsScheduleDataType string

// List of OverrideRelationshipsScheduleDataType.
const (
	OVERRIDERELATIONSHIPSSCHEDULEDATATYPE_SCHEDULES OverrideRelationshipsScheduleDataType = "schedules"
)

var allowedOverrideRelationshipsScheduleDataTypeEnumValues = []OverrideRelationshipsScheduleDataType{
	OVERRIDERELATIONSHIPSSCHEDULEDATATYPE_SCHEDULES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OverrideRelationshipsScheduleDataType) GetAllowedValues() []OverrideRelationshipsScheduleDataType {
	return allowedOverrideRelationshipsScheduleDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OverrideRelationshipsScheduleDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OverrideRelationshipsScheduleDataType(value)
	return nil
}

// NewOverrideRelationshipsScheduleDataTypeFromValue returns a pointer to a valid OverrideRelationshipsScheduleDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOverrideRelationshipsScheduleDataTypeFromValue(v string) (*OverrideRelationshipsScheduleDataType, error) {
	ev := OverrideRelationshipsScheduleDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OverrideRelationshipsScheduleDataType: valid values are %v", v, allowedOverrideRelationshipsScheduleDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OverrideRelationshipsScheduleDataType) IsValid() bool {
	for _, existing := range allowedOverrideRelationshipsScheduleDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OverrideRelationshipsScheduleDataType value.
func (v OverrideRelationshipsScheduleDataType) Ptr() *OverrideRelationshipsScheduleDataType {
	return &v
}
