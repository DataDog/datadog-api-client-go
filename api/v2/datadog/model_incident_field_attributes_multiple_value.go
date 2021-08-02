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

// IncidentFieldAttributesMultipleValue A field with potentially multiple values selected.
type IncidentFieldAttributesMultipleValue struct {
	Type *IncidentFieldAttributesValueType `json:"type,omitempty"`
	// The multiple values selected for this field.
	Value *[]string `json:"value,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:-`
}

// NewIncidentFieldAttributesMultipleValue instantiates a new IncidentFieldAttributesMultipleValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncidentFieldAttributesMultipleValue() *IncidentFieldAttributesMultipleValue {
	this := IncidentFieldAttributesMultipleValue{}
	var type_ IncidentFieldAttributesValueType = INCIDENTFIELDATTRIBUTESVALUETYPE_MULTISELECT
	this.Type = &type_
	return &this
}

// NewIncidentFieldAttributesMultipleValueWithDefaults instantiates a new IncidentFieldAttributesMultipleValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncidentFieldAttributesMultipleValueWithDefaults() *IncidentFieldAttributesMultipleValue {
	this := IncidentFieldAttributesMultipleValue{}
	var type_ IncidentFieldAttributesValueType = INCIDENTFIELDATTRIBUTESVALUETYPE_MULTISELECT
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IncidentFieldAttributesMultipleValue) GetType() IncidentFieldAttributesValueType {
	if o == nil || o.Type == nil {
		var ret IncidentFieldAttributesValueType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentFieldAttributesMultipleValue) GetTypeOk() (*IncidentFieldAttributesValueType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IncidentFieldAttributesMultipleValue) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given IncidentFieldAttributesValueType and assigns it to the Type field.
func (o *IncidentFieldAttributesMultipleValue) SetType(v IncidentFieldAttributesValueType) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *IncidentFieldAttributesMultipleValue) GetValue() []string {
	if o == nil || o.Value == nil {
		var ret []string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentFieldAttributesMultipleValue) GetValueOk() (*[]string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *IncidentFieldAttributesMultipleValue) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []string and assigns it to the Value field.
func (o *IncidentFieldAttributesMultipleValue) SetValue(v []string) {
	o.Value = &v
}

func (o IncidentFieldAttributesMultipleValue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

func (o *IncidentFieldAttributesMultipleValue) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Type  *IncidentFieldAttributesValueType `json:"type,omitempty"`
		Value *[]string                         `json:"value,omitempty"`
	}{}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	if v := all.Type; v != nil && !v.IsValid() {
		errr := json.Unmarshal(bytes, &raw)
		if errr != nil {
			return errr
		}
		o.UnparsedObject = raw
		return nil
	}
	o.Type = all.Type
	o.Value = all.Value
	return nil
}

type NullableIncidentFieldAttributesMultipleValue struct {
	value *IncidentFieldAttributesMultipleValue
	isSet bool
}

func (v NullableIncidentFieldAttributesMultipleValue) Get() *IncidentFieldAttributesMultipleValue {
	return v.value
}

func (v *NullableIncidentFieldAttributesMultipleValue) Set(val *IncidentFieldAttributesMultipleValue) {
	v.value = val
	v.isSet = true
}

func (v NullableIncidentFieldAttributesMultipleValue) IsSet() bool {
	return v.isSet
}

func (v *NullableIncidentFieldAttributesMultipleValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncidentFieldAttributesMultipleValue(val *IncidentFieldAttributesMultipleValue) *NullableIncidentFieldAttributesMultipleValue {
	return &NullableIncidentFieldAttributesMultipleValue{value: val, isSet: true}
}

func (v NullableIncidentFieldAttributesMultipleValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncidentFieldAttributesMultipleValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
