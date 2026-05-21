// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentJiraTemplateType Incident Jira template resource type.
type IncidentJiraTemplateType string

// List of IncidentJiraTemplateType.
const (
	INCIDENTJIRATEMPLATETYPE_INCIDENTS_JIRA_TEMPLATES IncidentJiraTemplateType = "incidents_jira_templates"
)

var allowedIncidentJiraTemplateTypeEnumValues = []IncidentJiraTemplateType{
	INCIDENTJIRATEMPLATETYPE_INCIDENTS_JIRA_TEMPLATES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentJiraTemplateType) GetAllowedValues() []IncidentJiraTemplateType {
	return allowedIncidentJiraTemplateTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentJiraTemplateType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentJiraTemplateType(value)
	return nil
}

// NewIncidentJiraTemplateTypeFromValue returns a pointer to a valid IncidentJiraTemplateType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentJiraTemplateTypeFromValue(v string) (*IncidentJiraTemplateType, error) {
	ev := IncidentJiraTemplateType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentJiraTemplateType: valid values are %v", v, allowedIncidentJiraTemplateTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentJiraTemplateType) IsValid() bool {
	for _, existing := range allowedIncidentJiraTemplateTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentJiraTemplateType value.
func (v IncidentJiraTemplateType) Ptr() *IncidentJiraTemplateType {
	return &v
}
