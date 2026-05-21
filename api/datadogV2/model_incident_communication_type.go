// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentCommunicationType Incident communication resource type.
type IncidentCommunicationType string

// List of IncidentCommunicationType.
const (
	INCIDENTCOMMUNICATIONTYPE_COMMUNICATION IncidentCommunicationType = "communication"
)

var allowedIncidentCommunicationTypeEnumValues = []IncidentCommunicationType{
	INCIDENTCOMMUNICATIONTYPE_COMMUNICATION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentCommunicationType) GetAllowedValues() []IncidentCommunicationType {
	return allowedIncidentCommunicationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentCommunicationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentCommunicationType(value)
	return nil
}

// NewIncidentCommunicationTypeFromValue returns a pointer to a valid IncidentCommunicationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentCommunicationTypeFromValue(v string) (*IncidentCommunicationType, error) {
	ev := IncidentCommunicationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentCommunicationType: valid values are %v", v, allowedIncidentCommunicationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentCommunicationType) IsValid() bool {
	for _, existing := range allowedIncidentCommunicationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentCommunicationType value.
func (v IncidentCommunicationType) Ptr() *IncidentCommunicationType {
	return &v
}
