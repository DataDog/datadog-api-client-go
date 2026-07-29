// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentOrgSettingsType Incident org settings resource type.
type IncidentOrgSettingsType string

// List of IncidentOrgSettingsType.
const (
	INCIDENTORGSETTINGSTYPE_INCIDENT_ORG_SETTINGS IncidentOrgSettingsType = "incident_org_settings"
)

var allowedIncidentOrgSettingsTypeEnumValues = []IncidentOrgSettingsType{
	INCIDENTORGSETTINGSTYPE_INCIDENT_ORG_SETTINGS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentOrgSettingsType) GetAllowedValues() []IncidentOrgSettingsType {
	return allowedIncidentOrgSettingsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentOrgSettingsType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentOrgSettingsType(value)
	return nil
}

// NewIncidentOrgSettingsTypeFromValue returns a pointer to a valid IncidentOrgSettingsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentOrgSettingsTypeFromValue(v string) (*IncidentOrgSettingsType, error) {
	ev := IncidentOrgSettingsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentOrgSettingsType: valid values are %v", v, allowedIncidentOrgSettingsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentOrgSettingsType) IsValid() bool {
	for _, existing := range allowedIncidentOrgSettingsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentOrgSettingsType value.
func (v IncidentOrgSettingsType) Ptr() *IncidentOrgSettingsType {
	return &v
}
