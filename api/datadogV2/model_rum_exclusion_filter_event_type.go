// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMExclusionFilterEventType The type of RUM events to filter on.
type RUMExclusionFilterEventType string

// List of RUMExclusionFilterEventType.
const (
	RUMEXCLUSIONFILTEREVENTTYPE_SESSION   RUMExclusionFilterEventType = "session"
	RUMEXCLUSIONFILTEREVENTTYPE_VIEW      RUMExclusionFilterEventType = "view"
	RUMEXCLUSIONFILTEREVENTTYPE_ACTION    RUMExclusionFilterEventType = "action"
	RUMEXCLUSIONFILTEREVENTTYPE_ERROR     RUMExclusionFilterEventType = "error"
	RUMEXCLUSIONFILTEREVENTTYPE_RESOURCE  RUMExclusionFilterEventType = "resource"
	RUMEXCLUSIONFILTEREVENTTYPE_LONG_TASK RUMExclusionFilterEventType = "long_task"
	RUMEXCLUSIONFILTEREVENTTYPE_VITAL     RUMExclusionFilterEventType = "vital"
)

var allowedRUMExclusionFilterEventTypeEnumValues = []RUMExclusionFilterEventType{
	RUMEXCLUSIONFILTEREVENTTYPE_SESSION,
	RUMEXCLUSIONFILTEREVENTTYPE_VIEW,
	RUMEXCLUSIONFILTEREVENTTYPE_ACTION,
	RUMEXCLUSIONFILTEREVENTTYPE_ERROR,
	RUMEXCLUSIONFILTEREVENTTYPE_RESOURCE,
	RUMEXCLUSIONFILTEREVENTTYPE_LONG_TASK,
	RUMEXCLUSIONFILTEREVENTTYPE_VITAL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RUMExclusionFilterEventType) GetAllowedValues() []RUMExclusionFilterEventType {
	return allowedRUMExclusionFilterEventTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RUMExclusionFilterEventType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RUMExclusionFilterEventType(value)
	return nil
}

// NewRUMExclusionFilterEventTypeFromValue returns a pointer to a valid RUMExclusionFilterEventType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRUMExclusionFilterEventTypeFromValue(v string) (*RUMExclusionFilterEventType, error) {
	ev := RUMExclusionFilterEventType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RUMExclusionFilterEventType: valid values are %v", v, allowedRUMExclusionFilterEventTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RUMExclusionFilterEventType) IsValid() bool {
	for _, existing := range allowedRUMExclusionFilterEventTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RUMExclusionFilterEventType value.
func (v RUMExclusionFilterEventType) Ptr() *RUMExclusionFilterEventType {
	return &v
}
