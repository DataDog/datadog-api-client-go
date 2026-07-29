// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentResponderRelationships Relationships for an incident responder.
type IncidentResponderRelationships struct {
	// Relationship to user.
	CreatedBy *RelationshipToUser `json:"created_by,omitempty"`
	// Relationship to user.
	LastModifiedBy *RelationshipToUser `json:"last_modified_by,omitempty"`
	// Relationship to role assignments for a responder.
	RoleAssignments *IncidentResponderRoleAssignmentsRelationship `json:"role_assignments,omitempty"`
	// Relationship to user.
	User NullableNullableRelationshipToUser `json:"user,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentResponderRelationships instantiates a new IncidentResponderRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentResponderRelationships() *IncidentResponderRelationships {
	this := IncidentResponderRelationships{}
	return &this
}

// NewIncidentResponderRelationshipsWithDefaults instantiates a new IncidentResponderRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentResponderRelationshipsWithDefaults() *IncidentResponderRelationships {
	this := IncidentResponderRelationships{}
	return &this
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *IncidentResponderRelationships) GetCreatedBy() RelationshipToUser {
	if o == nil || o.CreatedBy == nil {
		var ret RelationshipToUser
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentResponderRelationships) GetCreatedByOk() (*RelationshipToUser, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *IncidentResponderRelationships) HasCreatedBy() bool {
	return o != nil && o.CreatedBy != nil
}

// SetCreatedBy gets a reference to the given RelationshipToUser and assigns it to the CreatedBy field.
func (o *IncidentResponderRelationships) SetCreatedBy(v RelationshipToUser) {
	o.CreatedBy = &v
}

// GetLastModifiedBy returns the LastModifiedBy field value if set, zero value otherwise.
func (o *IncidentResponderRelationships) GetLastModifiedBy() RelationshipToUser {
	if o == nil || o.LastModifiedBy == nil {
		var ret RelationshipToUser
		return ret
	}
	return *o.LastModifiedBy
}

// GetLastModifiedByOk returns a tuple with the LastModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentResponderRelationships) GetLastModifiedByOk() (*RelationshipToUser, bool) {
	if o == nil || o.LastModifiedBy == nil {
		return nil, false
	}
	return o.LastModifiedBy, true
}

// HasLastModifiedBy returns a boolean if a field has been set.
func (o *IncidentResponderRelationships) HasLastModifiedBy() bool {
	return o != nil && o.LastModifiedBy != nil
}

// SetLastModifiedBy gets a reference to the given RelationshipToUser and assigns it to the LastModifiedBy field.
func (o *IncidentResponderRelationships) SetLastModifiedBy(v RelationshipToUser) {
	o.LastModifiedBy = &v
}

// GetRoleAssignments returns the RoleAssignments field value if set, zero value otherwise.
func (o *IncidentResponderRelationships) GetRoleAssignments() IncidentResponderRoleAssignmentsRelationship {
	if o == nil || o.RoleAssignments == nil {
		var ret IncidentResponderRoleAssignmentsRelationship
		return ret
	}
	return *o.RoleAssignments
}

// GetRoleAssignmentsOk returns a tuple with the RoleAssignments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentResponderRelationships) GetRoleAssignmentsOk() (*IncidentResponderRoleAssignmentsRelationship, bool) {
	if o == nil || o.RoleAssignments == nil {
		return nil, false
	}
	return o.RoleAssignments, true
}

// HasRoleAssignments returns a boolean if a field has been set.
func (o *IncidentResponderRelationships) HasRoleAssignments() bool {
	return o != nil && o.RoleAssignments != nil
}

// SetRoleAssignments gets a reference to the given IncidentResponderRoleAssignmentsRelationship and assigns it to the RoleAssignments field.
func (o *IncidentResponderRelationships) SetRoleAssignments(v IncidentResponderRoleAssignmentsRelationship) {
	o.RoleAssignments = &v
}

// GetUser returns the User field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentResponderRelationships) GetUser() NullableRelationshipToUser {
	if o == nil || o.User.Get() == nil {
		var ret NullableRelationshipToUser
		return ret
	}
	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentResponderRelationships) GetUserOk() (*NullableRelationshipToUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// HasUser returns a boolean if a field has been set.
func (o *IncidentResponderRelationships) HasUser() bool {
	return o != nil && o.User.IsSet()
}

// SetUser gets a reference to the given NullableNullableRelationshipToUser and assigns it to the User field.
func (o *IncidentResponderRelationships) SetUser(v NullableRelationshipToUser) {
	o.User.Set(&v)
}

// SetUserNil sets the value for User to be an explicit nil.
func (o *IncidentResponderRelationships) SetUserNil() {
	o.User.Set(nil)
}

// UnsetUser ensures that no value is present for User, not even an explicit nil.
func (o *IncidentResponderRelationships) UnsetUser() {
	o.User.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentResponderRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedBy != nil {
		toSerialize["created_by"] = o.CreatedBy
	}
	if o.LastModifiedBy != nil {
		toSerialize["last_modified_by"] = o.LastModifiedBy
	}
	if o.RoleAssignments != nil {
		toSerialize["role_assignments"] = o.RoleAssignments
	}
	if o.User.IsSet() {
		toSerialize["user"] = o.User.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentResponderRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedBy       *RelationshipToUser                           `json:"created_by,omitempty"`
		LastModifiedBy  *RelationshipToUser                           `json:"last_modified_by,omitempty"`
		RoleAssignments *IncidentResponderRoleAssignmentsRelationship `json:"role_assignments,omitempty"`
		User            NullableNullableRelationshipToUser            `json:"user,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_by", "last_modified_by", "role_assignments", "user"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.CreatedBy != nil && all.CreatedBy.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CreatedBy = all.CreatedBy
	if all.LastModifiedBy != nil && all.LastModifiedBy.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.LastModifiedBy = all.LastModifiedBy
	if all.RoleAssignments != nil && all.RoleAssignments.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.RoleAssignments = all.RoleAssignments
	o.User = all.User

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
