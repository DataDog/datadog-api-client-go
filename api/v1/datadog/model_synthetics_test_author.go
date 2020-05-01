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

// SyntheticsTestAuthor Object describing the author or a Synthetic test.
type SyntheticsTestAuthor struct {
	// Email of the author.
	Email *string `json:"email,omitempty"`
	// Handle of the author.
	Handle *string `json:"handle,omitempty"`
	// Unique ID of the author.
	Id *int64 `json:"id,omitempty"`
	// Name of the author.
	Name *string `json:"name,omitempty"`
}

// NewSyntheticsTestAuthor instantiates a new SyntheticsTestAuthor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyntheticsTestAuthor() *SyntheticsTestAuthor {
	this := SyntheticsTestAuthor{}
	return &this
}

// NewSyntheticsTestAuthorWithDefaults instantiates a new SyntheticsTestAuthor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyntheticsTestAuthorWithDefaults() *SyntheticsTestAuthor {
	this := SyntheticsTestAuthor{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *SyntheticsTestAuthor) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestAuthor) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *SyntheticsTestAuthor) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *SyntheticsTestAuthor) SetEmail(v string) {
	o.Email = &v
}

// GetHandle returns the Handle field value if set, zero value otherwise.
func (o *SyntheticsTestAuthor) GetHandle() string {
	if o == nil || o.Handle == nil {
		var ret string
		return ret
	}
	return *o.Handle
}

// GetHandleOk returns a tuple with the Handle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestAuthor) GetHandleOk() (*string, bool) {
	if o == nil || o.Handle == nil {
		return nil, false
	}
	return o.Handle, true
}

// HasHandle returns a boolean if a field has been set.
func (o *SyntheticsTestAuthor) HasHandle() bool {
	if o != nil && o.Handle != nil {
		return true
	}

	return false
}

// SetHandle gets a reference to the given string and assigns it to the Handle field.
func (o *SyntheticsTestAuthor) SetHandle(v string) {
	o.Handle = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SyntheticsTestAuthor) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestAuthor) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SyntheticsTestAuthor) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *SyntheticsTestAuthor) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SyntheticsTestAuthor) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestAuthor) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SyntheticsTestAuthor) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SyntheticsTestAuthor) SetName(v string) {
	o.Name = &v
}

func (o SyntheticsTestAuthor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.Handle != nil {
		toSerialize["handle"] = o.Handle
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableSyntheticsTestAuthor struct {
	value *SyntheticsTestAuthor
	isSet bool
}

func (v NullableSyntheticsTestAuthor) Get() *SyntheticsTestAuthor {
	return v.value
}

func (v *NullableSyntheticsTestAuthor) Set(val *SyntheticsTestAuthor) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticsTestAuthor) IsSet() bool {
	return v.isSet
}

func (v *NullableSyntheticsTestAuthor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticsTestAuthor(val *SyntheticsTestAuthor) *NullableSyntheticsTestAuthor {
	return &NullableSyntheticsTestAuthor{value: val, isSet: true}
}

func (v NullableSyntheticsTestAuthor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticsTestAuthor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
