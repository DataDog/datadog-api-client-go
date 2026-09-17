// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamsOwnershipMappingCreateDataAttributes The attributes of the teams ownership mapping to create.
type TeamsOwnershipMappingCreateDataAttributes struct {
	// The ID of the RUM application this mapping applies to.
	// For browser applications, provide the real application UUID — the team is applied to the view regardless of service.
	// For mobile applications, omit this field (or set it to the nil UUID `00000000-0000-0000-0000-000000000000`) — the team is applied to the view and service combination across all applications.
	ApplicationId *uuid.UUID `json:"application_id,omitempty"`
	// How the `view_name` is matched against RUM view names.
	MatchType *TeamsOwnershipMatchType `json:"match_type,omitempty"`
	// The RUM application's service name. For browser applications, this is optional. For mobile applications, this is required and scopes the ownership to a specific service.
	Service *string `json:"service,omitempty"`
	// The handle of the team that owns the matched RUM views.
	TeamHandle string `json:"team_handle"`
	// The RUM view name to match, or its prefix when `match_type` is `prefix`.
	ViewName string `json:"view_name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTeamsOwnershipMappingCreateDataAttributes instantiates a new TeamsOwnershipMappingCreateDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTeamsOwnershipMappingCreateDataAttributes(teamHandle string, viewName string) *TeamsOwnershipMappingCreateDataAttributes {
	this := TeamsOwnershipMappingCreateDataAttributes{}
	var matchType TeamsOwnershipMatchType = TEAMSOWNERSHIPMATCHTYPE_EXACT
	this.MatchType = &matchType
	this.TeamHandle = teamHandle
	this.ViewName = viewName
	return &this
}

// NewTeamsOwnershipMappingCreateDataAttributesWithDefaults instantiates a new TeamsOwnershipMappingCreateDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTeamsOwnershipMappingCreateDataAttributesWithDefaults() *TeamsOwnershipMappingCreateDataAttributes {
	this := TeamsOwnershipMappingCreateDataAttributes{}
	var matchType TeamsOwnershipMatchType = TEAMSOWNERSHIPMATCHTYPE_EXACT
	this.MatchType = &matchType
	return &this
}

// GetApplicationId returns the ApplicationId field value if set, zero value otherwise.
func (o *TeamsOwnershipMappingCreateDataAttributes) GetApplicationId() uuid.UUID {
	if o == nil || o.ApplicationId == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamsOwnershipMappingCreateDataAttributes) GetApplicationIdOk() (*uuid.UUID, bool) {
	if o == nil || o.ApplicationId == nil {
		return nil, false
	}
	return o.ApplicationId, true
}

// HasApplicationId returns a boolean if a field has been set.
func (o *TeamsOwnershipMappingCreateDataAttributes) HasApplicationId() bool {
	return o != nil && o.ApplicationId != nil
}

// SetApplicationId gets a reference to the given uuid.UUID and assigns it to the ApplicationId field.
func (o *TeamsOwnershipMappingCreateDataAttributes) SetApplicationId(v uuid.UUID) {
	o.ApplicationId = &v
}

// GetMatchType returns the MatchType field value if set, zero value otherwise.
func (o *TeamsOwnershipMappingCreateDataAttributes) GetMatchType() TeamsOwnershipMatchType {
	if o == nil || o.MatchType == nil {
		var ret TeamsOwnershipMatchType
		return ret
	}
	return *o.MatchType
}

// GetMatchTypeOk returns a tuple with the MatchType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamsOwnershipMappingCreateDataAttributes) GetMatchTypeOk() (*TeamsOwnershipMatchType, bool) {
	if o == nil || o.MatchType == nil {
		return nil, false
	}
	return o.MatchType, true
}

// HasMatchType returns a boolean if a field has been set.
func (o *TeamsOwnershipMappingCreateDataAttributes) HasMatchType() bool {
	return o != nil && o.MatchType != nil
}

// SetMatchType gets a reference to the given TeamsOwnershipMatchType and assigns it to the MatchType field.
func (o *TeamsOwnershipMappingCreateDataAttributes) SetMatchType(v TeamsOwnershipMatchType) {
	o.MatchType = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *TeamsOwnershipMappingCreateDataAttributes) GetService() string {
	if o == nil || o.Service == nil {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamsOwnershipMappingCreateDataAttributes) GetServiceOk() (*string, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *TeamsOwnershipMappingCreateDataAttributes) HasService() bool {
	return o != nil && o.Service != nil
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *TeamsOwnershipMappingCreateDataAttributes) SetService(v string) {
	o.Service = &v
}

// GetTeamHandle returns the TeamHandle field value.
func (o *TeamsOwnershipMappingCreateDataAttributes) GetTeamHandle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TeamHandle
}

// GetTeamHandleOk returns a tuple with the TeamHandle field value
// and a boolean to check if the value has been set.
func (o *TeamsOwnershipMappingCreateDataAttributes) GetTeamHandleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TeamHandle, true
}

// SetTeamHandle sets field value.
func (o *TeamsOwnershipMappingCreateDataAttributes) SetTeamHandle(v string) {
	o.TeamHandle = v
}

// GetViewName returns the ViewName field value.
func (o *TeamsOwnershipMappingCreateDataAttributes) GetViewName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ViewName
}

// GetViewNameOk returns a tuple with the ViewName field value
// and a boolean to check if the value has been set.
func (o *TeamsOwnershipMappingCreateDataAttributes) GetViewNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ViewName, true
}

// SetViewName sets field value.
func (o *TeamsOwnershipMappingCreateDataAttributes) SetViewName(v string) {
	o.ViewName = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TeamsOwnershipMappingCreateDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ApplicationId != nil {
		toSerialize["application_id"] = o.ApplicationId
	}
	if o.MatchType != nil {
		toSerialize["match_type"] = o.MatchType
	}
	if o.Service != nil {
		toSerialize["service"] = o.Service
	}
	toSerialize["team_handle"] = o.TeamHandle
	toSerialize["view_name"] = o.ViewName

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TeamsOwnershipMappingCreateDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ApplicationId *uuid.UUID               `json:"application_id,omitempty"`
		MatchType     *TeamsOwnershipMatchType `json:"match_type,omitempty"`
		Service       *string                  `json:"service,omitempty"`
		TeamHandle    *string                  `json:"team_handle"`
		ViewName      *string                  `json:"view_name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.TeamHandle == nil {
		return fmt.Errorf("required field team_handle missing")
	}
	if all.ViewName == nil {
		return fmt.Errorf("required field view_name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"application_id", "match_type", "service", "team_handle", "view_name"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ApplicationId = all.ApplicationId
	if all.MatchType != nil && !all.MatchType.IsValid() {
		hasInvalidField = true
	} else {
		o.MatchType = all.MatchType
	}
	o.Service = all.Service
	o.TeamHandle = *all.TeamHandle
	o.ViewName = *all.ViewName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
