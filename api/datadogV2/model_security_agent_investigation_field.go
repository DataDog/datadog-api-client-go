// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityAgentInvestigationField An attribute of an investigation that `fields[investigations]` can select.
type SecurityAgentInvestigationField string

// List of SecurityAgentInvestigationField.
const (
	SECURITYAGENTINVESTIGATIONFIELD_SIGNAL_ID            SecurityAgentInvestigationField = "signal_id"
	SECURITYAGENTINVESTIGATIONFIELD_RULE_ID              SecurityAgentInvestigationField = "rule_id"
	SECURITYAGENTINVESTIGATIONFIELD_STATE                SecurityAgentInvestigationField = "state"
	SECURITYAGENTINVESTIGATIONFIELD_VERDICT              SecurityAgentInvestigationField = "verdict"
	SECURITYAGENTINVESTIGATIONFIELD_CONCLUSION           SecurityAgentInvestigationField = "conclusion"
	SECURITYAGENTINVESTIGATIONFIELD_SIGNAL               SecurityAgentInvestigationField = "signal"
	SECURITYAGENTINVESTIGATIONFIELD_KEY_ENTITIES         SecurityAgentInvestigationField = "key_entities"
	SECURITYAGENTINVESTIGATIONFIELD_KEY_EVIDENCE         SecurityAgentInvestigationField = "key_evidence"
	SECURITYAGENTINVESTIGATIONFIELD_STEPS                SecurityAgentInvestigationField = "steps"
	SECURITYAGENTINVESTIGATIONFIELD_RESPONSE_SUGGESTIONS SecurityAgentInvestigationField = "response_suggestions"
	SECURITYAGENTINVESTIGATIONFIELD_CREATED_AT           SecurityAgentInvestigationField = "created_at"
	SECURITYAGENTINVESTIGATIONFIELD_UPDATED_AT           SecurityAgentInvestigationField = "updated_at"
)

var allowedSecurityAgentInvestigationFieldEnumValues = []SecurityAgentInvestigationField{
	SECURITYAGENTINVESTIGATIONFIELD_SIGNAL_ID,
	SECURITYAGENTINVESTIGATIONFIELD_RULE_ID,
	SECURITYAGENTINVESTIGATIONFIELD_STATE,
	SECURITYAGENTINVESTIGATIONFIELD_VERDICT,
	SECURITYAGENTINVESTIGATIONFIELD_CONCLUSION,
	SECURITYAGENTINVESTIGATIONFIELD_SIGNAL,
	SECURITYAGENTINVESTIGATIONFIELD_KEY_ENTITIES,
	SECURITYAGENTINVESTIGATIONFIELD_KEY_EVIDENCE,
	SECURITYAGENTINVESTIGATIONFIELD_STEPS,
	SECURITYAGENTINVESTIGATIONFIELD_RESPONSE_SUGGESTIONS,
	SECURITYAGENTINVESTIGATIONFIELD_CREATED_AT,
	SECURITYAGENTINVESTIGATIONFIELD_UPDATED_AT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityAgentInvestigationField) GetAllowedValues() []SecurityAgentInvestigationField {
	return allowedSecurityAgentInvestigationFieldEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityAgentInvestigationField) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityAgentInvestigationField(value)
	return nil
}

// NewSecurityAgentInvestigationFieldFromValue returns a pointer to a valid SecurityAgentInvestigationField
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityAgentInvestigationFieldFromValue(v string) (*SecurityAgentInvestigationField, error) {
	ev := SecurityAgentInvestigationField(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityAgentInvestigationField: valid values are %v", v, allowedSecurityAgentInvestigationFieldEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityAgentInvestigationField) IsValid() bool {
	for _, existing := range allowedSecurityAgentInvestigationFieldEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityAgentInvestigationField value.
func (v SecurityAgentInvestigationField) Ptr() *SecurityAgentInvestigationField {
	return &v
}
