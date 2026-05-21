// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentReservedRoleType Incident reserved role resource type.
type IncidentReservedRoleType string

// List of IncidentReservedRoleType.
const (
	INCIDENTRESERVEDROLETYPE_INCIDENT_RESERVED_ROLES IncidentReservedRoleType = "incident_reserved_roles"
)

var allowedIncidentReservedRoleTypeEnumValues = []IncidentReservedRoleType{
	INCIDENTRESERVEDROLETYPE_INCIDENT_RESERVED_ROLES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentReservedRoleType) GetAllowedValues() []IncidentReservedRoleType {
	return allowedIncidentReservedRoleTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentReservedRoleType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentReservedRoleType(value)
	return nil
}

// NewIncidentReservedRoleTypeFromValue returns a pointer to a valid IncidentReservedRoleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentReservedRoleTypeFromValue(v string) (*IncidentReservedRoleType, error) {
	ev := IncidentReservedRoleType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentReservedRoleType: valid values are %v", v, allowedIncidentReservedRoleTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentReservedRoleType) IsValid() bool {
	for _, existing := range allowedIncidentReservedRoleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentReservedRoleType value.
func (v IncidentReservedRoleType) Ptr() *IncidentReservedRoleType {
	return &v
}
