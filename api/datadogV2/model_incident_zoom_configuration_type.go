// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentZoomConfigurationType Zoom configuration resource type.
type IncidentZoomConfigurationType string

// List of IncidentZoomConfigurationType.
const (
	INCIDENTZOOMCONFIGURATIONTYPE_ZOOM_CONFIGURATIONS IncidentZoomConfigurationType = "zoom_configurations"
)

var allowedIncidentZoomConfigurationTypeEnumValues = []IncidentZoomConfigurationType{
	INCIDENTZOOMCONFIGURATIONTYPE_ZOOM_CONFIGURATIONS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentZoomConfigurationType) GetAllowedValues() []IncidentZoomConfigurationType {
	return allowedIncidentZoomConfigurationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentZoomConfigurationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentZoomConfigurationType(value)
	return nil
}

// NewIncidentZoomConfigurationTypeFromValue returns a pointer to a valid IncidentZoomConfigurationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentZoomConfigurationTypeFromValue(v string) (*IncidentZoomConfigurationType, error) {
	ev := IncidentZoomConfigurationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentZoomConfigurationType: valid values are %v", v, allowedIncidentZoomConfigurationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentZoomConfigurationType) IsValid() bool {
	for _, existing := range allowedIncidentZoomConfigurationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentZoomConfigurationType value.
func (v IncidentZoomConfigurationType) Ptr() *IncidentZoomConfigurationType {
	return &v
}
