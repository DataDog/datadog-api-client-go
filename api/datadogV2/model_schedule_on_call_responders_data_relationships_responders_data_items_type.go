// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScheduleOnCallRespondersDataRelationshipsRespondersDataItemsType Identifies the resource type for a responder group linked to a schedule's on-call responders lookup.
type ScheduleOnCallRespondersDataRelationshipsRespondersDataItemsType string

// List of ScheduleOnCallRespondersDataRelationshipsRespondersDataItemsType.
const (
	SCHEDULEONCALLRESPONDERSDATARELATIONSHIPSRESPONDERSDATAITEMSTYPE_SCHEDULE_ONCALL_RESPONDER ScheduleOnCallRespondersDataRelationshipsRespondersDataItemsType = "schedule_oncall_responder"
)

var allowedScheduleOnCallRespondersDataRelationshipsRespondersDataItemsTypeEnumValues = []ScheduleOnCallRespondersDataRelationshipsRespondersDataItemsType{
	SCHEDULEONCALLRESPONDERSDATARELATIONSHIPSRESPONDERSDATAITEMSTYPE_SCHEDULE_ONCALL_RESPONDER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ScheduleOnCallRespondersDataRelationshipsRespondersDataItemsType) GetAllowedValues() []ScheduleOnCallRespondersDataRelationshipsRespondersDataItemsType {
	return allowedScheduleOnCallRespondersDataRelationshipsRespondersDataItemsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ScheduleOnCallRespondersDataRelationshipsRespondersDataItemsType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ScheduleOnCallRespondersDataRelationshipsRespondersDataItemsType(value)
	return nil
}

// NewScheduleOnCallRespondersDataRelationshipsRespondersDataItemsTypeFromValue returns a pointer to a valid ScheduleOnCallRespondersDataRelationshipsRespondersDataItemsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewScheduleOnCallRespondersDataRelationshipsRespondersDataItemsTypeFromValue(v string) (*ScheduleOnCallRespondersDataRelationshipsRespondersDataItemsType, error) {
	ev := ScheduleOnCallRespondersDataRelationshipsRespondersDataItemsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ScheduleOnCallRespondersDataRelationshipsRespondersDataItemsType: valid values are %v", v, allowedScheduleOnCallRespondersDataRelationshipsRespondersDataItemsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ScheduleOnCallRespondersDataRelationshipsRespondersDataItemsType) IsValid() bool {
	for _, existing := range allowedScheduleOnCallRespondersDataRelationshipsRespondersDataItemsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ScheduleOnCallRespondersDataRelationshipsRespondersDataItemsType value.
func (v ScheduleOnCallRespondersDataRelationshipsRespondersDataItemsType) Ptr() *ScheduleOnCallRespondersDataRelationshipsRespondersDataItemsType {
	return &v
}
