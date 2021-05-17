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

// APIKeyCreateAttributes Attributes used to create an API Key.
type APIKeyCreateAttributes struct {
	// Name of the API key.
	Name string `json:"name"`
}

// NewAPIKeyCreateAttributes instantiates a new APIKeyCreateAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAPIKeyCreateAttributes(name string) *APIKeyCreateAttributes {
	this := APIKeyCreateAttributes{}
	this.Name = name
	return &this
}

// NewAPIKeyCreateAttributesWithDefaults instantiates a new APIKeyCreateAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAPIKeyCreateAttributesWithDefaults() *APIKeyCreateAttributes {
	this := APIKeyCreateAttributes{}
	return &this
}

// GetName returns the Name field value
func (o *APIKeyCreateAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *APIKeyCreateAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *APIKeyCreateAttributes) SetName(v string) {
	o.Name = v
}

func (o APIKeyCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

func (o *APIKeyCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	required := struct {
		Name *string `json:"name"`
	}{}
	all := struct {
		Name string `json:"name"}`
	}{}
	err = json.Unmarshal(bytes, &required)
	if err != nil {
		return err
	}
	if required.Name == nil {
		return fmt.Errorf("Required field name missing")
	}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		return err
	}
	o.Name = all.Name
	return nil
}

type NullableAPIKeyCreateAttributes struct {
	value *APIKeyCreateAttributes
	isSet bool
}

func (v NullableAPIKeyCreateAttributes) Get() *APIKeyCreateAttributes {
	return v.value
}

func (v *NullableAPIKeyCreateAttributes) Set(val *APIKeyCreateAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAPIKeyCreateAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAPIKeyCreateAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAPIKeyCreateAttributes(val *APIKeyCreateAttributes) *NullableAPIKeyCreateAttributes {
	return &NullableAPIKeyCreateAttributes{value: val, isSet: true}
}

func (v NullableAPIKeyCreateAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAPIKeyCreateAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
