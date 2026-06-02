// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumHardcodedRetentionFilterEventType The type of RUM events the hardcoded filter applies to. Read-only.
type RumHardcodedRetentionFilterEventType string

// List of RumHardcodedRetentionFilterEventType.
const (
	RUMHARDCODEDRETENTIONFILTEREVENTTYPE_SESSION   RumHardcodedRetentionFilterEventType = "session"
	RUMHARDCODEDRETENTIONFILTEREVENTTYPE_VIEW      RumHardcodedRetentionFilterEventType = "view"
	RUMHARDCODEDRETENTIONFILTEREVENTTYPE_ACTION    RumHardcodedRetentionFilterEventType = "action"
	RUMHARDCODEDRETENTIONFILTEREVENTTYPE_ERROR     RumHardcodedRetentionFilterEventType = "error"
	RUMHARDCODEDRETENTIONFILTEREVENTTYPE_RESOURCE  RumHardcodedRetentionFilterEventType = "resource"
	RUMHARDCODEDRETENTIONFILTEREVENTTYPE_LONG_TASK RumHardcodedRetentionFilterEventType = "long_task"
	RUMHARDCODEDRETENTIONFILTEREVENTTYPE_VITAL     RumHardcodedRetentionFilterEventType = "vital"
)

var allowedRumHardcodedRetentionFilterEventTypeEnumValues = []RumHardcodedRetentionFilterEventType{
	RUMHARDCODEDRETENTIONFILTEREVENTTYPE_SESSION,
	RUMHARDCODEDRETENTIONFILTEREVENTTYPE_VIEW,
	RUMHARDCODEDRETENTIONFILTEREVENTTYPE_ACTION,
	RUMHARDCODEDRETENTIONFILTEREVENTTYPE_ERROR,
	RUMHARDCODEDRETENTIONFILTEREVENTTYPE_RESOURCE,
	RUMHARDCODEDRETENTIONFILTEREVENTTYPE_LONG_TASK,
	RUMHARDCODEDRETENTIONFILTEREVENTTYPE_VITAL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RumHardcodedRetentionFilterEventType) GetAllowedValues() []RumHardcodedRetentionFilterEventType {
	return allowedRumHardcodedRetentionFilterEventTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RumHardcodedRetentionFilterEventType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RumHardcodedRetentionFilterEventType(value)
	return nil
}

// NewRumHardcodedRetentionFilterEventTypeFromValue returns a pointer to a valid RumHardcodedRetentionFilterEventType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRumHardcodedRetentionFilterEventTypeFromValue(v string) (*RumHardcodedRetentionFilterEventType, error) {
	ev := RumHardcodedRetentionFilterEventType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RumHardcodedRetentionFilterEventType: valid values are %v", v, allowedRumHardcodedRetentionFilterEventTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RumHardcodedRetentionFilterEventType) IsValid() bool {
	for _, existing := range allowedRumHardcodedRetentionFilterEventTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RumHardcodedRetentionFilterEventType value.
func (v RumHardcodedRetentionFilterEventType) Ptr() *RumHardcodedRetentionFilterEventType {
	return &v
}
