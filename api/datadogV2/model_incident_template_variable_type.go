// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentTemplateVariableType Template variable resource type.
type IncidentTemplateVariableType string

// List of IncidentTemplateVariableType.
const (
	INCIDENTTEMPLATEVARIABLETYPE_TEMPLATE_VARIABLES IncidentTemplateVariableType = "template_variables"
)

var allowedIncidentTemplateVariableTypeEnumValues = []IncidentTemplateVariableType{
	INCIDENTTEMPLATEVARIABLETYPE_TEMPLATE_VARIABLES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentTemplateVariableType) GetAllowedValues() []IncidentTemplateVariableType {
	return allowedIncidentTemplateVariableTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentTemplateVariableType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentTemplateVariableType(value)
	return nil
}

// NewIncidentTemplateVariableTypeFromValue returns a pointer to a valid IncidentTemplateVariableType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentTemplateVariableTypeFromValue(v string) (*IncidentTemplateVariableType, error) {
	ev := IncidentTemplateVariableType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentTemplateVariableType: valid values are %v", v, allowedIncidentTemplateVariableTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentTemplateVariableType) IsValid() bool {
	for _, existing := range allowedIncidentTemplateVariableTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentTemplateVariableType value.
func (v IncidentTemplateVariableType) Ptr() *IncidentTemplateVariableType {
	return &v
}
