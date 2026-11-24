// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityAgentOutreachPolicyType Type of the outreach policy resource.
type SecurityAgentOutreachPolicyType string

// List of SecurityAgentOutreachPolicyType.
const (
	SECURITYAGENTOUTREACHPOLICYTYPE_OUTREACH_POLICY SecurityAgentOutreachPolicyType = "outreach_policy"
)

var allowedSecurityAgentOutreachPolicyTypeEnumValues = []SecurityAgentOutreachPolicyType{
	SECURITYAGENTOUTREACHPOLICYTYPE_OUTREACH_POLICY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityAgentOutreachPolicyType) GetAllowedValues() []SecurityAgentOutreachPolicyType {
	return allowedSecurityAgentOutreachPolicyTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityAgentOutreachPolicyType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityAgentOutreachPolicyType(value)
	return nil
}

// NewSecurityAgentOutreachPolicyTypeFromValue returns a pointer to a valid SecurityAgentOutreachPolicyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityAgentOutreachPolicyTypeFromValue(v string) (*SecurityAgentOutreachPolicyType, error) {
	ev := SecurityAgentOutreachPolicyType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityAgentOutreachPolicyType: valid values are %v", v, allowedSecurityAgentOutreachPolicyTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityAgentOutreachPolicyType) IsValid() bool {
	for _, existing := range allowedSecurityAgentOutreachPolicyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityAgentOutreachPolicyType value.
func (v SecurityAgentOutreachPolicyType) Ptr() *SecurityAgentOutreachPolicyType {
	return &v
}
