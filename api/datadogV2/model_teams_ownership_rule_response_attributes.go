// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamsOwnershipRuleResponseAttributes The attributes of a teams ownership rule.
type TeamsOwnershipRuleResponseAttributes struct {
	// The ID of the RUM application this mapping applies to.
	// For browser applications, this is the real application UUID.
	// For mobile applications, this is the nil UUID `00000000-0000-0000-0000-000000000000` (wildcard), meaning the ownership applies across all applications.
	ApplicationId string `json:"application_id"`
	// How the `view_name` is matched against RUM view names.
	MatchType TeamsOwnershipMatchType `json:"match_type"`
	// The RUM application's service name. For browser applications, may be empty. For mobile applications, this is the service that scopes the ownership.
	Service string `json:"service"`
	// The teams that own the matched views, each paired with the ID of its underlying mapping.
	Teams []TeamsOwnershipRuleTeamMapping `json:"teams"`
	// The RUM view name to match, or its prefix when `match_type` is `prefix`.
	ViewName string `json:"view_name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTeamsOwnershipRuleResponseAttributes instantiates a new TeamsOwnershipRuleResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTeamsOwnershipRuleResponseAttributes(applicationId string, matchType TeamsOwnershipMatchType, service string, teams []TeamsOwnershipRuleTeamMapping, viewName string) *TeamsOwnershipRuleResponseAttributes {
	this := TeamsOwnershipRuleResponseAttributes{}
	this.ApplicationId = applicationId
	this.MatchType = matchType
	this.Service = service
	this.Teams = teams
	this.ViewName = viewName
	return &this
}

// NewTeamsOwnershipRuleResponseAttributesWithDefaults instantiates a new TeamsOwnershipRuleResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTeamsOwnershipRuleResponseAttributesWithDefaults() *TeamsOwnershipRuleResponseAttributes {
	this := TeamsOwnershipRuleResponseAttributes{}
	var matchType TeamsOwnershipMatchType = TEAMSOWNERSHIPMATCHTYPE_EXACT
	this.MatchType = matchType
	return &this
}

// GetApplicationId returns the ApplicationId field value.
func (o *TeamsOwnershipRuleResponseAttributes) GetApplicationId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value
// and a boolean to check if the value has been set.
func (o *TeamsOwnershipRuleResponseAttributes) GetApplicationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApplicationId, true
}

// SetApplicationId sets field value.
func (o *TeamsOwnershipRuleResponseAttributes) SetApplicationId(v string) {
	o.ApplicationId = v
}

// GetMatchType returns the MatchType field value.
func (o *TeamsOwnershipRuleResponseAttributes) GetMatchType() TeamsOwnershipMatchType {
	if o == nil {
		var ret TeamsOwnershipMatchType
		return ret
	}
	return o.MatchType
}

// GetMatchTypeOk returns a tuple with the MatchType field value
// and a boolean to check if the value has been set.
func (o *TeamsOwnershipRuleResponseAttributes) GetMatchTypeOk() (*TeamsOwnershipMatchType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MatchType, true
}

// SetMatchType sets field value.
func (o *TeamsOwnershipRuleResponseAttributes) SetMatchType(v TeamsOwnershipMatchType) {
	o.MatchType = v
}

// GetService returns the Service field value.
func (o *TeamsOwnershipRuleResponseAttributes) GetService() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Service
}

// GetServiceOk returns a tuple with the Service field value
// and a boolean to check if the value has been set.
func (o *TeamsOwnershipRuleResponseAttributes) GetServiceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Service, true
}

// SetService sets field value.
func (o *TeamsOwnershipRuleResponseAttributes) SetService(v string) {
	o.Service = v
}

// GetTeams returns the Teams field value.
func (o *TeamsOwnershipRuleResponseAttributes) GetTeams() []TeamsOwnershipRuleTeamMapping {
	if o == nil {
		var ret []TeamsOwnershipRuleTeamMapping
		return ret
	}
	return o.Teams
}

// GetTeamsOk returns a tuple with the Teams field value
// and a boolean to check if the value has been set.
func (o *TeamsOwnershipRuleResponseAttributes) GetTeamsOk() (*[]TeamsOwnershipRuleTeamMapping, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Teams, true
}

// SetTeams sets field value.
func (o *TeamsOwnershipRuleResponseAttributes) SetTeams(v []TeamsOwnershipRuleTeamMapping) {
	o.Teams = v
}

// GetViewName returns the ViewName field value.
func (o *TeamsOwnershipRuleResponseAttributes) GetViewName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ViewName
}

// GetViewNameOk returns a tuple with the ViewName field value
// and a boolean to check if the value has been set.
func (o *TeamsOwnershipRuleResponseAttributes) GetViewNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ViewName, true
}

// SetViewName sets field value.
func (o *TeamsOwnershipRuleResponseAttributes) SetViewName(v string) {
	o.ViewName = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TeamsOwnershipRuleResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["application_id"] = o.ApplicationId
	toSerialize["match_type"] = o.MatchType
	toSerialize["service"] = o.Service
	toSerialize["teams"] = o.Teams
	toSerialize["view_name"] = o.ViewName

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TeamsOwnershipRuleResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ApplicationId *string                          `json:"application_id"`
		MatchType     *TeamsOwnershipMatchType         `json:"match_type"`
		Service       *string                          `json:"service"`
		Teams         *[]TeamsOwnershipRuleTeamMapping `json:"teams"`
		ViewName      *string                          `json:"view_name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ApplicationId == nil {
		return fmt.Errorf("required field application_id missing")
	}
	if all.MatchType == nil {
		return fmt.Errorf("required field match_type missing")
	}
	if all.Service == nil {
		return fmt.Errorf("required field service missing")
	}
	if all.Teams == nil {
		return fmt.Errorf("required field teams missing")
	}
	if all.ViewName == nil {
		return fmt.Errorf("required field view_name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"application_id", "match_type", "service", "teams", "view_name"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ApplicationId = *all.ApplicationId
	if !all.MatchType.IsValid() {
		hasInvalidField = true
	} else {
		o.MatchType = *all.MatchType
	}
	o.Service = *all.Service
	o.Teams = *all.Teams
	o.ViewName = *all.ViewName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
