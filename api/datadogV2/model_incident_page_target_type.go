// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentPageTargetType Type of page target for incident pages.
type IncidentPageTargetType string

// List of IncidentPageTargetType.
const (
	INCIDENTPAGETARGETTYPE_TEAM_HANDLE IncidentPageTargetType = "team_handle"
	INCIDENTPAGETARGETTYPE_TEAM_UUID   IncidentPageTargetType = "team_uuid"
	INCIDENTPAGETARGETTYPE_USER_UUID   IncidentPageTargetType = "user_uuid"
)

var allowedIncidentPageTargetTypeEnumValues = []IncidentPageTargetType{
	INCIDENTPAGETARGETTYPE_TEAM_HANDLE,
	INCIDENTPAGETARGETTYPE_TEAM_UUID,
	INCIDENTPAGETARGETTYPE_USER_UUID,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentPageTargetType) GetAllowedValues() []IncidentPageTargetType {
	return allowedIncidentPageTargetTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentPageTargetType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentPageTargetType(value)
	return nil
}

// NewIncidentPageTargetTypeFromValue returns a pointer to a valid IncidentPageTargetType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentPageTargetTypeFromValue(v string) (*IncidentPageTargetType, error) {
	ev := IncidentPageTargetType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentPageTargetType: valid values are %v", v, allowedIncidentPageTargetTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentPageTargetType) IsValid() bool {
	for _, existing := range allowedIncidentPageTargetTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentPageTargetType value.
func (v IncidentPageTargetType) Ptr() *IncidentPageTargetType {
	return &v
}
