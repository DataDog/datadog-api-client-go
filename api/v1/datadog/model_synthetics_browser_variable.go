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

// SyntheticsBrowserVariable struct for SyntheticsBrowserVariable
type SyntheticsBrowserVariable struct {
	Example *string                       `json:"example,omitempty"`
	Id      *string                       `json:"id,omitempty"`
	Name    string                        `json:"name"`
	Pattern *string                       `json:"pattern,omitempty"`
	Type    SyntheticsBrowserVariableType `json:"type"`
}

// NewSyntheticsBrowserVariable instantiates a new SyntheticsBrowserVariable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyntheticsBrowserVariable(name string, type_ SyntheticsBrowserVariableType) *SyntheticsBrowserVariable {
	this := SyntheticsBrowserVariable{}
	this.Name = name
	this.Type = type_
	return &this
}

// NewSyntheticsBrowserVariableWithDefaults instantiates a new SyntheticsBrowserVariable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyntheticsBrowserVariableWithDefaults() *SyntheticsBrowserVariable {
	this := SyntheticsBrowserVariable{}
	return &this
}

// GetExample returns the Example field value if set, zero value otherwise.
func (o *SyntheticsBrowserVariable) GetExample() string {
	if o == nil || o.Example == nil {
		var ret string
		return ret
	}
	return *o.Example
}

// GetExampleOk returns a tuple with the Example field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsBrowserVariable) GetExampleOk() (string, bool) {
	if o == nil || o.Example == nil {
		var ret string
		return ret, false
	}
	return *o.Example, true
}

// HasExample returns a boolean if a field has been set.
func (o *SyntheticsBrowserVariable) HasExample() bool {
	if o != nil && o.Example != nil {
		return true
	}

	return false
}

// SetExample gets a reference to the given string and assigns it to the Example field.
func (o *SyntheticsBrowserVariable) SetExample(v string) {
	o.Example = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SyntheticsBrowserVariable) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsBrowserVariable) GetIdOk() (string, bool) {
	if o == nil || o.Id == nil {
		var ret string
		return ret, false
	}
	return *o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SyntheticsBrowserVariable) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SyntheticsBrowserVariable) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *SyntheticsBrowserVariable) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// SetName sets field value
func (o *SyntheticsBrowserVariable) SetName(v string) {
	o.Name = v
}

// GetPattern returns the Pattern field value if set, zero value otherwise.
func (o *SyntheticsBrowserVariable) GetPattern() string {
	if o == nil || o.Pattern == nil {
		var ret string
		return ret
	}
	return *o.Pattern
}

// GetPatternOk returns a tuple with the Pattern field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsBrowserVariable) GetPatternOk() (string, bool) {
	if o == nil || o.Pattern == nil {
		var ret string
		return ret, false
	}
	return *o.Pattern, true
}

// HasPattern returns a boolean if a field has been set.
func (o *SyntheticsBrowserVariable) HasPattern() bool {
	if o != nil && o.Pattern != nil {
		return true
	}

	return false
}

// SetPattern gets a reference to the given string and assigns it to the Pattern field.
func (o *SyntheticsBrowserVariable) SetPattern(v string) {
	o.Pattern = &v
}

// GetType returns the Type field value
func (o *SyntheticsBrowserVariable) GetType() SyntheticsBrowserVariableType {
	if o == nil {
		var ret SyntheticsBrowserVariableType
		return ret
	}

	return o.Type
}

// SetType sets field value
func (o *SyntheticsBrowserVariable) SetType(v SyntheticsBrowserVariableType) {
	o.Type = v
}

func (o SyntheticsBrowserVariable) MarshalJSON() ([]byte, error) {
	//TODO: serialize parents?
	toSerialize := map[string]interface{}{}
	if o.Example != nil {
		toSerialize["example"] = o.Example
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Pattern != nil {
		toSerialize["pattern"] = o.Pattern
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableSyntheticsBrowserVariable struct {
	value *SyntheticsBrowserVariable
	isSet bool
}

func (v NullableSyntheticsBrowserVariable) Get() *SyntheticsBrowserVariable {
	return v.value
}

func (v NullableSyntheticsBrowserVariable) Set(val *SyntheticsBrowserVariable) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticsBrowserVariable) IsSet() bool {
	return v.isSet
}

func (v NullableSyntheticsBrowserVariable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticsBrowserVariable(val *SyntheticsBrowserVariable) *NullableSyntheticsBrowserVariable {
	return &NullableSyntheticsBrowserVariable{value: val, isSet: true}
}

func (v NullableSyntheticsBrowserVariable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticsBrowserVariable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
