// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityAgentInvestigationNoRunReason Why `filter[signal_id]` returned no investigation. Set only on the first page of an empty result. The reason is
// based on the organization's current settings. Poll again only on `investigation_not_started`, after the number
// of seconds in the `Retry-After` header; treat the others as final for this signal.
//   - `investigation_not_started`: no investigation has started yet, but one may still start.
//   - `not_investigated`: no investigation started within an hour of the signal, so none is expected. The signal
//     was skipped for a reason this API doesn't report, such as its severity or a rate limit.
//   - `rule_not_enabled`: the signal's detection rule doesn't have Bits Security Analyst enabled, so no
//     investigation starts automatically.
//   - `signal_not_found`: the ID isn't a signal ID, or no signal with this ID exists in the organization.
//   - `org_not_activated`: the organization hasn't activated Bits Security Analyst.
type SecurityAgentInvestigationNoRunReason string

// List of SecurityAgentInvestigationNoRunReason.
const (
	SECURITYAGENTINVESTIGATIONNORUNREASON_INVESTIGATION_NOT_STARTED SecurityAgentInvestigationNoRunReason = "investigation_not_started"
	SECURITYAGENTINVESTIGATIONNORUNREASON_NOT_INVESTIGATED          SecurityAgentInvestigationNoRunReason = "not_investigated"
	SECURITYAGENTINVESTIGATIONNORUNREASON_RULE_NOT_ENABLED          SecurityAgentInvestigationNoRunReason = "rule_not_enabled"
	SECURITYAGENTINVESTIGATIONNORUNREASON_SIGNAL_NOT_FOUND          SecurityAgentInvestigationNoRunReason = "signal_not_found"
	SECURITYAGENTINVESTIGATIONNORUNREASON_ORG_NOT_ACTIVATED         SecurityAgentInvestigationNoRunReason = "org_not_activated"
)

var allowedSecurityAgentInvestigationNoRunReasonEnumValues = []SecurityAgentInvestigationNoRunReason{
	SECURITYAGENTINVESTIGATIONNORUNREASON_INVESTIGATION_NOT_STARTED,
	SECURITYAGENTINVESTIGATIONNORUNREASON_NOT_INVESTIGATED,
	SECURITYAGENTINVESTIGATIONNORUNREASON_RULE_NOT_ENABLED,
	SECURITYAGENTINVESTIGATIONNORUNREASON_SIGNAL_NOT_FOUND,
	SECURITYAGENTINVESTIGATIONNORUNREASON_ORG_NOT_ACTIVATED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityAgentInvestigationNoRunReason) GetAllowedValues() []SecurityAgentInvestigationNoRunReason {
	return allowedSecurityAgentInvestigationNoRunReasonEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityAgentInvestigationNoRunReason) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityAgentInvestigationNoRunReason(value)
	return nil
}

// NewSecurityAgentInvestigationNoRunReasonFromValue returns a pointer to a valid SecurityAgentInvestigationNoRunReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityAgentInvestigationNoRunReasonFromValue(v string) (*SecurityAgentInvestigationNoRunReason, error) {
	ev := SecurityAgentInvestigationNoRunReason(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityAgentInvestigationNoRunReason: valid values are %v", v, allowedSecurityAgentInvestigationNoRunReasonEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityAgentInvestigationNoRunReason) IsValid() bool {
	for _, existing := range allowedSecurityAgentInvestigationNoRunReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityAgentInvestigationNoRunReason value.
func (v SecurityAgentInvestigationNoRunReason) Ptr() *SecurityAgentInvestigationNoRunReason {
	return &v
}
