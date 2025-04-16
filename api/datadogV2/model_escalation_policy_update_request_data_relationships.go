// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EscalationPolicyUpdateRequestDataRelationships Represents relationships in an escalation policy update request, including references to teams.
type EscalationPolicyUpdateRequestDataRelationships struct {
	// Defines the relationship to teams within an escalation policy update request, referencing the teams to be associated with or removed from the policy.
	Teams *EscalationPolicyUpdateRequestDataRelationshipsTeams `json:"teams,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEscalationPolicyUpdateRequestDataRelationships instantiates a new EscalationPolicyUpdateRequestDataRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEscalationPolicyUpdateRequestDataRelationships() *EscalationPolicyUpdateRequestDataRelationships {
	this := EscalationPolicyUpdateRequestDataRelationships{}
	return &this
}

// NewEscalationPolicyUpdateRequestDataRelationshipsWithDefaults instantiates a new EscalationPolicyUpdateRequestDataRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEscalationPolicyUpdateRequestDataRelationshipsWithDefaults() *EscalationPolicyUpdateRequestDataRelationships {
	this := EscalationPolicyUpdateRequestDataRelationships{}
	return &this
}

// GetTeams returns the Teams field value if set, zero value otherwise.
func (o *EscalationPolicyUpdateRequestDataRelationships) GetTeams() EscalationPolicyUpdateRequestDataRelationshipsTeams {
	if o == nil || o.Teams == nil {
		var ret EscalationPolicyUpdateRequestDataRelationshipsTeams
		return ret
	}
	return *o.Teams
}

// GetTeamsOk returns a tuple with the Teams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EscalationPolicyUpdateRequestDataRelationships) GetTeamsOk() (*EscalationPolicyUpdateRequestDataRelationshipsTeams, bool) {
	if o == nil || o.Teams == nil {
		return nil, false
	}
	return o.Teams, true
}

// HasTeams returns a boolean if a field has been set.
func (o *EscalationPolicyUpdateRequestDataRelationships) HasTeams() bool {
	return o != nil && o.Teams != nil
}

// SetTeams gets a reference to the given EscalationPolicyUpdateRequestDataRelationshipsTeams and assigns it to the Teams field.
func (o *EscalationPolicyUpdateRequestDataRelationships) SetTeams(v EscalationPolicyUpdateRequestDataRelationshipsTeams) {
	o.Teams = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EscalationPolicyUpdateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Teams != nil {
		toSerialize["teams"] = o.Teams
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EscalationPolicyUpdateRequestDataRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Teams *EscalationPolicyUpdateRequestDataRelationshipsTeams `json:"teams,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"teams"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Teams != nil && all.Teams.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Teams = all.Teams

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
