// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScheduleTriggerOverlapBehavior Controls whether a scheduled workflow run may start while another instance is still running.
type ScheduleTriggerOverlapBehavior string

// List of ScheduleTriggerOverlapBehavior.
const (
	SCHEDULETRIGGEROVERLAPBEHAVIOR_EXCLUSIVE_RUN   ScheduleTriggerOverlapBehavior = "EXCLUSIVE_RUN"
	SCHEDULETRIGGEROVERLAPBEHAVIOR_OVERLAP_ALLOWED ScheduleTriggerOverlapBehavior = "OVERLAP_ALLOWED"
)

var allowedScheduleTriggerOverlapBehaviorEnumValues = []ScheduleTriggerOverlapBehavior{
	SCHEDULETRIGGEROVERLAPBEHAVIOR_EXCLUSIVE_RUN,
	SCHEDULETRIGGEROVERLAPBEHAVIOR_OVERLAP_ALLOWED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ScheduleTriggerOverlapBehavior) GetAllowedValues() []ScheduleTriggerOverlapBehavior {
	return allowedScheduleTriggerOverlapBehaviorEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ScheduleTriggerOverlapBehavior) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ScheduleTriggerOverlapBehavior(value)
	return nil
}

// NewScheduleTriggerOverlapBehaviorFromValue returns a pointer to a valid ScheduleTriggerOverlapBehavior
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewScheduleTriggerOverlapBehaviorFromValue(v string) (*ScheduleTriggerOverlapBehavior, error) {
	ev := ScheduleTriggerOverlapBehavior(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ScheduleTriggerOverlapBehavior: valid values are %v", v, allowedScheduleTriggerOverlapBehaviorEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ScheduleTriggerOverlapBehavior) IsValid() bool {
	for _, existing := range allowedScheduleTriggerOverlapBehaviorEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ScheduleTriggerOverlapBehavior value.
func (v ScheduleTriggerOverlapBehavior) Ptr() *ScheduleTriggerOverlapBehavior {
	return &v
}
