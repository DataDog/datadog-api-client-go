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

// FullApplicationKey Datadog application key.
type FullApplicationKey struct {
	Attributes *FullApplicationKeyAttributes `json:"attributes,omitempty"`
	// ID of the application key.
	Id            *string                      `json:"id,omitempty"`
	Relationships *ApplicationKeyRelationships `json:"relationships,omitempty"`
	Type          *ApplicationKeysType         `json:"type,omitempty"`
}

// NewFullApplicationKey instantiates a new FullApplicationKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFullApplicationKey() *FullApplicationKey {
	this := FullApplicationKey{}
	var type_ ApplicationKeysType = "application_keys"
	this.Type = &type_
	return &this
}

// NewFullApplicationKeyWithDefaults instantiates a new FullApplicationKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFullApplicationKeyWithDefaults() *FullApplicationKey {
	this := FullApplicationKey{}
	var type_ ApplicationKeysType = "application_keys"
	this.Type = &type_
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *FullApplicationKey) GetAttributes() FullApplicationKeyAttributes {
	if o == nil || o.Attributes == nil {
		var ret FullApplicationKeyAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullApplicationKey) GetAttributesOk() (*FullApplicationKeyAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *FullApplicationKey) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given FullApplicationKeyAttributes and assigns it to the Attributes field.
func (o *FullApplicationKey) SetAttributes(v FullApplicationKeyAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FullApplicationKey) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullApplicationKey) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FullApplicationKey) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *FullApplicationKey) SetId(v string) {
	o.Id = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *FullApplicationKey) GetRelationships() ApplicationKeyRelationships {
	if o == nil || o.Relationships == nil {
		var ret ApplicationKeyRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullApplicationKey) GetRelationshipsOk() (*ApplicationKeyRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *FullApplicationKey) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given ApplicationKeyRelationships and assigns it to the Relationships field.
func (o *FullApplicationKey) SetRelationships(v ApplicationKeyRelationships) {
	o.Relationships = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FullApplicationKey) GetType() ApplicationKeysType {
	if o == nil || o.Type == nil {
		var ret ApplicationKeysType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullApplicationKey) GetTypeOk() (*ApplicationKeysType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FullApplicationKey) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given ApplicationKeysType and assigns it to the Type field.
func (o *FullApplicationKey) SetType(v ApplicationKeysType) {
	o.Type = &v
}

func (o FullApplicationKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableFullApplicationKey struct {
	value *FullApplicationKey
	isSet bool
}

func (v NullableFullApplicationKey) Get() *FullApplicationKey {
	return v.value
}

func (v *NullableFullApplicationKey) Set(val *FullApplicationKey) {
	v.value = val
	v.isSet = true
}

func (v NullableFullApplicationKey) IsSet() bool {
	return v.isSet
}

func (v *NullableFullApplicationKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFullApplicationKey(val *FullApplicationKey) *NullableFullApplicationKey {
	return &NullableFullApplicationKey{value: val, isSet: true}
}

func (v NullableFullApplicationKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFullApplicationKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
