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

// OrgListResponse TODO.
type OrgListResponse struct {
	Orgs *[]Org `json:"orgs,omitempty"`
}

// NewOrgListResponse instantiates a new OrgListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgListResponse() *OrgListResponse {
	this := OrgListResponse{}
	return &this
}

// NewOrgListResponseWithDefaults instantiates a new OrgListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgListResponseWithDefaults() *OrgListResponse {
	this := OrgListResponse{}
	return &this
}

// GetOrgs returns the Orgs field value if set, zero value otherwise.
func (o *OrgListResponse) GetOrgs() []Org {
	if o == nil || o.Orgs == nil {
		var ret []Org
		return ret
	}
	return *o.Orgs
}

// GetOrgsOk returns a tuple with the Orgs field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *OrgListResponse) GetOrgsOk() ([]Org, bool) {
	if o == nil || o.Orgs == nil {
		var ret []Org
		return ret, false
	}
	return *o.Orgs, true
}

// HasOrgs returns a boolean if a field has been set.
func (o *OrgListResponse) HasOrgs() bool {
	if o != nil && o.Orgs != nil {
		return true
	}

	return false
}

// SetOrgs gets a reference to the given []Org and assigns it to the Orgs field.
func (o *OrgListResponse) SetOrgs(v []Org) {
	o.Orgs = &v
}

func (o OrgListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Orgs != nil {
		toSerialize["orgs"] = o.Orgs
	}
	return json.Marshal(toSerialize)
}

type NullableOrgListResponse struct {
	value *OrgListResponse
	isSet bool
}

func (v NullableOrgListResponse) Get() *OrgListResponse {
	return v.value
}

func (v NullableOrgListResponse) Set(val *OrgListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgListResponse) IsSet() bool {
	return v.isSet
}

func (v NullableOrgListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgListResponse(val *OrgListResponse) *NullableOrgListResponse {
	return &NullableOrgListResponse{value: val, isSet: true}
}

func (v NullableOrgListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
