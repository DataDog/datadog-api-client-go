// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentTimestampType The type of timestamp to override.
type IncidentTimestampType string

// List of IncidentTimestampType.
const (
	INCIDENTTIMESTAMPTYPE_CREATED  IncidentTimestampType = "created"
	INCIDENTTIMESTAMPTYPE_DETECTED IncidentTimestampType = "detected"
	INCIDENTTIMESTAMPTYPE_RESOLVED IncidentTimestampType = "resolved"
	INCIDENTTIMESTAMPTYPE_DECLARED IncidentTimestampType = "declared"
)

var allowedIncidentTimestampTypeEnumValues = []IncidentTimestampType{
	INCIDENTTIMESTAMPTYPE_CREATED,
	INCIDENTTIMESTAMPTYPE_DETECTED,
	INCIDENTTIMESTAMPTYPE_RESOLVED,
	INCIDENTTIMESTAMPTYPE_DECLARED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentTimestampType) GetAllowedValues() []IncidentTimestampType {
	return allowedIncidentTimestampTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentTimestampType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentTimestampType(value)
	return nil
}

// NewIncidentTimestampTypeFromValue returns a pointer to a valid IncidentTimestampType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentTimestampTypeFromValue(v string) (*IncidentTimestampType, error) {
	ev := IncidentTimestampType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentTimestampType: valid values are %v", v, allowedIncidentTimestampTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentTimestampType) IsValid() bool {
	for _, existing := range allowedIncidentTimestampTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentTimestampType value.
func (v IncidentTimestampType) Ptr() *IncidentTimestampType {
	return &v
}
