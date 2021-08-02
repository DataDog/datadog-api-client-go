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

// APIKeyCreateData Object used to create an API key.
type APIKeyCreateData struct {
	Attributes APIKeyCreateAttributes `json:"attributes"`
	Type       APIKeysType            `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:-`
}

// NewAPIKeyCreateData instantiates a new APIKeyCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAPIKeyCreateData(attributes APIKeyCreateAttributes, type_ APIKeysType) *APIKeyCreateData {
	this := APIKeyCreateData{}
	this.Attributes = attributes
	this.Type = type_
	return &this
}

// NewAPIKeyCreateDataWithDefaults instantiates a new APIKeyCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAPIKeyCreateDataWithDefaults() *APIKeyCreateData {
	this := APIKeyCreateData{}
	var type_ APIKeysType = APIKEYSTYPE_API_KEYS
	this.Type = type_
	return &this
}

// GetAttributes returns the Attributes field value
func (o *APIKeyCreateData) GetAttributes() APIKeyCreateAttributes {
	if o == nil {
		var ret APIKeyCreateAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *APIKeyCreateData) GetAttributesOk() (*APIKeyCreateAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *APIKeyCreateData) SetAttributes(v APIKeyCreateAttributes) {
	o.Attributes = v
}

// GetType returns the Type field value
func (o *APIKeyCreateData) GetType() APIKeysType {
	if o == nil {
		var ret APIKeysType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *APIKeyCreateData) GetTypeOk() (*APIKeysType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *APIKeyCreateData) SetType(v APIKeysType) {
	o.Type = v
}

func (o APIKeyCreateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

func (o *APIKeyCreateData) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	required := struct {
		Attributes *APIKeyCreateAttributes `json:"attributes"`
		Type       *APIKeysType            `json:"type"`
	}{}
	all := struct {
		Attributes APIKeyCreateAttributes `json:"attributes"`
		Type       APIKeysType            `json:"type"`
	}{}
	err = json.Unmarshal(bytes, &required)
	if err != nil {
		return err
	}
	if required.Attributes == nil {
		return fmt.Errorf("Required field attributes missing")
	}
	if required.Type == nil {
		return fmt.Errorf("Required field type missing")
	}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	if v := all.Type; !v.IsValid() {
		errr := json.Unmarshal(bytes, &raw)
		if errr != nil {
			return errr
		}
		o.UnparsedObject = raw
		return nil
	}
	o.Attributes = all.Attributes
	o.Type = all.Type
	return nil
}

type NullableAPIKeyCreateData struct {
	value *APIKeyCreateData
	isSet bool
}

func (v NullableAPIKeyCreateData) Get() *APIKeyCreateData {
	return v.value
}

func (v *NullableAPIKeyCreateData) Set(val *APIKeyCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableAPIKeyCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableAPIKeyCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAPIKeyCreateData(val *APIKeyCreateData) *NullableAPIKeyCreateData {
	return &NullableAPIKeyCreateData{value: val, isSet: true}
}

func (v NullableAPIKeyCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAPIKeyCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
