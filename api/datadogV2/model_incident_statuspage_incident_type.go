// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentStatuspageIncidentType Statuspage incident integration resource type.
type IncidentStatuspageIncidentType string

// List of IncidentStatuspageIncidentType.
const (
	INCIDENTSTATUSPAGEINCIDENTTYPE_INCIDENT_INTEGRATIONS IncidentStatuspageIncidentType = "incident_integrations"
)

var allowedIncidentStatuspageIncidentTypeEnumValues = []IncidentStatuspageIncidentType{
	INCIDENTSTATUSPAGEINCIDENTTYPE_INCIDENT_INTEGRATIONS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentStatuspageIncidentType) GetAllowedValues() []IncidentStatuspageIncidentType {
	return allowedIncidentStatuspageIncidentTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentStatuspageIncidentType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentStatuspageIncidentType(value)
	return nil
}

// NewIncidentStatuspageIncidentTypeFromValue returns a pointer to a valid IncidentStatuspageIncidentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentStatuspageIncidentTypeFromValue(v string) (*IncidentStatuspageIncidentType, error) {
	ev := IncidentStatuspageIncidentType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentStatuspageIncidentType: valid values are %v", v, allowedIncidentStatuspageIncidentTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentStatuspageIncidentType) IsValid() bool {
	for _, existing := range allowedIncidentStatuspageIncidentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentStatuspageIncidentType value.
func (v IncidentStatuspageIncidentType) Ptr() *IncidentStatuspageIncidentType {
	return &v
}
