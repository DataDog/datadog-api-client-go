// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentsTimestampOverridesType The JSON:API type for timestamp overrides.
type IncidentsTimestampOverridesType string

// List of IncidentsTimestampOverridesType.
const (
	INCIDENTSTIMESTAMPOVERRIDESTYPE_INCIDENTS_TIMESTAMP_OVERRIDES IncidentsTimestampOverridesType = "incidents_timestamp_overrides"
)

var allowedIncidentsTimestampOverridesTypeEnumValues = []IncidentsTimestampOverridesType{
	INCIDENTSTIMESTAMPOVERRIDESTYPE_INCIDENTS_TIMESTAMP_OVERRIDES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentsTimestampOverridesType) GetAllowedValues() []IncidentsTimestampOverridesType {
	return allowedIncidentsTimestampOverridesTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentsTimestampOverridesType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentsTimestampOverridesType(value)
	return nil
}

// NewIncidentsTimestampOverridesTypeFromValue returns a pointer to a valid IncidentsTimestampOverridesType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentsTimestampOverridesTypeFromValue(v string) (*IncidentsTimestampOverridesType, error) {
	ev := IncidentsTimestampOverridesType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentsTimestampOverridesType: valid values are %v", v, allowedIncidentsTimestampOverridesTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentsTimestampOverridesType) IsValid() bool {
	for _, existing := range allowedIncidentsTimestampOverridesTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentsTimestampOverridesType value.
func (v IncidentsTimestampOverridesType) Ptr() *IncidentsTimestampOverridesType {
	return &v
}
