// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScheduleOnCallRespondersDataType Represents the resource type for a schedule's grouped on-call responders across the previous, current, and next positions.
type ScheduleOnCallRespondersDataType string

// List of ScheduleOnCallRespondersDataType.
const (
	SCHEDULEONCALLRESPONDERSDATATYPE_SCHEDULE_ONCALL_RESPONDERS ScheduleOnCallRespondersDataType = "schedule_oncall_responders"
)

var allowedScheduleOnCallRespondersDataTypeEnumValues = []ScheduleOnCallRespondersDataType{
	SCHEDULEONCALLRESPONDERSDATATYPE_SCHEDULE_ONCALL_RESPONDERS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ScheduleOnCallRespondersDataType) GetAllowedValues() []ScheduleOnCallRespondersDataType {
	return allowedScheduleOnCallRespondersDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ScheduleOnCallRespondersDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ScheduleOnCallRespondersDataType(value)
	return nil
}

// NewScheduleOnCallRespondersDataTypeFromValue returns a pointer to a valid ScheduleOnCallRespondersDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewScheduleOnCallRespondersDataTypeFromValue(v string) (*ScheduleOnCallRespondersDataType, error) {
	ev := ScheduleOnCallRespondersDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ScheduleOnCallRespondersDataType: valid values are %v", v, allowedScheduleOnCallRespondersDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ScheduleOnCallRespondersDataType) IsValid() bool {
	for _, existing := range allowedScheduleOnCallRespondersDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ScheduleOnCallRespondersDataType value.
func (v ScheduleOnCallRespondersDataType) Ptr() *ScheduleOnCallRespondersDataType {
	return &v
}
