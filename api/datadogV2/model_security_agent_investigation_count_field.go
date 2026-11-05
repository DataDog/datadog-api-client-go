// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityAgentInvestigationCountField Field to count completed investigations by.
type SecurityAgentInvestigationCountField string

// List of SecurityAgentInvestigationCountField.
const (
	SECURITYAGENTINVESTIGATIONCOUNTFIELD_VERDICT       SecurityAgentInvestigationCountField = "verdict"
	SECURITYAGENTINVESTIGATIONCOUNTFIELD_RULE_TYPE     SecurityAgentInvestigationCountField = "rule_type"
	SECURITYAGENTINVESTIGATIONCOUNTFIELD_SIGNAL_SOURCE SecurityAgentInvestigationCountField = "signal_source"
	SECURITYAGENTINVESTIGATIONCOUNTFIELD_ENTITY        SecurityAgentInvestigationCountField = "entity"
)

var allowedSecurityAgentInvestigationCountFieldEnumValues = []SecurityAgentInvestigationCountField{
	SECURITYAGENTINVESTIGATIONCOUNTFIELD_VERDICT,
	SECURITYAGENTINVESTIGATIONCOUNTFIELD_RULE_TYPE,
	SECURITYAGENTINVESTIGATIONCOUNTFIELD_SIGNAL_SOURCE,
	SECURITYAGENTINVESTIGATIONCOUNTFIELD_ENTITY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityAgentInvestigationCountField) GetAllowedValues() []SecurityAgentInvestigationCountField {
	return allowedSecurityAgentInvestigationCountFieldEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityAgentInvestigationCountField) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityAgentInvestigationCountField(value)
	return nil
}

// NewSecurityAgentInvestigationCountFieldFromValue returns a pointer to a valid SecurityAgentInvestigationCountField
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityAgentInvestigationCountFieldFromValue(v string) (*SecurityAgentInvestigationCountField, error) {
	ev := SecurityAgentInvestigationCountField(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityAgentInvestigationCountField: valid values are %v", v, allowedSecurityAgentInvestigationCountFieldEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityAgentInvestigationCountField) IsValid() bool {
	for _, existing := range allowedSecurityAgentInvestigationCountFieldEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityAgentInvestigationCountField value.
func (v SecurityAgentInvestigationCountField) Ptr() *SecurityAgentInvestigationCountField {
	return &v
}
