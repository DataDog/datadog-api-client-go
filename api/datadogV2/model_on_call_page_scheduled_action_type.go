// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OnCallPageScheduledActionType Scheduled action resource type.
type OnCallPageScheduledActionType string

// List of OnCallPageScheduledActionType.
const (
	ONCALLPAGESCHEDULEDACTIONTYPE_SCHEDULED_ACTIONS OnCallPageScheduledActionType = "scheduled_actions"
)

var allowedOnCallPageScheduledActionTypeEnumValues = []OnCallPageScheduledActionType{
	ONCALLPAGESCHEDULEDACTIONTYPE_SCHEDULED_ACTIONS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OnCallPageScheduledActionType) GetAllowedValues() []OnCallPageScheduledActionType {
	return allowedOnCallPageScheduledActionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OnCallPageScheduledActionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OnCallPageScheduledActionType(value)
	return nil
}

// NewOnCallPageScheduledActionTypeFromValue returns a pointer to a valid OnCallPageScheduledActionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOnCallPageScheduledActionTypeFromValue(v string) (*OnCallPageScheduledActionType, error) {
	ev := OnCallPageScheduledActionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OnCallPageScheduledActionType: valid values are %v", v, allowedOnCallPageScheduledActionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OnCallPageScheduledActionType) IsValid() bool {
	for _, existing := range allowedOnCallPageScheduledActionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OnCallPageScheduledActionType value.
func (v OnCallPageScheduledActionType) Ptr() *OnCallPageScheduledActionType {
	return &v
}
