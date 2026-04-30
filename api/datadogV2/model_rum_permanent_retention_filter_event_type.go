// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumPermanentRetentionFilterEventType The type of RUM events the filter applies to. Read-only.
type RumPermanentRetentionFilterEventType string

// List of RumPermanentRetentionFilterEventType.
const (
	RUMPERMANENTRETENTIONFILTEREVENTTYPE_SESSION   RumPermanentRetentionFilterEventType = "session"
	RUMPERMANENTRETENTIONFILTEREVENTTYPE_VIEW      RumPermanentRetentionFilterEventType = "view"
	RUMPERMANENTRETENTIONFILTEREVENTTYPE_ACTION    RumPermanentRetentionFilterEventType = "action"
	RUMPERMANENTRETENTIONFILTEREVENTTYPE_ERROR     RumPermanentRetentionFilterEventType = "error"
	RUMPERMANENTRETENTIONFILTEREVENTTYPE_RESOURCE  RumPermanentRetentionFilterEventType = "resource"
	RUMPERMANENTRETENTIONFILTEREVENTTYPE_LONG_TASK RumPermanentRetentionFilterEventType = "long_task"
	RUMPERMANENTRETENTIONFILTEREVENTTYPE_VITAL     RumPermanentRetentionFilterEventType = "vital"
	RUMPERMANENTRETENTIONFILTEREVENTTYPE_OPERATION RumPermanentRetentionFilterEventType = "operation"
)

var allowedRumPermanentRetentionFilterEventTypeEnumValues = []RumPermanentRetentionFilterEventType{
	RUMPERMANENTRETENTIONFILTEREVENTTYPE_SESSION,
	RUMPERMANENTRETENTIONFILTEREVENTTYPE_VIEW,
	RUMPERMANENTRETENTIONFILTEREVENTTYPE_ACTION,
	RUMPERMANENTRETENTIONFILTEREVENTTYPE_ERROR,
	RUMPERMANENTRETENTIONFILTEREVENTTYPE_RESOURCE,
	RUMPERMANENTRETENTIONFILTEREVENTTYPE_LONG_TASK,
	RUMPERMANENTRETENTIONFILTEREVENTTYPE_VITAL,
	RUMPERMANENTRETENTIONFILTEREVENTTYPE_OPERATION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RumPermanentRetentionFilterEventType) GetAllowedValues() []RumPermanentRetentionFilterEventType {
	return allowedRumPermanentRetentionFilterEventTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RumPermanentRetentionFilterEventType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RumPermanentRetentionFilterEventType(value)
	return nil
}

// NewRumPermanentRetentionFilterEventTypeFromValue returns a pointer to a valid RumPermanentRetentionFilterEventType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRumPermanentRetentionFilterEventTypeFromValue(v string) (*RumPermanentRetentionFilterEventType, error) {
	ev := RumPermanentRetentionFilterEventType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RumPermanentRetentionFilterEventType: valid values are %v", v, allowedRumPermanentRetentionFilterEventTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RumPermanentRetentionFilterEventType) IsValid() bool {
	for _, existing := range allowedRumPermanentRetentionFilterEventTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RumPermanentRetentionFilterEventType value.
func (v RumPermanentRetentionFilterEventType) Ptr() *RumPermanentRetentionFilterEventType {
	return &v
}
