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

// Permissions Payload with API-returned permissions.
type Permissions struct {
	// Array of permissions.
	Data *[]Permission `json:"data,omitempty"`
}

// NewPermissions instantiates a new Permissions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermissions() *Permissions {
	this := Permissions{}
	return &this
}

// NewPermissionsWithDefaults instantiates a new Permissions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionsWithDefaults() *Permissions {
	this := Permissions{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *Permissions) GetData() []Permission {
	if o == nil || o.Data == nil {
		var ret []Permission
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Permissions) GetDataOk() (*[]Permission, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *Permissions) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []Permission and assigns it to the Data field.
func (o *Permissions) SetData(v []Permission) {
	o.Data = &v
}

func (o Permissions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePermissions struct {
	value *Permissions
	isSet bool
}

func (v NullablePermissions) Get() *Permissions {
	return v.value
}

func (v *NullablePermissions) Set(val *Permissions) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissions) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissions(val *Permissions) *NullablePermissions {
	return &NullablePermissions{value: val, isSet: true}
}

func (v NullablePermissions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
