// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentResponderType Incident responder resource type.
type IncidentResponderType string

// List of IncidentResponderType.
const (
	INCIDENTRESPONDERTYPE_INCIDENT_RESPONDERS IncidentResponderType = "incident_responders"
)

var allowedIncidentResponderTypeEnumValues = []IncidentResponderType{
	INCIDENTRESPONDERTYPE_INCIDENT_RESPONDERS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentResponderType) GetAllowedValues() []IncidentResponderType {
	return allowedIncidentResponderTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentResponderType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentResponderType(value)
	return nil
}

// NewIncidentResponderTypeFromValue returns a pointer to a valid IncidentResponderType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentResponderTypeFromValue(v string) (*IncidentResponderType, error) {
	ev := IncidentResponderType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentResponderType: valid values are %v", v, allowedIncidentResponderTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentResponderType) IsValid() bool {
	for _, existing := range allowedIncidentResponderTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentResponderType value.
func (v IncidentResponderType) Ptr() *IncidentResponderType {
	return &v
}
