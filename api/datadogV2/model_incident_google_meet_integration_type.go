// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentGoogleMeetIntegrationType Incident integration resource type.
type IncidentGoogleMeetIntegrationType string

// List of IncidentGoogleMeetIntegrationType.
const (
	INCIDENTGOOGLEMEETINTEGRATIONTYPE_INCIDENT_INTEGRATIONS IncidentGoogleMeetIntegrationType = "incident_integrations"
)

var allowedIncidentGoogleMeetIntegrationTypeEnumValues = []IncidentGoogleMeetIntegrationType{
	INCIDENTGOOGLEMEETINTEGRATIONTYPE_INCIDENT_INTEGRATIONS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentGoogleMeetIntegrationType) GetAllowedValues() []IncidentGoogleMeetIntegrationType {
	return allowedIncidentGoogleMeetIntegrationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentGoogleMeetIntegrationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentGoogleMeetIntegrationType(value)
	return nil
}

// NewIncidentGoogleMeetIntegrationTypeFromValue returns a pointer to a valid IncidentGoogleMeetIntegrationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentGoogleMeetIntegrationTypeFromValue(v string) (*IncidentGoogleMeetIntegrationType, error) {
	ev := IncidentGoogleMeetIntegrationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentGoogleMeetIntegrationType: valid values are %v", v, allowedIncidentGoogleMeetIntegrationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentGoogleMeetIntegrationType) IsValid() bool {
	for _, existing := range allowedIncidentGoogleMeetIntegrationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentGoogleMeetIntegrationType value.
func (v IncidentGoogleMeetIntegrationType) Ptr() *IncidentGoogleMeetIntegrationType {
	return &v
}
