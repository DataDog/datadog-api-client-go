// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringRuleProposalInputType The type of input to generate the proposed rule from: `natural_language` for a
// plain-English description of the behavior to detect, `spl` for a Splunk SPL search or
// saved search definition, or `kql` for a Microsoft KQL query. SPL inputs must be
// self-contained, with search macros expanded.
type SecurityMonitoringRuleProposalInputType string

// List of SecurityMonitoringRuleProposalInputType.
const (
	SECURITYMONITORINGRULEPROPOSALINPUTTYPE_NATURAL_LANGUAGE SecurityMonitoringRuleProposalInputType = "natural_language"
	SECURITYMONITORINGRULEPROPOSALINPUTTYPE_SPL              SecurityMonitoringRuleProposalInputType = "spl"
	SECURITYMONITORINGRULEPROPOSALINPUTTYPE_KQL              SecurityMonitoringRuleProposalInputType = "kql"
)

var allowedSecurityMonitoringRuleProposalInputTypeEnumValues = []SecurityMonitoringRuleProposalInputType{
	SECURITYMONITORINGRULEPROPOSALINPUTTYPE_NATURAL_LANGUAGE,
	SECURITYMONITORINGRULEPROPOSALINPUTTYPE_SPL,
	SECURITYMONITORINGRULEPROPOSALINPUTTYPE_KQL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityMonitoringRuleProposalInputType) GetAllowedValues() []SecurityMonitoringRuleProposalInputType {
	return allowedSecurityMonitoringRuleProposalInputTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityMonitoringRuleProposalInputType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringRuleProposalInputType(value)
	return nil
}

// NewSecurityMonitoringRuleProposalInputTypeFromValue returns a pointer to a valid SecurityMonitoringRuleProposalInputType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityMonitoringRuleProposalInputTypeFromValue(v string) (*SecurityMonitoringRuleProposalInputType, error) {
	ev := SecurityMonitoringRuleProposalInputType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringRuleProposalInputType: valid values are %v", v, allowedSecurityMonitoringRuleProposalInputTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityMonitoringRuleProposalInputType) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringRuleProposalInputTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringRuleProposalInputType value.
func (v SecurityMonitoringRuleProposalInputType) Ptr() *SecurityMonitoringRuleProposalInputType {
	return &v
}
