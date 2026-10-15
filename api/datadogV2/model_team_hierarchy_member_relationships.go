// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamHierarchyMemberRelationships User and requested team linkage.
type TeamHierarchyMemberRelationships struct {
	// Linkage to the requested team.
	Team TeamHierarchyMemberTeamRelationship `json:"team"`
	// Linkage to the user.
	User TeamHierarchyMemberUserRelationship `json:"user"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTeamHierarchyMemberRelationships instantiates a new TeamHierarchyMemberRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTeamHierarchyMemberRelationships(team TeamHierarchyMemberTeamRelationship, user TeamHierarchyMemberUserRelationship) *TeamHierarchyMemberRelationships {
	this := TeamHierarchyMemberRelationships{}
	this.Team = team
	this.User = user
	return &this
}

// NewTeamHierarchyMemberRelationshipsWithDefaults instantiates a new TeamHierarchyMemberRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTeamHierarchyMemberRelationshipsWithDefaults() *TeamHierarchyMemberRelationships {
	this := TeamHierarchyMemberRelationships{}
	return &this
}

// GetTeam returns the Team field value.
func (o *TeamHierarchyMemberRelationships) GetTeam() TeamHierarchyMemberTeamRelationship {
	if o == nil {
		var ret TeamHierarchyMemberTeamRelationship
		return ret
	}
	return o.Team
}

// GetTeamOk returns a tuple with the Team field value
// and a boolean to check if the value has been set.
func (o *TeamHierarchyMemberRelationships) GetTeamOk() (*TeamHierarchyMemberTeamRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Team, true
}

// SetTeam sets field value.
func (o *TeamHierarchyMemberRelationships) SetTeam(v TeamHierarchyMemberTeamRelationship) {
	o.Team = v
}

// GetUser returns the User field value.
func (o *TeamHierarchyMemberRelationships) GetUser() TeamHierarchyMemberUserRelationship {
	if o == nil {
		var ret TeamHierarchyMemberUserRelationship
		return ret
	}
	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *TeamHierarchyMemberRelationships) GetUserOk() (*TeamHierarchyMemberUserRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value.
func (o *TeamHierarchyMemberRelationships) SetUser(v TeamHierarchyMemberUserRelationship) {
	o.User = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TeamHierarchyMemberRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["team"] = o.Team
	toSerialize["user"] = o.User

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TeamHierarchyMemberRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Team *TeamHierarchyMemberTeamRelationship `json:"team"`
		User *TeamHierarchyMemberUserRelationship `json:"user"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Team == nil {
		return fmt.Errorf("required field team missing")
	}
	if all.User == nil {
		return fmt.Errorf("required field user missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"team", "user"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Team.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Team = *all.Team
	if all.User.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.User = *all.User

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
