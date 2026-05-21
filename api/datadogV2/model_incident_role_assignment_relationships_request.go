// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentRoleAssignmentRelationshipsRequest Relationships for creating a role assignment.
type IncidentRoleAssignmentRelationshipsRequest struct {
	// Relationship to a role.
	ReservedRole *IncidentRoleAssignmentRoleRelationship `json:"reserved_role,omitempty"`
	// Relationship to a responder.
	Responder IncidentRoleAssignmentResponderRelationship `json:"responder"`
	// Relationship to a role.
	UserDefinedRole *IncidentRoleAssignmentRoleRelationship `json:"user_defined_role,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentRoleAssignmentRelationshipsRequest instantiates a new IncidentRoleAssignmentRelationshipsRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentRoleAssignmentRelationshipsRequest(responder IncidentRoleAssignmentResponderRelationship) *IncidentRoleAssignmentRelationshipsRequest {
	this := IncidentRoleAssignmentRelationshipsRequest{}
	this.Responder = responder
	return &this
}

// NewIncidentRoleAssignmentRelationshipsRequestWithDefaults instantiates a new IncidentRoleAssignmentRelationshipsRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentRoleAssignmentRelationshipsRequestWithDefaults() *IncidentRoleAssignmentRelationshipsRequest {
	this := IncidentRoleAssignmentRelationshipsRequest{}
	return &this
}

// GetReservedRole returns the ReservedRole field value if set, zero value otherwise.
func (o *IncidentRoleAssignmentRelationshipsRequest) GetReservedRole() IncidentRoleAssignmentRoleRelationship {
	if o == nil || o.ReservedRole == nil {
		var ret IncidentRoleAssignmentRoleRelationship
		return ret
	}
	return *o.ReservedRole
}

// GetReservedRoleOk returns a tuple with the ReservedRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentRoleAssignmentRelationshipsRequest) GetReservedRoleOk() (*IncidentRoleAssignmentRoleRelationship, bool) {
	if o == nil || o.ReservedRole == nil {
		return nil, false
	}
	return o.ReservedRole, true
}

// HasReservedRole returns a boolean if a field has been set.
func (o *IncidentRoleAssignmentRelationshipsRequest) HasReservedRole() bool {
	return o != nil && o.ReservedRole != nil
}

// SetReservedRole gets a reference to the given IncidentRoleAssignmentRoleRelationship and assigns it to the ReservedRole field.
func (o *IncidentRoleAssignmentRelationshipsRequest) SetReservedRole(v IncidentRoleAssignmentRoleRelationship) {
	o.ReservedRole = &v
}

// GetResponder returns the Responder field value.
func (o *IncidentRoleAssignmentRelationshipsRequest) GetResponder() IncidentRoleAssignmentResponderRelationship {
	if o == nil {
		var ret IncidentRoleAssignmentResponderRelationship
		return ret
	}
	return o.Responder
}

// GetResponderOk returns a tuple with the Responder field value
// and a boolean to check if the value has been set.
func (o *IncidentRoleAssignmentRelationshipsRequest) GetResponderOk() (*IncidentRoleAssignmentResponderRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Responder, true
}

// SetResponder sets field value.
func (o *IncidentRoleAssignmentRelationshipsRequest) SetResponder(v IncidentRoleAssignmentResponderRelationship) {
	o.Responder = v
}

// GetUserDefinedRole returns the UserDefinedRole field value if set, zero value otherwise.
func (o *IncidentRoleAssignmentRelationshipsRequest) GetUserDefinedRole() IncidentRoleAssignmentRoleRelationship {
	if o == nil || o.UserDefinedRole == nil {
		var ret IncidentRoleAssignmentRoleRelationship
		return ret
	}
	return *o.UserDefinedRole
}

// GetUserDefinedRoleOk returns a tuple with the UserDefinedRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentRoleAssignmentRelationshipsRequest) GetUserDefinedRoleOk() (*IncidentRoleAssignmentRoleRelationship, bool) {
	if o == nil || o.UserDefinedRole == nil {
		return nil, false
	}
	return o.UserDefinedRole, true
}

// HasUserDefinedRole returns a boolean if a field has been set.
func (o *IncidentRoleAssignmentRelationshipsRequest) HasUserDefinedRole() bool {
	return o != nil && o.UserDefinedRole != nil
}

// SetUserDefinedRole gets a reference to the given IncidentRoleAssignmentRoleRelationship and assigns it to the UserDefinedRole field.
func (o *IncidentRoleAssignmentRelationshipsRequest) SetUserDefinedRole(v IncidentRoleAssignmentRoleRelationship) {
	o.UserDefinedRole = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentRoleAssignmentRelationshipsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ReservedRole != nil {
		toSerialize["reserved_role"] = o.ReservedRole
	}
	toSerialize["responder"] = o.Responder
	if o.UserDefinedRole != nil {
		toSerialize["user_defined_role"] = o.UserDefinedRole
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentRoleAssignmentRelationshipsRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ReservedRole    *IncidentRoleAssignmentRoleRelationship      `json:"reserved_role,omitempty"`
		Responder       *IncidentRoleAssignmentResponderRelationship `json:"responder"`
		UserDefinedRole *IncidentRoleAssignmentRoleRelationship      `json:"user_defined_role,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Responder == nil {
		return fmt.Errorf("required field responder missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"reserved_role", "responder", "user_defined_role"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.ReservedRole != nil && all.ReservedRole.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ReservedRole = all.ReservedRole
	if all.Responder.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Responder = *all.Responder
	if all.UserDefinedRole != nil && all.UserDefinedRole.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.UserDefinedRole = all.UserDefinedRole

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
