// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentTimelineEntryType Incident timeline entry resource type.
type IncidentTimelineEntryType string

// List of IncidentTimelineEntryType.
const (
	INCIDENTTIMELINEENTRYTYPE_INCIDENT_TIMELINE_CELLS IncidentTimelineEntryType = "incident_timeline_cells"
)

var allowedIncidentTimelineEntryTypeEnumValues = []IncidentTimelineEntryType{
	INCIDENTTIMELINEENTRYTYPE_INCIDENT_TIMELINE_CELLS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentTimelineEntryType) GetAllowedValues() []IncidentTimelineEntryType {
	return allowedIncidentTimelineEntryTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentTimelineEntryType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentTimelineEntryType(value)
	return nil
}

// NewIncidentTimelineEntryTypeFromValue returns a pointer to a valid IncidentTimelineEntryType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentTimelineEntryTypeFromValue(v string) (*IncidentTimelineEntryType, error) {
	ev := IncidentTimelineEntryType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentTimelineEntryType: valid values are %v", v, allowedIncidentTimelineEntryTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentTimelineEntryType) IsValid() bool {
	for _, existing := range allowedIncidentTimelineEntryTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentTimelineEntryType value.
func (v IncidentTimelineEntryType) Ptr() *IncidentTimelineEntryType {
	return &v
}
