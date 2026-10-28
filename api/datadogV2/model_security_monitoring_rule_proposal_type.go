// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringRuleProposalType The rule proposal resource type. The value is always `rule_proposals`.
type SecurityMonitoringRuleProposalType string

// List of SecurityMonitoringRuleProposalType.
const (
	SECURITYMONITORINGRULEPROPOSALTYPE_RULE_PROPOSALS SecurityMonitoringRuleProposalType = "rule_proposals"
)

var allowedSecurityMonitoringRuleProposalTypeEnumValues = []SecurityMonitoringRuleProposalType{
	SECURITYMONITORINGRULEPROPOSALTYPE_RULE_PROPOSALS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityMonitoringRuleProposalType) GetAllowedValues() []SecurityMonitoringRuleProposalType {
	return allowedSecurityMonitoringRuleProposalTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityMonitoringRuleProposalType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringRuleProposalType(value)
	return nil
}

// NewSecurityMonitoringRuleProposalTypeFromValue returns a pointer to a valid SecurityMonitoringRuleProposalType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityMonitoringRuleProposalTypeFromValue(v string) (*SecurityMonitoringRuleProposalType, error) {
	ev := SecurityMonitoringRuleProposalType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringRuleProposalType: valid values are %v", v, allowedSecurityMonitoringRuleProposalTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityMonitoringRuleProposalType) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringRuleProposalTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringRuleProposalType value.
func (v SecurityMonitoringRuleProposalType) Ptr() *SecurityMonitoringRuleProposalType {
	return &v
}
