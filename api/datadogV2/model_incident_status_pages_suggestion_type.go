// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentStatusPagesSuggestionType Incident status pages suggestion resource type.
type IncidentStatusPagesSuggestionType string

// List of IncidentStatusPagesSuggestionType.
const (
	INCIDENTSTATUSPAGESSUGGESTIONTYPE_INCIDENT_STATUSPAGES_SUGGESTION IncidentStatusPagesSuggestionType = "incident_statuspages_suggestion"
)

var allowedIncidentStatusPagesSuggestionTypeEnumValues = []IncidentStatusPagesSuggestionType{
	INCIDENTSTATUSPAGESSUGGESTIONTYPE_INCIDENT_STATUSPAGES_SUGGESTION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentStatusPagesSuggestionType) GetAllowedValues() []IncidentStatusPagesSuggestionType {
	return allowedIncidentStatusPagesSuggestionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentStatusPagesSuggestionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentStatusPagesSuggestionType(value)
	return nil
}

// NewIncidentStatusPagesSuggestionTypeFromValue returns a pointer to a valid IncidentStatusPagesSuggestionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentStatusPagesSuggestionTypeFromValue(v string) (*IncidentStatusPagesSuggestionType, error) {
	ev := IncidentStatusPagesSuggestionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentStatusPagesSuggestionType: valid values are %v", v, allowedIncidentStatusPagesSuggestionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentStatusPagesSuggestionType) IsValid() bool {
	for _, existing := range allowedIncidentStatusPagesSuggestionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentStatusPagesSuggestionType value.
func (v IncidentStatusPagesSuggestionType) Ptr() *IncidentStatusPagesSuggestionType {
	return &v
}
