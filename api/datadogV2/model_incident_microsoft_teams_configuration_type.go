// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentMicrosoftTeamsConfigurationType Microsoft Teams configuration resource type.
type IncidentMicrosoftTeamsConfigurationType string

// List of IncidentMicrosoftTeamsConfigurationType.
const (
	INCIDENTMICROSOFTTEAMSCONFIGURATIONTYPE_MICROSOFT_TEAMS_CONFIGURATIONS IncidentMicrosoftTeamsConfigurationType = "microsoft_teams_configurations"
)

var allowedIncidentMicrosoftTeamsConfigurationTypeEnumValues = []IncidentMicrosoftTeamsConfigurationType{
	INCIDENTMICROSOFTTEAMSCONFIGURATIONTYPE_MICROSOFT_TEAMS_CONFIGURATIONS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentMicrosoftTeamsConfigurationType) GetAllowedValues() []IncidentMicrosoftTeamsConfigurationType {
	return allowedIncidentMicrosoftTeamsConfigurationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentMicrosoftTeamsConfigurationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentMicrosoftTeamsConfigurationType(value)
	return nil
}

// NewIncidentMicrosoftTeamsConfigurationTypeFromValue returns a pointer to a valid IncidentMicrosoftTeamsConfigurationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentMicrosoftTeamsConfigurationTypeFromValue(v string) (*IncidentMicrosoftTeamsConfigurationType, error) {
	ev := IncidentMicrosoftTeamsConfigurationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentMicrosoftTeamsConfigurationType: valid values are %v", v, allowedIncidentMicrosoftTeamsConfigurationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentMicrosoftTeamsConfigurationType) IsValid() bool {
	for _, existing := range allowedIncidentMicrosoftTeamsConfigurationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentMicrosoftTeamsConfigurationType value.
func (v IncidentMicrosoftTeamsConfigurationType) Ptr() *IncidentMicrosoftTeamsConfigurationType {
	return &v
}
