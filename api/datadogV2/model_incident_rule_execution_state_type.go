// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentRuleExecutionStateType Incident rule execution state resource type.
type IncidentRuleExecutionStateType string

// List of IncidentRuleExecutionStateType.
const (
	INCIDENTRULEEXECUTIONSTATETYPE_INCIDENT_RULE_EXECUTION_STATES IncidentRuleExecutionStateType = "incident_rule_execution_states"
)

var allowedIncidentRuleExecutionStateTypeEnumValues = []IncidentRuleExecutionStateType{
	INCIDENTRULEEXECUTIONSTATETYPE_INCIDENT_RULE_EXECUTION_STATES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentRuleExecutionStateType) GetAllowedValues() []IncidentRuleExecutionStateType {
	return allowedIncidentRuleExecutionStateTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentRuleExecutionStateType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentRuleExecutionStateType(value)
	return nil
}

// NewIncidentRuleExecutionStateTypeFromValue returns a pointer to a valid IncidentRuleExecutionStateType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentRuleExecutionStateTypeFromValue(v string) (*IncidentRuleExecutionStateType, error) {
	ev := IncidentRuleExecutionStateType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentRuleExecutionStateType: valid values are %v", v, allowedIncidentRuleExecutionStateTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentRuleExecutionStateType) IsValid() bool {
	for _, existing := range allowedIncidentRuleExecutionStateTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentRuleExecutionStateType value.
func (v IncidentRuleExecutionStateType) Ptr() *IncidentRuleExecutionStateType {
	return &v
}
