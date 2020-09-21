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

// UserRelationship A relationship reference for users.
type UserRelationship struct {
	Data *UserRelationshipData `json:"data,omitempty"`
}

// NewUserRelationship instantiates a new UserRelationship object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserRelationship() *UserRelationship {
	this := UserRelationship{}
	return &this
}

// NewUserRelationshipWithDefaults instantiates a new UserRelationship object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserRelationshipWithDefaults() *UserRelationship {
	this := UserRelationship{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *UserRelationship) GetData() UserRelationshipData {
	if o == nil || o.Data == nil {
		var ret UserRelationshipData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRelationship) GetDataOk() (*UserRelationshipData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *UserRelationship) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given UserRelationshipData and assigns it to the Data field.
func (o *UserRelationship) SetData(v UserRelationshipData) {
	o.Data = &v
}

func (o UserRelationship) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableUserRelationship struct {
	value *UserRelationship
	isSet bool
}

func (v NullableUserRelationship) Get() *UserRelationship {
	return v.value
}

func (v *NullableUserRelationship) Set(val *UserRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableUserRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableUserRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserRelationship(val *UserRelationship) *NullableUserRelationship {
	return &NullableUserRelationship{value: val, isSet: true}
}

func (v NullableUserRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
