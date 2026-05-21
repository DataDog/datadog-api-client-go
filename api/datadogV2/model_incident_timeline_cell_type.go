// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentTimelineCellType The type of a timeline cell.
type IncidentTimelineCellType string

// List of IncidentTimelineCellType.
const (
	INCIDENTTIMELINECELLTYPE_MARKDOWN               IncidentTimelineCellType = "markdown"
	INCIDENTTIMELINECELLTYPE_INCIDENT_STATUS_CHANGE IncidentTimelineCellType = "incident_status_change"
	INCIDENTTIMELINECELLTYPE_TIMESTAMP_CHANGE       IncidentTimelineCellType = "timestamp_change"
	INCIDENTTIMELINECELLTYPE_MEETING_SUMMARY        IncidentTimelineCellType = "meeting_summary"
	INCIDENTTIMELINECELLTYPE_MEETING_CHAT           IncidentTimelineCellType = "meeting_chat"
	INCIDENTTIMELINECELLTYPE_ROLE_ASSIGNMENT_CHANGE IncidentTimelineCellType = "role_assignment_change"
	INCIDENTTIMELINECELLTYPE_POSTMORTEM_CHANGE      IncidentTimelineCellType = "postmortem_change"
)

var allowedIncidentTimelineCellTypeEnumValues = []IncidentTimelineCellType{
	INCIDENTTIMELINECELLTYPE_MARKDOWN,
	INCIDENTTIMELINECELLTYPE_INCIDENT_STATUS_CHANGE,
	INCIDENTTIMELINECELLTYPE_TIMESTAMP_CHANGE,
	INCIDENTTIMELINECELLTYPE_MEETING_SUMMARY,
	INCIDENTTIMELINECELLTYPE_MEETING_CHAT,
	INCIDENTTIMELINECELLTYPE_ROLE_ASSIGNMENT_CHANGE,
	INCIDENTTIMELINECELLTYPE_POSTMORTEM_CHANGE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentTimelineCellType) GetAllowedValues() []IncidentTimelineCellType {
	return allowedIncidentTimelineCellTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentTimelineCellType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentTimelineCellType(value)
	return nil
}

// NewIncidentTimelineCellTypeFromValue returns a pointer to a valid IncidentTimelineCellType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentTimelineCellTypeFromValue(v string) (*IncidentTimelineCellType, error) {
	ev := IncidentTimelineCellType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentTimelineCellType: valid values are %v", v, allowedIncidentTimelineCellTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentTimelineCellType) IsValid() bool {
	for _, existing := range allowedIncidentTimelineCellTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentTimelineCellType value.
func (v IncidentTimelineCellType) Ptr() *IncidentTimelineCellType {
	return &v
}
