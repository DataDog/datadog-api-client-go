// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentPostmortemRelationships The postmortem's relationships.
type IncidentPostmortemRelationships struct {
	// Relationship to incident.
	Incident *RelationshipToIncident `json:"incident,omitempty"`
	// Relationship to user.
	LastModifiedByUser *RelationshipToUser `json:"last_modified_by_user,omitempty"`
	// A relationship reference for a single incident responder.
	PostmortemOwnerResponder NullableRelationshipToIncidentResponder `json:"postmortem_owner_responder,omitempty"`
	// Relationship to user.
	PostmortemOwnerUser NullableNullableRelationshipToUser `json:"postmortem_owner_user,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentPostmortemRelationships instantiates a new IncidentPostmortemRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentPostmortemRelationships() *IncidentPostmortemRelationships {
	this := IncidentPostmortemRelationships{}
	return &this
}

// NewIncidentPostmortemRelationshipsWithDefaults instantiates a new IncidentPostmortemRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentPostmortemRelationshipsWithDefaults() *IncidentPostmortemRelationships {
	this := IncidentPostmortemRelationships{}
	return &this
}

// GetIncident returns the Incident field value if set, zero value otherwise.
func (o *IncidentPostmortemRelationships) GetIncident() RelationshipToIncident {
	if o == nil || o.Incident == nil {
		var ret RelationshipToIncident
		return ret
	}
	return *o.Incident
}

// GetIncidentOk returns a tuple with the Incident field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentPostmortemRelationships) GetIncidentOk() (*RelationshipToIncident, bool) {
	if o == nil || o.Incident == nil {
		return nil, false
	}
	return o.Incident, true
}

// HasIncident returns a boolean if a field has been set.
func (o *IncidentPostmortemRelationships) HasIncident() bool {
	return o != nil && o.Incident != nil
}

// SetIncident gets a reference to the given RelationshipToIncident and assigns it to the Incident field.
func (o *IncidentPostmortemRelationships) SetIncident(v RelationshipToIncident) {
	o.Incident = &v
}

// GetLastModifiedByUser returns the LastModifiedByUser field value if set, zero value otherwise.
func (o *IncidentPostmortemRelationships) GetLastModifiedByUser() RelationshipToUser {
	if o == nil || o.LastModifiedByUser == nil {
		var ret RelationshipToUser
		return ret
	}
	return *o.LastModifiedByUser
}

// GetLastModifiedByUserOk returns a tuple with the LastModifiedByUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentPostmortemRelationships) GetLastModifiedByUserOk() (*RelationshipToUser, bool) {
	if o == nil || o.LastModifiedByUser == nil {
		return nil, false
	}
	return o.LastModifiedByUser, true
}

// HasLastModifiedByUser returns a boolean if a field has been set.
func (o *IncidentPostmortemRelationships) HasLastModifiedByUser() bool {
	return o != nil && o.LastModifiedByUser != nil
}

// SetLastModifiedByUser gets a reference to the given RelationshipToUser and assigns it to the LastModifiedByUser field.
func (o *IncidentPostmortemRelationships) SetLastModifiedByUser(v RelationshipToUser) {
	o.LastModifiedByUser = &v
}

// GetPostmortemOwnerResponder returns the PostmortemOwnerResponder field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentPostmortemRelationships) GetPostmortemOwnerResponder() RelationshipToIncidentResponder {
	if o == nil || o.PostmortemOwnerResponder.Get() == nil {
		var ret RelationshipToIncidentResponder
		return ret
	}
	return *o.PostmortemOwnerResponder.Get()
}

// GetPostmortemOwnerResponderOk returns a tuple with the PostmortemOwnerResponder field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentPostmortemRelationships) GetPostmortemOwnerResponderOk() (*RelationshipToIncidentResponder, bool) {
	if o == nil {
		return nil, false
	}
	return o.PostmortemOwnerResponder.Get(), o.PostmortemOwnerResponder.IsSet()
}

// HasPostmortemOwnerResponder returns a boolean if a field has been set.
func (o *IncidentPostmortemRelationships) HasPostmortemOwnerResponder() bool {
	return o != nil && o.PostmortemOwnerResponder.IsSet()
}

// SetPostmortemOwnerResponder gets a reference to the given NullableRelationshipToIncidentResponder and assigns it to the PostmortemOwnerResponder field.
func (o *IncidentPostmortemRelationships) SetPostmortemOwnerResponder(v RelationshipToIncidentResponder) {
	o.PostmortemOwnerResponder.Set(&v)
}

// SetPostmortemOwnerResponderNil sets the value for PostmortemOwnerResponder to be an explicit nil.
func (o *IncidentPostmortemRelationships) SetPostmortemOwnerResponderNil() {
	o.PostmortemOwnerResponder.Set(nil)
}

// UnsetPostmortemOwnerResponder ensures that no value is present for PostmortemOwnerResponder, not even an explicit nil.
func (o *IncidentPostmortemRelationships) UnsetPostmortemOwnerResponder() {
	o.PostmortemOwnerResponder.Unset()
}

// GetPostmortemOwnerUser returns the PostmortemOwnerUser field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentPostmortemRelationships) GetPostmortemOwnerUser() NullableRelationshipToUser {
	if o == nil || o.PostmortemOwnerUser.Get() == nil {
		var ret NullableRelationshipToUser
		return ret
	}
	return *o.PostmortemOwnerUser.Get()
}

// GetPostmortemOwnerUserOk returns a tuple with the PostmortemOwnerUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentPostmortemRelationships) GetPostmortemOwnerUserOk() (*NullableRelationshipToUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.PostmortemOwnerUser.Get(), o.PostmortemOwnerUser.IsSet()
}

// HasPostmortemOwnerUser returns a boolean if a field has been set.
func (o *IncidentPostmortemRelationships) HasPostmortemOwnerUser() bool {
	return o != nil && o.PostmortemOwnerUser.IsSet()
}

// SetPostmortemOwnerUser gets a reference to the given NullableNullableRelationshipToUser and assigns it to the PostmortemOwnerUser field.
func (o *IncidentPostmortemRelationships) SetPostmortemOwnerUser(v NullableRelationshipToUser) {
	o.PostmortemOwnerUser.Set(&v)
}

// SetPostmortemOwnerUserNil sets the value for PostmortemOwnerUser to be an explicit nil.
func (o *IncidentPostmortemRelationships) SetPostmortemOwnerUserNil() {
	o.PostmortemOwnerUser.Set(nil)
}

// UnsetPostmortemOwnerUser ensures that no value is present for PostmortemOwnerUser, not even an explicit nil.
func (o *IncidentPostmortemRelationships) UnsetPostmortemOwnerUser() {
	o.PostmortemOwnerUser.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentPostmortemRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Incident != nil {
		toSerialize["incident"] = o.Incident
	}
	if o.LastModifiedByUser != nil {
		toSerialize["last_modified_by_user"] = o.LastModifiedByUser
	}
	if o.PostmortemOwnerResponder.IsSet() {
		toSerialize["postmortem_owner_responder"] = o.PostmortemOwnerResponder.Get()
	}
	if o.PostmortemOwnerUser.IsSet() {
		toSerialize["postmortem_owner_user"] = o.PostmortemOwnerUser.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentPostmortemRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Incident                 *RelationshipToIncident                 `json:"incident,omitempty"`
		LastModifiedByUser       *RelationshipToUser                     `json:"last_modified_by_user,omitempty"`
		PostmortemOwnerResponder NullableRelationshipToIncidentResponder `json:"postmortem_owner_responder,omitempty"`
		PostmortemOwnerUser      NullableNullableRelationshipToUser      `json:"postmortem_owner_user,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"incident", "last_modified_by_user", "postmortem_owner_responder", "postmortem_owner_user"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Incident != nil && all.Incident.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Incident = all.Incident
	if all.LastModifiedByUser != nil && all.LastModifiedByUser.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.LastModifiedByUser = all.LastModifiedByUser
	o.PostmortemOwnerResponder = all.PostmortemOwnerResponder
	o.PostmortemOwnerUser = all.PostmortemOwnerUser

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
