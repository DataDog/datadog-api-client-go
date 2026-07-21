// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentImpactFieldValueType The type of an impact field.
type IncidentImpactFieldValueType string

// List of IncidentImpactFieldValueType.
const (
	INCIDENTIMPACTFIELDVALUETYPE_DROPDOWN    IncidentImpactFieldValueType = "dropdown"
	INCIDENTIMPACTFIELDVALUETYPE_TEXT        IncidentImpactFieldValueType = "text"
	INCIDENTIMPACTFIELDVALUETYPE_TEXTARRAY   IncidentImpactFieldValueType = "textarray"
	INCIDENTIMPACTFIELDVALUETYPE_METRICTAG   IncidentImpactFieldValueType = "metrictag"
	INCIDENTIMPACTFIELDVALUETYPE_NUMBER      IncidentImpactFieldValueType = "number"
	INCIDENTIMPACTFIELDVALUETYPE_DATETIME    IncidentImpactFieldValueType = "datetime"
	INCIDENTIMPACTFIELDVALUETYPE_MULTISELECT IncidentImpactFieldValueType = "multiselect"
)

var allowedIncidentImpactFieldValueTypeEnumValues = []IncidentImpactFieldValueType{
	INCIDENTIMPACTFIELDVALUETYPE_DROPDOWN,
	INCIDENTIMPACTFIELDVALUETYPE_TEXT,
	INCIDENTIMPACTFIELDVALUETYPE_TEXTARRAY,
	INCIDENTIMPACTFIELDVALUETYPE_METRICTAG,
	INCIDENTIMPACTFIELDVALUETYPE_NUMBER,
	INCIDENTIMPACTFIELDVALUETYPE_DATETIME,
	INCIDENTIMPACTFIELDVALUETYPE_MULTISELECT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentImpactFieldValueType) GetAllowedValues() []IncidentImpactFieldValueType {
	return allowedIncidentImpactFieldValueTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentImpactFieldValueType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentImpactFieldValueType(value)
	return nil
}

// NewIncidentImpactFieldValueTypeFromValue returns a pointer to a valid IncidentImpactFieldValueType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentImpactFieldValueTypeFromValue(v string) (*IncidentImpactFieldValueType, error) {
	ev := IncidentImpactFieldValueType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentImpactFieldValueType: valid values are %v", v, allowedIncidentImpactFieldValueTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentImpactFieldValueType) IsValid() bool {
	for _, existing := range allowedIncidentImpactFieldValueTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentImpactFieldValueType value.
func (v IncidentImpactFieldValueType) Ptr() *IncidentImpactFieldValueType {
	return &v
}
