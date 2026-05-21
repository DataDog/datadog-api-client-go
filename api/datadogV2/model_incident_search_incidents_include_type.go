// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentSearchIncidentsIncludeType Types of related objects to include.
type IncidentSearchIncidentsIncludeType string

// List of IncidentSearchIncidentsIncludeType.
const (
	INCIDENTSEARCHINCIDENTSINCLUDETYPE_INCIDENT_TYPE IncidentSearchIncidentsIncludeType = "incident_type"
	INCIDENTSEARCHINCIDENTSINCLUDETYPE_IMPACTS       IncidentSearchIncidentsIncludeType = "impacts"
	INCIDENTSEARCHINCIDENTSINCLUDETYPE_USERS         IncidentSearchIncidentsIncludeType = "users"
	INCIDENTSEARCHINCIDENTSINCLUDETYPE_RESPONDERS    IncidentSearchIncidentsIncludeType = "responders"
	INCIDENTSEARCHINCIDENTSINCLUDETYPE_INTEGRATIONS  IncidentSearchIncidentsIncludeType = "integrations"
	INCIDENTSEARCHINCIDENTSINCLUDETYPE_ATTACHMENTS   IncidentSearchIncidentsIncludeType = "attachments"
)

var allowedIncidentSearchIncidentsIncludeTypeEnumValues = []IncidentSearchIncidentsIncludeType{
	INCIDENTSEARCHINCIDENTSINCLUDETYPE_INCIDENT_TYPE,
	INCIDENTSEARCHINCIDENTSINCLUDETYPE_IMPACTS,
	INCIDENTSEARCHINCIDENTSINCLUDETYPE_USERS,
	INCIDENTSEARCHINCIDENTSINCLUDETYPE_RESPONDERS,
	INCIDENTSEARCHINCIDENTSINCLUDETYPE_INTEGRATIONS,
	INCIDENTSEARCHINCIDENTSINCLUDETYPE_ATTACHMENTS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentSearchIncidentsIncludeType) GetAllowedValues() []IncidentSearchIncidentsIncludeType {
	return allowedIncidentSearchIncidentsIncludeTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentSearchIncidentsIncludeType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentSearchIncidentsIncludeType(value)
	return nil
}

// NewIncidentSearchIncidentsIncludeTypeFromValue returns a pointer to a valid IncidentSearchIncidentsIncludeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentSearchIncidentsIncludeTypeFromValue(v string) (*IncidentSearchIncidentsIncludeType, error) {
	ev := IncidentSearchIncidentsIncludeType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentSearchIncidentsIncludeType: valid values are %v", v, allowedIncidentSearchIncidentsIncludeTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentSearchIncidentsIncludeType) IsValid() bool {
	for _, existing := range allowedIncidentSearchIncidentsIncludeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentSearchIncidentsIncludeType value.
func (v IncidentSearchIncidentsIncludeType) Ptr() *IncidentSearchIncidentsIncludeType {
	return &v
}
