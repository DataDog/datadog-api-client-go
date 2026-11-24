// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityAgentOutreachDomainAddType Type of an allowed outreach domain add request.
type SecurityAgentOutreachDomainAddType string

// List of SecurityAgentOutreachDomainAddType.
const (
	SECURITYAGENTOUTREACHDOMAINADDTYPE_OUTREACH_POLICY_DOMAIN_ADD SecurityAgentOutreachDomainAddType = "outreach_policy_domain_add"
)

var allowedSecurityAgentOutreachDomainAddTypeEnumValues = []SecurityAgentOutreachDomainAddType{
	SECURITYAGENTOUTREACHDOMAINADDTYPE_OUTREACH_POLICY_DOMAIN_ADD,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityAgentOutreachDomainAddType) GetAllowedValues() []SecurityAgentOutreachDomainAddType {
	return allowedSecurityAgentOutreachDomainAddTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityAgentOutreachDomainAddType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityAgentOutreachDomainAddType(value)
	return nil
}

// NewSecurityAgentOutreachDomainAddTypeFromValue returns a pointer to a valid SecurityAgentOutreachDomainAddType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityAgentOutreachDomainAddTypeFromValue(v string) (*SecurityAgentOutreachDomainAddType, error) {
	ev := SecurityAgentOutreachDomainAddType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityAgentOutreachDomainAddType: valid values are %v", v, allowedSecurityAgentOutreachDomainAddTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityAgentOutreachDomainAddType) IsValid() bool {
	for _, existing := range allowedSecurityAgentOutreachDomainAddTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityAgentOutreachDomainAddType value.
func (v SecurityAgentOutreachDomainAddType) Ptr() *SecurityAgentOutreachDomainAddType {
	return &v
}
