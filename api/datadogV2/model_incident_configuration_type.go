// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentConfigurationType Incident configuration resource type.
type IncidentConfigurationType string

// List of IncidentConfigurationType.
const (
	INCIDENTCONFIGURATIONTYPE_INCIDENTS_CONFIGURATIONS IncidentConfigurationType = "incidents_configurations"
)

var allowedIncidentConfigurationTypeEnumValues = []IncidentConfigurationType{
	INCIDENTCONFIGURATIONTYPE_INCIDENTS_CONFIGURATIONS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentConfigurationType) GetAllowedValues() []IncidentConfigurationType {
	return allowedIncidentConfigurationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentConfigurationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentConfigurationType(value)
	return nil
}

// NewIncidentConfigurationTypeFromValue returns a pointer to a valid IncidentConfigurationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentConfigurationTypeFromValue(v string) (*IncidentConfigurationType, error) {
	ev := IncidentConfigurationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentConfigurationType: valid values are %v", v, allowedIncidentConfigurationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentConfigurationType) IsValid() bool {
	for _, existing := range allowedIncidentConfigurationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentConfigurationType value.
func (v IncidentConfigurationType) Ptr() *IncidentConfigurationType {
	return &v
}
