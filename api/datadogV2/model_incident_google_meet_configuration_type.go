// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentGoogleMeetConfigurationType Google Meet configuration resource type.
type IncidentGoogleMeetConfigurationType string

// List of IncidentGoogleMeetConfigurationType.
const (
	INCIDENTGOOGLEMEETCONFIGURATIONTYPE_GOOGLE_MEET_CONFIGURATIONS IncidentGoogleMeetConfigurationType = "google_meet_configurations"
)

var allowedIncidentGoogleMeetConfigurationTypeEnumValues = []IncidentGoogleMeetConfigurationType{
	INCIDENTGOOGLEMEETCONFIGURATIONTYPE_GOOGLE_MEET_CONFIGURATIONS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentGoogleMeetConfigurationType) GetAllowedValues() []IncidentGoogleMeetConfigurationType {
	return allowedIncidentGoogleMeetConfigurationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentGoogleMeetConfigurationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentGoogleMeetConfigurationType(value)
	return nil
}

// NewIncidentGoogleMeetConfigurationTypeFromValue returns a pointer to a valid IncidentGoogleMeetConfigurationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentGoogleMeetConfigurationTypeFromValue(v string) (*IncidentGoogleMeetConfigurationType, error) {
	ev := IncidentGoogleMeetConfigurationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentGoogleMeetConfigurationType: valid values are %v", v, allowedIncidentGoogleMeetConfigurationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentGoogleMeetConfigurationType) IsValid() bool {
	for _, existing := range allowedIncidentGoogleMeetConfigurationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentGoogleMeetConfigurationType value.
func (v IncidentGoogleMeetConfigurationType) Ptr() *IncidentGoogleMeetConfigurationType {
	return &v
}
