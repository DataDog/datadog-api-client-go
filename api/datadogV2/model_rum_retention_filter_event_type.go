// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMRetentionFilterEventType The type of RUM events to filter on.
type RUMRetentionFilterEventType string

// List of RUMRetentionFilterEventType.
const (
	RUMRETENTIONFILTEREVENTTYPE_SESSION   RUMRetentionFilterEventType = "session"
	RUMRETENTIONFILTEREVENTTYPE_VIEW      RUMRetentionFilterEventType = "view"
	RUMRETENTIONFILTEREVENTTYPE_ACTION    RUMRetentionFilterEventType = "action"
	RUMRETENTIONFILTEREVENTTYPE_ERROR     RUMRetentionFilterEventType = "error"
	RUMRETENTIONFILTEREVENTTYPE_RESOURCE  RUMRetentionFilterEventType = "resource"
	RUMRETENTIONFILTEREVENTTYPE_LONG_TASK RUMRetentionFilterEventType = "long_task"
	RUMRETENTIONFILTEREVENTTYPE_VITAL     RUMRetentionFilterEventType = "vital"
)

var allowedRUMRetentionFilterEventTypeEnumValues = []RUMRetentionFilterEventType{
	RUMRETENTIONFILTEREVENTTYPE_SESSION,
	RUMRETENTIONFILTEREVENTTYPE_VIEW,
	RUMRETENTIONFILTEREVENTTYPE_ACTION,
	RUMRETENTIONFILTEREVENTTYPE_ERROR,
	RUMRETENTIONFILTEREVENTTYPE_RESOURCE,
	RUMRETENTIONFILTEREVENTTYPE_LONG_TASK,
	RUMRETENTIONFILTEREVENTTYPE_VITAL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RUMRetentionFilterEventType) GetAllowedValues() []RUMRetentionFilterEventType {
	return allowedRUMRetentionFilterEventTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RUMRetentionFilterEventType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RUMRetentionFilterEventType(value)
	return nil
}

// NewRUMRetentionFilterEventTypeFromValue returns a pointer to a valid RUMRetentionFilterEventType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRUMRetentionFilterEventTypeFromValue(v string) (*RUMRetentionFilterEventType, error) {
	ev := RUMRetentionFilterEventType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RUMRetentionFilterEventType: valid values are %v", v, allowedRUMRetentionFilterEventTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RUMRetentionFilterEventType) IsValid() bool {
	for _, existing := range allowedRUMRetentionFilterEventTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RUMRetentionFilterEventType value.
func (v RUMRetentionFilterEventType) Ptr() *RUMRetentionFilterEventType {
	return &v
}
