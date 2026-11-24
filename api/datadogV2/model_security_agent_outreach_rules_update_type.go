// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityAgentOutreachRulesUpdateType Type of an outreach policy rules update request.
type SecurityAgentOutreachRulesUpdateType string

// List of SecurityAgentOutreachRulesUpdateType.
const (
	SECURITYAGENTOUTREACHRULESUPDATETYPE_OUTREACH_POLICY_RULES_UPDATE SecurityAgentOutreachRulesUpdateType = "outreach_policy_rules_update"
)

var allowedSecurityAgentOutreachRulesUpdateTypeEnumValues = []SecurityAgentOutreachRulesUpdateType{
	SECURITYAGENTOUTREACHRULESUPDATETYPE_OUTREACH_POLICY_RULES_UPDATE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityAgentOutreachRulesUpdateType) GetAllowedValues() []SecurityAgentOutreachRulesUpdateType {
	return allowedSecurityAgentOutreachRulesUpdateTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityAgentOutreachRulesUpdateType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityAgentOutreachRulesUpdateType(value)
	return nil
}

// NewSecurityAgentOutreachRulesUpdateTypeFromValue returns a pointer to a valid SecurityAgentOutreachRulesUpdateType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityAgentOutreachRulesUpdateTypeFromValue(v string) (*SecurityAgentOutreachRulesUpdateType, error) {
	ev := SecurityAgentOutreachRulesUpdateType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityAgentOutreachRulesUpdateType: valid values are %v", v, allowedSecurityAgentOutreachRulesUpdateTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityAgentOutreachRulesUpdateType) IsValid() bool {
	for _, existing := range allowedSecurityAgentOutreachRulesUpdateTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityAgentOutreachRulesUpdateType value.
func (v SecurityAgentOutreachRulesUpdateType) Ptr() *SecurityAgentOutreachRulesUpdateType {
	return &v
}
