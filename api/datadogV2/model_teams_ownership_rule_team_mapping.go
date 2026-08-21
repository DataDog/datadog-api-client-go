// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamsOwnershipRuleTeamMapping An individual team's ownership entry within a teams ownership rule.
type TeamsOwnershipRuleTeamMapping struct {
	// The ID of the underlying mapping, used to delete this team's ownership individually.
	MappingId string `json:"mapping_id"`
	// The handle of the owning team.
	TeamHandle string `json:"team_handle"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTeamsOwnershipRuleTeamMapping instantiates a new TeamsOwnershipRuleTeamMapping object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTeamsOwnershipRuleTeamMapping(mappingId string, teamHandle string) *TeamsOwnershipRuleTeamMapping {
	this := TeamsOwnershipRuleTeamMapping{}
	this.MappingId = mappingId
	this.TeamHandle = teamHandle
	return &this
}

// NewTeamsOwnershipRuleTeamMappingWithDefaults instantiates a new TeamsOwnershipRuleTeamMapping object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTeamsOwnershipRuleTeamMappingWithDefaults() *TeamsOwnershipRuleTeamMapping {
	this := TeamsOwnershipRuleTeamMapping{}
	return &this
}

// GetMappingId returns the MappingId field value.
func (o *TeamsOwnershipRuleTeamMapping) GetMappingId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.MappingId
}

// GetMappingIdOk returns a tuple with the MappingId field value
// and a boolean to check if the value has been set.
func (o *TeamsOwnershipRuleTeamMapping) GetMappingIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MappingId, true
}

// SetMappingId sets field value.
func (o *TeamsOwnershipRuleTeamMapping) SetMappingId(v string) {
	o.MappingId = v
}

// GetTeamHandle returns the TeamHandle field value.
func (o *TeamsOwnershipRuleTeamMapping) GetTeamHandle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TeamHandle
}

// GetTeamHandleOk returns a tuple with the TeamHandle field value
// and a boolean to check if the value has been set.
func (o *TeamsOwnershipRuleTeamMapping) GetTeamHandleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TeamHandle, true
}

// SetTeamHandle sets field value.
func (o *TeamsOwnershipRuleTeamMapping) SetTeamHandle(v string) {
	o.TeamHandle = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TeamsOwnershipRuleTeamMapping) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["mapping_id"] = o.MappingId
	toSerialize["team_handle"] = o.TeamHandle

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TeamsOwnershipRuleTeamMapping) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		MappingId  *string `json:"mapping_id"`
		TeamHandle *string `json:"team_handle"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.MappingId == nil {
		return fmt.Errorf("required field mapping_id missing")
	}
	if all.TeamHandle == nil {
		return fmt.Errorf("required field team_handle missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"mapping_id", "team_handle"})
	} else {
		return err
	}
	o.MappingId = *all.MappingId
	o.TeamHandle = *all.TeamHandle

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
