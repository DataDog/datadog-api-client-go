// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentRenderedTemplateType Rendered template resource type.
type IncidentRenderedTemplateType string

// List of IncidentRenderedTemplateType.
const (
	INCIDENTRENDEREDTEMPLATETYPE_RENDERED_TEMPLATES IncidentRenderedTemplateType = "rendered_templates"
)

var allowedIncidentRenderedTemplateTypeEnumValues = []IncidentRenderedTemplateType{
	INCIDENTRENDEREDTEMPLATETYPE_RENDERED_TEMPLATES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentRenderedTemplateType) GetAllowedValues() []IncidentRenderedTemplateType {
	return allowedIncidentRenderedTemplateTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentRenderedTemplateType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentRenderedTemplateType(value)
	return nil
}

// NewIncidentRenderedTemplateTypeFromValue returns a pointer to a valid IncidentRenderedTemplateType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentRenderedTemplateTypeFromValue(v string) (*IncidentRenderedTemplateType, error) {
	ev := IncidentRenderedTemplateType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentRenderedTemplateType: valid values are %v", v, allowedIncidentRenderedTemplateTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentRenderedTemplateType) IsValid() bool {
	for _, existing := range allowedIncidentRenderedTemplateTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentRenderedTemplateType value.
func (v IncidentRenderedTemplateType) Ptr() *IncidentRenderedTemplateType {
	return &v
}
