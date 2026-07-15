// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScheduleOnCallResponderDataRelationshipsShiftsDataItemsType Indicates that the related resource is of type `shifts`.
type ScheduleOnCallResponderDataRelationshipsShiftsDataItemsType string

// List of ScheduleOnCallResponderDataRelationshipsShiftsDataItemsType.
const (
	SCHEDULEONCALLRESPONDERDATARELATIONSHIPSSHIFTSDATAITEMSTYPE_SHIFTS ScheduleOnCallResponderDataRelationshipsShiftsDataItemsType = "shifts"
)

var allowedScheduleOnCallResponderDataRelationshipsShiftsDataItemsTypeEnumValues = []ScheduleOnCallResponderDataRelationshipsShiftsDataItemsType{
	SCHEDULEONCALLRESPONDERDATARELATIONSHIPSSHIFTSDATAITEMSTYPE_SHIFTS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ScheduleOnCallResponderDataRelationshipsShiftsDataItemsType) GetAllowedValues() []ScheduleOnCallResponderDataRelationshipsShiftsDataItemsType {
	return allowedScheduleOnCallResponderDataRelationshipsShiftsDataItemsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ScheduleOnCallResponderDataRelationshipsShiftsDataItemsType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ScheduleOnCallResponderDataRelationshipsShiftsDataItemsType(value)
	return nil
}

// NewScheduleOnCallResponderDataRelationshipsShiftsDataItemsTypeFromValue returns a pointer to a valid ScheduleOnCallResponderDataRelationshipsShiftsDataItemsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewScheduleOnCallResponderDataRelationshipsShiftsDataItemsTypeFromValue(v string) (*ScheduleOnCallResponderDataRelationshipsShiftsDataItemsType, error) {
	ev := ScheduleOnCallResponderDataRelationshipsShiftsDataItemsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ScheduleOnCallResponderDataRelationshipsShiftsDataItemsType: valid values are %v", v, allowedScheduleOnCallResponderDataRelationshipsShiftsDataItemsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ScheduleOnCallResponderDataRelationshipsShiftsDataItemsType) IsValid() bool {
	for _, existing := range allowedScheduleOnCallResponderDataRelationshipsShiftsDataItemsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ScheduleOnCallResponderDataRelationshipsShiftsDataItemsType value.
func (v ScheduleOnCallResponderDataRelationshipsShiftsDataItemsType) Ptr() *ScheduleOnCallResponderDataRelationshipsShiftsDataItemsType {
	return &v
}
