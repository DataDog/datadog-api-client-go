// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RoutingRuleRerouteToTeamAction Reroutes the page to another team, which then evaluates it against its own routing rules. Each routing rule can include this action only once. It can be combined only with `send_slack_message` and `send_teams_message` actions. It can't be used with `escalation_policy` or `workflow` actions, or when the routing rule sets `policy_id`.
type RoutingRuleRerouteToTeamAction struct {
	// The ID of the team to reroute the page to.
	DestinationTeamId uuid.UUID `json:"destination_team_id"`
	// Indicates that the action reroutes the page to another team's routing rules.
	Type RoutingRuleRerouteToTeamActionType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRoutingRuleRerouteToTeamAction instantiates a new RoutingRuleRerouteToTeamAction object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRoutingRuleRerouteToTeamAction(destinationTeamId uuid.UUID, typeVar RoutingRuleRerouteToTeamActionType) *RoutingRuleRerouteToTeamAction {
	this := RoutingRuleRerouteToTeamAction{}
	this.DestinationTeamId = destinationTeamId
	this.Type = typeVar
	return &this
}

// NewRoutingRuleRerouteToTeamActionWithDefaults instantiates a new RoutingRuleRerouteToTeamAction object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRoutingRuleRerouteToTeamActionWithDefaults() *RoutingRuleRerouteToTeamAction {
	this := RoutingRuleRerouteToTeamAction{}
	var typeVar RoutingRuleRerouteToTeamActionType = ROUTINGRULEREROUTETOTEAMACTIONTYPE_REROUTE_TO_TEAM
	this.Type = typeVar
	return &this
}

// GetDestinationTeamId returns the DestinationTeamId field value.
func (o *RoutingRuleRerouteToTeamAction) GetDestinationTeamId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.DestinationTeamId
}

// GetDestinationTeamIdOk returns a tuple with the DestinationTeamId field value
// and a boolean to check if the value has been set.
func (o *RoutingRuleRerouteToTeamAction) GetDestinationTeamIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DestinationTeamId, true
}

// SetDestinationTeamId sets field value.
func (o *RoutingRuleRerouteToTeamAction) SetDestinationTeamId(v uuid.UUID) {
	o.DestinationTeamId = v
}

// GetType returns the Type field value.
func (o *RoutingRuleRerouteToTeamAction) GetType() RoutingRuleRerouteToTeamActionType {
	if o == nil {
		var ret RoutingRuleRerouteToTeamActionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RoutingRuleRerouteToTeamAction) GetTypeOk() (*RoutingRuleRerouteToTeamActionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *RoutingRuleRerouteToTeamAction) SetType(v RoutingRuleRerouteToTeamActionType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RoutingRuleRerouteToTeamAction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["destination_team_id"] = o.DestinationTeamId
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RoutingRuleRerouteToTeamAction) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DestinationTeamId *uuid.UUID                          `json:"destination_team_id"`
		Type              *RoutingRuleRerouteToTeamActionType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DestinationTeamId == nil {
		return fmt.Errorf("required field destination_team_id missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"destination_team_id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.DestinationTeamId = *all.DestinationTeamId
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
