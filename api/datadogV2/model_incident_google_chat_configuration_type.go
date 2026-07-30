// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentGoogleChatConfigurationType Google Chat configuration resource type.
type IncidentGoogleChatConfigurationType string

// List of IncidentGoogleChatConfigurationType.
const (
	INCIDENTGOOGLECHATCONFIGURATIONTYPE_GOOGLE_CHAT_CONFIGURATIONS IncidentGoogleChatConfigurationType = "google_chat_configurations"
)

var allowedIncidentGoogleChatConfigurationTypeEnumValues = []IncidentGoogleChatConfigurationType{
	INCIDENTGOOGLECHATCONFIGURATIONTYPE_GOOGLE_CHAT_CONFIGURATIONS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentGoogleChatConfigurationType) GetAllowedValues() []IncidentGoogleChatConfigurationType {
	return allowedIncidentGoogleChatConfigurationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentGoogleChatConfigurationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentGoogleChatConfigurationType(value)
	return nil
}

// NewIncidentGoogleChatConfigurationTypeFromValue returns a pointer to a valid IncidentGoogleChatConfigurationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentGoogleChatConfigurationTypeFromValue(v string) (*IncidentGoogleChatConfigurationType, error) {
	ev := IncidentGoogleChatConfigurationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentGoogleChatConfigurationType: valid values are %v", v, allowedIncidentGoogleChatConfigurationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentGoogleChatConfigurationType) IsValid() bool {
	for _, existing := range allowedIncidentGoogleChatConfigurationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentGoogleChatConfigurationType value.
func (v IncidentGoogleChatConfigurationType) Ptr() *IncidentGoogleChatConfigurationType {
	return &v
}
