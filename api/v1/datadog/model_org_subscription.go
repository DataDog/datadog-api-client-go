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

// OrgSubscription A JSON array of subscription type. Types available are `trial`, `free`, and `pro`.
type OrgSubscription struct {
	Type *string `json:"type,omitempty"`
}

// NewOrgSubscription instantiates a new OrgSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgSubscription() *OrgSubscription {
	this := OrgSubscription{}
	return &this
}

// NewOrgSubscriptionWithDefaults instantiates a new OrgSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgSubscriptionWithDefaults() *OrgSubscription {
	this := OrgSubscription{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *OrgSubscription) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSubscription) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *OrgSubscription) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *OrgSubscription) SetType(v string) {
	o.Type = &v
}

func (o OrgSubscription) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableOrgSubscription struct {
	value *OrgSubscription
	isSet bool
}

func (v NullableOrgSubscription) Get() *OrgSubscription {
	return v.value
}

func (v *NullableOrgSubscription) Set(val *OrgSubscription) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgSubscription(val *OrgSubscription) *NullableOrgSubscription {
	return &NullableOrgSubscription{value: val, isSet: true}
}

func (v NullableOrgSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
