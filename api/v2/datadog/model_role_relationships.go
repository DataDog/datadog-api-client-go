/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
)

// RoleRelationships Relationships of the role object.
type RoleRelationships struct {
	Permissions *RelationshipToPermissions `json:"permissions,omitempty"`
	Users       *RelationshipToUsers       `json:"users,omitempty"`
}

// NewRoleRelationships instantiates a new RoleRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleRelationships() *RoleRelationships {
	this := RoleRelationships{}
	return &this
}

// NewRoleRelationshipsWithDefaults instantiates a new RoleRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleRelationshipsWithDefaults() *RoleRelationships {
	this := RoleRelationships{}
	return &this
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *RoleRelationships) GetPermissions() RelationshipToPermissions {
	if o == nil || o.Permissions == nil {
		var ret RelationshipToPermissions
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleRelationships) GetPermissionsOk() (*RelationshipToPermissions, bool) {
	if o == nil || o.Permissions == nil {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *RoleRelationships) HasPermissions() bool {
	if o != nil && o.Permissions != nil {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given RelationshipToPermissions and assigns it to the Permissions field.
func (o *RoleRelationships) SetPermissions(v RelationshipToPermissions) {
	o.Permissions = &v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *RoleRelationships) GetUsers() RelationshipToUsers {
	if o == nil || o.Users == nil {
		var ret RelationshipToUsers
		return ret
	}
	return *o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleRelationships) GetUsersOk() (*RelationshipToUsers, bool) {
	if o == nil || o.Users == nil {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *RoleRelationships) HasUsers() bool {
	if o != nil && o.Users != nil {
		return true
	}

	return false
}

// SetUsers gets a reference to the given RelationshipToUsers and assigns it to the Users field.
func (o *RoleRelationships) SetUsers(v RelationshipToUsers) {
	o.Users = &v
}

func (o RoleRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Permissions != nil {
		toSerialize["permissions"] = o.Permissions
	}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}
	return json.Marshal(toSerialize)
}

type NullableRoleRelationships struct {
	value *RoleRelationships
	isSet bool
}

func (v NullableRoleRelationships) Get() *RoleRelationships {
	return v.value
}

func (v *NullableRoleRelationships) Set(val *RoleRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleRelationships(val *RoleRelationships) *NullableRoleRelationships {
	return &NullableRoleRelationships{value: val, isSet: true}
}

func (v NullableRoleRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
