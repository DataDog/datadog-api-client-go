// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamsOwnershipMappingResponseAttributes The attributes of a teams ownership mapping.
type TeamsOwnershipMappingResponseAttributes struct {
	// The ID of the RUM application this mapping applies to.
	// For browser applications, this is the real application UUID.
	// For mobile applications, this is the nil UUID `00000000-0000-0000-0000-000000000000` (wildcard), meaning the ownership applies across all applications.
	ApplicationId string `json:"application_id"`
	// Timestamp when the mapping was created.
	CreatedAt time.Time `json:"created_at"`
	// The UUID of the user who created the mapping.
	CreatedBy string `json:"created_by"`
	// How the `view_name` is matched against RUM view names.
	MatchType TeamsOwnershipMatchType `json:"match_type"`
	// The ID of the organization that owns this mapping.
	OrgId int64 `json:"org_id"`
	// The RUM application's service name. For browser applications, may be empty. For mobile applications, this is the service that scopes the ownership.
	Service string `json:"service"`
	// The handle of the team that owns the matched RUM views.
	TeamHandle string `json:"team_handle"`
	// The RUM view name to match, or its prefix when `match_type` is `prefix`.
	ViewName string `json:"view_name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTeamsOwnershipMappingResponseAttributes instantiates a new TeamsOwnershipMappingResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTeamsOwnershipMappingResponseAttributes(applicationId string, createdAt time.Time, createdBy string, matchType TeamsOwnershipMatchType, orgId int64, service string, teamHandle string, viewName string) *TeamsOwnershipMappingResponseAttributes {
	this := TeamsOwnershipMappingResponseAttributes{}
	this.ApplicationId = applicationId
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.MatchType = matchType
	this.OrgId = orgId
	this.Service = service
	this.TeamHandle = teamHandle
	this.ViewName = viewName
	return &this
}

// NewTeamsOwnershipMappingResponseAttributesWithDefaults instantiates a new TeamsOwnershipMappingResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTeamsOwnershipMappingResponseAttributesWithDefaults() *TeamsOwnershipMappingResponseAttributes {
	this := TeamsOwnershipMappingResponseAttributes{}
	var matchType TeamsOwnershipMatchType = TEAMSOWNERSHIPMATCHTYPE_EXACT
	this.MatchType = matchType
	return &this
}

// GetApplicationId returns the ApplicationId field value.
func (o *TeamsOwnershipMappingResponseAttributes) GetApplicationId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value
// and a boolean to check if the value has been set.
func (o *TeamsOwnershipMappingResponseAttributes) GetApplicationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApplicationId, true
}

// SetApplicationId sets field value.
func (o *TeamsOwnershipMappingResponseAttributes) SetApplicationId(v string) {
	o.ApplicationId = v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *TeamsOwnershipMappingResponseAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *TeamsOwnershipMappingResponseAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *TeamsOwnershipMappingResponseAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value.
func (o *TeamsOwnershipMappingResponseAttributes) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *TeamsOwnershipMappingResponseAttributes) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value.
func (o *TeamsOwnershipMappingResponseAttributes) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetMatchType returns the MatchType field value.
func (o *TeamsOwnershipMappingResponseAttributes) GetMatchType() TeamsOwnershipMatchType {
	if o == nil {
		var ret TeamsOwnershipMatchType
		return ret
	}
	return o.MatchType
}

// GetMatchTypeOk returns a tuple with the MatchType field value
// and a boolean to check if the value has been set.
func (o *TeamsOwnershipMappingResponseAttributes) GetMatchTypeOk() (*TeamsOwnershipMatchType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MatchType, true
}

// SetMatchType sets field value.
func (o *TeamsOwnershipMappingResponseAttributes) SetMatchType(v TeamsOwnershipMatchType) {
	o.MatchType = v
}

// GetOrgId returns the OrgId field value.
func (o *TeamsOwnershipMappingResponseAttributes) GetOrgId() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *TeamsOwnershipMappingResponseAttributes) GetOrgIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value.
func (o *TeamsOwnershipMappingResponseAttributes) SetOrgId(v int64) {
	o.OrgId = v
}

// GetService returns the Service field value.
func (o *TeamsOwnershipMappingResponseAttributes) GetService() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Service
}

// GetServiceOk returns a tuple with the Service field value
// and a boolean to check if the value has been set.
func (o *TeamsOwnershipMappingResponseAttributes) GetServiceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Service, true
}

// SetService sets field value.
func (o *TeamsOwnershipMappingResponseAttributes) SetService(v string) {
	o.Service = v
}

// GetTeamHandle returns the TeamHandle field value.
func (o *TeamsOwnershipMappingResponseAttributes) GetTeamHandle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TeamHandle
}

// GetTeamHandleOk returns a tuple with the TeamHandle field value
// and a boolean to check if the value has been set.
func (o *TeamsOwnershipMappingResponseAttributes) GetTeamHandleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TeamHandle, true
}

// SetTeamHandle sets field value.
func (o *TeamsOwnershipMappingResponseAttributes) SetTeamHandle(v string) {
	o.TeamHandle = v
}

// GetViewName returns the ViewName field value.
func (o *TeamsOwnershipMappingResponseAttributes) GetViewName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ViewName
}

// GetViewNameOk returns a tuple with the ViewName field value
// and a boolean to check if the value has been set.
func (o *TeamsOwnershipMappingResponseAttributes) GetViewNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ViewName, true
}

// SetViewName sets field value.
func (o *TeamsOwnershipMappingResponseAttributes) SetViewName(v string) {
	o.ViewName = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TeamsOwnershipMappingResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["application_id"] = o.ApplicationId
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["created_by"] = o.CreatedBy
	toSerialize["match_type"] = o.MatchType
	toSerialize["org_id"] = o.OrgId
	toSerialize["service"] = o.Service
	toSerialize["team_handle"] = o.TeamHandle
	toSerialize["view_name"] = o.ViewName

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TeamsOwnershipMappingResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ApplicationId *string                  `json:"application_id"`
		CreatedAt     *time.Time               `json:"created_at"`
		CreatedBy     *string                  `json:"created_by"`
		MatchType     *TeamsOwnershipMatchType `json:"match_type"`
		OrgId         *int64                   `json:"org_id"`
		Service       *string                  `json:"service"`
		TeamHandle    *string                  `json:"team_handle"`
		ViewName      *string                  `json:"view_name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ApplicationId == nil {
		return fmt.Errorf("required field application_id missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.CreatedBy == nil {
		return fmt.Errorf("required field created_by missing")
	}
	if all.MatchType == nil {
		return fmt.Errorf("required field match_type missing")
	}
	if all.OrgId == nil {
		return fmt.Errorf("required field org_id missing")
	}
	if all.Service == nil {
		return fmt.Errorf("required field service missing")
	}
	if all.TeamHandle == nil {
		return fmt.Errorf("required field team_handle missing")
	}
	if all.ViewName == nil {
		return fmt.Errorf("required field view_name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"application_id", "created_at", "created_by", "match_type", "org_id", "service", "team_handle", "view_name"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ApplicationId = *all.ApplicationId
	o.CreatedAt = *all.CreatedAt
	o.CreatedBy = *all.CreatedBy
	if !all.MatchType.IsValid() {
		hasInvalidField = true
	} else {
		o.MatchType = *all.MatchType
	}
	o.OrgId = *all.OrgId
	o.Service = *all.Service
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
