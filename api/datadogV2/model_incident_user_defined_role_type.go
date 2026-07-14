// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentUserDefinedRoleType Incident user-defined role resource type.
type IncidentUserDefinedRoleType string

// List of IncidentUserDefinedRoleType.
const (
	INCIDENTUSERDEFINEDROLETYPE_INCIDENT_USER_DEFINED_ROLES IncidentUserDefinedRoleType = "incident_user_defined_roles"
)

var allowedIncidentUserDefinedRoleTypeEnumValues = []IncidentUserDefinedRoleType{
	INCIDENTUSERDEFINEDROLETYPE_INCIDENT_USER_DEFINED_ROLES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentUserDefinedRoleType) GetAllowedValues() []IncidentUserDefinedRoleType {
	return allowedIncidentUserDefinedRoleTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentUserDefinedRoleType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentUserDefinedRoleType(value)
	return nil
}

// NewIncidentUserDefinedRoleTypeFromValue returns a pointer to a valid IncidentUserDefinedRoleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentUserDefinedRoleTypeFromValue(v string) (*IncidentUserDefinedRoleType, error) {
	ev := IncidentUserDefinedRoleType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentUserDefinedRoleType: valid values are %v", v, allowedIncidentUserDefinedRoleTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentUserDefinedRoleType) IsValid() bool {
	for _, existing := range allowedIncidentUserDefinedRoleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentUserDefinedRoleType value.
func (v IncidentUserDefinedRoleType) Ptr() *IncidentUserDefinedRoleType {
	return &v
}
