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

// SyntheticsTestRequestBasicAuth struct for SyntheticsTestRequestBasicAuth
type SyntheticsTestRequestBasicAuth struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

// NewSyntheticsTestRequestBasicAuth instantiates a new SyntheticsTestRequestBasicAuth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyntheticsTestRequestBasicAuth(password string, username string) *SyntheticsTestRequestBasicAuth {
	this := SyntheticsTestRequestBasicAuth{}
	this.Password = password
	this.Username = username
	return &this
}

// NewSyntheticsTestRequestBasicAuthWithDefaults instantiates a new SyntheticsTestRequestBasicAuth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyntheticsTestRequestBasicAuthWithDefaults() *SyntheticsTestRequestBasicAuth {
	this := SyntheticsTestRequestBasicAuth{}
	return &this
}

// GetPassword returns the Password field value
func (o *SyntheticsTestRequestBasicAuth) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *SyntheticsTestRequestBasicAuth) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *SyntheticsTestRequestBasicAuth) SetPassword(v string) {
	o.Password = v
}

// GetUsername returns the Username field value
func (o *SyntheticsTestRequestBasicAuth) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *SyntheticsTestRequestBasicAuth) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *SyntheticsTestRequestBasicAuth) SetUsername(v string) {
	o.Username = v
}

func (o SyntheticsTestRequestBasicAuth) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["password"] = o.Password
	}
	if true {
		toSerialize["username"] = o.Username
	}
	return json.Marshal(toSerialize)
}

type NullableSyntheticsTestRequestBasicAuth struct {
	value *SyntheticsTestRequestBasicAuth
	isSet bool
}

func (v NullableSyntheticsTestRequestBasicAuth) Get() *SyntheticsTestRequestBasicAuth {
	return v.value
}

func (v *NullableSyntheticsTestRequestBasicAuth) Set(val *SyntheticsTestRequestBasicAuth) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticsTestRequestBasicAuth) IsSet() bool {
	return v.isSet
}

func (v *NullableSyntheticsTestRequestBasicAuth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticsTestRequestBasicAuth(val *SyntheticsTestRequestBasicAuth) *NullableSyntheticsTestRequestBasicAuth {
	return &NullableSyntheticsTestRequestBasicAuth{value: val, isSet: true}
}

func (v NullableSyntheticsTestRequestBasicAuth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticsTestRequestBasicAuth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
