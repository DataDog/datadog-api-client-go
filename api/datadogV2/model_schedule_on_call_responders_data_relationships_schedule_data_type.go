// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScheduleOnCallRespondersDataRelationshipsScheduleDataType Identifies the resource type for the schedule associated with this on-call responders lookup.
type ScheduleOnCallRespondersDataRelationshipsScheduleDataType string

// List of ScheduleOnCallRespondersDataRelationshipsScheduleDataType.
const (
	SCHEDULEONCALLRESPONDERSDATARELATIONSHIPSSCHEDULEDATATYPE_SCHEDULES ScheduleOnCallRespondersDataRelationshipsScheduleDataType = "schedules"
)

var allowedScheduleOnCallRespondersDataRelationshipsScheduleDataTypeEnumValues = []ScheduleOnCallRespondersDataRelationshipsScheduleDataType{
	SCHEDULEONCALLRESPONDERSDATARELATIONSHIPSSCHEDULEDATATYPE_SCHEDULES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ScheduleOnCallRespondersDataRelationshipsScheduleDataType) GetAllowedValues() []ScheduleOnCallRespondersDataRelationshipsScheduleDataType {
	return allowedScheduleOnCallRespondersDataRelationshipsScheduleDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ScheduleOnCallRespondersDataRelationshipsScheduleDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ScheduleOnCallRespondersDataRelationshipsScheduleDataType(value)
	return nil
}

// NewScheduleOnCallRespondersDataRelationshipsScheduleDataTypeFromValue returns a pointer to a valid ScheduleOnCallRespondersDataRelationshipsScheduleDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewScheduleOnCallRespondersDataRelationshipsScheduleDataTypeFromValue(v string) (*ScheduleOnCallRespondersDataRelationshipsScheduleDataType, error) {
	ev := ScheduleOnCallRespondersDataRelationshipsScheduleDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ScheduleOnCallRespondersDataRelationshipsScheduleDataType: valid values are %v", v, allowedScheduleOnCallRespondersDataRelationshipsScheduleDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ScheduleOnCallRespondersDataRelationshipsScheduleDataType) IsValid() bool {
	for _, existing := range allowedScheduleOnCallRespondersDataRelationshipsScheduleDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ScheduleOnCallRespondersDataRelationshipsScheduleDataType value.
func (v ScheduleOnCallRespondersDataRelationshipsScheduleDataType) Ptr() *ScheduleOnCallRespondersDataRelationshipsScheduleDataType {
	return &v
}
