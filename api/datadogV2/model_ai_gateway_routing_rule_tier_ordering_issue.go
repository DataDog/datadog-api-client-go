// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayRoutingRuleTierOrderingIssue Stored tier ordering is inconsistent. mixed means targets have different supported orderings, and unknown means a target has an unsupported ordering. The tier is reported as configured without weights. Replace the rule to repair it. Omitted for consistent tiers. A forced deletion cannot rewrite a surviving tier with an unknown issue.
type AiGatewayRoutingRuleTierOrderingIssue string

// List of AiGatewayRoutingRuleTierOrderingIssue.
const (
	AIGATEWAYROUTINGRULETIERORDERINGISSUE_MIXED   AiGatewayRoutingRuleTierOrderingIssue = "mixed"
	AIGATEWAYROUTINGRULETIERORDERINGISSUE_UNKNOWN AiGatewayRoutingRuleTierOrderingIssue = "unknown"
)

var allowedAiGatewayRoutingRuleTierOrderingIssueEnumValues = []AiGatewayRoutingRuleTierOrderingIssue{
	AIGATEWAYROUTINGRULETIERORDERINGISSUE_MIXED,
	AIGATEWAYROUTINGRULETIERORDERINGISSUE_UNKNOWN,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AiGatewayRoutingRuleTierOrderingIssue) GetAllowedValues() []AiGatewayRoutingRuleTierOrderingIssue {
	return allowedAiGatewayRoutingRuleTierOrderingIssueEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiGatewayRoutingRuleTierOrderingIssue) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiGatewayRoutingRuleTierOrderingIssue(value)
	return nil
}

// NewAiGatewayRoutingRuleTierOrderingIssueFromValue returns a pointer to a valid AiGatewayRoutingRuleTierOrderingIssue
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiGatewayRoutingRuleTierOrderingIssueFromValue(v string) (*AiGatewayRoutingRuleTierOrderingIssue, error) {
	ev := AiGatewayRoutingRuleTierOrderingIssue(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiGatewayRoutingRuleTierOrderingIssue: valid values are %v", v, allowedAiGatewayRoutingRuleTierOrderingIssueEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiGatewayRoutingRuleTierOrderingIssue) IsValid() bool {
	for _, existing := range allowedAiGatewayRoutingRuleTierOrderingIssueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiGatewayRoutingRuleTierOrderingIssue value.
func (v AiGatewayRoutingRuleTierOrderingIssue) Ptr() *AiGatewayRoutingRuleTierOrderingIssue {
	return &v
}
