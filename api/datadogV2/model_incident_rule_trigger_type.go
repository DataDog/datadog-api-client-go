// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentRuleTriggerType The trigger event for an incident rule.
type IncidentRuleTriggerType string

// List of IncidentRuleTriggerType.
const (
	INCIDENTRULETRIGGERTYPE_INCIDENT_SAVED_TRIGGER    IncidentRuleTriggerType = "incident_saved_trigger"
	INCIDENTRULETRIGGERTYPE_INCIDENT_CREATED_TRIGGER  IncidentRuleTriggerType = "incident_created_trigger"
	INCIDENTRULETRIGGERTYPE_INCIDENT_MODIFIED_TRIGGER IncidentRuleTriggerType = "incident_modified_trigger"
)

var allowedIncidentRuleTriggerTypeEnumValues = []IncidentRuleTriggerType{
	INCIDENTRULETRIGGERTYPE_INCIDENT_SAVED_TRIGGER,
	INCIDENTRULETRIGGERTYPE_INCIDENT_CREATED_TRIGGER,
	INCIDENTRULETRIGGERTYPE_INCIDENT_MODIFIED_TRIGGER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentRuleTriggerType) GetAllowedValues() []IncidentRuleTriggerType {
	return allowedIncidentRuleTriggerTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentRuleTriggerType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentRuleTriggerType(value)
	return nil
}

// NewIncidentRuleTriggerTypeFromValue returns a pointer to a valid IncidentRuleTriggerType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentRuleTriggerTypeFromValue(v string) (*IncidentRuleTriggerType, error) {
	ev := IncidentRuleTriggerType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentRuleTriggerType: valid values are %v", v, allowedIncidentRuleTriggerTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentRuleTriggerType) IsValid() bool {
	for _, existing := range allowedIncidentRuleTriggerTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentRuleTriggerType value.
func (v IncidentRuleTriggerType) Ptr() *IncidentRuleTriggerType {
	return &v
}
