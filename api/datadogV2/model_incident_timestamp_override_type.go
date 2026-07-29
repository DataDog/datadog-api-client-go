// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentTimestampOverrideType Incident timestamp override resource type.
type IncidentTimestampOverrideType string

// List of IncidentTimestampOverrideType.
const (
	INCIDENTTIMESTAMPOVERRIDETYPE_INCIDENTS_TIMESTAMP_OVERRIDES IncidentTimestampOverrideType = "incidents_timestamp_overrides"
)

var allowedIncidentTimestampOverrideTypeEnumValues = []IncidentTimestampOverrideType{
	INCIDENTTIMESTAMPOVERRIDETYPE_INCIDENTS_TIMESTAMP_OVERRIDES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentTimestampOverrideType) GetAllowedValues() []IncidentTimestampOverrideType {
	return allowedIncidentTimestampOverrideTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentTimestampOverrideType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentTimestampOverrideType(value)
	return nil
}

// NewIncidentTimestampOverrideTypeFromValue returns a pointer to a valid IncidentTimestampOverrideType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentTimestampOverrideTypeFromValue(v string) (*IncidentTimestampOverrideType, error) {
	ev := IncidentTimestampOverrideType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentTimestampOverrideType: valid values are %v", v, allowedIncidentTimestampOverrideTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentTimestampOverrideType) IsValid() bool {
	for _, existing := range allowedIncidentTimestampOverrideTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentTimestampOverrideType value.
func (v IncidentTimestampOverrideType) Ptr() *IncidentTimestampOverrideType {
	return &v
}
