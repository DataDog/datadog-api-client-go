// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RoutingRuleRerouteToTeamActionType Indicates that the action reroutes the page to another team's routing rules.
type RoutingRuleRerouteToTeamActionType string

// List of RoutingRuleRerouteToTeamActionType.
const (
	ROUTINGRULEREROUTETOTEAMACTIONTYPE_REROUTE_TO_TEAM RoutingRuleRerouteToTeamActionType = "reroute_to_team"
)

var allowedRoutingRuleRerouteToTeamActionTypeEnumValues = []RoutingRuleRerouteToTeamActionType{
	ROUTINGRULEREROUTETOTEAMACTIONTYPE_REROUTE_TO_TEAM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RoutingRuleRerouteToTeamActionType) GetAllowedValues() []RoutingRuleRerouteToTeamActionType {
	return allowedRoutingRuleRerouteToTeamActionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RoutingRuleRerouteToTeamActionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RoutingRuleRerouteToTeamActionType(value)
	return nil
}

// NewRoutingRuleRerouteToTeamActionTypeFromValue returns a pointer to a valid RoutingRuleRerouteToTeamActionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRoutingRuleRerouteToTeamActionTypeFromValue(v string) (*RoutingRuleRerouteToTeamActionType, error) {
	ev := RoutingRuleRerouteToTeamActionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RoutingRuleRerouteToTeamActionType: valid values are %v", v, allowedRoutingRuleRerouteToTeamActionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RoutingRuleRerouteToTeamActionType) IsValid() bool {
	for _, existing := range allowedRoutingRuleRerouteToTeamActionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RoutingRuleRerouteToTeamActionType value.
func (v RoutingRuleRerouteToTeamActionType) Ptr() *RoutingRuleRerouteToTeamActionType {
	return &v
}
