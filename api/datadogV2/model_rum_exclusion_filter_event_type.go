// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumExclusionFilterEventType The type of RUM events to filter on.
type RumExclusionFilterEventType string

// List of RumExclusionFilterEventType.
const (
	RUMEXCLUSIONFILTEREVENTTYPE_SESSION   RumExclusionFilterEventType = "session"
	RUMEXCLUSIONFILTEREVENTTYPE_VIEW      RumExclusionFilterEventType = "view"
	RUMEXCLUSIONFILTEREVENTTYPE_ACTION    RumExclusionFilterEventType = "action"
	RUMEXCLUSIONFILTEREVENTTYPE_ERROR     RumExclusionFilterEventType = "error"
	RUMEXCLUSIONFILTEREVENTTYPE_RESOURCE  RumExclusionFilterEventType = "resource"
	RUMEXCLUSIONFILTEREVENTTYPE_LONG_TASK RumExclusionFilterEventType = "long_task"
	RUMEXCLUSIONFILTEREVENTTYPE_VITAL     RumExclusionFilterEventType = "vital"
)

var allowedRumExclusionFilterEventTypeEnumValues = []RumExclusionFilterEventType{
	RUMEXCLUSIONFILTEREVENTTYPE_SESSION,
	RUMEXCLUSIONFILTEREVENTTYPE_VIEW,
	RUMEXCLUSIONFILTEREVENTTYPE_ACTION,
	RUMEXCLUSIONFILTEREVENTTYPE_ERROR,
	RUMEXCLUSIONFILTEREVENTTYPE_RESOURCE,
	RUMEXCLUSIONFILTEREVENTTYPE_LONG_TASK,
	RUMEXCLUSIONFILTEREVENTTYPE_VITAL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RumExclusionFilterEventType) GetAllowedValues() []RumExclusionFilterEventType {
	return allowedRumExclusionFilterEventTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RumExclusionFilterEventType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RumExclusionFilterEventType(value)
	return nil
}

// NewRumExclusionFilterEventTypeFromValue returns a pointer to a valid RumExclusionFilterEventType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRumExclusionFilterEventTypeFromValue(v string) (*RumExclusionFilterEventType, error) {
	ev := RumExclusionFilterEventType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RumExclusionFilterEventType: valid values are %v", v, allowedRumExclusionFilterEventTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RumExclusionFilterEventType) IsValid() bool {
	for _, existing := range allowedRumExclusionFilterEventTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RumExclusionFilterEventType value.
func (v RumExclusionFilterEventType) Ptr() *RumExclusionFilterEventType {
	return &v
}
