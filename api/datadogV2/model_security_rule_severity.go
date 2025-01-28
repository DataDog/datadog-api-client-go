// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityRuleSeverity Severity of a security rule
type SecurityRuleSeverity string

// List of SecurityRuleSeverity.
const (
	SECURITYRULESEVERITY_CRITICAL SecurityRuleSeverity = "critical"
	SECURITYRULESEVERITY_HIGH     SecurityRuleSeverity = "high"
	SECURITYRULESEVERITY_MEDIUM   SecurityRuleSeverity = "medium"
	SECURITYRULESEVERITY_LOW      SecurityRuleSeverity = "low"
	SECURITYRULESEVERITY_UNKNOWN  SecurityRuleSeverity = "unknown"
	SECURITYRULESEVERITY_INFO     SecurityRuleSeverity = "info"
)

var allowedSecurityRuleSeverityEnumValues = []SecurityRuleSeverity{
	SECURITYRULESEVERITY_CRITICAL,
	SECURITYRULESEVERITY_HIGH,
	SECURITYRULESEVERITY_MEDIUM,
	SECURITYRULESEVERITY_LOW,
	SECURITYRULESEVERITY_UNKNOWN,
	SECURITYRULESEVERITY_INFO,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityRuleSeverity) GetAllowedValues() []SecurityRuleSeverity {
	return allowedSecurityRuleSeverityEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityRuleSeverity) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityRuleSeverity(value)
	return nil
}

// NewSecurityRuleSeverityFromValue returns a pointer to a valid SecurityRuleSeverity
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityRuleSeverityFromValue(v string) (*SecurityRuleSeverity, error) {
	ev := SecurityRuleSeverity(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityRuleSeverity: valid values are %v", v, allowedSecurityRuleSeverityEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityRuleSeverity) IsValid() bool {
	for _, existing := range allowedSecurityRuleSeverityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityRuleSeverity value.
func (v SecurityRuleSeverity) Ptr() *SecurityRuleSeverity {
	return &v
}
