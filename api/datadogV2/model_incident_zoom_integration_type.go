// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentZoomIntegrationType Incident integration resource type.
type IncidentZoomIntegrationType string

// List of IncidentZoomIntegrationType.
const (
	INCIDENTZOOMINTEGRATIONTYPE_INCIDENT_INTEGRATIONS IncidentZoomIntegrationType = "incident_integrations"
)

var allowedIncidentZoomIntegrationTypeEnumValues = []IncidentZoomIntegrationType{
	INCIDENTZOOMINTEGRATIONTYPE_INCIDENT_INTEGRATIONS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentZoomIntegrationType) GetAllowedValues() []IncidentZoomIntegrationType {
	return allowedIncidentZoomIntegrationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentZoomIntegrationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentZoomIntegrationType(value)
	return nil
}

// NewIncidentZoomIntegrationTypeFromValue returns a pointer to a valid IncidentZoomIntegrationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentZoomIntegrationTypeFromValue(v string) (*IncidentZoomIntegrationType, error) {
	ev := IncidentZoomIntegrationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentZoomIntegrationType: valid values are %v", v, allowedIncidentZoomIntegrationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentZoomIntegrationType) IsValid() bool {
	for _, existing := range allowedIncidentZoomIntegrationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentZoomIntegrationType value.
func (v IncidentZoomIntegrationType) Ptr() *IncidentZoomIntegrationType {
	return &v
}
