/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
	"fmt"
)

// RelationshipToOrganizationData Relationship to organization object.
type RelationshipToOrganizationData struct {
	// ID of the organization.
	Id   string            `json:"id"`
	Type OrganizationsType `json:"type"`
}

// NewRelationshipToOrganizationData instantiates a new RelationshipToOrganizationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipToOrganizationData(id string, type_ OrganizationsType) *RelationshipToOrganizationData {
	this := RelationshipToOrganizationData{}
	this.Id = id
	this.Type = type_
	return &this
}

// NewRelationshipToOrganizationDataWithDefaults instantiates a new RelationshipToOrganizationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipToOrganizationDataWithDefaults() *RelationshipToOrganizationData {
	this := RelationshipToOrganizationData{}
	var type_ OrganizationsType = ORGANIZATIONSTYPE_ORGS
	this.Type = type_
	return &this
}

// GetId returns the Id field value
func (o *RelationshipToOrganizationData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RelationshipToOrganizationData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RelationshipToOrganizationData) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *RelationshipToOrganizationData) GetType() OrganizationsType {
	if o == nil {
		var ret OrganizationsType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RelationshipToOrganizationData) GetTypeOk() (*OrganizationsType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RelationshipToOrganizationData) SetType(v OrganizationsType) {
	o.Type = v
}

func (o RelationshipToOrganizationData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

func (o *RelationshipToOrganizationData) UnmarshalJSON(bytes []byte) (err error) {
	required := struct {
		Id   *string            `json:"id"`
		Type *OrganizationsType `json:"type"`
	}{}
	all := struct {
		Id   string            `json:"id"}`
		Type OrganizationsType `json:"type"}`
	}{}
	err = json.Unmarshal(bytes, &required)
	if err != nil {
		return err
	}
	if required.Id == nil {
		return fmt.Errorf("Required field id missing")
	}
	if required.Type == nil {
		return fmt.Errorf("Required field type missing")
	}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		return err
	}
	o.Id = all.Id
	o.Type = all.Type
	return nil
}

type NullableRelationshipToOrganizationData struct {
	value *RelationshipToOrganizationData
	isSet bool
}

func (v NullableRelationshipToOrganizationData) Get() *RelationshipToOrganizationData {
	return v.value
}

func (v *NullableRelationshipToOrganizationData) Set(val *RelationshipToOrganizationData) {
	v.value = val
	v.isSet = true
}

func (v NullableRelationshipToOrganizationData) IsSet() bool {
	return v.isSet
}

func (v *NullableRelationshipToOrganizationData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRelationshipToOrganizationData(val *RelationshipToOrganizationData) *NullableRelationshipToOrganizationData {
	return &NullableRelationshipToOrganizationData{value: val, isSet: true}
}

func (v NullableRelationshipToOrganizationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRelationshipToOrganizationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
