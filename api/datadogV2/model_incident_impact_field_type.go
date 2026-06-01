// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentImpactFieldType Impact field resource type.
type IncidentImpactFieldType string

// List of IncidentImpactFieldType.
const (
	INCIDENTIMPACTFIELDTYPE_IMPACT_FIELDS IncidentImpactFieldType = "impact_fields"
)

var allowedIncidentImpactFieldTypeEnumValues = []IncidentImpactFieldType{
	INCIDENTIMPACTFIELDTYPE_IMPACT_FIELDS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentImpactFieldType) GetAllowedValues() []IncidentImpactFieldType {
	return allowedIncidentImpactFieldTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentImpactFieldType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentImpactFieldType(value)
	return nil
}

// NewIncidentImpactFieldTypeFromValue returns a pointer to a valid IncidentImpactFieldType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentImpactFieldTypeFromValue(v string) (*IncidentImpactFieldType, error) {
	ev := IncidentImpactFieldType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentImpactFieldType: valid values are %v", v, allowedIncidentImpactFieldTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentImpactFieldType) IsValid() bool {
	for _, existing := range allowedIncidentImpactFieldTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentImpactFieldType value.
func (v IncidentImpactFieldType) Ptr() *IncidentImpactFieldType {
	return &v
}
