// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentPageRoleType The type of incident role for a page.
type IncidentPageRoleType string

// List of IncidentPageRoleType.
const (
	INCIDENTPAGEROLETYPE_INCIDENT_USER_DEFINED_ROLES IncidentPageRoleType = "incident_user_defined_roles"
	INCIDENTPAGEROLETYPE_INCIDENT_RESERVED_ROLES     IncidentPageRoleType = "incident_reserved_roles"
)

var allowedIncidentPageRoleTypeEnumValues = []IncidentPageRoleType{
	INCIDENTPAGEROLETYPE_INCIDENT_USER_DEFINED_ROLES,
	INCIDENTPAGEROLETYPE_INCIDENT_RESERVED_ROLES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentPageRoleType) GetAllowedValues() []IncidentPageRoleType {
	return allowedIncidentPageRoleTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentPageRoleType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentPageRoleType(value)
	return nil
}

// NewIncidentPageRoleTypeFromValue returns a pointer to a valid IncidentPageRoleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentPageRoleTypeFromValue(v string) (*IncidentPageRoleType, error) {
	ev := IncidentPageRoleType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentPageRoleType: valid values are %v", v, allowedIncidentPageRoleTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentPageRoleType) IsValid() bool {
	for _, existing := range allowedIncidentPageRoleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentPageRoleType value.
func (v IncidentPageRoleType) Ptr() *IncidentPageRoleType {
	return &v
}
