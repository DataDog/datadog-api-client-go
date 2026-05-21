// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentTimelineCellSource The source of a timeline cell.
type IncidentTimelineCellSource string

// List of IncidentTimelineCellSource.
const (
	INCIDENTTIMELINECELLSOURCE_SLACK           IncidentTimelineCellSource = "slack"
	INCIDENTTIMELINECELLSOURCE_MICROSOFT_TEAMS IncidentTimelineCellSource = "microsoft_teams"
	INCIDENTTIMELINECELLSOURCE_DATADOG         IncidentTimelineCellSource = "datadog"
	INCIDENTTIMELINECELLSOURCE_API             IncidentTimelineCellSource = "api"
)

var allowedIncidentTimelineCellSourceEnumValues = []IncidentTimelineCellSource{
	INCIDENTTIMELINECELLSOURCE_SLACK,
	INCIDENTTIMELINECELLSOURCE_MICROSOFT_TEAMS,
	INCIDENTTIMELINECELLSOURCE_DATADOG,
	INCIDENTTIMELINECELLSOURCE_API,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentTimelineCellSource) GetAllowedValues() []IncidentTimelineCellSource {
	return allowedIncidentTimelineCellSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentTimelineCellSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentTimelineCellSource(value)
	return nil
}

// NewIncidentTimelineCellSourceFromValue returns a pointer to a valid IncidentTimelineCellSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentTimelineCellSourceFromValue(v string) (*IncidentTimelineCellSource, error) {
	ev := IncidentTimelineCellSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentTimelineCellSource: valid values are %v", v, allowedIncidentTimelineCellSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentTimelineCellSource) IsValid() bool {
	for _, existing := range allowedIncidentTimelineCellSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentTimelineCellSource value.
func (v IncidentTimelineCellSource) Ptr() *IncidentTimelineCellSource {
	return &v
}
