// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentRoleAssignmentType Incident role assignment resource type.
type IncidentRoleAssignmentType string

// List of IncidentRoleAssignmentType.
const (
	INCIDENTROLEASSIGNMENTTYPE_INCIDENT_ROLE_ASSIGNMENTS IncidentRoleAssignmentType = "incident_role_assignments"
)

var allowedIncidentRoleAssignmentTypeEnumValues = []IncidentRoleAssignmentType{
	INCIDENTROLEASSIGNMENTTYPE_INCIDENT_ROLE_ASSIGNMENTS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentRoleAssignmentType) GetAllowedValues() []IncidentRoleAssignmentType {
	return allowedIncidentRoleAssignmentTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentRoleAssignmentType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentRoleAssignmentType(value)
	return nil
}

// NewIncidentRoleAssignmentTypeFromValue returns a pointer to a valid IncidentRoleAssignmentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentRoleAssignmentTypeFromValue(v string) (*IncidentRoleAssignmentType, error) {
	ev := IncidentRoleAssignmentType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentRoleAssignmentType: valid values are %v", v, allowedIncidentRoleAssignmentTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentRoleAssignmentType) IsValid() bool {
	for _, existing := range allowedIncidentRoleAssignmentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentRoleAssignmentType value.
func (v IncidentRoleAssignmentType) Ptr() *IncidentRoleAssignmentType {
	return &v
}
