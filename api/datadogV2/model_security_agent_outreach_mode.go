// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityAgentOutreachMode The organization's outreach mode. `off` disables user outreach. `manual` only sends outreach an analyst
// triggers. `automated` auto-sends automated interviews for rules with outreach enabled, to users whose email
// domain is allowed.
type SecurityAgentOutreachMode string

// List of SecurityAgentOutreachMode.
const (
	SECURITYAGENTOUTREACHMODE_OFF       SecurityAgentOutreachMode = "off"
	SECURITYAGENTOUTREACHMODE_MANUAL    SecurityAgentOutreachMode = "manual"
	SECURITYAGENTOUTREACHMODE_AUTOMATED SecurityAgentOutreachMode = "automated"
)

var allowedSecurityAgentOutreachModeEnumValues = []SecurityAgentOutreachMode{
	SECURITYAGENTOUTREACHMODE_OFF,
	SECURITYAGENTOUTREACHMODE_MANUAL,
	SECURITYAGENTOUTREACHMODE_AUTOMATED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityAgentOutreachMode) GetAllowedValues() []SecurityAgentOutreachMode {
	return allowedSecurityAgentOutreachModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityAgentOutreachMode) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityAgentOutreachMode(value)
	return nil
}

// NewSecurityAgentOutreachModeFromValue returns a pointer to a valid SecurityAgentOutreachMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityAgentOutreachModeFromValue(v string) (*SecurityAgentOutreachMode, error) {
	ev := SecurityAgentOutreachMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityAgentOutreachMode: valid values are %v", v, allowedSecurityAgentOutreachModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityAgentOutreachMode) IsValid() bool {
	for _, existing := range allowedSecurityAgentOutreachModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityAgentOutreachMode value.
func (v SecurityAgentOutreachMode) Ptr() *SecurityAgentOutreachMode {
	return &v
}
