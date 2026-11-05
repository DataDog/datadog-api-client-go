// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityAgentInvestigationVerdict Verdict of the investigation. Only completed investigations have one.
type SecurityAgentInvestigationVerdict string

// List of SecurityAgentInvestigationVerdict.
const (
	SECURITYAGENTINVESTIGATIONVERDICT_BENIGN     SecurityAgentInvestigationVerdict = "benign"
	SECURITYAGENTINVESTIGATIONVERDICT_SUSPICIOUS SecurityAgentInvestigationVerdict = "suspicious"
	SECURITYAGENTINVESTIGATIONVERDICT_MALICIOUS  SecurityAgentInvestigationVerdict = "malicious"
)

var allowedSecurityAgentInvestigationVerdictEnumValues = []SecurityAgentInvestigationVerdict{
	SECURITYAGENTINVESTIGATIONVERDICT_BENIGN,
	SECURITYAGENTINVESTIGATIONVERDICT_SUSPICIOUS,
	SECURITYAGENTINVESTIGATIONVERDICT_MALICIOUS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityAgentInvestigationVerdict) GetAllowedValues() []SecurityAgentInvestigationVerdict {
	return allowedSecurityAgentInvestigationVerdictEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityAgentInvestigationVerdict) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityAgentInvestigationVerdict(value)
	return nil
}

// NewSecurityAgentInvestigationVerdictFromValue returns a pointer to a valid SecurityAgentInvestigationVerdict
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityAgentInvestigationVerdictFromValue(v string) (*SecurityAgentInvestigationVerdict, error) {
	ev := SecurityAgentInvestigationVerdict(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityAgentInvestigationVerdict: valid values are %v", v, allowedSecurityAgentInvestigationVerdictEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityAgentInvestigationVerdict) IsValid() bool {
	for _, existing := range allowedSecurityAgentInvestigationVerdictEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityAgentInvestigationVerdict value.
func (v SecurityAgentInvestigationVerdict) Ptr() *SecurityAgentInvestigationVerdict {
	return &v
}
