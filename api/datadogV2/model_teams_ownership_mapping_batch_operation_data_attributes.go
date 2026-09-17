// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamsOwnershipMappingBatchOperationDataAttributes The attributes of the mapping to add. `team_handle` and `view_name` are required
// when `op` is `add`. At least one of `service` or `application_id` must be provided.
type TeamsOwnershipMappingBatchOperationDataAttributes struct {
	// The ID of the RUM application this mapping applies to.
	// For browser applications, provide the real application UUID — the team is applied to the view regardless of service.
	// For mobile applications, omit this field (or set it to the nil UUID `00000000-0000-0000-0000-000000000000`) — the team is applied to the view and service combination across all applications.
	ApplicationId *uuid.UUID `json:"application_id,omitempty"`
	// How the `view_name` is matched against RUM view names.
	MatchType *TeamsOwnershipMatchType `json:"match_type,omitempty"`
	// The RUM application's service name. For browser applications, this is optional. For mobile applications, this is required and scopes the ownership to a specific service.
	Service *string `json:"service,omitempty"`
	// The handle of the team that owns the matched RUM views.
	TeamHandle *string `json:"team_handle,omitempty"`
	// The RUM view name to match, or its prefix when `match_type` is `prefix`.
	ViewName *string `json:"view_name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTeamsOwnershipMappingBatchOperationDataAttributes instantiates a new TeamsOwnershipMappingBatchOperationDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTeamsOwnershipMappingBatchOperationDataAttributes() *TeamsOwnershipMappingBatchOperationDataAttributes {
	this := TeamsOwnershipMappingBatchOperationDataAttributes{}
	var matchType TeamsOwnershipMatchType = TEAMSOWNERSHIPMATCHTYPE_EXACT
	this.MatchType = &matchType
	return &this
}

// NewTeamsOwnershipMappingBatchOperationDataAttributesWithDefaults instantiates a new TeamsOwnershipMappingBatchOperationDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTeamsOwnershipMappingBatchOperationDataAttributesWithDefaults() *TeamsOwnershipMappingBatchOperationDataAttributes {
	this := TeamsOwnershipMappingBatchOperationDataAttributes{}
	var matchType TeamsOwnershipMatchType = TEAMSOWNERSHIPMATCHTYPE_EXACT
	this.MatchType = &matchType
	return &this
}

// GetApplicationId returns the ApplicationId field value if set, zero value otherwise.
func (o *TeamsOwnershipMappingBatchOperationDataAttributes) GetApplicationId() uuid.UUID {
	if o == nil || o.ApplicationId == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamsOwnershipMappingBatchOperationDataAttributes) GetApplicationIdOk() (*uuid.UUID, bool) {
	if o == nil || o.ApplicationId == nil {
		return nil, false
	}
	return o.ApplicationId, true
}

// HasApplicationId returns a boolean if a field has been set.
func (o *TeamsOwnershipMappingBatchOperationDataAttributes) HasApplicationId() bool {
	return o != nil && o.ApplicationId != nil
}

// SetApplicationId gets a reference to the given uuid.UUID and assigns it to the ApplicationId field.
func (o *TeamsOwnershipMappingBatchOperationDataAttributes) SetApplicationId(v uuid.UUID) {
	o.ApplicationId = &v
}

// GetMatchType returns the MatchType field value if set, zero value otherwise.
func (o *TeamsOwnershipMappingBatchOperationDataAttributes) GetMatchType() TeamsOwnershipMatchType {
	if o == nil || o.MatchType == nil {
		var ret TeamsOwnershipMatchType
		return ret
	}
	return *o.MatchType
}

// GetMatchTypeOk returns a tuple with the MatchType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamsOwnershipMappingBatchOperationDataAttributes) GetMatchTypeOk() (*TeamsOwnershipMatchType, bool) {
	if o == nil || o.MatchType == nil {
		return nil, false
	}
	return o.MatchType, true
}

// HasMatchType returns a boolean if a field has been set.
func (o *TeamsOwnershipMappingBatchOperationDataAttributes) HasMatchType() bool {
	return o != nil && o.MatchType != nil
}

// SetMatchType gets a reference to the given TeamsOwnershipMatchType and assigns it to the MatchType field.
func (o *TeamsOwnershipMappingBatchOperationDataAttributes) SetMatchType(v TeamsOwnershipMatchType) {
	o.MatchType = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *TeamsOwnershipMappingBatchOperationDataAttributes) GetService() string {
	if o == nil || o.Service == nil {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamsOwnershipMappingBatchOperationDataAttributes) GetServiceOk() (*string, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *TeamsOwnershipMappingBatchOperationDataAttributes) HasService() bool {
	return o != nil && o.Service != nil
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *TeamsOwnershipMappingBatchOperationDataAttributes) SetService(v string) {
	o.Service = &v
}

// GetTeamHandle returns the TeamHandle field value if set, zero value otherwise.
func (o *TeamsOwnershipMappingBatchOperationDataAttributes) GetTeamHandle() string {
	if o == nil || o.TeamHandle == nil {
		var ret string
		return ret
	}
	return *o.TeamHandle
}

// GetTeamHandleOk returns a tuple with the TeamHandle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamsOwnershipMappingBatchOperationDataAttributes) GetTeamHandleOk() (*string, bool) {
	if o == nil || o.TeamHandle == nil {
		return nil, false
	}
	return o.TeamHandle, true
}

// HasTeamHandle returns a boolean if a field has been set.
func (o *TeamsOwnershipMappingBatchOperationDataAttributes) HasTeamHandle() bool {
	return o != nil && o.TeamHandle != nil
}

// SetTeamHandle gets a reference to the given string and assigns it to the TeamHandle field.
func (o *TeamsOwnershipMappingBatchOperationDataAttributes) SetTeamHandle(v string) {
	o.TeamHandle = &v
}

// GetViewName returns the ViewName field value if set, zero value otherwise.
func (o *TeamsOwnershipMappingBatchOperationDataAttributes) GetViewName() string {
	if o == nil || o.ViewName == nil {
		var ret string
		return ret
	}
	return *o.ViewName
}

// GetViewNameOk returns a tuple with the ViewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamsOwnershipMappingBatchOperationDataAttributes) GetViewNameOk() (*string, bool) {
	if o == nil || o.ViewName == nil {
		return nil, false
	}
	return o.ViewName, true
}

// HasViewName returns a boolean if a field has been set.
func (o *TeamsOwnershipMappingBatchOperationDataAttributes) HasViewName() bool {
	return o != nil && o.ViewName != nil
}

// SetViewName gets a reference to the given string and assigns it to the ViewName field.
func (o *TeamsOwnershipMappingBatchOperationDataAttributes) SetViewName(v string) {
	o.ViewName = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TeamsOwnershipMappingBatchOperationDataAttributes) MarshalJSON() ([]byte, error) {
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
	if o.TeamHandle != nil {
		toSerialize["team_handle"] = o.TeamHandle
	}
	if o.ViewName != nil {
		toSerialize["view_name"] = o.ViewName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TeamsOwnershipMappingBatchOperationDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ApplicationId *uuid.UUID               `json:"application_id,omitempty"`
		MatchType     *TeamsOwnershipMatchType `json:"match_type,omitempty"`
		Service       *string                  `json:"service,omitempty"`
		TeamHandle    *string                  `json:"team_handle,omitempty"`
		ViewName      *string                  `json:"view_name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
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
	o.TeamHandle = all.TeamHandle
	o.ViewName = all.ViewName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
