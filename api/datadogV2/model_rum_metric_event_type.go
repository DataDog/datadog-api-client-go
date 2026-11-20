// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMMetricEventType The type of RUM events the metric is generated from. It is set at creation and cannot be changed.
type RUMMetricEventType string

// List of RUMMetricEventType.
const (
	RUMMETRICEVENTTYPE_SESSION   RUMMetricEventType = "session"
	RUMMETRICEVENTTYPE_VIEW      RUMMetricEventType = "view"
	RUMMETRICEVENTTYPE_ACTION    RUMMetricEventType = "action"
	RUMMETRICEVENTTYPE_ERROR     RUMMetricEventType = "error"
	RUMMETRICEVENTTYPE_RESOURCE  RUMMetricEventType = "resource"
	RUMMETRICEVENTTYPE_LONG_TASK RUMMetricEventType = "long_task"
	RUMMETRICEVENTTYPE_VITAL     RUMMetricEventType = "vital"
)

var allowedRUMMetricEventTypeEnumValues = []RUMMetricEventType{
	RUMMETRICEVENTTYPE_SESSION,
	RUMMETRICEVENTTYPE_VIEW,
	RUMMETRICEVENTTYPE_ACTION,
	RUMMETRICEVENTTYPE_ERROR,
	RUMMETRICEVENTTYPE_RESOURCE,
	RUMMETRICEVENTTYPE_LONG_TASK,
	RUMMETRICEVENTTYPE_VITAL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RUMMetricEventType) GetAllowedValues() []RUMMetricEventType {
	return allowedRUMMetricEventTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RUMMetricEventType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RUMMetricEventType(value)
	return nil
}

// NewRUMMetricEventTypeFromValue returns a pointer to a valid RUMMetricEventType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRUMMetricEventTypeFromValue(v string) (*RUMMetricEventType, error) {
	ev := RUMMetricEventType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RUMMetricEventType: valid values are %v", v, allowedRUMMetricEventTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RUMMetricEventType) IsValid() bool {
	for _, existing := range allowedRUMMetricEventTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RUMMetricEventType value.
func (v RUMMetricEventType) Ptr() *RUMMetricEventType {
	return &v
}
