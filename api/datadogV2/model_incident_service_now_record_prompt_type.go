// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentServiceNowRecordPromptType ServiceNow record prompt resource type.
type IncidentServiceNowRecordPromptType string

// List of IncidentServiceNowRecordPromptType.
const (
	INCIDENTSERVICENOWRECORDPROMPTTYPE_INCIDENT_SERVICENOW_RECORD_PROMPT IncidentServiceNowRecordPromptType = "incident_servicenow_record_prompt"
)

var allowedIncidentServiceNowRecordPromptTypeEnumValues = []IncidentServiceNowRecordPromptType{
	INCIDENTSERVICENOWRECORDPROMPTTYPE_INCIDENT_SERVICENOW_RECORD_PROMPT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentServiceNowRecordPromptType) GetAllowedValues() []IncidentServiceNowRecordPromptType {
	return allowedIncidentServiceNowRecordPromptTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentServiceNowRecordPromptType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentServiceNowRecordPromptType(value)
	return nil
}

// NewIncidentServiceNowRecordPromptTypeFromValue returns a pointer to a valid IncidentServiceNowRecordPromptType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentServiceNowRecordPromptTypeFromValue(v string) (*IncidentServiceNowRecordPromptType, error) {
	ev := IncidentServiceNowRecordPromptType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentServiceNowRecordPromptType: valid values are %v", v, allowedIncidentServiceNowRecordPromptTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentServiceNowRecordPromptType) IsValid() bool {
	for _, existing := range allowedIncidentServiceNowRecordPromptTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentServiceNowRecordPromptType value.
func (v IncidentServiceNowRecordPromptType) Ptr() *IncidentServiceNowRecordPromptType {
	return &v
}
